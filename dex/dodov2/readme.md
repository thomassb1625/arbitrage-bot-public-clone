DODO guide

Relevant files: pool.go and market.go

pool.go: pulls addresses, updates reserves returning a pair struct

market.go: calculates swap outputs

How to implement:
1. Pull the addresses for all the pools using the following function: 
    getDVMPoolAddresses(_client string) []*common.Address
    - getDSP is in progress

2. Once you have an array of pointers to common.Address, pass these into:
    updateReserves(_client string, pools []*common.Address, isDVM bool) []pair

    - set'isDVM' to true for now, as only DVM pools work currently so you should only be passing in DVM pool addresses
    - This function returns a pair struct which holds all the data necessary for swaps, but note that a baseToken 
        and quoteToken address are specified as the swap math is different in each direction

3. Finally, the following two functions are used to calculate swaps:
    querySellBase(p *pair, payBaseAmount *uint256.Int) *uint256.Int
    querySellQuote(p *pair, payQuoteAmount *uint256.Int) *uint256.Int

    There are two seperate functions for the two directions each pool can go in. For querySellBase, all it is calculating is the quoteToken amount that will be output if we pay baseToken into the pool, and vice verse for querySellQuote. Just make sure that we know we are inputting baseToken into querySellBase and quoteToken into querySellQuote or else the outputs will be totally wrong

Testing notes:
The testiing file I made is pretty comprehensive, but the following functions in market_test.go will probably be most useful for you
    - TestDVMMarket - tests all pools swapping in each direction for DVM, just change the input and CLIENT variables at the top of   the file if you want to play around with different cases
    - TestDSPMarket - same but with the stable pools if you want to checkout what is erroring
    - TestPool - function I made to help with debugging specific pools that are erroring, just flip 'testSellBase := true' to false if you want to test selling the quote instead, and change the testPoolAddr at the top to try it with different pools

DSP Update - over half the pools are functioning properly but some weird errors are occuring for the others, debugging them is proving harder but I'll keep at it and they should be wrapped up soon
