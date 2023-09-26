package decimalMath

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestDivCeil(t *testing.T) {
	a := uint256.NewInt(100)
	b := uint256.NewInt(10)
	out := DivCeil(a, b)
	if out.Cmp(uint256.NewInt(10000000000000000000)) != 0 {
		t.Fatalf("Got %v expected 10000000000000000000 for DivCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(100)
	b = uint256.NewInt(11)
	out = DivCeil(a, b)
	if out.Cmp(uint256.NewInt(9090909090909090910)) != 0 {
		t.Fatalf("Got %v expected 9090909090909090910 for DivCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(100)
	b = uint256.NewInt(19)
	out = DivCeil(a, b)
	if out.Cmp(uint256.NewInt(5263157894736842106)) != 0 {
		t.Fatalf("Got %v expected 5263157894736842106 for DivCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(51241)
	out = DivCeil(a, b)
	v, _ := uint256.FromHex("0x363324C6D1D3F84A7B") //999809468979918424699
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v expected 999809468979918424699 for DivCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(1e18)
	out = DivCeil(a, b)
	if out.Cmp(uint256.NewInt(51231237)) != 0 {
		t.Fatalf("Got %v expected 51231237 for DivCeil(%v, %v)", out, a, b)
	}
}

func TestDivFloor(t *testing.T) {
	a := uint256.NewInt(100)
	b := uint256.NewInt(10)
	out := DivFloor(a, b)
	if out.Cmp(uint256.NewInt(10000000000000000000)) != 0 {
		t.Fatalf("Got %v expected 10000000000000000000 for DivFloor(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(100)
	b = uint256.NewInt(11)
	out = DivFloor(a, b)
	if out.Cmp(uint256.NewInt(9090909090909090909)) != 0 {
		t.Fatalf("Got %v expected 9090909090909090909 for DivFloor(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(100)
	b = uint256.NewInt(19)
	out = DivFloor(a, b)
	if out.Cmp(uint256.NewInt(5263157894736842105)) != 0 {
		t.Fatalf("Got %v expected 5263157894736842105 for DivFloor(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(51241)
	out = DivFloor(a, b)
	v, _ := uint256.FromHex("0x363324C6D1D3F84A7A") //999809468979918424698
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v expected 999809468979918424698 for DivFloor(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(1e18)
	out = DivFloor(a, b)
	if out.Cmp(uint256.NewInt(51231237)) != 0 {
		t.Fatalf("Got %v expected 51231237 for DivFloor(%v, %v)", out, a, b)
	}
}

func TestMulCeil(t *testing.T) {
	a := uint256.NewInt(100)
	b := uint256.NewInt(10)
	out := MulCeil(a, b)
	if out.Cmp(uint256.NewInt(1)) != 0 {
		t.Fatalf("Got %v expected 1 for MulCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(100)
	b = uint256.NewInt(11)
	out = MulCeil(a, b)
	if out.Cmp(uint256.NewInt(1)) != 0 {
		t.Fatalf("Got %v expected 1 for MulCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(100)
	b = uint256.NewInt(19)
	out = MulCeil(a, b)
	if out.Cmp(uint256.NewInt(1)) != 0 {
		t.Fatalf("Got %v expected 1 for MulCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(51241)
	out = MulCeil(a, b)
	v, _ := uint256.FromHex("0x1") //1
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v expected 1 for MulCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(1e18)
	out = MulCeil(a, b)
	if out.Cmp(uint256.NewInt(51231237)) != 0 {
		t.Fatalf("Got %v expected 51231237 for MulCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(5112321213231237)
	b = uint256.NewInt(7e18)
	out = MulCeil(a, b)
	if out.Cmp(uint256.NewInt(35786248492618659)) != 0 {
		t.Fatalf("Got %v expected 35786248492618659 for MulCeil(%v, %v)", out, a, b)
	}
}

func TestMulFloor(t *testing.T) {
	a := uint256.NewInt(100)
	b := uint256.NewInt(10)
	out := MulFloor(a, b)
	if out.Cmp(uint256.NewInt(0)) != 0 {
		t.Fatalf("Got %v expected 0 for MulFloor(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(100)
	b = uint256.NewInt(11)
	out = MulFloor(a, b)
	if out.Cmp(uint256.NewInt(0)) != 0 {
		t.Fatalf("Got %v expected 0 for MulFloor(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(35786248492618659)
	b = uint256.NewInt(35786248492618659)
	out = MulFloor(a, b)
	if out.Cmp(uint256.NewInt(1280655581175451)) != 0 {
		t.Fatalf("Got %v expected 1280655581175451 for MulFloor(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(51241)
	out = MulFloor(a, b)
	v, _ := uint256.FromHex("0x0") //1
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v expected 0 for MulFloor(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(7e18)
	out = MulFloor(a, b)
	if out.Cmp(uint256.NewInt(358618659)) != 0 {
		t.Fatalf("Got %v expected 358618659 for MulFloor(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(5112321213231237)
	b = uint256.NewInt(7e18)
	out = MulFloor(a, b)
	if out.Cmp(uint256.NewInt(35786248492618659)) != 0 {
		t.Fatalf("Got %v expected 35786248492618659 for MulFloor(%v, %v)", out, a, b)
	}
}

func TestReciprocalCeil(t *testing.T) {
	a := uint256.NewInt(1000)
	out := ReciprocalCeil(a)

	v, _ := uint256.FromHex("0x314DC6448D9338C15B0A00000000") //1000000000000000000000000000000000
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v expected %v for ReciprocalCeil(%v)", out, v, a)
	}

	a = uint256.NewInt(78787878)
	out = ReciprocalCeil(a)
	v, _ = uint256.FromHex("0x2902D4B992BFB6C89FEFD995") // 12692307819230770500000012693
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v expected %v for ReciprocalCeil(%v)", out, v, a)
	}

	a = uint256.NewInt(1e18)
	a.Mul(a, uint256.NewInt(420420696))
	out = ReciprocalCeil(a)
	if out.Cmp(uint256.NewInt(2378569870)) != 0 {
		t.Fatalf("Got %v expected 2378569870 for ReciprocalCeil(%v)", out, a)
	}

}
func TestReciprocalFloor(t *testing.T) {
	a := uint256.NewInt(1000)
	out := ReciprocalFloor(a)

	v, _ := uint256.FromHex("0x314DC6448D9338C15B0A00000000") //1000000000000000000000000000000000
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v expected %v for ReciprocalFloor(%v)", out, v, a)
	}

	a = uint256.NewInt(78787878)
	out = ReciprocalFloor(a)
	v, _ = uint256.FromHex("0x2902D4B992BFB6C89FEFD994") //12692307819230770500000012692
	if out.Cmp(v) != 0 {
		t.Fatalf("Got %v expected %v for ReciprocalFloor(%v)", out, v, a)
	}

	a = uint256.NewInt(1e18)
	a.Mul(a, uint256.NewInt(420420696))
	out = ReciprocalFloor(a)
	if out.Cmp(uint256.NewInt(2378569869)) != 0 {
		t.Fatalf("Got %v expected 2378569869 for ReciprocalFloor(%v)", out, a)
	}

}
