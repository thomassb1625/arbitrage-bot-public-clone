package dex

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

type Pair interface {
	//GetTokensIn(tokenIn common.Address, tokenOut common.Address, amountOut *uint256.Int) *uint256.Int
	GetTokensOut(tokenIn common.Address, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int
	//Later split this into two functions
	SellTokens(tokenIn common.Address, amountIn *uint256.Int, recipient, sender common.Address) ([]common.Address, [][]byte) //No longer used
	ReceiveDirectly() bool                                                                                                   //No longer used
	PrepareReceive(tokenAddress common.Address, amountIn *uint256.Int) ([]common.Address, [][]byte)                          //No longer used
	Reserves() (*uint256.Int, *uint256.Int)
	MarketAddress() common.Address
	SpenderAddress() common.Address
	Protocol() string
	Tokens() (Token, Token)
	GetOtherToken(knownToken common.Address) Token
	HeavyComputation() bool
}

var Logger log.Logger = *log.New(os.Stdout, "", 0)

func SetLogger(logger log.Logger) {
	Logger = logger
}
