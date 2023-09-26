package uniswapV3

import (
	"errors"
	"math/big"

	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/holiman/uint256"
)

// Errors for swap math functions
var (
	ErrAddOverflow           = errors.New("add overflow")
	ErrMulOverflow           = errors.New("mul overflow")
	ErrInvariant             = errors.New("invariant")
	ErrSqrtLessThanZero      = errors.New("sqrt less than zero")
	ErrLiquidityLessThanZero = errors.New("liquidity less than zero")
)

func computeSwapStep(sqrtRatioCurrentX96, sqrtRatioTargetX96, liquidity, amountRemaining, feePips uint256.Int, exactInput bool) (sqrtRatioNextX96, amountIn, feeAmount, amountOut uint256.Int, err error) {
	zeroForOne := sqrtRatioCurrentX96.Cmp(&sqrtRatioTargetX96) >= 0

	if exactInput {
		amountRemainingLessFee := new(uint256.Int).Div(new(uint256.Int).Mul(&amountRemaining, new(uint256.Int).Sub(FeeMultiplier, &feePips)), FeeMultiplier)
		//fmt.Println("AMOUNT REMAININNG LESS: ", amountRemainingLessFee.ToBig().String())
		if zeroForOne {
			amountIn = getAmount0Delta(&sqrtRatioTargetX96, &sqrtRatioCurrentX96, &liquidity, true)
		} else {
			amountIn = getAmount1Delta(&sqrtRatioCurrentX96, &sqrtRatioTargetX96, &liquidity, true)
		}
		//fmt.Println("AMOUNT IN: ", amountIn.ToBig().String())

		sqrtRatioNextX96 = sqrtRatioTargetX96
		if amountRemainingLessFee.Cmp(&amountIn) < 0 {
			sqrtRatioNextX96, err = getNextSqrtPriceFromInput(&sqrtRatioCurrentX96, &liquidity, amountRemainingLessFee, zeroForOne)
			if err != nil {
				return
			}
		}
	} else {
		if zeroForOne {
			amountOut = getAmount1Delta(&sqrtRatioTargetX96, &sqrtRatioCurrentX96, &liquidity, false)
		} else {
			amountOut = getAmount0Delta(&sqrtRatioCurrentX96, &sqrtRatioTargetX96, &liquidity, false)
		}

		if amountRemaining.Cmp(&amountOut) >= 0 {
			sqrtRatioNextX96 = sqrtRatioTargetX96
		} else {
			sqrtRatioNextX96, err = getNextSqrtPriceFromOutput(&sqrtRatioCurrentX96, &liquidity, &amountRemaining, zeroForOne)
			if err != nil {
				return
			}
		}
	}

	max := sqrtRatioTargetX96.Cmp(&sqrtRatioNextX96) == 0

	if zeroForOne {
		if !(max && exactInput) {
			amountIn = getAmount0Delta(&sqrtRatioNextX96, &sqrtRatioCurrentX96, &liquidity, true)
		}
		if !(max && !exactInput) {
			amountOut = getAmount1Delta(&sqrtRatioNextX96, &sqrtRatioCurrentX96, &liquidity, false)
		}
	} else {
		if !(max && exactInput) {
			amountIn = getAmount1Delta(&sqrtRatioCurrentX96, &sqrtRatioNextX96, &liquidity, true)
		}
		if !(max && !exactInput) {
			amountOut = getAmount0Delta(&sqrtRatioCurrentX96, &sqrtRatioNextX96, &liquidity, false)
		}
	}
	//WARNING: ONLY WORKS WITH EXACT INPUT

	//Different than univ3-go-sdk
	//if exactInput && amountOut.Cmp(&amountRemaining) > 0 {
	//	amountOut = amountRemaining
	//}

	//if !exactInput && amountOut.Cmp(new(big.Int).Mul(amountRemaining, constants.NegativeOne)) > 0 {
	//	amountOut = new(big.Int).Mul(amountRemaining, constants.NegativeOne)
	//}

	if exactInput && sqrtRatioNextX96.Cmp(&sqrtRatioTargetX96) != 0 {
		// we didn't reach the target, so take the remainder of the maximum input as fee
		feeAmount = *new(uint256.Int).Sub(&amountRemaining, &amountIn)
	} else {
		feeAmount = mulDivRoundingUp(&amountIn, &feePips, new(uint256.Int).Sub(FeeMultiplier, &feePips))
	}

	return
}

func getNextSqrtPriceFromOutput(sqrtPX96, liquidity, amountOut *uint256.Int, zeroForOne bool) (uint256.Int, error) {
	if sqrtPX96.Cmp(ZeroUint256) <= 0 {
		return *uint256.NewInt(0), ErrSqrtLessThanZero
	}
	if liquidity.Cmp(ZeroUint256) <= 0 {
		return *uint256.NewInt(0), ErrLiquidityLessThanZero
	}
	if zeroForOne {
		return getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountOut, false)
	}
	return getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountOut, false)
}

