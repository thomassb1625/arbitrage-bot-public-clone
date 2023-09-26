package arbitrage

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/flashbots"
	"github.com/jmjac/ethGoTest/helperClient"
)

type ArbBot struct {
	//Log files
	revertedLog *log.Logger
	genLog      *log.Logger
	submitedLog *log.Logger
	noProfitLog *log.Logger
	errLog      *log.Logger

	//Sending tx
	mevRelays    []flashbots.MevRelay
	signerPureV2 flashbots.Signer
	signerMixed  flashbots.Signer
	ethclient    *helperClient.Client

	//Block specific
	baseFee      *big.Int
	blockNum     uint64
	pureV2Nounce uint64
	mixedNounce  uint64

	//Bribe
	bribeNumerator   *big.Int
	bribeDenominator *big.Int

	//Counters
	SimTotal       int32
	SimReverts     int32
	flashbotErrors int32
	Submitted      int32
	NotProfitable  int32
	AccessListFail int32

	//Blacklist
	BLACKLIST      map[common.Address]bool
	TokenReverts   map[common.Address]int
	mutexBlackList sync.RWMutex

	//Channels
	successSubmit chan SimResult
	successSim    chan SimResult
}

type SimResult struct {
	success bool
	profit  *uint256.Int
}

func NewArbBot(revertedLog, generalLog, submittedLog, errLog, noProfitLog *log.Logger, signerPureV2, signerMixed flashbots.Signer, client *helperClient.Client, mevRelays []flashbots.MevRelay, BLACKLIST map[common.Address]bool, tokenReverts map[common.Address]int) *ArbBot {
	bot := &ArbBot{signerPureV2: signerPureV2, signerMixed: signerMixed, mevRelays: mevRelays, baseFee: big.NewInt(0), bribeNumerator: big.NewInt(0), bribeDenominator: big.NewInt(0)}
	bot.ethclient = client
	bot.revertedLog = revertedLog
	bot.genLog = generalLog
	bot.submitedLog = submittedLog
	bot.noProfitLog = noProfitLog
	bot.errLog = errLog
	bot.BLACKLIST = BLACKLIST
	bot.TokenReverts = tokenReverts
	bot.mutexBlackList = sync.RWMutex{}
	return bot
}

func (bot *ArbBot) UpdateForNewBlock(blockNum, pureV2Nounce, mixedNounce uint64, baseFee, bribeNumerator, bribeDenominator *big.Int) {
	bot.baseFee.Set(baseFee)
	bot.bribeNumerator.Set(bribeNumerator)
	bot.bribeDenominator.Set(bribeDenominator)

	//Update block vars
	bot.blockNum = blockNum
	bot.pureV2Nounce = pureV2Nounce
	bot.mixedNounce = mixedNounce

	//Reset counters
	bot.SimTotal = 0
	bot.SimReverts = 0
	bot.Submitted = 0
	bot.AccessListFail = 0
	bot.flashbotErrors = 0
	bot.NotProfitable = 0
}

func (bot *ArbBot) CheckProfits(minNumArbs, maxNumArbs int, arbs []*Arbitrage) bool {
	startTime := time.Now()
	bot.successSubmit = make(chan SimResult, 10)
	bot.successSim = make(chan SimResult, 10)
	guard := make(chan bool, 10) //Limits to 10 concurrent simulations
	var wg sync.WaitGroup

	submitted := false
	highestProfit := uint256.NewInt(0)
	for i, arbOpp := range arbs {

		select {
		case res := <-bot.successSim:
			if highestProfit.Cmp(res.profit) < 0 {
				highestProfit.Set(res.profit)
			}
		case res := <-bot.successSubmit:
			submitted = res.success
			if highestProfit.Cmp(res.profit) < 0 {
				highestProfit.Set(res.profit)
			}
		default:

		}
		if (i > minNumArbs && submitted) || i > maxNumArbs {
			break
		}

		arb := *arbOpp
		arbNum := i
		wg.Add(1)
		guard <- true
		//Add the gorutine
		go func(arb *Arbitrage, arbNum int) {
			txData, msgValue, onlyV2 := arb.NewPrepareTx()
			if onlyV2 {
				bot.checkArb(arbNum, arb, txData, msgValue, arb.Profit.ToBig(), int(bot.blockNum), bot.signerPureV2, bot.pureV2Nounce)
			} else {
				bot.checkArb(arbNum, arb, txData, msgValue, arb.Profit.ToBig(), int(bot.blockNum), bot.signerMixed, bot.mixedNounce)
			}
			wg.Done()
			<-guard
		}(&arb, arbNum)

		bot.SimTotal++
	}

	wg.Wait()
	close(guard)
	//Make sure we udpate profit
	select {
	case res := <-bot.successSubmit:
		if highestProfit.Cmp(res.profit) < 0 {
			highestProfit.Set(res.profit)
		}
	default:
	}
	close(bot.successSubmit)
	close(bot.successSim)

	bot.genLog.Printf("Highest Simulated Profit: \t%v\n", highestProfit)
	bot.genLog.Printf("Reverted Simulations: \t%v\n", bot.SimReverts)
	bot.genLog.Printf("Acccess list failure : \t%v\n", bot.AccessListFail)
	bot.genLog.Printf("NotProfitable: \t%v\n", bot.NotProfitable)
	bot.genLog.Printf("Flashbots errors: \t%v\n", bot.flashbotErrors)
	bot.genLog.Printf("Submited Tx: \t%v\n", bot.Submitted)
	bot.genLog.Printf("Total Simulations: \t%v\n", bot.SimTotal)
	bot.genLog.Printf("Simulation: \t%v\n", time.Since(startTime))
	return bot.Submitted != 0
}

