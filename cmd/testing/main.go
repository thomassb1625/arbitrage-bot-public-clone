package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jmjac/ethGoTest/abis"
	ADDR "github.com/jmjac/ethGoTest/addr"
	"github.com/jmjac/ethGoTest/flashbots"
)

func main() {
	//ADDR.MainnetAddresses()
	//url := "https://mainnet.infura.io/v3/b21211a97c6e49cb92a084c041f30468"
	//client, _ := helperClient.New(url)

	//txData := withdrawlWETH(signer.Address, big.NewInt(0), *client.Client)

	//signer := flashbots.NewSigner(big.NewInt(1), "900f71941b926d9fe1fff18ece3661c8cfb745554556efa9e7f4d4c698a7943a", ADDR.BUNDLE_EXECUTOR, ADDR.WETH, client)
	fmt.Println(flashbots.GetBundleStats("0x28589b5712dd1c05118a633638380c075c75691495a4721c8c8b7c51b35bf1f5", fmt.Sprintf("0x%x", 16095195)))

}
func withdrawlWETH(to common.Address, amount *big.Int, client ethclient.Client) []byte {
	IERC20Abi, err := abi.JSON(strings.NewReader(abis.IERC20ABI))
	if err != nil {
		log.Fatal(err)
	}
	transferData, _ := IERC20Abi.Pack("transfer", to, amount)
	bundleAbi, _ := abi.JSON(strings.NewReader(flashbots.BundleExecutorABI))
	txData, _ := bundleAbi.Pack("call", ADDR.WETH, big.NewInt(0), transferData)
	return txData
}
