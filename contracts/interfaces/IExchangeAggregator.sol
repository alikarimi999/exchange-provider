// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

interface IExchangeAggregator {
    struct swapInput{
        address tokenIn;
        address tokenOut;
        uint totalAmount;
        uint feeAmount;
        uint amountIn;
        bool fromContract;
        address swapper;
        bytes swapperData;
        address sender;
        address receiver;
        bool native;
    }

    function Swap(swapInput calldata data,bytes calldata sig) external payable;
}