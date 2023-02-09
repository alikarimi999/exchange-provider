// SPDX-License-Identifier: MIT
pragma solidity >=0.7.6;

interface PriceProvider {
    function Price(address router, address t1,address t2) external returns (uint256);
}