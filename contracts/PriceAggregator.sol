// SPDX-License-Identifier: MIT
pragma solidity 0.7.6;
pragma abicoder v2;

import '@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router01.sol';
import './interfaces/IPriceAggregator.sol';
import './libraries/UniswapV2.sol';
import './libraries/UniswapV3.sol';


contract PriceAggregator is IPriceAggregator {

    function getPrices(priceIn[] memory inputs) public override view returns (priceOut[] memory) {
        priceOut[] memory outputs = new priceOut[](inputs.length);
       for (uint i=0;i<inputs.length;i++){
        if (inputs[i].providerVersion == 2) {
            uint256 price = UniswapV2.Price(inputs[i].provider,inputs[i].t0,inputs[i].t1);
            outputs[i] = priceOut(inputs[i].index,price,0); 
        
        } else if (inputs[i].providerVersion == 3) {
            (uint256 price,uint24 fee) = UniswapV3.Price(inputs[i].provider,inputs[i].t0,inputs[i].t1);
            outputs[i] = priceOut(inputs[i].index,price,fee);            
        }
       }
    
       return outputs;
    }

    function poolsExists(existsIn[] memory inputs) external override view returns (existsOut[] memory){
        existsOut[] memory outputs = new existsOut[](inputs.length);
    for (uint i=0;i<inputs.length;i++){
        if (inputs[i].providerVersion == 2) {
            bool exists = UniswapV2.PoolExists(inputs[i].provider,inputs[i].t0,inputs[i].t1,inputs[i].min0,inputs[i].min1);
            outputs[i] = existsOut(inputs[i].index,exists);
        } if (inputs[i].providerVersion == 3) {
            bool exists = UniswapV3.PoolExists(inputs[i].provider,inputs[i].t0,inputs[i].t1,inputs[i].min0,inputs[i].min1);
            outputs[i] = existsOut(inputs[i].index,exists);
        }
                      
    }
    return outputs;
    }

}