func (bot *ArbBot) checkArb(arbNum int, arb *Arbitrage, txData []byte, msgValue, profit *big.Int, blockNum int, signer flashbots.Signer, nounce uint64) {
	estimatedGas, err := signer.EstimateGasForTxWithBlock(msgValue, txData, "latest")

	// Simulation fail
	if err != nil {
		msg := "======================================================\n"
		msg += fmt.Sprintf("Block: \t%v\n", blockNum)
		msg += fmt.Sprintf("Profit: \t%v\n", profit)
		msg += formatPath(arb)
		msg += fmt.Sprintf("Simulated using \t%v\n", signer.Address)
		msg += fmt.Sprintf("TxData: \t%v\n", hexutil.Encode(txData))
		msg += fmt.Sprintf("MsgValue: \t%v\n", msgValue)
		bot.revertedLog.Println(msg)
		atomic.AddInt32(&bot.SimReverts, 1)

		//Indicate that token failed simulation
		bot.mutexBlackList.Lock()
		for _, p := range arb.Path {
			t0, t1 := p.Tokens()
			bot.TokenReverts[t0.Addr]++
			bot.TokenReverts[t1.Addr]++
		}
		bot.mutexBlackList.Unlock()
		return
	}

	accessList, accessListGas, err := bot.ethclient.CreateAccessList(signer.Address, ADDR.BUNDLE_EXECUTOR, txData, msgValue)
	if err != nil {
		msg := "======================================================\n"
		msg += fmt.Sprintf("Failed to create access list\n")
		msg += fmt.Sprintf("Block: \t%v\n", blockNum)
		msg += fmt.Sprintf("Profit: \t%v\n", profit)
		msg += formatPath(arb)
		msg += fmt.Sprintf("Simulated using \t%v\n", signer.Address)
		msg += fmt.Sprintf("TxData: \t%v\n", hexutil.Encode(txData))
		msg += fmt.Sprintf("MsgValue: \t%v\n", msgValue)
		bot.revertedLog.Println(msg)
		atomic.AddInt32(&bot.AccessListFail, 1)
		return
	}

	// Token simulated successfuly
	bot.mutexBlackList.Lock()
	for _, p := range arb.Path {
		t0, t1 := p.Tokens()
		bot.TokenReverts[t0.Addr] -= 3000
		bot.TokenReverts[t1.Addr] -= 3000
	}
	bot.mutexBlackList.Unlock()

	tempBaseFee := new(big.Int).Mul(bot.baseFee, big.NewInt(100))
	tempBaseFee.Div(tempBaseFee, big.NewInt(100))
	totalCost := new(big.Int).Mul(big.NewInt(int64(accessListGas)), tempBaseFee)

	//Tx not profitalbe so we don't submit
	if totalCost.Cmp(profit) > 0 {
		msg := "======================================================\n"
		msg += fmt.Sprintf("Estimated Erigon Gas: \t%v\n", estimatedGas)
		msg += fmt.Sprintf("Estimated Access List Gas: \t%v\n", accessListGas)
		msg += fmt.Sprintf("BaseFee block: \t%v\n", bot.baseFee)
		msg += fmt.Sprintf("Block: \t%v\n", blockNum)
		msg += fmt.Sprintf("Profit: \t%v\n", profit)
		msg += formatPath(arb)
		msg += fmt.Sprintf("Simulated using \t%v\n", signer.Address)
		msg += fmt.Sprintf("TxData: \t%v\n", hexutil.Encode(txData))
		msg += fmt.Sprintf("MsgValue: \t%v\n", msgValue)
		bot.noProfitLog.Println(msg)
		atomic.AddInt32(&bot.NotProfitable, 1)
		return
	}
	tempBaseFee = new(big.Int).Mul(bot.baseFee, big.NewInt(115))
	tempBaseFee.Div(tempBaseFee, big.NewInt(100))

	bot.successSim <- SimResult{true, arb.Profit} //Indicate that we have successful simulation

	//tx, _ := signer.SignTx(txData, msgValue, nounce, estimatedGas*3/2, tempBaseFee, tempBaseFee)
	//signedRawTx := flashbots.GetRawSignedTx(tx)
	txAccess, _ := signer.SignTxWithAccessList(txData, *accessList, msgValue, nounce, accessListGas*3/2, tempBaseFee, tempBaseFee)
	signedAccessTx := flashbots.GetRawSignedTx(txAccess)

	flashbotsEstimatedGas, err := bot.mevRelays[0].CallBundle(signedAccessTx, fmt.Sprintf("0x%x", blockNum+1), *bot.errLog)
	if err != nil || flashbotsEstimatedGas == 0 {
		bot.errLog.Printf("Failed on callbundleStats with tradeNum %v at block %v with error %v and flashbots gas %v\n", arbNum, blockNum, err, flashbotsEstimatedGas)
		atomic.AddInt32(&bot.flashbotErrors, 1)
		return
	}

	gasLimit := flashbotsEstimatedGas * 1065 / 1000

	gasFeeCap := new(big.Int).Div(profit, big.NewInt(gasLimit))
	gasFeeCap.Mul(gasFeeCap, bot.bribeNumerator)
	gasFeeCap.Div(gasFeeCap, bot.bribeDenominator)
	gasTipCap := new(big.Int).Set(gasFeeCap)

	//Tx not profitalbe so we don't submit
	if gasFeeCap.Cmp(bot.baseFee) < 0 {
		atomic.AddInt32(&bot.NotProfitable, 1)
		return
	}

	bot.successSubmit <- SimResult{true, arb.Profit} //Indicate that we have successful submitted

	txAccess, _ = signer.SignTxWithAccessList(txData, *accessList, msgValue, nounce, uint64(gasLimit), gasFeeCap, gasFeeCap)
	signedAccessTx = flashbots.GetRawSignedTx(txAccess)

	for _, relay := range bot.mevRelays {
		//WARN: Has to be +1. If bigger next block can have higher basefee and we lose money
		_, err := relay.SendBundleSingleTx(signedAccessTx, fmt.Sprintf("0x%x", blockNum+1), *bot.errLog)
		if err != nil {
			bot.errLog.Printf("WARNING RELAY %v SEND FAILED at %v arbNum at %v block with err %v. Retrying\n", relay.Name, arbNum, blockNum, err)
			time.Sleep(time.Millisecond * time.Duration(rand.Float64()) * 150)
			_, err := relay.SendBundleSingleTx(signedAccessTx, fmt.Sprintf("0x%x", blockNum+1), *bot.errLog)
			if err != nil {
				bot.errLog.Printf("SECOND ATTEMPT FAILED RELAY %v SEND FAILED at %v arbNum at %v block with err %v\n", relay.Name, arbNum, blockNum, err)
			}
			//NOTE: This error should be handeled better
		}
	}

	// NOTE: Sprintf is slow so it maybe better to do it all in one go
	msg := "======================================================\n"
	msg += fmt.Sprintf("Submited for Block: \t%v\n", blockNum+1)
	msg += fmt.Sprintf("ArbNum: \t%v\n", arbNum)
	msg += fmt.Sprintf("Erigon Estimated Gas: \t%v\n", estimatedGas)
	msg += fmt.Sprintf("Access List Estimated Gas: \t%v\n", accessListGas)
	msg += fmt.Sprintf("Flashbots Estimated Gas: \t%v\n", flashbotsEstimatedGas)
	msg += fmt.Sprintf("Submited Gas Limit: \t%v\n", gasLimit)
	msg += fmt.Sprintf("GasFeeCap: \t%v\n", gasFeeCap)
	msg += fmt.Sprintf("GasTipCap: \t%v\n", gasTipCap)
	msg += fmt.Sprintf("BribeNum: \t%v\n", bot.bribeNumerator)
	msg += fmt.Sprintf("BribeDen: \t%v\n", bot.bribeDenominator)
	msg += fmt.Sprintf("Profit: \t%v\n", profit)
	msg += fmt.Sprintf("Sender: \t%v\n", signer.Address)
	msg += formatPath(arb)
	msg += fmt.Sprintf("TxData: \t%v\n", hexutil.Encode(txData))
	msg += fmt.Sprintf("MsgValue: \t%v\n", msgValue)
	msg += fmt.Sprintf("SignedTx: \t%v\n", signedAccessTx)
	msg += fmt.Sprintf("TxHash: \t%v\n", txAccess.Hash())
	bot.submitedLog.Println(msg)
	atomic.AddInt32(&bot.Submitted, 1)
}

func formatPath(arb *Arbitrage) string {
	path := ""
	tokenIn := arb.StartToken
	in := arb.AmountIn
	for i, p := range arb.Path {
		tokenOut := p.GetOtherToken(tokenIn).Addr
		out := p.GetTokensOut(tokenIn, tokenOut, in)
		path += fmt.Sprintf("Swap: %v\nExchange: %v\nTokenIn: %v\nAmountIn: %v\nTokenOut: %v\nAmountOut: %v\n", i, p.MarketAddress(), tokenIn, in, tokenOut, out)
		in = out
		tokenIn = tokenOut
	}
	return path
}
