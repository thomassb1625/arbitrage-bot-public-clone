package safeMath

import "github.com/holiman/uint256"

func DivCeil(a, b *uint256.Int) *uint256.Int {
	q := new(uint256.Int).Div(a, b)
	remainder := new(uint256.Int).Sub(a, new(uint256.Int).Mul(q, b))
	if remainder.IsZero() {
		return q
	} else {
		return q.Add(q, uint256.NewInt(1))
	}
}

func Sqrt(x *uint256.Int) *uint256.Int {
	z := new(uint256.Int).Div(x, uint256.NewInt(2))
	z.Add(z, uint256.NewInt(1))
	y := new(uint256.Int).Set(x)
	for y.Cmp(z) > 0 {
		y.Set(z)
		temp := new(uint256.Int).Div(x, z)
		z.Add(temp, z)
		z.Div(z, uint256.NewInt(2))
	}
	return y
}
