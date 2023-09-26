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
        - DODO/Curve
        - Min amount out
        - UniswapV3 reverts
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
                    transferWethPure(pair0Addr, shr(184, calldataload(20)))
                    uniswapPure(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49))) // 9b eth value
                    uniswapPure(pair1Addr, address(), and(callvalue(), 0x10), shr(184, calldataload(63)))
                }
                case 2 /* univ2 three swaps*/ { //Works
                    let pair0Addr := shr(96, calldataload(0)) // extract 20b addr
                    let pair1Addr := shr(96, calldataload(29)) // extract 20b addr
                    let pair2Addr := shr(96, calldataload(63)) // extract 20b addr
                    transferWethPure(pair0Addr, shr(184, calldataload(20)))
                    uniswapPure(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49))) 
                    uniswapPure(pair1Addr, pair2Addr, and(callvalue(), 0x10), shr(144, calldataload(83)))
                    uniswapPure(pair2Addr, address(), and(callvalue(), 0x20), shr(184, calldataload(97)))
                }
                case 3 /* univ2 four swaps*/{ // Works
                    let pair0Addr := shr(96, calldataload(0)) // extract 20b addr
                    let pair1Addr := shr(96, calldataload(29)) // extract 20b addr
                    let pair2Addr := shr(96, calldataload(63)) // extract 20b addr
                    let pair3Addr := shr(96, calldataload(97)) // extract 20b addr
                    transferWethPure(pair0Addr, shr(184, calldataload(20)))
                    uniswapPure(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49)))
                    uniswapPure(pair1Addr, pair2Addr, and(callvalue(), 0x10), shr(144, calldataload(83)))
                    uniswapPure(pair2Addr, pair3Addr, and(callvalue(), 0x20), shr(144, calldataload(117)))
                    uniswapPure(pair3Addr, address(), and(callvalue(), 0x40), shr(184, calldataload(131)))
                }
                case 4 /* univ2 five swaps*/{//Doesn't work
                    let pair0Addr := shr(96, calldataload(0)) // extract 20b addr
                    let pair1Addr := shr(96, calldataload(29)) // extract 20b addr
                    let pair2Addr := shr(96, calldataload(63)) // extract 20b addr
                    let pair3Addr := shr(96, calldataload(97)) // extract 20b addr
                    let pair4Addr := shr(96, calldataload(131)) // extract 20b addr
                    transferWethPure(pair0Addr, shr(184, calldataload(20))) 
                    uniswapPure(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49))) // 9b eth value
                    uniswapPure(pair1Addr, pair2Addr, and(callvalue(), 0x10), shr(144, calldataload(83)))
                    uniswapPure(pair2Addr, pair3Addr, and(callvalue(), 0x20), shr(144, calldataload(117)))
                    uniswapPure(pair3Addr, pair4Addr, and(callvalue(), 0x40), shr(144, calldataload(151)))
                    uniswapPure(pair4Addr, address(), and(callvalue(), 0x80), shr(184, calldataload(165)))
                }
                case 5 /* univ2 six swaps*/{ //?
                    let pair0Addr := shr(96, calldataload(0)) // extract 20b addr
                    let pair1Addr := shr(96, calldataload(29)) // extract 20b addr
                    let pair2Addr := shr(96, calldataload(63)) // extract 20b addr
                    let pair3Addr := shr(96, calldataload(97)) // extract 20b addr
                    let pair4Addr := shr(96, calldataload(131)) // extract 20b addr
                    let pair5Addr := shr(96, calldataload(165)) // extract 20b addr
                    transferWethPure(pair0Addr, shr(184, calldataload(20))) 
                    uniswapPure(pair0Addr, pair1Addr, and(callvalue(), 0x8), shr(144, calldataload(49))) // 9b eth value
                    uniswapPure(pair1Addr, pair2Addr, and(callvalue(), 0x10), shr(144, calldataload(83)))
                    uniswapPure(pair2Addr, pair3Addr, and(callvalue(), 0x20), shr(144, calldataload(117)))
                    uniswapPure(pair3Addr, pair4Addr, and(callvalue(), 0x40), shr(144, calldataload(151)))
                    uniswapPure(pair4Addr, pair5Addr, and(callvalue(), 0x80), shr(144, calldataload(185)))
                    uniswapPure(pair5Addr, address(), and(callvalue(), 0x100), shr(184, calldataload(199)))
                }
                case 6 /* supports multiple pure v2 swaps */{
                    /*
                        The idea of this is to support multiple univ2 arbitarge in same tx. Allows for a higher bribe. Not implemented in this relase

                        msg.value = use it to indicate how many weth transfers we have to do and how many swaps each of them involves
                        20b pair0 address
                        9b  pair0 input amount up to 4722 ETH - used for eth transfer
                        20b pair1 address - 320gas
                        14b pair0 output amount
                        9b pair1 output amount 
                        

                    */
                    revert(0, 0)

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
                if lt(gas(), 400000){
                    revert(0,0)
                }
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
                            9B Weth Amount
                            14B Pair 0 out
                            20B Pair 0 Address

                        if first swap is univ3
                            14B Amount In
                            20B pair 0 address
                            20B to addr
                            2B data len 
                            XB bytes for data

                        If first swap is dodo:
                            9B Weth Amount
                            20B Pair 0 Address
                        
                        If first swap id curve
                            9B Weth amount
                            20B Pair 0 Address
                            1B i
                            1B j
                            20B Token Out
                            14B Amount Out


                        Last Pair:
                            univ3:
                            20b - pair x address
                            14B - pair x in
                            20B - pair to 
                            2B - data len
                            XB - data bytes
                            9b - amountMinOut!!! //Also for dodo and curve if they are the last one

                            univ2:
                            20B - pair x address
                            14B - pair x out   

                            dodo: 
                            20b pair x address
                            
                            curve:
                            20B - Pair 0 Address
                            1B - i
                            1B - j
                            14B - dx(Amount in)
                            20B - Token Out - only if we sent to other address
                            14B - Amount Out - only if we sent to other address


                        msgValue:
                            b[2-0] - num of swaps
                            b[3-4] - 00 univ2, 01 univ3, 10 dodo, 11 curve
                            b[5] - Swap 0 - ZeroForOne
                            b[6] - Swap 0 - T0 - o us, 1 to next pool 
                            b[7-8] - 00 univ2, 01 univ3, 10 dodo, 11 curve
                            b[9]  - Swap 1 Zero for one
                            b[10] - Swap 1 to us or to next pool
                            b[11-12] - 00 univ2, 01 univ3, 10 dodo, 11 curve
                            b[13] - Swap 2 zeroForOne
                            b[14] - Swap 2 to us or to next pool
                        */

                    let offset := 0
                    switch and(callvalue(), 0x8)
                        case 0x0{
                            switch and(callvalue(), 0x10)
                            case 0x0{ //UniV2
                                let addr := shr(96, calldataload(23))
                                transferWeth(addr, shr(184, calldataload(0)))
                                uniswap(addr, shr(96, calldataload(43)), and(callvalue(), 0x20), shr(144, calldataload(9)))
                                offset := 43
                            }
                            default{ //UniV3
                                let x := mload(0x40)
                                mstore(x, 0x128acb0800000000000000000000000000000000000000000000000000000000) //0x128acb08
                                mstore(add(x, 0x4), shr(96,calldataload(34))) //toAddr
                                switch and(callvalue(), 0x20) //ZeroForOne
                                    case 0{ // False
                                        mstore(add(x, 0x24), 0)
                                        mstore(add(x, 0x44), shr(144,calldataload(0))) //Amount
                                        mstore(add(x, 0x64), 1461446703485210103287273052203988822378723970341)
                                    }
                                default { // True
                                    mstore(add(x, 0x24), 1)
                                    mstore(add(x, 0x44), shr(144,calldataload(0))) // Amount
                                    mstore(add(x, 0x64), 4295128740)
                                }
                            mstore(add(x,0x84) , 160 /* position of where length of "bytes data" is stored from first arg (excluding func signature) */)
                            let dataLen := shr(240, calldataload(54))// 2bytes
                            mstore(add(x,0xA4), dataLen) // length of "bytes data", //This should include remaining data needed for the rest of callbacks
                            let memOffset := 0xC4
                            for {let k:=56} lt(k, add(dataLen,56)) {k := add(k, 32)}{
                                mstore(add(x, memOffset), calldataload(k)) // Change this
                                memOffset := add(32, memOffset)
                            }
                            // mstore(add(x,0xA5), token)
                            if iszero(call(gas(), shr(96,calldataload(14)), 0 /* wei */, x /* in pos */, memOffset /* in size */, 0 /* out pos */, 0 /* out size */)) {
                                revert(0, 0)
                            }
                            mstore(0x40,add(x,memOffset))
                            offset := add(80, dataLen)
                            }
                        }
                        default{
                            switch and(callvalue(), 0x10)
                            case 0x0{ //Dodo
                                transferWeth(shr(96,calldataload(9)), shr(184, calldataload(0)))
                                dodoSwap(shr(96,calldataload(9)), shr(96, calldataload(29)), and(callvalue(), 0x20))
                                offset := 29
                            }
                            default{ //Curve
                                //curveSwap(pairAddr, i,j,dx)
                                transferWeth(shr(96,calldataload(9)), shr(184, calldataload(0)))
                                curveSwap(shr(96,calldataload(9)), shr(248,calldataload(29)),shr(248,calldataload(30)),shr(184, calldataload(0))) // addr, i,j, dx
                                transferAsset(shr(96,calldataload(65)),shr(144,calldataload(51)),shr(96,calldataload(31))) //transferAsset(toAddr, amount, asset)
                                offset := 65
                            }
                        }
 
                
                    // Num of swaps on this stack levelfor example if we have v2 -> v3 (2 swaps in callback) -> dodo this would be 3
                    let numSwaps := and(callvalue(), 0x7)
                    for { let i := 0 } lt(i, numSwaps) { i := add(i,1) } {
                        let shift := shl(mul(i,4),128)
                        switch and(callvalue(), shift)
                        case 0x0{
                            switch and(callvalue(), mul(shift, 0x2))
                            case 0x0{ //UniV2
                                switch and(callvalue(), mul(shift, 0x8))
                                case 0x0{
                                    uniswap(shr(96, calldataload(offset)), address(), and(callvalue(), mul(shift, 0x4)), shr(144, calldataload(add(offset,20))))
                                }
                                default{
                                    uniswap(shr(96, calldataload(offset)), shr(96, calldataload(add(offset, 34))), and(callvalue(), mul(shift, 0x4)), shr(144, calldataload(add(offset,20))))
                                }
                                offset := add(offset, 34)
                            }
                            default{ //UniV3
                                let x := mload(0x40)
                                mstore(x, 0x128acb0800000000000000000000000000000000000000000000000000000000) //0x128acb08
                                switch and(callvalue(), mul(shift, 0x8))
                                case 0x0{
                                    mstore(add(x, 0x4), address()) //toAddr
                                }
                                default{
                                    mstore(add(x, 0x4), shr(96,calldataload(add(offset, 34)))) //toAddr
                                }
                                switch and(callvalue(), 0x4) //ZeroForOne
                                case 0{ // False
                                    mstore(add(x, 0x24), 0)
                                    mstore(add(x, 0x44), shr(144,calldataload(add(offset, 20)))) //Amount
                                    mstore(add(x, 0x64), 1461446703485210103287273052203988822378723970341)
                                }
                                default { // True
                                    mstore(add(x, 0x24), 1)
                                    mstore(add(x, 0x44), shr(144,calldataload(add(offset, 20)))) // Amount
                                    mstore(add(x, 0x64), 4295128740)
                                }
                                mstore(add(x,0x84) , 160 /* position of where length of "bytes data" is stored from first arg (excluding func signature) */)
                                let dataLen := shr(240, calldataload(add(offset,54)))// 2bytes
                                mstore(add(x,0xA4), dataLen) // length of "bytes data", //This should include remaining data needed for the rest of callbacks
                                let memOffset := 0xC4
                                for {let k:=56} lt(k, add(dataLen,56)) {k := add(k, 32)}{
                                    mstore(add(x, memOffset), calldataload(k)) // Change this
                                    memOffset := add(32, memOffset)
                                }
                                // mstore(add(x,0xA5), token)
                                if iszero(call(gas(), shr(96,calldataload(offset)), 0 /* wei */, x /* in pos */, memOffset /* in size */, 0 /* out pos */, 0 /* out size */)) {
                                    revert(0, 0)
                                }
                                mstore(0x40,add(x,memOffset))
                                offset := add(offset,add(56, dataLen))
                            }
                        }
                        default{
                            switch and(callvalue(), mul(shift, 0x2))
                            case 0x0{ //Dodo
                                switch and(callvalue(), mul(shift, 0x8))
                                case 0x0{
                                    dodoSwap(shr(96,calldataload(offset)), address(), and(callvalue(), mul(shift, 0x4))) //dodoSwap(pairAddr, toAddr, zeroForOne)
                                }
                                default{
                                    dodoSwap(shr(96,calldataload(offset)), shr(96, calldataload(add(offset,20))), and(callvalue(), mul(shift, 0x4))) //dodoSwap(pairAddr, toAddr, zeroForOne)
                                }
                                offset := add(offset, 20)
                            }
                            default{ //Curve
                                //curveSwap(pairAddr, i,j,dx)
                                curveSwap(shr(96,calldataload(offset)), shr(248,calldataload(add(offset,20))),shr(248,calldataload(add(offset,21))),shr(184, calldataload(add(offset,22)))) // addr, i,j, dx
                                switch and(callvalue(), mul(shift, 0x8))
                                case 0x0{
                                    offset := add(offset, 36)
                                }default{
                                    transferAsset(shr(96,calldataload(70)),shr(144,calldataload(56)),shr(96,calldataload(36))) //transferAsset(toAddr, amount, asset)
                                    offset := add(offset, 56)
                                }
                            }
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

                // if balance after<before+minAmountIn
                if lt(after,add(before, shr(184,calldataload(offset)))){
                    revert(0,0)
                }

                //This is so we can have 200k higher gas limit while still having a limitation on how much can be used. 
                //Useful when swap makes a lot of calls on the stack
                if lt(gas(), 400000){
                    revert(0,0)
                }
                    
                return(0,0)
            }


            //Any other function goes here, like univ3 callback
            //Check tx.origin == owner. Warning this is not enough if we want to have flashloans
            //Change this to the second address that initiates mixed v2 and v3 swaps
            if eq(origin(), 0xC49995D326b9438CF4b0AFA1045f2bAD777Cf37F){
                //Functions called here all need to return at the end
                callbackLogic()
                // switch selector()
                // case 0xfa461e33 /* uniswapV3SwapCallback(amount0, amount1, data) */{
                //     /*
                //         This needs some additional logic depending on the case. 
                //         If we need to make another swap here we need to indicate this somehow in calladata
                //     */
                //     switch sgt(calldataload(4),0)
                //         //We need to send amount0 to pool
                //         case 1{
                //             transferAsset(caller(), calldataload(4), calldataload(132))// amount, toAddr, token
                //         }
                //         //We need to send amount1 to pool
                //         case 0{
                //             transferAsset(caller(), calldataload(36), calldataload(132))// amount, toAddr, token 
                //         }
                //  }
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


            /*
                Used to save on transfer with v3 swaps.
                DataLayout (calldata):
                    [0-4]B - function sig
                    [5]B - num calls that callback has to make
                    Every swap then:
                    [1]B - type(v2,v3,curve,dodo) (up to 256 types)

                    If v2:
                        [0]B - type
                        [1-20]B - pair x address
                        [21-40]B - destination address
                        [41]B - zerForOne
                        [42-62]B - pair x out

                    If v3:
                        [0]B - type
                        [1-20]B pair x address
                        [21-35]B pair x in
                        [36-56]B toAddr
                        [57]B - zero for one 
                        [58-60]B - how many bytes to read for v3 data
                        [61 - X]Defined in the previous line - data bytes for swaps

                    If dodo:
                        [0]B - type
                        [1-20]B - pair x address
                        [21-40]B - destination address
                        [41]B - zerForOne

                    If curve:
                        [0]B - type
                        [1-20]B - pairAddr
                        [21]B - i
                        [22]B - j
                        [23-37]B - dx
                        [38]B -  send to next pair
                        if 1:
                            [39-59]B - pairToAddr
                            [60-80]B - token addr
                        
                    if transferWeth
                        20b Pair to Addr
                        9b Amount
                        */ 

            function callbackLogic(){
                /*
                        4B - func sig
                        32B - int256 amount0Delta,
                        32B - int256 amount1Delta,
                        X - ??? bytes location 
                        32B - data len
                        XB - data 
                */
                let numCallbackSwaps := shr(248,calldataload(132))
                let callBackOffset := 133
                for { let i := 0 } lt(i, numCallbackSwaps) { i := add(i,1) } {
                    switch shr(248,calldataload(callBackOffset)) //read swap type
                    case 0{ // UniswapV2
                        uniswap(shr(96,calldataload(add(callBackOffset,1))), shr(96,calldataload(add(callBackOffset,21))), shr(248,calldataload(add(callBackOffset, 41))), shr(144, calldataload(add(callBackOffset,42))))
                        callBackOffset := add(callBackOffset, 42)
                    }
                    case 1{ // UniswapV3

                        let x := mload(0x40)
                        mstore(x, 0x128acb0800000000000000000000000000000000000000000000000000000000) //0x128acb08
                        mstore(add(x, 0x4), shr(96,calldataload(add(callBackOffset, 36)))) //toAddr
                        switch shr(248, calldataload(add(callBackOffset,57))) //ZeroForOne
                        case 0{ // False
                            mstore(add(x, 0x24), 0)
                            mstore(add(x, 0x44), shr(144,calldataload(add(callBackOffset, 41)))) //Amount
                            mstore(add(x, 0x64), 1461446703485210103287273052203988822378723970341)
                        }
                        default { // True
                            mstore(add(x, 0x24), 1)
                            mstore(add(x, 0x44), shr(144,calldataload(add(callBackOffset, 41)))) // Amount
                            mstore(add(x, 0x64), 4295128740)
                        }
                        mstore(add(x,0x84) , 160 /* position of where length of "bytes data" is stored from first arg (excluding func signature) */)
                        let dataLen := shr(240, calldataload(add(callBackOffset, 58)))// 2bytes
                        mstore(add(x,0xA4), dataLen) // length of "bytes data", //This should include remaining data needed for the rest of callbacks
                        let memOffset := 0xC4
                        for {let k:=80} lt(k, add(dataLen,61)) {k := add(k, 32)}{
                            mstore(add(x, memOffset), calldataload(add(callBackOffset, k))) // Change this
                            memOffset := add(32, memOffset)
                        }
                        // mstore(add(x,0xA5), token)
                        if iszero(call(gas(), shr(96, calldataload(add(callBackOffset, 1))), 0 /* wei */, x /* in pos */, memOffset /* in size */, 0 /* out pos */, 0 /* out size */)) {
                            revert(0, 0)
                        }
                        mstore(0x40,add(x,memOffset))
                        callBackOffset := add(callBackOffset, add(61, dataLen))
                    }
                    case 2{ // Dodo
                        dodoSwap(shr(96,calldataload(add(callBackOffset,1))), shr(96,calldataload(add(callBackOffset,21))), shr(248,calldataload(add(callBackOffset, 41))))
                        callBackOffset := add(callBackOffset, 42)
                    }
                    case 3{ //Curve
                        /*curveSwap(pairAddr, i,j,dx) 
                        20b - pairAddr
                        1b - i
                        1b - j
                        14b - dx
                        1b - Send to next pair
                        if 1:
                            20b - pairToAddr
                            20b - token addr
                        */ 
                        curveSwap(shr(96, calldataload(add(callBackOffset, 1))), shr(248, calldataload(add(callBackOffset, 21))),shr(248, calldataload(add(callBackOffset, 22))),shr(144, calldataload(add(callBackOffset, 23))))
                        //Transfer addr
                        switch shr(248,calldataload(add(callBackOffset, 38))) //read swap type
                        case 1{
                            transferAsset(caller(), shr(96, calldataload(add(callBackOffset,39))), shr(96, calldataload(add(callBackOffset,60))))// amount, toAddr, token
                            callBackOffset := add(callBackOffset, 81)
                        }
                        default{
                            callBackOffset := add(callBackOffset, 38)
                        }
                    }
                    case 4{ //Weth transfer(usually the case when v3 is the first swap

                        //ALL OFFSETS NEED TO BE CHANGED TO TAKE INTO ACCOUNT LOCATTION OF BYTES LEN
                        switch sgt(calldataload(4),0)
                        //We need to send amount0 to pool
                        case 1{
                            transferAsset(caller(), calldataload(4), calldataload(132))// amount, toAddr, token
                        }
                        //We need to send amount1 to pool
                        case 0{
                            transferAsset(caller(), calldataload(36),calldataload(132))// amount, toAddr, token 
                        }
                        callBackOffset := add(callBackOffset, 162)

                    }
                    case 5{
                        //transferWeth
                        // 20B Pair Addr
                        // 9B amount
                        transferWeth(shr(96, calldataload(add(callBackOffset,1))), shr(184, calldataload(add(callBackOffset, 21))))
                        callBackOffset := add(callBackOffset, 30)
                    }
                    default{
                        revert(0,0)
                    }
                }

            }

            /* 
                Add way to call other swaps in reverse
                Most likely will need to use 
            */
            function v3uniswapInitial(pairAddr, toAddr, amount, zeroForOne){
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
                mstore(add(x,0xA4), 0x1) // length of "bytes data", //This indicates that this is the inital v3 swap which needs just weth transfer(type 4)
                mstore8(add(x,0xC4), 0x4)
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

            //Can be changed to transferFrom to save some gas
            function transferWethPure(toAddr, amount) {
                //0xa9059cbb = function transfer(address to, uint value) external returns (bool)
                // mstore(0x0, 0xa9059cbb)
                // mstore(4, toAddr)
                // mstore(36, amount)

                //let x := mload(0x40) 
                mstore(0x40, 0xa9059cbb00000000000000000000000000000000000000000000000000000000)
                mstore(0x44, toAddr)
                mstore(0x64, amount)
                if iszero(call(gas(), 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2, 0 /* wei */, 0x40/* in pos */, 0x44 /* in len */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                //mstore(0x40,add(x,0x44))
            }

            function uniswapPure(pairAddr, toAddr, isToken0OutAppliedMask, tokenOutAmount) {
                switch isToken0OutAppliedMask
                case 0x0 {
                    swapPure(pairAddr, tokenOutAmount, 0, toAddr)
                }
                default {
                    swapPure(pairAddr, 0, tokenOutAmount, toAddr)
                }
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

            function swapPure(pairAddr, amount0Out, amount1Out, toAddr) {
                // build swap call input
                //0x090f344e - function swap(uint amount0Out, uint amount1Out, address to, bytes data)

                //let x := mload(0x40) 
                mstore(0x40, 0x022c0d9f00000000000000000000000000000000000000000000000000000000)
                mstore(0x44, amount0Out)
                mstore(0x64 /* 4+32 */, amount1Out)
                mstore(0x84 /* 4+32+32 */, toAddr)
                mstore(0xA4 /* 4+32+32+32 */, 128 /* position of where length of "bytes data" is stored from first arg (excluding func signature) */)
                mstore(0xD4, 0x0) // length of "bytes data"
                if iszero(call(gas(), pairAddr, 0 /* wei */, 0x40 /* in pos */, 0xA4 /* in size */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                //mstore(0x40,add(x,0xA4))
            }


            function dodoSwap(pairAddr, toAddr, zeroForOne){
                switch zeroForOne
                case 0x0{
                    dodoSellBase(pairAddr, toAddr)
                }
                default{
                    dodoSellQuote(pairAddr, toAddr)
                }
            }

            //Curve sends tokens to msg.sender, min_dy probalby can be 0
            // i and j can be max 8, so we only need 5 bits in calldata from them
            function curveSwap(pairAddr, i,j,dx){
                let x := mload(0x40)  
                mstore(x, 0x5b41b90800000000000000000000000000000000000000000000000000000000) // Exchange without sent_eth
                mstore(add(x,0x04), i) 
                mstore(add(x,0x24), j)
                mstore(add(x,0x44), dx)
                mstore(add(x,0x64), 0) //min_dy
                if iszero(call(gas(), pairAddr, 0 /* wei */, x /* in pos */, 0x84 /* in size */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                mstore(0x40,add(x,0x84))
            }

            //TODO: Figure out what makes Base Token and Quote Token
            function dodoSellBase(pairAddr, toAddr){
                let x := mload(0x40) 
                mstore(x, 0xbd6015b400000000000000000000000000000000000000000000000000000000)
                mstore(add(x,0x04), toAddr)
                if iszero(call(gas(), pairAddr, 0 /* wei */, x /* in pos */, 0x24 /* in size */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                mstore(0x40,add(x,0x24))
            }

            function dodoSellQuote(pairAddr, toAddr){
                let x := mload(0x40) 
                mstore(x, 0xdd93f59a00000000000000000000000000000000000000000000000000000000)
                mstore(add(x,0x04), toAddr)
                if iszero(call(gas(), pairAddr, 0 /* wei */, x /* in pos */, 0x24 /* in size */, 0 /* out pos */, 0 /* out size */)) {
                    revert(0, 0)
                }
                mstore(0x40,add(x,0x24))
            }


            //This may be optimized probalby
            function selector() -> s {
                s := div(calldataload(0), 0x100000000000000000000000000000000000000000000000000000000)
            }
        }
    }
}
