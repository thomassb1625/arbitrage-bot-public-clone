package uniswapV3

import (
	"errors"

	"github.com/holiman/uint256"
)

// Errors for tick math functions
var (
	ErrInvalidInput     = errors.New("invalid input")
	ErrInvalidSqrtRatio = errors.New("invalid sqrt ratio")
)

var (
	uint256AbsTick, _        = uint256.FromHex("0x100000000000000000000000000000000")
	uint256AbsTick0x1, _     = uint256.FromHex("0xfffcb933bd6fad37aa2d162d1a594001")
	uint256AbsTick0x2, _     = uint256.FromHex("0xfff97272373d413259a46990580e213a")
	uint256AbsTick0x4, _     = uint256.FromHex("0xfff2e50f5f656932ef12357cf3c7fdcc")
	uint256AbsTick0x8, _     = uint256.FromHex("0xffe5caca7e10e4e61c3624eaa0941cd0")
	uint256AbsTick0x10, _    = uint256.FromHex("0xffcb9843d60f6159c9db58835c926644")
	uint256AbsTick0x20, _    = uint256.FromHex("0xff973b41fa98c081472e6896dfb254c0")
	uint256AbsTick0x40, _    = uint256.FromHex("0xff2ea16466c96a3843ec78b326b52861")
	uint256AbsTick0x80, _    = uint256.FromHex("0xfe5dee046a99a2a811c461f1969c3053")
	uint256AbsTick0x100, _   = uint256.FromHex("0xfcbe86c7900a88aedcffc83b479aa3a4")
	uint256AbsTick0x200, _   = uint256.FromHex("0xf987a7253ac413176f2b074cf7815e54")
	uint256AbsTick0x400, _   = uint256.FromHex("0xf3392b0822b70005940c7a398e4b70f3")
	uint256AbsTick0x800, _   = uint256.FromHex("0xe7159475a2c29b7443b29c7fa6e889d9")
	uint256AbsTick0x1000, _  = uint256.FromHex("0xd097f3bdfd2022b8845ad8f792aa5825")
	uint256AbsTick0x2000, _  = uint256.FromHex("0xa9f746462d870fdf8a65dc1f90e061e5")
	uint256AbsTick0x4000, _  = uint256.FromHex("0x70d869a156d2a1b890bb3df62baf32f7")
	uint256AbsTick0x8000, _  = uint256.FromHex("0x31be135f97d08fd981231505542fcfa6")
	uint256AbsTick0x10000, _ = uint256.FromHex("0x9aa508b5b7a84e1c677de54f3e99bc9")
	uint256AbsTick0x20000, _ = uint256.FromHex("0x5d6af8dedb81196699c329225ee604")
	uint256AbsTick0x40000, _ = uint256.FromHex("0x2216e584f5fa1ea926041bedfe98")
	uint256AbsTick0x80000, _ = uint256.FromHex("0x48a170391f7dc42444e8fa2")
)

func getSqrtRatioAtTick(tick int) uint256.Int {
	absTick := tick
	if tick < 0 {
		absTick = -tick
	}

	var ratio *uint256.Int
	if absTick&0x1 != 0 {
		ratio = uint256AbsTick0x1
	} else {
		ratio = uint256AbsTick
	}
	if absTick&0x2 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x2), 128)
	}
	if absTick&0x4 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x4), 128)
	}
	if absTick&0x8 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x8), 128)
	}
	if absTick&0x10 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x10), 128)
	}
	if absTick&0x20 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x20), 128)
	}
	if absTick&0x40 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x40), 128)
	}
	if absTick&0x80 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x80), 128)
	}
	if absTick&0x100 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x100), 128)
	}
	if absTick&0x200 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x200), 128)
	}
	if absTick&0x400 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x400), 128)
	}
	if absTick&0x800 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x800), 128)
	}
	if absTick&0x1000 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x1000), 128)
	}
	if absTick&0x2000 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x2000), 128)
	}
	if absTick&0x4000 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x4000), 128)
	}
	if absTick&0x8000 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x8000), 128)
	}
	if absTick&0x10000 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x10000), 128)
	}
	if absTick&0x20000 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x20000), 128)
	}
	if absTick&0x40000 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x40000), 128)
	}
	if absTick&0x80000 != 0 {
		ratio = new(uint256.Int).Rsh(new(uint256.Int).Mul(ratio, uint256AbsTick0x80000), 128)
	}

	if tick > 0 {
		ratio = new(uint256.Int).Div(MaxUint256, ratio)
	}

	// back to Q96
	if new(uint256.Int).Mod(ratio, Q32Uint256).Eq(ZeroUint256) {
		return *new(uint256.Int).Rsh(ratio, 32)
	}
	return *new(uint256.Int).Add(new(uint256.Int).Rsh(ratio, 32), OneUint256)
}

