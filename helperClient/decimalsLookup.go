package helperClient

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jmjac/ethGoTest/abis"
)

var decimalsBytecode = "0x608060405234801561001057600080fd5b50600436106100365760003560e01c8063162790551461003b578063d0cbe79714610075575b600080fd5b6100616004803603602081101561005157600080fd5b50356001600160a01b0316610135565b604080519115158252519081900360200190f35b6100e56004803603602081101561008b57600080fd5b8101906020810181356401000000008111156100a657600080fd5b8201836020820111156100b857600080fd5b803590602001918460208302840111640100000000831117156100da57600080fd5b50909250905061013b565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610121578181015183820152602001610109565b505050509050019250505060405180910390f35b3b151590565b6060808267ffffffffffffffff8111801561015557600080fd5b5060405190808252806020026020018201604052801561017f578160200160208202803683370190505b50905060005b8381101561027e576101b185858381811061019c57fe5b905060200201356001600160a01b0316610135565b6101d5576000198282815181106101c457fe5b602002602001018181525050610276565b8484828181106101e157fe5b905060200201356001600160a01b03166001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561022957600080fd5b505afa15801561023d573d6000803e3d6000fd5b505050506040513d602081101561025357600080fd5b5051825160ff9091169083908390811061026957fe5b6020026020010181815250505b600101610185565b50939250505056fea2646970667358221220493f13024f093186ee25989351a8d0cfde3de76d36c1d55c1d37a97d53186ed264736f6c634300060c0033"

//Decimales are -1 if contract was deleted
func (h Client) LookupDecimals(tokens []common.Address) ([]int64, error) {
	decimalsLookup, _ := abi.JSON(strings.NewReader(abis.DecimalLookupABI))
	approveData, _ := decimalsLookup.Pack("lookupDecimals", tokens)
	fakeLookupAddr := common.HexToAddress("0x07b237558c1B1ce38CAF4443C29d149f2b04d815")
	senderAddr := common.HexToAddress("0x5F5F3F32482504550B25e437444115d64517d4a8")

	resp, err := h.EthCallInjectBytecode(senderAddr, fakeLookupAddr, approveData, decimalsBytecode, fakeLookupAddr, "")
	if err != nil {
		return nil, err
	}

	var out []interface{}
	results := &out
	if results == nil {
		results = new([]interface{})
	}

	if len(*results) == 0 {
		res, err0 := decimalsLookup.Unpack("lookupDecimals", resp)
		err = err0
		*results = res
	} else {
		res := *results
		err = decimalsLookup.UnpackIntoInterface(res[0], "lookupDecimals", resp)
	}

	if err != nil {
		return nil, err
	}
	bigDec := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	decim := make([]int64, len(bigDec))

	for i, n := range bigDec {
		decim[i] = n.Int64()
	}
	return decim, nil
}
