package flashbots

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jmjac/ethGoTest/helperClient"
)

type Signer struct {
	PrivateKey     *ecdsa.PrivateKey
	publicKey      *ecdsa.PublicKey
	Address        common.Address
	bundleExecutor common.Address
	wethAddress    common.Address
	ChainID        *big.Int
	client         *helperClient.Client
}

//TODO: Add separate key for pure v2 swaps

func NewSigner(chainID *big.Int, pk string, bundleExecutor, wethAddress common.Address, client *helperClient.Client) *Signer {
	s := new(Signer)
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
	}
	s.ChainID = chainID
	s.wethAddress = wethAddress
	s.client = client
	s.bundleExecutor = bundleExecutor
	s.PrivateKey = privateKey
	s.publicKey = &privateKey.PublicKey
	s.Address = crypto.PubkeyToAddress(*s.publicKey)
	return s
}

//Signs and sends the transactions
func (s Signer) SentTxToBundle(txData []byte, value *big.Int, nounce, gasLimit uint64, gasFeeCap, gasTipCap *big.Int) error {
	signedTx, err := s.SignTx(txData, value, nounce, gasLimit, gasFeeCap, gasTipCap)
	err = s.client.SendTransaction(context.Background(), signedTx)
	return err
}

//func (s Signer) SentAccessListTxToBundle(txData []byte, accessList types.AccessList, value *big.Int, nounce, gasLimit uint64, gasFeeCap, gasTipCap *big.Int) error {
//	signedTx, err := s.SignTxWithAccessList(txData, accessList, value, nounce, gasLimit, gasFeeCap, gasTipCap)
//	err = s.client.SendTransaction(context.Background(), signedTx)
//	return err
//}

func GetRawSignedTx(signedTx *types.Transaction) string {
	byts, _ := signedTx.MarshalBinary()
	return hexutil.Encode(byts)
	//	ts := types.Transactions{signedTx}
	//	b := new(bytes.Buffer)
	//	ts.EncodeIndex(0, b)
	//return "0x" + hex.EncodeToString(b.Bytes())
}

func (s Signer) SignTxToAddress(txData []byte, to common.Address, value *big.Int, nonce, gasLimit uint64, gasFeeCap, gasTipCap *big.Int) (*types.Transaction, error) {
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.ChainID,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap, // a.k.a. maxFeePerGas
		GasTipCap: gasTipCap, // a.k.a. maxPriorityFeePerGas
		Gas:       gasLimit,
		To:        &to,
		Value:     value,
		Data:      txData,
	})
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(s.ChainID), s.PrivateKey)
	return signedTx, err
}

func (s Signer) SentTxToAddress(txData []byte, to common.Address, value *big.Int, nounce, gasLimit uint64, gasFeeCap, gasTipCap *big.Int) error {
	signedTx, err := s.SignTxToAddress(txData, to, value, nounce, gasLimit, gasFeeCap, gasTipCap)
	err = s.client.SendTransaction(context.Background(), signedTx)
	return err
}

/*
Retruns signedTx
GasFeeCap - maxFeePerGas
GasTipCap - maxPriorityFeePerGas
*/
func (s Signer) SignTx(txData []byte, value *big.Int, nonce, gasLimit uint64, gasFeeCap, gasTipCap *big.Int) (*types.Transaction, error) {
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   s.ChainID,
		Nonce:     nonce,
		GasFeeCap: gasFeeCap, // a.k.a. maxFeePerGas
		GasTipCap: gasTipCap, // a.k.a. maxPriorityFeePerGas
		Gas:       gasLimit,
		To:        &s.bundleExecutor,
		Value:     value,
		Data:      txData,
	})
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(s.ChainID), s.PrivateKey)
	return signedTx, err
}

