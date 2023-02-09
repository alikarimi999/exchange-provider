// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;

import '../interfaces/IERC20.sol';
import '@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router01.sol';
import '@uniswap/v2-core/contracts/interfaces/IUniswapV2Pair.sol';
import '@uniswap/v2-core/contracts/interfaces/IUniswapV2Factory.sol';

library UniswapV2 {
    function Price(address router, address t0,address t1) external view returns (uint256){
        uint amountIn = 10**IERC20(t0).decimals();
        address[] memory path = new address[](2);
        path[0] = t0;
        path[1] = t1;
        uint[] memory amounts = IUniswapV2Router01(router).getAmountsOut(amountIn,path);
       return amounts[1];
    }

    function PoolExists(address factory,address t0,address t1,uint min0,uint min1) external view returns (bool){
        address pool = IUniswapV2Factory(factory).getPair(t0,t1);
        if (pool == address(0)) return false;
        (uint112 r0, uint112 r1,) = IUniswapV2Pair(pool).getReserves();
        if ( min0 > r0) {
            return false;
        } else if (min1 > r1) {
            return false;
        }
        return true;
    }

}
