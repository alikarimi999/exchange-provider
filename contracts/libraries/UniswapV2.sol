// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;

import '../interfaces/IERC20.sol';
import '@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router01.sol';
import '@uniswap/v2-core/contracts/interfaces/IUniswapV2Pair.sol';
import '@uniswap/v2-core/contracts/interfaces/IUniswapV2Factory.sol';

library UniswapV2 { 
    function EstimateAmountOut(address router,address tA, address tB,uint256 amountIn) external view returns (uint256){
        address[] memory path = new address[](2);
        path[0] = tA;
        path[1] = tB;
        uint[] memory amounts = IUniswapV2Router01(router).getAmountsOut(amountIn,path);
       return amounts[1];
    }
    
}
