package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/arbitrage"
	"github.com/jmjac/ethGoTest/dex"
	"github.com/jmjac/ethGoTest/dex/balancerV2"
	"github.com/jmjac/ethGoTest/dex/uniswapV2"
	"github.com/jmjac/ethGoTest/dex/uniswapV3"
	"github.com/jmjac/ethGoTest/flashbots"
	"github.com/jmjac/ethGoTest/helperClient"
	"github.com/jmjac/ethGoTest/secrets"
)

func main() {
	//Setup log files
	genf, err := os.OpenFile("general.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	genLog := log.New(genf, "", 0)

	submitf, err := os.OpenFile("submitted.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	subLog := log.New(submitf, "", 0)
	revf, err := os.OpenFile("reverted.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	revLog := log.New(revf, "", 0)
	errorf, err := os.OpenFile("error.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	errLog := log.New(errorf, "", 0)
	noProfitf, err := os.OpenFile("noProfit.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	noProfitLog := log.New(noProfitf, "", 0)
	defer genf.Close()
	defer submitf.Close()
	defer revf.Close()
	defer errorf.Close()
	defer noProfitf.Close()

	dex.SetLogger(*genLog)

	var provider string
	//Read from env in the future
	url := "https://xyznode.space/api/mainnet/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"
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

	//Setup adreses
	if mainnet {
		ADDR.MainnetAddresses(*genLog)
	} else {
		ADDR.GoerliAddresses(*genLog)
	}

	// =============================== SETUP BLACKLIST ==========================
	BLACKLIST := arbitrage.InitBlackList("blacklist.json")
	genLog.Printf("Blacklist has %v tokens\n", len(BLACKLIST))
	const BLACKLIST_THRESHOLD = 150
	toBlackList := make(map[common.Address]int)
	toBlackList = arbitrage.AddWhiteListed(toBlackList)
	// ============================= END SETUP BLACKLIST ========================

	// ============================= LOAD TOKENS WITH TAX ======================
	tokens := dex.LoadTokens("tokensTax.json")
	tokenMap := dex.MakeTokenMap(tokens)
	count := 0
	for _, t := range tokens {
		if !t.IsTradable { //Add not tradable tokens to blacklist
			BLACKLIST[t.Addr] = true
			count++
		}
	}
	genLog.Printf("Removed %v not tradable tokens\n", count)
	// ============================= END TOKENS WITH TAX ======================

	TOKEN_DECIMALS := arbitrage.LoadDecimalsFromJson("token_decimals.json")

	// ============================= LOAD BALANCERV2 POOLS ======================
	pairsBalancer2Token := balancerV2.Get2TokenWeightedMarket(ethClient, ADDR.BALANCERv2_VAULT, ADDR.BALANCERv2_WEIGHTED2TOKENS_FACTORY, big.NewInt(0), BLACKLIST, tokenMap)
	pairsBalancerWeighted := balancerV2.GetWeightedMarkets(ethClient, ADDR.BALANCERv2_VAULT, ADDR.BALANCERv2_WEIGHTED_FACTORY, big.NewInt(0), BLACKLIST, tokenMap)

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

	genLog.Printf("Unique %v trading pairs from BalancerV2Weighted\n", len(allBalancerPairs))
	// ============================= END LOAD BALANCERV2 POOLS ======================

	// ============================= LOAD UNIV2 POOLS ======================
	pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsPancakeSwap, pairsRadioShack, pairsLuaSwap, pairsSaitaSwap := setupUniV2Pairs(ethClient, BLACKLIST, tokenMap)
	allUniV2 := mergeUniV2Markets(pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsPancakeSwap, pairsRadioShack, pairsLuaSwap, pairsSaitaSwap)
	uniV2, sushi, shiba, orion, sake, dxSwap, cryptDefi, frax, pancake, radioshack, luaSwap, saitaSwap := castUniV2Pairs(pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsPancakeSwap, pairsRadioShack, pairsLuaSwap, pairsSaitaSwap)
	// ============================= END LOAD UNIV2 POOLS ======================

	// ============================= LOAD DODO DVM POOLS ======================
	// ============================= END LOAD DODO DVM POOLS ======================

	// ============================================= UNISWAPV3 POOLS =======================================================
	bNum, _ := ethClient.BlockNumber(context.Background())
	bNum -= 3
	now := time.Now()

	pairsUniV3 := make([]*uniswapV3.Pair, 0)
	if mainnet {
		pairsUniV3 = uniswapV3.InitialSyncWithSubgraph(ethClient, ADDR.UNIv3_FACTORY, BLACKLIST, bNum, tokenMap)
	}
	// ============================================= END UNISWAPV3 POOLS =======================================================

	// ============================================= CAST EACH POOL TO PAIR =========================================================
	balancerWeighted := make([]dex.Pair, 0, len(allBalancerPairs))
	for _, p := range allBalancerPairs {
		balancerWeighted = append(balancerWeighted, p)
	}
	uniV3 := make([]dex.Pair, 0, len(pairsUniV3))
	for _, p := range pairsUniV3 {
		uniV3 = append(uniV3, p)
	}
	// ============================================= END CAST EACH POOL TO PAIR =========================================================

	genLog.Println("Generating trading paths")
	now = time.Now()
	targetToken := ADDR.WETH
	toBlackList[targetToken] = -1000000000
	genLog.Println("NOT USING BALANCER")
	//TODO: Update logging for this
	tradePaths := make([][]dex.Pair, 0)
	//tradePaths := arbitrage.GeneratePaths(1, ethClient, targetToken, 4, BLACKLIST, TOKEN_DECIMALS, sushi, shiba, orion, sake, dxSwap, uniV2, cryptDefi, frax, pancake, radioshack, luaSwap, uniV3) //BalancerV2Weighted
	tradePaths2 := arbitrage.GeneratePaths(0, ethClient, targetToken, 5, BLACKLIST, TOKEN_DECIMALS, shiba, orion, sake, dxSwap, uniV2, cryptDefi, frax, pancake, radioshack, sushi, luaSwap, saitaSwap) //BalancerV2Weighted

	for _, p := range tradePaths2 {
		//		if len(p) < 5 {
		//			continue
		//		}
		tradePaths = append(tradePaths, p)
	}

	genLog.Println("Trading paths: \t", len(tradePaths))
	arbitrage.SaveTokenDecimals(TOKEN_DECIMALS, "token_decimals.json")
	genLog.Printf("Found %v paths\n", len(tradePaths))
	genLog.Println("Took: ", time.Since(now))

	blockNum, _ := ethClient.BlockNumber(context.Background())

	uniswapV3.UpdateReserves(pairsUniV3, blockNum, ethClient)

	if !mainnet {
		panic("This is not updated to work on goerli")
	}

	mevReyals := make([]flashbots.MevRelay, 0)
	mevReyals = append(mevReyals, flashbots.NewFlashbotsMainnet())
	mevReyals = append(mevReyals, flashbots.NewMevRelay("https://rpc.beaverbuild.org/", "beaverbuilder"))
	mevReyals = append(mevReyals, flashbots.NewMevRelay("https://builder0x69.io", "builder0x69"))
	mevReyals = append(mevReyals, flashbots.NewMevRelay("https://api.edennetwork.io/v1/bundle", "edenBuilder"))
	mevReyals = append(mevReyals, flashbots.NewMevRelay("https://eth-builder.com", "eth-builder"))

	signerPureV2 := flashbots.NewSigner(big.NewInt(1), secrets.PRIVATE_KEY_PURE_V2, ADDR.BUNDLE_EXECUTOR, ADDR.WETH, ethClient)
	signerMixed := flashbots.NewSigner(big.NewInt(1), secrets.PRIVATE_KEY_MIXED, ADDR.BUNDLE_EXECUTOR, ADDR.WETH, ethClient)

	blockNum, _ = ethClient.BlockNumber(context.Background())
	newBlock, _ := ethClient.BlockNumber(context.Background())
	genLog.Println("Syncing with blocktime")
	for newBlock-blockNum < 1 {
		newBlock, _ = ethClient.BlockNumber(context.Background())
		time.Sleep(time.Millisecond * 100)
	}
	genLog.Printf("Found new block at %v", time.Now())

	//TODO: THIS SHOULD BE SAVED
	bribeNumerator := int64(9966944)
	bribeDenominator := big.NewInt(10000000)
	if !mainnet {
		bribeNumerator = int64(1000)
	}

	nounceMixed, _ := ethClient.NonceAt(context.Background(), signerMixed.Address, nil)
	nouncePureV2, _ := ethClient.NonceAt(context.Background(), signerPureV2.Address, nil)
	lastNounceMixed := nounceMixed
	lastNouncePure := nouncePureV2
	submitedToFlashbots := false

	MIN_AMOUNT_IN := uint256.NewInt(1e16)
	MIN_AMOUNT_IN.Mul(MIN_AMOUNT_IN, uint256.NewInt(1))
	MAX_AMOUNT_IN := uint256.NewInt(84)
	MAX_AMOUNT_IN.Mul(MAX_AMOUNT_IN, uint256.NewInt(1e16))

	genLog.Printf("\tArb token: %v\n", targetToken)
	genLog.Printf("\tMIN_AMOUNT_IN: %v\n", MIN_AMOUNT_IN)
	genLog.Printf("\tMAX_AMOUNT_IN: %v\n", MAX_AMOUNT_IN)

	//targetTokenContract, err := abis.NewIERC20(targetToken, ethClient)
	bot := arbitrage.NewArbBot(revLog, genLog, subLog, errLog, noProfitLog, *signerPureV2, *signerMixed, ethClient, mevReyals, BLACKLIST, toBlackList)

	for i := 0; ; i++ {

		now := time.Now()
		uniswapV2.UpdateReservesWithoutLookup(allUniV2, ethClient)
		//curve.UpdateReserves(ethClient, ADDR.CURVE_REGISTRY, pairsCurve)
		genLog.Println("Balancer reserves disabled")
		//balancerV2.UpdateReservesWeighted(ethClient, ADDR.BALANCERv2_VAULT, pairsBalancerWeighted)
		//balancerV2.UpdateReserves2TokensWeighted(ethClient, ADDR.BALANCERv2_VAULT, pairsBalancer2Token)
		err = uniswapV3.UpdateReserves(pairsUniV3, blockNum, ethClient)

		//TODO: Add recovery from failed update of reserves
		if err != nil {
			genLog.Println("UniswapV3 failed to update reserves, exiting")
			os.Exit(1)
		}

		genLog.Println("Updating reserves took: ", time.Since(now))

		now = time.Now()
		block, _ := ethClient.BlockByNumber(context.Background(), nil)
		baseFee := ethClient.CalcBaseFee(params.MainnetChainConfig, block.Header())
		nouncePureV2, _ := ethClient.NonceAt(context.Background(), signerPureV2.Address, nil)
		nounceMixed, _ := ethClient.NonceAt(context.Background(), signerMixed.Address, nil)

		arbs := arbitrage.LookForArb(threads, targetToken, MIN_AMOUNT_IN, MAX_AMOUNT_IN, uint256.NewInt(1025), uint256.NewInt(1000), tradePaths)
		arbs = arbitrage.MoveArbsAround(arbs) //Change ordering to have a better split between short and longs arbs
		genLog.Println("Arbitrage search took: ", time.Since(now))

		if submitedToFlashbots {
			bribeNumerator = updateBribe(lastNounceMixed, nounceMixed, lastNouncePure, nouncePureV2, bribeNumerator, *genLog)
		}

		lastNounceMixed = nounceMixed
		lastNouncePure = nouncePureV2

		//Update bot state
		bot.UpdateForNewBlock(blockNum, nouncePureV2, nounceMixed, baseFee, big.NewInt(bribeNumerator), bribeDenominator)

		// Check Arbs and submit if profitable
		submitedToFlashbots = bot.CheckProfits(200, 400, arbs)

		// TODO: Make this a function
		if bot.BlacklistWasUpdated(BLACKLIST_THRESHOLD) {
			genLog.Println("Recreating paths after blacklist updates")
			newTradingPaths := make([][]dex.Pair, 0)
			for _, path := range tradePaths {
				blacklisted := false
				for _, p := range path {
					t0, t1 := p.Tokens()
					if bot.BLACKLIST[t0.Addr] || bot.BLACKLIST[t1.Addr] {
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

		genLog.Printf("Sleeping for the next 1 block\n")
		newBlock, _ := ethClient.BlockNumber(context.Background())
		//Wait for new block
		for newBlock-blockNum < 1 {
			newBlock, _ = ethClient.BlockNumber(context.Background())
			time.Sleep(time.Millisecond * 50)
		}
		blockNum, _ = ethClient.BlockNumber(context.Background())
		genLog.Println()
		genLog.Println()
		genLog.Printf("%v New block: %v", time.Now().Format("2006-01-02 15:04:05"), blockNum)
	}

}

func updateReserves(client bind.ContractBackend, pairsUniV2, pairsPancakeSwap, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsRadioShack []*uniswapV2.Pair, pairsBalancerWeighted []*balancerV2.WeightedPair, pairsBalancer2Token []*balancerV2.Weighted2TokenPair) {
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsUniV2)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsPancakeSwap)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsSushi)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsShibaSwap)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsOrion)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsSake)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsDXSwap)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsCryptDefi)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsFraxSwap)
	uniswapV2.UpdateReserves(client, ADDR.UNISWAPv2_LOOKUP, pairsRadioShack)
	balancerV2.UpdateReservesWeighted(client, ADDR.BALANCERv2_VAULT, pairsBalancerWeighted)
	balancerV2.UpdateReserves2TokensWeighted(client, ADDR.BALANCERv2_VAULT, pairsBalancer2Token)
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

func setupUniV2Pairs(client bind.ContractBackend, BLACKLIST map[common.Address]bool, tokenMap map[common.Address]dex.Token) ([]*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair) {
	pairsSushi := uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.SUSHI_FACTORY, "sushi-swap", BLACKLIST, tokenMap)
	pairsUniV2 := uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.UNIv2_FACTORY, "uniswapV2", BLACKLIST, tokenMap)
	pairsShibaSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.SHIBA_FACTORY != common.HexToAddress("") {
		pairsShibaSwap = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.SHIBA_FACTORY, "shiba-swap", BLACKLIST, tokenMap)
	}
	pairsOrion := make([]*uniswapV2.Pair, 0)
	if ADDR.ORION_FACTORY != common.HexToAddress("") {
		pairsOrion = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.ORION_FACTORY, "orion-swap", BLACKLIST, tokenMap)
	}
	pairsSake := make([]*uniswapV2.Pair, 0)
	if ADDR.SAKE_FACTORY != common.HexToAddress("") {
		pairsSake = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.SAKE_FACTORY, "sake-swap", BLACKLIST, tokenMap)
	}
	pairsDXSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.DX_FACTORY != common.HexToAddress("") {
		pairsDXSwap = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.DX_FACTORY, "dx-swap", BLACKLIST, tokenMap)
	}
	pairsCryptDefi := make([]*uniswapV2.Pair, 0)
	if ADDR.CRYPTDEFI_FACTORY != common.HexToAddress("") {
		pairsCryptDefi = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.CRYPTDEFI_FACTORY, "cryptdefi-swap", BLACKLIST, tokenMap)
	}
	pairsFraxSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.FRAX_FACTORY != common.HexToAddress("") {
		pairsFraxSwap = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.FRAX_FACTORY, "frax-swap", BLACKLIST, tokenMap)
	}
	pairsPancakeSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.PANCAKESWAP_FACTORY != common.HexToAddress("") {
		pairsPancakeSwap = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.PANCAKESWAP_FACTORY, "pancake-swap", BLACKLIST, tokenMap)
	}
	pairsRadioShack := make([]*uniswapV2.Pair, 0)
	if ADDR.RADIOSHACK_FACTORY != common.HexToAddress("") {
		pairsRadioShack = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.RADIOSHACK_FACTORY, "radioshack-swap", BLACKLIST, tokenMap)
	}
	pairsLuaSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.LUASWAP_FACTORY != common.HexToAddress("") {
		pairsLuaSwap = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.LUASWAP_FACTORY, "luaSwap", BLACKLIST, tokenMap)
	}
	pairsSaitaSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.LUASWAP_FACTORY != common.HexToAddress("") {
		pairsSaitaSwap = uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.SAITASWAP_FACTORY, "saitaSwap", BLACKLIST, tokenMap)
	}

	return pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsPancakeSwap, pairsRadioShack, pairsLuaSwap, pairsSaitaSwap
}

