package uniswapV3

import (
	"math/big"

	"github.com/holiman/uint256"
)

var big0 = big.NewInt(0)

func (p *Pair) burn(tickLower, tickUpper int, amount *big.Int) {
	//DON'T switch the order, AND only works for positive numbers
	amount.And(amount, bigu128)
	amount.Mul(amount, NegativeOneBig)
	p._modifyPosition(tickLower, tickUpper, amount)

}

func (p *Pair) mint(tickLower, tickUpper int, amount *big.Int) {
	amount.And(amount, bigu128)
	p._modifyPosition(tickLower, tickUpper, amount)
}

func (p *Pair) _modifyPosition(tickLower, tickUpper int, liquidityDelta *big.Int) {
	p._updatePosition(tickLower, tickUpper, liquidityDelta, p.Tick)
	//Implement this
	if liquidityDelta.Cmp(big0) != 0 {
		if p.Tick < tickLower {
		} else if p.Tick < tickUpper {
			//liquidityBefore, _ := new(big.Int).Set(p.liquidity)
			//v, _ := uint256.FromBig(new(big.Int).Mul(liquidityDelta, NegativeOneBig))
			p.Liquidity.Set(addDelta(p.Liquidity, liquidityDelta))
		} else {
		}
	}
}

func (p *Pair) _updatePosition(tickLower, tickUpper int, liquidityDelta *big.Int, tick int) {
	var flippedLower bool
	var flippedUpper bool

	if liquidityDelta.Cmp(big0) != 0 {
		//tick false update
		flippedLower = p.updateTick(tickLower, tick, liquidityDelta, false)
		//tick true updae
		flippedUpper = p.updateTick(tickUpper, tick, liquidityDelta, true)

		//		if flippedLower {
		//			//if tick%p.tickSpacing != 0 {
		//			//	panic("TICK IS NOT SPACED")
		//			//}
		//			wordPos, bitPos := position(tickLower / p.TickSpacing)
		//			mask := new(uint256.Int).Lsh(OneUint256, uint(bitPos))
		//			v := p.Tickbitmap[int16(wordPos)]
		//			p.Tickbitmap[int16(wordPos)] = *new(uint256.Int).Xor(&v, mask)
		//		}
		//		if flippedUpper {
		//			//if tick%p.tickSpacing != 0 {
		//			//	panic("TICK IS NOT SPACED")
		//			//}
		//			wordPos, bitPos := position(tickUpper / p.TickSpacing)
		//			mask := new(uint256.Int).Lsh(OneUint256, uint(bitPos))
		//			v := p.Tickbitmap[int16(wordPos)]
		//			p.Tickbitmap[int16(wordPos)] = *new(uint256.Int).Xor(&v, mask)
		//		}
	}

	if liquidityDelta.Cmp(big0) < 0 {
		if flippedLower {
			//Tick clear
			p.clearTick(tickLower)
		}
		if flippedUpper {
			p.clearTick(tickUpper)
		}
	}

}

//Most likely this is broken
func (p *Pair) updateTick(tick int, tickCurrent int, liquidityDelta *big.Int, upper bool) bool {
	info := p.Ticks[tick]
	lg := info.LiquidityGross
	liquidityGrossBefore := new(uint256.Int).And(&lg, uintu128)
	liquidityGrossAfter := addDelta(liquidityGrossBefore, liquidityDelta)
	//liquidityGrossAfter.And(liquidityGrossAfter, uintu128)
	flipped := liquidityGrossAfter.IsZero() != liquidityGrossBefore.IsZero()
	if (liquidityGrossBefore).IsZero() {
		info.Initialized = true
	}

	info.LiquidityGross = *liquidityGrossAfter
	if upper {
		info.LiquidityNet = *new(big.Int).Sub(&info.LiquidityNet, liquidityDelta)
		//fmt.Printf("%v tick upper Liqned = %v\n", tick, &info.LiquidityNet)
	} else {
		info.LiquidityNet = *new(big.Int).Add(&info.LiquidityNet, liquidityDelta)
		//fmt.Printf("%v tick NOT Liqned = %v\n", tick, &info.LiquidityNet)
	}
	p.Ticks[tick] = info
	return flipped
}

var bigu128, _ = big.NewInt(0).SetString("340282366920938463463374607431768211455", 10)
var uintu128, _ = uint256.FromBig(bigu128)

func addDelta(x *uint256.Int, liquidityDetla *big.Int) *uint256.Int {
	if liquidityDetla.Cmp(ZeroBig) < 0 {
		xBig := x.ToBig()
		ld := new(big.Int).Mul(liquidityDetla, NegativeOneBig)
		ld.And(ld, bigu128)
		delta, _ := uint256.FromBig(new(big.Int).Sub(xBig, ld))
		return delta
	} else {
		ld := new(big.Int).And(liquidityDetla, bigu128)
		liquidityNetBig, _ := uint256.FromBig(ld)
		delta := new(uint256.Int).Add(x, liquidityNetBig)
		return delta
	}
}

func addDeltaOld(x *uint256.Int, y *big.Int) *uint256.Int {
	if y.Cmp(big0) >= 0 {
		uy, _ := uint256.FromBig(y)
		d, e := new(uint256.Int).AddOverflow(x, uy)
		if e {
			panic("OVERFLOW IN ADD")
		}
		return d
	} else {
		uy, _ := uint256.FromBig(new(big.Int).Mul(y, NegativeOneBig))
		d, e := new(uint256.Int).SubOverflow(x, uy)
		if e {
			panic("OVERFLOW IN Sub")
		}
		return d
	}
}

func (p *Pair) clearTick(tick int) {
	delete(p.Ticks, tick)
}
