package flashbots

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

//TODO: Change URL to use for Goerli for now
//var flashbotURL = "https://relay.flashbots.net"

const (
	json_cont = "application/json"
	//flashbotURL     = "https://relay-goerli.flashbots.net"
	stats           = "flashbots_getUserStats"
	privateTx       = "eth_sendPrivateTransaction"
	sendBundle      = "eth_sendBundle"
	callBundle      = "eth_callBundle"
	getBundleStats  = "flashbots_getBundleStats"
	flashbotXHeader = "X-Flashbots-Signature"
	p               = "POST"
)

type MevRelay struct {
	url  string
	Name string
}

func NewMevRelay(url, name string) MevRelay {
	return MevRelay{url: url, Name: name}
}

func NewFlashbotsMainnet() MevRelay {
	flashbotURL := "https://relay.flashbots.net"
	m := MevRelay{url: flashbotURL, Name: "Flashbots"}
	//fmt.Printf("Using %v as flashbotUrl\n", flashbotURL)
	return m
}

//func FlashbotsGoerli() {
//	flashbotURL = "https://relay-goerli.flashbots.net"
//	fmt.Printf("Using %v as flashbotUrl\n", flashbotURL)
//}

//TODO: Change this
var (
	privateKey, err = crypto.HexToECDSA(
		"4a085c580a18b8fcd428d38aeb7dfd2edaf2764b3b5783e1b621f74a54008293",
		//"3f379b6f7fe75be780760f7717ea4dda52d9a95b1ad1689895bc1d9204e8f1f5",
	)
)

func flashbotHeader(signature []byte, privateKey *ecdsa.PrivateKey) string {
	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex() +
		":" + hexutil.Encode(signature)
}

func sign(payload []byte) []byte {
	hashedBody := crypto.Keccak256Hash(payload).Hex()
	sig, _ := crypto.Sign(accounts.TextHash([]byte(hashedBody)), privateKey)
	return sig
}

func addHeader(req *http.Request, signedHeader []byte) {
	req.Header.Add("content-type", json_cont)
	req.Header.Add("Accept", json_cont)
	req.Header.Add(flashbotXHeader, flashbotHeader(signedHeader, privateKey))
}

func prepareCall(params interface{}, method string) []byte {
	req := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  params,
	}
	payload, _ := json.Marshal(req)
	return payload
}

