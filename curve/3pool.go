package curve

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/curve/reg1"
)

type Pair struct {
	address  common.Address
	tokens   []common.Address
	reserves []*uint256.Int
	protocol string

	A                 *uint256.Int
	rates             []*uint256.Int
	LENDING_PRECISION *uint256.Int
	PRECISION         *uint256.Int
	FEE_DENOMINATOR   *uint256.Int
	fee               *uint256.Int
	admin_fee         *uint256.Int
	N_COINS           *uint256.Int
}

func newPair(client bind.ContractBackend, pairAddr common.Address, tokens []common.Address, A *uint256.Int, rates []*uint256.Int, fee *uint256.Int, admin_fee *uint256.Int, LENDING_PRECISION *uint256.Int, PRECISION *uint256.Int, FEE_DENOMINATOR *uint256.Int, N_COINS *uint256.Int, protocol string) Pair {
	pair := Pair{address: pairAddr, tokens: tokens, A: A, rates: rates, fee: fee, admin_fee: admin_fee, LENDING_PRECISION: LENDING_PRECISION, PRECISION: PRECISION, FEE_DENOMINATOR: FEE_DENOMINATOR, N_COINS: N_COINS, protocol: protocol}
	pair.reserves = make([]*uint256.Int, 0)

	return pair
}

func GetCurveMarkets(client bind.ContractBackend, CURVE_ADDRESS_PROVIDER common.Address, CURVE_REGISTRY common.Address, protocol string, BLACKLIST map[common.Address]bool) []Pair {
	fmt.Printf("=============== Setting up %v ===================\n", protocol)
	now := time.Now()

	marketPairs := make([]Pair, 0)
	registry, _ := reg1.NewReg1Caller(CURVE_REGISTRY, client)

	pools_count, _ := registry.PoolCount(nil)

	var num = pools_count.Int64()

	var j int64 = 0
	for j < num {
		pool, _ := registry.PoolList(nil, big.NewInt(j))
		tokens, _ := registry.GetCoins(nil, pool)

		tokensNum := 0
		for _, i := range tokens {
			if i == common.HexToAddress("") {
				break
			}
			tokensNum++
		}

		factor3 := uint256.NewInt(10000000000)
		factor2, _ := uint256.FromHex("0xC9F2C9CD04674EDEA40000000")
		factor := uint256.NewInt(1000000000000000000)
		raw_A, _ := registry.GetA(nil, pool)
		A, _ := uint256.FromBig(raw_A)
		rates := []*uint256.Int{factor, factor2, factor2}
		fee := uint256.NewInt(1000000)
		admin_fee := uint256.NewInt(5000000000)
		LENDING_PRECISION := factor
		PRECISION := factor
		FEE_DENOMINATOR := factor3
		N_COINS := uint256.NewInt(3)

		final_tokens := make([]common.Address, 0)
		for i := 0; i < tokensNum; i++ {

			// Only grabs initialized addresses
			if tokens[i] != common.HexToAddress("0x0000000000000000000000000000000000000000") {
				final_tokens = append(final_tokens, tokens[i])
			}
		}

		pair := newPair(client, pool, final_tokens, A, rates, fee, admin_fee, LENDING_PRECISION, PRECISION, FEE_DENOMINATOR, N_COINS, protocol)
		marketPairs = append(marketPairs, pair)
		j += 1
	}

	UpdateReserves(client, CURVE_REGISTRY, marketPairs, protocol)
	fmt.Printf("Loaded %v pools from %v\n", len(marketPairs), protocol)
	fmt.Printf("Took: %v\n\n", time.Since(now))
	return marketPairs
}

