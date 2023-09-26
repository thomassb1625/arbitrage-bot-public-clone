package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/helperClient"
	"github.com/jmjac/ethGoTest/mempool/abis"
)

/*
TODO:
	- Figure out how to check which pools are involved in tx
	- Check if the tx in the mempool has a valid nounce
	- Update those pool reserves and look for simple arbs with these pools, maybe have a tool to check which trades now will return different values
	- Simulate the bundle (some tx and our backrun)
*/

/*
	Ideas:
		- Check function signature to detect changes in liquidity for different pools
		- Check routers for all protocols
		- Start with only simple arbs like ETH->x->ETH or ETH -> x -> y -> ETH


*/
func main() {
	//Add websocket functionality to xyznode
	provider := "ws://127.0.0.1:8111"
	flag.StringVar(&provider, "provider", provider, "RPC endpoint")
	flag.Parse()

	//	client, err := ethclient.Dial(provider)
	//	if err != nil {
	//		fmt.Println("Coudn't create client,", err)
	//		return
	//	}
	helperClient, err := helperClient.New(provider)
	if err != nil {
		fmt.Println("Coudn't create helper client,", err)
		return
	}

	//Setup adreses
	ADDR.MainnetAddresses()

	geth := gethclient.New(helperClient.RPC)
	txnsHash := make(chan common.Hash)
	_, err = geth.SubscribePendingTransactions(context.Background(), txnsHash)

	if err != nil {
		log.Fatal("Error when creating subscriber: ", err)
	}

	//signer := types.NewLondonSigner(big.NewInt(1))
	routerV2 := common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D") //uniswapv2 router
	routerV3 := common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564") //uniswapv3 router

	//WARNING: Erigon updates mempool every 100ms
	abiV202, err := abi.JSON(strings.NewReader(abis.UniV2Router02ABI))
	if err != nil {
		log.Fatal("Failed with router02 abi: ", err)
	}

	fmt.Printf("\nStarted reading mempool!\n\n")
	for txnHash := range txnsHash {
		txn, _, err := helperClient.TransactionByHash(context.Background(), txnHash)
		if err != nil {
			continue
		}

		//	message, err := txn.AsMessage(signer, nil)
		//	if err != nil {
		//		log.Fatalln(err)
		//	}
		if txn.To() == nil {
			continue
		}

		if *txn.To() == routerV3 {
			fmt.Printf("Tx to v3 router - %v\n", txn.Hash())
		}
		if *txn.To() == routerV2 {
			now := time.Now()
			fmt.Printf("Tx to v2 router - %v\n", txn.Hash())
			fmt.Printf("Method sig: %v\n", hexutil.Encode(txn.Data()[:4]))
			DecodeTx(&abiV202, txn.Data())
			fmt.Println("Took:", time.Since(now))
		}
	}

}
