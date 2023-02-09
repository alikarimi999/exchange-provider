// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

interface IExchangeAggregator {
    struct swapData{
        address input;
        uint totalAmount;
        uint feeAmount;
        address swapper;
        bytes data;
        address sender;
    }

    function swap(swapData calldata data,bytes calldata sig) external;
    function swapNativeIn(swapData calldata data,bytes calldata sig) external payable;
}