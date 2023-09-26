package uniswapV2

import (
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/dex"
	"github.com/jmjac/ethGoTest/helperClient"
)

const BATCH_COUNT_LIMIT = 300
const UNISWAP_BATCH_SIZE = 1000
const UNISWAP_BATCH_SIZE_GAS_FREE = 20000
const UNISWAP_RESERVES_BATCH_SIZE = 1200
const UNISWAP_RESERVES_BATCH_SIZE_GAS_FREE = 6000
const GORUTINES_LIMIT = 200
const GORUTINES_LIMIT_GAS_FREE = 150

type Pair struct {
	address  common.Address
	token0   dex.Token
	token1   dex.Token
	reserve0 *uint256.Int
	reserve1 *uint256.Int
	protocol string
}

func newPair(client bind.ContractBackend, pairAddr common.Address, token0Addr, token1Addr common.Address, protocol string, tokenMap map[common.Address]dex.Token) Pair {
	t0, ok := tokenMap[token0Addr]
	if !ok {
		t0 = dex.NewNotTaxed(token0Addr)
	}

	t1, ok := tokenMap[token1Addr]
	if !ok {
		t1 = dex.NewNotTaxed(token1Addr)
	}
	pair := Pair{address: pairAddr, token0: t0, token1: t1, protocol: protocol}
	pair.reserve0 = new(uint256.Int)
	pair.reserve1 = new(uint256.Int)
	return pair
}

func GetUniswappyPairsWithoutLookup(client *helperClient.Client, addrFactory common.Address, protocol string, BLACKLIST map[common.Address]bool, tokenMap map[common.Address]dex.Token) []*Pair {
	now := time.Now()
	dex.Logger.Printf("=============== Setting up %v ===================\n", protocol)
	marketPairs := make([]*Pair, 0)
	factory, _ := NewUniswapV2FactoryCaller(addrFactory, client)
	numPairs, _ := factory.AllPairsLength(nil)

	for i := int64(0); i < numPairs.Int64(); i += UNISWAP_BATCH_SIZE_GAS_FREE {

		//pairs, err := lookup.GetPairsByIndexRange(nil, addrFactory, big.NewInt(i), big.NewInt(i+UNISWAP_BATCH_SIZE))
		end := i + UNISWAP_BATCH_SIZE_GAS_FREE
		if end > numPairs.Int64() {
			end = numPairs.Int64()
		}
		pairs, err := injectedLookupPairsByIndex(client, addrFactory, big.NewInt(i), big.NewInt(end))
		if err != nil {
			panic(err)
		}

		for _, p := range pairs {
			pair := newPair(client, p[2], p[0], p[1], protocol, tokenMap)
			marketPairs = append(marketPairs, &pair)
		}
	}
	UpdateReservesWithoutLookup(marketPairs, client)
	dex.Logger.Printf("Loaded %v pools from %v\n", len(marketPairs), protocol)
	dex.Logger.Printf("Took: %v\n\n", time.Since(now))
	return marketPairs
}

func UpdateReservesWithoutLookup(pairs []*Pair, client *helperClient.Client) {
	//lookup, _ := NewUniswapV2Query(addrUniswapLookup, client)
	pairAddresses := make([]common.Address, 0, len(pairs))

	guard := make(chan struct{}, GORUTINES_LIMIT_GAS_FREE)
	var wg sync.WaitGroup

	for _, pair := range pairs {
		pairAddresses = append(pairAddresses, pair.address)
	}

	for slice := 0; slice < len(pairAddresses); slice += UNISWAP_RESERVES_BATCH_SIZE_GAS_FREE {
		guard <- struct{}{}
		wg.Add(1)
		slice := slice //Make sure varaible is not overwritten before the start of gorutine
		go func(start int) {
			end := start + UNISWAP_RESERVES_BATCH_SIZE_GAS_FREE
			if end > len(pairAddresses) {
				end = len(pairAddresses)
			}

			//TODO: Add error handling
			//reserves, _ := lookup.GetReservesByPairs(nil, pairAddresses[start:end])
			var err error
			var reserves [][3]*big.Int

			reserves, err = injectedReservesByPairs(pairAddresses[start:end], client)
			count := 0
			for err != nil {
				count++
				time.Sleep(time.Millisecond * 5)
				reserves, err = injectedReservesByPairs(pairAddresses[start:end], client)
				if count > 3 {
					panic(err)
				}
			}
			for i := start; i < end; i++ {
				r0, _ := uint256.FromBig(reserves[i%UNISWAP_RESERVES_BATCH_SIZE_GAS_FREE][0])
				r1, _ := uint256.FromBig(reserves[i%UNISWAP_RESERVES_BATCH_SIZE_GAS_FREE][1])
				pairs[i].reserve0.Set(r0)
				pairs[i].reserve1.Set(r1)
			}

			<-guard
			wg.Done()
		}(slice)

	}
	wg.Wait()
	close(guard)
	return
}

