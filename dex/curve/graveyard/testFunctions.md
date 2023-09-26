package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/arbitrage"
	"github.com/jmjac/ethGoTest/curve"
	"github.com/jmjac/ethGoTest/curve/factory"
	"github.com/jmjac/ethGoTest/curve/reg1"
	"github.com/jmjac/ethGoTest/curve/swap"
	"github.com/jmjac/ethGoTest/curve/vyper/get_D"
	get_dy "github.com/jmjac/ethGoTest/curve/vyper/get_Dy"
	get_Y "github.com/jmjac/ethGoTest/curve/vyper/get_y"
	"github.com/jmjac/ethGoTest/curve/vyper/xp"
)

func main() {
	mainnet_url := "https://xyznode.space/api/mainnet/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"
	mainnet_client, err := ethclient.Dial(mainnet_url)

	if err != nil {
		panic(err)
	}

	goerli_url := "https://xyznode.space/api/goerli/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"
	goerli_client, err := ethclient.Dial(goerli_url)

	if err != nil {
		panic(err)
	}

	// Running Pool Stat Function
	// poolStats(mainnet_client)

	ADDR.MainnetAddresses()

	BLACKLIST := arbitrage.InitBlackList("blacklist.json")
	fmt.Printf("Blacklist has %v tokens\n", len(BLACKLIST))

	pairsCurve := curve.GetCurveMarkets(mainnet_client, ADDR.CURVE_ADDRESS_PROVIDER, ADDR.CURVE_REGISTRY, "curve-finance", BLACKLIST)

	initial_count := 0
	initial_total := 0
	test_swap_count := 0
	test_swap_total := 0
	main_swap_count := 0
	main_swap_total := 0
	all_count := 0
	all_total := 0
	for i := 0; i < /* len(pairsCurve) */ 1; i++ {
		// Set up params
		p := pairsCurve[0]
		address := p.MarketAddress()
		tokens := p.Tokens()
		reserves := p.Reserves()
		N_COINS := p.N_COINS
		RATES := p.Rates()

		// These parameters control all the trading
		i := 1
		j := 0
		dx := big.NewInt(1000000000000000000)
		trade_amount := uint256.NewInt(1)

		big_reserves := [8]*big.Int{}
		for i := uint64(0); i < 8; i++ {
			if N_COINS.Cmp(uint256.NewInt(i)) > 0 {
				if i != 0 {
					sketchy_reserve := new(uint256.Int).Mul(reserves[i], uint256.NewInt(1000000000000))
					big_reserves[i] = sketchy_reserve.ToBig()
				} else {
					big_reserves[i] = reserves[i].ToBig()
				}

			} else {
				big_reserves[i] = big.NewInt(0)
			}
		}

		big_rates := [8]*big.Int{}
		for i := uint64(0); i < 8; i++ {
			if N_COINS.Cmp(uint256.NewInt(i)) > 0 {
				big_rates[i] = RATES[i].ToBig()
			} else {
				big_rates[i] = big.NewInt(0)
			}
		}

		// Testing functions
		built_xp, goerli_xp := getXpTest(goerli_client, p, big_reserves, big_rates)
		var check_xp bool
		for i := 0; i < int(N_COINS.Uint64()); i++ {
			initial_total += 1
			if built_xp[i].Uint64() == goerli_xp[i].Uint64() {
				check_xp = true
				initial_count += 1
			} else {
				check_xp = false
			}
		}
		built_d, goerli_d := getDTest(goerli_client, p, built_xp, goerli_xp)
		var check_d bool
		initial_total += 1
		if built_d.Uint64() == goerli_d.Uint64() {
			check_d = true
			initial_count += 1
		} else {
			check_d = false
		}

		goerli_x := new(big.Int).Add(goerli_xp[i], dx)
		built_x := new(uint256.Int).Add(built_xp[j], uint256.NewInt(dx.Uint64()))
		var check_x bool
		initial_total += 1
		if built_x.Uint64() == goerli_x.Uint64() {
			check_x = true
			initial_count += 1
		} else {
			check_x = false
		}

		built_y, goerli_y := getYTest(goerli_client, p, big.NewInt(int64(i)), big.NewInt(int64(j)), built_xp, goerli_xp, built_x, goerli_x, goerli_d)
		var check_y bool
		initial_total += 1
		if built_y.Uint64() == goerli_y.Uint64() {
			check_y = true
			initial_count += 1
		} else {
			check_y = false
		}

		built_dy, goerli_dy := getDyTest(goerli_client, p, big.NewInt(int64(i)), big.NewInt(int64(j)), tokens[1], tokens[0], dx, big_rates, goerli_xp, goerli_y)
		main_dy := mainnetTest(mainnet_client, address, tokens, trade_amount)

		var check_dy_test bool
		test_swap_total += 1
		if built_dy.Uint64() == goerli_dy.Uint64() {
			check_dy_test = true
			test_swap_count += 1
		} else {
			check_dy_test = false
		}

		var check_dy_main bool
		main_swap_total += 1
		if built_dy.Uint64() == main_dy.Uint64() {
			check_dy_main = true
			main_swap_count += 1
		} else {
			check_dy_main = false
		}

		var check_all bool
		all_total += 1
		if built_dy.Uint64() == main_dy.Uint64() && built_dy.Uint64() == goerli_dy.Uint64() {
			check_all = true
			all_count += 1
		} else {
			check_all = false
		}

		fmt.Println("===================== INITAL INPUT VALUES ====================")
		fmt.Printf("Swaps from %v pool \n", address)
		fmt.Printf("Swaps from %v ----> %v \n", tokens[1], tokens[0])
		fmt.Printf("rates: %v \n", big_rates)
		fmt.Printf("Underlying & Meta: %v %v \n", p.IS_UNDERLYING, p.IS_META)
		fmt.Println("===================== SWAP RETURN VALUES ====================")
		fmt.Printf("XP built vs goerli [%v, %v] outcome: %v \n", built_xp, goerli_xp, check_xp)
		fmt.Printf("D built vs goerli [%v, %v] outcome: %v\n", built_d, goerli_d, check_d)
		fmt.Printf("X built vs goerli [%v, %v] outcome: %v\n", built_x, goerli_x, check_x)
		fmt.Printf("Y built vs goerli [%v, %v] outcome: %v\n", built_y, goerli_y, check_y)
		fmt.Printf("DY built vs goerli vs main [%v, %v, %v] \n", built_dy, goerli_dy, main_dy)
		fmt.Printf("Outcome of Dy test [%v, %v, %v] \n", check_dy_test, check_dy_main, check_all)
		fmt.Println("")

	}
	fmt.Println("===================== TESTING STATS ====================")
	fmt.Printf("Initial True to Total %v / %v \n", initial_count, initial_total)
	percentage := float64(initial_count) / float64(initial_total)
	fmt.Printf("Percentage of Trues: %v \n", percentage)
	fmt.Printf("Test Swap True to Total %v / %v \n", test_swap_count, test_swap_total)
	percentage1 := float64(test_swap_count) / float64(test_swap_total)
	fmt.Printf("Percentage of Trues: %v \n", percentage1)
	fmt.Printf("Main Swap True to Total %v / %v \n", main_swap_count, main_swap_total)
	percentage2 := float64(main_swap_count) / float64(main_swap_total)
	fmt.Printf("Percentage of Trues: %v \n", percentage2)
	fmt.Printf("All Swap True to Total %v / %v \n", all_count, all_total)
	percentage3 := float64(all_count) / float64(all_total)
	fmt.Printf("Percentage of Trues: %v \n", percentage3)
}

