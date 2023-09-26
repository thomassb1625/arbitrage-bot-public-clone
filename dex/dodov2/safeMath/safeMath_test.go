package safeMath

import (
	"testing"

	"github.com/holiman/uint256"
)

func TestSqrt(t *testing.T) {
	a := uint256.NewInt(1000)
	out := Sqrt(a)

	if out.Cmp(uint256.NewInt(31)) != 0 {
		t.Fatalf("Got %v expected 31 for Sqrt(%v)", out, a)
	}

	a = uint256.NewInt(78787878)
	out = Sqrt(a)
	if out.Cmp(uint256.NewInt(8876)) != 0 {
		t.Fatalf("Got %v expected 8876 for Sqrt(%v)", out, a)
	}

	a = uint256.NewInt(1e18)
	out = Sqrt(a)
	if out.Cmp(uint256.NewInt(1000000000)) != 0 {
		t.Fatalf("Got %v expected 1000000000 for Sqrt(%v)", out, a)
	}

	a = uint256.NewInt(1e18)
	a.Mul(a, uint256.NewInt(420420696))
	out = Sqrt(a)
	if out.Cmp(uint256.NewInt(20504162894397)) != 0 {
		t.Fatalf("Got %v expected 20504162894397 for Sqrt(%v)", out, a)
	}

	a = uint256.NewInt(3)
	out = Sqrt(a)
	if out.Cmp(uint256.NewInt(1)) != 0 {
		t.Fatalf("Got %v expected 1 for Sqrt(%v)", out, a)
	}

}

func TestDivCeil(t *testing.T) {
	a := uint256.NewInt(100)
	b := uint256.NewInt(10)
	out := DivCeil(a, b)
	if out.Cmp(uint256.NewInt(10)) != 0 {
		t.Fatalf("Got %v expected 10 for DivCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(100)
	b = uint256.NewInt(11)
	out = DivCeil(a, b)
	if out.Cmp(uint256.NewInt(10)) != 0 {
		t.Fatalf("Got %v expected 10 for DivCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(100)
	b = uint256.NewInt(19)
	out = DivCeil(a, b)
	if out.Cmp(uint256.NewInt(6)) != 0 {
		t.Fatalf("Got %v expected 6 for DivCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(51241)
	out = DivCeil(a, b)
	if out.Cmp(uint256.NewInt(1000)) != 0 {
		t.Fatalf("Got %v expected 1000 for DivCeil(%v, %v)", out, a, b)
	}

	a = uint256.NewInt(51231237)
	b = uint256.NewInt(1e18)
	out = DivCeil(a, b)
	if out.Cmp(uint256.NewInt(1)) != 0 {
		t.Fatalf("Got %v expected 1000 for DivCeil(%v, %v)", out, a, b)
	}

}
