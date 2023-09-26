package arbitrage

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/abis"
	"github.com/jmjac/ethGoTest/dex"
	"github.com/jmjac/ethGoTest/dex/uniswapV2"
	"github.com/jmjac/ethGoTest/dex/uniswapV3"
)

var approveLimit = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1e18))
var IERC20ABI, _ = abi.JSON(strings.NewReader(abis.IERC20ABI))

//MAX 10 bytes
func WithdrawlWETH(amount *uint256.Int) ([]byte, int64) {
	payload := "0x"
	payload += encodeTo10Bytes(amount)
	v, _ := hexutil.Decode(payload)
	return v, 0
}

//For dodoPools we need to know for every pool which token is Base and which is Quote
//Only works with WETH
func (a Arbitrage) NewPrepareTx() ([]byte, *big.Int, bool) {
	onlyV2 := true
	for _, p := range a.Path {
		switch p.(type) {
		case *uniswapV2.Pair:
		default:
			onlyV2 = false
			break
		}
	}

	var data []byte
	var msgValue uint64

	if onlyV2 {
		data, msgValue = a.handleUniswapV2()
	} else {
		data, msgValue = a.handleMixed()
	}

	return data, big.NewInt(int64(msgValue)), onlyV2
}

//This maybe slow
func isV3Pool(pool dex.Pair) bool {
	switch pool.(type) {
	case *uniswapV3.Pair:
		return true
	default:
		return false
	}
}

//TODO: Add tax calculation
//Supports only up to 8 swaps
func (a Arbitrage) handleMixed() ([]byte, uint64) {
	//Use the other sender
	if len(a.Path) > 8 {
		panic("Mixed swap supports up to 8 swaps")
	}

	payload := "0x"
	msgValue := uint64(len(a.Path) - 1)
	in := a.AmountIn

	startPool := a.Path[0]
	t0, _ := startPool.Tokens()
	tokenOut := startPool.GetOtherToken(a.StartToken).Addr
	out := startPool.GetTokensOut(a.StartToken, tokenOut, in)

	if a.StartToken == t0.Addr {
		msgValue = msgValue | 16 //Flip the bit for ZeroForOne
	}

	switch startPool.(type) {
	case *uniswapV2.Pair:
		payload += encodeTo9Bytes(in)   // Weth in amount
		payload += encodeTo14Bytes(out) // Pair 0 Out
		payload += startPool.MarketAddress().Hex()[2:]
	case *uniswapV3.Pair:
		msgValue = msgValue | 8 //Indicate we start with v3
		payload += encodeTo9Bytes(in)
		payload += startPool.MarketAddress().Hex()[2:]
	default:
		panic("Only supports V2 and V3")
	}

	//Flip the bit if we sent to the next pool
	if !isV3Pool(a.Path[1]) {
		msgValue = msgValue | 32
	}
	valueOffset := uint64(128)

	tokenIn := tokenOut
	//Add tax calculation

	tOut := startPool.GetOtherToken(a.StartToken)
	if tOut.IsTaxed {
		out.Mul(out, tOut.AfterTaxNum)
		out.Div(out, tOut.AfterTaxDen)
		if isV3Pool(a.Path[1]) {
			out.Mul(out, tOut.AfterTaxNum)
			out.Div(out, tOut.AfterTaxDen)
		}
	}

	in = out

	for i, p := range a.Path {
		if i == 0 {
			continue
		}
		//TODO: Add tax calculation
		outToken := p.GetOtherToken(tokenIn)
		outTokenAddr := outToken.Addr
		outAmount := p.GetTokensOut(tokenIn, outTokenAddr, in)

		if outToken.IsTaxed {
			outAmount.Mul(outAmount, outToken.AfterTaxNum)
			outAmount.Div(outAmount, outToken.AfterTaxDen)
			if i+1 < len(a.Path) && isV3Pool(a.Path[i+1]) {
				outAmount.Mul(outAmount, outToken.AfterTaxNum)
				outAmount.Div(outAmount, outToken.AfterTaxDen)
			}
		}

		if isV3Pool(p) {
			msgValue = msgValue | valueOffset //flip bit for v3 swap
			payload += p.MarketAddress().Hex()[2:]
			payload += tokenIn.Hex()[2:]
			payload += encodeTo14Bytes(in) // Pair x In for v3
		} else {
			payload += p.MarketAddress().Hex()[2:]
			payload += encodeTo14Bytes(outAmount) // Pair x out for v2
		}

		t0, _ := p.Tokens()

		if tokenIn == t0.Addr {
			msgValue = msgValue | (valueOffset * 2) //Flip the bit for ZeroForOne
		}

		if i+1 < len(a.Path) {
			if !isV3Pool(a.Path[i+1]) {
				msgValue = msgValue | (valueOffset * 4) //Flip the bit because we sent money to next pool
			}

		}

		valueOffset *= 8 // shift 3 bits

		in = outAmount
		tokenIn = outTokenAddr
	}

	v, _ := hexutil.Decode(payload)

	return v, msgValue
}

