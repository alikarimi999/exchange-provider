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
import './interfaces/IBridgeAggregator.sol';
import './interfaces/IBridge.sol';
import './Multicall.sol';



contract ExchangeAggregator is 
    IExchangeAggregator,
    IBridgeAggregator,
    Ownable,
    Multicall {

    address public feeReciever;
    
    mapping(bytes => bool) private processedSignatures; // Mapping to track processed signatures
      modifier signatureNotProcessed(bytes calldata sig) {
        require(!processedSignatures[sig], "this txData already processed");
        _;
        processedSignatures[sig] = true;
    }

       constructor(){
        feeReciever = msg.sender;
    }

    receive() external payable {}

    function estimateAmountOut(address priceProvider, address provider,address tA,address tB,uint256 amountIn,uint8 version) external view returns (uint256 amountOut,uint24 fee){
        return IPriceProvider(priceProvider).estimateAmountOut(provider,tA,tB,amountIn,version);
    }

    uint256 public swapAmountOut;
    function Swap(swapInput calldata data,bytes calldata sig) public payable signatureNotProcessed(sig) {
        require(data.sender == msg.sender,"invalid sender");
        utils.checkSig(owner(),abi.encode(data), sig);
        uint value;
        if (!data.native) {
            if (!data.fromContract) {
                TransferHelper.safeTransferFrom(data.tokenIn,msg.sender,address(this),data.totalAmount);
            }
            if (data.feeAmount > 0){
                TransferHelper.safeTransfer(data.tokenIn,feeReciever,data.feeAmount);
            }
            TransferHelper.safeApprove(data.tokenIn,data.swapper,data.amountIn);
        }else {
            require(msg.value >= data.totalAmount,"insufficient value");
            TransferHelper.safeTransferETH(feeReciever,data.feeAmount);
            value = data.amountIn;
        }

        uint256 balance0 = IERC20(data.tokenOut).balanceOf(address(this));
        SafeCaller.safeCall(data.swapper,value,data.swapperData);
        uint256 balance1 = IERC20(data.tokenOut).balanceOf(address(this));
        
        if (balance1 > balance0){      
        swapAmountOut = balance1 - balance0;
        }else {
            swapAmountOut = 0;
        }
        processedSignatures[sig] = true;
    }


    function Bridge(bridgeInput calldata data,bytes calldata sig) public payable signatureNotProcessed(sig) {
        require(data.sender == msg.sender,"invalid sender");
        require(msg.value >= data.bridgeFee,"insufficient value");
        utils.checkSig(owner(),abi.encode(data), sig);
        uint amountIn;
        if (data.afterSwap) {
            amountIn = swapAmountOut;
            swapAmountOut = 0;
           
        }else {
            TransferHelper.safeTransferFrom(data.tokenIn,msg.sender,address(this),data.amountIn+data.feeAmount);
            if (data.feeAmount > 0) {
                    TransferHelper.safeTransfer(data.tokenIn,feeReciever,data.feeAmount);
            }
            amountIn = data.amountIn;
        }


        TransferHelper.safeApprove(data.tokenIn,data.bridge,amountIn);
        IBridge(data.bridge).Bridge{value:data.bridgeFee}(data.bridgeData,amountIn);
    }

    function ChangeFeeReciever(address _feeReciever) public onlyOwner {
        feeReciever = _feeReciever;
    }

    function Balance(address token) public view returns(uint256) {
        return IERC20(token).balanceOf(address(this));
    }

    function BalanceETH() public view returns(uint256){
        return address(this).balance;
    }

    function Withdraw(address token,address to,uint256 amount) public onlyOwner {
        TransferHelper.safeTransfer(token,to,amount);
    }

    function WithdrawETH(address to,uint256 amount) public payable onlyOwner {
        TransferHelper.safeTransferETH(to,amount); 
   }

}
