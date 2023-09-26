package balancerV2

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/dex"
)

func TestLn(t *testing.T) {
	return
	output, _ := big.NewInt(0).SetString("42815917693433210574557461317804", 10)
	fmt.Println("Need to get: ", output)
	t0B, _ := new(big.Int).SetString("430640987615090677971821901444426061856309964507521523970387181494137014684", 10)
	b0, _ := uint256.FromBig(t0B)
	fmt.Println(b0)
	t1B, _ := new(big.Int).SetString("430640987615090677971821901444426061856309964507521824334833757182405668275", 10)
	b1, _ := uint256.FromBig(t1B)
	fmt.Println(b1)

	weight0, _ := new(big.Int).SetString("500000000000000000", 10)
	weight1, _ := new(big.Int).SetString("500000000000000000", 10)
	fee, _ := new(big.Int).SetString("900000000000000", 10)
	feeUint, _ := uint256.FromBig(fee)
	w0, _ := uint256.FromBig(weight0)
	w1, _ := uint256.FromBig(weight1)
	sc0 := uint256.NewInt(1e18)
	sc1 := uint256.NewInt(1e18)
	t0 := common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6")
	t1 := common.HexToAddress("0xC728bF49cB1c44C1f0e39A672BB5AB894429bCEA")
	pool := Weighted2TokenPair{token0: dex.NewNotTaxed(t0), token1: dex.NewNotTaxed(t1), b0: b0, b1: b1, w0: w0, w1: w1, fee: feeUint, sc0: sc0, sc1: sc1}
	v := pool.GetTokensOut(t0, t1, uint256.NewInt(1e17))
	fmt.Printf("Got %v \n", v)
}

func TestComplement(t *testing.T) {
	ab, _ := new(big.Int).SetString("11111", 10)
	a, _ := uint256.FromBig(ab)
	out := complement(a)
	vb, _ := new(big.Int).SetString("999999999999988889", 10)
	v, _ := uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for complement(%v)", out, v, a)
	}

	ab, _ = new(big.Int).SetString("58232847523432234", 10)
	a, _ = uint256.FromBig(ab)
	out = complement(a)
	vb, _ = new(big.Int).SetString("941767152476567766", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for complement(%v)", out, v, a)
	}

	ab, _ = new(big.Int).SetString("158232847523432234", 10)
	a, _ = uint256.FromBig(ab)
	out = complement(a)
	vb, _ = new(big.Int).SetString("841767152476567766", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for complement(%v)", out, v, a)
	}

	ab, _ = new(big.Int).SetString("0", 10)
	a, _ = uint256.FromBig(ab)
	out = complement(a)
	vb, _ = new(big.Int).SetString("1000000000000000000", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for complement(%v)", out, v, a)
	}

	ab, _ = new(big.Int).SetString("1000000000000000000", 10)
	a, _ = uint256.FromBig(ab)
	out = complement(a)
	vb, _ = new(big.Int).SetString("0", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for complement(%v)", out, v, a)
	}

	ab, _ = new(big.Int).SetString("154000000000000000000", 10)
	a, _ = uint256.FromBig(ab)
	out = complement(a)
	vb, _ = new(big.Int).SetString("0", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for complement(%v)", out, v, a)
	}
}

