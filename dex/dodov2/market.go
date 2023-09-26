package dodo

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/dex"
)

func (p Pair) GetTokensOut(tokenIn common.Address, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int {
	if p.state.B.IsZero() || p.state.Q.IsZero() {
		return uint256.NewInt(0)
	}

	if tokenIn == p.baseToken.Addr {
		return querySellBase(p, amountIn)
	} else {
		return querySellQuote(p, amountIn)
	}
}

func (pair Pair) PrepareReceive(tokenAddress common.Address, amountIn *uint256.Int) ([]common.Address, [][]byte) {
	panic("Not implemented")
}

func (pair Pair) ReceiveDirectly() bool {
	return true
}

func (pair Pair) GetOtherToken(knowAddres common.Address) dex.Token {
	if pair.baseToken.Addr == knowAddres {
		return pair.quoteToken
	} else {
		return pair.baseToken
	}
}

func (pair Pair) SellTokens(tokenIn common.Address, amountIn *uint256.Int, recipient, sender common.Address) ([]common.Address, [][]byte) {
	panic("Not Implemented")
}

func (pair Pair) MarketAddress() common.Address {
	return pair.addr
}

func (pair Pair) SpenderAddress() common.Address {
	panic("Dodo doesn't use spender address")
	return common.HexToAddress("")
}

func (pair Pair) Tokens() (dex.Token, dex.Token) {
	return pair.baseToken, pair.quoteToken
}

func (pair Pair) Protocol() string {
	return pair.protocol
}

func (pair Pair) Reserves() (*uint256.Int, *uint256.Int) {
	r0 := pair.state.B
	r1 := pair.state.Q
	return new(uint256.Int).Set(r0), new(uint256.Int).Set(r1)

}

func (p Pair) HeavyComputation() bool {
	return false
}
