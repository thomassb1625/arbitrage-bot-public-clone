package uniswapV3

import (
	"context"
	"errors"
	"log"
	"math/big"
	"sort"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/abis"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/dex"
)

const maxGoroutines = 100

// Pair represents an UniswapV3 pair representation
type Pair struct {
	Address      common.Address
	Token0       dex.Token
	Token1       dex.Token
	ProtocolName string
	// init
	TickSpacing int         // int24 tickSpacing - This value is an int24 to avoid casting even though it is always positive.
	Fee         uint256.Int // uint24 fee
	HackyRounds int         // hacky canceller for speed-boost in result of possible trade lost

	// state
	Tick         int          // int24 tick
	SqrtPriceX96 *uint256.Int // uint160 sqrtPriceX96
	Liquidity    *uint256.Int // uint128 liquidity

	// current Ticks in pool
	//tickList   []Tick
	Ticks    map[int]Tick
	TickList []Tick

	//mutex sync.RWMutex

	LastBlock uint64
}

func parseContractLogsForAddresses(client bind.ContractBackend, addrFactoryV3 common.Address, creationBlock *big.Int) []*Pair {
	pools := make([]*Pair, 0)

	query := ethereum.FilterQuery{
		FromBlock: creationBlock,
		Addresses: []common.Address{
			addrFactoryV3,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(UniswapFactoryV3ABI)))
	contract := bind.NewBoundContract(addrFactoryV3, contractAbi, client, client, client)

	for _, vLog := range logs {
		if vLog.Removed || len(vLog.Data) == 0 {
			continue
		}

		event := new(UniswapFactoryV3PoolCreated)
		err := contract.UnpackLog(event, "PoolCreated", vLog)
		if err != nil {
			panic(err)
		}
		fee := new(big.Int).Set(event.Fee)
		feeUint, _ := uint256.FromBig(fee)
		poolAddr := event.Pool

		pair := Pair{Address: poolAddr, Fee: *feeUint, TickSpacing: int(event.TickSpacing.Int64())}

		pools = append(pools, &pair)
	}
	return pools
}

var (
	BURN       = common.HexToHash("0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c")
	SWAP       = common.HexToHash("0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67")
	INITIALIZE = common.HexToHash("0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95")
	MINT       = common.HexToHash("0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde")
)

func UpdateReserves(pairs []*Pair, blockNum uint64, client bind.ContractBackend) error {
	abort := false
	limit := make(chan bool, 5)
	var wg sync.WaitGroup
	for i := range pairs {
		if abort {
			break
		}
		limit <- true
		wg.Add(1)

		pairNum := i
		go func(pn int) {
			err := pairs[pn].UpdateReserves(big.NewInt(int64(blockNum)), client)
			if err != nil {
				abort = true
			}
			<-limit
			wg.Done()
		}(pairNum)
	}
	wg.Wait()
	if abort {
		return errors.New("Update reserves failed")
	}
	return nil
}

func (p *Pair) UpdateReserves(block *big.Int, client bind.ContractBackend) error {
	if block.Uint64() < p.LastBlock {
		return nil
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(p.LastBlock)),
		ToBlock:   block,
		Addresses: []common.Address{
			p.Address,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Printf("Error with %v", p.Address)
		return errors.New("Timed out")
	}

	// Initialization in GetUniswappyPairs for small speedup
	//CONTRACT_ABI, _ = abi.JSON(strings.NewReader(string(UniswapPoolV3ABI)))
	//CONTRACT = bind.NewBoundContract(common.HexToAddress(""), CONTRACT_ABI, client, client, client)

	for _, vLog := range logs {
		if vLog.Removed {
			continue
		}
		switch vLog.Topics[0] {
		case INITIALIZE:
			event := new(UniswapPoolV3Initialize)
			if err := CONTRACT.UnpackLog(event, "Initialize", vLog); err != nil {
				panic(err)
			}
			p.Tick = int(event.Tick.Int64())
			p.SqrtPriceX96, _ = uint256.FromBig(event.SqrtPriceX96)
			p.Liquidity = uint256.NewInt(0)
			p.Ticks = make(map[int]Tick)
		case MINT:
			event := new(UniswapPoolV3Mint)
			CONTRACT.UnpackLog(event, "Mint", vLog)
			p.mint(int(event.TickLower.Int64()), int(event.TickUpper.Int64()), event.Amount)
		case BURN:
			event := new(UniswapPoolV3Burn)
			CONTRACT.UnpackLog(event, "Burn", vLog)
			p.burn(int(event.TickLower.Int64()), int(event.TickUpper.Int64()), event.Amount)
		case SWAP:
			event := new(UniswapPoolV3Swap)
			CONTRACT.UnpackLog(event, "Swap", vLog)
			p.Tick = int(event.Tick.Int64())
			//			l, _ := uint256.FromBig(event.Liquidity)
			//			p.Liquidity.Set(l)
			p.Liquidity.SetBytes(event.Liquidity.Bytes())
			//			sqrt, _ := uint256.FromBig(event.SqrtPriceX96)
			//			p.SqrtPriceX96.Set(sqrt)
			p.SqrtPriceX96.SetBytes(event.SqrtPriceX96.Bytes())
		}
	}
	p.UpdateTickList()

	//Update tickList for swap calulcations
	p.LastBlock = uint64(block.Int64() + 1)
	return nil
}

