// SPDX-License-Identifier: MIT
pragma solidity >=0.7.6;
pragma abicoder v2;

interface IPriceAggregator {
    
    struct priceIn{
        uint index;
        address t0;
        address t1;
        address provider;
        uint8 providerVersion;
    }

    struct priceOut {
        uint index;
        uint256 price;
        uint24 fee;
    }

    struct existsIn{
        uint index;
        address t0;
        address t1;
        address provider;
        uint8 providerVersion;
        uint min0;
        uint min1;

    }

    struct existsOut {
        uint index;
        bool exists;
    } 

    function getPrices(priceIn[] memory inputs) external view returns (priceOut[] memory);
    function poolsExists(existsIn[] memory inputs) external view returns (existsOut[] memory);
}

