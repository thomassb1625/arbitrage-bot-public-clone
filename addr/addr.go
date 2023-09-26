package ADDR

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
)

var WETH common.Address
var WBNB common.Address
var WGLMR common.Address
var DAI common.Address
var USDC common.Address
var WBTC common.Address
var RAI common.Address
var UNIv2_FACTORY common.Address
var UNIv3_FACTORY common.Address
var SUSHI_FACTORY common.Address
var SHIBA_FACTORY common.Address
var ORION_FACTORY common.Address
var SAKE_FACTORY common.Address
var DX_FACTORY common.Address
var FRAX_FACTORY common.Address
var PANCAKESWAP_FACTORY common.Address
var CRYPTDEFI_FACTORY common.Address
var UNISWAPv2_LOOKUP common.Address
var UNISWAPV3_ROUTER common.Address
var NOMISWAP_FACTORY common.Address
var APESWAP_FACTORY common.Address
var RADIOSHACK_FACTORY common.Address
var LUASWAP_FACTORY common.Address
var SAITASWAP_FACTORY common.Address

var SOLARFLARE_FACTORY common.Address // Moonbeam
var STELLASWAP_FACTORY common.Address // Moonbeam
var BEAMSWAP_FACTORY common.Address   // Moonbeam
var PADSWAP_FACTORY common.Address    // Moonbeam

var BALANCERv1_FACTORY common.Address
var BALANCERv2_WEIGHTED_FACTORY common.Address
var BALANCERv2_WEIGHTED2TOKENS_FACTORY common.Address
var BALANCERv2_STABLE_FACTORY common.Address
var BALANCERv2_VAULT common.Address
var BUNDLE_EXECUTOR common.Address
var CURVE_ADDRESS_PROVIDER common.Address
var CURVE_REGISTRY common.Address
var CURVE_FACTORY common.Address

var SUBGRAPH_URL string

func MainnetAddresses(logger log.Logger) {
	logger.Println("================= MAINNET ======================")
	WETH = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	DAI = common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")
	USDC = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	RAI = common.HexToAddress("0x03ab458634910aad20ef5f1c8ee96f1d6ac54919")
	WBTC = common.HexToAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")
	UNIv2_FACTORY = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	UNIv3_FACTORY = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	SHIBA_FACTORY = common.HexToAddress("0x115934131916c8b277dd010ee02de363c09d037c")
	ORION_FACTORY = common.HexToAddress("0x5FA0060FcfEa35B31F7A5f6025F0fF399b98Edf1")
	SUSHI_FACTORY = common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac")
	SAKE_FACTORY = common.HexToAddress("0x75e48C954594d64ef9613AeEF97Ad85370F13807")
	DX_FACTORY = common.HexToAddress("0xd34971BaB6E5E356fd250715F5dE0492BB070452")
	FRAX_FACTORY = common.HexToAddress("0xB076b06F669e682609fb4a8C6646D2619717Be4b")
	CRYPTDEFI_FACTORY = common.HexToAddress("0x9DEB29c9a4c7A88a3C0257393b7f3335338D9A9D")
	PANCAKESWAP_FACTORY = common.HexToAddress("0x1097053Fd2ea711dad45caCcc45EfF7548fCB362")
	RADIOSHACK_FACTORY = common.HexToAddress("0x91fAe1bc94A9793708fbc66aDcb59087C46dEe10")
	LUASWAP_FACTORY = common.HexToAddress("0x0388C1E0f210AbAe597B7DE712B9510C6C36C857")
	SAITASWAP_FACTORY = common.HexToAddress("0x35113a300ca0D7621374890ABFEAC30E88f214b1")
	UNISWAPv2_LOOKUP = common.HexToAddress("0x5EF1009b9FCD4fec3094a5564047e190D72Bd511")
	UNISWAPV3_ROUTER = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
	//BALANCERv1_FACTORY = common.HexToAddress("0x9424B1412450D0f8Fc2255FAf6046b98213B76Bd")
	BALANCERv2_WEIGHTED_FACTORY = common.HexToAddress("0x8E9aa87E45e92bad84D5F8DD1bff34Fb92637dE9")
	BALANCERv2_WEIGHTED2TOKENS_FACTORY = common.HexToAddress("0xA5bf2ddF098bb0Ef6d120C98217dD6B141c74EE0")
	BALANCERv2_STABLE_FACTORY = common.HexToAddress("0x8df6EfEc5547e31B0eb7d1291B511FF8a2bf987c")
	BALANCERv2_VAULT = common.HexToAddress("0xBA12222222228d8Ba445958a75a0704d566BF2C8")
	CURVE_ADDRESS_PROVIDER = common.HexToAddress("0x0000000022D53366457F9d5E68Ec105046FC4383")
	CURVE_REGISTRY = common.HexToAddress("0x90E00ACe148ca3b23Ac1bC8C240C2a7Dd9c2d7f5")
	CURVE_FACTORY = common.HexToAddress("0xB9fC157394Af804a3578134A6585C0dc9cc990d4")
	BUNDLE_EXECUTOR = common.HexToAddress("0x0000008200bE7f00920000d500eCd5F6fafb8386")
	SUBGRAPH_URL = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"
	PrintAddresses(logger)
}