//Should be called
func (p *Pair) UpdateTickList() {
	updatedTicks := make([]Tick, 0)
	for k, v := range p.Ticks {
		v.Index = k
		updatedTicks = append(updatedTicks, v)
	}
	p.TickList = updatedTicks
	sort.Slice(p.TickList, func(i, j int) bool {
		return p.TickList[i].Index < p.TickList[j].Index
	})

}

func InitialSyncWithSubgraph(client bind.ContractBackend, factory common.Address, BLACKLIST map[common.Address]bool, blockNum uint64, tokenMap map[common.Address]dex.Token) []*Pair {
	uniswapV3Creation, _ := new(big.Int).SetString("0", 10) //To save setup time change this
	pairsUniV3 := GetUniswappyPairs(client, factory, uniswapV3Creation, "uniswapV3", BLACKLIST, blockNum, tokenMap)

	var wg sync.WaitGroup
	var listWg sync.WaitGroup
	pairChan := make(chan *Pair, 10)
	limit := make(chan bool, 16)

	failed := make([]*Pair, 0)
	allUniv3Pairs := make([]*Pair, 0)

	listWg.Add(1)
	go func() {
		for p := range pairChan {
			allUniv3Pairs = append(allUniv3Pairs, p)
			if p.MarketAddress() == common.HexToAddress("0x87986ae1e99f99da1f955d16930dc8914ffbed56") {
				dex.Logger.Println("Sanity check")
				t0, t1 := p.Tokens()
				//pool := &Pair{Address: p.Address, Token0Addr: t0, Token1Addr: t1, Fee: p.Fee, ProtocolName: p.ProtocolName, TickSpacing: p.TickSpacing}
				//t0, t1 = pool.Tokens()
				p.UpdateTickList()
				dex.Logger.Printf("%v and %v\n", t0.Addr, t1.Addr)
				dex.Logger.Println(p.GetTokensOut(t1.Addr, t0.Addr, uint256.NewInt(1e18)))
				dex.Logger.Println(p.GetTokensOut(t0.Addr, t1.Addr, uint256.NewInt(1e18)))
			}
		}
		listWg.Done()
	}()

	for i := range pairsUniV3 {
		limit <- true
		wg.Add(1)
		pairNum := i
		go func(pn int) {
			p := pairsUniV3[pn]
			tick, liquidity, sqrtPriceX96, ticks, err := GetStateFromSubgraph(p.Address, big.NewInt(int64(blockNum)))
			if err != nil {
				dex.Logger.Printf("Failed on: %v with err: %v\n", p.Address, err)
			} else {
				p.Tick = tick
				p.Liquidity = liquidity
				p.SqrtPriceX96 = sqrtPriceX96
				p.Ticks = ticks
				p.Ticks = ticks
				p.LastBlock = blockNum
				p.UpdateTickList()
				pairChan <- p
			}
			wg.Done()
			<-limit
		}(pairNum)
	}

	wg.Wait()
	close(pairChan)
	listWg.Wait()
	close(limit)

	for i, p := range failed {
		dex.Logger.Println(i)

		tick, liquidity, sqrtPriceX96, ticks, err := GetStateFromSubgraph(p.Address, big.NewInt(int64(blockNum)))
		if err != nil {
			dex.Logger.Println("Failed on: ", p.Address)
			continue
		} else {
			p.Tick = tick
			p.Liquidity = liquidity
			p.SqrtPriceX96 = sqrtPriceX96
			p.Ticks = ticks
			p.Ticks = ticks
			p.LastBlock = blockNum
			p.UpdateTickList()
			allUniv3Pairs = append(allUniv3Pairs, p)
		}
	}

	return allUniv3Pairs
}

var CONTRACT_ABI abi.ABI
var CONTRACT *bind.BoundContract

