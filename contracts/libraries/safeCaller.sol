// SPDX-License-Identifier: MIT
pragma solidity >=0.7.6;

library SafeCaller {
    function safeCall(
        address _contract,
        uint value,
        bytes memory data
    ) internal {
        (bool succeed,bytes memory result) = _contract.call{value: value}(data);
        if (!succeed) {
                if (result.length < 68) revert("ExchangeAggregator::SafeCaller:safeCall");
                assembly {
                    result := add(result, 0x04)
                }
                revert(abi.decode(result, (string)));
            }
    }
}