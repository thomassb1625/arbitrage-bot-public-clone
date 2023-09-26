package balancerV2

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/jmjac/ethGoTest/abis"
	"github.com/jmjac/ethGoTest/dex"
)

//Pool can have at most 8 tokens
const TOKEN_LIMIT = 8

var MAX_IN_RATIO = uint256.NewInt(3e17)

//Check if the pool is paused
type WeightedPair struct {
	address common.Address
	poolID  [32]byte
	//Should already be normalized
	scallingFactor [TOKEN_LIMIT]*uint256.Int
	weights        [TOKEN_LIMIT]*uint256.Int
	balances       [TOKEN_LIMIT]*uint256.Int // Needed every block
	tokens         [TOKEN_LIMIT]dex.Token
	fee            *uint256.Int
	protocol       string
}

func parseWeightedLogsForAddresses(client bind.ContractBackend, addrWeighted common.Address, creationBlock *big.Int) []*WeightedPair {
	pools := make([]*WeightedPair, 0)

	query := ethereum.FilterQuery{
		FromBlock: creationBlock,
		Addresses: []common.Address{
			addrWeighted,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		panic(err)
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
			panic(err)
		}
		state, err := pool.GetPausedState(nil)
		if err != nil {
			panic(err)
		}
		if state.Paused {
			continue
		}
		pair := WeightedPair{address: poolAddr, fee: uint256.NewInt(0), protocol: "balancerV2Weighted"}
		pools = append(pools, &pair)
	}
	return pools
}

func GetWeightedMarkets(client bind.ContractBackend, addrVault, addrWeighted common.Address, creationBlock *big.Int, BLACKLIST map[common.Address]bool, tokensMap map[common.Address]dex.Token) []*WeightedPair {
	now := time.Now()
	dex.Logger.Println("\n============== Setting up BalancerV2 Weighted Pools ===================")
	//TODO: USE THE BLACKLIST
	pairs := parseWeightedLogsForAddresses(client, addrWeighted, creationBlock)
	vault, _ := NewVaultCaller(addrVault, client)
	for index, p := range pairs {
		pool, err := NewBBasedWeightedPool(p.address, client)

		if err != nil {
			panic(err)
		}

		weights, _ := pool.GetNormalizedWeights(nil)
		for i, w := range weights {
			v, _ := uint256.FromBig(w)
			pairs[index].weights[i] = v
		}

		id, _ := pool.GetPoolId(nil)
		pairs[index].poolID = id
		fee, _ := pool.GetSwapFeePercentage(nil)
		feeUint, _ := uint256.FromBig(fee)
		pairs[index].fee = feeUint

		poolTokens, err := vault.GetPoolTokens(nil, id)
		if err != nil {
			panic(err)
		}
		for i := range poolTokens.Balances {
			balance, _ := uint256.FromBig(poolTokens.Balances[i])
			pairs[index].balances[i] = balance

			ti, ok := tokensMap[poolTokens.Tokens[i]]
			if !ok {
				ti = dex.NewNotTaxed(poolTokens.Tokens[i])
			}

			pairs[index].tokens[i] = ti

			token, _ := abis.NewIERC20Caller(poolTokens.Tokens[i], client)
			decimals, _ := token.Decimals(nil)
			decimalDifference := uint64(18 - decimals)
			scallingFactor := new(uint256.Int).Exp(uint256.NewInt(10), uint256.NewInt(decimalDifference))
			scallingFactor.Mul(ONE_E18, scallingFactor)
			pairs[index].scallingFactor[i] = scallingFactor
		}

		//		fmt.Println("==================================================")
		//		fmt.Printf("%+v\n", pairs[index].balances)
		//		fmt.Printf("%+v\n", pairs[index].weights)
		//		fmt.Printf("%+v\n", pairs[index].scallingFactor)
		//		fmt.Printf("%+v\n", pairs[index].Address)
		//		fmt.Printf("%+v\n", pairs[index].Tokens)
		//Get values with solidity contract
		//Valut is where you execute the swaps
	}
	dex.Logger.Printf("Loaded %v pools from BalancerV2 Weighted\n", len(pairs))
	dex.Logger.Printf("Took: %v\n", time.Since(now))
	return pairs
}