func getNextSqrtPriceFromInput(sqrtPX96, liquidity, amountIn *uint256.Int, zeroForOne bool) (uint256.Int, error) {
	if sqrtPX96.Cmp(ZeroUint256) <= 0 {
		return *uint256.NewInt(0), ErrSqrtLessThanZero
	}
	if liquidity.Cmp(ZeroUint256) <= 0 {
		return *uint256.NewInt(0), ErrLiquidityLessThanZero
	}
	if zeroForOne {
		return getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amountIn, true)
	}
	return getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amountIn, true)
}

func getNextSqrtPriceFromAmount0RoundingUp(sqrtPX96, liquidity, amount *uint256.Int, add bool) (uint256.Int, error) {
	if amount.Cmp(ZeroUint256) == 0 {
		return *sqrtPX96, nil
	}

	numerator1 := new(uint256.Int).Lsh(liquidity, 96)
	if add {
		product, isOverflow := new(uint256.Int).MulOverflow(amount, sqrtPX96)
		if isOverflow {
			return *uint256.NewInt(0), ErrMulOverflow
		}
		if new(uint256.Int).Div(product, amount).Cmp(sqrtPX96) == 0 {
			denominator, isOverflow := new(uint256.Int).AddOverflow(numerator1, product)
			if isOverflow {
				return *uint256.NewInt(0), ErrAddOverflow
			}
			if denominator.Cmp(numerator1) >= 0 {
				return mulDivRoundingUp(numerator1, sqrtPX96, denominator), nil
			}
		}
		return mulDivRoundingUp(numerator1, OneUint256, new(uint256.Int).Add(new(uint256.Int).Div(numerator1, sqrtPX96), amount)), nil
	}

	product, isOverflow := new(uint256.Int).MulOverflow(amount, sqrtPX96)
	if isOverflow {
		return *uint256.NewInt(0), ErrMulOverflow
	}
	if new(uint256.Int).Div(product, amount).Cmp(sqrtPX96) != 0 {
		return *uint256.NewInt(0), ErrInvariant
	}
	if numerator1.Cmp(product) <= 0 {
		return *uint256.NewInt(0), ErrInvariant
	}
	denominator := new(uint256.Int).Sub(numerator1, product)
	return mulDivRoundingUp(numerator1, sqrtPX96, denominator), nil
}

func getNextSqrtPriceFromAmount1RoundingDown(sqrtPX96, liquidity, amount *uint256.Int, add bool) (uint256.Int, error) {
	if add {
		var quotient *uint256.Int
		if amount.Cmp(MaxUint160) <= 0 {
			quotient = new(uint256.Int).Div(new(uint256.Int).Lsh(amount, 96), liquidity)
		} else {
			quotient = new(uint256.Int).Div(new(uint256.Int).Mul(amount, Q96Uint256), liquidity)
		}
		return *new(uint256.Int).Add(sqrtPX96, quotient), nil
	}

	quotient := mulDivRoundingUp(amount, Q96Uint256, liquidity)
	if sqrtPX96.Cmp(&quotient) <= 0 {
		return *uint256.NewInt(0), ErrInvariant
	}
	return *new(uint256.Int).Sub(sqrtPX96, &quotient), nil
}

func getAmount0Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity *uint256.Int, roundUp bool) uint256.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) >= 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	numerator1 := new(uint256.Int).Lsh(liquidity, 96)
	numerator2 := new(uint256.Int).Sub(sqrtRatioBX96, sqrtRatioAX96)

	if roundUp {
		_mulDiv := mulDivRoundingUp(numerator1, numerator2, sqrtRatioBX96)
		return mulDivRoundingUp(&_mulDiv, OneUint256, sqrtRatioAX96)
	}
	return *new(uint256.Int).Div(new(uint256.Int).Div(new(uint256.Int).Mul(numerator1, numerator2), sqrtRatioBX96), sqrtRatioAX96)
}

func getAmount1Delta(sqrtRatioAX96, sqrtRatioBX96, liquidity *uint256.Int, roundUp bool) uint256.Int {
	if sqrtRatioAX96.Cmp(sqrtRatioBX96) >= 0 {
		sqrtRatioAX96, sqrtRatioBX96 = sqrtRatioBX96, sqrtRatioAX96
	}

	if roundUp {
		return mulDivRoundingUp(liquidity, new(uint256.Int).Sub(sqrtRatioBX96, sqrtRatioAX96), Q96Uint256)
	}
	return *new(uint256.Int).Div(new(uint256.Int).Mul(liquidity, new(uint256.Int).Sub(sqrtRatioBX96, sqrtRatioAX96)), Q96Uint256)
}

func mulDivRoundingUp(a, b, denominator *uint256.Int) uint256.Int {
	product := new(big.Int).Mul(a.ToBig(), b.ToBig())
	result := new(big.Int).Div(product, denominator.ToBig())
	if new(big.Int).Rem(product, denominator.ToBig()).Cmp(big.NewInt(0)) != 0 {
		result.Add(result, constants.One)
	}
	r, _ := uint256.FromBig(result)
	return *r
}
