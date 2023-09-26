package uniswapV3

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

// Errors for pools
var (
	ErrPoolState  = errors.New("PoolState - SqrtPriceX96 | Liquidity not set")
	ErrTickList   = errors.New("TickList is empty")
	ErrPoolCreate = errors.New("TickSpacing | Fee nil or 0")
)

// Tick is a single Tick from the pool
type Tick struct {
	Index          int
	LiquidityGross uint256.Int // uint128
	LiquidityNet   big.Int     // int128
	Initialized    bool
}

// NewPool creates a new UniswapV3 pair that holds all the needed information
//
// - a good hackyRounds default is 10, you can set a higher number or 0 for more iterations to get probably to a later tick, but in exchange for speed
func NewPool(tickSpacing int, fee *uint256.Int, hackyRounds int) (*Pair, error) {
	if fee == nil || fee.IsZero() {
		return nil, ErrPoolCreate
	}

	return &Pair{
		TickSpacing: tickSpacing,
		Fee:         *fee,
		HackyRounds: hackyRounds,
		//mutex:       sync.RWMutex{},
	}, nil
}

// SetTickList sets the tickList for the pool
//func (p *Pair) SetTickList(tickList []Tick) {
//	p.mutex.Lock()
//	p.tickList = tickList
//	p.mutex.Unlock()
//}
//
//// GetTickList get the current tickList from the pool
//func (p *Pair) GetTickList() (tickList []Tick) {
//	p.mutex.RLock()
//	defer p.mutex.RUnlock()
//
//	return p.tickList
//}

// SetState sets the current state for the pool
func (p *Pair) SetState(tick int, sqrtPriceX96, liquidity *uint256.Int) {
	//p.mutex.Lock()
	p.Tick = tick
	p.SqrtPriceX96 = sqrtPriceX96
	p.Liquidity = liquidity
	//p.mutex.Unlock()
}

// GetState get the current state from the pool
func (p *Pair) GetState() (tick int, sqrtPriceX96, liquidity *uint256.Int) {
	//	p.mutex.RLock()
	//	defer p.mutex.RUnlock()

	return p.Tick, p.SqrtPriceX96, p.Liquidity
}

// SetHackyRounds is used to stop the swap loop if we didn't find liquidity in the next x ticks
func (p *Pair) SetHackyRounds(hackyRounds int) {
	//p.mutex.Lock()
	p.HackyRounds = hackyRounds
	//p.mutex.Unlock()
}

// GetHackyRounds get the current max rounds the pool is using in the calculation
func (p *Pair) GetHackyRounds() (hackyRounds int) {
	//p.mutex.RLock()
	//defer p.mutex.RUnlock()

	return p.HackyRounds
}

// swapState the top level state of the swap, the results will not update the state
type swapState struct {
	amountSpecifiedRemaining     uint256.Int
	amountSpecifiedRemainingLast uint256.Int
	amountCalculated             uint256.Int
	sqrtPriceX96                 uint256.Int
	tick                         int
	liquidity                    uint256.Int
}

// stepComputations swap step states
type stepComputations struct {
	sqrtPriceStartX96 uint256.Int
	tickNext          int
	initialized       bool
	sqrtPriceNextX96  uint256.Int
	amountIn          uint256.Int
	amountOut         uint256.Int
	feeAmount         uint256.Int
}

// GetAmountIn calculates the amount you need to pay to get the amount out
func (p *Pair) GetAmountIn(amountOut *uint256.Int, zeroForOne bool) (*uint256.Int, error) {
	//	p.mutex.RLock()
	//	defer p.mutex.RUnlock()

	return p.swap(amountOut, zeroForOne)
}

