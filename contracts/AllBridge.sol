// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

import '@openzeppelin/contracts/access/Ownable.sol';
import './interfaces/IBridge.sol';
import './libraries/transferHelper.sol';
import './libraries/safeCaller.sol';
import './interfaces/IERC20.sol';

interface IAllBridge {
    enum MessengerProtocol {
        None,
        Allbridge,
        Wormhole,
        LayerZero
    }

    function swapAndBridge(
        bytes32 token,
        uint amount,
        bytes32 recipient,
        uint destinationChainId,
        bytes32 receiveToken,
        uint nonce,
        MessengerProtocol messenger,
        uint feeTokenAmount
    ) external payable;

}

contract ALLBridge is IBridge,Ownable {    
    struct allBridgeInput {
        address bridge;
        bytes32 tokenAddress;
        bytes32 recipient;
        uint destinationChainId;
        bytes32 receiveTokenAddress;
        uint nonce;
        IAllBridge.MessengerProtocol messenger;
        uint feeTokenAmount;
    }

      constructor(address[] memory tokenAddresses, address bridgeAddress) {
        // Approve smart contract for each token
        for (uint i = 0; i < tokenAddresses.length; i++) {
            approveToken(tokenAddresses[i], bridgeAddress);
        }
    }

//0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174   0x7775d63836987f444E2F14AA0fA2602204D7D3E0
    function Bridge(bytes calldata data,uint amountIn) public payable {
            allBridgeInput memory input = decodeInput(data);
            address tokenAddress = address(uint160(uint256(input.tokenAddress)));
            TransferHelper.safeTransferFrom(tokenAddress,msg.sender,address(this),amountIn);
            IAllBridge(input.bridge).swapAndBridge{value:msg.value}(
                input.tokenAddress,
                amountIn,
                input.recipient,
                input.destinationChainId,
                input.receiveTokenAddress,
                input.nonce,
                input.messenger,
                input.feeTokenAmount);
    }
    
    function approveToken(address token,address bridge) public onlyOwner {
        TransferHelper.safeApprove(token,bridge,type(uint256).max);
    }

    function allowance(address token,address bridge) public view returns(uint256){
        return IERC20(token).allowance(address(this),bridge);
    }

    function encode(allBridgeInput memory input) public pure returns (bytes memory) {
        return abi.encode(input);
    }

    function decodeInput(bytes calldata data) public pure returns (allBridgeInput memory) {
        allBridgeInput memory decodedInput;

        (decodedInput.bridge, decodedInput.tokenAddress, decodedInput.recipient, decodedInput.destinationChainId,
         decodedInput.receiveTokenAddress, decodedInput.nonce, decodedInput.messenger, decodedInput.feeTokenAmount) =
            abi.decode(data, (address, bytes32, bytes32, uint256, bytes32, uint256, IAllBridge.MessengerProtocol, uint256));

        return decodedInput;
    }
    
}