func (m MevRelay) SendPrivateTx(signedTx string, hexMaxBlockNumber string) error {
	mevHTTPClient := &http.Client{
		Timeout: time.Second * 3,
	}
	payload := prepareCall([]map[string]string{{"tx": signedTx, "maxBlockNumber": hexMaxBlockNumber}}, privateTx)
	//fmt.Println("Payload: ", string(payload))
	req, _ := http.NewRequest(p, m.url, bytes.NewBuffer(payload))
	addHeader(req, sign(payload))
	resp, err := mevHTTPClient.Do(req)
	if err != nil {
		return err
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Resp: ", string(res))
	return nil
}

type BundleHash struct {
	BundleHash string `json:"bundleHash"`
}

type SendBundleResp struct {
	Jsonrpc string     `json:"jsonrpc"`
	ID      int64      `json:"id"`
	Result  BundleHash `json:"result"`
}

func (m MevRelay) SendBundleSingleTx(signedTx string, hexMaxBlockNumber string, errorLog log.Logger) (string, error) {

	mevHTTPClient := &http.Client{
		Timeout: time.Second * 3,
	}
	payload := prepareCall([]map[string]interface{}{{"txs": []string{signedTx}, "blockNumber": hexMaxBlockNumber}}, sendBundle)

	//fmt.Println("Payload: ", string(payload))
	req, _ := http.NewRequest(p, m.url, bytes.NewBuffer(payload))
	addHeader(req, sign(payload))
	resp, err := mevHTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//fmt.Printf("Resp from %v: %v\n", m.Name, string(res))
	var bundleResp SendBundleResp
	err = json.Unmarshal(res, &bundleResp)
	if err != nil {
		errorLog.Printf("Failed bundle with payload: %v\nresponse %v\n", string(payload), string(res))
	}
	return bundleResp.Result.BundleHash, err
}

func (m MevRelay) GetBundleStats(bundleHash string, hexBlockNumber string) error {
	mevHTTPClient := &http.Client{
		Timeout: time.Second * 3,
	}
	payload := prepareCall([]map[string]string{{"bundleHash": bundleHash, "blockNumber": hexBlockNumber}}, getBundleStats)

	fmt.Println("GetBundleStats payload: ", string(payload))
	req, _ := http.NewRequest(p, m.url, bytes.NewBuffer(payload))
	addHeader(req, sign(payload))
	resp, err := mevHTTPClient.Do(req)
	if err != nil {
		return err
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("GetBundleStats resp: ", string(res))
	return nil
}

type CallBundleResp struct {
	Jsonrpc string               `json:"jsonrpc"`
	ID      int64                `json:"id"`
	Result  CallBundleRespResult `json:"result"`
}

type CallBundleRespResult struct {
	BundleGasPrice    string              `json:"bundleGasPrice"`
	BundleHash        string              `json:"bundleHash"`
	CoinbaseDiff      string              `json:"coinbaseDiff"`
	EthSentToCoinbase string              `json:"ethSentToCoinbase"`
	GasFees           string              `json:"gasFees"`
	Results           []CallResultElement `json:"results"`
	StateBlockNumber  int64               `json:"stateBlockNumber"`
	TotalGasUsed      int64               `json:"totalGasUsed"`
}

type CallResultElement struct {
	CoinbaseDiff      string `json:"coinbaseDiff"`
	Error             string `json:"error"`
	EthSentToCoinbase string `json:"ethSentToCoinbase"`
	FromAddress       string `json:"fromAddress"`
	GasFees           string `json:"gasFees"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           int64  `json:"gasUsed"`
	ToAddress         string `json:"toAddress"`
	TxHash            string `json:"txHash"`
}

func (m MevRelay) CallBundle(signedTx string, hexMaxBlockNumber string, errorLog log.Logger) (int64, error) {
	mevHTTPClient := &http.Client{
		Timeout: time.Second * 3,
	}
	payload := prepareCall([]map[string]interface{}{{"txs": []string{signedTx}, "blockNumber": hexMaxBlockNumber, "stateBlockNumber": "latest"}}, callBundle)

	//fmt.Println("CallBundle payload: ", string(payload))
	req, _ := http.NewRequest(p, m.url, bytes.NewBuffer(payload))
	addHeader(req, sign(payload))
	resp, err := mevHTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var callBundleResp CallBundleResp
	err = json.Unmarshal(res, &callBundleResp)
	if err != nil {
		errorLog.Printf("Failed bundle with payload: %v\nresponse %v\n", string(payload), string(res))
	}
	if callBundleResp.Result.TotalGasUsed == 0 {
		return 0, errors.New(fmt.Sprintf("Got 0 gas used with:\n\tpayload: %v\n\tresp: %v\n", string(payload), string(res)))
	}
	return callBundleResp.Result.TotalGasUsed, nil
}

func (m MevRelay) GetUserStats(blockNum *big.Int) string {
	mevHTTPClient := &http.Client{
		Timeout: time.Second * 3,
	}
	currentBlock := blockNum
	params := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  stats,
		"params": []interface{}{
			fmt.Sprintf("0x%x", currentBlock.Uint64()),
		},
	}
	payload, _ := json.Marshal(params)
	req, _ := http.NewRequest(p, m.url, bytes.NewBuffer(payload))
	headerReady, _ := crypto.Sign(
		accounts.TextHash([]byte(hexutil.Encode(crypto.Keccak256(payload)))),
		privateKey,
	)
	req.Header.Add("content-type", json_cont)
	req.Header.Add("Accept", json_cont)
	req.Header.Add(flashbotXHeader, flashbotHeader(headerReady, privateKey))
	resp, _ := mevHTTPClient.Do(req)
	res, _ := ioutil.ReadAll(resp.Body)
	return string(res)
}
