package dodoMath

import (
	DecimalMath "github.com/jmjac/ethGoTest/dex/dodov2/decimalMath"
	SafeMath "github.com/jmjac/ethGoTest/dex/dodov2/safeMath"

	"github.com/holiman/uint256"
)

func GeneralIntegrate(
	V0 *uint256.Int,
	V1 *uint256.Int,
	V2 *uint256.Int,
	i *uint256.Int,
	k *uint256.Int) *uint256.Int {

	if V0.Cmp(uint256.NewInt(0)) <= 0 {
		return nil
	}

	fairAmount := new(uint256.Int).Sub(V1, V2)
	fairAmount.Mul(i, fairAmount)

	if k.Cmp(uint256.NewInt(0)) == 0 {
		return fairAmount.Div(fairAmount, DecimalMath.ONE)
	}

	V0V0V1V2 := new(uint256.Int).Mul(V0, V0)
	V0V0V1V2.Div(V0V0V1V2, V1)
	V0V0V1V2 = DecimalMath.DivFloor(V0V0V1V2, V2)

	penalty := DecimalMath.MulFloor(k, V0V0V1V2)

	final := new(uint256.Int).Sub(DecimalMath.ONE, k)
	final.Add(final, penalty)
	final.Mul(final, fairAmount)

	return final.Div(final, DecimalMath.ONE2)
}

func SolveQuadraticFunctionForTarget(
	V1 *uint256.Int,
	delta *uint256.Int,
	i *uint256.Int,
	k *uint256.Int) *uint256.Int {
	if V1.Cmp(uint256.NewInt(0)) == 0 {
		return uint256.NewInt(0)
	}
	if k.Cmp(uint256.NewInt(0)) == 0 {
		return new(uint256.Int).Add(V1, DecimalMath.MulFloor(i, delta))
	}
	// V0 = V1*(1+(sqrt-1)/2k)
	// sqrt = √(1+4kidelta/V1)
	// premium = 1+(sqrt-1)/2k
	// uint256 sqrt = (4 * k).mul(i).mul(delta).div(V1).add(DecimalMath.ONE2).sqrt();
	sqrt := new(uint256.Int)
	ki := new(uint256.Int).Mul(uint256.NewInt(4), k)
	ki.Mul(ki, i)

	if ki.Cmp(uint256.NewInt(0)) == 0 {
		sqrt = DecimalMath.ONE

	} else if ki.Cmp(uint256.NewInt(1)) == 0 {
		//sqrt = (ki * delta).div(V1).add(DecimalMath.ONE2).sqrt();
		sqrt.Mul(ki, delta)
		sqrt.Div(sqrt, V1)
		sqrt.Add(sqrt, DecimalMath.ONE2)
		sqrt = SafeMath.Sqrt(sqrt)
	} else {
		//sqrt = ki.div(V1).mul(delta).add(DecimalMath.ONE2).sqrt();
		sqrt.Div(ki, V1)
		sqrt.Mul(sqrt, delta)
		sqrt.Add(sqrt, DecimalMath.ONE2)
		sqrt = SafeMath.Sqrt(sqrt)
	}
	//premium = DecimalMath.divFloor(sqrt.sub(DecimalMath.ONE), k * 2).add(DecimalMath.ONE);

	premium := new(uint256.Int).Sub(sqrt, DecimalMath.ONE)
	premium = DecimalMath.DivFloor(premium, k.Mul(k, uint256.NewInt(2)))
	premium.Add(premium, DecimalMath.ONE)
	// V0 is greater than or equal to V1 according to the solution
	return DecimalMath.MulFloor(V1, premium)
}

