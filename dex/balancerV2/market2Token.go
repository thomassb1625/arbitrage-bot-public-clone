package balancerV2

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/abis"
	"github.com/jmjac/ethGoTest/dex"
)

//Pool can have at most 8 tokens
//Check if the pool is paused
type Weighted2TokenPair struct {
	spender  common.Address //Vault address
	address  common.Address
	PoolId   [32]byte
	sc0      *uint256.Int
	sc1      *uint256.Int
	b0       *uint256.Int
	b1       *uint256.Int
	w0       *uint256.Int
	w1       *uint256.Int
	token0   dex.Token
	token1   dex.Token
	fee      *uint256.Int
	protocol string
}

func parse2TokenLogs(client bind.ContractBackend, addrWeighted common.Address, creationBlock *big.Int) []*Weighted2TokenPair {
	pools := make([]*Weighted2TokenPair, 0)

	query := ethereum.FilterQuery{
		FromBlock: creationBlock,
		Addresses: []common.Address{
			addrWeighted,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//contractAbi, err := abi.JSON(strings.NewReader(string(BWeightedFactoryABI)))
	//TODO: Change to use event topics or something
	LOG_POOLCREATED := common.HexToHash("0x83a48fbcfc991335314e74d0496aab6a1987e992ddc85dddbcc4d6dd6ef2e9fc")

	for _, vLog := range logs {
		if vLog.Removed || len(vLog.Topics) == 0 {
			continue
		}

		if vLog.Topics[0] != LOG_POOLCREATED {
			continue
		}

		addr := vLog.Topics[1]
		poolAddr := common.HexToAddress(addr.Hex())
		pool, err := NewBBasedWeightedPoolCaller(poolAddr, client)
		if err != nil {
			log.Fatal("PoolCreation ", err)
		}
		state, err := pool.GetPausedState(nil)
		if err != nil {
			log.Fatal("Pool paused states ", err)
		}
		if state.Paused {
			continue
		}
		pair := Weighted2TokenPair{address: poolAddr, fee: uint256.NewInt(0), protocol: "balancerV2"}
		pools = append(pools, &pair)
	}
	return pools
}

func Get2TokenWeightedMarket(client bind.ContractBackend, addrVault, addrWeighted common.Address, creationBlock *big.Int, BLACKLIST map[common.Address]bool, tokenMap map[common.Address]dex.Token) []*Weighted2TokenPair {
	now := time.Now()
	dex.Logger.Println("\n============== Setting up BalancerV2 2 Tokens Weighted Pools ===================")
	//TODO: USE THE BLACKLIST
	pairs := parse2TokenLogs(client, addrWeighted, creationBlock)
	vault, _ := NewVaultCaller(addrVault, client)
	for index, p := range pairs {
		pool, err := NewBBasedWeightedPool(p.address, client)

		if err != nil {
			log.Fatal("Pool creation in weighted ", err)
		}

		weights, _ := pool.GetNormalizedWeights(nil)
		w0, _ := uint256.FromBig(weights[0])
		w1, _ := uint256.FromBig(weights[1])
		pairs[index].w0 = w0
		pairs[index].w1 = w1

		id, _ := pool.GetPoolId(nil)
		pairs[index].PoolId = id
		fee, _ := pool.GetSwapFeePercentage(nil)
		feeUint, _ := uint256.FromBig(fee)
		pairs[index].fee = feeUint

		poolTokens, err := vault.GetPoolTokens(nil, id)
		if err != nil {
			log.Fatal("Vault get pool tokens", err)
		}
		b0, _ := uint256.FromBig(poolTokens.Balances[0])
		pairs[index].b0 = b0

		t0Token, ok := tokenMap[poolTokens.Tokens[0]]
		if !ok {
			t0Token = dex.NewNotTaxed(poolTokens.Tokens[0])
		}
		pairs[index].token0 = t0Token

		b1, _ := uint256.FromBig(poolTokens.Balances[1])
		pairs[index].b1 = b1

		t1Token, ok := tokenMap[poolTokens.Tokens[1]]
		if !ok {
			t1Token = dex.NewNotTaxed(poolTokens.Tokens[1])
		}
		pairs[index].token1 = t1Token

		t0, _ := abis.NewIERC20Caller(poolTokens.Tokens[0], client)
		decimals, _ := t0.Decimals(nil)
		decimalDifference := uint64(18 - decimals)
		scallingFactor := new(uint256.Int).Exp(uint256.NewInt(10), uint256.NewInt(decimalDifference))
		scallingFactor.Mul(ONE_E18, scallingFactor)
		pairs[index].sc0 = scallingFactor

		t1, _ := abis.NewIERC20Caller(poolTokens.Tokens[0], client)
		decimals, _ = t1.Decimals(nil)
		decimalDifference = uint64(18 - decimals)
		scallingFactor = new(uint256.Int).Exp(uint256.NewInt(10), uint256.NewInt(decimalDifference))
		scallingFactor.Mul(ONE_E18, scallingFactor)
		pairs[index].sc1 = scallingFactor
		pairs[index].spender = addrVault

	}
	dex.Logger.Printf("Loaded %v pools from BalancerV2 2 Tokens Weighted Pools\n", len(pairs))
	dex.Logger.Printf("Took: %v\n", time.Since(now))
	return pairs
}
func UpdateReserves2TokensWeighted(client bind.ContractBackend, addrVault common.Address, pairs []*Weighted2TokenPair) {
	vault, _ := NewVaultCaller(addrVault, client)
	var wg sync.WaitGroup
	for index := 0; index < len(pairs); index++ {
		wg.Add(1)
		go func(index int) {
			poolTokens, _ := vault.GetPoolTokens(nil, pairs[index].PoolId)
			b0, _ := uint256.FromBig(poolTokens.Balances[0])
			pairs[index].b0 = b0

			b1, _ := uint256.FromBig(poolTokens.Balances[1])
			pairs[index].b1 = b1
			wg.Done()
		}(index)
	}
	wg.Wait()
}

/* TODO: Add this later as a safety check
uint256 internal constant _MAX_IN_RATIO = 0.3e18;
uint256 internal constant _MAX_OUT_RATIO = 0.3e18;
*/
func (wp Weighted2TokenPair) GetTokensOut(tokenIn, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int {
	bi, wi, sci, bo, wo, sco := wp.tokensToBalancesAndWeights(tokenIn, tokenOut)

	if amountIn.IsZero() || bi.IsZero() || bo.IsZero() {
		return uint256.NewInt(0)
	}

	bi = upscale(bi, sci)
	if amountIn.Cmp(mulDown(bi, MAX_IN_RATIO)) > 0 {
		return uint256.NewInt(0)
	}
	//bi.Mul(bi, sci)
	//bi.Div(bi, ONE_E18)
	bo = upscale(bo, sco)
	//bo.Mul(bo, sco)
	//bo.Div(bo, ONE_E18)
	feeAmount := mulUp(amountIn, wp.fee)

	//ai := upscale(feeAmount.Sub(amountIn, feeAmount), sci)
	_, underflow := feeAmount.SubOverflow(amountIn, feeAmount)
	if underflow {
		return uint256.NewInt(0)
	}
	feeAmount.Mul(feeAmount, sci)
	feeAmount.Div(feeAmount, ONE_E18)

	res := downscaleDown(wp.getAmountOut(feeAmount, bi, wi, bo, wo), sco)
	//dex.Logger.Printf("res = %v, result = %v\n", res, resultUint)
	//WARNING: Temporary fix when balancer overestimes slightly
	return res
}

func (wp Weighted2TokenPair) getAmountOut(ai, bi, wi *uint256.Int, bo, wo *uint256.Int) *uint256.Int {
	denominator := ai.Add(ai, bi)
	base := divUp(bi, denominator)
	exponent := divDown(wi, wo)
	power := powUp(base, exponent)
	return mulDown(bo, complement(power))
}

func (wp Weighted2TokenPair) tokensToBalancesAndWeights(t0, t1 common.Address) (*uint256.Int, *uint256.Int, *uint256.Int, *uint256.Int, *uint256.Int, *uint256.Int) {
	var weightT0, balanceT0, scallingT0, weightT1, balanceT1, scallingT1 *uint256.Int
	if wp.token0.Addr == t0 {
		balanceT0 = wp.b0
		weightT0 = wp.w0
		scallingT0 = wp.sc0
		balanceT1 = wp.b1
		weightT1 = wp.w1
		scallingT1 = wp.sc1
	} else if wp.token0.Addr == t1 {
		balanceT0 = wp.b1
		weightT0 = wp.w1
		scallingT0 = wp.sc1
		balanceT1 = wp.b0
		weightT1 = wp.w0
		scallingT1 = wp.sc0
	}

	if balanceT0 == nil || balanceT1 == nil {
		err := fmt.Sprintf("Wrong tokenAddress passed to the pool! Got %v and %v. Pool has %v and %v\n", t0, t1, wp.token0, wp.token1)
		panic(err)
	}
	return balanceT0, weightT0, scallingT0, balanceT1, weightT1, scallingT1
}

func (wp Weighted2TokenPair) Protocol() string {
	return wp.protocol
}

func (wp Weighted2TokenPair) Reserves() (*uint256.Int, *uint256.Int) {
	return new(uint256.Int).Set(wp.b0), new(uint256.Int).Set(wp.b1)
}

//TODO: This should be vautl adderss
func (wp Weighted2TokenPair) MarketAddress() common.Address {
	return wp.address
}

func (wp Weighted2TokenPair) SpenderAddress() common.Address {
	return wp.spender
}

func (wp Weighted2TokenPair) GetOtherToken(tokenIn common.Address) dex.Token {
	if tokenIn == wp.token0.Addr {
		return wp.token1
	} else {
		return wp.token0
	}
}

func (wp Weighted2TokenPair) Tokens() (dex.Token, dex.Token) {
	return wp.token0, wp.token1
}

//Approve anything else as 1e36
var approveAmount, _ = new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 10)
var deadline, _ = new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 10)

//TODO: Implement batch swaps
//WARNING: This will all utlize vault address as the destination
//TODO: Add and approval for hughe amount like 2^240 and then remove set the approval to 0
func (wp Weighted2TokenPair) SellTokens(tokenIn common.Address, amountIn *uint256.Int, recipient common.Address, sender common.Address) ([]common.Address, [][]byte) {
	//Kind = 0 GivenIn, KIND = 1 GivenOut
	//Approve token
	//fmt.Printf("Balancer %v Prepared swap for %v\n", wp.MarketAddress(), amountIn)
	targets := make([]common.Address, 0)
	payloads := make([][]byte, 0)

	IERC20ABI, _ := abi.JSON(strings.NewReader(abis.IERC20ABI))
	approveData, _ := IERC20ABI.Pack("approve", wp.spender, approveAmount)
	targets = append(targets, tokenIn)
	payloads = append(payloads, approveData)

	funds := IVaultFundManagement{Recipient: recipient, Sender: sender, FromInternalBalance: false, ToInternalBalance: false}
	singleSwap := IVaultSingleSwap{PoolId: wp.PoolId, AssetIn: tokenIn, AssetOut: wp.GetOtherToken(tokenIn).Addr, Amount: amountIn.ToBig(), Kind: 0}
	parsed, _ := abi.JSON(strings.NewReader(VaultABI))
	txData, err := parsed.Pack("swap", singleSwap, funds, amountIn.ToBig(), deadline)
	if err != nil {
		panic(err)
	}
	targets = append(targets, wp.spender)
	payloads = append(payloads, txData)
	//return txData
	return targets, payloads
}

func (wp Weighted2TokenPair) ReceiveDirectly() bool {
	return false
}

func (wp Weighted2TokenPair) PrepareReceive(tokenAddress common.Address, amountIn *uint256.Int) ([]common.Address, [][]byte) {
	panic("Not implemented")
	return nil, nil
}

func (wp Weighted2TokenPair) HeavyComputation() bool {
	return true
}