func UpdateReserves(client bind.ContractBackend, CURVE_REGISTRY common.Address, allMarketPairs []Pair, protocol string) []Pair {
	now := time.Now()

	registry, _ := reg1.NewReg1Caller(CURVE_REGISTRY, client)

	for i := 0; i < len(allMarketPairs); i++ {
		pool := allMarketPairs[i]

		balances, _ := registry.GetBalances(nil, pool.address)
		final_balances := make([]*uint256.Int, 0)
		for i := 0; i < int(pool.N_COINS.Uint64()); i++ {
			vUint, err := uint256.FromBig(balances[i])
			if err {
				panic("overflow")
			}
			final_balances = append(final_balances, vUint)
		}
		allMarketPairs[i].reserves = final_balances
	}

	fmt.Printf("=============== Updating Reserves for %v ===================\n", protocol)
	fmt.Printf("Updated %v pools from Curve \n", len(allMarketPairs))
	fmt.Printf("Took: %v\n\n", time.Since(now))
	return allMarketPairs
}

func (pair Pair) Get_Dy(i, j uint64, dx *uint256.Int) *uint256.Int {
	rates := pair.rates
	xp := pair.Xp(pair.rates, pair.reserves, pair.LENDING_PRECISION)

	t2 := new(uint256.Int).Mul(dx, rates[i]) //(dx*rates[i]/PRECISION)
	t2.Div(t2, pair.PRECISION)
	x := new(uint256.Int).Add(xp[i], t2)

	y := pair.Get_Y(i, j, x, xp)
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

func (pair Pair) Xp(rates []*uint256.Int, balances []*uint256.Int, LENDING_PRECISION *uint256.Int) []*uint256.Int {
	result := make([]*uint256.Int, len(rates))
	for i := 0; i < len(balances); i++ {
		result[i] = new(uint256.Int).Mul(rates[i], balances[i])
		result[i].Div(result[i], LENDING_PRECISION)
	}
	return result
}

func (pair Pair) Get_Y(i, j uint64, x *uint256.Int, xp []*uint256.Int) *uint256.Int {
	amp := pair.A
	D := pair.Get_D(xp, amp)
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

func (pair Pair) Get_D(xp []*uint256.Int, amp *uint256.Int) *uint256.Int {
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

// This could be expanded to include the is_underlying
func (pair Pair) ConvertAddresses(tokenIn common.Address, tokenOut common.Address) (uint64, uint64) {
	var i, j uint64
	for k, t := range pair.Tokens() {
		if t == tokenIn {
			i = uint64(k)
		}
		if t == tokenOut {
			j = uint64(k)
		}
	}
	return i, j
}

func (pair Pair) tokensToReserves(tokenIn common.Address, tokenOut common.Address) (*uint256.Int, *uint256.Int) {
	panic("Not implemented")
}

func (pair Pair) GetTokensOut(tokenIn common.Address, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int {
	i, j := pair.ConvertAddresses(tokenIn, tokenOut)
	return pair.Get_Dy(i, j, amountIn)
}

func (pair Pair) getAmountOut(reserveIn *uint256.Int, reserveOut *uint256.Int, amountIn *uint256.Int) *uint256.Int {
	panic("Not implemented")
}

func (pair Pair) Fee() *uint256.Int {
	return pair.fee
}

func SellTokensToNextMarket(tokenIn common.Address, amountIn *uint256.Int, ethMarketAddress common.Address) []byte {
	panic("Not implemented")
}

func (pair Pair) SellTokens(tokenIn common.Address, amountIn *uint256.Int, recipient, sender common.Address) ([]common.Address, [][]byte) {
	panic("Not implemented")
}

func (pair Pair) ReceiveDirectly() bool {
	return true
}

func (pair Pair) PrepareReceive(tokenAddress common.Address, amountIn *uint256.Int) ([]common.Address, [][]byte) {
	panic("Not implemented")
}

func (pair Pair) GetOtherTokenAddress(knowAddress common.Address) common.Address {
	panic("Not implemented")
}

func (pair Pair) MarketAddress() common.Address {
	return pair.address
}

func (pair Pair) SpenderAddress() common.Address {
	panic("Not implemented")
}

func (pair Pair) Tokens() []common.Address {
	return pair.tokens
}

func (pair Pair) Protocol() string {
	return pair.protocol
}

func (pair Pair) Reserves() []*uint256.Int {
	return pair.reserves
}

func (pair Pair) Rates() []*uint256.Int {
	return pair.rates
}

func (pair *Pair) FakeReserves(fake *uint256.Int) {
	panic("Not implemented")
}