func mainnetTest(forked_client *ethclient.Client, address common.Address, tokens []common.Address, trade_amount *uint256.Int) *big.Int {

	swap_address := common.HexToAddress("0x55B916Ce078eA594c10a874ba67eCc3d62e29822")
	swap, _ := swap.NewSwapCaller(swap_address, forked_client)

	main, e := swap.GetExchangeAmount(nil, address, tokens[1], tokens[0], trade_amount.ToBig())

	if e != nil {
		fmt.Println("Mainnet transaction reverted")
		return big.NewInt(0)
	} else {
		return main
	}

}

func getXpTest(goerli_client *ethclient.Client, p curve.Pair, big_reserves, big_rates [8]*big.Int) ([]*uint256.Int, [8]*big.Int) {
	getXp_address := common.HexToAddress("0x1a170f097eD348c8393b6DD5aC177dE83528f161")
	getxp, _ := xp.NewXpCaller(getXp_address, goerli_client)
	goerli_XP, _ := getxp.Xp(nil, big_reserves, p.N_COINS.ToBig(), big_rates, p.LENDING_PRECISION.ToBig())
	built_XP := p.Xp(p.Rates(), p.Reserves(), p.LENDING_PRECISION)

	return built_XP, goerli_XP
}

func getDTest(goerli_client *ethclient.Client, p curve.Pair, built_xp []*uint256.Int, goerli_xp [8]*big.Int) (*uint256.Int, *big.Int) {
	getD_address := common.HexToAddress("0xb8209ad70fd58E30235bbb0AAa075381DCc50462")
	getD, _ := get_D.NewGetDCaller(getD_address, goerli_client)
	goerli_D, _ := getD.GetD(nil, goerli_xp, p.A.ToBig(), p.N_COINS.ToBig())
	built_D := p.Get_D(built_xp, p.A)

	return built_D, goerli_D
}

func getYTest(goerli_client *ethclient.Client, p curve.Pair, i, j *big.Int, built_xp []*uint256.Int, goerli_xp [8]*big.Int, built_x *uint256.Int, goerli_x *big.Int, goerli_d *big.Int) (*uint256.Int, *big.Int) {
	getY_address := common.HexToAddress("0x04Ba145228A7229Bf98CD3982168e950432e758F")
	getY, _ := get_Y.NewGetYCaller(getY_address, goerli_client)
	goerli_Y, _ := getY.GetY(nil, i, j, goerli_x, goerli_xp, p.A.ToBig(), p.N_COINS.ToBig(), goerli_d)
	built_Y := p.Get_Y(i.Uint64(), j.Uint64(), built_x, built_xp)

	return built_Y, goerli_Y
}

