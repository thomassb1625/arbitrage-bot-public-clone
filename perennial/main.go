package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/flashbots"
	"github.com/jmjac/ethGoTest/helperClient"
	perennial "github.com/jmjac/ethGoTest/perennial/abis"
	"github.com/jmjac/ethGoTest/secrets"
)

func main() {
	ADDR.MainnetAddresses(*log.Default())
	provider := "http://127.0.0.1:8111"
	//provider = "http://127.0.0.1:8545"
	collateralAddr := common.HexToAddress("0x2d264EBDb6632A06A1726193D4d37FeF1E5dbDcd")
	ethClient, err := helperClient.New(provider)
	if err != nil {
		panic(err)
	}

	fmt.Println("Syncing with block time")
	blockNum, _ := ethClient.BlockNumber(context.Background())
	for {
		newBlock, _ := ethClient.BlockNumber(context.Background())
		if blockNum != newBlock {
			blockNum = newBlock
			break
		}
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("Started checking for liquidations")

	userProducts := make(map[common.Address]map[common.Address]bool)
	userProducts = ParseLogsForAccounts(collateralAddr, nil, nil, ethClient, userProducts)

	mevReyals := make([]flashbots.MevRelay, 0)
	mevReyals = append(mevReyals, flashbots.NewFlashbotsMainnet())
	mevReyals = append(mevReyals, flashbots.NewMevRelay("https://rpc.beaverbuild.org/", "beaverbuilder"))
	mevReyals = append(mevReyals, flashbots.NewMevRelay("https://builder0x69.io", "builder0x69"))
	signer := flashbots.NewSigner(big.NewInt(1), secrets.PRIVATE_KEY_MIXED, ADDR.BUNDLE_EXECUTOR, ADDR.WETH, ethClient)
	for {
		nounce, _ := ethClient.NonceAt(context.Background(), signer.Address, nil)
		block, _ := ethClient.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
		baseFee := ethClient.CalcBaseFee(params.MainnetChainConfig, block.Header())

		now := time.Now()
		userProducts = ParseLogsForAccounts(collateralAddr, big.NewInt(int64(blockNum-1)), nil, ethClient, userProducts)
		fmt.Printf("Got %v users\n", len(userProducts))
		fmt.Println("Updating logs took: ", time.Since(now))

		now = time.Now()
		users := make([]common.Address, 0)
		products := make([]common.Address, 0)
		for k, v := range userProducts {
			for i := range v {
				users = append(users, k)
				products = append(products, i)
			}
		}
		profits := make([]*big.Int, 0, len(users))
		for i := 0; i < len(users); i += 200 {
			end := i + 200
			if end > len(users) {
				end = len(users)
			}
			partProfits, err := checkLiquidable(users[i:end], products[i:end], nil, ethClient)
			if err != nil {
				panic(err)
			}
			for _, p := range partProfits {
				profits = append(profits, p)
			}
		}

		//fmt.Printf("%v\n\n", profits)
		for i, profit := range profits {
			if profit.Cmp(big.NewInt(0)) != 0 {
				userToLiq := users[i]
				productToLiq := products[i]
				profitForLiq := convertDsuToEth(profits[i])
				fmt.Printf("Liquidate %v with product %v for profit %v\n", userToLiq, productToLiq, profitForLiq)
				txData, err := PrepareLiquidation(userToLiq, productToLiq)
				if err != nil {
					panic(err)
				}
				estimatedGas, err := signer.EstimateGasForTxToAddress(big.NewInt(0), txData, collateralAddr)
				estimatedGas = estimatedGas * 110 / 100
				if err != nil {
					log.Printf("\tFailed to estimate gas with error: %v\n", err)
					continue
				}
				bribeFlashbots := new(big.Int).Div(profitForLiq, big.NewInt(int64(estimatedGas)))
				bribeFlashbots.Mul(bribeFlashbots, big.NewInt(30)) //5% of profit as gas
				bribeFlashbots.Div(bribeFlashbots, big.NewInt(1000))
				if bribeFlashbots.Cmp(baseFee) <= 0 {
					log.Println("Not profitable")
					continue
				}

				//				bribeMempool := new(big.Int).Div(profitForLiq, big.NewInt(int64(estimatedGas)))
				//				bribeMempool.Mul(bribeMempool, big.NewInt(15)) //5% of profit as gas
				//				bribeMempool.Div(bribeMempool, big.NewInt(1000))

				totalCostFlashbots := new(big.Int).Mul(big.NewInt(int64(estimatedGas)), bribeFlashbots)
				//totalCostMempool := new(big.Int).Mul(big.NewInt(int64(estimatedGas)), bribeMempool)
				//fmt.Printf("GasUsage: \t%v\nGasFeeCap mempool: \t%v\nTotal Gas Fee Mempool: \t%v\nSubmitted for Block: \t%v\n", estimatedGas, bribeMempool, totalCostMempool, blockNum+1)
				//log.Println("Submit tx to mempool with 1.5% bribe")
				//signer.SentTxToAddress(txData, collateralAddr, big.NewInt(0), nounce, estimatedGas, bribeMempool, bribeMempool)
				log.Println("Submit tx to flashbots with 3.0% bribe")
				fmt.Printf("GasUsage: \t%v\nGasFeeCap mempool: \t%v\nTotal Gas Fee Mempool: \t%v\nSubmitted for Block: \t%v\nTx Data: \t%v\n", estimatedGas, bribeFlashbots, totalCostFlashbots, blockNum+1, hexutil.Encode(txData))
				signedTx, err := signer.SignTxToAddress(txData, collateralAddr, big.NewInt(0), nounce, estimatedGas, bribeFlashbots, bribeFlashbots)
				if err != nil {
					panic(err)
				}
				signedRawTx := flashbots.GetRawSignedTx(signedTx)
				for _, relay := range mevReyals {
					relay.SendBundleSingleTx(signedRawTx, fmt.Sprintf("0x%x", blockNum+1), *log.Default())
				}
				//fmt.Println("Exit")
				//os.Exit(0)
			}
		}

		fmt.Println("Checking for liquidations took: ", time.Since(now))
		log.Println("Waiting for new block")
		for {
			newBlock, _ := ethClient.BlockNumber(context.Background())
			if blockNum != newBlock {
				blockNum = newBlock
				break
			}
			time.Sleep(300 * time.Millisecond)
		}
		log.Printf("Found new block %v\n\n", blockNum)
	}

}

var (
	deposit = common.HexToHash("0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62")
)

func convertDsuToEth(dsu *big.Int) *big.Int {
	convertionRate := big.NewInt(800000000000000) //  0.return0008 ~ 1250DSU = 1 ETH
	convertionRate.Mul(dsu, convertionRate)
	convertionRate.Div(convertionRate, big.NewInt(1e18))
	return convertionRate
}

func ParseLogsForAccounts(coll common.Address, fromBlock, blockNum *big.Int, client *helperClient.Client, userProducts map[common.Address]map[common.Address]bool) map[common.Address]map[common.Address]bool {

	query := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   blockNum,
		Addresses: []common.Address{coll},
	}
	contractABI, _ := abi.JSON(strings.NewReader(string(perennial.CollateralABI)))
	contract := bind.NewBoundContract(coll, contractABI, client, client, client)

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		if vLog.Topics[0] == deposit {
			event := new(perennial.CollateralDeposit)
			if err := contract.UnpackLog(event, "Deposit", vLog); err != nil {
				panic(err)
			}
			if _, ok := userProducts[event.User]; !ok {
				userProducts[event.User] = make(map[common.Address]bool)
			}
			userProducts[event.User][event.Product] = true
		}
	}

	return userProducts
}

func SubmitLiquidation() {
}

func PrepareLiquidation(user, product common.Address) ([]byte, error) {
	contractABI, _ := abi.JSON(strings.NewReader(string(perennial.CollateralABI)))
	data, err := contractABI.Pack("liquidate", user, product)
	if err != nil {
		return nil, err
	}
	return data, nil
}
