/*
1. packing of call input:
- input data length:// we won't be using that for now

- call ether value: //We can use 40bits of call data so 5 bytes - 0.0022$ with 2000$ ETH
    bytes[0]:
        000 = ?
        001 = swap 2 univ2
        010 = swap 3 univ2
        011 = swap 4 univ2
        100 = swap 5 univ2
        101, 110, 111 and 000 are free to use for now
    If pure uniV2 swaps:
    bytes[1]:
        8  - 001 - First swap
        16 - 010 - Second swap
        32 - 100 - Third swap
    bytes[2]:
        64  - 001 - Fourth Swap
        128 - 010 - Fifth Swap
        256 - 100 - 
- swap2 data:
        20b pair0 address
        9b  pair0 input amount up to 4722 ETH - used for eth transfer
        20b pair1 address - 320gas
        14b pair1 input amount == pair0 output amount (up to 10^33 tokens)
        9b pair1 output amount == exepcted ETH output
        --
        72b input data total 
- swap3 data:
        20b pair0 address
        9b  pair0 input amount up to 4722 ETH - used for eth transfer
        20b pair1 address - 320gas
        14b pair1 input amount == pair0 output amount (up to 10^33 tokens)
        20b pair2 address - 320gas
        14b pair2 input amount == pair1 output amount 
        9b pair1 output amount == exepcted ETH output
        --
        106b input data total 
- swap4 data:
        20b pair0 address
        9b  pair0 input amount up to 4722 ETH - used for eth transfer
        20b pair1 address - 320gas
        14b pair1 input amount == pair0 output amount (up to 10^33 tokens)
        20b pair2 address - 320gas
        14b pair2 input amount == pair1 output amount 
        20b pair3 address - 320gas
        14b pair3 input amount == pair2 output amount
        9b pair1 output amount == exepcted ETH output
        --
        106b input data total 
2. restictions to take into account in caller:
  - 9b max input amount for pair0 (buy)
  - 14b max swap amounts other then WETH
*/

/*
    TODO:
        - Add ability to call arbitrary function by the owner
        - Add suicide if we want reinit
        - Balancer/Curve ?
*/

/*
    IDEAS:
        - Gas left can be used as a function selector given that it doesn't affect if we are included in the block - can get free 2 bytes
*/

//0xfa461e33 - univ3callback

