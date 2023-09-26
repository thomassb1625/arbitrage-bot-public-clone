package curve

import (
	"github.com/holiman/uint256"
)

func (pair Pair) BasePool_Dy(i, j uint64, dx *uint256.Int) *uint256.Int {
	rates := pair.rates
	xp := pair.base_xp(pair.rates, pair.reserves, pair.LENDING_PRECISION)

	t2 := new(uint256.Int).Mul(dx, rates[i]) //(dx*rates[i]/PRECISION)
	t2.Div(t2, pair.PRECISION)
	x := new(uint256.Int).Add(xp[i], t2)

	y := pair.base_get_Y(i, j, x, xp)
	zero := uint256.NewInt(0)
	if y.Cmp(zero) == 0 {
		return zero
	}

	// Formula before conversion
	// dy := (xp[j] - y - 1) * PRECISION / rates[j]
	dy := new(uint256.Int).Sub(xp[j], y)
	dy.Sub(dy, uint256.NewInt(1))
	dy.Mul(dy, pair.PRECISION)
	dy.Div(dy, rates[j])

	// Formula before conversion
	// fee := pair.fee * dy / FEE_DENOMINATOR
	_fee := new(uint256.Int).Mul(pair.fee, dy)
	_fee.Div(_fee, pair.FEE_DENOMINATOR)
	_fee.Sub(dy, _fee)

	return _fee
}

func (pair Pair) base_xp(rates []*uint256.Int, balances []*uint256.Int, LENDING_PRECISION *uint256.Int) []*uint256.Int {
	result := make([]*uint256.Int, len(rates))
	for i := 0; i < len(balances); i++ {
		result[i] = new(uint256.Int).Mul(rates[i], balances[i])
		result[i].Div(result[i], LENDING_PRECISION)
	}
	return result
}

func (pair Pair) base_get_Y(i, j uint64, x *uint256.Int, xp []*uint256.Int) *uint256.Int {
	amp := pair.A
	D := pair.base_get_D(xp, amp)
	c := D.Clone()
	S_ := uint256.NewInt(0)
	Ann := new(uint256.Int).Mul(amp, pair.N_COINS)

	_x := uint256.NewInt(0)
	for _i := uint64(0); _i < uint64(3); _i++ {
		if _i == i {
			_x.Set(x)
		} else if _i != j {
			_x.Set(xp[_i])
		} else {
			continue
		}
		S_.Add(S_, _x)

		// Formula before conversion
		//c = c * D / (_x * N_COINS)
		numerator := new(uint256.Int).Mul(c, D)
		denominator := new(uint256.Int).Mul(_x, pair.N_COINS)
		c.Set(new(uint256.Int).Div(numerator, denominator))
	}

	// Formula before conversion
	// c = c * D / (Ann * N_COINS)
	numerator := new(uint256.Int).Mul(c, D)
	denominator := new(uint256.Int).Mul(Ann, pair.N_COINS)
	c.Set(new(uint256.Int).Div(numerator, denominator))

	// Formula before conversion
	// b := S_ + D/Ann
	adder := new(uint256.Int).Div(D, Ann)
	b := new(uint256.Int).Add(S_, adder)
	y_prev := uint256.NewInt(0)
	y := D.Clone()

	for i := 0; i < 255; i++ {
		y_prev.Set(y)

		// Formula before conversion
		// y = (y*y + c) / (2 * y + b - D)
		n1 := new(uint256.Int).Mul(y, y)
		n1.Add(n1, c)
		d1 := new(uint256.Int).Mul(uint256.NewInt(2), y)
		d1.Add(d1, b)
		d1.Sub(d1, D)
		y.Set(new(uint256.Int).Div(n1, d1))

		temp := uint256.NewInt(0)
		if y.Cmp(y_prev) > 0 {
			temp.Sub(y, y_prev)
		} else {
			temp.Sub(y_prev, y)
		}
		if temp.Cmp(uint256.NewInt(1)) <= 0 {
			break
		}

	}

	return y
}

func (pair Pair) base_get_D(xp []*uint256.Int, amp *uint256.Int) *uint256.Int {
	S := uint256.NewInt(0)
	for _, _x := range xp {
		S.Add(S, _x)
	}
	if S.IsZero() {
		return uint256.NewInt(0) //It's zero
	}
	Dprev := uint256.NewInt(0)
	D := S.Clone()
	Ann := new(uint256.Int).Mul(amp, pair.N_COINS)

	for _i := 0; _i < 255; _i++ {
		D_P := D.Clone()
		for _, _x := range xp {
			D_P.Mul(D_P, D)

			div := new(uint256.Int).Mul(_x, pair.N_COINS)
			if div.IsZero() {
				panic("Division by zero, look up the docs to see what we need to do then")
			}
			D_P.Div(D_P, div) // D_P = D_P * D / (_x * N_COINS)
		}
		Dprev.Set(D) // Dprev = D

		num := new(uint256.Int).Mul(Ann, S)
		numerator := new(uint256.Int).Mul(D, num.Add(num, new(uint256.Int).Mul(D_P, pair.N_COINS))) // Ann*S + D_P * N_COINS * D
		denominator := new(uint256.Int).Sub(Ann, uint256.NewInt(1))
		denominator.Mul(denominator, D)
		denominator.Add(denominator, new(uint256.Int).Mul((new(uint256.Int).Add(pair.N_COINS, uint256.NewInt(1))), D_P))
		D.Div(numerator, denominator) // (Ann * S + D_P * N_COINS) * D / ((Ann - 1) * D + (N_COINS + 1) * D_P)

		temp := uint256.NewInt(0)
		if D.Cmp(Dprev) >= 0 {
			temp.Sub(D, Dprev)
		} else {
			temp.Sub(Dprev, D)
		}
		if temp.Cmp(uint256.NewInt(1)) <= 0 {
			break
		}
	}

	return D
}