//Use token pairs from uniswapV2
func GetUniswappyPairs(client bind.ContractBackend, addrFactoryV3 common.Address, creationBlock *big.Int, protocol string, BLACKLIST map[common.Address]bool, blockNum uint64, tokenMap map[common.Address]dex.Token) []*Pair {
	dex.Logger.Println("=========== Setting up UniswapV3 ===============")
	//To avoid subgraph being out of sync
	blockNum -= 50

	//More overwhelmes the node maxGoroutines := 200
	guard := make(chan struct{}, maxGoroutines)
	var wg = sync.WaitGroup{}

	pairs := parseContractLogsForAddresses(client, addrFactoryV3, creationBlock)

	for i := range pairs {
		guard <- struct{}{}
		wg.Add(1)
		pairNum := i
		go func(pn int) {
			pool, _ := NewUniswapPoolV3Caller(pairs[pn].Address, client)
			t0Addr, _ := pool.Token0(nil)
			t1Addr, _ := pool.Token1(nil)
			t0, ok := tokenMap[t0Addr]
			if !ok {
				t0 = dex.NewNotTaxed(t0Addr)
			}
			pairs[pn].Token0 = t0

			t1, ok := tokenMap[t1Addr]
			if !ok {
				t1 = dex.NewNotTaxed(t1Addr)
			}
			pairs[pn].Token1 = t1
			pairs[pn].ProtocolName = protocol
			<-guard
			wg.Done()
		}(pairNum)

	}
	wg.Wait()
	close(guard)

	//Remove bllacklisted tokens
	fliteredPairs := make([]*Pair, 0)
	for _, p := range pairs {
		t0, t1 := p.Tokens()
		_, blackT0 := BLACKLIST[t0.Addr]
		_, blackT1 := BLACKLIST[t1.Addr]
		if blackT0 || blackT1 {
			continue
		}

		p.HackyRounds = 10
		fliteredPairs = append(fliteredPairs, p)
		////TODO: Remove this
		//if (i+1)%50 == 0 {
		//	break
		//}
	}

	//TODO: Do the reserves
	//UpdateReserves(client, addrUniswapLookup, fliteredPairs)

	CONTRACT_ABI, _ = abi.JSON(strings.NewReader(string(UniswapPoolV3ABI)))
	CONTRACT = bind.NewBoundContract(common.HexToAddress(""), CONTRACT_ABI, client, client, client)
	dex.Logger.Printf("Loaded %v pools from uniswapV3\n", len(fliteredPairs))
	return fliteredPairs
}

func (p Pair) GetTokensIn(tokenIn common.Address, tokenOut common.Address, amountOut *uint256.Int) *uint256.Int {
	panic("Not implemented")
	return nil
}

//INFO: Implemented in temporary
//func (p Pair) GetTokensOut(tokenIn common.Address, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int {
//return p.GetAmountOut(amountIn, tokenIn == p.Token0Addr)
//}

var approveAmount, _ = new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 10)
var deadline, _ = new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 10)

func (p Pair) SellTokens(tokenIn common.Address, amountIn *uint256.Int, recipient, sender common.Address) ([]common.Address, [][]byte) {
	targets := make([]common.Address, 0)
	payloads := make([][]byte, 0)

	IERC20ABI, _ := abi.JSON(strings.NewReader(abis.IERC20ABI))
	approveData, _ := IERC20ABI.Pack("approve", ADDR.UNISWAPV3_ROUTER, approveAmount)
	targets = append(targets, tokenIn)
	payloads = append(payloads, approveData)

	var tokenOut common.Address
	if p.Token0.Addr == tokenIn {
		tokenOut = p.Token1.Addr
	} else {
		tokenOut = p.Token0.Addr
	}

	params := ISwapRouterExactInputSingleParams{TokenIn: tokenIn, TokenOut: tokenOut, Fee: p.Fee.ToBig(), Recipient: recipient, Deadline: deadline, AmountIn: amountIn.ToBig(), AmountOutMinimum: big.NewInt(0), SqrtPriceLimitX96: big.NewInt(0)}
	parsed, _ := abi.JSON(strings.NewReader(SwapRouterABI))
	txData, err := parsed.Pack("exactInputSingle", params)
	if err != nil {
		panic(err)
	}
	targets = append(targets, ADDR.UNISWAPV3_ROUTER)
	payloads = append(payloads, txData)
	return targets, payloads
}

func (p Pair) ReceiveDirectly() bool {
	return false
}

func (p Pair) PrepareReceive(tokenAddress common.Address, amountIn *uint256.Int) ([]common.Address, [][]byte) {
	panic("Not implemented")
	return nil, nil
}

func (p Pair) MarketAddress() common.Address {
	return p.Address
}

func (p Pair) SpenderAddress() common.Address {
	return ADDR.UNISWAPV3_ROUTER
}

func (p Pair) Tokens() (dex.Token, dex.Token) {
	return p.Token0, p.Token1
}

func (p Pair) Protocol() string {
	return p.ProtocolName
}

func (p Pair) GetOtherToken(token common.Address) dex.Token {
	if token == p.Token0.Addr {
		return p.Token1
	}
	return p.Token0
}

//THIS SHOULD RETURN LIQUIDITY IN CASE OF UNIV3 Or always some high number if we want V3 to always be included
func (p Pair) Reserves() (*uint256.Int, *uint256.Int) {
	return new(uint256.Int).Set(p.Liquidity), new(uint256.Int).Set(p.Liquidity)
}

func (p Pair) HeavyComputation() bool {
	return true
}
