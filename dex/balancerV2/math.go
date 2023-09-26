package balancerV2

import (
	"log"
	"math/big"

	"github.com/holiman/uint256"
)

var ZERO = uint256.NewInt(0)
var ONE = uint256.NewInt(1)
var ONE_E18 = uint256.NewInt(1e18)
var TWO_E18 = uint256.NewInt(2e18)
var FOUR_E18 = uint256.NewInt(4e18)
var MAX_POW_RELATIVE_ERROR = uint256.NewInt(10000)
var MIN_POW_BASE_FREE_EXPONENT = uint256.NewInt(7e17)

var BIG_0 = big.NewInt(0)
var BIG_1 = big.NewInt(1)
var BIG_2 = big.NewInt(2)
var BIG_3 = big.NewInt(3)
var BIG_4 = big.NewInt(4)
var BIG_5 = big.NewInt(5)
var BIG_6 = big.NewInt(6)
var BIG_7 = big.NewInt(7)
var BIG_8 = big.NewInt(8)
var BIG_9 = big.NewInt(9)
var BIG_10 = big.NewInt(10)
var BIG_11 = big.NewInt(11)
var BIG_12 = big.NewInt(12)
var BIG_13 = big.NewInt(13)
var BIG_14 = big.NewInt(14)
var BIG_15 = big.NewInt(15)
var BIG_100 = big.NewInt(100)
var ONE_18 = big.NewInt(1e18)
var ONE_19, _ = new(big.Int).SetString("10000000000000000000", 10)                  // 2ˆ7
var ONE_20, _ = new(big.Int).SetString("100000000000000000000", 10)                 // 2ˆ7
var ONE_36, _ = new(big.Int).SetString("1000000000000000000000000000000000000", 10) // 2ˆ7
var LN_36_LOWER_BOUND = big.NewInt(1e18 - 1e17)
var LN_36_UPPER_BOUND = big.NewInt(1e18 + 1e17)

// 18 decimal constants
var x0, _ = new(big.Int).SetString("128000000000000000000", 10)                                    // 2ˆ7
var a0, _ = new(big.Int).SetString("38877084059945950922200000000000000000000000000000000000", 10) // eˆ(x0) (no decimals)
var a0_ONE18 = new(big.Int).Mul(a0, ONE_18)
var x1, _ = new(big.Int).SetString("64000000000000000000", 10)         // 2ˆ6
var a1, _ = new(big.Int).SetString("6235149080811616882910000000", 10) // eˆ(x1) (no decimals)
var a1_ONE18 = new(big.Int).Mul(a1, ONE_18)

// 20 decimal constants
var x2, _ = new(big.Int).SetString("3200000000000000000000", 10)             // 2ˆ5
var a2, _ = new(big.Int).SetString("7896296018268069516100000000000000", 10) // eˆ(x2)
var x3, _ = new(big.Int).SetString("1600000000000000000000", 10)             // 2ˆ4
var a3, _ = new(big.Int).SetString("888611052050787263676000000", 10)        // eˆ(x3)
var x4, _ = new(big.Int).SetString("800000000000000000000", 10)              // 2ˆ3
var a4, _ = new(big.Int).SetString("298095798704172827474000", 10)           // eˆ(x4)
var x5, _ = new(big.Int).SetString("400000000000000000000", 10)              // 2ˆ2
var a5, _ = new(big.Int).SetString("5459815003314423907810", 10)             // eˆ(x5)
var x6, _ = new(big.Int).SetString("200000000000000000000", 10)              // 2ˆ1
var a6, _ = new(big.Int).SetString("738905609893065022723", 10)              // eˆ(x6)
var x7, _ = new(big.Int).SetString("100000000000000000000", 10)              // 2ˆ0
var a7, _ = new(big.Int).SetString("271828182845904523536", 10)              // eˆ(x7)
var x8, _ = new(big.Int).SetString("50000000000000000000", 10)               // 2ˆ-1
var a8, _ = new(big.Int).SetString("164872127070012814685", 10)              // eˆ(x8)
var x9, _ = new(big.Int).SetString("25000000000000000000", 10)               // 2ˆ-2
var a9, _ = new(big.Int).SetString("128402541668774148407", 10)              // eˆ(x9)
var x10, _ = new(big.Int).SetString("12500000000000000000", 10)              // 2ˆ-3
var a10, _ = new(big.Int).SetString("113314845306682631683", 10)             // eˆ(x10)
var x11, _ = new(big.Int).SetString("6250000000000000000", 10)               // 2ˆ-4
var a11, _ = new(big.Int).SetString("106449445891785942956", 10)             // eˆ(x11)

