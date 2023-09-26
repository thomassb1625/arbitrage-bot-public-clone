package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/arbitrage"
	"github.com/jmjac/ethGoTest/dex"
	"github.com/jmjac/ethGoTest/dex/uniswapV2"
	"github.com/jmjac/ethGoTest/flashbots"
	"github.com/jmjac/ethGoTest/helperClient"
	"github.com/jmjac/ethGoTest/secrets"
)

func main() {
	//Setup log files
	genf, err := os.OpenFile("general-moonbeam.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	genLog := log.New(genf, "", 0)

	submitf, err := os.OpenFile("submitted-moonbeam.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	ADDR.Moonbeam(*genLog)
	//subLog := log.New(submitf, "", 0)
	//revf, err := os.OpenFile("reverted-moonbeam.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//revLog := log.New(revf, "", 0)
	//errorf, err := os.OpenFile("error-moonbeam.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//errLog := log.New(errorf, "", 0)
	defer genf.Close()
	defer submitf.Close()
	//defer revf.Close()
	//defer errorf.Close()

	dex.SetLogger(*genLog)

	var provider string
	//Read from env in the future
	url := "https://xyznode.space/api/moonbeam/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"
	mainnet := false
	threads := 256

	flag.StringVar(&provider, "provider", url, "RPC endpoint")
	flag.BoolVar(&mainnet, "mainnet", mainnet, "If true use mainnet, otherwise goerli addreses")
	flag.IntVar(&threads, "threads", 256, "Num of gorutines(default 256)")
	flag.Parse()

	ethClient, err := helperClient.New(provider)
	if err != nil {
		genLog.Println("Coudn't create client,", err)
		return
	}

	// =============================== SETUP BLACKLIST ==========================
	BLACKLIST := arbitrage.InitBlackList("blacklist_moonbeam.json")
	genLog.Printf("Blacklist has %v tokens\n", len(BLACKLIST))
	const BLACKLIST_THRESHOLD = 200
	toBlackList := make(map[common.Address]int)
	toBlackList = arbitrage.AddWhiteListed(toBlackList)
	// ============================= END SETUP BLACKLIST ========================

	TOKEN_DECIMALS := arbitrage.LoadDecimalsFromJson("token_decimals_moonbeam.json")

	// ============================= LOAD UNIV2 POOLS ======================
	genLog.Println(ethClient.BlockNumber(context.Background()))
	stellaSwap := uniswapV2.GetUniswappyPairs(ethClient, ADDR.UNISWAPv2_LOOKUP, ADDR.STELLASWAP_FACTORY, "stellaSwap", BLACKLIST, make(map[common.Address]dex.Token))
	solarflareSwap := uniswapV2.GetUniswappyPairs(ethClient, ADDR.UNISWAPv2_LOOKUP, ADDR.SOLARFLARE_FACTORY, "solarflareSwap", BLACKLIST, make(map[common.Address]dex.Token))
	padSwap := uniswapV2.GetUniswappyPairs(ethClient, ADDR.UNISWAPv2_LOOKUP, ADDR.PADSWAP_FACTORY, "padSwap", BLACKLIST, make(map[common.Address]dex.Token))
	beamSwap := uniswapV2.GetUniswappyPairs(ethClient, ADDR.UNISWAPv2_LOOKUP, ADDR.BEAMSWAP_FACTORY, "beamSwap", BLACKLIST, make(map[common.Address]dex.Token))
	allUniV2 := mergeUniV2Markets(stellaSwap, solarflareSwap, padSwap, beamSwap)

	stella := make([]dex.Pair, 0)
	for _, p := range stellaSwap {
		stella = append(stella, p)
	}
	solar := make([]dex.Pair, 0)
	for _, p := range solarflareSwap {
		solar = append(solar, p)
	}
	pad := make([]dex.Pair, 0)
	for _, p := range padSwap {
		pad = append(pad, p)
	}
	beam := make([]dex.Pair, 0)
	for _, p := range beamSwap {
		beam = append(beam, p)
	}

	// ============================= END LOAD UNIV2 POOLS ======================

	genLog.Println("Generating trading paths")
	now := time.Now()
	targetToken := ADDR.WGLMR
	toBlackList[targetToken] = -1000000000

	tradePaths := arbitrage.GeneratePaths(2, ethClient, targetToken, 5, BLACKLIST, TOKEN_DECIMALS, stella, solar, pad, beam) //BalancerV2Weighted
	genLog.Println("Trading paths: \t", len(tradePaths))
	arbitrage.SaveTokenDecimals(TOKEN_DECIMALS, "token_decimals_moonbeam.json")
	genLog.Printf("Found %v paths\n", len(tradePaths))
	genLog.Println("Took: ", time.Since(now))

	blockNum, _ := ethClient.BlockNumber(context.Background())

	blockNum, _ = ethClient.BlockNumber(context.Background())
	newBlock, _ := ethClient.BlockNumber(context.Background())
	genLog.Println("Syncing with blocktime")
	for newBlock-blockNum < 1 {
		newBlock, _ = ethClient.BlockNumber(context.Background())
		time.Sleep(time.Millisecond * 100)
	}
	genLog.Printf("Found new block at %v", time.Now())

	MIN_AMOUNT_IN := uint256.NewInt(1e16)
	MIN_AMOUNT_IN.Mul(MIN_AMOUNT_IN, uint256.NewInt(1))
	MAX_AMOUNT_IN := uint256.NewInt(10)
	MAX_AMOUNT_IN.Mul(MAX_AMOUNT_IN, uint256.NewInt(1e18))

	genLog.Printf("\tArb token: %v\n", targetToken)
	genLog.Printf("\tMIN_AMOUNT_IN: %v\n", MIN_AMOUNT_IN)
	genLog.Printf("\tMAX_AMOUNT_IN: %v\n", MAX_AMOUNT_IN)

	signer := flashbots.NewSigner(big.NewInt(1284), secrets.PRIVATE_KEY_MOONBEAM, ADDR.BUNDLE_EXECUTOR, ADDR.WGLMR, ethClient)
	//targetTokenContract, err := abis.NewIERC20(targetToken, ethClient)

	baseFee := big.NewInt(100000000000) //TODO: Later change this
	for i := 0; ; i++ {
		nounce, err := ethClient.NonceAt(context.Background(), signer.Address, nil)
		if err != nil {
			log.Fatal("Coudn't get nounce with err: ", err)
		}

		now := time.Now()
		uniswapV2.UpdateReserves(ethClient, ADDR.UNISWAPv2_LOOKUP, allUniV2)
		genLog.Println("Updating reserves took: ", time.Since(now))

		now = time.Now()

		arbs := arbitrage.LookForArb(1, targetToken, MIN_AMOUNT_IN, MAX_AMOUNT_IN, uint256.NewInt(105), uint256.NewInt(100), tradePaths)
		genLog.Println("Arbitrage search took: ", time.Since(now))
		now = time.Now()
		reverts := 0
		notProfitable := 0
		for i, arb := range arbs {
			if i > 40 {
				break
			}
			_, targets, payloads := arb.MoonbeamPrepareTx(targetToken, ADDR.BUNDLE_EXECUTOR)

			txData, err := signer.MoonbeamPrepareDirectArbTx(arb.AmountIn.ToBig(), big.NewInt(0), targets, payloads)
			if err != nil {
				fmt.Println("Error when creating tx data")
				continue
			}

			estimatedGas, err := signer.EstimateGasForTx(big.NewInt(0), txData)
			if err != nil {
				for _, p := range arb.Path {
					t0, t1 := p.Tokens()
					toBlackList[t0.Addr]++
					toBlackList[t1.Addr]++
				}
				reverts++
				continue
			}

			cost := new(big.Int).Mul(big.NewInt(int64(estimatedGas*125/100)), baseFee)
			if cost.Cmp(arb.Profit.ToBig()) >= 0 {
				notProfitable++
				continue
			}

			fmt.Println("Submiting!")
			err = signer.SentTxToBundle(txData, big.NewInt(0), nounce, estimatedGas*3/2, baseFee, baseFee)
			if err != nil {
				log.Fatal("Failed to send tx. Err: ", err)
			}
			break
		}
		genLog.Println("Simulations took: ", time.Since(now))
		genLog.Println("Not profitable: ", notProfitable)
		genLog.Println("Reverts: ", reverts)

		if arbitrage.BlackistWasUpdated(toBlackList, BLACKLIST, BLACKLIST_THRESHOLD, *genLog) {
			genLog.Println("Recreating paths after blacklist updates")
			newTradingPaths := make([][]dex.Pair, 0)
			for _, path := range tradePaths {
				blacklisted := false
				for _, p := range path {
					t0, t1 := p.Tokens()
					if BLACKLIST[t0.Addr] || BLACKLIST[t1.Addr] {
						blacklisted = true
						break
					}
				}
				if !blacklisted {
					newTradingPaths = append(newTradingPaths, path)
				}
			}
			tradePaths = newTradingPaths
			genLog.Printf("Have %v paths\n", len(tradePaths))
		}

		genLog.Println("Waiting for new block")
		newBlock, _ := ethClient.BlockNumber(context.Background())
		for newBlock-blockNum < 1 {
			newBlock, _ = ethClient.BlockNumber(context.Background())
			time.Sleep(time.Millisecond * 100)
		}
		blockNum = newBlock
		genLog.Printf("Got new block: %v\n\n", blockNum)
	}

}

func mergeUniV2Markets(markets ...[]*uniswapV2.Pair) []*uniswapV2.Pair {
	allUniV2 := make([]*uniswapV2.Pair, 0)
	for _, m := range markets {
		for _, p := range m {
			allUniV2 = append(allUniV2, p)
		}
	}
	return allUniV2
}
