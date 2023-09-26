package main

import (
	"fmt"
	"log"
	"math/big"
	"sort"

	core "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/arbitrage"
	"github.com/jmjac/ethGoTest/dex"
	"github.com/jmjac/ethGoTest/dex/uniswapV3"
	"github.com/jmjac/ethGoTest/helperClient"
)

//Utility to sync univ3 pool data from subgraph
func main() {
	provider := "https://xyznode.space/api/mainnet/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"
	mainnet := true
	client, err := helperClient.New(provider)
	if err != nil {
		panic(err)
	}

	BLACKLIST := arbitrage.InitBlackList("blacklist.json")
	//BLACKLIST := make(map[common.Address]bool) //arbitrage.InitBlackList("blacklist.json")
	fmt.Printf("Blacklist has %v tokens\n", len(BLACKLIST))
	TOKEN_DECIMALS := arbitrage.LoadDecimalsFromJson("token_decimals.json")

	if mainnet {
		ADDR.MainnetAddresses()
	} else {
		ADDR.GoerliAddresses()
	}

	//pools := uniswapV2.GetUniswappyPairs(client, ADDR.UNISWAPv2_LOOKUP, ADDR.UNIv2_FACTORY, "uniswapV2", make(map[common.Address]bool))
	blockNum := uint64(15937666 - 50) //client.BlockNumber(context.Background())
	//Maybe add some -50 or something
	pairs := uniswapV3.InitialSyncWithSubgraph(client.Client, ADDR.UNIv3_FACTORY, BLACKLIST, blockNum, make(map[common.Address]dex.Token))
	tar1 := common.HexToAddress("0x090185f2135308BaD17527004364eBcC2D37e5F6")
	tar2 := common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	var pool *uniswapV3.Pair
	for _, p := range pairs {
		t0, t1 := p.Tokens()
		if t0.Addr == tar1 && t1.Addr == tar2 || t0.Addr == tar2 && t1.Addr == tar1 {
			pool = p
			break
		}
	}

	failedBlock := uint64(15938875 - 2)
	for i := blockNum; i <= failedBlock; i++ {
		//Maybe should update all reserves
		pool.UpdateReserves(big.NewInt(int64(blockNum)), client)
	}

	in, _ := uint256.FromHex("0x3D1AA2A1B2F1D9A6E10D")
	fmt.Printf("Got out %v when estimated at %v\n", pool.GetTokensOut(tar1, tar2, in), failedBlock)
	swapWithSdk(pool, TOKEN_DECIMALS, in.ToBig())
	failedBlock++
	pool.UpdateReserves(big.NewInt(int64(failedBlock)), client)
	fmt.Printf("Got out %v when estimated at %v\n", pool.GetTokensOut(tar1, tar2, in), failedBlock)
	swapWithSdk(pool, TOKEN_DECIMALS, in.ToBig())
	failedBlock++
	pool.UpdateReserves(big.NewInt(int64(failedBlock)), client)
	fmt.Printf("Got out %v when estimated at %v\n", pool.GetTokensOut(tar1, tar2, in), failedBlock)
	swapWithSdk(pool, TOKEN_DECIMALS, in.ToBig())
	failedBlock++
	pool.UpdateReserves(big.NewInt(int64(failedBlock)), client)
	fmt.Printf("Got out %v when estimated at %v\n", pool.GetTokensOut(tar1, tar2, in), failedBlock)
	swapWithSdk(pool, TOKEN_DECIMALS, in.ToBig())

	//uniswapV2.UpdateReservesWithoutLookup(pools, client)

}

func swapWithSdk(pair *uniswapV3.Pair, TOKEN_DECIMALS map[common.Address]uint64, inP *big.Int) {
	ticks := make([]entities.Tick, 0)
	for index, t := range pair.Ticks {
		g := t.LiquidityGross.ToBig()
		n := new(big.Int).Set(&t.LiquidityNet)
		ticks = append(ticks, entities.Tick{Index: index, LiquidityGross: g, LiquidityNet: n})
		//fmt.Printf("Tick: %v\n\tGross: %v\n\tNet: %v\n", index, g, n)
	}
	sort.Slice(ticks, func(i, j int) bool {
		return ticks[i].Index < ticks[j].Index
	})

	fee := constants.FeeAmount(pair.Fee.Uint64())
	p, err := entities.NewTickListDataProvider(ticks, constants.TickSpacings[fee])
	if err != nil {
		log.Fatal(err)
	}
	d0 := uint(TOKEN_DECIMALS[pair.Token0.Addr])
	t0 := core.NewToken(1, pair.Token0.Addr, d0, "t0", "t0")
	d1 := uint(TOKEN_DECIMALS[pair.Token1.Addr])
	t1 := core.NewToken(1, pair.Token1.Addr, d1, "t1", "t1")
	pool, _ := entities.NewPool(t0, t1, fee, pair.SqrtPriceX96.ToBig(), pair.Liquidity.ToBig(), pair.Tick, p)

	in := core.FromRawAmount(t0, new(big.Int).Set(inP))
	v1, _, _ := pool.GetOutputAmount(in, nil)
	//v1u, _ := uint256.FromBig(v1.Numerator)
	//out1 := arbitrage.FormatUintToFloat(v1u, 1/math.Pow(10, float64(d1)))

	//v2u, _ := uint256.FromBig(v2.Numerator)
	//out2 := arbitrage.FormatUintToFloat(v2u, 1/math.Pow(10, float64(d0)))

	fmt.Printf("For %v got %v of %v\n", in.ToExact(), v1.ToExact(), pair.Token1.Addr)
	fmt.Printf("For %v got %v of %v\n", in.Numerator, v1.Numerator, pair.Token1.Addr)
}