//Can be changed to MulOverflow for safety
func divUp(a, b *uint256.Int) *uint256.Int {
	aInflated, overflow := new(uint256.Int).MulOverflow(a, ONE_E18)
	if overflow {
		log.Println("OVERFLOW in divup")
		return uint256.NewInt(0)
	}
	aInflated.Sub(aInflated, ONE)
	aInflated.Div(aInflated, b)
	aInflated.Add(aInflated, ONE)
	return aInflated
}

func divDown(a, b *uint256.Int) *uint256.Int {
	aInflated := new(uint256.Int).Mul(a, ONE_E18)
	return aInflated.Div(aInflated, b)
}

func mulDown(a, b *uint256.Int) *uint256.Int {
	product := new(uint256.Int).Mul(a, b)
	return product.Div(product, ONE_E18)
}

func mulUp(a, b *uint256.Int) *uint256.Int {
	product := new(uint256.Int).Mul(a, b)
	if product.IsZero() {
		return product
	} else {
		product.Sub(product, ONE)
		product.Div(product, ONE_E18)
		product.Add(product, ONE)
		return product
	}
}

func powUp(x, y *uint256.Int) *uint256.Int {
	if y.Cmp(ONE_E18) == 0 {
		return x
	} else if y.Cmp(TWO_E18) == 0 {
		return mulUp(x, x)
	} else if y.Cmp(FOUR_E18) == 0 {
		square := mulUp(x, x)
		return mulUp(square, square)
	} else {
		//THIS EXP IS SLOW AS FUCK, maybe rewrite in uint256?
		raw := pow(x, y)
		//fa, _ := new(big.Float).SetInt(x.ToBig()).Float64()
		//fb, _ := new(big.Float).SetInt(y.ToBig()).Float64()
		//ff := math.Pow(fa/1e18, fb/1e18) * 1e18
		//f := new(big.Float).SetFloat64(ff)
		//d := new(big.Int)
		//f.Int(d)
		//r, _ := uint256.FromBig(d)
		ans := new(uint256.Int).Add(mulUp(raw, MAX_POW_RELATIVE_ERROR), ONE)
		return raw.Add(raw, ans)
	}
}

//Overwrites x
func complement(x *uint256.Int) *uint256.Int {
	if x.Cmp(ONE_E18) < 0 {
		return x.Sub(ONE_E18, x)
	} else {
		return uint256.NewInt(0)
	}
}

func upscale(x, scallingFactor *uint256.Int) *uint256.Int {
	return mulDown(x, scallingFactor)
}

func downscaleDown(x, scallingFactor *uint256.Int) *uint256.Int {
	return divDown(x, scallingFactor)
}

func pow(x, y *uint256.Int) *uint256.Int {
	if y.IsZero() {
		return new(uint256.Int).Set(ONE_E18)
	}

	if x.IsZero() {
		return uint256.NewInt(0)
	}

	//THIS IS SLOW
	intX := x.ToBig()
	intY := y.ToBig()
	var logx_times_y *big.Int

	if intX.Cmp(LN_36_LOWER_BOUND) > 0 && intX.Cmp(LN_36_UPPER_BOUND) < 0 {
		ln_36_x := _ln_36(intX)

		//Can be optimized in the future given no overflows in bigInt
		//First term, reusing intX to avoid additional allocations
		intX.Quo(ln_36_x, ONE_18)
		intX.Mul(intX, intY)

		secondTerm := new(big.Int).Mod(ln_36_x, ONE_18)
		secondTerm.Mul(secondTerm, intY)
		secondTerm.Quo(secondTerm, ONE_18)
		logx_times_y = intX.Add(intX, secondTerm)
	} else {
		v := _ln(intX)
		logx_times_y = v.Mul(v, intY)
	}

	logx_times_y.Quo(logx_times_y, ONE_18)
	//Add a check later

	e, _ := uint256.FromBig(exp(logx_times_y))
	return e
}

