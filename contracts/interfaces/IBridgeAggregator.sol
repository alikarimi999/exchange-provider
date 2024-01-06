// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

interface IBridgeAggregator {
    struct bridgeInput {
        address bridge;
        address tokenIn;
        address sender;
        uint bridgeFee;
        bool afterSwap;
        uint amountIn;
        uint feeAmount;
        bytes bridgeData;
    }

    function Bridge(bridgeInput calldata data,bytes calldata sig) external payable;
}