func TestPow(t *testing.T) {
	ab, _ := new(big.Int).SetString("10", 10)
	a, _ := uint256.FromBig(ab)
	bb, _ := new(big.Int).SetString("4", 10)
	b, _ := uint256.FromBig(bb)
	out := pow(a, b)
	vb, _ := new(big.Int).SetString("999999999999999844", 10)
	v, _ := uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for pow(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("1041212421124421421412421", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("444122412421414", 10)
	b, _ = uint256.FromBig(bb)
	out = pow(a, b)
	vb, _ = new(big.Int).SetString("1006172687124229394", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for pow(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("1041212421124421421412421", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("2", 10)
	b, _ = uint256.FromBig(bb)
	out = pow(a, b)
	vb, _ = new(big.Int).SetString("1000000000000000027", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for pow(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("1041212421124421421412421", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("19481248141244124", 10)
	b, _ = uint256.FromBig(bb)
	out = pow(a, b)
	vb, _ = new(big.Int).SetString("1309872960287062586", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for pow(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("8", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("41241224241124721", 10)
	b, _ = uint256.FromBig(bb)
	out = pow(a, b)
	vb, _ = new(big.Int).SetString("197197997434394467", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for pow(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("83414141214221412414224214142424", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("1", 10)
	b, _ = uint256.FromBig(bb)
	out = pow(a, b)
	vb, _ = new(big.Int).SetString("1000000000000000032", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for pow(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("834141412142214124142242312312354545214142424", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("0", 10)
	b, _ = uint256.FromBig(bb)
	out = pow(a, b)
	vb, _ = new(big.Int).SetString("1000000000000000000", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for pow(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("0", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("5124124122144214", 10)
	b, _ = uint256.FromBig(bb)
	out = pow(a, b)
	vb, _ = new(big.Int).SetString("0", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for pow(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("834141412142214124142242312312354545214142424", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("0", 10)
	b, _ = uint256.FromBig(bb)
	out = pow(a, b)
	vb, _ = new(big.Int).SetString("1000000000000000000", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for pow(%v,%v)", out, v, a, b)
	}
}

func TestDivUp(t *testing.T) {
	a := uint256.NewInt(1000)
	b := uint256.NewInt(10)
	out := divUp(a, b)
	vb, _ := new(big.Int).SetString("100000000000000000000", 10)
	v, _ := uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(13123123000)
	b = uint256.NewInt(1412)
	out = divUp(a, b)
	vb, _ = new(big.Int).SetString("9293996458923512747875355", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(1004901883832123210)
	b = uint256.NewInt(4)
	out = divUp(a, b)
	vb, _ = new(big.Int).SetString("251225470958030802500000000000000000", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(4123213421)
	b = uint256.NewInt(14123123132131232130)
	out = divUp(a, b)
	vb, _ = new(big.Int).SetString("291947708", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(10999999999999999990)
	a.Mul(a, uint256.NewInt(3131223131313213211))
	b = uint256.NewInt(5918739239)
	out = divUp(a, b)
	vb, _ = new(big.Int).SetString("5819390423130844958937338426452658913970445320", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divUp(%v,%v)", out, v, a, b)
	}
}

func TestMulDown(t *testing.T) {
	a := uint256.NewInt(10999999999999999990)
	b := uint256.NewInt(5918739239)
	out := mulDown(a, b)
	vb, _ := new(big.Int).SetString("65106131628", 10)
	v, _ := uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for mulDown(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(1131234)
	b = uint256.NewInt(50189219321908)
	out = mulDown(a, b)
	vb, _ = new(big.Int).SetString("56", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for mulDown(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(198898098123121234)
	a.Mul(a, uint256.NewInt(1239123213))
	b = uint256.NewInt(50189219321999908)
	b.Mul(b, uint256.NewInt(13123123239123213))
	out = mulDown(a, b)
	vb, _ = new(big.Int).SetString("162327750738412994267909346356098079089087", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for mulDown(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(1123123131234)
	b = uint256.NewInt(0)
	out = mulDown(a, b)
	vb, _ = new(big.Int).SetString("0", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for mulDown(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(0)
	b = uint256.NewInt(123213213)
	out = mulDown(a, b)
	vb, _ = new(big.Int).SetString("0", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for mulDown(%v,%v)", out, v, a, b)
	}
}

func TestPowUp(t *testing.T) {
	a := uint256.NewInt(10999999999999999990)
	b := uint256.NewInt(5918739239)
	out := powUp(a, b)
	vb, _ := new(big.Int).SetString("1000000014192526944", 10)
	v, _ := uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(1123123131234)
	b = uint256.NewInt(0)
	out = powUp(a, b)
	vb, _ = new(big.Int).SetString("1000000000000010001", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(0)
	b = uint256.NewInt(123213213)
	out = powUp(a, b)
	vb, _ = new(big.Int).SetString("1", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(3131321413123143)
	a.Mul(a, uint256.NewInt(321321314123))
	b = uint256.NewInt(1)
	out = powUp(a, b)
	vb, _ = new(big.Int).SetString("1000000000000010022", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(3131321413123143)
	a.Mul(a, uint256.NewInt(321321314123))
	b = uint256.NewInt(2)
	out = powUp(a, b)
	vb, _ = new(big.Int).SetString("1000000000000010043", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(3131321413123143)
	a.Mul(a, uint256.NewInt(321321314123))
	b = uint256.NewInt(4)
	out = powUp(a, b)
	vb, _ = new(big.Int).SetString("1000000000000010084", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(9884192999991937999)
	a.Mul(a, uint256.NewInt(7e18))
	b = uint256.NewInt(1)
	out = powUp(a, b)
	vb, _ = new(big.Int).SetString("1000000000000010047", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(9884192731937999)
	a.Mul(a, uint256.NewInt(1e18))
	b = uint256.NewInt(2)
	out = powUp(a, b)
	vb, _ = new(big.Int).SetString("1000000000000010075", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(9884192731937999)
	a.Mul(a, uint256.NewInt(1e18))
	b = uint256.NewInt(5)
	out = powUp(a, b)
	vb, _ = new(big.Int).SetString("1000000000000010186", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}

	a = uint256.NewInt(95232343431937999)
	a.Mul(a, uint256.NewInt(7e18))
	b = uint256.NewInt(541241424)
	out = powUp(a, b)
	vb, _ = new(big.Int).SetString("1000000022213102869", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for powUp(%v,%v)", out, v, a, b)
	}
}

func TestMulUp(t *testing.T) {
	a := uint256.NewInt(1)
	a.Mul(a, uint256.NewInt(1))
	b := uint256.NewInt(1)
	out := mulUp(a, b)
	vb, _ := new(big.Int).SetString("1", 10)
	v, _ := uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for MulUp(%v,%v)", out, v, a, b)
	}

	ab, _ := new(big.Int).SetString("112313213131331321313123", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ := new(big.Int).SetString("112123123123", 10)
	b, _ = uint256.FromBig(bb)
	out = mulUp(a, b)
	vb, _ = new(big.Int).SetString("12592908224264003", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for MulUp(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("112313213131331321313123", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("1121231231235345435435", 10)
	b, _ = uint256.FromBig(bb)
	out = mulUp(a, b)
	vb, _ = new(big.Int).SetString("125929082243240384114266568", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for MulUp(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("112313213131331321313123", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("0", 10)
	b, _ = uint256.FromBig(bb)
	out = mulUp(a, b)
	vb, _ = new(big.Int).SetString("0", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for MulUp(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("0", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("1121231231235345435435", 10)
	b, _ = uint256.FromBig(bb)
	out = mulUp(a, b)
	vb, _ = new(big.Int).SetString("0", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for MulUp(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("3123", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("52342342423424535435", 10)
	b, _ = uint256.FromBig(bb)
	out = mulUp(a, b)
	vb, _ = new(big.Int).SetString("163466", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for MulUp(%v,%v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("1482941849018490124890248140", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("1231231", 10)
	b, _ = uint256.FromBig(bb)
	out = mulUp(a, b)
	vb, _ = new(big.Int).SetString("1825843975708885", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for MulUp(%v,%v)", out, v, a, b)
	}
}

func Test_Ln(t *testing.T) {
	a, _ := new(big.Int).SetString("12313213131", 10)
	out := _ln(a)
	v, _ := new(big.Int).SetString("-18212592912856424453", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("1314812981904819", 10)
	out = _ln(a)
	v, _ = new(big.Int).SetString("-6634060842550289863", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("1", 10)
	out = _ln(a)
	v, _ = new(big.Int).SetString("-41446531673892822312", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("29875894759823743237444", 10)
	out = _ln(a)
	v, _ = new(big.Int).SetString("10304807238904939117", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("4215894759823743237444444", 10)
	out = _ln(a)
	v, _ = new(big.Int).SetString("15254372406800634628", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("421589475412414242149823743237444444", 10)
	out = _ln(a)
	v, _ = new(big.Int).SetString("40582808428383205678", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("100000000", 10)
	out = _ln(a)
	v, _ = new(big.Int).SetString("-23025850929940456840", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("9999999999999999999999", 10)
	out = _ln(a)
	v, _ = new(big.Int).SetString("9210340371976182736", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("122414141242141241249999999999999999999999", 10)
	out = _ln(a)
	v, _ = new(big.Int).SetString("53161696849303788995", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("59827574857298572357834574545735739576485735743857649735765873598", 10)
	out = _ln(a)
	v, _ = new(big.Int).SetString("107707795857421427417", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln(%v)", out, v, a)
	}
}

func Test_ln_36(t *testing.T) {
	a, _ := new(big.Int).SetString("1000000000000000000", 10)
	out := _ln_36(a)
	v, _ := new(big.Int).SetString("0", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln_36(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("1100000000000000000", 10)
	out = _ln_36(a)
	v, _ = new(big.Int).SetString("95310179804324860043948199225536944", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln_36(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("1014144141244878870", 10)
	out = _ln_36(a)
	v, _ = new(big.Int).SetString("14045046195548914443781922870226942", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln_36(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("914144141244878870", 10)
	out = _ln_36(a)
	v, _ = new(big.Int).SetString("-89767016192905108307170668794567936", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for _ln_36(%v)", out, v, a)
	}
}

func TestExp(t *testing.T) {
	a, _ := new(big.Int).SetString("100", 10)
	out := exp(a)
	v, _ := new(big.Int).SetString("1000000000000000100", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("914144141244878870", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("2494639279121861344", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("1000000000000000000", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("2718281828459045235", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("7000000000000000000", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("1096633158428458599263", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("70000000000000000000", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("2515438670919167006265420014324008882556494800000", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("74124508912389123897", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("155548273283893860201735563143998479083015165800000", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("14124508912389123897", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("1362060129129276204080834", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("129000000000000000000", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("105678871143625881256404495171560308048992000000000000000000000000000000000", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("-40000000000000000000", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("4", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("-4000000000000000000", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("18315638888734180", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("-2123516768934899021", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("119610246395163060", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("-216768934899021", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("999783254557789022", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("-68923", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("999999999999931077", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("0", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("1000000000000000000", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("1", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("1000000000000000001", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}

	a, _ = new(big.Int).SetString("-1", 10)
	out = exp(a)
	v, _ = new(big.Int).SetString("999999999999999999", 10)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for exp(%v)", out, v, a)
	}
}

func TestDivDown(t *testing.T) {
	ab, _ := new(big.Int).SetString("11111", 10)
	a, _ := uint256.FromBig(ab)
	bb, _ := new(big.Int).SetString("11111", 10)
	b, _ := uint256.FromBig(bb)
	out := divDown(a, b)
	vb, _ := new(big.Int).SetString("1000000000000000000", 10)
	v, _ := uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divDown(%v, %v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("1111152342342341", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("11111", 10)
	b, _ = uint256.FromBig(bb)
	out = divDown(a, b)
	vb, _ = new(big.Int).SetString("100004710857919269192691926919", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divDown(%v, %v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("1111152342342341", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("543523523414124141241", 10)
	b, _ = uint256.FromBig(bb)
	out = divDown(a, b)
	vb, _ = new(big.Int).SetString("2044350050137", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divDown(%v, %v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("0", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("543523523414124141241", 10)
	b, _ = uint256.FromBig(bb)
	out = divDown(a, b)
	vb, _ = new(big.Int).SetString("0", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divDown(%v, %v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("4124124124141414124214124141241414214", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("543523523414124141241", 10)
	b, _ = uint256.FromBig(bb)
	out = divDown(a, b)
	vb, _ = new(big.Int).SetString("7587756456677112386259797543043838", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divDown(%v, %v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("15293482347247842472147238472472424724724724718479124", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("4", 10)
	b, _ = uint256.FromBig(bb)
	out = divDown(a, b)
	vb, _ = new(big.Int).SetString("3823370586811960618036809618118106181181181179619781000000000000000000", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divDown(%v, %v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("14241284719431", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("1", 10)
	b, _ = uint256.FromBig(bb)
	out = divDown(a, b)
	vb, _ = new(big.Int).SetString("14241284719431000000000000000000", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divDown(%v, %v)", out, v, a, b)
	}

	ab, _ = new(big.Int).SetString("14241284719431", 10)
	a, _ = uint256.FromBig(ab)
	bb, _ = new(big.Int).SetString("333333333333333333333333", 10)
	b, _ = uint256.FromBig(bb)
	out = divDown(a, b)
	vb, _ = new(big.Int).SetString("42723854", 10)
	v, _ = uint256.FromBig(vb)
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v exptected %v for divDown(%v, %v)", out, v, a, b)
	}

}