// Goerli Addresses initialize all addresses for Goerli testnet
func GoerliAddresses(logger log.Logger) {
	WETH = common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6")
	DAI = common.HexToAddress("0x11fE4B6AE13d2a6055C8D9cF65c55bac32B5d844")
	USDC = common.HexToAddress("")
	UNIv2_FACTORY = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	UNIv3_FACTORY = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	SUSHI_FACTORY = common.HexToAddress("0xc35DADB65012eC5796536bD9864eD8773aBc74C4")
	SHIBA_FACTORY = common.HexToAddress("")
	ORION_FACTORY = common.HexToAddress("")
	SAKE_FACTORY = common.HexToAddress("")
	DX_FACTORY = common.HexToAddress("")
	UNISWAPV3_ROUTER = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
	FRAX_FACTORY = common.HexToAddress("")
	CRYPTDEFI_FACTORY = common.HexToAddress("")
	PANCAKESWAP_FACTORY = common.HexToAddress("")
	RADIOSHACK_FACTORY = common.HexToAddress("")
	UNISWAPv2_LOOKUP = common.HexToAddress("0x8B13bc595309C345A358A3975330FB530434a54C")
	BALANCERv2_WEIGHTED_FACTORY = common.HexToAddress("0x8e9aa87e45e92bad84d5f8dd1bff34fb92637de9")
	BALANCERv2_WEIGHTED2TOKENS_FACTORY = common.HexToAddress("0xA5bf2ddF098bb0Ef6d120C98217dD6B141c74EE0")
	BALANCERv2_STABLE_FACTORY = common.HexToAddress("0x44afeb87c871D8fEA9398a026DeA2BD3A13F5769")
	BALANCERv2_VAULT = common.HexToAddress("0xBA12222222228d8Ba445958a75a0704d566BF2C8")
	//BUNDLE_EXECUTOR = common.HexToAddress("0x098Ff4f36441C9A16E707C1e02dd1DA485f164db") //old
	// BUNDLE_EXECUTOR = common.HexToAddress("0x92E9A40aEEB5bF032661E5f156eF12f430203e85") //testing this
	//BUNDLE_EXECUTOR = common.HexToAddress("0xb2521D93CF36e374789f04d283C912E497FE0f8b") // 200 runs
	//BUNDLE_EXECUTOR_MAX_RUNS = common.HexToAddress("0x060e13EC63550d63d23C0524C104827F762059ea") // MAX RUNS
	//BUNDLE_EXECUTOR_MAX_RUNS = common.HexToAddress("0x977C80a450950a5A52eb0b72430693fadd530042") // Removed bribe payment
	BUNDLE_EXECUTOR = common.HexToAddress("0x977C80a450950a5A52eb0b72430693fadd530042") // Removed bribe payment and flashloanGuard

	SUBGRAPH_URL = "https://api.thegraph.com/subgraphs/name/liqwiz/uniswap-v3-goerli"
	PrintAddresses(logger)
}