func GetUniswappyPairs(client bind.ContractBackend, addrUniswapLookup common.Address, addrFactoryV2 common.Address, protocol string, BLACKLIST map[common.Address]bool, tokenMap map[common.Address]dex.Token) []*Pair {
	now := time.Now()
	dex.Logger.Printf("=============== Setting up %v ===================\n", protocol)

	lookup, _ := NewUniswapV2Query(addrUniswapLookup, client)
	marketPairs := make([]*Pair, 0)

	var i int64
	for ; i < BATCH_COUNT_LIMIT*UNISWAP_BATCH_SIZE; i += UNISWAP_BATCH_SIZE {
		pairs, err := lookup.GetPairsByIndexRange(nil, addrFactoryV2, big.NewInt(i), big.NewInt(i+UNISWAP_BATCH_SIZE))
		if err != nil {
			panic(err)
		}

		for _, pair := range pairs {
			addrMarket := pair[2]
			if _, ok := BLACKLIST[pair[0]]; ok {
				continue
			}
			if _, ok := BLACKLIST[pair[1]]; ok {
				continue
			}

			uniswapV2Pair := newPair(client, addrMarket, pair[0], pair[1], protocol, tokenMap)
			marketPairs = append(marketPairs, &uniswapV2Pair)
		}

		if len(pairs) < UNISWAP_BATCH_SIZE {
			break
		}
	}
	UpdateReserves(client, addrUniswapLookup, marketPairs)
	dex.Logger.Printf("Loaded %v pools from %v\n", len(marketPairs), protocol)
	dex.Logger.Printf("Took: %v\n\n", time.Since(now))
	return marketPairs
}

func UpdateReserves(client bind.ContractBackend, addrUniswapLookup common.Address, allMarketPairs []*Pair) {
	lookup, _ := NewUniswapV2Query(addrUniswapLookup, client)
	pairAddresses := make([]common.Address, 0, len(allMarketPairs))

	guard := make(chan struct{}, GORUTINES_LIMIT)
	var wg sync.WaitGroup

	for _, pair := range allMarketPairs {
		pairAddresses = append(pairAddresses, pair.address)
	}

	for slice := 0; slice < len(pairAddresses); slice += UNISWAP_RESERVES_BATCH_SIZE {
		guard <- struct{}{}
		wg.Add(1)

		slice := slice //Don't overwritte the variable
		go func(start int) {
			end := start + UNISWAP_RESERVES_BATCH_SIZE
			if end > len(allMarketPairs) {
				end = len(allMarketPairs)
			}

			//TODO: Add error handling
			reserves, err := lookup.GetReservesByPairs(nil, pairAddresses[start:end])
			if err != nil {
				panic(err)
			}
			for i := start; i < start+UNISWAP_RESERVES_BATCH_SIZE; i++ {
				if i >= len(allMarketPairs) {
					break
				}
				r0, _ := uint256.FromBig(reserves[i%UNISWAP_RESERVES_BATCH_SIZE][0])
				r1, _ := uint256.FromBig(reserves[i%UNISWAP_RESERVES_BATCH_SIZE][1])
				allMarketPairs[i].reserve0.Set(r0)
				allMarketPairs[i].reserve1.Set(r1)
			}

			<-guard
			wg.Done()
		}(slice)

	}
	wg.Wait()
	close(guard)
	return
}

func (pair Pair) tokensToReserves(tokenIn common.Address, tokenOut common.Address) (*uint256.Int, *uint256.Int) {
	if tokenIn == pair.token0.Addr {
		return new(uint256.Int).Set(pair.reserve0), new(uint256.Int).Set(pair.reserve1)
	} else {
		return new(uint256.Int).Set(pair.reserve1), new(uint256.Int).Set(pair.reserve0)
	}
}

