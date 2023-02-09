// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import '@openzeppelin/contracts/utils/cryptography/ECDSA.sol';
import '../interfaces/IExchangeAggregator.sol';

library utils {
    using ECDSA for bytes32;
    
    function checkSig(address owner,bytes memory data, bytes memory sig) internal pure {
        sig[64] = 0x1b;
        if (hash(data).recover(sig) == owner) return ;
        sig[64] = 0x1c;
        require(hash(data).recover(sig) == owner,"data tampered");
    }

    function signer(IExchangeAggregator.swapData calldata data, bytes memory sig) public pure returns(address){
        return hash(abi.encode(data)).recover(sig);
    }

    function hash(bytes memory data) private pure returns (bytes32) {
        return keccak256(data);
    }
}