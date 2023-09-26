# Curve Finance

# File Overview:

# Main Files:
1. pair.go - Provides all the v2 methods implemented in the older version of the bot. This files also contains the main Pair struct that will hold pool data.
2. market.go - Retrieves the necessary data and reserves. Built to accomodate 3pool and other basepools.
3. 3pool.go - Working swap logic for the curve 3pool.
4. basePool.go - Will hold the implementation for curve base pool swaps.

# Testing Files:
1. 3poolTest.go - Tests 3pool swap functions.
2. basePoolTest.go - In progress.
3. poolStats.go - Returns information about the number of different implementation types.

# Abi Files:
1. factory.go - basepool data.
2. registry.go - 3pool data.
3. swap.go - Mainnet swaps.

## Data Structures

File: pair.go

type Pair struct {
	address  common.Address
	tokens   []common.Address
	reserves []*uint256.Int
	protocol string

	A                 *uint256.Int
	rates             []*uint256.Int
	LENDING_PRECISION *uint256.Int
	PRECISION         *uint256.Int
	FEE_DENOMINATOR   *uint256.Int
	fee               *uint256.Int
	admin_fee         *uint256.Int
	N_COINS           *uint256.Int
	implemention      common.Address
}

Each curve pair is structured like this. The top section is build to mimic our Univ2 implementation. The major difference being that tokens and reserves supports a list of tokens. This is due to the fact that the swap math requires all the tokens in the pool to make the calculations.

# Integration Notes

## Getting Market Data

File: market.go

func GetCurveMarkets(client bind.ContractBackend, CURVE_ADDRESS_PROVIDER, CURVE_REGISTRY, CURVE_FACTORY common.Address, protocol string) []Pair 

Desc: This function will return a list of structs of type Pair that holds all the data from above. This will initialize the reserves to zero and call update reserves at the end of the function. You can use this function to retrieve all the pairs and the necessary data for swaps.

## Updating Reserves

File: market.go

func UpdateReserves(client bind.ContractBackend, CURVE_REGISTRY, CURVE_FACTORY common.Address, marketPairs []Pair, protocol string) []Pair

Desc: This function returns a list of structs with all the updated reserves. Functionality has not been included to sync the reserve amounts with a certain block.

## Making Swaps

File: pair.go

func (pair Pair) GetTokensOut(tokenIn common.Address, tokenOut common.Address, amountIn *uint256.Int) *uint256.Int 

Desc: The GetTokensOut Function can be used to calculate swaps on a given Pair. You will need to pass in the tokenIn and tokenOut as pair.tokens[index] to make sure the function has access to all the data stored in the struct.

## Integrating After Additional pools are added

File: basePool.go

Desc: This file will be used to house the swap implementation for the next set of pools. I have already adjusted the market.go and pair.go files for accomodating the new types as they may need different data. All the functions mentioned above will still work the same, and I will update the README.md once progress has been made.