package arbitrage

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/jmjac/ethGoTest/dex"
)

func TokenToTokenMarkets(token0, token1 common.Address, markets ...[]dex.Pair) []dex.Pair {
	filtered := make([]dex.Pair, 0)
	for _, m := range markets {
		for _, p := range m {
			t0, t1 := p.Tokens()
			if (t0.Addr == token0 && t1.Addr == token1) || (t1.Addr == token0 && t0.Addr == token1) {
				filtered = append(filtered, p)
			}
		}
	}
	return filtered
}

func FilterPairs() {}
