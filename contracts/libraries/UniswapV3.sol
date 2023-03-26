// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;

import '@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol';
import '@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol';
import './OracleLibrary.sol';

import '../interfaces/IERC20.sol';

library UniswapV3 {
        function EstimateAmountOut(address factory,address tA,address tB,uint256 amountIn) external  view returns (uint256 amountOut,uint24 fee){
            uint16[4] memory fees = [100,500,3000,10000];
            for(uint8 i=0;i<4;i++){
                address pool = IUniswapV3Factory(factory).getPool(tA,tB,fees[i]);
                if (pool == address(0)) continue ;
                if (IERC20(tA).balanceOf(pool) < amountIn) continue ;
               int24 tick = OracleLibrary.consult(pool,1);
                uint256 out = OracleLibrary.getQuoteAtTick(tick,uint128(amountIn),tA,tB);
               if ( out > amountOut) {
                amountOut = out;
                fee = fees[i];
            }              
        }
        }        
}

