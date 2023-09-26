package uniswapV3

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	ADDR "github.com/jmjac/ethGoTest/addr"
)

/*
Subgraph query
{
  pool(id: "0x87986ae1e99f99da1f955d16930dc8914ffbed56", block: {number: 15713777}){
    id
    ticks{
      tickIdx
      liquidityNet
      liquidityGross
    }
  }
}
*/

func UnmarshalResponse(data []byte) (graphResponse, error) {
	var r graphResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *graphResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type graphResponse struct {
	Data graphData `json:"data"`
}

type graphData struct {
	Pool graphPool `json:"pool"`
}

type graphPool struct {
	Liquidity string      `json:"liquidity"`
	Tick      string      `json:"tick"`
	SqrtPrice string      `json:"sqrtPrice"`
	Ticks     []graphTick `json:"ticks"`
}

type graphTick struct {
	TickIdx        string `json:"tickIdx"`
	LiquidityNet   string `json:"liquidityNet"`
	LiquidityGross string `json:"liquidityGross"`
}

func GetStateFromSubgraph(address common.Address, block *big.Int) (int, *uint256.Int, *uint256.Int, map[int]Tick, error) {
	liquidity := new(uint256.Int)
	var tick int64
	ticks := make(map[int]Tick)
	sqrtPriceX96 := new(uint256.Int)

	client := &http.Client{Timeout: time.Second * 10}

	first := 1000
	skip := 0
	for {
		jsonData := map[string]string{
			"query": fmt.Sprintf(`{
  		pool(id: "%v", block: {number: %v}){
    		liquidity
				sqrtPrice
				tick
				ticks(first: %v, skip: %v){
      		tickIdx
      		liquidityNet
      		liquidityGross
    		}
  		}
		}`, strings.ToLower(address.Hex()), block, first, skip),
		}
		jsonValue, _ := json.Marshal(jsonData)
		request, err := http.NewRequest("POST", ADDR.SUBGRAPH_URL, bytes.NewBuffer(jsonValue))
		response, err := client.Do(request)

		if err != nil {
			return 0, nil, nil, nil, err
		}

		data, err := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		response.Body.Close()
		//TODO: There are some errors wtih the data returned from subgraph
		if string(data) == `{"data":{"pool":null}}` {
			//fmt.Println(string(data))
			return 0, nil, nil, nil, errors.New("Got response as pool:null")
		}
		v, _ := UnmarshalResponse(data)
		poolData := v.Data.Pool

		liqBig, _ := new(big.Int).SetString(poolData.Liquidity, 10)
		if liqBig == nil {
			errMsg := fmt.Sprintf("Got nil liquidity. Resp was %v", string(data))
			return 0, nil, nil, nil, errors.New(errMsg)
		}
		liquidity, _ = uint256.FromBig(liqBig)
		tick, _ = strconv.ParseInt(poolData.Tick, 10, 64)
		bigSqrt, _ := new(big.Int).SetString(poolData.SqrtPrice, 10)
		sqrtPriceX96, _ = uint256.FromBig(bigSqrt)

		for _, t := range poolData.Ticks {
			index, _ := strconv.ParseInt(t.TickIdx, 10, 64)
			liquidityGross, _ := new(big.Int).SetString(t.LiquidityGross, 10)
			uintLiquidtyGross, _ := uint256.FromBig(liquidityGross)
			liquidityNet, _ := new(big.Int).SetString(t.LiquidityNet, 10)
			ticks[int(index)] = Tick{Initialized: true, LiquidityGross: *uintLiquidtyGross, LiquidityNet: *liquidityNet}
		}
		//fmt.Println(len(ticks))

		if len(poolData.Ticks) < 1000 {
			break
		} else {
			skip += first
		}
	}
	return int(tick), liquidity, sqrtPriceX96, ticks, nil
}
