package arbitrage

import (
	"fmt"
	"log"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/abis"
	"github.com/jmjac/ethGoTest/dex"
)

type Arbitrage struct {
	Path       []dex.Pair
	StartToken common.Address
	AmountIn   *uint256.Int
	Profit     *uint256.Int
}

type Found struct {
	profit   *uint256.Int
	path     []dex.Pair
	amountIn *uint256.Int
}

var MAX_AMOUNT_IN = uint256.NewInt(0).Mul(uint256.NewInt(1e18), uint256.NewInt(20))
var MUL_CONST = uint256.NewInt(101)
var DIV_CONST = uint256.NewInt(100)

func LookForArb(threads int, tokenIn common.Address, minAmountIn, maxAmountIn, mulConst, divConst *uint256.Int, paths [][]dex.Pair) []*Arbitrage {
	if minAmountIn.Cmp(maxAmountIn) >= 0 {
		log.Fatal("minAmountIn must be smaller than maxAmountIn")
	}

	if mulConst.Cmp(divConst) <= 0 {
		log.Fatal("mulConst must be bigger than divConst")
	}

	MAX_AMOUNT_IN.Set(maxAmountIn)
	MUL_CONST.Set(mulConst)
	DIV_CONST.Set(divConst)

	//step := len(paths) / threads
	profitChan := make(chan Found, 128*512)
	arbs := make([]*Arbitrage, 0)
	var wgProfit sync.WaitGroup
	wgProfit.Add(1)

	go func(c <-chan Found) {
		for f := range c {
			a := &Arbitrage{Path: f.path, AmountIn: f.amountIn, StartToken: tokenIn, Profit: new(uint256.Int).Set(f.profit)}
			arbs = append(arbs, a)
		}
		wgProfit.Done()
	}(profitChan)

	pathsLimit := make(chan bool, 2048)
	stepSize := 500

	var wg sync.WaitGroup
	for i := 0; i < len(paths); i += stepSize {
		end := i + stepSize
		if end > len(paths) {
			end = len(paths)
		}
		wg.Add(1)
		pathsLimit <- true
		start := i
		go func(s, e int) {
			checkTradingPaths(tokenIn, minAmountIn, paths[start:end], profitChan)
			<-pathsLimit
			wg.Done()
		}(start, end)
	}
	wg.Wait()
	close(profitChan)
	close(pathsLimit)
	wgProfit.Wait()

	if len(arbs) <= 1 {
		return arbs
	}
	sort.Slice(arbs, func(i, j int) bool {
		return arbs[i].Profit.Cmp(arbs[j].Profit) > 0
	})
	return arbs
}

//Moves around arbs for a betetr split between short and long trades
func MoveArbsAround(arbs []*Arbitrage) []*Arbitrage {
	shortArbs := 0
	arbsMoved := make([]*Arbitrage, 0)
	added := make(map[int]bool)

	//Put the 40 most profitable short arbs first
	//TODO: Move this to bot logic instead of arb search
	for i, a := range arbs {
		if len(a.Path) <= 2 {
			shortArbs++
			added[i] = true
			arbsMoved = append(arbsMoved, a)
		}
		if shortArbs >= 10 {
			break
		}
	}

	for i, a := range arbs {
		if len(a.Path) <= 3 && !added[i] {
			shortArbs++
			added[i] = true
			arbsMoved = append(arbsMoved, a)
		}
		if shortArbs >= 44 {
			break
		}
	}

	for i, a := range arbs {
		if !added[i] {
			arbsMoved = append(arbsMoved, a)
		}
	}
	return arbsMoved
}

var gorutineLimit = make(chan bool, 1)

func GeneratePaths(heavyLimit int, client bind.ContractBackend, token common.Address, maxHops int, BLACKLIST map[common.Address]bool, tokenDecimals map[common.Address]uint64, markets ...[]dex.Pair) [][]dex.Pair {
	fmt.Printf("Looking up to depth = %v\n", maxHops)
	//TODO: Change this if needed
	minReserve := uint256.NewInt(10)
	poolMap := make(map[common.Address][]dex.Pair)
	for _, m := range markets {
		for _, p := range m {
			t0, t1 := p.Tokens()
			if BLACKLIST[t0.Addr] || BLACKLIST[t1.Addr] {
				continue
			}
			r0, r1 := p.Reserves()
			var d uint64
			dec, ok := tokenDecimals[t0.Addr]
			if !ok {
				contract0, _ := abis.NewIERC20Caller(t0.Addr, client)
				d, _ := contract0.Decimals(nil)
				tokenDecimals[t0.Addr] = uint64(d)
				dec = uint64(d)
			}
			minReserve.Exp(uint256.NewInt(10), uint256.NewInt(uint64(dec)))
			minReserve.Div(minReserve, uint256.NewInt(3000000))

			//minReserve.Mul(minReserve, uint256.NewInt(8))
			if r0.Cmp(minReserve) <= 0 {
				continue
			}

			dec, ok = tokenDecimals[t1.Addr]
			if !ok {
				contract1, _ := abis.NewIERC20Caller(t1.Addr, client)
				d, _ := contract1.Decimals(nil)
				tokenDecimals[t1.Addr] = uint64(d)
				dec = uint64(d)
			}
			minReserve.Exp(uint256.NewInt(10), uint256.NewInt(uint64(d)))
			minReserve.Div(minReserve, uint256.NewInt(3000000))

			//minReserve.Mul(minReserve, uint256.NewInt(8))
			if r1.Cmp(minReserve) <= 0 {
				continue
			}
			if _, ok := poolMap[t0.Addr]; !ok {
				poolMap[t0.Addr] = make([]dex.Pair, 0)
			}
			if _, ok := poolMap[t1.Addr]; !ok {
				poolMap[t1.Addr] = make([]dex.Pair, 0)
			}

			poolMap[t0.Addr] = append(poolMap[t0.Addr], p)
			poolMap[t1.Addr] = append(poolMap[t1.Addr], p)
		}
	}
	fmt.Println("Got ", len(poolMap), " pairs with enough liquidity")
	fmt.Println("Finding trades paths now")

	var wg sync.WaitGroup
	allPaths := make(chan []dex.Pair, 2048)
	gorutineLimit = make(chan bool, 64000)
	wg.Add(1)
	go findTradingPaths(poolMap, token, token, heavyLimit, maxHops, maxHops, make([]dex.Pair, maxHops, maxHops), allPaths, &wg)

	paths := make([][]dex.Pair, 0)
	var wgPaths sync.WaitGroup
	wgPaths.Add(1)

	go func(all chan []dex.Pair) {
		for path := range all {
			paths = append(paths, path)
		}
		wgPaths.Done()
	}(allPaths)

	wg.Wait()
	close(allPaths)
	close(gorutineLimit)
	wgPaths.Wait()

	//Sort by length so that, maybe useful?

	//TODO: Uncomment when cache is implemented
	//	sort.Slice(paths, func(i, j int) bool {
	//		return len(paths[i]) < len(paths[j])
	//	})
	//	for _, p := range paths {
	//		val := ""
	//		for _, k := range p[:len(p)-1] {
	//			val += k.MarketAddress().Hex()
	//		}
	//		cache[val] = new(uint256.Int)
	//	}

	return paths
}

