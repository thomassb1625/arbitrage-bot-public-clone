package arbitrage

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
)

func PrintPath(token common.Address, arb *Arbitrage, w io.Writer) {
	in := arb.AmountIn
	for _, p := range arb.Path {
		outToken := p.GetOtherToken(token).Addr
		out := p.GetTokensOut(token, outToken, in)
		fmt.Fprintf(w, "%v %v -> (%s) %v %v\n", token, in, p.Protocol(), out, outToken)
		token = outToken
		in = out
	}
	//fmt.Printf("Best profit found %v starting with %v ETH\n", FormatUintToFloat(bestFound.profit, 1/1e18), FormatUintToFloat(bestFound.amountIn, 1/1e18))
}

func FormatUintToFloat(val *uint256.Int, decimals float64) *big.Float {
	f := new(big.Float).SetInt(val.ToBig())
	den := new(big.Float).SetFloat64(decimals)
	return f.Mul(f, den)
}

func LoadDecimalsFromJson(filename string) map[common.Address]uint64 {
	tokenDecimals := make(map[common.Address]uint64)
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &tokenDecimals)
	return tokenDecimals
}

func SaveTokenDecimals(decimals map[common.Address]uint64, filename string) {
	data, _ := json.Marshal(decimals)
	ioutil.WriteFile(filename, data, 0666)
}
