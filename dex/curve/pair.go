package curve

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type Pair struct {
	address  common.Address
	tokens   []common.Address
	reserves []*uint256.Int
	protocol string

	A                 *uint256.Int
	rates             []*uint256.Int
	LENDING_PRECISION *uint256.Int
	PRECISION         *uint256.Int
	FEE_DENOMINATOR   *uint256.Int
	fee               *uint256.Int
	admin_fee         *uint256.Int
	N_COINS           *uint256.Int
	implemention      common.Address
}

func newPair(
	pairAddr common.Address,
	tokens []common.Address,

	A *uint256.Int,
	rates []*uint256.Int,
	fee *uint256.Int,
	admin_fee *uint256.Int,
	LENDING_PRECISION *uint256.Int,
	PRECISION *uint256.Int,
	FEE_DENOMINATOR *uint256.Int,
	N_COINS *uint256.Int,
	implemenation common.Address,
	protocol string) Pair {

	pair := Pair{
		address:           pairAddr,
		tokens:            tokens,
		A:                 A,
		rates:             rates,
		fee:               fee,
		admin_fee:         admin_fee,
		LENDING_PRECISION: LENDING_PRECISION,
		PRECISION:         PRECISION,
		FEE_DENOMINATOR:   FEE_DENOMINATOR,
		N_COINS:           N_COINS,
		protocol:          protocol,
		implemention:      implemenation,
	}

	pair.reserves = make([]*uint256.Int, 0)
	return pair
}

func (pair Pair) GetTokensOut(tokenIn common.Address, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int {
	i, j := pair.ConvertAddresses(tokenIn, tokenOut)
	if pair.implemention == common.HexToAddress("0x0000000000000000000000000000000000000000") {
		return pair.ThreePool_Dy(i, j, amountIn)
	} else if pair.implemention == common.HexToAddress("0x6523Ac15EC152Cb70a334230F6c5d62C5Bd963f1") {
		return pair.BasePool_Dy(i, j, amountIn)
	} else {
		return uint256.NewInt(0)
	}

	// TO DO: Finish Switch Statement for multiple implementations
	/* switch pair.implemention {
	case common.HexToAddress("0x0000000000000000000000000000000000000000"):
		return pair.ThreePool_Dy(i, j, amountIn)
	case common.HexToAddress("0x6523Ac15EC152Cb70a334230F6c5d62C5Bd963f1"):
		return pair.BasePool_Dy(i, j, amountIn)
	} */
}

func (pair Pair) ConvertAddresses(tokenIn common.Address, tokenOut common.Address) (uint64, uint64) {
	var i, j uint64
	for k, t := range pair.Tokens() {
		if t == tokenIn {
			i = uint64(k)
		}
		if t == tokenOut {
			j = uint64(k)
		}
	}
	return i, j
}

func (pair Pair) tokensToReserves(tokenIn common.Address, tokenOut common.Address) (*uint256.Int, *uint256.Int) {
	panic("Not implemented")
}

func (pair Pair) getAmountOut(reserveIn *uint256.Int, reserveOut *uint256.Int, amountIn *uint256.Int) *uint256.Int {
	panic("Not implemented")
}

func (pair Pair) Fee() *uint256.Int {
	return pair.fee
}

func SellTokensToNextMarket(tokenIn common.Address, amountIn *uint256.Int, ethMarketAddress common.Address) []byte {
	panic("Not implemented")
}

func (pair Pair) SellTokens(tokenIn common.Address, amountIn *uint256.Int, recipient, sender common.Address) ([]common.Address, [][]byte) {
	panic("Not implemented")
}

func (pair Pair) ReceiveDirectly() bool {
	return true
}

func (pair Pair) PrepareReceive(tokenAddress common.Address, amountIn *uint256.Int) ([]common.Address, [][]byte) {
	panic("Not implemented")
}

func (pair Pair) GetOtherTokenAddress(knowAddress common.Address) common.Address {
	panic("Not implemented")
}

func (pair Pair) MarketAddress() common.Address {
	return pair.address
}

func (pair Pair) SpenderAddress() common.Address {
	panic("Not implemented")
}

func (pair Pair) Tokens() []common.Address {
	return pair.tokens
}

func (pair Pair) Protocol() string {
	return pair.protocol
}

func (pair Pair) Reserves() []*uint256.Int {
	return pair.reserves
}

func (pair Pair) Rates() []*uint256.Int {
	return pair.rates
}

func (pair *Pair) FakeReserves(fake *uint256.Int) {
	panic("Not implemented")
}