func castUniV2Pairs(pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsPancakeSwap, pairsRadioShack, pairsLuaSwap, pairsSaitaSwap []*uniswapV2.Pair) ([]dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair) {
	uniV2 := make([]dex.Pair, 0, len(pairsUniV2))
	for _, p := range pairsUniV2 {
		uniV2 = append(uniV2, p)
	}
	sushi := make([]dex.Pair, 0, len(pairsSushi))
	for _, p := range pairsSushi {
		sushi = append(sushi, p)
	}
	shiba := make([]dex.Pair, 0, len(pairsShibaSwap))
	for _, p := range pairsShibaSwap {
		shiba = append(shiba, p)
	}
	orion := make([]dex.Pair, 0, len(pairsOrion))
	for _, p := range pairsOrion {
		orion = append(orion, p)
	}
	sake := make([]dex.Pair, 0, len(pairsSake))
	for _, p := range pairsSake {
		sake = append(sake, p)
	}
	dxSwap := make([]dex.Pair, 0, len(pairsDXSwap))
	for _, p := range pairsDXSwap {
		dxSwap = append(dxSwap, p)
	}
	cryptDefi := make([]dex.Pair, 0, len(pairsCryptDefi))
	for _, p := range pairsCryptDefi {
		cryptDefi = append(cryptDefi, p)
	}
	frax := make([]dex.Pair, 0, len(pairsFraxSwap))
	for _, p := range pairsFraxSwap {
		frax = append(frax, p)
	}
	pancake := make([]dex.Pair, 0, len(pairsPancakeSwap))
	for _, p := range pairsPancakeSwap {
		pancake = append(pancake, p)
	}
	radioShack := make([]dex.Pair, 0, len(pairsRadioShack))
	for _, p := range pairsRadioShack {
		radioShack = append(radioShack, p)
	}

	luaSwap := make([]dex.Pair, 0, len(pairsLuaSwap))
	for _, p := range pairsLuaSwap {
		luaSwap = append(luaSwap, p)
	}

	saitaSwap := make([]dex.Pair, 0, len(pairsSaitaSwap))
	for _, p := range pairsSaitaSwap {
		saitaSwap = append(saitaSwap, p)
	}

	return uniV2, sushi, shiba, orion, sake, dxSwap, cryptDefi, frax, pancake, radioShack, luaSwap, saitaSwap
}

func updateBribe(lastPure, newPure, lastMixed, newMixed uint64, bribeNumerator int64, genLog log.Logger) int64 {
	//Stats from last submission to flashbots
	//flashbots.GetBundleStats(bundleHash, bundleBlock)
	if lastPure == newPure && lastMixed == newMixed {
		//Incrase the bribe - we didn't get included
		if bribeNumerator < 9000000 {
			bribeNumerator += 10000
		} else if bribeNumerator < 9500000 {
			bribeNumerator += 500
		} else if bribeNumerator < 9650000 {
			bribeNumerator += 250
		} else if bribeNumerator < 9700000 {
			bribeNumerator += 80
		} else {
			bribeNumerator += 10
		}

		if bribeNumerator > 9999999 {
			genLog.Println("Bribe reached the limit!")
			bribeNumerator = 9950000
		} else {
			genLog.Println("Didn't get into the block. Increasing bribe to: ", bribeNumerator)
		}
	} else {
		//Decrease the bribe
		bribeNumerator -= 4
		genLog.Println("GOT INTO THE BLOCK!. Decreasing bribe to: ", bribeNumerator)
	}

	return bribeNumerator
}