func getTickAtSqrtRatio(sqrtRatioX96 *uint256.Int) (int, error) {
	if sqrtRatioX96.Cmp(MinSqrtRatio) < 0 || sqrtRatioX96.Cmp(MaxSqrtRatio) >= 0 {
		return 0, ErrInvalidSqrtRatio
	}
	sqrtRatioX128 := new(uint256.Int).Lsh(sqrtRatioX96, 32)
	msb := mostSignificantBit(sqrtRatioX128)

	var r *uint256.Int
	if uint256.NewInt(msb).Cmp(uint256.NewInt(128)) >= 0 {
		r = new(uint256.Int).Rsh(sqrtRatioX128, uint(msb-127))
	} else {
		r = new(uint256.Int).Lsh(sqrtRatioX128, uint(127-msb))
	}

	//log2 := new(uint256.Int).Lsh(new(uint256.Int).Sub(uint256.NewInt(msb), uint256.NewInt(128)), 64)
	log2 := new(uint256.Int).Sub(uint256.NewInt(msb), uint256.NewInt(128))
	log2.Lsh(log2, 64)

	for i := 0; i < 14; i++ {
		//_r := new(uint256.Int).Rsh(new(uint256.Int).Mul(r, r), 127)
		_r := new(uint256.Int).Mul(r, r)
		_r.Rsh(_r, 127)
		*r = *_r
		f := new(uint256.Int).Rsh(r, 128)
		_log2 := new(uint256.Int).Or(log2, new(uint256.Int).Lsh(f, uint(63-i)))
		*log2 = *_log2
		__r2 := new(uint256.Int).Rsh(r, uint(f.Uint64()))
		*r = *__r2
	}

	logSqrt10001 := new(uint256.Int).Mul(log2, MagicSqrt10001)

	//tickLow := new(uint256.Int).Rsh(new(uint256.Int).Sub(logSqrt10001, MagicTickLow), 128).Uint64()
	temp := new(uint256.Int).Sub(logSqrt10001, MagicTickLow)
	tickLow := temp.Rsh(temp, 128).Uint64()
	//tickHigh := new(uint256.Int).Rsh(new(uint256.Int).Add(logSqrt10001, MagicTickHigh), 128).Uint64()
	temp.Add(logSqrt10001, MagicTickHigh)
	tickHigh := temp.Rsh(temp, 128).Uint64()

	if tickLow == tickHigh {
		return int(tickLow), nil
	}

	sqrtRatio := getSqrtRatioAtTick(int(tickHigh))
	//And to make 24bit
	if sqrtRatio.Cmp(sqrtRatioX96) <= 0 {
		return int(tickHigh) & 8388607, nil
	}

	return int(tickLow) & 8388607, nil
}

func mostSignificantBit(x *uint256.Int) uint64 {
	//	if x.Cmp(ZeroUint256) <= 0 {
	//		return 0, ErrInvalidInput
	//	}
	//	if x.Cmp(MaxUint256) > 0 {
	//		return 0, ErrInvalidInput
	//	}
	var msb uint64
	min := new(uint256.Int)
	for _, power := range []uint64{128, 64, 32, 16, 8, 4, 2, 1} {
		min.Exp(uint256.NewInt(2), uint256.NewInt(power))
		if x.Cmp(min) >= 0 {
			_x := min.Rsh(x, uint(power))
			*x = *_x
			msb += power
		}
	}
	return msb
}

var (
	maxu128, _ = uint256.FromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	maxu64, _  = uint256.FromHex("0xFFFFFFFFFFFFFFFF")
	maxu32, _  = uint256.FromHex("0xFFFFFFFF")
	maxu16, _  = uint256.FromHex("0xFFFF")
	maxu8, _   = uint256.FromHex("0xFF")
	maxu4, _   = uint256.FromHex("0xF")
	maxu2, _   = uint256.FromHex("0x3")
	maxu1, _   = uint256.FromHex("0x1")
)

func leastSignificantBit(x *uint256.Int) uint64 {
	r := uint8(255)
	temp := new(uint256.Int)
	if !temp.And(x, maxu128).IsZero() {
		r -= 128
	} else {
		x.Rsh(x, 128)
	}

	if !temp.And(x, maxu64).IsZero() {
		r -= 64
	} else {
		x.Rsh(x, 64)
	}

	if !temp.And(x, maxu32).IsZero() {
		r -= 32
	} else {
		x.Rsh(x, 32)
	}

	if !temp.And(x, maxu16).IsZero() {
		r -= 16
	} else {
		x.Rsh(x, 16)
	}

	if !temp.And(x, maxu8).IsZero() {
		r -= 8
	} else {
		x.Rsh(x, 8)
	}

	if !temp.And(x, maxu4).IsZero() {
		r -= 4
	} else {
		x.Rsh(x, 4)
	}

	if !temp.And(x, maxu2).IsZero() {
		r -= 2
	} else {
		x.Rsh(x, 2)
	}

	if !temp.And(x, maxu1).IsZero() {
		r -= 1
	}

	return uint64(r)
}
