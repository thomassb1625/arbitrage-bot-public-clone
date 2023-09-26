package uniswapV3

import (
	"math/big"

	"github.com/holiman/uint256"
)

var (
	// MinTick The minimum tick that may be passed to #getSqrtRatioAtTick computed from log base 1.0001 of 2**-128
	MinTick = -887272
	// MaxTick The maximum tick that may be passed to #getSqrtRatioAtTick computed from log base 1.0001 of 2**128
	MaxTick = -MinTick
)

var (
	// NegativeOneBig is -1 in big.Int
	NegativeOneBig = big.NewInt(-1)
	// ZeroBig is 0 in big.Int
	ZeroBig = big.NewInt(0)
)

var (
	// FeeMultiplier uint256.Int multiplier for calculating fees
	FeeMultiplier = uint256.NewInt(1000000)
	// ZeroUint256 is 0 in uint256.Int
	ZeroUint256 = uint256.NewInt(0)
	// OneUint256 is 1 in uint256.Int
	OneUint256 = uint256.NewInt(1)
	// Q32Uint256 is the q32 representation in uint256.Int
	Q32Uint256 = uint256.NewInt(1 << 32)
	// MaxUint256 is the maximum value for an uint256.Int
	MaxUint256, _ = uint256.FromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	// MinSqrtRatio The minimum value that can be returned from #getSqrtRatioAtTick. Equivalent to getSqrtRatioAtTick(MIN_TICK)
	MinSqrtRatio = uint256.NewInt(4295128739)
	// MinSqrtRatioAddOne is MinSqrtRatio + 1
	MinSqrtRatioAddOne = uint256.NewInt(4295128740)
	// MaxSqrtRatio The maximum value (1461446703485210103287273052203988822378723970342) that can be returned from #getSqrtRatioAtTick. Equivalent to getSqrtRatioAtTick(MAX_TICK)
	MaxSqrtRatio, _ = uint256.FromHex("0xFFFD8963EFD1FC6A506488495D951D5263988D26")
	// MaxSqrtRatioSubOne is MaxSqrtRatio - 1 (1461446703485210103287273052203988822378723970341)
	MaxSqrtRatioSubOne, _ = uint256.FromHex("0xFFFD8963EFD1FC6A506488495D951D5263988D25")
	// Q96Uint256 is the Q96 (79228162514264337593543950336) representation in uint256.Int
	Q96Uint256, _ = uint256.FromHex("0x1000000000000000000000000")
	// MaxUint160 max value for an uint160 (1461501637330902918203684832716283019655932542975) in uint256.Int
	MaxUint160, _ = uint256.FromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	// MagicSqrt10001 tick-math magic (255738958999603826347141)
	MagicSqrt10001, _ = uint256.FromHex("0x3627A301D71055774C85")
	// MagicTickLow tick-math magic (3402992956809132418596140100660247210)
	MagicTickLow, _ = uint256.FromHex("0x28F6481AB7F045A5AF012A19D003AAA")
	// MagicTickHigh tick-math magic (291339464771989622907027621153398088495)
	MagicTickHigh, _ = uint256.FromHex("0xDB2DF09E81959A81455E260799A0632F")
)