func (pair Pair) GetTokensOut(tokenIn common.Address, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int {
	reserveIn, reserveOut := pair.tokensToReserves(tokenIn, tokenOut)
	if amountIn.IsZero() || reserveIn.IsZero() || reserveOut.IsZero() {
		return uint256.NewInt(0)
	}
	//Benchmarked, it's slighlty faster when inlined. Go compiler dosen't fix this
	//return pair.getAmountOut(reserveIn, reserveOut, amountIn)
	amountInWithFee := uint256.NewInt(997)
	amountInWithFee.Mul(amountInWithFee, amountIn)
	denominator := uint256.NewInt(1000)
	denominator.Mul(reserveIn, denominator)
	denominator.Add(denominator, amountInWithFee)
	numerator := amountInWithFee.Mul(amountInWithFee, reserveOut)
	return numerator.Div(numerator, denominator)
}

// WARNING: No longer used
func (pair Pair) getAmountOut(reserveIn *uint256.Int, reserveOut *uint256.Int, amountIn *uint256.Int) *uint256.Int {
	//	amountInWithFee := new(uint256.Int).Mul(amountIn, uint256.NewInt(997))
	//	numerator := new(uint256.Int).Mul(amountInWithFee, reserveOut)
	//	denominator := new(uint256.Int).Mul(reserveIn, uint256.NewInt(1000))
	//	denominator.Add(denominator, amountInWithFee)
	//TODO: Check if this is faster
	amountInWithFee := uint256.NewInt(997)
	amountInWithFee.Mul(amountInWithFee, amountIn)
	denominator := uint256.NewInt(1000)
	denominator.Mul(reserveIn, denominator)
	denominator.Add(denominator, amountInWithFee)
	numerator := amountInWithFee.Mul(amountInWithFee, reserveOut)
	return numerator.Div(numerator, denominator)
}

// TODO: Implement
func SellTokensToNextMarket(tokenIn common.Address, amountIn *uint256.Int, ethMarketAddress common.Address) []byte {
	return nil
}

func (pair Pair) SellTokens(tokenIn common.Address, amountIn *uint256.Int, recipient, sender common.Address) ([]common.Address, [][]byte) {
	//dex.Logger.Printf("UniswappyV2 %v Prepared swap for %v\n", pair.MarketAddress(), amountIn)
	targets := make([]common.Address, 0)
	payloads := make([][]byte, 0)
	amount0Out := big.NewInt(0)
	amount1Out := big.NewInt(0)

	var tokenOut common.Address
	if tokenIn == pair.token0.Addr {
		tokenOut = pair.token1.Addr
		amount1Out = pair.GetTokensOut(tokenIn, tokenOut, amountIn).ToBig()
	} else {
		tokenOut = pair.token0.Addr
		amount0Out = pair.GetTokensOut(tokenIn, tokenOut, amountIn).ToBig()
	}

	//Not sure if this is the right way to do it
	parsed, _ := abi.JSON(strings.NewReader(UniswapV2PairABI))
	txData, _ := parsed.Pack("swap", amount0Out, amount1Out, recipient, new([]byte))

	targets = append(targets, pair.MarketAddress())
	payloads = append(payloads, txData)

	return targets, payloads
}

func (pair Pair) ReceiveDirectly() bool {
	return true
}

func (pair Pair) PrepareReceive(tokenAddress common.Address, amountIn *uint256.Int) ([]common.Address, [][]byte) {
	panic("Not implemented")
	return nil, nil
}

//	func (pair *UniswapV2Pair) updatetReserves(token common.Address, balance *uint256.Int) {
//	 pair.balancesMap[token] = balance
//	}
func (pair Pair) GetOtherToken(knowAddres common.Address) dex.Token {
	if pair.token0.Addr == knowAddres {
		return pair.token1
	} else {
		return pair.token0
	}

}

func (pair Pair) MarketAddress() common.Address {
	return pair.address
}

func (pair Pair) SpenderAddress() common.Address {
	panic("UniswapV2 doesn't use spender address")
	return common.HexToAddress("")
}

func (pair Pair) Tokens() (dex.Token, dex.Token) {
	return pair.token0, pair.token1
}

func (pair Pair) Protocol() string {
	return pair.protocol
}

func (pair Pair) Reserves() (*uint256.Int, *uint256.Int) {
	r0 := pair.reserve0
	r1 := pair.reserve1
	return new(uint256.Int).Set(r0), new(uint256.Int).Set(r1)

}

func (p Pair) HeavyComputation() bool {
	return false
}
