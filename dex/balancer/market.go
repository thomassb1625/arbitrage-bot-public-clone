package balancer

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Pair struct {
	Address common.Address
	Fee     *big.Int
}

func ParseContractLogsForAddresses(client bind.ContractBackend, addrBalancer common.Address, creationBlock *big.Int) []*Pair {
	pools := make([]*Pair, 0)

	query := ethereum.FilterQuery{
		FromBlock: creationBlock,
		Addresses: []common.Address{
			addrBalancer,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//contractAbi, err := abi.JSON(strings.NewReader(string(BalancerFactoryABI)))
	LOG_NEW_POOL := common.HexToHash("0x8ccec77b0cb63ac2cafd0f5de8cdfadab91ce656d262240ba8a6343bccc5f945")

	for _, vLog := range logs {
		if vLog.Removed || len(vLog.Topics) == 0 {
			continue
		}

		if vLog.Topics[0] != LOG_NEW_POOL {
			continue
		}

		addr := vLog.Topics[2]
		poolAddr := common.HexToAddress(addr.Hex())
		pair := Pair{Address: poolAddr}
		pools = append(pools, &pair)
	}
	return pools
}
