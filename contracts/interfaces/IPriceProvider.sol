// SPDX-License-Identifier: MIT
pragma solidity >=0.7.6;

interface IPriceProvider {
 function estimateAmountOut(address provider,address tA,address tB,uint256 amountIn,uint8 version) external view returns (uint256 amountOut,uint24 fee);
}