func SolveQuadraticFunctionForTrade(
	V0 *uint256.Int,
	V1 *uint256.Int,
	delta *uint256.Int,
	i *uint256.Int,
	k *uint256.Int) *uint256.Int {

	if V0.Cmp(uint256.NewInt(0)) <= 0 {
		return nil
	}

	if delta.Cmp(uint256.NewInt(0)) == 0 {
		return uint256.NewInt(0)
	}

	if k.Cmp(uint256.NewInt(0)) == 0 {
		//return DecimalMath.mulFloor(i, delta) > V1 ? V1 : DecimalMath.mulFloor(i, delta);
		condition := DecimalMath.MulFloor(i, delta)

		if condition.Cmp(V1) > 0 {
			return V1
		} else {
			return condition
		}
	}

	if k.Cmp(DecimalMath.ONE) == 0 {
		// if k==1
		// Q2=Q1/(1+ideltaBQ1/Q0/Q0)
		// temp = ideltaBQ1/Q0/Q0
		// Q2 = Q1/(1+temp)
		// Q1-Q2 = Q1*(1-1/(1+temp)) = Q1*(temp/(1+temp))
		// uint256 temp = i.mul(delta).mul(V1).div(V0.mul(V0));
		temp := new(uint256.Int)

		idelta := new(uint256.Int).Mul(i, delta)

		cond2 := new(uint256.Int).Mul(idelta, V1)
		cond2.Div(cond2, idelta)

		if idelta.Cmp(uint256.NewInt(0)) == 0 {
			temp = uint256.NewInt(0)
		} else if V1.Cmp(cond2) == 0 {
			//temp = (idelta * V1).div(V0.mul(V0));
			temp.Mul(idelta, V1)
			temp.Div(temp, new(uint256.Int).Mul(V0, V0))
		} else {
			//temp = delta.mul(V1).div(V0).mul(i).div(V0);
			temp.Mul(delta, V1)
			temp.Div(temp, V0)
			temp.Mul(temp, i)
			temp.Div(temp, V0)
		}

		//return V1.mul(temp).div(temp.add(DecimalMath.ONE));
		V1.Mul(V1, temp)
		return V1.Div(V1, temp.Add(temp, DecimalMath.ONE))
	}

	// calculate -b value and sig
	// b = kQ0^2/Q1-i*deltaB-(1-k)Q1
	// part1 = (1-k)Q1 >=0
	// part2 = kQ0^2/Q1-i*deltaB >=0
	// bAbs = abs(part1-part2)
	// if part1>part2 => b is negative => bSig is false
	// if part2>part1 => b is positive => bSig is true

	//uint256 part2 = k.mul(V0).div(V1).mul(V0).add(i.mul(delta)); // kQ0^2/Q1-i*deltaB
	//uint256 bAbs = DecimalMath.ONE.sub(k).mul(V1); // (1-k)Q1

	part2 := new(uint256.Int).Mul(k, V0)
	part2.Div(part2, V1)
	part2.Mul(part2, V0)
	part2.Add(part2, new(uint256.Int).Mul(i, delta))

	bAbs := new(uint256.Int).Sub(DecimalMath.ONE, k)
	bAbs.Mul(bAbs, V1)

	var bSig bool

	if bAbs.Cmp(part2) >= 0 {
		bAbs.Sub(bAbs, part2)
		bSig = false
	} else {
		bAbs.Sub(part2, bAbs)
		bSig = true
	}

	bAbs.Div(bAbs, DecimalMath.ONE)

	// calculate sqrt
	// uint256 squareRoot =
	// 	DecimalMath.mulFloor(
	// 		DecimalMath.ONE.sub(k).mul(4),
	// 		DecimalMath.mulFloor(k, V0).mul(V0)
	// 	); // 4(1-k)kQ0^2

	squareRoot := DecimalMath.MulFloor(
		new(uint256.Int).Mul(new(uint256.Int).Sub(DecimalMath.ONE, k), uint256.NewInt(4)),
		new(uint256.Int).Mul(DecimalMath.MulFloor(k, V0), V0))

	//squareRoot = bAbs.mul(bAbs).add(squareRoot).sqrt(); // sqrt(b*b+4(1-k)kQ0*Q0)
	squareRoot.Add(new(uint256.Int).Mul(bAbs, bAbs), squareRoot)
	squareRoot = SafeMath.Sqrt(squareRoot)

	// final res
	//uint256 denominator = DecimalMath.ONE.sub(k).mul(2); // 2(1-k)

	denominator := new(uint256.Int).Sub(DecimalMath.ONE, k)
	denominator.Mul(denominator, uint256.NewInt(2))

	var numerator *uint256.Int

	if bSig {
		//numerator = squareRoot.sub(bAbs);
		numerator = new(uint256.Int).Sub(squareRoot, bAbs)
	} else {
		//numerator = bAbs.add(squareRoot);
		numerator = new(uint256.Int).Add(bAbs, squareRoot)
	}

	V2 := DecimalMath.DivCeil(numerator, denominator)

	if V2.Cmp(V1) > 0 {
		return uint256.NewInt(0)
	} else {
		return new(uint256.Int).Sub(V1, V2)
	}
}