func clearCache() {
	for k := range cache {
		cache[k].Clear()
	}
}

var cache = make(map[string]*uint256.Int)

// Limits amount in to max 50 ETH

//TODO: Add caching for values already computed, maybe sort paths by lenght first
func checkTradingPaths(tokenIn common.Address, v *uint256.Int, tradingPaths [][]dex.Pair, f chan Found) {
	//	two := uint256.NewInt(2)

	amountIn := new(uint256.Int).Set(v)
	prevProfit := uint256.NewInt(0)
	prevAmountIn := uint256.NewInt(0)

	for i := 0; i < len(tradingPaths); i++ {
		profitIncreased := false
		path := tradingPaths[i]
		token := tokenIn
		out := amountIn

		for _, p := range path {
			//TODO: Here we should add trade tax logic
			outToken := p.GetOtherToken(token)
			out = p.GetTokensOut(token, outToken.Addr, out)
			if outToken.IsTaxed {
				out.Mul(out, outToken.AfterTaxNum)
				out.Div(out, outToken.AfterTaxDen)

				//Tax it twice because we recive the token and send it again. Two transfers in total
				if i < len(path)-1 && !path[i+1].ReceiveDirectly() {
					out.Mul(out, outToken.AfterTaxNum)
					out.Div(out, outToken.AfterTaxDen)
				}
			}
			token = outToken.Addr
		}

		if out.Cmp(amountIn) > 0 {
			tempProfit := new(uint256.Int).Sub(out, amountIn)
			if tempProfit.Cmp(prevProfit) > 0 {
				prevProfit.Set(tempProfit)
				prevAmountIn.Set(amountIn)
				amountIn.Mul(amountIn, MUL_CONST)
				amountIn.Div(amountIn, DIV_CONST)
				if amountIn.Cmp(MAX_AMOUNT_IN) < 0 {
					profitIncreased = true
				}
			}
		}

		if profitIncreased {
			i--
		} else {
			if !prevProfit.IsZero() {
				setPath := make([]dex.Pair, 0, len(path))
				for _, p := range path {
					setPath = append(setPath, p)
				}
				f <- Found{amountIn: new(uint256.Int).Set(prevAmountIn), profit: new(uint256.Int).Set(prevProfit), path: setPath}
				//Reset the variables
			}
			prevAmountIn.Set(v)
			amountIn.Set(v)
			prevProfit.Clear()
		}

	}

}

func findTradingPaths(poolMap map[common.Address][]dex.Pair, tokenIn, tokenOut common.Address, heavyLimit, maxHops, hopsLimit int, path []dex.Pair, allPaths chan []dex.Pair, wg *sync.WaitGroup) {
	//TODO: Change the path algorithm to not create a new list everytime
	for _, pair := range poolMap[tokenIn] {
		repetition := false
		heavy := 0
		for i := 0; i <= hopsLimit-maxHops-1; i++ {
			if path[i].MarketAddress() == pair.MarketAddress() {
				repetition = true
				break
			}
			if path[i].HeavyComputation() {
				heavy++
			}
		}

		if pair.HeavyComputation() {
			heavy++
		}

		if repetition || heavy > heavyLimit {
			continue
		}
		path[hopsLimit-maxHops] = pair
		otherToken := pair.GetOtherToken(tokenIn).Addr
		if otherToken == tokenOut {
			newPath := make([]dex.Pair, 0, hopsLimit-maxHops)
			for i := 0; i <= hopsLimit-maxHops; i++ {
				newPath = append(newPath, path[i])
			}
			allPaths <- newPath
		} else if maxHops > 1 {
			wg.Add(1)
			inTok := otherToken
			outTok := tokenOut
			mHop := maxHops
			hopLim := hopsLimit
			copyPath := make([]dex.Pair, 0, len(path))
			for _, p := range path {
				copyPath = append(copyPath, p)
			}
			gorutineLimit <- true
			go func(iT common.Address, oT common.Address, mHop, hopLim int, p []dex.Pair) {
				findTradingPaths(poolMap, iT, oT, heavyLimit, mHop-1, hopLim, p, allPaths, wg)
				<-gorutineLimit
			}(inTok, outTok, mHop, hopLim, copyPath)
		}
	}
	wg.Done()
}