func _ln_36(xIn *big.Int) *big.Int {
	x := new(big.Int).Mul(xIn, ONE_18)

	xSub := new(big.Int).Sub(x, ONE_36)
	xSub.Mul(xSub, ONE_36)
	xAdd := new(big.Int).Add(x, ONE_36)
	z := xSub.Quo(xSub, xAdd)
	zSquared := new(big.Int).Mul(z, z)
	zSquared.Quo(zSquared, ONE_36)

	num := new(big.Int).Set(z)
	seriesSum := new(big.Int).Set(num)

	temp := new(big.Int)
	num.Mul(num, zSquared)
	num.Quo(num, ONE_36)
	temp.Quo(num, BIG_3)
	seriesSum.Add(seriesSum, temp)

	num.Mul(num, zSquared)
	num.Quo(num, ONE_36)
	temp.Quo(num, BIG_5)
	seriesSum.Add(seriesSum, temp)

	num.Mul(num, zSquared)
	num.Quo(num, ONE_36)
	temp.Quo(num, BIG_7)
	seriesSum.Add(seriesSum, temp)

	num.Mul(num, zSquared)
	num.Quo(num, ONE_36)
	temp.Quo(num, BIG_9)
	seriesSum.Add(seriesSum, temp)

	num.Mul(num, zSquared)
	num.Quo(num, ONE_36)
	temp.Quo(num, BIG_11)
	seriesSum.Add(seriesSum, temp)

	num.Mul(num, zSquared)
	num.Quo(num, ONE_36)
	temp.Quo(num, BIG_13)
	seriesSum.Add(seriesSum, temp)

	num.Mul(num, zSquared)
	num.Quo(num, ONE_36)
	temp.Quo(num, BIG_15)
	seriesSum.Add(seriesSum, temp)

	seriesSum.Mul(seriesSum, BIG_2)
	return seriesSum
}

func _ln(a *big.Int) *big.Int {
	//fmt.Println("THIS LOG WAS TAKEN, Mostlikey implementation error in this function")
	if a.Cmp(ONE_18) < 0 {
		v := new(big.Int).Mul(ONE_18, ONE_18)
		v.Quo(v, a)
		x := _ln(v)
		return x.Neg(x)
	}
	sum := big.NewInt(0)

	if a.Cmp(a0_ONE18) >= 0 {
		a.Quo(a, a0)
		sum.Add(sum, x0)
	}

	if a.Cmp(a1_ONE18) >= 0 {
		a.Quo(a, a1)
		sum.Add(sum, x1)
	}
	sum.Mul(sum, BIG_100)
	a.Mul(a, BIG_100)

	if a.Cmp(a2) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a2)
		sum.Add(sum, x2)
	}

	if a.Cmp(a3) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a3)
		sum.Add(sum, x3)
	}

	if a.Cmp(a4) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a4)
		sum.Add(sum, x4)
	}

	if a.Cmp(a5) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a5)
		sum.Add(sum, x5)
	}

	if a.Cmp(a6) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a6)
		sum.Add(sum, x6)
	}

	if a.Cmp(a7) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a7)
		sum.Add(sum, x7)
	}

	if a.Cmp(a8) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a8)
		sum.Add(sum, x8)
	}

	if a.Cmp(a9) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a9)
		sum.Add(sum, x9)
	}

	if a.Cmp(a10) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a10)
		sum.Add(sum, x10)
	}

	if a.Cmp(a11) >= 0 {
		a.Mul(a, ONE_20)
		a.Quo(a, a11)
		sum.Add(sum, x11)
	}

	aSub := new(big.Int).Sub(a, ONE_20)
	aAdd := a.Add(a, ONE_20)
	z := aSub.Mul(aSub, ONE_20)
	z.Quo(z, aAdd)
	zSquared := aAdd.Mul(z, z)
	zSquared.Quo(zSquared, ONE_20)

	//One allocation here can be removed
	temp := new(big.Int)
	seriesSum := new(big.Int).Set(z)

	z.Mul(z, zSquared)
	z.Quo(z, ONE_20)
	temp.Quo(z, BIG_3)
	seriesSum.Add(seriesSum, temp)

	z.Mul(z, zSquared)
	z.Quo(z, ONE_20)
	temp.Quo(z, BIG_5)
	seriesSum.Add(seriesSum, temp)

	z.Mul(z, zSquared)
	z.Quo(z, ONE_20)
	temp.Quo(z, BIG_7)
	seriesSum.Add(seriesSum, temp)

	z.Mul(z, zSquared)
	z.Quo(z, ONE_20)
	temp.Quo(z, BIG_9)
	seriesSum.Add(seriesSum, temp)

	z.Mul(z, zSquared)
	z.Quo(z, ONE_20)
	temp.Quo(z, BIG_11)
	seriesSum.Add(seriesSum, temp)

	seriesSum.Mul(seriesSum, BIG_2)
	seriesSum.Add(sum, seriesSum)
	seriesSum.Quo(seriesSum, BIG_100)

	return seriesSum
}

