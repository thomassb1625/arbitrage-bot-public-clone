package decimalMath

import (
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/dex/dodov2/safeMath"
)

var ONE, _ = uint256.FromHex("0xDE0B6B3A7640000")
var ONE2, _ = uint256.FromHex("0xC097CE7BC90715B34B9F1000000000")

func MulFloor(target, d *uint256.Int) *uint256.Int {
	v := new(uint256.Int).Mul(target, d)
	return v.Div(v, ONE)
}

func MulCeil(target, d *uint256.Int) *uint256.Int {
	v := new(uint256.Int).Mul(target, d)
	v = safeMath.DivCeil(v, ONE)
	return v
}

func DivFloor(target, d *uint256.Int) *uint256.Int {
	v := new(uint256.Int).Mul(target, ONE)
	return v.Div(v, d)
}

func DivCeil(target, d *uint256.Int) *uint256.Int {
	v := new(uint256.Int).Mul(target, ONE)
	v = safeMath.DivCeil(v, d)
	return v
}

func ReciprocalFloor(target *uint256.Int) *uint256.Int {
	v := new(uint256.Int).Set(ONE2)
	v.Div(v, target)
	return v
}

func ReciprocalCeil(target *uint256.Int) *uint256.Int {
	v := new(uint256.Int).Set(ONE2)
	v = safeMath.DivCeil(v, target)
	return v
}
