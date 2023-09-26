package curve

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	ADDR "github.com/jmjac/ethGoTest/addr"
)

var (
	main_swap_count = 0
	main_swap_total = 0
)

func TestStats3Pool(client bind.ContractBackend, CURVE_REGISTRY common.Address) {
	marketPairs := GetCurveMarkets(client, ADDR.CURVE_ADDRESS_PROVIDER, CURVE_REGISTRY, ADDR.CURVE_FACTORY, "curve-finance")

	var x int64
	var i uint64
	var j uint64
	for x = 1; x < 100; x++ {
		for i = 0; i < 3; i++ {
			for j = 0; j < 3; j++ {
				if i != j {
					dx := uint256.NewInt(uint64(x))
					trade_amount := big.NewInt(x)
					testFunction(client, marketPairs[0], i, j, dx, trade_amount)
				}
			}
		}
	}

	fmt.Println("===================== TESTING STATS ====================")
	fmt.Printf("Main Swap True to Total %v / %v \n", main_swap_count, main_swap_total)
	percentage := float64(main_swap_count) / float64(main_swap_total)
	fmt.Printf("Percentage of Trues: %v \n", percentage)
}

func testFunction(client bind.ContractBackend, pair Pair, i, j uint64, dx *uint256.Int, trade_amount *big.Int) {
	// Set up params
	address := pair.MarketAddress()
	tokens := pair.Tokens()
	rates := pair.Rates()
	now := time.Now()
	built_dy := pair.GetTokensOut(tokens[i], tokens[j], dx)
	fmt.Printf("Took: %v\n\n", time.Since(now))
	main_dy := mainnetTest(client, address, tokens, i, j, trade_amount)

	main_convert, _ := uint256.FromBig(main_dy)
	diff := new(uint256.Int).Sub(built_dy, main_convert)

	var check_dy_main bool
	main_swap_total += 1
	if built_dy.Uint64() == main_dy.Uint64() {
		check_dy_main = true
		main_swap_count += 1
	} else {
		check_dy_main = false
	}

	fmt.Println("===================== INITAL INPUT VALUES ====================")
	fmt.Printf("Swaps from %v pool \n", address)
	fmt.Printf("Swaps from Index %v ----> %v \n", i, j)
	fmt.Printf("Swaps from Tokens %v ----> %v \n", tokens[i], tokens[j])
	fmt.Printf("Trade Size: %v \n", trade_amount)
	fmt.Printf("rates: %v \n", rates)
	fmt.Println("===================== SWAP RETURN VALUES ====================")
	fmt.Printf("DY built vs main [%v, %v] \n", built_dy, main_dy)
	fmt.Println("Outcome of Dy test:", check_dy_main)
	fmt.Println("Difference:", diff)
	fmt.Println("")
}

func mainnetTest(client bind.ContractBackend, address common.Address, tokens []common.Address, i, j uint64, trade_amount *big.Int) *big.Int {

	swap_address := common.HexToAddress("0x55B916Ce078eA594c10a874ba67eCc3d62e29822")
	swap, _ := NewSwapCaller(swap_address, client)

	main, e := swap.GetExchangeAmount(nil, address, tokens[i], tokens[j], trade_amount)

	if e != nil {
		fmt.Println("Mainnet transaction reverted")
		return big.NewInt(0)
	} else {
		return main
	}

}
