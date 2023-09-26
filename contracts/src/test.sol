
pragma solidity ^0.8.13;

contract Test{
    uint256 public count;

    function incrase() public{
        count++;
    }
    function incraseByX(uint256 x) public{
        count = count + x;
    }

    function checkFree() public view returns(uint256){
        return count;
    }

    function checkNotFree() public returns(uint256){
        count += 1;
        return count;
    }
}