/*
	MsgValue selects how many swaps

	Layout:
	20b pair0 Address
	9b ETH input  - First swap input
	20b pair1 Address
	14b Second swap input
	20b pair2 Address
	...
	20b pair4 Address
	14b Last swap input
*/

//TODO: Add tax calculation
func (a Arbitrage) handleUniswapV2() ([]byte, uint64) {
	msgValue := uint64(len(a.Path) - 1) // For 2 swaps should be one, for 3 two and so one. Look at yul contract
	if msgValue >= 6 {
		panic("More than 6 swaps are not supported by the contract")
	}
	//fmt.Printf("With %v len set msgValue to %v\n", len(a.Path), msgValue)
	zeroForOne := uint64(8) //Initial bit position for univ2 swaps
	tokenIn := a.StartToken
	in := a.AmountIn

	payload := "0x"

	for i, p := range a.Path {
		t0, _ := p.Tokens()

		if tokenIn == t0.Addr {
			msgValue = msgValue | zeroForOne //Flip the bit
		}

		//fmt.Printf("Added %v as adddress of 20 bytes\n", p.MarketAddress().Hex()[2:])
		payload += p.MarketAddress().Hex()[2:] //Add pair address to payload. Should be 20 bytes

		if i == 0 {
			//fmt.Printf("Added %v as 9 bytes value\n", encodeTo9Bytes(in))
			payload += encodeTo9Bytes(in) //Add WETH value for transfer
		} else {
			//fmt.Printf("Added %v as 14 bytes value\n", encodeTo14Bytes(in))
			payload += encodeTo14Bytes(in) //Add expected output of next trade?
		}

		tokenOut := p.GetOtherToken(tokenIn)
		out := p.GetTokensOut(tokenIn, tokenOut.Addr, in)

		if tokenOut.IsTaxed {
			out.Mul(out, tokenOut.AfterTaxNum)
			out.Div(out, tokenOut.AfterTaxDen)
		}
		in = out
		tokenIn = tokenOut.Addr

		zeroForOne = zeroForOne * 2 //Move to next bit
	}

	payload += encodeTo9Bytes(in) //Add exected value out of eth

	data, err := hexutil.Decode(payload)
	if err != nil {
		panic(err)
	}
	return data, msgValue
}

var MAX_9_BYTES, _ = uint256.FromHex("0xFFFFFFFFFFFFFFFFFF")

func encodeTo9Bytes(v *uint256.Int) string {
	if v.Cmp(MAX_9_BYTES) > 0 {
		log.Fatalf("Can't encode %v in 9 bytes", v)
	}
	return fmt.Sprintf("%018x", v)
}

var MAX_10_BYTES, _ = uint256.FromHex("0xFFFFFFFFFFFFFFFFFFFFFF")

