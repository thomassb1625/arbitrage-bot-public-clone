package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/arbitrage"
	"github.com/jmjac/ethGoTest/dex"
	"github.com/jmjac/ethGoTest/dex/uniswapV2"
	"github.com/jmjac/ethGoTest/flashbots"
	"github.com/jmjac/ethGoTest/helperClient"
	"github.com/jmjac/ethGoTest/secrets"
)

var bytecode = "0x60806040526004361061001e5760003560e01c80635cdffdca14610023575b600080fd5b610036610031366004610277565b61004c565b60405161004391906103d4565b60405180910390f35b6000805b8451811015610112576000606086838151811061006957fe5b60200260200101516001600160a01b031686848151811061008657fe5b602002602001015160405161009b9190610362565b6000604051808303816000865af19150503d80600081146100d8576040519150601f19603f3d011682016040523d82523d6000602084013e6100dd565b606091505b5091509150816101085760405162461bcd60e51b81526004016100ff906103af565b60405180910390fd5b5050600101610050565b506040516370a0823160e01b81526001600160a01b038316906370a082319061013f90309060040161039b565b60206040518083038186803b15801561015757600080fd5b505afa15801561016b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061018f919061034a565b949350505050565b80356001600160a01b03811681146101ae57600080fd5b92915050565b6000601f83818401126101c5578182fd5b82356101d86101d382610404565b6103dd565b818152925060208084019085810160005b8481101561026b578135880189603f82011261020457600080fd5b8381013567ffffffffffffffff81111561021d57600080fd5b61022e818901601f191686016103dd565b81815260408c8184860101111561024457600080fd5b828185018884013750600091810186019190915285525092820192908201906001016101e9565b50505050505092915050565b60008060006060848603121561028b578283fd5b833567ffffffffffffffff808211156102a2578485fd5b818601915086601f8301126102b5578485fd5b81356102c36101d382610404565b80828252602080830192508086018b8283870289010111156102e357898afd5b8996505b8487101561030d576102f98c82610197565b8452600196909601959281019281016102e7565b509097508801359350505080821115610324578384fd5b50610331868287016101b4565b9250506103418560408601610197565b90509250925092565b60006020828403121561035b578081fd5b5051919050565b60008251815b818110156103825760208186018101518583015201610368565b818111156103905782828501525b509190910192915050565b6001600160a01b0391909116815260200190565b6020808252600b908201526a4661696c6564206865726560a81b604082015260600190565b90815260200190565b60405181810167ffffffffffffffff811182821017156103fc57600080fd5b604052919050565b600067ffffffffffffffff82111561041a578081fd5b506020908102019056fea26469706673582212204135bebfda2af6e9a5ec538d8d04d2af3713423c2d2e036142049b7b0599695264736f6c634300060c0033"

