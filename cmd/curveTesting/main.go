package main

import (
	"fmt"
	"log"

	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/arbitrage"
	"github.com/jmjac/ethGoTest/dex/curve"
	"github.com/jmjac/ethGoTest/helperClient"
)

func main() {
	url := "https://xyznode.space/api/mainnet/b3a768b5d2709d46e3099760e19c9c43b1148b8fffeb06c354dcf2b6bcaff5a7"
	ethClient, err := helperClient.New(url)
	if err != nil {
		fmt.Println("Coudn't create client,", err)
		return
	}

	ADDR.MainnetAddresses(*log.Default())

	BLACKLIST := arbitrage.InitBlackList("blacklist.json")
	fmt.Printf("Blacklist has %v tokens\n", len(BLACKLIST))

	curve.TestStats3Pool(ethClient, ADDR.CURVE_REGISTRY)
	// curve.PoolStats(ethClient)
	curve.GetCurveMarkets(ethClient, ADDR.CURVE_ADDRESS_PROVIDER, ADDR.CURVE_REGISTRY, ADDR.CURVE_FACTORY, "curve-finance")
}
