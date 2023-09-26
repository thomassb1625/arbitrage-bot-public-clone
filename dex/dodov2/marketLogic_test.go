package dodo

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
)

var providerUrl = "https://mainnet.infura.io/v3/cd5fd6cadfd34f269fe890bfedd5d23a"
var INPUT = uint256.NewInt(1)
var DSP_SELL_BASE_0_ERROR = common.HexToAddress("0x138c825d993D5fFb7F028408e870ebf50A019543")

func TestDVMMarket(t *testing.T) {
	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		t.Fatal(err)
	}

	pools, err := GetDVMPools(client)
	if err != nil {
		t.Fatal(err)
	}

	pairs := UpdateReserves(client, pools, true)

	totalPools := 0
	succesfulSellBase := 0
	succesfulSellQuote := 0

	for _, p := range pairs {

		testPool, err := NewDVM(p.addr, client)

		if err != nil {
			t.Fatal(err)
		}

		// Test sell base
		quote, err := testPool.QuerySellBase(nil, CALLER, INPUT.ToBig())

		var q *uint256.Int

		if err != nil {
			q = uint256.NewInt(0)
		} else {
			newQuote, _ := uint256.FromBig(quote.ReceiveQuoteAmount)
			q = newQuote
		}

		testQuote := querySellBase(&p, INPUT)

		if q.Cmp(testQuote) != 0 {
			println("The following pool failed after attempting to sell base: ")
			println(p.addr.String())
			println("Got ", testQuote.ToBig().String(), " but expected ", q.ToBig().String())
			println()
		} else {
			succesfulSellBase += 1
		}

		// Test sell quote
		base, err := testPool.QuerySellQuote(nil, CALLER, INPUT.ToBig())

		var b *uint256.Int

		if err != nil {
			b = uint256.NewInt(0)
		} else {
			newBase, _ := uint256.FromBig(base.ReceiveBaseAmount)
			b = newBase
		}

		testBase := querySellQuote(&p, INPUT)

		if b.Cmp(testBase) != 0 {
			println("The following pool failed after attempting to sell quote: ")
			println(p.addr.String())
			println("Got ", testBase.ToBig().String(), " but expected ", b.ToBig().String())
			println()
		} else {
			succesfulSellQuote += 1
		}

		totalPools += 1
	}
	println("Final testing results: ")
	println("Sell base succesful tests: ", succesfulSellBase, "/", totalPools)
	println("Sell quote succesful tests: ", succesfulSellQuote, "/", totalPools)
}

//func TestDSPMarket(t *testing.T) {
//	client, err := ethclient.Dial(CLIENT)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	pools := getDSPPoolAddresses(client)
//
//	pairs := UpdateReserves(client, pools, false)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	totalPools := 0
//	succesfulSellBase := 0
//	succesfulSellQuote := 0
//
//	for _, p := range pairs {
//
//		testPool, err := NewDVM(*p.addr, client)
//
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		// Test sell base
//		quote, err := testPool.QuerySellBase(nil, CALLER, INPUT.ToBig())
//
//		var q *uint256.Int
//
//		if err != nil {
//			q = uint256.NewInt(0)
//		} else {
//			newQuote, _ := uint256.FromBig(quote.ReceiveQuoteAmount)
//			q = newQuote
//		}
//
//		testQuote := querySellBase(&p, INPUT)
//
//		if q.Cmp(testQuote) != 0 {
//			println("The following pool failed after attempting to sell base: ")
//			println(p.addr.String())
//			println("Got ", testQuote.ToBig().String(), " but expected ", q.ToBig().String())
//			println()
//		} else {
//			succesfulSellBase += 1
//		}
//
//		// Test sell quote
//		base, err := testPool.QuerySellQuote(nil, CALLER, INPUT.ToBig())
//
//		var b *uint256.Int
//
//		if err != nil {
//			b = uint256.NewInt(0)
//		} else {
//			newBase, _ := uint256.FromBig(base.ReceiveBaseAmount)
//			b = newBase
//		}
//
//		testBase := querySellQuote(&p, INPUT)
//
//		if b.Cmp(testBase) != 0 {
//			println("The following pool failed after attempting to sell quote: ")
//			println(p.addr.String())
//			println("Got ", testBase.ToBig().String(), " but expected ", b.ToBig().String())
//			println("R value was: ", p.state.R.ToBig().String())
//			println()
//		} else {
//			succesfulSellQuote += 1
//		}
//
//		totalPools += 1
//	}
//	println("Final testing results: ")
//	println("Sell base succesful tests: ", succesfulSellBase, "/", totalPools)
//	println("Sell quote succesful tests: ", succesfulSellQuote, "/", totalPools)
//}
//
//func TestPool(t *testing.T) {
//	client, err := ethclient.Dial(CLIENT)
//	if err != nil {
//		log.Fatal(err)
//	}
//	testPoolAddr := DSP_SELL_BASE_0_ERROR
//
//	pool, err := NewDVM(testPoolAddr, client)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	testPair := UpdateReserves(client, []*common.Address{&testPoolAddr}, true)[0]
//
//	testSellBase := true
//
//	print("B: ")
//	println(testPair.state.B.ToBig().String())
//	print("B0: ")
//	println(testPair.state.B0.ToBig().String())
//	print("Q: ")
//	println(testPair.state.Q.ToBig().String())
//	print("Q0: ")
//	println(testPair.state.Q0.ToBig().String())
//	print("R: ")
//	println(testPair.state.R.ToBig().String())
//
//	if testSellBase {
//		quote, err := pool.QuerySellBase(nil, CALLER, INPUT.ToBig())
//
//		var q *uint256.Int
//
//		if err != nil {
//			q = uint256.NewInt(0)
//		} else {
//			newQuote, _ := uint256.FromBig(quote.ReceiveQuoteAmount)
//			q = newQuote
//		}
//
//		testQuote := querySellBase(&testPair, INPUT)
//
//		if q.Cmp(testQuote) != 0 {
//			t.Fatal("Sell base failed: \n Got ", testQuote.ToBig().String(), " but expected ", q.ToBig().String())
//		}
//
//	} else {
//		base, err := pool.QuerySellQuote(nil, CALLER, INPUT.ToBig())
//
//		var b *uint256.Int
//
//		if err != nil {
//			b = uint256.NewInt(0)
//		} else {
//			newBase, _ := uint256.FromBig(base.ReceiveBaseAmount)
//			b = newBase
//		}
//
//		testBase := querySellQuote(&testPair, INPUT)
//
//		if b.Cmp(testBase) != 0 {
//			t.Fatal("Sell quote failed: \n Got ", testBase.ToBig().String(), " but expected ", b.ToBig().String())
//		}
//	}
//}
