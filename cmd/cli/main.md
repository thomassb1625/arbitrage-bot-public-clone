package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/abis"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/arbitrage"
	"github.com/jmjac/ethGoTest/balancerV2"
	curve "github.com/jmjac/ethGoTest/curve"
	"github.com/jmjac/ethGoTest/flashbots"
	secrets "github.com/jmjac/ethGoTest/secrets"
	"github.com/jmjac/ethGoTest/uniswapV2"
)

func main() {
	var provider string
	//Read from env in the future
	// url := "https://xyznode.space/api/mainnet/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"
	// url := "https://goerli.infura.io/v3/a7e04ad8464640009197f8a83eb0a686"
	// url := "http://127.0.0.1:8545"
	url := "https://xyznode.space/api/forked/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"
	mainnet := true

	flag.StringVar(&provider, "provider", url, "RPC endpoint")
	flag.BoolVar(&mainnet, "mainnet", mainnet, "If true use mainnet, otherwise goerli addreses")
	flag.Parse()

	if provider == "" {
		panic("Provider can't be empty. Use --provider")
	}
	client, err := ethclient.Dial(provider)
	if err != nil {
		panic(err)
	}

	//Setup adreses
	if mainnet {
		ADDR.MainnetAddresses()
	} else {
		ADDR.GoerliAddresses()
	}

	// =============================== SETUP BLACKLIST ==========================
	BLACKLIST := arbitrage.InitBlackList("blacklist.json")
	fmt.Printf("Blacklist has %v tokens\n", len(BLACKLIST))
	const BLACKLIST_THRESHOLD = 50
	blacklistedTokens := 0
	toBlackList := make(map[common.Address]int)
	toBlackList = arbitrage.AddWhiteListed(toBlackList)
	toBlackList[ADDR.WETH] = -100000000
	toBlackList[ADDR.DAI] = -100000000
	// ============================= END SETUP BLACKLIST ========================

	// ============================= SETUP TOKEN_DECIMALS ==================
	TOKEN_DECIMALS := arbitrage.LoadDecimalsFromJson("token_decimals.json")
	// ============================= END SETUP TOKEN_DECIMALS ===================

	// ============================= LOAD DIFFERENT POOLS ======================

	pairsBalancer2Token := balancerV2.Get2TokenWeightedMarket(client, ADDR.BALANCERv2_VAULT, ADDR.BALANCERv2_WEIGHTED2TOKENS_FACTORY, big.NewInt(0), BLACKLIST)
	pairsBalancerWeighted := balancerV2.GetWeightedMarkets(client, ADDR.BALANCERv2_VAULT, ADDR.BALANCERv2_WEIGHTED_FACTORY, big.NewInt(0), BLACKLIST)

	//Covert Weighted Pairs to 2TokenPairs for better compatibility with the market API for now
	allBalancerPairs := make([]*balancerV2.Weighted2TokenPair, 0)
	for _, p := range pairsBalancerWeighted {
		for _, converted := range p.ConvertTo2TokenWeightedPair(ADDR.BALANCERv2_VAULT) {
			allBalancerPairs = append(allBalancerPairs, converted)
		}
	}

	for _, p := range pairsBalancer2Token {
		allBalancerPairs = append(allBalancerPairs, p)
	}

	fmt.Printf("Unique %v trading pairs from BalancerV2Weighted\n", len(allBalancerPairs))

	pairsSushi := uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.SUSHI_FACTORY, "sushi-swap", BLACKLIST)
	pairsUniV2 := uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.UNIv2_FACTORY, "uniswapV2", BLACKLIST)
	pairsShibaSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.SHIBA_FACTORY != common.HexToAddress("") {
		pairsShibaSwap = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.SHIBA_FACTORY, "shiba-swap", BLACKLIST)
	}
	pairsOrion := make([]*uniswapV2.Pair, 0)
	if ADDR.ORION_FACTORY != common.HexToAddress("") {
		pairsOrion = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.ORION_FACTORY, "orion-swap", BLACKLIST)
	}
	pairsSake := make([]*uniswapV2.Pair, 0)
	if ADDR.SAKE_FACTORY != common.HexToAddress("") {
		pairsSake = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.SAKE_FACTORY, "sake-swap", BLACKLIST)
	}
	pairsDXSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.DX_FACTORY != common.HexToAddress("") {
		pairsDXSwap = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.DX_FACTORY, "dx-swap", BLACKLIST)
	}
	pairsCryptDefi := make([]*uniswapV2.Pair, 0)
	if ADDR.CRYPTDEFI_FACTORY != common.HexToAddress("") {
		pairsCryptDefi = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.CRYPTDEFI_FACTORY, "cryptdefi-swap", BLACKLIST)
	}

	pairsCurve := curve.GetCurveMarkets(client, ADDR.CURVE_ADDRESS_PROVIDER, ADDR.CURVE_GETTER, "curve-finance", BLACKLIST)
	fmt.Println(pairsCurve)
	// ============================================= END LOAD DIFFERENT POOLS =======================================================

	// ============================================= CAST EACH POOL TO PAIR =========================================================
	uniV2 := make([]arbitrage.Pair, 0, len(pairsUniV2))
	for _, p := range pairsUniV2 {
		uniV2 = append(uniV2, p)
	}
	sushi := make([]arbitrage.Pair, 0, len(pairsSushi))
	for _, p := range pairsSushi {
		sushi = append(sushi, p)
	}
	shiba := make([]arbitrage.Pair, 0, len(pairsShibaSwap))
	for _, p := range pairsShibaSwap {
		shiba = append(shiba, p)
	}
	orion := make([]arbitrage.Pair, 0, len(pairsOrion))
	for _, p := range pairsOrion {
		orion = append(orion, p)
	}
	sake := make([]arbitrage.Pair, 0, len(pairsSake))
	for _, p := range pairsSake {
		sake = append(sake, p)
	}
	dxSwap := make([]arbitrage.Pair, 0, len(pairsDXSwap))
	for _, p := range pairsDXSwap {
		dxSwap = append(dxSwap, p)
	}
	cryptDefi := make([]arbitrage.Pair, 0, len(pairsCryptDefi))
	for _, p := range pairsCryptDefi {
		cryptDefi = append(cryptDefi, p)
	}
	balancerWeighted := make([]arbitrage.Pair, 0, len(allBalancerPairs))
	for _, p := range allBalancerPairs {
		balancerWeighted = append(balancerWeighted, p)
	}
	// ============================================= END CAST EACH POOL TO PAIR =========================================================

	fmt.Println("Generating trading paths")
	now := time.Now()
	targetToken := ADDR.WETH
	//Shoul be called after blacklist updates

	tradingPaths := arbitrage.GeneratePaths(client, targetToken, 5, BLACKLIST, TOKEN_DECIMALS, sushi, shiba, orion, sake, dxSwap, balancerWeighted, uniV2, cryptDefi)
	arbitrage.SaveTokenDecimals(TOKEN_DECIMALS, "token_decimals.json")
	fmt.Printf("Found %v paths\n", len(tradingPaths))
	fmt.Println("Took: ", time.Since(now))

	blockNum, _ := client.BlockNumber(context.Background())
	newBlock, _ := client.BlockNumber(context.Background())
	fmt.Println("Syncing with blocktime")
	for newBlock-blockNum < 1 {
		newBlock, _ = client.BlockNumber(context.Background())
		time.Sleep(time.Millisecond * 200)
	}
	blockNum = newBlock

	wethContract, err := abis.NewIERC20(ADDR.WETH, client)
	signer := flashbots.NewSigner(big.NewInt(1), secrets.PRIVATE_KEY, ADDR.BUNDLE_EXECUTOR, ADDR.WETH, client) //Has 0.6 ETH on Goerli

	bribeFlat := int64(99911)
	bribeDenominator := big.NewInt(100000)
	nounce, _ := client.NonceAt(context.Background(), signer.Address, nil)
	lastNounce := nounce
	submitedToFlashbots := false

	for i := 0; ; i++ {
		now := time.Now()
		updateReserves(client, pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsBalancerWeighted, pairsBalancer2Token)
		fmt.Println("Updating reserves took ", time.Since(now))
		now = time.Now()
		arbs := arbitrage.LookForArb(targetToken, tradingPaths)

		//TODO: Rework the gas price estimation to use tipcap and basefee

		successfulSimulation := false
		block, _ := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
		baseFee := block.Header().BaseFee
		baseFee.Mul(baseFee, big.NewInt(114))
		baseFee.Div(baseFee, big.NewInt(100))
		gasTipCap, _ := client.SuggestGasTipCap(context.Background())
		gasTipCap.Mul(baseFee, big.NewInt(110))
		gasTipCap.Div(baseFee, big.NewInt(100))

		nounce, _ := client.NonceAt(context.Background(), signer.Address, nil)

		fmt.Println("Arbitrage search took ", time.Since(now))

		if submitedToFlashbots {
			if lastNounce == nounce {
				//Incrase the bribe - we didn't get included
				if bribeFlat > 99900 {
					bribeFlat += 15
				} else {
					bribeFlat += 1000
				}
				if bribeFlat >= 99999 {
					bribeFlat = 99993
				}
				fmt.Println("Increasing bribe to: ", bribeFlat)
			} else {
				//Decrease the bribe
				bribeFlat -= 3
				fmt.Println("Decreasing bribe to: ", bribeFlat)
			}
		}
		submitedToFlashbots = false
		lastNounce = nounce

		failedCount := 0
		now = time.Now()

		var arb *arbitrage.Arbitrage
		var count int
		for count, arb = range arbs {
			if count > 20 && successfulSimulation {
				break
			} else if count > 200 {
				break
			}

			amountIn, targets, payloads := arb.PrepareTx(ADDR.BUNDLE_EXECUTOR)

			//TODO: Add toogle for private vs PGA
			//Set the bribe amonut
			bribe := big.NewInt(0)
			bribe.Div(arb.Profit.ToBig(), bribeDenominator)
			bribe.Mul(bribe, big.NewInt(bribeFlat))

			executorBalance, _ := wethContract.BalanceOf(nil, ADDR.BUNDLE_EXECUTOR)

			var txData []byte
			if executorBalance.Cmp(amountIn) > 0 {
				//Using without FlashLoan
				txData, err = signer.PrepareDirectArbTx(bribe, targets, payloads)
			} else {
				//Using flasloan
				loanAmounts := []*big.Int{new(big.Int).Sub(amountIn, executorBalance)}
				modes := []*big.Int{big.NewInt(0)}
				loanTokens := make([]common.Address, 0)
				loanTokens = append(loanTokens, ADDR.WETH)
				txData, err = signer.PreperareFlashLoanArbTx(loanTokens, bribe, loanAmounts, modes, targets, payloads)
			}

			if err != nil {
				panic(err)
			}

			_, err = signer.SimulateBundleTx(big.NewInt(0), txData)
			if err != nil {
				failedCount++
				//fmt.Printf("Simulation failed with %v\n", err)
				for _, p := range arb.Path {
					t0, t1 := p.Tokens()
					toBlackList[t0]++
					toBlackList[t1]++
				}
				continue
			}
			if !successfulSimulation {
				profit := arbitrage.FormatUintToFloat(arb.Profit, 1/1e18)
				fmt.Println("Highest simulated profit: ", profit)
			}
			successfulSimulation = true
			//Add some overhead
			estimatedGas, err := signer.EstimateGasForTx(big.NewInt(0), txData)
			if err != nil {
				fmt.Println("Failed to estimate gas cost")
				continue
			}
			estimatedGas = estimatedGas * 110 / 100
			//fmt.Printf("EstimateGas cost %v\n", estimatedGas)
			//fmt.Println("WARNING: IGNORING GAS COST")
			gasPrice := new(big.Int).Add(baseFee, gasTipCap)
			totalCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(estimatedGas)))
			//fmt.Printf("Estimated gas fee: %v for %v gas\n", totalCost, estimatedGas)
			totalCost.Add(totalCost, bribe)
			net := totalCost.Sub(arb.Profit.ToBig(), totalCost)

			if !(net.Cmp(big.NewInt(0)) > 0) {
				//fmt.Println("No longer profitable")
				//Check if proftitable without bribe
				gasPrice := new(big.Int).Add(baseFee, gasTipCap)
				totalCost := new(big.Int).Mul(gasPrice, big.NewInt(int64(estimatedGas)))
				if arb.Profit.ToBig().Cmp(totalCost) > 0 {
					net.Sub(arb.Profit.ToBig(), totalCost)
					bribe = bribe.Mul(net, big.NewInt(bribeFlat))
					bribe = bribe.Div(bribe, bribeDenominator)
					if executorBalance.Cmp(amountIn) > 0 {
						//Using without FlashLoan
						txData, err = signer.PrepareDirectArbTx(bribe, targets, payloads)
					} else {
						//Using flasloan
						loanAmounts := []*big.Int{new(big.Int).Sub(amountIn, executorBalance)}
						modes := []*big.Int{big.NewInt(0)}
						loanTokens := make([]common.Address, 0)
						loanTokens = append(loanTokens, ADDR.WETH)
						txData, err = signer.PreperareFlashLoanArbTx(loanTokens, bribe, loanAmounts, modes, targets, payloads)
					}
				} else {
					continue
				}
			}

			//err = signer.SentTxToBundle(txData, big.NewInt(0), nounce, estimatedGas, gasPrice)
			tx, _ := signer.SignTx(txData, big.NewInt(0), nounce, estimatedGas, baseFee, gasTipCap)
			signedRawTx := flashbots.GetRawSignedTx(tx)
			flashbots.SendPrivateTx(signedRawTx, fmt.Sprintf("0x%x", blockNum+1))

			for _, p := range arb.Path {
				t0, t1 := p.Tokens()
				toBlackList[t0]--
				toBlackList[t1]--
			}
			if err != nil {
				log.Println(err)
				continue
			}
			profit := arbitrage.FormatUintToFloat(arb.Profit, 1/1e18)
			in := arbitrage.FormatUintToFloat(arb.AmountIn, 1/1e18)
			bribeUint, _ := uint256.FromBig(bribe)
			fmt.Printf("\nTrying trade path for %v with %v profit and %v bribe\n", in, profit, arbitrage.FormatUintToFloat(bribeUint, 1/1e18))
			fmt.Printf("Estimated gas fee: %v for %v gas\n", totalCost, estimatedGas)
			arbitrage.PrintPath(ADDR.WETH, arb.Path)
			log.Printf("Succesfull execution at: %v\n", blockNum+1)
			submitedToFlashbots = true
			break
		}

		highest := 0
		badToken := common.HexToAddress("")
		for k, v := range toBlackList {
			if v > BLACKLIST_THRESHOLD {
				toBlackList[k] = 0
				if v > highest {
					highest = v
					badToken = k
				}
			}
		}

		if highest != 0 {
			fmt.Printf("Blacklisted %v \n", badToken)
			BLACKLIST[badToken] = true
			delete(toBlackList, badToken)
			arbitrage.SaveBlackList(BLACKLIST)
			blacklistedTokens++
		}

		if blacklistedTokens >= 1 {
			fmt.Println("Recreating paths after blacklist updates")
			//tradingPaths = arbitrage.GeneratePaths(client, targetToken, 5, BLACKLIST, TOKEN_DECIMALS, uniV2, sushi, shiba, orion, balancerWeighted)
			newTradingPaths := make([][]arbitrage.Pair, 0)
			for _, path := range tradingPaths {
				blacklisted := false
				for _, p := range path {
					t0, t1 := p.Tokens()
					if BLACKLIST[t0] || BLACKLIST[t1] {
						blacklisted = true
						break
					}
				}
				if !blacklisted {
					newTradingPaths = append(newTradingPaths, path)
				}
			}
			blacklistedTokens = 0
			tradingPaths = newTradingPaths
			fmt.Printf("Have %v paths\n", len(tradingPaths))
			blockNum, _ = client.BlockNumber(context.Background())
		}

		//TODO: Change this for tx confirmation
		fmt.Printf("%v/%v simulations reverted\n", failedCount, count)
		failedCount = 0
		fmt.Println("Simulations took: ", time.Since(now))
		fmt.Printf("\n\nSleeping for the next 1 block\n")
		newBlock, _ := client.BlockNumber(context.Background())
		for newBlock-blockNum < 1 {
			newBlock, _ = client.BlockNumber(context.Background())
			time.Sleep(time.Millisecond * 50)
		}
		blockNum, _ = client.BlockNumber(context.Background())
		log.Println("New block")
	}

}

func updateReserves(client *ethclient.Client, pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi []*uniswapV2.Pair, pairsBalancerWeighted []*balancerV2.WeightedPair, pairsBalancer2Token []*balancerV2.Weighted2TokenPair) {
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsUniV2)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsSushi)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsShibaSwap)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsOrion)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsSake)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsDXSwap)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsCryptDefi)
	balancerV2.UpdateReservesWeighted(client, ADDR.BALANCERv2_VAULT, pairsBalancerWeighted)
	balancerV2.UpdateReserves2TokensWeighted(client, ADDR.BALANCERv2_VAULT, pairsBalancer2Token)
	// curve.UpdateReserves(client, ADDR.CURVE_REGISTRY, pairsCurve)

}