func UpdateReservesWeighted(client bind.ContractBackend, addrVault common.Address, pairs []*WeightedPair) {
	vault, _ := NewVaultCaller(addrVault, client)
	var wg sync.WaitGroup
	for index := 0; index < len(pairs); index++ {
		wg.Add(1)
		go func(index int) {
			poolTokens, _ := vault.GetPoolTokens(nil, pairs[index].poolID)
			for i := range poolTokens.Balances {
				balance, _ := uint256.FromBig(poolTokens.Balances[i])
				pairs[index].balances[i] = balance
			}
			wg.Done()
		}(index)
	}
	wg.Wait()
}

/* TODO: Add this later as a safety check
uint256 internal constant _MAX_IN_RATIO = 0.3e18;
uint256 internal constant _MAX_OUT_RATIO = 0.3e18;
*/
func (wp WeightedPair) GetTokensOut(tokenIn, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int {
	//sci= scallinfFactory for tokenIn
	panic("Unimplemented")
	bi, wi, sci, bo, wo, sco := wp.tokensToBalancesAndWeights(tokenIn, tokenOut)
	//fmt.Printf("%v, %v, %v, %v, %v, %v\n", bi, wi, sci, bo, wo, sco)

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

	return downscaleDown(wp.getAmountOut(feeAmount, bi, wi, bo, wo), sco)
}

func (wp WeightedPair) getAmountOut(ai, bi, wi *uint256.Int, bo, wo *uint256.Int) *uint256.Int {
	denominator := ai.Add(ai, bi)
	base := divUp(bi, denominator)
	exponent := divDown(wi, wo)
	power := powUp(base, exponent)
	return mulDown(bo, complement(power))
}
func (wp WeightedPair) tokensToBalancesAndWeights(t0, t1 common.Address) (*uint256.Int, *uint256.Int, *uint256.Int, *uint256.Int, *uint256.Int, *uint256.Int) {
	//MAYBE A BETTER IDEA IS TO CREATE A POOL FOR EVERY PAIR. Would be faster than this but use more memory
	var weightT0, balanceT0, scallingT0, weightT1, balanceT1, scallingT1 *uint256.Int
	for i := 0; i < TOKEN_LIMIT; i++ {
		if wp.tokens[i].Addr == t0 {
			balanceT0 = wp.balances[i]
			weightT0 = wp.weights[i]
			scallingT0 = wp.scallingFactor[i]
		} else if wp.tokens[i].Addr == t1 {
			balanceT1 = wp.balances[i]
			weightT1 = wp.weights[i]
			scallingT1 = wp.scallingFactor[i]
		}
	}
	if balanceT0 == nil || balanceT1 == nil {
		panic("Wrong tokenAddress passed to the pool")
	}
	return balanceT0, weightT0, scallingT0, balanceT1, weightT1, scallingT1
}

func (wp *WeightedPair) ConvertTo2TokenWeightedPair(vault common.Address) []*Weighted2TokenPair {
	pairs := make([]*Weighted2TokenPair, 0)
	numTokens := 0
	for i := 0; i < TOKEN_LIMIT; i++ {
		if wp.weights[i] != nil {
			numTokens++
		} else {
			break
		}
	}

	for i := 0; i < numTokens; i++ {
		for j := i + 1; j < numTokens; j++ {
			p := Weighted2TokenPair{address: wp.address, spender: vault}
			p.b0 = wp.balances[i]
			p.w0 = wp.weights[i]
			p.b1 = wp.balances[j]
			p.w1 = wp.weights[j]
			p.sc0 = wp.scallingFactor[i]
			p.sc1 = wp.scallingFactor[j]
			p.token0 = wp.tokens[i]
			p.token1 = wp.tokens[j]
			p.fee = wp.fee
			p.PoolId = wp.poolID
			p.protocol = "balancerV2"
			pairs = append(pairs, &p)
		}
	}

	return pairs
}
