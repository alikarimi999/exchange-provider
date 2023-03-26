// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;

import './interfaces/IPriceProvider.sol';
import './libraries/UniswapV2.sol';
import './libraries/UniswapV3.sol';


contract PriceProvider is IPriceProvider {
        function estimateAmountOut(address provider,address tA,address tB,uint256 amountIn,uint8 version) external override view returns (uint256 amountOut,uint24 fee){
        if (version == 2) {
            return (UniswapV2.EstimateAmountOut(provider, tA, tB,amountIn),0);
        }else if (version == 3) {
            return UniswapV3.EstimateAmountOut(provider, tA, tB, amountIn);
        }
    }
}