func main() {
	provider := "https://xyznode.space/api/mainnet/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"
	mainnet := true

	helperClient, err := helperClient.New(provider)
	if err != nil {
		fmt.Println("Coudn't create helper client,", err)
		return
	}

	//Setup adreses
	if mainnet {
		ADDR.MainnetAddresses()
	} else {
		ADDR.GoerliAddresses()
	}

	pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsPancakeSwap := setupUniV2Pairs(helperClient.Client, make(map[common.Address]bool), make(map[common.Address]dex.Token))
	allUniV2 := mergeUniV2Markets(pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsPancakeSwap)

	blockNum, _ := helperClient.BlockNumber(context.Background())
	newBlock, _ := helperClient.BlockNumber(context.Background())
	fmt.Println("Syncing with blocktime")
	for newBlock-blockNum < 1 {
		newBlock, _ = helperClient.BlockNumber(context.Background())
		fmt.Println(blockNum)
		time.Sleep(time.Millisecond * 330)
	}
	fmt.Println("Found new block")

	uniswapV2.UpdateReservesWithoutLookup(pairsUniV2, helperClient)
	uniswapV2.UpdateReservesWithoutLookup(pairsPancakeSwap, helperClient)
	uniswapV2.UpdateReservesWithoutLookup(pairsSushi, helperClient)
	uniswapV2.UpdateReservesWithoutLookup(pairsShibaSwap, helperClient)
	uniswapV2.UpdateReservesWithoutLookup(pairsOrion, helperClient)
	uniswapV2.UpdateReservesWithoutLookup(pairsSake, helperClient)
	uniswapV2.UpdateReservesWithoutLookup(pairsDXSwap, helperClient)
	uniswapV2.UpdateReservesWithoutLookup(pairsCryptDefi, helperClient)
	uniswapV2.UpdateReservesWithoutLookup(pairsFraxSwap, helperClient)

	time.Sleep(time.Second * 12)

	ethPools := make([]dex.Pair, 0)
	for _, p := range allUniV2 {
		t0, t1 := p.Tokens()
		if t0.Addr == ADDR.WETH || t1.Addr == ADDR.WETH {
			ethPools = append(ethPools, p)
		}
	}

	failed := make([]common.Address, 0)
	tokensWithTax := make([]dex.Token, 0)

	fmt.Println(len(ethPools))
	signer := flashbots.NewSigner(big.NewInt(1), secrets.PRIVATE_KEY_GOERLI, ADDR.BUNDLE_EXECUTOR, ADDR.WETH, helperClient)

	checkForTax, _ := abi.JSON(strings.NewReader(CheckTransferTaxABI))

	v := uint256.NewInt(3e17)
	for _, p := range ethPools {
		r0, r1 := p.Reserves()
		if r0.IsZero() || r1.IsZero() {
			continue
		}
		targets := make([]common.Address, 0)
		payloads := make([][]byte, 0)
		tar, pay := arbitrage.TransferToken(ADDR.WETH, p.MarketAddress(), v.ToBig())
		targets = append(targets, tar)
		payloads = append(payloads, pay)

		targs, pays := p.SellTokens(ADDR.WETH, v, ADDR.BUNDLE_EXECUTOR, signer.Address)
		for i := range targs {
			targets = append(targets, targs[i])
			payloads = append(payloads, pays[i])
		}

		data, _ := checkForTax.Pack("swapWithoutFlashLoan", targets, payloads, p.GetOtherToken(ADDR.WETH).Addr)
		resp, err := helperClient.EthCallInjectBytecode(signer.Address, ADDR.BUNDLE_EXECUTOR, data, bytecode, ADDR.BUNDLE_EXECUTOR, fmt.Sprintf("0x%x", blockNum))

		if err != nil {
			failed = append(failed, p.GetOtherToken(ADDR.WETH).Addr)
			//fmt.Printf("Failed at %v\n=======================\n", p.MarketAddress())
			continue
		}

		var out []interface{}
		results := &out
		if results == nil {
			results = new([]interface{})
		}

		if len(*results) == 0 {
			res, err0 := checkForTax.Unpack("swapWithoutFlashLoan", resp)
			err = err0
			*results = res
		} else {
			res := *results
			err = checkForTax.UnpackIntoInterface(res[0], "swapWithoutFlashLoan", resp)
		}

		outReal := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

		simulated := p.GetTokensOut(ADDR.WETH, p.GetOtherToken(ADDR.WETH).Addr, v)

		if outReal.Cmp(simulated.ToBig()) != 0 {
			fmt.Printf("For token %v\n", p.GetOtherToken(ADDR.WETH).Addr)
			fmt.Printf("Got %v\n", outReal)
			fmt.Printf("Got simulated %v\n", simulated)
			fReal := new(big.Float).SetInt(outReal)
			fSim := new(big.Float).SetInt(simulated.ToBig())
			afterTax, _ := fReal.Quo(fReal, fSim).Float64()
			numerator := uint256.NewInt(uint64(math.Floor(afterTax * 1000)))
			denominator := uint256.NewInt(1000)
			simulated.Mul(simulated, numerator)
			simulated.Div(simulated, denominator)
			fmt.Printf("With tax it should be %v\n", simulated)
			fmt.Printf("Estimated amount after tax of %v%%\n==========================\n", afterTax)

			//Tokens with more money should sent
			if afterTax > 0.99999999 {
				continue
			}

			taxedToken := dex.NewTaxed(p.GetOtherToken(ADDR.WETH).Addr, numerator, denominator, afterTax > 0.03)
			tokensWithTax = append(tokensWithTax, taxedToken)
		}
	}

	//Filter out repetition
	tmap := dex.MakeTokenMap(tokensWithTax)
	tokensWithTax = make([]dex.Token, 0)
	for _, t := range tmap {
		tokensWithTax = append(tokensWithTax, t)
	}
	dex.SaveTokens(tokensWithTax, "tokensTax.json")
	fmt.Printf("Got %v tokens with Tax\n", len(tokensWithTax))
	fmt.Printf("Failed with %v tokens\n", len(failed))

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

func setupUniV2Pairs(client *ethclient.Client, BLACKLIST map[common.Address]bool, tokenMap map[common.Address]dex.Token) ([]*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair, []*uniswapV2.Pair) {
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

	return pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsPancakeSwap
}

func castUniV2Pairs(pairsUniV2, pairsSushi, pairsShibaSwap, pairsOrion, pairsSake, pairsDXSwap, pairsCryptDefi, pairsFraxSwap, pairsPancakeSwap []*uniswapV2.Pair) ([]dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair, []dex.Pair) {
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

	return uniV2, sushi, shiba, orion, sake, dxSwap, cryptDefi, frax, pancake
}