func (p *Pair) GetTokensOut(tokenIn, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int {
	if amountIn.IsZero() || p.Liquidity.IsZero() {
		return uint256.NewInt(0)
	}

	var v *uint256.Int
	if tokenIn == p.Token0.Addr {
		v, _ = p.swap(amountIn, true)

	} else {
		v, _ = p.swap(amountIn, false)
	}
	return v
}

// GetAmountOut calculates the amount you will receive for the amount you give in
func (p *Pair) GetAmountOut(amountIn *uint256.Int, zeroForOne bool) *uint256.Int {
	//	p.mutex.RLock()
	//	defer p.mutex.RUnlock()
	if amountIn.IsZero() || p.Liquidity.IsZero() {
		return uint256.NewInt(0)
	}

	v, _ := p.swap(amountIn, zeroForOne)
	return v
}

//WARNING: Only works with exactInput, meaming amountSpecified can't be 0
func (p *Pair) swap(amountSpecified *uint256.Int, zeroForOne bool) (*uint256.Int, error) {
	if p.SqrtPriceX96 == nil || p.Liquidity == nil {
		return nil, ErrPoolState
	}

	if len(p.TickList) == 0 {
		return nil, ErrTickList
	}

	// as we default to min or max we don't need the check regarding sqrtPriceLimitX96 + slot0Start.sqrtPriceX96
	sqrtPriceLimitX96 := new(uint256.Int).Set(MaxSqrtRatioSubOne)
	if zeroForOne {
		sqrtPriceLimitX96 = new(uint256.Int).Set(MinSqrtRatioAddOne)
	}

	exactInput := !amountSpecified.IsZero()

	// keep track of swap state
	state := swapState{
		amountSpecifiedRemaining:     *amountSpecified,
		amountSpecifiedRemainingLast: *ZeroUint256,
		amountCalculated:             *ZeroUint256,
		sqrtPriceX96:                 *p.SqrtPriceX96,
		tick:                         p.Tick,
		liquidity:                    *p.Liquidity,
	}

	hackyCounter := 0
	// swap loop
	for state.amountSpecifiedRemaining.Cmp(ZeroUint256) != 0 && state.sqrtPriceX96.Cmp(sqrtPriceLimitX96) != 0 {

		// hacky counter handle
		if p.HackyRounds > 0 {
			if state.amountSpecifiedRemaining.Cmp(&state.amountSpecifiedRemainingLast) == 0 {
				hackyCounter++
			} else {
				hackyCounter = 0
			}
			state.amountSpecifiedRemainingLast = state.amountSpecifiedRemaining
			if hackyCounter == p.HackyRounds {
				break
			}
		}

		var step stepComputations
		//		fmt.Println("STATE")
		//		fmt.Printf("amountSpecifiedRemaining: %v\n", state.amountSpecifiedRemaining.ToBig().String())
		//		fmt.Printf("amountCaclulated: %v\n", state.amountCalculated.ToBig().String())
		//		fmt.Printf("sqrtPX96: %v\n", state.sqrtPriceX96.ToBig().String())
		//		fmt.Printf("tick: %v\n", state.tick)
		//		fmt.Printf("liquidity: %v\n", state.liquidity.ToBig().String())
		//		fmt.Printf("fee: %v\n", p.Fee.String())

		step.sqrtPriceStartX96 = state.sqrtPriceX96
		step.tickNext, step.initialized = p.nextInitializedTickWithinOneWord(
			state.tick,
			zeroForOne,
		)

		if step.tickNext < MinTick {
			step.tickNext = MinTick
		} else if step.tickNext > MaxTick {
			step.tickNext = MaxTick
		}

		step.sqrtPriceNextX96 = getSqrtRatioAtTick(step.tickNext)
		//		fmt.Println("STEP")
		//		fmt.Printf("tickNext: %v\n", step.tickNext)
		//		fmt.Printf("sqrtPX96Start: %v\n", step.sqrtPriceStartX96.ToBig().String())
		//		fmt.Printf("sqrtPX96Next: %v\n", step.sqrtPriceNextX96.ToBig().String())

		var sqrtRatioTargetX96 *uint256.Int
		if zeroForOne {
			if step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) < 0 {
				sqrtRatioTargetX96 = sqrtPriceLimitX96
			} else {
				sqrtRatioTargetX96 = &step.sqrtPriceNextX96
			}
		} else {
			if step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) > 0 {
				sqrtRatioTargetX96 = sqrtPriceLimitX96
			} else {
				sqrtRatioTargetX96 = &step.sqrtPriceNextX96
			}
		}

		//		fmt.Println("sqrtPriceX96Target: ", sqrtRatioTargetX96.ToBig().String())
		// compute values to swap to the target tick, price limit, or point where input/output amount is exhausted
		var err error
		state.sqrtPriceX96, step.amountIn, step.feeAmount, step.amountOut, err = computeSwapStep(
			state.sqrtPriceX96,
			*sqrtRatioTargetX96,
			state.liquidity,
			state.amountSpecifiedRemaining,
			p.Fee,
			exactInput,
		)
		if err != nil {
			return nil, err
		}

		//		fmt.Println("After computeSwapStep state.SqrtPriceX96: ", state.sqrtPriceX96.ToBig().String())
		//		fmt.Println("After computeSwapStep step.amountIn: ", step.amountIn.ToBig().String())
		//		fmt.Println("After computeSwapStep step.feeAmount: ", step.feeAmount.ToBig().String())
		//		fmt.Println("After computeSwapStep step.amountOut: ", step.amountOut.ToBig().String())

		if exactInput {
			_amountSpecifiedRemaining, isOverflow := new(uint256.Int).SubOverflow(&state.amountSpecifiedRemaining, new(uint256.Int).Add(&step.amountIn, &step.feeAmount))
			state.amountSpecifiedRemaining = *_amountSpecifiedRemaining
			if isOverflow {
				state.amountSpecifiedRemaining = *ZeroUint256
			}
			state.amountCalculated = *new(uint256.Int).Add(&state.amountCalculated, &step.amountOut)
		} else {
			_amountSpecifiedRemaining, isOverflow := new(uint256.Int).SubOverflow(&state.amountSpecifiedRemaining, &step.amountOut)
			state.amountSpecifiedRemaining = *_amountSpecifiedRemaining
			if isOverflow {
				state.amountSpecifiedRemaining = *ZeroUint256
			}
			state.amountCalculated = *new(uint256.Int).Add(&state.amountCalculated, new(uint256.Int).Add(&step.amountIn, &step.feeAmount))
		}
		//		fmt.Println("STATE amountSpecifiedRemaining after: ", state.amountSpecifiedRemaining.ToBig().String())
		//		fmt.Println("STATE amountCalculated after: ", state.amountCalculated.ToBig().String())

		if state.sqrtPriceX96.Cmp(&step.sqrtPriceNextX96) == 0 {
			// if the tick is initialized, run the tick transition
			if step.initialized {
				liquidityNetTick := p.TickList[p.binarySearch(step.tickNext)]

				liquidityNet := liquidityNetTick.LiquidityNet
				if zeroForOne {
					liquidityNet = *new(big.Int).Mul(&liquidityNet, NegativeOneBig)
				}
				state.liquidity.Set(addDelta(&state.liquidity, &liquidityNet))

				//if liquidityNet.Cmp(ZeroBig) < 0 {
				//	liquidityBig := state.liquidity.ToBig()
				//	_liquidityCalc, _ := uint256.FromBig(new(big.Int).Sub(liquidityBig, new(big.Int).Mul(liquidityNet, NegativeOneBig)))
				//	state.liquidity = *_liquidityCalc
				//} else {
				//	liquidityNetBig, _ := uint256.FromBig(liquidityNet)
				//	_liquidityCalc := new(uint256.Int).Add(&state.liquidity, liquidityNetBig)
				//	state.liquidity = *_liquidityCalc
				//}
			}
			if zeroForOne {
				state.tick = step.tickNext - 1
			} else {
				state.tick = step.tickNext
			}
		} else if state.sqrtPriceX96.Cmp(&step.sqrtPriceStartX96) != 0 {
			state.tick, err = getTickAtSqrtRatio(&state.sqrtPriceX96)
			if err != nil {
				return nil, err
			}
		}
	}

	return &state.amountCalculated, nil
}

