package helperClient

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

type Client struct {
	*ethclient.Client
	geth *gethclient.Client
	RPC  *rpc.Client
}

func New(url string) (*Client, error) {
	clt := new(Client)
	c, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}
	clt.Client = ethclient.NewClient(c)
	clt.RPC = c
	clt.geth = gethclient.New(c)
	return clt, nil
}

type result struct {
	AccessList []types.AccessTuple `json:"accessList"`
	GasUsed    hexutil.Uint64      `json:"gasUsed"`
}

type acccessListResponse struct {
	AccessList []types.AccessTuple
	GasUsed    uint64
}

func (h Client) CreateAccessList(from, to common.Address, data []byte, amountIn *big.Int) (*types.AccessList, uint64, error) {
	msg := ethereum.CallMsg{}
	msg.Value = amountIn
	msg.From = from
	msg.Data = data
	msg.To = &to
	accessList, estimatedGas, _, err := h.geth.CreateAccessList(context.Background(), msg)
	if err != nil {
		return nil, 0, err
	}
	return accessList, estimatedGas, nil

	//	var resp result
	//	arg := map[string]interface{}{
	//		"from": from.Hex(),
	//		"to":   to.Hex(),
	//	}
	//	arg["data"] = hexutil.Bytes(data)
	//
	//	err := h.RPC.CallContext(context.Background(), &resp, "eth_createAccessList", arg)
	//	if err != nil {
	//		return nil, err
	//	}
	//	response := acccessListResponse{AccessList: resp.AccessList, GasUsed: uint64(resp.GasUsed)}
	//	return &response, nil
}

func toCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func (h Client) CalcBaseFee(config *params.ChainConfig, parent *types.Header) *big.Int {
	// If the current block is the first EIP-1559 block, return the InitialBaseFee.
	if !config.IsLondon(parent.Number) {
		return new(big.Int).SetUint64(params.InitialBaseFee)
	}

	parentGasTarget := parent.GasLimit / params.ElasticityMultiplier
	// If the parent gasUsed is the same as the target, the baseFee remains unchanged.
	if parent.GasUsed == parentGasTarget {
		return new(big.Int).Set(parent.BaseFee)
	}

	var (
		num   = new(big.Int)
		denom = new(big.Int)
	)

	if parent.GasUsed > parentGasTarget {
		// If the parent block used more gas than its target, the baseFee should increase.
		// max(1, parentBaseFee * gasUsedDelta / parentGasTarget / baseFeeChangeDenominator)
		num.SetUint64(parent.GasUsed - parentGasTarget)
		num.Mul(num, parent.BaseFee)
		num.Div(num, denom.SetUint64(parentGasTarget))
		num.Div(num, denom.SetUint64(params.BaseFeeChangeDenominator))
		baseFeeDelta := math.BigMax(num, common.Big1)

		return num.Add(parent.BaseFee, baseFeeDelta)
	} else {
		// Otherwise if the parent block used less gas than its target, the baseFee should decrease.
		// max(0, parentBaseFee * gasUsedDelta / parentGasTarget / baseFeeChangeDenominator)
		num.SetUint64(parentGasTarget - parent.GasUsed)
		num.Mul(num, parent.BaseFee)
		num.Div(num, denom.SetUint64(parentGasTarget))
		num.Div(num, denom.SetUint64(params.BaseFeeChangeDenominator))
		baseFee := num.Sub(parent.BaseFee, num)

		return math.BigMax(baseFee, common.Big0)
	}
}

func (h Client) EthCallInjectBytecode(from, to common.Address, data []byte, bytecode string, addrToOveride common.Address, blockNum string) ([]byte, error) {
	if blockNum == "" {
		blockNum = "latest"
	}
	var resp hexutil.Bytes
	arg := map[string]interface{}{
		"from": from.Hex(),
		"to":   to.Hex(),
	}
	arg["data"] = hexutil.Bytes(data)

	overideStates := map[string]interface{}{
		addrToOveride.Hex(): map[string]string{
			"code": bytecode,
		},
	}

	err := h.RPC.CallContext(context.Background(), &resp, "eth_call", arg, blockNum, overideStates)
	return resp, err
}

func (h Client) EstimateGasWithBlock(ctx context.Context, msg ethereum.CallMsg, block string) (uint64, error) {
	var hex hexutil.Uint64
	args := toCallArg(msg)

	err := h.RPC.CallContext(ctx, &hex, "eth_estimateGas", args, block)
	if err != nil {
		return 0, err
	}
	return uint64(hex), nil
}