//NOTE: CHANGE ALL ADDR TO WETH FOR MAINNET BEFORE DEPLOYMENT
object "UniswapOptimized" {
    code {
        codecopy(0, dataoffset("runtime"), datasize("runtime"))
        //May need to be changed if deployed with create2
        //setimmutable(0, "owner", caller())
        return(0, datasize("runtime"))
    }
    object "runtime" {
        code {
            //Make this hardcoded so we can deploy with Create2
            if eq(0xCAC9325A179B4dbB2FfbF76BD3524EEF966fd81A, caller()) {
                // Balance checks are not needed for pure v2 swaps I think
                // //Move the pointer to leave space for balance before and after
                // mstore(0x40,add(0x80,0x40))
                // balanceOf(0x80) // Before the call

                switch and(callvalue(), 0x7)
                // case 0 /* withdraw */ {
                //     widthdrawWeth() 
                // }
                case 1 /* univ2 two swaps */ { //Works
                    let pair0Addr := shr(96, calldataload(0)) // extract 20b addr
                    let pair1Addr := shr(96, calldataload(29)) // extract 20b addr
                    transferWeth(pair0Addr, shr(184, calldataload(20)))
                    uniswap(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49))) // 9b eth value
                    uniswap(pair1Addr, address(), and(callvalue(), 0x10), shr(184, calldataload(63)))
                }
                case 2 /* univ2 three swaps*/ { //Works
                    let pair0Addr := shr(96, calldataload(0)) // extract 20b addr
                    let pair1Addr := shr(96, calldataload(29)) // extract 20b addr
                    let pair2Addr := shr(96, calldataload(63)) // extract 20b addr
                    transferWeth(pair0Addr, shr(184, calldataload(20)))
                    uniswap(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49))) 
                    uniswap(pair1Addr, pair2Addr, and(callvalue(), 0x10), shr(144, calldataload(83)))
                    uniswap(pair2Addr, address(), and(callvalue(), 0x20), shr(184, calldataload(97)))
                }
                case 3 /* univ2 four swaps*/{ // Works
                    let pair0Addr := shr(96, calldataload(0)) // extract 20b addr
                    let pair1Addr := shr(96, calldataload(29)) // extract 20b addr
                    let pair2Addr := shr(96, calldataload(63)) // extract 20b addr
                    let pair3Addr := shr(96, calldataload(97)) // extract 20b addr
                    transferWeth(pair0Addr, shr(184, calldataload(20)))
                    uniswap(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49)))
                    uniswap(pair1Addr, pair2Addr, and(callvalue(), 0x10), shr(144, calldataload(83)))
                    uniswap(pair2Addr, pair3Addr, and(callvalue(), 0x20), shr(144, calldataload(117)))
                    uniswap(pair3Addr, address(), and(callvalue(), 0x40), shr(184, calldataload(131)))
                }
                case 4 /* univ2 five swaps*/{//Doesn't work
                    let pair0Addr := shr(96, calldataload(0)) // extract 20b addr
                    let pair1Addr := shr(96, calldataload(29)) // extract 20b addr
                    let pair2Addr := shr(96, calldataload(63)) // extract 20b addr
                    let pair3Addr := shr(96, calldataload(97)) // extract 20b addr
                    let pair4Addr := shr(96, calldataload(131)) // extract 20b addr
                    transferWeth(pair0Addr, shr(184, calldataload(20))) 
                    uniswap(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49))) // 9b eth value
                    uniswap(pair1Addr, pair2Addr, and(callvalue(), 0x10), shr(144, calldataload(83)))
                    uniswap(pair2Addr, pair3Addr, and(callvalue(), 0x20), shr(144, calldataload(117)))
                    uniswap(pair3Addr, pair4Addr, and(callvalue(), 0x40), shr(144, calldataload(151)))
                    uniswap(pair4Addr, address(), and(callvalue(), 0x80), shr(184, calldataload(165)))
                }
                case 5 /* univ2 six swaps*/{ //?
                    let pair0Addr := shr(96, calldataload(0)) // extract 20b addr
                    let pair1Addr := shr(96, calldataload(29)) // extract 20b addr
                    let pair2Addr := shr(96, calldataload(63)) // extract 20b addr
                    let pair3Addr := shr(96, calldataload(97)) // extract 20b addr
                    let pair4Addr := shr(96, calldataload(131)) // extract 20b addr
                    let pair5Addr := shr(96, calldataload(165)) // extract 20b addr
                    transferWeth(pair0Addr, shr(184, calldataload(20))) 
                    uniswap(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49))) // 9b eth value
                    uniswap(pair1Addr, pair2Addr, and(callvalue(), 0x10), shr(144, calldataload(83)))
                    uniswap(pair2Addr, pair3Addr, and(callvalue(), 0x20), shr(144, calldataload(117)))
                    uniswap(pair3Addr, pair4Addr, and(callvalue(), 0x40), shr(144, calldataload(151)))
                    uniswap(pair4Addr, pair5Addr, and(callvalue(), 0x80), shr(144, calldataload(185)))
                    uniswap(pair5Addr, address(), and(callvalue(), 0x100), shr(184, calldataload(199)))
                }
                // case 6 /* withdraw */ {
                //     transferAsset(0xC49995D326b9438CF4b0AFA1045f2bAD777Cf37F, calldataload(0), calldataload(32))
                // }
                //Anything higher than 7 requires reworking callvalue for eariler swaps
                default {
                    revert(0, 0)
                }
                //Balance checks are not needed for v2 I think
                // balanceOf(0x100) //After the call
                    
                // let before := mload(0x80) // balance before
                // let after := mload(0x100) // balance after

                // // if balance after<before
                // if lt(after,before){
                //     revert(0,0)
                // }
                    
                return(0,0)
                //Add here check if the balance after is bigger than initial balance
            }

            //Univ3, balancer and other mixed swaps go here
            if eq(0xC49995D326b9438CF4b0AFA1045f2bAD777Cf37F, caller()){
                mstore(0x40,add(0x80,0x40)) // Move memeory pointer to leave space for balance before and after
                balanceOf(0x80) // Get WETH balance before the swaps
                    /*
                        bytes: 
                        If first swap is univ2
                            9b Weth Amount
                            14b Pair 0 out
                            20b Pair 0 Address
                        if first swap is univ3
                            9b Weth amount
                            20b pair 0 address

                        Last Pair:
                            univ3:
                            20b pair x address
                            20b pair x token
                            14 pair x in

                            univ2:
                            20b pair x address
                            14b apair x token                        

                        msgValue:
                            b[2-0] - num of swaps
                            b[3] - No Initial Weth trasfer
                            b[4] - Swap 0 - 0 for univ2, 1 for univ3 // This bit is wasted
                            b[5] - Swap 0 - ZeroForOne
                            b[6] - Swap 0 - T0 - o us, 1 to next pool
                            b[7] - Swap 1 univ2 or univ3
                            b[8] - Swap 1 Zero for one
                            b[9] - Swap 1 to us or to next pool
                            b[10] - Swap 2 univ2 or univ3
                            b[11] - Swap 2 zeroForOne
                            b[12] - Swap 2 to us or to next pool
                        */
                    let offset := 0 // used to load addresses
                    //Check if we need to first transfer WETH
                    switch and(callvalue(), 0x8)
                        case 0x0{
                            let addr := shr(96, calldataload(23))
                            transferWeth(addr, shr(184, calldataload(0)))
                            switch and(callvalue(), 0x20)
                            case 0x0{
                                uniswap(addr, address(), and(callvalue(), 16), shr(144, calldataload(9)))
                            }
                            default{
                                uniswap(addr, shr(96, calldataload(43)), and(callvalue(), 16), shr(144, calldataload(9)))
                            }
                            offset := 43
                        }
                        default{
                            switch and(callvalue(), 0x20)
                            case 0x0{
                                v3uniswap(shr(96, calldataload(9)), address(),  shr(184, calldataload(0)), and(callvalue(), 16), 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)
                            }
                            default{
                                v3uniswap(shr(96, calldataload(9)), shr(96, calldataload(29)),  shr(184, calldataload(0)), and(callvalue(), 16), 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2)
                            }
                            offset := 29
                        }

                    let numSwaps := and(callvalue(), 0x7)
                    for { let i := 0 } lt(i, numSwaps) { i := add(i,1) } {
                        let shift := shl(mul(i,3),128)
                        switch and(callvalue(), shift)
                            case 0x0 {
                            switch and(callvalue(), mul(shift, 0x4))
                                case 0x0{
                                    uniswap(shr(96, calldataload(offset)), address(), and(callvalue(), mul(shift, 0x2)), shr(144, calldataload(add(offset,20))))
                                }
                                default{
                                    uniswap(shr(96, calldataload(offset)), shr(96, calldataload(add(offset, 34))), and(callvalue(), mul(shift, 0x2)), shr(144, calldataload(add(offset,20))))
                                }
                                offset := add(offset, 34)
                            }
                            default {
                            switch and(callvalue(), mul(shift, 0x4))
                                case 0x0{
                                    v3uniswap(shr(96, calldataload(offset)), address(),  shr(144, calldataload(add(offset,40))),and(callvalue(), mul(shift, 0x2)), shr(96, calldataload(add(offset, 20))))
                                }
                                default{
                                    v3uniswap(shr(96, calldataload(offset)), shr(96, calldataload(add(offset, 54))),shr(144, calldataload(add(offset,40))), and(callvalue(), mul(shift, 0x2)),  shr(96, calldataload(add(offset, 20))))
                                }
                                offset := add(offset, 54)
                            }     
                    }


                // }

                // //More than 31 requires change in fuction selector and encoding of zero for one

                // default{
                //     revert(0,0)
                // }

                
                balanceOf(0x100) //After the call
                // This needs to be the last thing in the call
                let before := mload(0x80) // balance before
                let after := mload(0x100) // balance after

                // if balance after<before
                if lt(after,before){
                    revert(0,0)
                }
                    
                return(0,0)
            }

            //Any other function goes here, like univ3 callback
            //Check tx.origin == owner. Warning this is not enough if we want to have flashloans
            //Change this to the second address that initiates mixed v2 and v3 swaps
            if eq(origin(), 0xC49995D326b9438CF4b0AFA1045f2bAD777Cf37F){
                //Functions called here all need to return at the end
                switch selector()
                case 0xfa461e33 /* uniswapV3SwapCallback(amount0, amount1, data) */{
                    /*
                        This needs some additional logic depending on the case. 
                        If we need to make another swap here we need to indicate this somehow in calladata
                    */
                    switch sgt(calldataload(4),0)
                        //We need to send amount0 to pool
                        case 1{
                            transferAsset(caller(), calldataload(4), calldataload(132))// amount, toAddr, token
                        }
                        //We need to send amount1 to pool
                        case 0{
                            transferAsset(caller(), calldataload(36), calldataload(132))// amount, toAddr, token 
                        }
                 }
                return(0,0)
            }

            if eq(0x98827093418C47C7fD5F161eC5B8DF1dA474bD9E, caller()) {
                switch and(callvalue(), 0x7)
                case 0 /* withdraw */ {
                    widthdrawWeth() 
                }
                case 6 /* withdraw */ {
                    transferAsset(0x98827093418C47C7fD5F161eC5B8DF1dA474bD9E, calldataload(0), calldataload(32))
                }
                default{
                    revert(0,0)
                }
                return(0,0)
            }

            //Otherwise we revert
            revert(0,0)

            function v3uniswap(pairAddr, toAddr, amount, zeroForOne, token){
                // MinSqrtRatio + 1 = 4295128740
                // MaxSqrtRatio - 1 = 1461446703485210103287273052203988822378723970341
                //0x128acb08 - swap(recipient address, zeroForOne bool, amountSpecified uin256, sqrtPriceLimitX96 uin160, data bytes) 
                let x := mload(0x40)
                mstore(x, 0x128acb0800000000000000000000000000000000000000000000000000000000) //0x128acb08
                mstore(add(x, 0x4), toAddr)
                switch zeroForOne
                case 0{ // False
                    mstore(add(x, 0x24), 0)
                    mstore(add(x, 0x44), amount)
                    mstore(add(x, 0x64), 1461446703485210103287273052203988822378723970341)
                }
                default { // True
                    mstore(add(x, 0x24), 1)
                    mstore(add(x, 0x44), amount)
                    mstore(add(x, 0x64), 4295128740)
                }
                mstore(add(x,0x84) , 160 /* position of where length of "bytes data" is stored from first arg (excluding func signature) */)
                mstore(add(x,0xA4), 0x20) // length of "bytes data", //This should include the pair address , token and amount
                mstore(add(x,0xC4), token)
                // mstore(add(x,0xA5), token)
                if iszero(call(gas(), pairAddr, 0 /* wei */, x /* in pos */, 0xE4 /* in size */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                mstore(0x40,add(x,0xE4))
            }

            function balanceOf(storeAddr){
                let x := mload(0x40) 
                mstore(x, 0x70a0823100000000000000000000000000000000000000000000000000000000) //0x70a08231 - balanceOf(address)
                mstore(add(x,0x4),address())
                if iszero(call(gas(), 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2, 0 /* wei */, x /* in pos */, 0x24 /* in size */, storeAddr /* out pos */, 0x20 /* out size */)) {
                    revert(0, 0)
                }
                mstore(0x40,add(x,0x24))
            }

            function widthdrawWeth() {
                //0xa9059cbb = function transfer(address to, uint value) external returns (bool)
                let x := mload(0x40) 
                mstore(x, 0xa9059cbb00000000000000000000000000000000000000000000000000000000)
                mstore(add(x,0x04), 0x98827093418C47C7fD5F161eC5B8DF1dA474bD9E)
                mstore(add(x,0x24), shr(176, calldataload(0)) /* extract 10b = uint64 */) //This may need to be changed
                if iszero(call(gas(), 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2, 0 /* wei */, x /* in pos */, 0x44 /* in len */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                let b := selfbalance()
                if iszero(call(gas(), 0x98827093418C47C7fD5F161eC5B8DF1dA474bD9E, b, 0,0,0,0)){
                    revert(0,0)
                }
                mstore(0x40,add(x,0x44))
            }

            function transferAsset(toAddr, amount, asset){
                let x := mload(0x40) 
                mstore(x, 0xa9059cbb00000000000000000000000000000000000000000000000000000000)
                mstore(add(x,0x04), toAddr)
                mstore(add(x,0x24), amount)
                if iszero(call(gas(), asset, 0 /* wei */, x /* in pos */, 0x44 /* in len */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                mstore(0x40,add(x,0x44))
            }

            //Can be changed to transferFrom to save some gas
            function transferWeth(toAddr, amount) {
                //0xa9059cbb = function transfer(address to, uint value) external returns (bool)
                // mstore(0x0, 0xa9059cbb)
                // mstore(4, toAddr)
                // mstore(36, amount)

                let x := mload(0x40) 
                mstore(x, 0xa9059cbb00000000000000000000000000000000000000000000000000000000)
                mstore(add(x,0x04), toAddr)
                mstore(add(x,0x24), amount)
                if iszero(call(gas(), 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2, 0 /* wei */, x /* in pos */, 0x44 /* in len */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                mstore(0x40,add(x,0x44))
            }

            function uniswap(pairAddr, toAddr, isToken0OutAppliedMask, tokenOutAmount) {
                switch isToken0OutAppliedMask
                case 0x0 {
                    swap(pairAddr, tokenOutAmount, 0, toAddr)
                }
                default {
                    swap(pairAddr, 0, tokenOutAmount, toAddr)
                }
            }

            function swap(pairAddr, amount0Out, amount1Out, toAddr) {
                // build swap call input
                //0x090f344e - function swap(uint amount0Out, uint amount1Out, address to, bytes data)

                //THIS WILL HAVE TO BE SHIFTED
                let x := mload(0x40) 
                mstore(x, 0x022c0d9f00000000000000000000000000000000000000000000000000000000)
                mstore(add(x,0x04), amount0Out)
                mstore(add(x,0x24) /* 4+32 */, amount1Out)
                mstore(add(x,0x44) /* 4+32+32 */, toAddr)
                mstore(add(x,0x64) /* 4+32+32+32 */, 128 /* position of where length of "bytes data" is stored from first arg (excluding func signature) */)
                mstore(add(x,0x84), 0x0) // length of "bytes data"
                if iszero(call(gas(), pairAddr, 0 /* wei */, x /* in pos */, 0xA4 /* in size */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                mstore(0x40,add(x,0xA4))
            }

            //This may be optimized probalby
            function selector() -> s {
                s := div(calldataload(0), 0x100000000000000000000000000000000000000000000000000000000)
            }
        }
    }
}
