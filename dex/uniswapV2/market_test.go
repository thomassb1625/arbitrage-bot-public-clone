package uniswapV2

import (
	"math/rand"
	"testing"

	"github.com/holiman/uint256"
)

func BenchmarkDirect(b *testing.B) {
	reserveIn := uint256.NewInt(1e17)
	reserveOut := uint256.NewInt(7e15)
	rand.Seed(1)
	var v *uint256.Int
	for i := 0; i < b.N; i++ {
		v = uint256.NewInt(uint64(rand.Float64() * 1000))
		v.Mul(v, uint256.NewInt(1e16))
		ver2(v, reserveIn, reserveOut)
	}
	//fmt.Printf("\n\nOUT = %v = %v\n\n", ver2(v, reserveIn, reserveOut), ver1(v, reserveIn, reserveOut))
}

func BenchmarkInDirect(b *testing.B) {
	reserveIn := uint256.NewInt(1e17)
	reserveOut := uint256.NewInt(7e15)
	rand.Seed(1)
	var v *uint256.Int
	for i := 0; i < b.N; i++ {
		v = uint256.NewInt(uint64(rand.Float64() * 1000))
		v.Mul(v, uint256.NewInt(1e16))
		InDirect(v, reserveIn, reserveOut)
	}
	//fmt.Printf("\n\nOUT = %v = %v\n\n", ver2(v, reserveIn, reserveOut), ver1(v, reserveIn, reserveOut))
}

func BenchmarkV2(b *testing.B) {
	reserveIn := uint256.NewInt(1e17)
	reserveOut := uint256.NewInt(7e15)
	rand.Seed(1)
	var v *uint256.Int
	for i := 0; i < b.N; i++ {
		v = uint256.NewInt(uint64(rand.Float64() * 1000))
		v.Mul(v, uint256.NewInt(1e16))
		ver2(v, reserveIn, reserveOut)
	}
}

func ver1(amountIn, reserveIn, reserveOut *uint256.Int) *uint256.Int {
	amountInWithFee := uint256.NewInt(997)
	amountInWithFee.Mul(amountInWithFee, amountIn)
	denominator := uint256.NewInt(1000)
	denominator.Mul(reserveIn, denominator)
	denominator.Add(denominator, amountInWithFee)
	amountInWithFee.Mul(amountInWithFee, reserveOut)
	return amountInWithFee.Div(amountInWithFee, denominator)
}

func ver2(amountIn, reserveIn, reserveOut *uint256.Int) *uint256.Int {
	amountInWithFee := uint256.NewInt(997)
	amountInWithFee.Mul(amountInWithFee, amountIn)
	denominator := uint256.NewInt(1000)
	denominator.Mul(reserveIn, denominator)
	denominator.Add(denominator, amountInWithFee)
	numerator := amountInWithFee.Mul(amountInWithFee, reserveOut)
	return numerator.Div(numerator, denominator)
}

func InDirect(amountIn, reserveIn, reserveOut *uint256.Int) *uint256.Int {
	return ver2(reserveIn, reserveOut, amountIn)
}
