package dodo

import (
	DecimalMath "github.com/jmjac/ethGoTest/dex/dodov2/decimalMath"
	DodoMath "github.com/jmjac/ethGoTest/dex/dodov2/dodoMath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

//WARNING: Change this to use ADDR.BUNDLE_EXECUTOR
var CALLER = common.HexToAddress("0x0000008200bE7f00920000d500eCd5F6fafb8386")

// ============ final amount out ============

func querySellBase(p Pair, payBaseAmount *uint256.Int) *uint256.Int {
	receiveQuoteAmount := sellBaseToken(p.state, payBaseAmount)
	println("Recieve quote amount is", receiveQuoteAmount.ToBig().String())
	mtFee := DecimalMath.MulFloor(receiveQuoteAmount, p.mtFeeRate)
	lpFee := DecimalMath.MulFloor(receiveQuoteAmount, p.lpFeeRate)
	fees := new(uint256.Int).Add(mtFee, lpFee)

	return new(uint256.Int).Sub(receiveQuoteAmount, fees)
}

func querySellQuote(p Pair, payQuoteAmount *uint256.Int) *uint256.Int {
	receiveBaseAmount := sellQuoteToken(p.state, payQuoteAmount)

	mtFee := DecimalMath.MulFloor(receiveBaseAmount, p.mtFeeRate)
	lpFee := DecimalMath.MulFloor(receiveBaseAmount, p.lpFeeRate)
	fees := new(uint256.Int).Add(lpFee, mtFee)

	return new(uint256.Int).Sub(receiveBaseAmount, fees)
}

// ============ buy & sell token outputs ============

func sellBaseToken(state *pmmState, payBaseAmount *uint256.Int) *uint256.Int {

	var receiveQuoteAmount *uint256.Int

	if state.R.Cmp(uint256.NewInt(1)) == 0 {
		// case 1: R=1
		// R falls below one
		receiveQuoteAmount = ROneSellBaseToken(state, payBaseAmount)

	} else if state.R.Cmp(uint256.NewInt(1)) > 0 {

		var backToOnePayBase *uint256.Int
		if state.B0.Cmp(state.B) > 0 {
			backToOnePayBase = new(uint256.Int).Sub(state.B0, state.B)
		} else {
			backToOnePayBase = uint256.NewInt(0)
		}

		var backToOneReceiveQuote *uint256.Int
		if state.Q.Cmp(state.Q0) > 0 {
			backToOneReceiveQuote = new(uint256.Int).Sub(state.Q, state.Q0)
		} else {
			backToOneReceiveQuote = uint256.NewInt(0)
		}

		// case 2: R>1
		// complex case, R status depends on trading amount
		if payBaseAmount.Cmp(backToOnePayBase) < 0 {
			// case 2.1: R status do not change
			receiveQuoteAmount = RAboveSellBaseToken(state, payBaseAmount)
			if receiveQuoteAmount.Cmp(backToOneReceiveQuote) > 0 {
				// [Important corner case!] may enter this branch when some precision problem happens. And consequently contribute to negative spare quote amount
				// to make sure spare quote>=0, mannually set receiveQuote=backToOneReceiveQuote
				receiveQuoteAmount = backToOneReceiveQuote
			}

		} else if payBaseAmount.Cmp(backToOnePayBase) == 0 {
			// case 2.2: R status changes to ONE
			receiveQuoteAmount = backToOneReceiveQuote

		} else {
			// case 2.3: R status changes to BELOW_ONE
			rOne := ROneSellBaseToken(state, new(uint256.Int).Sub(payBaseAmount, backToOnePayBase))

			if rOne != nil {
				receiveQuoteAmount = new(uint256.Int).Add(backToOneReceiveQuote, rOne)
			} else {
				receiveQuoteAmount = uint256.NewInt(0)
			}
		}

	} else {
		// state.R == RState.BELOW_ONE
		// case 3: R<1
		receiveQuoteAmount = RBelowSellBaseToken(state, payBaseAmount)
	}

	if receiveQuoteAmount != nil {
		return receiveQuoteAmount
	} else {
		return uint256.NewInt(0)
	}
}

func sellQuoteToken(state *pmmState, payQuoteAmount *uint256.Int) *uint256.Int {

	var receiveBaseAmount *uint256.Int

	if state.R.Cmp(uint256.NewInt(1)) == 0 {
		receiveBaseAmount = ROneSellQuoteToken(state, payQuoteAmount)

	} else if state.R.Cmp(uint256.NewInt(1)) > 0 {
		receiveBaseAmount = RAboveSellQuoteToken(state, payQuoteAmount)

	} else {
		var backToOnePayQuote *uint256.Int
		if state.Q0.Cmp(state.Q) > 0 {
			backToOnePayQuote = new(uint256.Int).Sub(state.Q0, state.Q)
		} else {
			backToOnePayQuote = uint256.NewInt(0)
		}

		var backToOneReceiveBase *uint256.Int
		if state.B.Cmp(state.B0) > 0 {
			backToOneReceiveBase = new(uint256.Int).Sub(state.B, state.B0)
		} else {
			backToOneReceiveBase = uint256.NewInt(0)
		}

		if payQuoteAmount.Cmp(backToOnePayQuote) < 0 {
			receiveBaseAmount = RBelowSellQuoteToken(state, payQuoteAmount)

			if receiveBaseAmount.Cmp(backToOneReceiveBase) > 0 {
				receiveBaseAmount = backToOneReceiveBase
			}

		} else if payQuoteAmount.Cmp(backToOnePayQuote) == 0 {
			receiveBaseAmount = backToOneReceiveBase

		} else {
			rOne := ROneSellQuoteToken(state, new(uint256.Int).Sub(payQuoteAmount, backToOnePayQuote))

			if rOne != nil {
				receiveBaseAmount = new(uint256.Int).Add(backToOneReceiveBase, rOne)
			} else {
				receiveBaseAmount = uint256.NewInt(0)
			}
		}
	}

	if receiveBaseAmount != nil {
		return receiveBaseAmount
	} else {
		return uint256.NewInt(0)
	}
}

// ============ R = 1 cases ============

func ROneSellBaseToken(state *pmmState, payBaseAmount *uint256.Int) *uint256.Int {
	// in theory Q2 <= targetQuoteTokenAmount
	// however when amount is close to 0, precision problems may cause Q2 > targetQuoteTokenAmount
	return DodoMath.SolveQuadraticFunctionForTrade(
		state.Q0,
		state.Q0,
		payBaseAmount,
		state.I,
		state.K)
}

func ROneSellQuoteToken(state *pmmState, payQuoteAmount *uint256.Int) *uint256.Int {
	return DodoMath.SolveQuadraticFunctionForTrade(
		state.B0,
		state.B0,
		payQuoteAmount,
		DecimalMath.ReciprocalFloor(state.I),
		state.K)
}

// ============ R < 1 cases ============

func RBelowSellBaseToken(state *pmmState, payBaseAmount *uint256.Int) *uint256.Int {
	return DodoMath.SolveQuadraticFunctionForTrade(
		state.Q0,
		state.Q,
		payBaseAmount,
		state.I,
		state.K)
}

func RBelowSellQuoteToken(state *pmmState, payQuoteAmount *uint256.Int) *uint256.Int {
	return DodoMath.GeneralIntegrate(
		state.Q0,
		new(uint256.Int).Add(state.Q, payQuoteAmount),
		state.Q,
		DecimalMath.ReciprocalFloor(state.I),
		state.K)
}

// ============ R > 1 cases ============

func RAboveSellBaseToken(state *pmmState, payBaseAmount *uint256.Int) *uint256.Int {
	return DodoMath.GeneralIntegrate(
		state.B0,
		new(uint256.Int).Add(state.B, payBaseAmount),
		state.B,
		state.I,
		state.K)
}

func RAboveSellQuoteToken(state *pmmState, payQuoteAmount *uint256.Int) *uint256.Int {
	return DodoMath.SolveQuadraticFunctionForTrade(
		state.B0,
		state.B,
		payQuoteAmount,
		DecimalMath.ReciprocalFloor(state.I),
		state.K)
}