func exp(xIn *big.Int) *big.Int {
	x := new(big.Int).Set(xIn)
	if x.Cmp(BIG_0) < 0 {
		v := new(big.Int).Mul(ONE_18, ONE_18)
		e := exp(x.Neg(x))
		return e.Quo(v, e)
	}

	firstAn := new(big.Int)
	if x.Cmp(x0) >= 0 {
		x.Sub(x, x0)
		firstAn.Set(a0)
	} else if x.Cmp(x1) >= 0 {
		x.Sub(x, x1)
		firstAn.Set(a1)
	} else {
		firstAn.Set(BIG_1)
	}

	x.Mul(x, BIG_100)

	product := new(big.Int).Set(ONE_20)

	if x.Cmp(x2) >= 0 {
		x.Sub(x, x2)
		product.Mul(product, a2)
		product.Quo(product, ONE_20)
	}

	if x.Cmp(x3) >= 0 {
		x.Sub(x, x3)
		product.Mul(product, a3)
		product.Quo(product, ONE_20)
	}

	if x.Cmp(x4) >= 0 {
		x.Sub(x, x4)
		product.Mul(product, a4)
		product.Quo(product, ONE_20)
	}
	if x.Cmp(x5) >= 0 {
		x.Sub(x, x5)
		product.Mul(product, a5)
		product.Quo(product, ONE_20)
	}
	if x.Cmp(x6) >= 0 {
		x.Sub(x, x6)
		product.Mul(product, a6)
		product.Quo(product, ONE_20)
	}
	if x.Cmp(x7) >= 0 {
		x.Sub(x, x7)
		product.Mul(product, a7)
		product.Quo(product, ONE_20)
	}
	if x.Cmp(x8) >= 0 {
		x.Sub(x, x8)
		product.Mul(product, a8)
		product.Quo(product, ONE_20)
	}
	if x.Cmp(x9) >= 0 {
		x.Sub(x, x9)
		product.Mul(product, a9)
		product.Quo(product, ONE_20)
	}

	seriesSum := new(big.Int).Set(ONE_20)
	term := new(big.Int).Set(x)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_2)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_3)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_4)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_5)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_6)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_7)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_8)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_9)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_10)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_11)
	seriesSum.Add(seriesSum, term)

	term.Mul(term, x)
	term.Quo(term, ONE_20)
	term.Quo(term, BIG_12)
	seriesSum.Add(seriesSum, term)

	product.Mul(product, seriesSum)
	product.Quo(product, ONE_20)
	product.Mul(product, firstAn)
	product.Quo(product, BIG_100)
	return product
}
