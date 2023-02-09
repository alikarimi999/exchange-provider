// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;

import '@uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol';
import '@uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol';
import './OracleLibrary.sol';

import '../interfaces/IERC20.sol';

library UniswapV3 {

        function Price(address factory,address t0,address t1) external view returns (uint256 price,uint24 fee){
        uint16[4] memory fees = [100,500,3000,10000];
        for(uint8 i=0;i<4;i++){
        address pool = IUniswapV3Factory(factory).getPool(t0,t1,fees[i]);
        if (pool == address(0)) continue ;
        
        uint8 d0 = IERC20(t0).decimals();
        uint8 d1 = IERC20(t1).decimals();
        
        if (IERC20(t0).balanceOf(pool) < 10**d0) continue ;
        if (IERC20(t1).balanceOf(pool) < 10**d1) continue ; 

        int24 tick = OracleLibrary.consult(pool,1);
        uint256 amountOut = OracleLibrary.getQuoteAtTick(tick,uint128(10**d0),t0,t1);
        if (i == 0){
            price = amountOut;
            fee = fees[i];
        }
        if (i > 0 && (price == 0 || amountOut < price)) {
            price = amountOut;
            fee = fees[i];
            }
        }
        return (price,fee);
    }


    function PoolExists(address factory,address t0,address t1,uint min0,uint min1) external view returns (bool){
        uint16[4] memory fees = [100,500,3000,10000];
        uint16[4] memory fs;
        for(uint8 i=0;i<4;i++){
           address pool = IUniswapV3Factory(factory).getPool(t0,t1,fees[i]);
           if (pool == address(0)) continue ;
           if (IERC20(t0).balanceOf(pool) < min0) continue ;
           if (IERC20(t1).balanceOf(pool) < min1) continue ; 
           return true;
        }
        return false;
    }

}

