package curve

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

func GetCurveMarkets(client bind.ContractBackend, CURVE_ADDRESS_PROVIDER, CURVE_REGISTRY, CURVE_FACTORY common.Address, protocol string) []Pair {
	fmt.Printf("=============== Setting up %v ===================\n", protocol)
	now := time.Now()

	marketPairs := make([]Pair, 0)

	registryMarkets := RegistryMarkets(client, CURVE_REGISTRY, protocol)
	factoryMarkets := FactoryMarkets(client, CURVE_FACTORY, protocol)

	for i := 0; i < len(registryMarkets); i++ {
		marketPairs = append(marketPairs, registryMarkets[i])
	}
	for i := 0; i < len(factoryMarkets); i++ {
		marketPairs = append(marketPairs, factoryMarkets[i])
	}

	fmt.Printf("Loaded %v pools from %v\n", len(marketPairs), protocol)
	fmt.Printf("Took: %v\n\n", time.Since(now))

	UpdateReserves(client, CURVE_REGISTRY, CURVE_FACTORY, marketPairs, protocol)
	return marketPairs
}

func UpdateReserves(client bind.ContractBackend, CURVE_REGISTRY, CURVE_FACTORY common.Address, marketPairs []Pair, protocol string) []Pair {
	now := time.Now()
	for i := 0; i < len(marketPairs); i++ {
		if marketPairs[i].implemention != common.HexToAddress("0x0000000000000000000000000000000000000000") {
			factoryReserves := FactoryReserves(client, CURVE_FACTORY, marketPairs[i], protocol)
			marketPairs[i].reserves = factoryReserves
		} else {
			registryReserves := RegistryReserves(client, CURVE_REGISTRY, marketPairs[i], protocol)
			marketPairs[i].reserves = registryReserves
		}
	}

	fmt.Printf("=============== Updating Reserves for %v ===================\n", protocol)
	fmt.Printf("Updated %v pools from Curve \n", len(marketPairs))
	fmt.Printf("Took: %v\n\n", time.Since(now))
	return marketPairs
}

func FactoryReserves(client bind.ContractBackend, CURVE_FACTORY common.Address, pool Pair, protocol string) []*uint256.Int {
	factory, _ := NewFactoryCaller(CURVE_FACTORY, client)
	balances, _ := factory.GetBalances(nil, pool.address)
	final_balances := make([]*uint256.Int, 0)
	var i uint64
	for i = 0; i < pool.N_COINS.Uint64(); i++ {
		vUint, err := uint256.FromBig(balances[i])
		if err {
			panic("overflow")
		}
		final_balances = append(final_balances, vUint)
	}
	return final_balances
}

func RegistryReserves(client bind.ContractBackend, CURVE_REGISTRY common.Address, pool Pair, protocol string) []*uint256.Int {
	registry, _ := NewRegistryCaller(CURVE_REGISTRY, client)
	balances, _ := registry.GetBalances(nil, pool.address)
	final_balances := make([]*uint256.Int, 0)
	var i uint64
	for i = 0; i < pool.N_COINS.Uint64(); i++ {
		vUint, err := uint256.FromBig(balances[i])
		if err {
			panic("overflow")
		}
		final_balances = append(final_balances, vUint)
	}
	return final_balances
}

// TO DO: Change for necessary data
func FactoryMarkets(client bind.ContractBackend, CURVE_FACTORY common.Address, protocol string) []Pair {
	factoryMarkets := make([]Pair, 0)
	factory, _ := NewFactoryCaller(CURVE_FACTORY, client)

	pools_count, _ := factory.PoolCount(nil)

	var num = pools_count.Int64()
	var j int64 = 0
	for j < num {
		pool, _ := factory.PoolList(nil, big.NewInt(j))
		tokens, _ := factory.GetCoins(nil, pool)

		tokensNum := 0
		for _, i := range tokens {
			if i == common.HexToAddress("") {
				break
			}
			tokensNum++
		}

		factor := uint256.NewInt(0)
		raw_A, _ := factory.GetA(nil, pool)
		A, _ := uint256.FromBig(raw_A)
		rates := []*uint256.Int{factor, factor, factor}
		fee := uint256.NewInt(1000000)
		admin_fee := uint256.NewInt(5000000000)
		LENDING_PRECISION := factor
		PRECISION := factor
		FEE_DENOMINATOR := factor
		N_COINS := uint256.NewInt(3)
		implementation, _ := factory.GetImplementationAddress(nil, pool)

		final_tokens := make([]common.Address, 0)
		for i := 0; i < tokensNum; i++ {

			// Only grabs initialized addresses
			if tokens[i] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
				final_tokens = append(final_tokens, tokens[i])
			}
		}

		pair := newPair(pool, final_tokens, A, rates, fee, admin_fee, LENDING_PRECISION, PRECISION, FEE_DENOMINATOR, N_COINS, implementation, protocol)
		factoryMarkets = append(factoryMarkets, pair)
		j += 1
	}
	return factoryMarkets
}

func RegistryMarkets(client bind.ContractBackend, CURVE_REGISTRY common.Address, protocol string) []Pair {
	registryMarkets := make([]Pair, 0)
	registry, _ := NewRegistryCaller(CURVE_REGISTRY, client)

	/* pools_count, _ := registry.PoolCount(nil)
	var num = pools_count.Int64() */

	var num int64 = 1
	var j int64 = 0
	for j < num {
		pool, _ := registry.PoolList(nil, big.NewInt(j))
		tokens, _ := registry.GetCoins(nil, pool)

		tokensNum := 0
		for _, i := range tokens {
			if i == common.HexToAddress("") {
				break
			}
			tokensNum++
		}

		factor3 := uint256.NewInt(10000000000)
		factor2, _ := uint256.FromHex("0xC9F2C9CD04674EDEA40000000")
		factor := uint256.NewInt(1000000000000000000)
		raw_A, _ := registry.GetA(nil, pool)
		A, _ := uint256.FromBig(raw_A)
		rates := []*uint256.Int{factor, factor2, factor2}
		fee := uint256.NewInt(1000000)
		admin_fee := uint256.NewInt(5000000000)
		LENDING_PRECISION := factor
		PRECISION := factor
		FEE_DENOMINATOR := factor3
		N_COINS := uint256.NewInt(3)
		implementation := common.HexToAddress("0x0000000000000000000000000000000000000000")

		final_tokens := make([]common.Address, 0)
		for i := 0; i < tokensNum; i++ {

			if tokens[i] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
				final_tokens = append(final_tokens, tokens[i])
			}
		}

		pair := newPair(pool, final_tokens, A, rates, fee, admin_fee, LENDING_PRECISION, PRECISION, FEE_DENOMINATOR, N_COINS, implementation, protocol)
		registryMarkets = append(registryMarkets, pair)
		j += 1
	}

	return registryMarkets
}
