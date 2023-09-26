package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/arbitrage"
	"github.com/jmjac/ethGoTest/dex"
	"github.com/jmjac/ethGoTest/dex/uniswapV2"
	"github.com/jmjac/ethGoTest/helperClient"
)

func main() {
	var provider string
	//Read from env in the future
	url := "https://xyznode.space/api/bsc/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"

	flag.StringVar(&provider, "provider", url, "RPC endpoint")
	flag.Parse()

	if provider == "" {
		fmt.Println("Provider can't be empty. Use --provider")
		return
	}
	client, err := ethclient.Dial(provider)
	if err != nil {
		fmt.Println("Coudn't create client,", err)
		return
	}

	//Setup adreses
	ADDR.BnbAddresses()

	//signer := flashbots.NewSigner(big.NewInt(56), secrets.PRIVATE_KEY_GOERLI, ADDR.BUNDLE_EXECUTOR, ADDR.WETH, client)

	helper, _ := helperClient.New(url)
	//err := _UniswapV2Query.contract.Call(opts, &out, "getPairsByIndexRange", _uniswapFactory, _start, _stop)

	// =============================== SETUP BLACKLIST ==========================
	BLACKLIST := arbitrage.InitBlackList("blacklist_bnb.json")
	fmt.Printf("Blacklist has %v tokens\n", len(BLACKLIST))
	const BLACKLIST_THRESHOLD = 50
	toBlackList := make(map[common.Address]int)
	toBlackList = arbitrage.AddWhiteListed(toBlackList)
	toBlackList[ADDR.WETH] = -100000000
	toBlackList[ADDR.DAI] = -100000000
	// ============================= END SETUP BLACKLIST ========================

	TOKEN_DECIMALS := arbitrage.LoadDecimalsFromJson("token_decimals_bnb.json")

	pairsNomiSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.NOMISWAP_FACTORY != common.HexToAddress("") {
		pairsNomiSwap = uniswapV2.GetUniswappyPairsWithoutLookup(helper, ADDR.NOMISWAP_FACTORY, "nomi-swap", BLACKLIST, make(map[common.Address]dex.Token))
	}

	pairsApeSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.APESWAP_FACTORY != common.HexToAddress("") {
		pairsApeSwap = uniswapV2.GetUniswappyPairsWithoutLookup(helper, ADDR.APESWAP_FACTORY, "ape-swap", BLACKLIST, make(map[common.Address]dex.Token))
	}
	pairsPancakeSwap := make([]*uniswapV2.Pair, 0)
	if ADDR.PANCAKESWAP_FACTORY != common.HexToAddress("") {
		pairsPancakeSwap = uniswapV2.GetUniswappyPairsWithoutLookup(helper, ADDR.PANCAKESWAP_FACTORY, "pancake-swap", BLACKLIST, make(map[common.Address]dex.Token))
	}
	//TODO: ADD NomiSwap, bitswap, safeswap, apeswap

	toLookup := make(map[common.Address]bool)

	for _, m := range [][]*uniswapV2.Pair{pairsApeSwap, pairsNomiSwap, pairsPancakeSwap} {
		for _, p := range m {
			t0, t1 := p.Tokens()
			if _, ok := TOKEN_DECIMALS[t0.Addr]; !ok && !BLACKLIST[t0.Addr] {
				toLookup[t0.Addr] = true
			}
			if _, ok := TOKEN_DECIMALS[t1.Addr]; !ok && !BLACKLIST[t1.Addr] {
				toLookup[t1.Addr] = true
			}
		}
	}

	addresses := make([]common.Address, 0)
	for k := range toLookup {
		addresses = append(addresses, k)
	}

	fmt.Println(len(addresses))

	step := 1000
	for i := 0; i < len(addresses); i += step {
		end := i + step
		if end > len(addresses) {
			end = len(addresses)
		}

		dec, err := helper.LookupDecimals(addresses[i:end])
		if err != nil {
			for check := i; check < end; check++ {
				dec, err = helper.LookupDecimals(addresses[check : check+1])
				if err != nil || dec[0] < 0 {
					fmt.Println("Problem with: ", addresses[check])
					BLACKLIST[addresses[check]] = true
				} else {
					TOKEN_DECIMALS[addresses[check]] = uint64(dec[0])
				}
			}
			continue
		}

		for j, t := range addresses[i:end] {
			if dec[0] >= 0 {
				TOKEN_DECIMALS[t] = uint64(dec[j])
			} else {
				//ADD to blacklist
				BLACKLIST[t] = true
				fmt.Println("Problem with: ", t)
			}
		}
	}
	arbitrage.SaveTokenDecimals(TOKEN_DECIMALS, "token_decimals_bnb.json")
	arbitrage.SaveBlackList(BLACKLIST, "blacklist_bnb.json")

	//TODO: Remove
	r0, r1 := pairsApeSwap[0].Reserves()
	fmt.Printf("%v and %v\n", r0, r1)

	// ============================================= END LOAD DIFFERENT POOLS =======================================================

	// ============================================= CAST EACH POOL TO PAIR =========================================================
	pancake := make([]dex.Pair, 0, len(pairsPancakeSwap))
	for _, p := range pairsPancakeSwap {
		pancake = append(pancake, p)
	}
	nomi := make([]dex.Pair, 0, len(pairsNomiSwap))
	for _, p := range pairsNomiSwap {
		nomi = append(nomi, p)
	}
	ape := make([]dex.Pair, 0, len(pairsApeSwap))
	for _, p := range pairsPancakeSwap {
		ape = append(ape, p)
	}

	// ============================================= END CAST EACH POOL TO PAIR =========================================================
	allUniPairs := make([]*uniswapV2.Pair, 0)
	for _, m := range [][]*uniswapV2.Pair{pairsApeSwap, pairsNomiSwap, pairsPancakeSwap} {
		for _, p := range m {
			t0, t1 := p.Tokens()
			if !BLACKLIST[t0.Addr] && !BLACKLIST[t1.Addr] {
				allUniPairs = append(allUniPairs, p)
			}
		}
	}

	now := time.Now()
	//uniswapV2.UpdateReservesWithoutLookup(pairsPancakeSwap, helper)
	//uniswapV2.UpdateReservesWithoutLookup(pairsNomiSwap, helper)
	//uniswapV2.UpdateReservesWithoutLookup(pairsApeSwap, helper)
	uniswapV2.UpdateReservesWithoutLookup(allUniPairs, helper)
	fmt.Println("Updating reserves took: ", time.Since(now))

	fmt.Println("Generating trading paths")
	now = time.Now()

	targetToken := ADDR.WBNB
	tradingPaths := arbitrage.GeneratePaths(4, client, targetToken, 4, BLACKLIST, TOKEN_DECIMALS, pancake, ape, nomi)

	//TODO: Separete function to load token decimals using a contract
	arbitrage.SaveTokenDecimals(TOKEN_DECIMALS, "token_decimals_bnb.json")

	fmt.Printf("Found %v paths\n", len(tradingPaths))
	fmt.Println("Took: ", time.Since(now))
	return

	fmt.Println("Syncing with blocktime")

	blockChanged := checkIfBlockChanged(client)

	minAmountIn := uint256.NewInt(1e16)
	maxAmountIn := uint256.NewInt(5e18)
	mulConst := uint256.NewInt(103)
	divConst := uint256.NewInt(100)

	for {
		select {
		case <-blockChanged:
			now := time.Now()
			uniswapV2.UpdateReservesWithoutLookup(pairsPancakeSwap, helper)
			fmt.Println("Updating reserves took: ", time.Since(now))
			now = time.Now()
			arbs := arbitrage.LookForArb(256, ADDR.WBNB, minAmountIn, maxAmountIn, mulConst, divConst, tradingPaths)
			fmt.Println("Arb search took: ", time.Since(now))
			fmt.Println(arbs[0].Profit)
		default:
			//Submit with higher gas price
		}
	}

}

func checkIfBlockChanged(client *ethclient.Client) chan bool {
	changed := make(chan bool, 3)
	go func() {
		current, _ := client.BlockNumber(context.Background())
		for {
			blockNum, _ := client.BlockNumber(context.Background())
			if current == blockNum {
				time.Sleep(30 * time.Millisecond)
			} else {
				changed <- true
				current = blockNum
			}
		}
	}()

	return changed
}
