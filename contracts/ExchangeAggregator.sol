// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import '@openzeppelin/contracts/access/Ownable.sol';
import './interfaces/IWETH.sol';
import './interfaces/IPriceProvider.sol';
import './interfaces/IERC20.sol';
import './libraries/transferHelper.sol';
import './libraries/safeCaller.sol';
import './libraries/utils.sol';
import './interfaces/IExchangeAggregator.sol';


contract ExchangeAggregator is IExchangeAggregator,Ownable,IPriceProvider {
    address public WETH;
    address public PriceProvider;
    
    constructor(address _WETH,address _PriceProvider){
        WETH = _WETH;
        PriceProvider = _PriceProvider;
    }


    function estimateAmountOut(address provider,address tA,address tB,uint256 amountIn,uint8 version) external view returns (uint256 amountOut,uint24 fee){
        return IPriceProvider(PriceProvider).estimateAmountOut(provider,tA,tB,amountIn,version);
    }

    function swap(swapData calldata data,bytes calldata sig) public {
        require(data.sender == msg.sender,"invaled sender");
        utils.checkSig(owner(),abi.encode(data), sig);
        TransferHelper.safeTransferFrom(data.input,msg.sender,address(this),data.totalAmount);
        TransferHelper.safeApprove(data.input,data.swapper,data.totalAmount-data.feeAmount);
        SafeCaller.safeCall(data.swapper,0,data.data);
    }

    function swapNativeIn(swapData calldata data,bytes calldata sig) public payable {
        require(data.sender == msg.sender,"invaled sender");
        utils.checkSig(owner(),abi.encode(data), sig);
        require(msg.value >= data.totalAmount,"insufficient input amount");
        uint amount = msg.value - data.feeAmount;  
        SafeCaller.safeCall(data.swapper,amount,data.data);
    }

    function balanceToken(address token) public view returns(uint){
       return IERC20(token).balanceOf(address(this));
    }

    function balanceETH() public view returns(uint){
        return address(this).balance;
    }

    function withdrawETH(address to,uint amount) public onlyOwner {
        TransferHelper.safeTransferETH(to,amount);
    }

    function withdrawToken(address token,address to,uint amount) public onlyOwner {
        TransferHelper.safeTransfer(token,to,amount);
    }

    function changePriceProvider(address _PriceProvider) public onlyOwner {
        PriceProvider = _PriceProvider;
    }

}