func (s Signer) SignTxWithAccessList(txData []byte, accessList types.AccessList, value *big.Int, nonce, gasLimit uint64, gasFeeCap, gasTipCap *big.Int) (*types.Transaction, error) {
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:    s.ChainID,
		Nonce:      nonce,
		GasFeeCap:  gasFeeCap, // a.k.a. maxFeePerGas
		GasTipCap:  gasTipCap, // a.k.a. maxPriorityFeePerGas
		Gas:        gasLimit,
		To:         &s.bundleExecutor,
		AccessList: accessList,
		Value:      value,
		Data:       txData,
	})
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(s.ChainID), s.PrivateKey)
	return signedTx, err
}

func (s Signer) SignTxWithGasPrice(txData []byte, value *big.Int, nonce, gasLimit uint64, gasPrice *big.Int) (*types.Transaction, error) {
	//tx := types.NewTx(&types.LegacyTx{
	//	Nonce:    nonce,
	//	GasPrice: gasPrice,
	//	Gas:      gasLimit,
	//	To:       &s.bundleExecutor,
	//	Value:    value,
	//	Data:     txData,
	//})
	tx := types.NewTransaction(nonce, s.bundleExecutor, value, gasLimit, gasPrice, txData)
	fmt.Printf("Setting gasPrice to %v\n", tx.GasPrice())
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(s.ChainID), s.PrivateKey)
	return signedTx, err
}

func (s Signer) PrepareOptsWithoutSend() (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(s.PrivateKey, s.ChainID)
	if err != nil {
		return nil, err
	}

	auth.NoSend = true
	auth.From = s.Address
	return auth, nil
}

func (s Signer) PrepareOpts(client bind.ContractBackend) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(s.PrivateKey, s.ChainID)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice
	return auth, err
}

func (s Signer) SimulateBundleTx(amountIn *big.Int, txData []byte, blockNum *big.Int) ([]byte, error) {
	msg := ethereum.CallMsg{}
	msg.Value = amountIn
	msg.From = s.Address
	msg.Data = txData
	msg.To = &s.bundleExecutor
	resp, err := s.client.CallContract(context.Background(), msg, blockNum)
	//s.client.PendingCallContract()
	return resp, err
}

func (s Signer) MoonbeamPrepareDirectArbTx(amountToFirstMarket *big.Int, coinbaseFee *big.Int, targets []common.Address, payloads [][]byte) ([]byte, error) {
	//fmt.Printf("Prepared tx with %v wei bribe\n", coinbaseFee)
	parsed, _ := abi.JSON(strings.NewReader(MoonbeamExecutorABI))
	//fmt.Println("Calling swapWithoutFlashLoan")
	txData, err := parsed.Pack("swapWithoutFlashLoan", amountToFirstMarket, coinbaseFee, targets, payloads)
	return txData, err
}

func (s Signer) EstimateGasForTxWithBlock(amountIn *big.Int, txData []byte, block string) (uint64, error) {
	msg := ethereum.CallMsg{}
	msg.Value = amountIn
	msg.From = s.Address
	msg.Data = txData
	msg.To = &s.bundleExecutor
	msg.GasPrice = big.NewInt(0)
	estimatedGas, err := s.client.EstimateGasWithBlock(context.Background(), msg, block)
	return estimatedGas, err
}

func (s Signer) EstimateGasForTx(amountIn *big.Int, txData []byte) (uint64, error) {
	msg := ethereum.CallMsg{}
	msg.Value = amountIn
	msg.From = s.Address
	msg.Data = txData
	msg.To = &s.bundleExecutor
	msg.GasPrice = big.NewInt(0)
	estimatedGas, err := s.client.EstimateGas(context.Background(), msg)
	return estimatedGas, err
}

func (s Signer) EstimateGasForTxToAddress(amountIn *big.Int, txData []byte, to common.Address) (uint64, error) {
	msg := ethereum.CallMsg{}
	msg.Value = amountIn
	msg.From = s.Address
	msg.Data = txData
	msg.To = &to
	msg.GasPrice = big.NewInt(0)
	estimatedGas, err := s.client.EstimateGas(context.Background(), msg)
	return estimatedGas, err
}