func getDyTest(goerli_client *ethclient.Client, p curve.Pair, i, j *big.Int, tokenIn, tokenOut common.Address, dx *big.Int, big_rates [8]*big.Int, goerli_xp [8]*big.Int, goerli_y *big.Int) (*uint256.Int, *big.Int) {
	getDy_address := common.HexToAddress("0xB7463425C6Cbc412af0DA97530a6BE28e9C85BAa")
	getDy, _ := get_dy.NewGetDyCaller(getDy_address, goerli_client)
	goerli_DY, _ := getDy.GetDy(nil, i, j, dx, goerli_xp, goerli_y, big_rates, p.N_COINS.ToBig(), p.PRECISION.ToBig(), p.FEE_DENOMINATOR.ToBig(), p.Fee().ToBig())
	built_DY := p.GetTokensOut(tokenIn, tokenOut, uint256.NewInt(dx.Uint64()))

	if goerli_y.Cmp(big.NewInt(0)) == 0 {
		return built_DY, big.NewInt(0)
	} else {
		return built_DY, goerli_DY
	}
}

func poolStats(mainnet_client *ethclient.Client) {
	factory_address := common.HexToAddress("0xB9fC157394Af804a3578134A6585C0dc9cc990d4")
	factory, _ := factory.NewFactoryCaller(factory_address, mainnet_client)

	registry_address := common.HexToAddress("0x90E00ACe148ca3b23Ac1bC8C240C2a7Dd9c2d7f5")
	registry, _ := reg1.NewReg1Caller(registry_address, mainnet_client)

	factory_pools, _ := factory.PoolCount(nil)
	registry_pools, _ := registry.PoolCount(nil)
	total_pools := new(big.Int).Add(factory_pools, registry_pools)

	var i int64
	meta_count := 0
	base_count := 0
	unsorted_implementations := 0
	meta_count_map := make(map[common.Address]int)
	base_count_map := make(map[common.Address]int)

	for i = 0; i < factory_pools.Int64(); i++ {
		pool_address, _ := factory.PoolList(nil, big.NewInt(i))
		implementation_address, _ := factory.GetImplementationAddress(nil, pool_address)
		is_meta, _ := factory.IsMeta(nil, pool_address)

		if is_meta {
			meta_count += 1
			meta_count_map[implementation_address] += 1
		} else {
			base_count += 1
			base_count_map[implementation_address] += 1
		}

		fmt.Println("================ POOL and IMPLEMENTATION ==================")
		fmt.Printf("Pool Adress: %v \n", pool_address)
		fmt.Printf("Implementation Adress: %v \n", implementation_address)
		fmt.Printf("Meta: %v \n", is_meta)
	}

	for i = 0; i < registry_pools.Int64(); i++ {
		pool_address, _ := registry.PoolList(nil, big.NewInt(i))
		is_meta, _ := registry.IsMeta(nil, pool_address)

		if is_meta {
			meta_count += 1
			implementation_address, _ := factory.GetImplementationAddress(nil, pool_address)
			if implementation_address == common.HexToAddress("0x0000000000000000000000000000000000000000") {
				unsorted_implementations += 1
			} else {
				meta_count_map[implementation_address] += 1
			}
			fmt.Println("================ POOL and IMPLEMENTATION ==================")
			fmt.Printf("Pool Adress: %v \n", pool_address)
			fmt.Printf("Implementation Adress: %v \n", implementation_address)
			fmt.Printf("Meta: %v \n", is_meta)
		} else {
			base_count += 1
			implementation_address, _ := factory.GetImplementationAddress(nil, pool_address)
			if implementation_address == common.HexToAddress("0x0000000000000000000000000000000000000000") {
				unsorted_implementations += 1
			} else {
				base_count_map[implementation_address] += 1
			}
			fmt.Println("================ POOL and IMPLEMENTATION ==================")
			fmt.Printf("Pool Adress: %v \n", pool_address)
			fmt.Printf("Implementation Adress: %v \n", implementation_address)
			fmt.Printf("Meta: %v \n", is_meta)
		}
	}

	fmt.Println("================ POOL STATS ==================")
	fmt.Printf("Factory Pool Count: %v \n", factory_pools)
	fmt.Printf("Registry Pool Count: %v \n", registry_pools)
	fmt.Printf("Total Pool Count: %v \n", total_pools)
	fmt.Printf("Base Pool Count: %v \n", base_count)
	fmt.Printf("Meta Pool Count: %v \n", meta_count)
	fmt.Printf("Unsorted Pool Count: %v \n", unsorted_implementations)
	fmt.Printf("Meta Pool Map: %v \n", meta_count_map)
	fmt.Printf("Base Pool Map: %v \n", base_count_map)
	fmt.Printf("Total Meta Implementations: %v \n", len(meta_count_map))
	fmt.Printf("Total Base Implementations: %v \n", len(base_count_map))
}