func BnbAddresses(logger log.Logger) {
	logger.Println("================= BNB ======================")
	WETH = common.HexToAddress("0x4DB5a66E937A9F4473fA95b1cAF1d1E1D62E29EA")
	WBNB = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
	PANCAKESWAP_FACTORY = common.HexToAddress("0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73")
	NOMISWAP_FACTORY = common.HexToAddress("0xd6715A8be3944ec72738F0BFDC739d48C3c29349")
	APESWAP_FACTORY = common.HexToAddress("0x0841BD0B734E4F5853f0dD8d7Ea041c241fb0Da6")

	BUNDLE_EXECUTOR = common.HexToAddress("") // Bundle ExecutorV2
	PrintAddresses(logger)
}

func Moonbeam(logger log.Logger) {
	WETH = common.HexToAddress("0x30D2a9F5FDf90ACe8c17952cbb4eE48a55D916A7")
	WGLMR = common.HexToAddress("0xAcc15dC74880C9944775448304B263D191c6077F")
	DAI = common.HexToAddress("0x765277EebeCA2e31912C9946eAe1021199B39C61")
	USDC = common.HexToAddress("")
	UNISWAPv2_LOOKUP = common.HexToAddress("0xF6330429e82CbEf117eE46cBEC39529a411e66e4")
	SOLARFLARE_FACTORY = common.HexToAddress("0x19B85ae92947E0725d5265fFB3389e7E4F191FDa")
	STELLASWAP_FACTORY = common.HexToAddress("0x68A384D826D3678f78BB9FB1533c7E9577dACc0E")
	PADSWAP_FACTORY = common.HexToAddress("0x663a07a2648296f1A3C02EE86A126fE1407888E5")
	BEAMSWAP_FACTORY = common.HexToAddress("0x985BcA32293A7A496300a48081947321177a86FD")
	BUNDLE_EXECUTOR = common.HexToAddress("0x2eb0f72d3ba6c6d5cd2a1ebae682034caa2ef0cf")
}

func PrintAddresses(logger log.Logger) {
	logger.Println("================= TOKENS ======================")
	logger.Println("WETH address = \t", WETH)
	logger.Println("DAI address = \t", DAI)
	logger.Println("================= FACTORIES ======================")
	logger.Println("================= UNIV2 and clones ======================")
	logger.Println("UNIv2_FACTORY address = \t", UNIv2_FACTORY)
	logger.Println("UNIv3_FACTORY address = \t", UNIv3_FACTORY)
	logger.Println("SUSHI_FACTORY address = \t", SUSHI_FACTORY)
	logger.Println("SHIBA_FACTORY address = \t", SHIBA_FACTORY)
	logger.Println("ORION_FACTORY address = \t", ORION_FACTORY)
	logger.Println("SAKE_FACTORY address = \t", ORION_FACTORY)
	logger.Println("SAKE_FACTORY address = \t", SAKE_FACTORY)
	logger.Println("DX_FACTORY address = \t", DX_FACTORY)
	logger.Println("FRAX_FACTORY address = \t", FRAX_FACTORY)
	logger.Println("CRYPTDEFI_FACTORY address = \t", CRYPTDEFI_FACTORY)
	logger.Println("PANCAKESWAP_FACTORY address = \t", PANCAKESWAP_FACTORY)
	logger.Println("================= OTHER ======================")
	logger.Println("BALANCERv1_FACTORY address = \t", BALANCERv1_FACTORY)
	logger.Println("BALANCERv2_VAULT address = \t", BALANCERv2_VAULT)
	logger.Println("BALANCERv2_STABLE_FACTORY address = \t", BALANCERv2_STABLE_FACTORY)
	logger.Println("BALANCERv2_WEIGHTED_FACTORY address = \t", BALANCERv2_WEIGHTED_FACTORY)
	logger.Println("BALANCERv2_WEIGHTED2TOKENS_FACTORY address = \t", BALANCERv2_WEIGHTED2TOKENS_FACTORY)
	logger.Println("CURVE_ADDRESS_PROVIDER address = \t", CURVE_ADDRESS_PROVIDER)
	logger.Println("CURVE_ADDRESS_PROVIDER address = \t", CURVE_REGISTRY)
	logger.Println("================= UTILITIES ======================")
	logger.Println("UNISWAPv2_LOOKUP address = \t", UNISWAPv2_LOOKUP)
	logger.Println("BundleExecutor address = \t", BUNDLE_EXECUTOR)
}
