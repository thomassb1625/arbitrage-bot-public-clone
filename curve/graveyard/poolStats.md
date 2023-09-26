// For checking implemenation types and number of pools
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