func (p *Pair) swapOld(amountSpecified *uint256.Int, zeroForOne bool, exactInput bool) (*uint256.Int, error) {
	if p.SqrtPriceX96 == nil || p.Liquidity == nil {
		return nil, ErrPoolState
	}

	//	if len(p.ticks) == 0 {
	//		return nil, ErrTickList
	//	}

	// as we default to min or max we don't need the check regarding sqrtPriceLimitX96 + slot0Start.sqrtPriceX96
	sqrtPriceLimitX96 := new(uint256.Int).Set(MaxSqrtRatioSubOne)
	if zeroForOne {
		sqrtPriceLimitX96.Set(MinSqrtRatioAddOne)
	}

	// keep track of swap state
	state := swapState{
		amountSpecifiedRemaining:     *amountSpecified,
		amountSpecifiedRemainingLast: *ZeroUint256,
		amountCalculated:             *ZeroUint256,
		sqrtPriceX96:                 *p.SqrtPriceX96,
		tick:                         p.Tick,
		liquidity:                    *p.Liquidity,
	}

	hackyCounter := 0
	// swap loop
	for state.amountSpecifiedRemaining.Cmp(ZeroUint256) != 0 && state.sqrtPriceX96.Cmp(sqrtPriceLimitX96) != 0 {

		// hacky counter handle
		if p.HackyRounds > 0 {
			if state.amountSpecifiedRemaining.Cmp(&state.amountSpecifiedRemainingLast) == 0 {
				hackyCounter++
			} else {
				hackyCounter = 0
			}
			state.amountSpecifiedRemainingLast = state.amountSpecifiedRemaining
			if hackyCounter == p.HackyRounds {
				break
			}
		}

		var step stepComputations
		step.sqrtPriceStartX96 = state.sqrtPriceX96
		step.tickNext, step.initialized = p.nextInitializedTickWithinOneWord(
			state.tick,
			zeroForOne,
		)

		if step.tickNext < MinTick {
			step.tickNext = MinTick
		} else if step.tickNext > MaxTick {
			step.tickNext = MaxTick
		}

		step.sqrtPriceNextX96 = getSqrtRatioAtTick(step.tickNext)

		var sqrtRatioTargetX96 *uint256.Int
		if zeroForOne {
			if step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) < 0 {
				sqrtRatioTargetX96 = sqrtPriceLimitX96
			} else {
				sqrtRatioTargetX96 = &step.sqrtPriceNextX96
			}
		} else {
			if step.sqrtPriceNextX96.Cmp(sqrtPriceLimitX96) > 0 {
				sqrtRatioTargetX96 = sqrtPriceLimitX96
			} else {
				sqrtRatioTargetX96 = &step.sqrtPriceNextX96
			}
		}

		// compute values to swap to the target tick, price limit, or point where input/output amount is exhausted
		var err error
		state.sqrtPriceX96, step.amountIn, step.feeAmount, step.amountOut, err = computeSwapStep(
			state.sqrtPriceX96,
			*sqrtRatioTargetX96,
			state.liquidity,
			state.amountSpecifiedRemaining,
			p.Fee,
			exactInput,
		)
		if err != nil {
			return nil, err
		}

		if exactInput {
			_amountSpecifiedRemaining, isOverflow := new(uint256.Int).SubOverflow(&state.amountSpecifiedRemaining, new(uint256.Int).Add(&step.amountIn, &step.feeAmount))
			state.amountSpecifiedRemaining = *_amountSpecifiedRemaining
			if isOverflow {
				state.amountSpecifiedRemaining = *ZeroUint256
			}
			state.amountCalculated = *new(uint256.Int).Add(&state.amountCalculated, &step.amountOut)
		} else {
			_amountSpecifiedRemaining, isOverflow := new(uint256.Int).SubOverflow(&state.amountSpecifiedRemaining, &step.amountOut)
			state.amountSpecifiedRemaining = *_amountSpecifiedRemaining
			if isOverflow {
				state.amountSpecifiedRemaining = *ZeroUint256
			}
			state.amountCalculated = *new(uint256.Int).Add(&state.amountCalculated, new(uint256.Int).Add(&step.amountIn, &step.feeAmount))
		}

		if state.sqrtPriceX96.Cmp(&step.sqrtPriceNextX96) == 0 {
			// if the tick is initialized, run the tick transition
			if step.initialized {
				//NOTE: This maybe tick.cross
				liquidityNetTick := p.Ticks[step.tickNext]
				liquidityNet := liquidityNetTick.LiquidityNet
				if zeroForOne {
					liquidityNet = *new(big.Int).Mul(&liquidityNet, NegativeOneBig)
				}

				state.liquidity.Set(addDelta(&state.liquidity, &liquidityNet))
				//				if liquidityNet.Cmp(ZeroBig) < 0 {
				//					liquidityBig := state.liquidity.ToBig()
				//					_liquidityCalc, _ := uint256.FromBig(new(big.Int).Sub(liquidityBig, new(big.Int).Mul(&liquidityNet, NegativeOneBig)))
				//					state.liquidity = *_liquidityCalc
				//				} else {
				//					liquidityNetBig, _ := uint256.FromBig(&liquidityNet)
				//					_liquidityCalc := new(uint256.Int).Add(&state.liquidity, liquidityNetBig)
				//					state.liquidity = *_liquidityCalc
				//				}
			}
			if zeroForOne {
				state.tick = step.tickNext - 1
			} else {
				state.tick = step.tickNext
			}
		} else if state.sqrtPriceX96.Cmp(&step.sqrtPriceStartX96) != 0 {
			state.tick, err = getTickAtSqrtRatio(&state.sqrtPriceX96)
			if err != nil {
				return nil, err
			}
		}
	}

	return &state.amountCalculated, nil
}
