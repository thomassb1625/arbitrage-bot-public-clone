package dodo

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/dex"
)

var (
	DVM_FACTORY = common.HexToAddress("0x72d220cE168C4f361dD4deE5D826a01AD8598f6C")
	DVM_NEW     = common.HexToHash("0xaf5c5f12a80fc937520df6fcaed66262a4cc775e0f3fceaf7a7cfe476d9a751d")
	DVM_REMOVE  = common.HexToHash("0x63971a172c674ce2e9da5e027e9e81a54fd3aa74a2c246a2eb473dc0aa7f5cdd")
	DSP_FACTORY = common.HexToAddress("0x6fdDB76c93299D985f4d3FC7ac468F9A168577A4")
	DSP_NEW     = common.HexToHash("0xbc1083a2c1c5ef31e13fb436953d22b47880cf7db279c2c5666b16083afd6b9d")
	DSP_REMOVE  = common.HexToHash("")
)

type Pair struct {
	addr       common.Address
	baseToken  dex.Token //common.Address
	quoteToken dex.Token //common.Address
	mtFeeRate  *uint256.Int
	lpFeeRate  *uint256.Int
	state      *pmmState
	protocol   string
}

type pmmState struct {
	I  *uint256.Int
	K  *uint256.Int
	B  *uint256.Int
	Q  *uint256.Int
	B0 *uint256.Int
	Q0 *uint256.Int
	R  *uint256.Int
}

func initEmptyPmmState() *pmmState {
	s := &pmmState{
		I:  new(uint256.Int),
		K:  new(uint256.Int),
		B:  new(uint256.Int),
		Q:  new(uint256.Int),
		B0: new(uint256.Int),
		Q0: new(uint256.Int),
		R:  new(uint256.Int),
	}
	return s
}

// Using addresses for pools, update their reserves and return a pair struct describing the pool
func UpdateReserves(client *ethclient.Client, pools []*Pair, isDVM bool) []Pair {
	var pairs []Pair

	for i, p := range pools {

		instance, err := NewDVM(p.addr, client)

		if err != nil {
			log.Fatal(err)
		}

		newState, err := instance.GetPMMStateForCall(nil)

		if err != nil {
			log.Fatal(err)
		}

		// Don't log pool if reserves are zero
		//		if len(newState.B.Bits()) == 0 {
		//			continue
		//		}
		//
		//		if len(newState.Q.Bits()) == 0 {
		//			continue
		//		}

		//base, err := instance.BASETOKEN(nil)
		//quote, err := instance.QUOTETOKEN(nil)

		fees, err := instance.GetUserFeeRate(nil, CALLER)

		if err != nil {
			log.Fatal(err)
		}

		lpFeeRate, _ := uint256.FromBig(fees.LpFeeRate)
		mtFeeRate, _ := uint256.FromBig(fees.MtFeeRate)
		I, _ := uint256.FromBig(newState.I)
		K, _ := uint256.FromBig(newState.K)
		B, _ := uint256.FromBig(newState.B)
		Q, _ := uint256.FromBig(newState.Q)
		B0, _ := uint256.FromBig(newState.B0)
		Q0, _ := uint256.FromBig(newState.Q0)

		// R-Value depends on pool type

		var R *uint256.Int

		if isDVM {
			R = uint256.NewInt(2)
		} else {
			R, _ = uint256.FromBig(newState.R)
		}
		pairs[i].lpFeeRate.Set(lpFeeRate)
		pairs[i].mtFeeRate.Set(mtFeeRate)
		pairs[i].state.I.Set(I)
		pairs[i].state.K.Set(K)
		pairs[i].state.B.Set(B)
		pairs[i].state.Q.Set(Q)
		pairs[i].state.B0.Set(B0)
		pairs[i].state.Q0.Set(Q0)
		pairs[i].state.R.Set(R)
	}

	return pairs
}

func GetDVMPools(client *ethclient.Client, tokensMap map[common.Address]dex.Token) ([]*Pair, error) {
	addresses := getDVMPoolAddresses(client)

	pairs := make([]*Pair, 0)
	for _, addr := range addresses {
		instance, err := NewDVM(addr, client)
		if err != nil {
			return nil, err
		}
		base, err := instance.BASETOKEN(nil)
		quote, err := instance.QUOTETOKEN(nil)
		var baseToken dex.Token
		if t, ok := tokensMap[base]; ok {
			baseToken = t
		} else {
			baseToken = dex.NewNotTaxed(base)
		}

		var quoteToken dex.Token
		if t, ok := tokensMap[quote]; ok {
			quoteToken = t
		} else {
			quoteToken = dex.NewNotTaxed(base)
		}

		p := &Pair{addr: addr, baseToken: baseToken, quoteToken: quoteToken, mtFeeRate: new(uint256.Int), lpFeeRate: new(uint256.Int), state: initEmptyPmmState(), protocol: "dodoDVM"}
		pairs = append(pairs, p)
	}
	UpdateReserves(client, pairs, true)
	return pairs, nil
}

// Get all adresses for DVM pools (THESE ARE NOT READY YET)
func getDVMPoolAddresses(client *ethclient.Client) []common.Address {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(10000000),
		Addresses: []common.Address{
			DVM_FACTORY,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)

	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(DVMFactoryABI)))

	if err != nil {
		log.Fatal(err)
	}

	contract := bind.NewBoundContract(DVM_FACTORY, contractAbi, client, client, client)

	newPools := make([]common.Address, 0)
	remPools := make([]common.Address, 0)

	for _, vLog := range logs {

		if vLog.Removed || len(vLog.Data) == 0 {
			continue
		}

		switch vLog.Topics[0] {
		case DVM_NEW:
			event := new(DVMFactoryNewDVM)
			err := contract.UnpackLog(event, "NewDVM", vLog)

			if err != nil {
				panic(err)
			}

			newPools = append(newPools, event.Dvm)

		case DVM_REMOVE:
			event := new(DVMFactoryRemoveDVM)
			err := contract.UnpackLog(event, "RemoveDVM", vLog)

			if err != nil {
				panic(err)
			}

			remPools = append(remPools, event.Dvm)
		}
	}

	pools := make([]common.Address, 0)

	//Loop through pools to figure out if they've been removed
	for _, pool := range newPools {
		poolRemoved := false

		for _, removed := range remPools {
			if pool == removed {
				poolRemoved = true
			}
		}

		if !poolRemoved {
			pools = append(pools, pool)
		}
	}
	return pools
}

// Get all adresses for DSP pools
func getDSPPoolAddresses(client *ethclient.Client) []common.Address {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(10000000),
		Addresses: []common.Address{
			DSP_FACTORY,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)

	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(DSPFactoryABI)))

	if err != nil {
		log.Fatal(err)
	}

	contract := bind.NewBoundContract(DSP_FACTORY, contractAbi, client, client, client)

	pools := make([]common.Address, 0)

	for _, vLog := range logs {

		if vLog.Removed || len(vLog.Data) == 0 {
			continue
		}

		event := new(DSPFactoryNewDSP)
		err := contract.UnpackLog(event, "NewDSP", vLog)

		if err != nil {
			panic(err)
		}

		pools = append(pools, event.DSP)

	}

	return pools
}