func encodeTo10Bytes(v *uint256.Int) string {
	if v.Cmp(MAX_10_BYTES) > 0 {
		log.Fatalf("Can't encode %v in 10 bytes", v)
	}
	return fmt.Sprintf("%020x", v)
}

var MAX_14_BYTES, _ = uint256.FromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFF")

func encodeTo14Bytes(v *uint256.Int) string {
	if v.Cmp(MAX_14_BYTES) > 0 {
		log.Fatalf("Can't encode %v in 14 bytes", v)
	}
	return fmt.Sprintf("%028x", v)
}

////TODO: Keep track of already approved tokens
//func (a Arbitrage) PrepareTx(arbToken, bundleAddr common.Address) (*big.Int, []common.Address, [][]byte) {
//	//PrintPath(arbToken, a.Path)
//	targets := make([]common.Address, 0)
//	payloads := make([][]byte, 0)
//	amountIn := a.AmountIn.ToBig()
//
//	var addr common.Address
//	var pay []byte
//	//wethWasApproved := false
//	//Depeding on the protocol either approve it or transfer directly
//
//	//TODO: Add here a check if initial token is taxed
//	if a.Path[0].ReceiveDirectly() {
//		addr, pay = TransferToken(arbToken, a.Path[0].MarketAddress(), amountIn)
//	} else {
//		addr, pay = ApproveToken(arbToken, a.Path[0].SpenderAddress(), approveLimit)
//		//wethWasApproved = true
//	}
//	targets = append(targets, addr)
//	payloads = append(payloads, pay)
//
//	in := a.AmountIn
//	token := a.StartToken
//	for i, pool := range a.Path[0 : len(a.Path)-1] {
//		tokenOut := pool.GetOtherToken(token)
//		out := pool.GetTokensOut(token, tokenOut.Addr, in)
//
//		var t []common.Address
//		var p [][]byte
//
//		if tokenOut.IsTaxed {
//			out.Mul(out, tokenOut.AfterTaxNum)
//			out.Div(out, tokenOut.AfterTaxDen)
//			if !a.Path[i+1].ReceiveDirectly() {
//				out.Mul(out, tokenOut.AfterTaxNum)
//				out.Div(out, tokenOut.AfterTaxDen)
//			}
//		}
//
//		if a.Path[i+1].ReceiveDirectly() {
//			//For uniswap sender dosen't matter
//			t, p = pool.SellTokens(token, in, a.Path[i+1].MarketAddress(), bundleAddr)
//		} else {
//			//For balancer we need to sent the money back to ourselves
//			t, p = pool.SellTokens(token, in, bundleAddr, bundleAddr)
//		}
//		for _, targ := range t {
//			targets = append(targets, targ)
//		}
//
//		for _, data := range p {
//			payloads = append(payloads, data)
//		}
//
//		in = out
//		token = tokenOut.Addr
//	}
//
//	pool := a.Path[len(a.Path)-1]
//	t, p := pool.SellTokens(token, in, bundleAddr, bundleAddr)
//	for _, targ := range t {
//		targets = append(targets, targ)
//	}
//
//	for _, data := range p {
//		payloads = append(payloads, data)
//	}
//
//	//Set approval for WETH to 0
//	//To save gas we don't remove approvals anymore
//	//	if wethWasApproved {
//	//		addr, pay = approveToken(ADDR.WETH, a.Path[0].SpenderAddress(), big.NewInt(0))
//	//		targets = append(targets, addr)
//	//		payloads = append(payloads, pay)
//	//	}
//	return amountIn, targets, payloads
//}

//Set the amount to 2^200 or somethinf for anything thats not WETH
func ApproveToken(token common.Address, spender common.Address, amount *big.Int) (common.Address, []byte) {
	approveData, _ := IERC20ABI.Pack("approve", spender, amount)
	return token, approveData
}

func TransferToken(token common.Address, destination common.Address, amount *big.Int) (common.Address, []byte) {
	transferData, _ := IERC20ABI.Pack("transfer", destination, amount)
	return token, transferData
}
