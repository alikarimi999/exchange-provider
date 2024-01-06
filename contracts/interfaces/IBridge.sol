// SPDX-License-Identifier: MIT
pragma solidity 0.8.18;

interface IBridge {
    function Bridge(bytes calldata data,uint amountIn) external payable;
}