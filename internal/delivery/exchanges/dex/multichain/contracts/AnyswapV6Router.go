// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"exchange-provider/pkg/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wNATIVE\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mpc\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txhash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapTradeTokensForNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapTradeTokensForTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMPC\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMPC\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"effectiveTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"LogChangeMPC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRouter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRouter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"LogChangeRouter\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"anySwapFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"txs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fromChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"anySwapIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapInAuto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapInExactTokensForNative\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapInExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txs\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapInUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"to\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"toChainIDs\",\"type\":\"uint256[]\"}],\"name\":\"anySwapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForNativeUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForNativeUnderlyingWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForNativeUnderlyingWithTransferPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForTokensUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForTokensUnderlyingWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutExactTokensForTokensUnderlyingWithTransferPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutUnderlyingWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutUnderlyingWithTransferPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMPC\",\"type\":\"address\"}],\"name\":\"changeMPC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"changeVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mpc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wNATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawNative\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b506040516200504d3803806200504d833981016040819052620000349162000095565b600180546001600160a01b039092166001600160a01b0319909216919091179055426002556001600160601b0319606092831b8116608052911b1660a052620000df565b80516001600160a01b03811681146200009057600080fd5b919050565b600080600060608486031215620000ab57600080fd5b620000b68462000078565b9250620000c66020850162000078565b9150620000d66040850162000078565b90509250925092565b60805160601c60a05160601c614eb6620001976000396000818161020d0152818161051601528181610820015281816108ef015281816116270152818161181001528181611cc401528181611d8501528181611e0b01528181611f7d015281816120d301528181612aa001528181612b610152612be70152600081816105e801528181610cbf01528181610e6401528181610f7c015281816116d30152818161302401528181613b0a0152613b5a0152614eb66000f3fe6080604052600436106101fd5760003560e01c8063832e94921161010d578063a5e56571116100a0578063d06ca61f1161006f578063d06ca61f1461062a578063d8b9f6101461064a578063dcfb77b11461066a578063edbdf5e21461068a578063f75c2664146106aa57600080fd5b8063a5e56571146105a3578063ad615dec146105b6578063c45a0155146105d6578063c8e174f61461060a57600080fd5b80638fd903f5116100dc5780638fd903f51461050457806399a2f2d71461055057806399cd84b5146105635780639aa1ac611461058357600080fd5b8063832e94921461048457806385f8c259146104a457806387cc6e2f146104c45780638d7d3eea146104e457600080fd5b80633f88de89116101905780635b7b018c1161015f5780635b7b018c146103f157806365782f56146104115780636a45397214610431578063701bb89114610451578063825bb13c1461046457600080fd5b80633f88de8914610361578063456862aa146103815780634d93bb94146103b157806352a397d5146103d157600080fd5b80631f00ca74116101cc5780631f00ca74146102d4578063241dc2df1461030157806325121b76146103215780632fc1e7281461034157600080fd5b80630175b1c414610241578063054d50d4146102615780630bb57203146102945780631b91a934146102b457600080fd5b3661023c57336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461023a5761023a614df0565b005b600080fd5b34801561024d57600080fd5b5061023a61025c36600461477d565b6106bf565b34801561026d57600080fd5b5061028161027c366004614a07565b6109ed565b6040519081526020015b60405180910390f35b3480156102a057600080fd5b5061023a6102af36600461498a565b610a02565b3480156102c057600080fd5b5061023a6102cf366004614371565b610b22565b3480156102e057600080fd5b506102f46102ef3660046148b2565b610cb8565b60405161028b9190614b9e565b34801561030d57600080fd5b5061023a61031c36600461443d565b610cee565b34801561032d57600080fd5b5061023a61033c36600461466a565b610d01565b34801561034d57600080fd5b506102f461035c3660046147cf565b610e04565b34801561036d57600080fd5b5061023a61037c36600461477d565b61105d565b34801561038d57600080fd5b506103a161039c366004614338565b61112a565b604051901515815260200161028b565b3480156103bd57600080fd5b5061023a6103cc3660046144f1565b611235565b3480156103dd57600080fd5b506102f46103ec3660046147cf565b6115c2565b3480156103fd57600080fd5b506103a161040c3660046142f7565b6118e4565b34801561041d57600080fd5b5061023a61042c36600461498a565b611a0c565b34801561043d57600080fd5b5061023a61044c36600461498a565b611b0b565b61028161045f366004614338565b611cc0565b34801561047057600080fd5b5061023a61047f36600461477d565b611f2d565b34801561049057600080fd5b5061028161049f3660046144af565b611f79565b3480156104b057600080fd5b506102816104bf366004614a07565b612149565b3480156104d057600080fd5b5061023a6104df366004614483565b612156565b3480156104f057600080fd5b5061023a6104ff366004614371565b61229c565b34801561051057600080fd5b506105387f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161028b565b34801561055c57600080fd5b5046610281565b34801561056f57600080fd5b5061023a61057e3660046144f1565b61242e565b34801561058f57600080fd5b5061023a61059e3660046144f1565b612776565b61023a6105b13660046143fc565b612a9e565b3480156105c257600080fd5b506102816105d1366004614a07565b612d11565b3480156105e257600080fd5b506105387f000000000000000000000000000000000000000000000000000000000000000081565b34801561061657600080fd5b5061023a6106253660046144f1565b612d1e565b34801561063657600080fd5b506102f46106453660046148b2565b61301d565b34801561065657600080fd5b5061023a61066536600461498a565b61304a565b34801561067657600080fd5b5061023a6106853660046145a6565b613140565b34801561069657600080fd5b5061023a6106a536600461443d565b6131e6565b3480156106b657600080fd5b506105386132b3565b6106c76132b3565b6001600160a01b0316336001600160a01b0316146107005760405162461bcd60e51b81526004016106f790614c64565b60405180910390fd5b61070d85858585856132de565b60008490506000816001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b15801561074d57600080fd5b505afa158015610761573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610785919061431b565b90506001600160a01b0381161580159061081957506040516370a0823160e01b81526001600160a01b0387811660048301528591908316906370a082319060240160206040518083038186803b1580156107de57600080fd5b505afa1580156107f2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108169190614899565b10155b156109e4577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316141561096257604051620e75bb60e21b81526001600160a01b038316906239d6ec9061088690889088903090600401614aa9565b602060405180830381600087803b1580156108a057600080fd5b505af11580156108b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108d89190614899565b50604051632e1a7d4d60e01b8152600481018590527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561093b57600080fd5b505af115801561094f573d6000803e3d6000fd5b5050505061095d85856133c4565b6109e4565b604051620e75bb60e21b81526001600160a01b038316906239d6ec9061099090889088908290600401614aa9565b602060405180830381600087803b1580156109aa57600080fd5b505af11580156109be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e29190614899565b505b50505050505050565b60006109fa848484613490565b949350505050565b8142811015610a235760405162461bcd60e51b81526004016106f790614c9b565b85856000818110610a3657610a36614e1c565b9050602002016020810190610a4b91906142f7565b6001600160a01b0316639dc29fac338a6040518363ffffffff1660e01b8152600401610a78929190614a90565b602060405180830381600087803b158015610a9257600080fd5b505af1158015610aa6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aca919061475b565b506001600160a01b038416337ffea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b388888c8c465b89604051610b1096959493929190614acc565b60405180910390a35050505050505050565b876001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015610b5b57600080fd5b505afa158015610b6f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b93919061431b565b6001600160a01b031663605629d68a8a89898989896040518863ffffffff1660e01b8152600401610bca9796959493929190614a4f565b602060405180830381600087803b158015610be457600080fd5b505af1158015610bf8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1c919061475b565b50604051630bebbf4d60e41b8152600481018790526001600160a01b038a8116602483015289169063bebbf4d090604401602060405180830381600087803b158015610c6757600080fd5b505af1158015610c7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c9f9190614899565b50610cad898989898561355e565b505050505050505050565b6060610ce57f00000000000000000000000000000000000000000000000000000000000000008484613642565b90505b92915050565b610cfb338585858561355e565b50505050565b610d096132b3565b6001600160a01b0316336001600160a01b031614610d395760405162461bcd60e51b81526004016106f790614c64565b60005b87811015610df757610de58b8b83818110610d5957610d59614e1c565b905060200201358a8a84818110610d7257610d72614e1c565b9050602002016020810190610d8791906142f7565b898985818110610d9957610d99614e1c565b9050602002016020810190610dae91906142f7565b888886818110610dc057610dc0614e1c565b90506020020135878787818110610dd957610dd9614e1c565b905060200201356132de565b80610def81614dd5565b915050610d3c565b5050505050505050505050565b6060610e0e6132b3565b6001600160a01b0316336001600160a01b031614610e3e5760405162461bcd60e51b81526004016106f790614c64565b8242811015610e5f5760405162461bcd60e51b81526004016106f790614c9b565b610ebd7f00000000000000000000000000000000000000000000000000000000000000008a8989808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506137d892505050565b9150878260018451610ecf9190614d7b565b81518110610edf57610edf614e1c565b60200260200101511015610f4b5760405162461bcd60e51b815260206004820152602d60248201527f5375736869737761705632526f757465723a20494e53554646494349454e545f60448201526c13d55514155517d05353d55395609a1b60648201526084016106f7565b6110118a88886000818110610f6257610f62614e1c565b9050602002016020810190610f7791906142f7565b610ff07f00000000000000000000000000000000000000000000000000000000000000008b8b6000818110610fae57610fae614e1c565b9050602002016020810190610fc391906142f7565b8c8c6001818110610fd657610fd6614e1c565b9050602002016020810190610feb91906142f7565b61394e565b8560008151811061100357611003614e1c565b6020026020010151876132de565b611050828888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250613a27915050565b5098975050505050505050565b6110656132b3565b6001600160a01b0316336001600160a01b0316146110955760405162461bcd60e51b81526004016106f790614c64565b6110a285858585856132de565b604051620e75bb60e21b81526001600160a01b038516906239d6ec906110d090869086908290600401614aa9565b602060405180830381600087803b1580156110ea57600080fd5b505af11580156110fe573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111229190614899565b505050505050565b60006111346132b3565b6001600160a01b0316336001600160a01b0316146111645760405162461bcd60e51b81526004016106f790614c64565b6001600160a01b0382166111ba5760405162461bcd60e51b815260206004820152601d60248201527f416e79737761705633526f757465723a2061646472657373283078302900000060448201526064016106f7565b6040516360e232a960e01b81526001600160a01b0383811660048301528416906360e232a990602401602060405180830381600087803b1580156111fd57600080fd5b505af1158015611211573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce5919061475b565b84428110156112565760405162461bcd60e51b81526004016106f790614c9b565b60008989600081811061126b5761126b614e1c565b905060200201602081019061128091906142f7565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156112b857600080fd5b505afa1580156112cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112f0919061431b565b9050806001600160a01b031663d505accf8e308f8b8b8b8b6040518863ffffffff1660e01b815260040161132a9796959493929190614a4f565b600060405180830381600087803b15801561134457600080fd5b505af1158015611358573d6000803e3d6000fd5b5050505061139a8d8b8b600081811061137357611373614e1c565b905060200201602081019061138891906142f7565b6001600160a01b03841691908f613c29565b898960008181106113ad576113ad614e1c565b90506020020160208101906113c291906142f7565b604051630bebbf4d60e41b8152600481018e90526001600160a01b038f81166024830152919091169063bebbf4d090604401602060405180830381600087803b15801561140e57600080fd5b505af1158015611422573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114469190614899565b508989600081811061145a5761145a614e1c565b905060200201602081019061146f91906142f7565b6001600160a01b0316639dc29fac8e8e6040518363ffffffff1660e01b815260040161149c929190614a90565b602060405180830381600087803b1580156114b657600080fd5b505af11580156114ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ee919061475b565b5060008a8a80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050905060008e905060008a905060008f905060008f905060006115524690565b90506000899050846001600160a01b0316866001600160a01b03167f278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f489878787876040516115a4959493929190614b36565b60405180910390a35050505050505050505050505050505050505050565b60606115cc6132b3565b6001600160a01b0316336001600160a01b0316146115fc5760405162461bcd60e51b81526004016106f790614c64565b824281101561161d5760405162461bcd60e51b81526004016106f790614c9b565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168787611654600182614d7b565b81811061166357611663614e1c565b905060200201602081019061167891906142f7565b6001600160a01b0316146116ce5760405162461bcd60e51b815260206004820152601d60248201527f416e79737761705633526f757465723a20494e56414c49445f5041544800000060448201526064016106f7565b61172c7f00000000000000000000000000000000000000000000000000000000000000008a8989808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506137d892505050565b915087826001845161173e9190614d7b565b8151811061174e5761174e614e1c565b602002602001015110156117b85760405162461bcd60e51b815260206004820152602b60248201527f416e79737761705633526f757465723a20494e53554646494349454e545f4f5560448201526a1514155517d05353d5539560aa1b60648201526084016106f7565b6117cf8a88886000818110610f6257610f62614e1c565b61180e82888880806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250309250613a27915050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316632e1a7d4d836001855161184c9190614d7b565b8151811061185c5761185c614e1c565b60200260200101516040518263ffffffff1660e01b815260040161188291815260200190565b600060405180830381600087803b15801561189c57600080fd5b505af11580156118b0573d6000803e3d6000fd5b505050506110508583600185516118c79190614d7b565b815181106118d7576118d7614e1c565b60200260200101516133c4565b60006118ee6132b3565b6001600160a01b0316336001600160a01b03161461191e5760405162461bcd60e51b81526004016106f790614c64565b6001600160a01b0382166119745760405162461bcd60e51b815260206004820152601d60248201527f416e79737761705633526f757465723a2061646472657373283078302900000060448201526064016106f7565b61197c6132b3565b600080546001600160a01b03199081166001600160a01b0393841617909155600180549091169184169190911790556119b8426202a300614d22565b60028190556001546000546001600160a01b0391821691167fcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a21981214660405190815260200160405180910390a45060015b919050565b8142811015611a2d5760405162461bcd60e51b81526004016106f790614c9b565b85856000818110611a4057611a40614e1c565b9050602002016020810190611a5591906142f7565b6001600160a01b0316639dc29fac338a6040518363ffffffff1660e01b8152600401611a82929190614a90565b602060405180830381600087803b158015611a9c57600080fd5b505af1158015611ab0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ad4919061475b565b506001600160a01b038416337f278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f488888c8c46610afd565b8142811015611b2c5760405162461bcd60e51b81526004016106f790614c9b565b611c023387876000818110611b4357611b43614e1c565b9050602002016020810190611b5891906142f7565b8a89896000818110611b6c57611b6c614e1c565b9050602002016020810190611b8191906142f7565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015611bb957600080fd5b505afa158015611bcd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bf1919061431b565b6001600160a01b0316929190613c29565b85856000818110611c1557611c15614e1c565b9050602002016020810190611c2a91906142f7565b604051630bebbf4d60e41b8152600481018a90523360248201526001600160a01b03919091169063bebbf4d090604401602060405180830381600087803b158015611c7457600080fd5b505af1158015611c88573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cac9190614899565b5085856000818110611a4057611a40614e1c565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015611d2557600080fd5b505afa158015611d39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d5d919061431b565b6001600160a01b031614611d835760405162461bcd60e51b81526004016106f790614c1a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b158015611dde57600080fd5b505af1158015611df2573d6000803e3d6000fd5b505060405163a9059cbb60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016935063a9059cbb9250611e45915086903490600401614a90565b602060405180830381600087803b158015611e5f57600080fd5b505af1158015611e73573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e97919061475b565b611ea357611ea3614df0565b604051630bebbf4d60e41b81523460048201526001600160a01b03838116602483015284169063bebbf4d090604401602060405180830381600087803b158015611eec57600080fd5b505af1158015611f00573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f249190614899565b50349392505050565b611f356132b3565b6001600160a01b0316336001600160a01b031614611f655760405162461bcd60e51b81526004016106f790614c64565b611f7285858585856132de565b5050505050565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316846001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015611fde57600080fd5b505afa158015611ff2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612016919061431b565b6001600160a01b03161461203c5760405162461bcd60e51b81526004016106f790614c1a565b604051620e75bb60e21b81526001600160a01b038516906239d6ec9061206a90339087903090600401614aa9565b602060405180830381600087803b15801561208457600080fd5b505af1158015612098573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120bc9190614899565b50604051632e1a7d4d60e01b8152600481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561211f57600080fd5b505af1158015612133573d6000803e3d6000fd5b5050505061214182846133c4565b509092915050565b60006109fa848484613c83565b61215e6132b3565b6001600160a01b0316336001600160a01b03161461218e5760405162461bcd60e51b81526004016106f790614c64565b60006121986132b3565b6040516340c10f1960e01b81529091506001600160a01b038416906340c10f19906121c99084908690600401614a90565b602060405180830381600087803b1580156121e357600080fd5b505af11580156121f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061221b919061475b565b50604051620e75bb60e21b81526001600160a01b038416906239d6ec9061224a90849086908290600401614aa9565b602060405180830381600087803b15801561226457600080fd5b505af1158015612278573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cfb9190614899565b6000886001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156122d757600080fd5b505afa1580156122eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061230f919061431b565b60405163d505accf60e01b81529091506001600160a01b0382169063d505accf9061234a908d9030908c908c908c908c908c90600401614a4f565b600060405180830381600087803b15801561236457600080fd5b505af1158015612378573d6000803e3d6000fd5b50612392925050506001600160a01b0382168b8b8a613c29565b604051630bebbf4d60e41b8152600481018890526001600160a01b038b811660248301528a169063bebbf4d090604401602060405180830381600087803b1580156123dc57600080fd5b505af11580156123f0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124149190614899565b506124228a8a8a8a8661355e565b50505050505050505050565b844281101561244f5760405162461bcd60e51b81526004016106f790614c9b565b60008989600081811061246457612464614e1c565b905060200201602081019061247991906142f7565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156124b157600080fd5b505afa1580156124c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124e9919061431b565b9050806001600160a01b031663d505accf8e308f8b8b8b8b6040518863ffffffff1660e01b81526004016125239796959493929190614a4f565b600060405180830381600087803b15801561253d57600080fd5b505af1158015612551573d6000803e3d6000fd5b5050505061256c8d8b8b600081811061137357611373614e1c565b8989600081811061257f5761257f614e1c565b905060200201602081019061259491906142f7565b604051630bebbf4d60e41b8152600481018e90526001600160a01b038f81166024830152919091169063bebbf4d090604401602060405180830381600087803b1580156125e057600080fd5b505af11580156125f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126189190614899565b508989600081811061262c5761262c614e1c565b905060200201602081019061264191906142f7565b6001600160a01b0316639dc29fac8e8e6040518363ffffffff1660e01b815260040161266e929190614a90565b602060405180830381600087803b15801561268857600080fd5b505af115801561269c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126c0919061475b565b5060008a8a80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050905060008e905060008a905060008f905060008f905060006127244690565b90506000899050846001600160a01b0316866001600160a01b03167ffea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b389878787876040516115a4959493929190614b36565b84428110156127975760405162461bcd60e51b81526004016106f790614c9b565b888860008181106127aa576127aa614e1c565b90506020020160208101906127bf91906142f7565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156127f757600080fd5b505afa15801561280b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061282f919061431b565b6001600160a01b031663605629d68d8b8b600081811061285157612851614e1c565b905060200201602081019061286691906142f7565b8e8a8a8a8a6040518863ffffffff1660e01b815260040161288d9796959493929190614a4f565b602060405180830381600087803b1580156128a757600080fd5b505af11580156128bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128df919061475b565b50888860008181106128f3576128f3614e1c565b905060200201602081019061290891906142f7565b604051630bebbf4d60e41b8152600481018d90526001600160a01b038e81166024830152919091169063bebbf4d090604401602060405180830381600087803b15801561295457600080fd5b505af1158015612968573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061298c9190614899565b50888860008181106129a0576129a0614e1c565b90506020020160208101906129b591906142f7565b6001600160a01b0316639dc29fac8d8d6040518363ffffffff1660e01b81526004016129e2929190614a90565b602060405180830381600087803b1580156129fc57600080fd5b505af1158015612a10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a34919061475b565b50866001600160a01b03168c6001600160a01b03167ffea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b38b8b8f8f612a754690565b89604051612a8896959493929190614acc565b60405180910390a3505050505050505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015612b0157600080fd5b505afa158015612b15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b39919061431b565b6001600160a01b031614612b5f5760405162461bcd60e51b81526004016106f790614c1a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b158015612bba57600080fd5b505af1158015612bce573d6000803e3d6000fd5b505060405163a9059cbb60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016935063a9059cbb9250612c21915086903490600401614a90565b602060405180830381600087803b158015612c3b57600080fd5b505af1158015612c4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c73919061475b565b612c7f57612c7f614df0565b604051630bebbf4d60e41b81523460048201523360248201526001600160a01b0384169063bebbf4d090604401602060405180830381600087803b158015612cc657600080fd5b505af1158015612cda573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612cfe9190614899565b50612d0c338484348561355e565b505050565b60006109fa848484613d4c565b8442811015612d3f5760405162461bcd60e51b81526004016106f790614c9b565b88886000818110612d5257612d52614e1c565b9050602002016020810190612d6791906142f7565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015612d9f57600080fd5b505afa158015612db3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dd7919061431b565b6001600160a01b031663605629d68d8b8b6000818110612df957612df9614e1c565b9050602002016020810190612e0e91906142f7565b8e8a8a8a8a6040518863ffffffff1660e01b8152600401612e359796959493929190614a4f565b602060405180830381600087803b158015612e4f57600080fd5b505af1158015612e63573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e87919061475b565b5088886000818110612e9b57612e9b614e1c565b9050602002016020810190612eb091906142f7565b604051630bebbf4d60e41b8152600481018d90526001600160a01b038e81166024830152919091169063bebbf4d090604401602060405180830381600087803b158015612efc57600080fd5b505af1158015612f10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f349190614899565b5088886000818110612f4857612f48614e1c565b9050602002016020810190612f5d91906142f7565b6001600160a01b0316639dc29fac8d8d6040518363ffffffff1660e01b8152600401612f8a929190614a90565b602060405180830381600087803b158015612fa457600080fd5b505af1158015612fb8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fdc919061475b565b50866001600160a01b03168c6001600160a01b03167f278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f48b8b8f8f612a754690565b6060610ce57f000000000000000000000000000000000000000000000000000000000000000084846137d8565b814281101561306b5760405162461bcd60e51b81526004016106f790614c9b565b6130823387876000818110611b4357611b43614e1c565b8585600081811061309557613095614e1c565b90506020020160208101906130aa91906142f7565b604051630bebbf4d60e41b8152600481018a90523360248201526001600160a01b03919091169063bebbf4d090604401602060405180830381600087803b1580156130f457600080fd5b505af1158015613108573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061312c9190614899565b5085856000818110610a3657610a36614e1c565b60005b87811015610cad576131d4338a8a8481811061316157613161614e1c565b905060200201602081019061317691906142f7565b89898581811061318857613188614e1c565b905060200201602081019061319d91906142f7565b8888868181106131af576131af614e1c565b905060200201358787878181106131c8576131c8614e1c565b9050602002013561355e565b806131de81614dd5565b915050613143565b613225338584876001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015611bb957600080fd5b604051630bebbf4d60e41b8152600481018390523360248201526001600160a01b0385169063bebbf4d090604401602060405180830381600087803b15801561326d57600080fd5b505af1158015613281573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132a59190614899565b50610cfb338585858561355e565b600060025442106132ce57506001546001600160a01b031690565b506000546001600160a01b031690565b6040516340c10f1960e01b81526001600160a01b038516906340c10f199061330c9086908690600401614a90565b602060405180830381600087803b15801561332657600080fd5b505af115801561333a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061335e919061475b565b50826001600160a01b0316846001600160a01b0316867faac9ce45fe3adf5143598c4f18a369591a20a3384aedaf1b525d29127e1fcd55858561339e4690565b604080519384526020840192909252908201526060015b60405180910390a45050505050565b604080516000808252602082019092526001600160a01b0384169083906040516133ee9190614a33565b60006040518083038185875af1925050503d806000811461342b576040519150601f19603f3d011682016040523d82523d6000602084013e613430565b606091505b5050905080612d0c5760405162461bcd60e51b815260206004820152602660248201527f5472616e7366657248656c7065723a204e41544956455f5452414e534645525f60448201526511905253115160d21b60648201526084016106f7565b60008084116134e55760405162461bcd60e51b815260206004820152602d6024820152600080516020614e6183398151915260448201526c17d25394155517d05353d55395609a1b60648201526084016106f7565b6000831180156134f55750600082115b6135115760405162461bcd60e51b81526004016106f790614be2565b600061351f856103e5613ddc565b9050600061352d8285613ddc565b9050600061354783613541886103e8613ddc565b90613e43565b90506135538183614d3a565b979650505050505050565b604051632770a7eb60e21b81526001600160a01b03851690639dc29fac9061358c9088908690600401614a90565b602060405180830381600087803b1580156135a657600080fd5b505af11580156135ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135de919061475b565b50826001600160a01b0316856001600160a01b0316856001600160a01b03167f97116cf6cd4f6412bb47914d6db18da9e16ab2142f543b86e207c24fbd16b23a856136264690565b60408051928352602083019190915281018690526060016133b5565b60606002825110156136965760405162461bcd60e51b815260206004820181905260248201527f53757368697377617056324c6962726172793a20494e56414c49445f5041544860448201526064016106f7565b815167ffffffffffffffff8111156136b0576136b0614e32565b6040519080825280602002602001820160405280156136d9578160200160208202803683370190505b5090508281600183516136ec9190614d7b565b815181106136fc576136fc614e1c565b6020026020010181815250506000600183516137189190614d7b565b90505b80156137d05760008061376b8786613734600187614d7b565b8151811061374457613744614e1c565b602002602001015187868151811061375e5761375e614e1c565b6020026020010151613e98565b9150915061379384848151811061378457613784614e1c565b60200260200101518383613c83565b8461379f600186614d7b565b815181106137af576137af614e1c565b602002602001018181525050505080806137c890614dbe565b91505061371b565b509392505050565b606060028251101561382c5760405162461bcd60e51b815260206004820181905260248201527f53757368697377617056324c6962726172793a20494e56414c49445f5041544860448201526064016106f7565b815167ffffffffffffffff81111561384657613846614e32565b60405190808252806020026020018201604052801561386f578160200160208202803683370190505b509050828160008151811061388657613886614e1c565b60200260200101818152505060005b600183516138a39190614d7b565b8110156137d0576000806138e9878685815181106138c3576138c3614e1c565b6020026020010151878660016138d99190614d22565b8151811061375e5761375e614e1c565b9150915061391184848151811061390257613902614e1c565b60200260200101518383613490565b8461391d856001614d22565b8151811061392d5761392d614e1c565b6020026020010181815250505050808061394690614dd5565b915050613895565b600080600061395d8585613f71565b6040516bffffffffffffffffffffffff19606084811b8216602084015283901b1660348201529193509150869060480160405160208183030381529060405280519060200120604051602001613a059291906001600160f81b0319815260609290921b6bffffffffffffffffffffffff1916600183015260158201527fe18a34eb0e04b04f7a0ac29a6e80748dca96319b42c54d679cb821dca90c6303603582015260550190565b60408051601f1981840301815291905280516020909101209695505050505050565b60005b60018351613a389190614d7b565b811015610cfb57600080848381518110613a5457613a54614e1c565b602002602001015185846001613a6a9190614d22565b81518110613a7a57613a7a614e1c565b6020026020010151915091506000613a928383613f71565b509050600087613aa3866001614d22565b81518110613ab357613ab3614e1c565b60200260200101519050600080836001600160a01b0316866001600160a01b031614613ae157826000613ae5565b6000835b91509150600060028a51613af99190614d7b565b8810613b055788613b53565b613b537f0000000000000000000000000000000000000000000000000000000000000000878c613b368c6002614d22565b81518110613b4657613b46614e1c565b602002602001015161394e565b9050613b807f0000000000000000000000000000000000000000000000000000000000000000888861394e565b6001600160a01b031663022c0d9f84848460006040519080825280601f01601f191660200182016040528015613bbd576020820181803683370190505b506040518563ffffffff1660e01b8152600401613bdd9493929190614cd2565b600060405180830381600087803b158015613bf757600080fd5b505af1158015613c0b573d6000803e3d6000fd5b50505050505050505050508080613c2190614dd5565b915050613a2a565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052610cfb90859061406b565b6000808411613cd95760405162461bcd60e51b815260206004820152602e6024820152600080516020614e6183398151915260448201526d17d3d55514155517d05353d5539560921b60648201526084016106f7565b600083118015613ce95750600082115b613d055760405162461bcd60e51b81526004016106f790614be2565b6000613d1d6103e8613d178688613ddc565b90613ddc565b90506000613d316103e5613d1786896141f2565b9050613d4260016135418385614d3a565b9695505050505050565b6000808411613d9b5760405162461bcd60e51b81526020600482015260276024820152600080516020614e6183398151915260448201526617d05353d5539560ca1b60648201526084016106f7565b600083118015613dab5750600082115b613dc75760405162461bcd60e51b81526004016106f790614be2565b82613dd28584613ddc565b6109fa9190614d3a565b6000811580613e0057508282613df28183614d5c565b9250613dfe9083614d3a565b145b610ce85760405162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6d756c2d6f766572666c6f7760601b60448201526064016106f7565b600082613e508382614d22565b9150811015610ce85760405162461bcd60e51b815260206004820152601460248201527364732d6d6174682d6164642d6f766572666c6f7760601b60448201526064016106f7565b6000806000613ea78585613f71565b509050600080613eb888888861394e565b6001600160a01b0316630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b158015613ef057600080fd5b505afa158015613f04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f289190614854565b506001600160701b031691506001600160701b03169150826001600160a01b0316876001600160a01b031614613f5f578082613f62565b81815b90999098509650505050505050565b600080826001600160a01b0316846001600160a01b03161415613fe65760405162461bcd60e51b815260206004820152602760248201527f53757368697377617056324c6962726172793a204944454e544943414c5f41446044820152664452455353455360c81b60648201526084016106f7565b826001600160a01b0316846001600160a01b031610614006578284614009565b83835b90925090506001600160a01b0382166140645760405162461bcd60e51b815260206004820181905260248201527f53757368697377617056324c6962726172793a205a45524f5f4144445245535360448201526064016106f7565b9250929050565b61407d826001600160a01b0316614248565b6140c95760405162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e74726163740060448201526064016106f7565b600080836001600160a01b0316836040516140e49190614a33565b6000604051808303816000865af19150503d8060008114614121576040519150601f19603f3d011682016040523d82523d6000602084013e614126565b606091505b5091509150816141785760405162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656460448201526064016106f7565b805115610cfb5780806020019051810190614193919061475b565b610cfb5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016106f7565b6000826141ff8382614d7b565b9150811115610ce85760405162461bcd60e51b815260206004820152601560248201527464732d6d6174682d7375622d756e646572666c6f7760581b60448201526064016106f7565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47081158015906109fa5750141592915050565b8035611a0781614e48565b60008083601f84011261429c57600080fd5b50813567ffffffffffffffff8111156142b457600080fd5b6020830191508360208260051b850101111561406457600080fd5b80516001600160701b0381168114611a0757600080fd5b803560ff81168114611a0757600080fd5b60006020828403121561430957600080fd5b813561431481614e48565b9392505050565b60006020828403121561432d57600080fd5b815161431481614e48565b6000806040838503121561434b57600080fd5b823561435681614e48565b9150602083013561436681614e48565b809150509250929050565b60008060008060008060008060006101208a8c03121561439057600080fd5b893561439b81614e48565b985060208a01356143ab81614e48565b975060408a01356143bb81614e48565b965060608a0135955060808a013594506143d760a08b016142e6565b935060c08a0135925060e08a013591506101008a013590509295985092959850929598565b60008060006060848603121561441157600080fd5b833561441c81614e48565b9250602084013561442c81614e48565b929592945050506040919091013590565b6000806000806080858703121561445357600080fd5b843561445e81614e48565b9350602085013561446e81614e48565b93969395505050506040820135916060013590565b6000806040838503121561449657600080fd5b82356144a181614e48565b946020939093013593505050565b6000806000606084860312156144c457600080fd5b83356144cf81614e48565b92506020840135915060408401356144e681614e48565b809150509250925092565b60008060008060008060008060008060006101408c8e03121561451357600080fd5b8b3561451e81614e48565b9a5060208c0135995060408c0135985060608c013567ffffffffffffffff81111561454857600080fd5b6145548e828f0161428a565b90995097505060808c013561456881614e48565b955060a08c0135945061457d60c08d016142e6565b935060e08c013592506101008c013591506101208c013590509295989b509295989b9093969950565b6000806000806000806000806080898b0312156145c257600080fd5b883567ffffffffffffffff808211156145da57600080fd5b6145e68c838d0161428a565b909a50985060208b01359150808211156145ff57600080fd5b61460b8c838d0161428a565b909850965060408b013591508082111561462457600080fd5b6146308c838d0161428a565b909650945060608b013591508082111561464957600080fd5b506146568b828c0161428a565b999c989b5096995094979396929594505050565b60008060008060008060008060008060a08b8d03121561468957600080fd5b8a3567ffffffffffffffff808211156146a157600080fd5b6146ad8e838f0161428a565b909c509a5060208d01359150808211156146c657600080fd5b6146d28e838f0161428a565b909a50985060408d01359150808211156146eb57600080fd5b6146f78e838f0161428a565b909850965060608d013591508082111561471057600080fd5b61471c8e838f0161428a565b909650945060808d013591508082111561473557600080fd5b506147428d828e0161428a565b915080935050809150509295989b9194979a5092959850565b60006020828403121561476d57600080fd5b8151801515811461431457600080fd5b600080600080600060a0868803121561479557600080fd5b8535945060208601356147a781614e48565b935060408601356147b781614e48565b94979396509394606081013594506080013592915050565b60008060008060008060008060e0898b0312156147eb57600080fd5b883597506020890135965060408901359550606089013567ffffffffffffffff81111561481757600080fd5b6148238b828c0161428a565b909650945050608089013561483781614e48565b979a969950949793969295929450505060a08201359160c0013590565b60008060006060848603121561486957600080fd5b614872846142cf565b9250614880602085016142cf565b9150604084015163ffffffff811681146144e657600080fd5b6000602082840312156148ab57600080fd5b5051919050565b600080604083850312156148c557600080fd5b8235915060208084013567ffffffffffffffff808211156148e557600080fd5b818601915086601f8301126148f957600080fd5b81358181111561490b5761490b614e32565b8060051b604051601f19603f8301168101818110858211171561493057614930614e32565b604052828152858101935084860182860187018b101561494f57600080fd5b600095505b83861015614979576149658161427f565b855260019590950194938601938601614954565b508096505050505050509250929050565b600080600080600080600060c0888a0312156149a557600080fd5b8735965060208801359550604088013567ffffffffffffffff8111156149ca57600080fd5b6149d68a828b0161428a565b90965094505060608801356149ea81614e48565b969995985093969295946080840135945060a09093013592915050565b600080600060608486031215614a1c57600080fd5b505081359360208301359350604090920135919050565b60008251614a45818460208701614d92565b9190910192915050565b6001600160a01b0397881681529590961660208601526040850193909352606084019190915260ff16608083015260a082015260c081019190915260e00190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0393841681526020810192909252909116604082015260600190565b60a0808252810186905260008760c08301825b89811015614b0f578235614af281614e48565b6001600160a01b0316825260209283019290910190600101614adf565b50602084019790975250506040810193909352606083019190915260809091015292915050565b60a0808252865190820181905260009060209060c0840190828a01845b82811015614b785781516001600160a01b031684529284019290840190600101614b53565b505050908301969096525060408101939093526060830191909152608090910152919050565b6020808252825182820181905260009190848201906040850190845b81811015614bd657835183529284019291840191600101614bba565b50909695505050505050565b6020808252602a90820152600080516020614e618339815191526040820152695f4c495155494449545960b01b606082015260800190565b6020808252602a908201527f416e79737761705633526f757465723a20756e6465726c79696e67206973206e6040820152696f7420774e415449564560b01b606082015260800190565b6020808252601a908201527f416e79737761705633526f757465723a20464f5242494444454e000000000000604082015260600190565b60208082526018908201527f416e79737761705633526f757465723a20455850495245440000000000000000604082015260600190565b84815283602082015260018060a01b03831660408201526080606082015260008251806080840152614d0b8160a0850160208701614d92565b601f01601f19169190910160a00195945050505050565b60008219821115614d3557614d35614e06565b500190565b600082614d5757634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615614d7657614d76614e06565b500290565b600082821015614d8d57614d8d614e06565b500390565b60005b83811015614dad578181015183820152602001614d95565b83811115610cfb5750506000910152565b600081614dcd57614dcd614e06565b506000190190565b6000600019821415614de957614de9614e06565b5060010190565b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114614e5d57600080fd5b5056fe53757368697377617056324c6962726172793a20494e53554646494349454e54a264697066735822122073407528f159bfea36d084e49afe223dc0f542434e791cc076c52d87bb5c136a64736f6c63430008060033000000000000000000000000c0aee478e3658e2610c5f7a4a2e1777ce9e4f2ac000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000f39fee2fdfe7db022591f4a82e3537fa0b55fb9c",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _factory common.Address, _wNATIVE common.Address, _mpc common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _factory, _wNATIVE, _mpc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// CID is a free data retrieval call binding the contract method 0x99a2f2d7.
//
// Solidity: function cID() view returns(uint256 id)
func (_Contracts *ContractsCaller) CID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "cID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CID is a free data retrieval call binding the contract method 0x99a2f2d7.
//
// Solidity: function cID() view returns(uint256 id)
func (_Contracts *ContractsSession) CID() (*big.Int, error) {
	return _Contracts.Contract.CID(&_Contracts.CallOpts)
}

// CID is a free data retrieval call binding the contract method 0x99a2f2d7.
//
// Solidity: function cID() view returns(uint256 id)
func (_Contracts *ContractsCallerSession) CID() (*big.Int, error) {
	return _Contracts.Contract.CID(&_Contracts.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Contracts *ContractsCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Contracts *ContractsSession) Factory() (common.Address, error) {
	return _Contracts.Contract.Factory(&_Contracts.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Contracts *ContractsCallerSession) Factory() (common.Address, error) {
	return _Contracts.Contract.Factory(&_Contracts.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Contracts *ContractsCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Contracts *ContractsSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAmountIn(&_Contracts.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Contracts *ContractsCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAmountIn(&_Contracts.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Contracts *ContractsCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Contracts *ContractsSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAmountOut(&_Contracts.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Contracts *ContractsCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetAmountOut(&_Contracts.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Contracts *ContractsCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Contracts *ContractsSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Contracts.Contract.GetAmountsIn(&_Contracts.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Contracts *ContractsCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Contracts.Contract.GetAmountsIn(&_Contracts.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Contracts *ContractsCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Contracts *ContractsSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Contracts.Contract.GetAmountsOut(&_Contracts.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Contracts *ContractsCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Contracts.Contract.GetAmountsOut(&_Contracts.CallOpts, amountIn, path)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Contracts *ContractsCaller) Mpc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "mpc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Contracts *ContractsSession) Mpc() (common.Address, error) {
	return _Contracts.Contract.Mpc(&_Contracts.CallOpts)
}

// Mpc is a free data retrieval call binding the contract method 0xf75c2664.
//
// Solidity: function mpc() view returns(address)
func (_Contracts *ContractsCallerSession) Mpc() (common.Address, error) {
	return _Contracts.Contract.Mpc(&_Contracts.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Contracts *ContractsCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Contracts *ContractsSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Quote(&_Contracts.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Contracts *ContractsCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Quote(&_Contracts.CallOpts, amountA, reserveA, reserveB)
}

// WNATIVE is a free data retrieval call binding the contract method 0x8fd903f5.
//
// Solidity: function wNATIVE() view returns(address)
func (_Contracts *ContractsCaller) WNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "wNATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WNATIVE is a free data retrieval call binding the contract method 0x8fd903f5.
//
// Solidity: function wNATIVE() view returns(address)
func (_Contracts *ContractsSession) WNATIVE() (common.Address, error) {
	return _Contracts.Contract.WNATIVE(&_Contracts.CallOpts)
}

// WNATIVE is a free data retrieval call binding the contract method 0x8fd903f5.
//
// Solidity: function wNATIVE() view returns(address)
func (_Contracts *ContractsCallerSession) WNATIVE() (common.Address, error) {
	return _Contracts.Contract.WNATIVE(&_Contracts.CallOpts)
}

// AnySwapFeeTo is a paid mutator transaction binding the contract method 0x87cc6e2f.
//
// Solidity: function anySwapFeeTo(address token, uint256 amount) returns()
func (_Contracts *ContractsTransactor) AnySwapFeeTo(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapFeeTo", token, amount)
}

// AnySwapFeeTo is a paid mutator transaction binding the contract method 0x87cc6e2f.
//
// Solidity: function anySwapFeeTo(address token, uint256 amount) returns()
func (_Contracts *ContractsSession) AnySwapFeeTo(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapFeeTo(&_Contracts.TransactOpts, token, amount)
}

// AnySwapFeeTo is a paid mutator transaction binding the contract method 0x87cc6e2f.
//
// Solidity: function anySwapFeeTo(address token, uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) AnySwapFeeTo(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapFeeTo(&_Contracts.TransactOpts, token, amount)
}

// AnySwapIn is a paid mutator transaction binding the contract method 0x25121b76.
//
// Solidity: function anySwapIn(bytes32[] txs, address[] tokens, address[] to, uint256[] amounts, uint256[] fromChainIDs) returns()
func (_Contracts *ContractsTransactor) AnySwapIn(opts *bind.TransactOpts, txs [][32]byte, tokens []common.Address, to []common.Address, amounts []*big.Int, fromChainIDs []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapIn", txs, tokens, to, amounts, fromChainIDs)
}

// AnySwapIn is a paid mutator transaction binding the contract method 0x25121b76.
//
// Solidity: function anySwapIn(bytes32[] txs, address[] tokens, address[] to, uint256[] amounts, uint256[] fromChainIDs) returns()
func (_Contracts *ContractsSession) AnySwapIn(txs [][32]byte, tokens []common.Address, to []common.Address, amounts []*big.Int, fromChainIDs []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapIn(&_Contracts.TransactOpts, txs, tokens, to, amounts, fromChainIDs)
}

// AnySwapIn is a paid mutator transaction binding the contract method 0x25121b76.
//
// Solidity: function anySwapIn(bytes32[] txs, address[] tokens, address[] to, uint256[] amounts, uint256[] fromChainIDs) returns()
func (_Contracts *ContractsTransactorSession) AnySwapIn(txs [][32]byte, tokens []common.Address, to []common.Address, amounts []*big.Int, fromChainIDs []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapIn(&_Contracts.TransactOpts, txs, tokens, to, amounts, fromChainIDs)
}

// AnySwapIn0 is a paid mutator transaction binding the contract method 0x825bb13c.
//
// Solidity: function anySwapIn(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapIn0(opts *bind.TransactOpts, txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapIn0", txs, token, to, amount, fromChainID)
}

// AnySwapIn0 is a paid mutator transaction binding the contract method 0x825bb13c.
//
// Solidity: function anySwapIn(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_Contracts *ContractsSession) AnySwapIn0(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapIn0(&_Contracts.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapIn0 is a paid mutator transaction binding the contract method 0x825bb13c.
//
// Solidity: function anySwapIn(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapIn0(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapIn0(&_Contracts.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapInAuto is a paid mutator transaction binding the contract method 0x0175b1c4.
//
// Solidity: function anySwapInAuto(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapInAuto(opts *bind.TransactOpts, txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapInAuto", txs, token, to, amount, fromChainID)
}

// AnySwapInAuto is a paid mutator transaction binding the contract method 0x0175b1c4.
//
// Solidity: function anySwapInAuto(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_Contracts *ContractsSession) AnySwapInAuto(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInAuto(&_Contracts.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapInAuto is a paid mutator transaction binding the contract method 0x0175b1c4.
//
// Solidity: function anySwapInAuto(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapInAuto(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInAuto(&_Contracts.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapInExactTokensForNative is a paid mutator transaction binding the contract method 0x52a397d5.
//
// Solidity: function anySwapInExactTokensForNative(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_Contracts *ContractsTransactor) AnySwapInExactTokensForNative(opts *bind.TransactOpts, txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapInExactTokensForNative", txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForNative is a paid mutator transaction binding the contract method 0x52a397d5.
//
// Solidity: function anySwapInExactTokensForNative(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_Contracts *ContractsSession) AnySwapInExactTokensForNative(txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInExactTokensForNative(&_Contracts.TransactOpts, txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForNative is a paid mutator transaction binding the contract method 0x52a397d5.
//
// Solidity: function anySwapInExactTokensForNative(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_Contracts *ContractsTransactorSession) AnySwapInExactTokensForNative(txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInExactTokensForNative(&_Contracts.TransactOpts, txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForTokens is a paid mutator transaction binding the contract method 0x2fc1e728.
//
// Solidity: function anySwapInExactTokensForTokens(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_Contracts *ContractsTransactor) AnySwapInExactTokensForTokens(opts *bind.TransactOpts, txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapInExactTokensForTokens", txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForTokens is a paid mutator transaction binding the contract method 0x2fc1e728.
//
// Solidity: function anySwapInExactTokensForTokens(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_Contracts *ContractsSession) AnySwapInExactTokensForTokens(txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInExactTokensForTokens(&_Contracts.TransactOpts, txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInExactTokensForTokens is a paid mutator transaction binding the contract method 0x2fc1e728.
//
// Solidity: function anySwapInExactTokensForTokens(bytes32 txs, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 fromChainID) returns(uint256[] amounts)
func (_Contracts *ContractsTransactorSession) AnySwapInExactTokensForTokens(txs [32]byte, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInExactTokensForTokens(&_Contracts.TransactOpts, txs, amountIn, amountOutMin, path, to, deadline, fromChainID)
}

// AnySwapInUnderlying is a paid mutator transaction binding the contract method 0x3f88de89.
//
// Solidity: function anySwapInUnderlying(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapInUnderlying(opts *bind.TransactOpts, txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapInUnderlying", txs, token, to, amount, fromChainID)
}

// AnySwapInUnderlying is a paid mutator transaction binding the contract method 0x3f88de89.
//
// Solidity: function anySwapInUnderlying(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_Contracts *ContractsSession) AnySwapInUnderlying(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInUnderlying(&_Contracts.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapInUnderlying is a paid mutator transaction binding the contract method 0x3f88de89.
//
// Solidity: function anySwapInUnderlying(bytes32 txs, address token, address to, uint256 amount, uint256 fromChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapInUnderlying(txs [32]byte, token common.Address, to common.Address, amount *big.Int, fromChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInUnderlying(&_Contracts.TransactOpts, txs, token, to, amount, fromChainID)
}

// AnySwapOut is a paid mutator transaction binding the contract method 0x241dc2df.
//
// Solidity: function anySwapOut(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOut(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOut", token, to, amount, toChainID)
}

// AnySwapOut is a paid mutator transaction binding the contract method 0x241dc2df.
//
// Solidity: function anySwapOut(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOut(token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOut(&_Contracts.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOut is a paid mutator transaction binding the contract method 0x241dc2df.
//
// Solidity: function anySwapOut(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOut(token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOut(&_Contracts.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOut0 is a paid mutator transaction binding the contract method 0xdcfb77b1.
//
// Solidity: function anySwapOut(address[] tokens, address[] to, uint256[] amounts, uint256[] toChainIDs) returns()
func (_Contracts *ContractsTransactor) AnySwapOut0(opts *bind.TransactOpts, tokens []common.Address, to []common.Address, amounts []*big.Int, toChainIDs []*big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOut0", tokens, to, amounts, toChainIDs)
}

// AnySwapOut0 is a paid mutator transaction binding the contract method 0xdcfb77b1.
//
// Solidity: function anySwapOut(address[] tokens, address[] to, uint256[] amounts, uint256[] toChainIDs) returns()
func (_Contracts *ContractsSession) AnySwapOut0(tokens []common.Address, to []common.Address, amounts []*big.Int, toChainIDs []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOut0(&_Contracts.TransactOpts, tokens, to, amounts, toChainIDs)
}

// AnySwapOut0 is a paid mutator transaction binding the contract method 0xdcfb77b1.
//
// Solidity: function anySwapOut(address[] tokens, address[] to, uint256[] amounts, uint256[] toChainIDs) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOut0(tokens []common.Address, to []common.Address, amounts []*big.Int, toChainIDs []*big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOut0(&_Contracts.TransactOpts, tokens, to, amounts, toChainIDs)
}

// AnySwapOutExactTokensForNative is a paid mutator transaction binding the contract method 0x65782f56.
//
// Solidity: function anySwapOutExactTokensForNative(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutExactTokensForNative(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutExactTokensForNative", amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNative is a paid mutator transaction binding the contract method 0x65782f56.
//
// Solidity: function anySwapOutExactTokensForNative(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutExactTokensForNative(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForNative(&_Contracts.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNative is a paid mutator transaction binding the contract method 0x65782f56.
//
// Solidity: function anySwapOutExactTokensForNative(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutExactTokensForNative(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForNative(&_Contracts.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlying is a paid mutator transaction binding the contract method 0x6a453972.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutExactTokensForNativeUnderlying(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutExactTokensForNativeUnderlying", amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlying is a paid mutator transaction binding the contract method 0x6a453972.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutExactTokensForNativeUnderlying(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForNativeUnderlying(&_Contracts.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlying is a paid mutator transaction binding the contract method 0x6a453972.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutExactTokensForNativeUnderlying(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForNativeUnderlying(&_Contracts.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x4d93bb94.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutExactTokensForNativeUnderlyingWithPermit(opts *bind.TransactOpts, from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutExactTokensForNativeUnderlyingWithPermit", from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x4d93bb94.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutExactTokensForNativeUnderlyingWithPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForNativeUnderlyingWithPermit(&_Contracts.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x4d93bb94.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutExactTokensForNativeUnderlyingWithPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForNativeUnderlyingWithPermit(&_Contracts.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0xc8e174f6.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(opts *bind.TransactOpts, from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutExactTokensForNativeUnderlyingWithTransferPermit", from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0xc8e174f6.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(&_Contracts.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0xc8e174f6.
//
// Solidity: function anySwapOutExactTokensForNativeUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForNativeUnderlyingWithTransferPermit(&_Contracts.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokens is a paid mutator transaction binding the contract method 0x0bb57203.
//
// Solidity: function anySwapOutExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutExactTokensForTokens", amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokens is a paid mutator transaction binding the contract method 0x0bb57203.
//
// Solidity: function anySwapOutExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForTokens(&_Contracts.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokens is a paid mutator transaction binding the contract method 0x0bb57203.
//
// Solidity: function anySwapOutExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForTokens(&_Contracts.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlying is a paid mutator transaction binding the contract method 0xd8b9f610.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutExactTokensForTokensUnderlying(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutExactTokensForTokensUnderlying", amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlying is a paid mutator transaction binding the contract method 0xd8b9f610.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutExactTokensForTokensUnderlying(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForTokensUnderlying(&_Contracts.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlying is a paid mutator transaction binding the contract method 0xd8b9f610.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlying(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutExactTokensForTokensUnderlying(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForTokensUnderlying(&_Contracts.TransactOpts, amountIn, amountOutMin, path, to, deadline, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x99cd84b5.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutExactTokensForTokensUnderlyingWithPermit(opts *bind.TransactOpts, from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutExactTokensForTokensUnderlyingWithPermit", from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x99cd84b5.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutExactTokensForTokensUnderlyingWithPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForTokensUnderlyingWithPermit(&_Contracts.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x99cd84b5.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutExactTokensForTokensUnderlyingWithPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForTokensUnderlyingWithPermit(&_Contracts.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x9aa1ac61.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(opts *bind.TransactOpts, from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutExactTokensForTokensUnderlyingWithTransferPermit", from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x9aa1ac61.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(&_Contracts.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x9aa1ac61.
//
// Solidity: function anySwapOutExactTokensForTokensUnderlyingWithTransferPermit(address from, uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(from common.Address, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutExactTokensForTokensUnderlyingWithTransferPermit(&_Contracts.TransactOpts, from, amountIn, amountOutMin, path, to, deadline, v, r, s, toChainID)
}

// AnySwapOutNative is a paid mutator transaction binding the contract method 0xa5e56571.
//
// Solidity: function anySwapOutNative(address token, address to, uint256 toChainID) payable returns()
func (_Contracts *ContractsTransactor) AnySwapOutNative(opts *bind.TransactOpts, token common.Address, to common.Address, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutNative", token, to, toChainID)
}

// AnySwapOutNative is a paid mutator transaction binding the contract method 0xa5e56571.
//
// Solidity: function anySwapOutNative(address token, address to, uint256 toChainID) payable returns()
func (_Contracts *ContractsSession) AnySwapOutNative(token common.Address, to common.Address, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutNative(&_Contracts.TransactOpts, token, to, toChainID)
}

// AnySwapOutNative is a paid mutator transaction binding the contract method 0xa5e56571.
//
// Solidity: function anySwapOutNative(address token, address to, uint256 toChainID) payable returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutNative(token common.Address, to common.Address, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutNative(&_Contracts.TransactOpts, token, to, toChainID)
}

// AnySwapOutUnderlying is a paid mutator transaction binding the contract method 0xedbdf5e2.
//
// Solidity: function anySwapOutUnderlying(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutUnderlying(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutUnderlying", token, to, amount, toChainID)
}

// AnySwapOutUnderlying is a paid mutator transaction binding the contract method 0xedbdf5e2.
//
// Solidity: function anySwapOutUnderlying(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutUnderlying(token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlying(&_Contracts.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOutUnderlying is a paid mutator transaction binding the contract method 0xedbdf5e2.
//
// Solidity: function anySwapOutUnderlying(address token, address to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutUnderlying(token common.Address, to common.Address, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlying(&_Contracts.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOutUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x8d7d3eea.
//
// Solidity: function anySwapOutUnderlyingWithPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutUnderlyingWithPermit(opts *bind.TransactOpts, from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutUnderlyingWithPermit", from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x8d7d3eea.
//
// Solidity: function anySwapOutUnderlyingWithPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutUnderlyingWithPermit(from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlyingWithPermit(&_Contracts.TransactOpts, from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithPermit is a paid mutator transaction binding the contract method 0x8d7d3eea.
//
// Solidity: function anySwapOutUnderlyingWithPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutUnderlyingWithPermit(from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlyingWithPermit(&_Contracts.TransactOpts, from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x1b91a934.
//
// Solidity: function anySwapOutUnderlyingWithTransferPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutUnderlyingWithTransferPermit(opts *bind.TransactOpts, from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutUnderlyingWithTransferPermit", from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x1b91a934.
//
// Solidity: function anySwapOutUnderlyingWithTransferPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutUnderlyingWithTransferPermit(from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlyingWithTransferPermit(&_Contracts.TransactOpts, from, token, to, amount, deadline, v, r, s, toChainID)
}

// AnySwapOutUnderlyingWithTransferPermit is a paid mutator transaction binding the contract method 0x1b91a934.
//
// Solidity: function anySwapOutUnderlyingWithTransferPermit(address from, address token, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutUnderlyingWithTransferPermit(from common.Address, token common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlyingWithTransferPermit(&_Contracts.TransactOpts, from, token, to, amount, deadline, v, r, s, toChainID)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address newMPC) returns(bool)
func (_Contracts *ContractsTransactor) ChangeMPC(opts *bind.TransactOpts, newMPC common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changeMPC", newMPC)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address newMPC) returns(bool)
func (_Contracts *ContractsSession) ChangeMPC(newMPC common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeMPC(&_Contracts.TransactOpts, newMPC)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address newMPC) returns(bool)
func (_Contracts *ContractsTransactorSession) ChangeMPC(newMPC common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeMPC(&_Contracts.TransactOpts, newMPC)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x456862aa.
//
// Solidity: function changeVault(address token, address newVault) returns(bool)
func (_Contracts *ContractsTransactor) ChangeVault(opts *bind.TransactOpts, token common.Address, newVault common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changeVault", token, newVault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x456862aa.
//
// Solidity: function changeVault(address token, address newVault) returns(bool)
func (_Contracts *ContractsSession) ChangeVault(token common.Address, newVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeVault(&_Contracts.TransactOpts, token, newVault)
}

// ChangeVault is a paid mutator transaction binding the contract method 0x456862aa.
//
// Solidity: function changeVault(address token, address newVault) returns(bool)
func (_Contracts *ContractsTransactorSession) ChangeVault(token common.Address, newVault common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeVault(&_Contracts.TransactOpts, token, newVault)
}

// DepositNative is a paid mutator transaction binding the contract method 0x701bb891.
//
// Solidity: function depositNative(address token, address to) payable returns(uint256)
func (_Contracts *ContractsTransactor) DepositNative(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "depositNative", token, to)
}

// DepositNative is a paid mutator transaction binding the contract method 0x701bb891.
//
// Solidity: function depositNative(address token, address to) payable returns(uint256)
func (_Contracts *ContractsSession) DepositNative(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DepositNative(&_Contracts.TransactOpts, token, to)
}

// DepositNative is a paid mutator transaction binding the contract method 0x701bb891.
//
// Solidity: function depositNative(address token, address to) payable returns(uint256)
func (_Contracts *ContractsTransactorSession) DepositNative(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.DepositNative(&_Contracts.TransactOpts, token, to)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x832e9492.
//
// Solidity: function withdrawNative(address token, uint256 amount, address to) returns(uint256)
func (_Contracts *ContractsTransactor) WithdrawNative(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawNative", token, amount, to)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x832e9492.
//
// Solidity: function withdrawNative(address token, uint256 amount, address to) returns(uint256)
func (_Contracts *ContractsSession) WithdrawNative(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawNative(&_Contracts.TransactOpts, token, amount, to)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x832e9492.
//
// Solidity: function withdrawNative(address token, uint256 amount, address to) returns(uint256)
func (_Contracts *ContractsTransactorSession) WithdrawNative(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawNative(&_Contracts.TransactOpts, token, amount, to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contracts *ContractsTransactorSession) Receive() (*types.Transaction, error) {
	return _Contracts.Contract.Receive(&_Contracts.TransactOpts)
}

// ContractsLogAnySwapInIterator is returned from FilterLogAnySwapIn and is used to iterate over the raw logs and unpacked data for LogAnySwapIn events raised by the Contracts contract.
type ContractsLogAnySwapInIterator struct {
	Event *ContractsLogAnySwapIn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogAnySwapInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogAnySwapIn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogAnySwapIn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogAnySwapInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogAnySwapInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogAnySwapIn represents a LogAnySwapIn event raised by the Contracts contract.
type ContractsLogAnySwapIn struct {
	Txhash      [32]byte
	Token       common.Address
	To          common.Address
	Amount      *big.Int
	FromChainID *big.Int
	ToChainID   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapIn is a free log retrieval operation binding the contract event 0xaac9ce45fe3adf5143598c4f18a369591a20a3384aedaf1b525d29127e1fcd55.
//
// Solidity: event LogAnySwapIn(bytes32 indexed txhash, address indexed token, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) FilterLogAnySwapIn(opts *bind.FilterOpts, txhash [][32]byte, token []common.Address, to []common.Address) (*ContractsLogAnySwapInIterator, error) {

	var txhashRule []interface{}
	for _, txhashItem := range txhash {
		txhashRule = append(txhashRule, txhashItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogAnySwapIn", txhashRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogAnySwapInIterator{contract: _Contracts.contract, event: "LogAnySwapIn", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapIn is a free log subscription operation binding the contract event 0xaac9ce45fe3adf5143598c4f18a369591a20a3384aedaf1b525d29127e1fcd55.
//
// Solidity: event LogAnySwapIn(bytes32 indexed txhash, address indexed token, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) WatchLogAnySwapIn(opts *bind.WatchOpts, sink chan<- *ContractsLogAnySwapIn, txhash [][32]byte, token []common.Address, to []common.Address) (event.Subscription, error) {

	var txhashRule []interface{}
	for _, txhashItem := range txhash {
		txhashRule = append(txhashRule, txhashItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogAnySwapIn", txhashRule, tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogAnySwapIn)
				if err := _Contracts.contract.UnpackLog(event, "LogAnySwapIn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogAnySwapIn is a log parse operation binding the contract event 0xaac9ce45fe3adf5143598c4f18a369591a20a3384aedaf1b525d29127e1fcd55.
//
// Solidity: event LogAnySwapIn(bytes32 indexed txhash, address indexed token, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) ParseLogAnySwapIn(log types.Log) (*ContractsLogAnySwapIn, error) {
	event := new(ContractsLogAnySwapIn)
	if err := _Contracts.contract.UnpackLog(event, "LogAnySwapIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogAnySwapOutIterator is returned from FilterLogAnySwapOut and is used to iterate over the raw logs and unpacked data for LogAnySwapOut events raised by the Contracts contract.
type ContractsLogAnySwapOutIterator struct {
	Event *ContractsLogAnySwapOut // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogAnySwapOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogAnySwapOut)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogAnySwapOut)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogAnySwapOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogAnySwapOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogAnySwapOut represents a LogAnySwapOut event raised by the Contracts contract.
type ContractsLogAnySwapOut struct {
	Token       common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	FromChainID *big.Int
	ToChainID   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapOut is a free log retrieval operation binding the contract event 0x97116cf6cd4f6412bb47914d6db18da9e16ab2142f543b86e207c24fbd16b23a.
//
// Solidity: event LogAnySwapOut(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) FilterLogAnySwapOut(opts *bind.FilterOpts, token []common.Address, from []common.Address, to []common.Address) (*ContractsLogAnySwapOutIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogAnySwapOut", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogAnySwapOutIterator{contract: _Contracts.contract, event: "LogAnySwapOut", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapOut is a free log subscription operation binding the contract event 0x97116cf6cd4f6412bb47914d6db18da9e16ab2142f543b86e207c24fbd16b23a.
//
// Solidity: event LogAnySwapOut(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) WatchLogAnySwapOut(opts *bind.WatchOpts, sink chan<- *ContractsLogAnySwapOut, token []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogAnySwapOut", tokenRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogAnySwapOut)
				if err := _Contracts.contract.UnpackLog(event, "LogAnySwapOut", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogAnySwapOut is a log parse operation binding the contract event 0x97116cf6cd4f6412bb47914d6db18da9e16ab2142f543b86e207c24fbd16b23a.
//
// Solidity: event LogAnySwapOut(address indexed token, address indexed from, address indexed to, uint256 amount, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) ParseLogAnySwapOut(log types.Log) (*ContractsLogAnySwapOut, error) {
	event := new(ContractsLogAnySwapOut)
	if err := _Contracts.contract.UnpackLog(event, "LogAnySwapOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogAnySwapTradeTokensForNativeIterator is returned from FilterLogAnySwapTradeTokensForNative and is used to iterate over the raw logs and unpacked data for LogAnySwapTradeTokensForNative events raised by the Contracts contract.
type ContractsLogAnySwapTradeTokensForNativeIterator struct {
	Event *ContractsLogAnySwapTradeTokensForNative // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogAnySwapTradeTokensForNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogAnySwapTradeTokensForNative)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogAnySwapTradeTokensForNative)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogAnySwapTradeTokensForNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogAnySwapTradeTokensForNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogAnySwapTradeTokensForNative represents a LogAnySwapTradeTokensForNative event raised by the Contracts contract.
type ContractsLogAnySwapTradeTokensForNative struct {
	Path         []common.Address
	From         common.Address
	To           common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	FromChainID  *big.Int
	ToChainID    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapTradeTokensForNative is a free log retrieval operation binding the contract event 0x278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f4.
//
// Solidity: event LogAnySwapTradeTokensForNative(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) FilterLogAnySwapTradeTokensForNative(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsLogAnySwapTradeTokensForNativeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogAnySwapTradeTokensForNative", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogAnySwapTradeTokensForNativeIterator{contract: _Contracts.contract, event: "LogAnySwapTradeTokensForNative", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapTradeTokensForNative is a free log subscription operation binding the contract event 0x278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f4.
//
// Solidity: event LogAnySwapTradeTokensForNative(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) WatchLogAnySwapTradeTokensForNative(opts *bind.WatchOpts, sink chan<- *ContractsLogAnySwapTradeTokensForNative, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogAnySwapTradeTokensForNative", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogAnySwapTradeTokensForNative)
				if err := _Contracts.contract.UnpackLog(event, "LogAnySwapTradeTokensForNative", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogAnySwapTradeTokensForNative is a log parse operation binding the contract event 0x278277e0209c347189add7bd92411973b5f6b8644f7ac62ea1be984ce993f8f4.
//
// Solidity: event LogAnySwapTradeTokensForNative(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) ParseLogAnySwapTradeTokensForNative(log types.Log) (*ContractsLogAnySwapTradeTokensForNative, error) {
	event := new(ContractsLogAnySwapTradeTokensForNative)
	if err := _Contracts.contract.UnpackLog(event, "LogAnySwapTradeTokensForNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogAnySwapTradeTokensForTokensIterator is returned from FilterLogAnySwapTradeTokensForTokens and is used to iterate over the raw logs and unpacked data for LogAnySwapTradeTokensForTokens events raised by the Contracts contract.
type ContractsLogAnySwapTradeTokensForTokensIterator struct {
	Event *ContractsLogAnySwapTradeTokensForTokens // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogAnySwapTradeTokensForTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogAnySwapTradeTokensForTokens)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogAnySwapTradeTokensForTokens)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogAnySwapTradeTokensForTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogAnySwapTradeTokensForTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogAnySwapTradeTokensForTokens represents a LogAnySwapTradeTokensForTokens event raised by the Contracts contract.
type ContractsLogAnySwapTradeTokensForTokens struct {
	Path         []common.Address
	From         common.Address
	To           common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	FromChainID  *big.Int
	ToChainID    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapTradeTokensForTokens is a free log retrieval operation binding the contract event 0xfea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b3.
//
// Solidity: event LogAnySwapTradeTokensForTokens(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) FilterLogAnySwapTradeTokensForTokens(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsLogAnySwapTradeTokensForTokensIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogAnySwapTradeTokensForTokens", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogAnySwapTradeTokensForTokensIterator{contract: _Contracts.contract, event: "LogAnySwapTradeTokensForTokens", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapTradeTokensForTokens is a free log subscription operation binding the contract event 0xfea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b3.
//
// Solidity: event LogAnySwapTradeTokensForTokens(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) WatchLogAnySwapTradeTokensForTokens(opts *bind.WatchOpts, sink chan<- *ContractsLogAnySwapTradeTokensForTokens, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogAnySwapTradeTokensForTokens", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogAnySwapTradeTokensForTokens)
				if err := _Contracts.contract.UnpackLog(event, "LogAnySwapTradeTokensForTokens", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogAnySwapTradeTokensForTokens is a log parse operation binding the contract event 0xfea6abdf4fd32f20966dff7619354cd82cd43dc78a3bee479f04c74dbfc585b3.
//
// Solidity: event LogAnySwapTradeTokensForTokens(address[] path, address indexed from, address indexed to, uint256 amountIn, uint256 amountOutMin, uint256 fromChainID, uint256 toChainID)
func (_Contracts *ContractsFilterer) ParseLogAnySwapTradeTokensForTokens(log types.Log) (*ContractsLogAnySwapTradeTokensForTokens, error) {
	event := new(ContractsLogAnySwapTradeTokensForTokens)
	if err := _Contracts.contract.UnpackLog(event, "LogAnySwapTradeTokensForTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogChangeMPCIterator is returned from FilterLogChangeMPC and is used to iterate over the raw logs and unpacked data for LogChangeMPC events raised by the Contracts contract.
type ContractsLogChangeMPCIterator struct {
	Event *ContractsLogChangeMPC // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogChangeMPCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogChangeMPC)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogChangeMPC)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogChangeMPCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogChangeMPCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogChangeMPC represents a LogChangeMPC event raised by the Contracts contract.
type ContractsLogChangeMPC struct {
	OldMPC        common.Address
	NewMPC        common.Address
	EffectiveTime *big.Int
	ChainID       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogChangeMPC is a free log retrieval operation binding the contract event 0xcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a2198121.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 indexed effectiveTime, uint256 chainID)
func (_Contracts *ContractsFilterer) FilterLogChangeMPC(opts *bind.FilterOpts, oldMPC []common.Address, newMPC []common.Address, effectiveTime []*big.Int) (*ContractsLogChangeMPCIterator, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogChangeMPC", oldMPCRule, newMPCRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogChangeMPCIterator{contract: _Contracts.contract, event: "LogChangeMPC", logs: logs, sub: sub}, nil
}

// WatchLogChangeMPC is a free log subscription operation binding the contract event 0xcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a2198121.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 indexed effectiveTime, uint256 chainID)
func (_Contracts *ContractsFilterer) WatchLogChangeMPC(opts *bind.WatchOpts, sink chan<- *ContractsLogChangeMPC, oldMPC []common.Address, newMPC []common.Address, effectiveTime []*big.Int) (event.Subscription, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogChangeMPC", oldMPCRule, newMPCRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogChangeMPC)
				if err := _Contracts.contract.UnpackLog(event, "LogChangeMPC", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogChangeMPC is a log parse operation binding the contract event 0xcda32bc39904597666dfa9f9c845714756e1ffffad55b52e0d344673a2198121.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 indexed effectiveTime, uint256 chainID)
func (_Contracts *ContractsFilterer) ParseLogChangeMPC(log types.Log) (*ContractsLogChangeMPC, error) {
	event := new(ContractsLogChangeMPC)
	if err := _Contracts.contract.UnpackLog(event, "LogChangeMPC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogChangeRouterIterator is returned from FilterLogChangeRouter and is used to iterate over the raw logs and unpacked data for LogChangeRouter events raised by the Contracts contract.
type ContractsLogChangeRouterIterator struct {
	Event *ContractsLogChangeRouter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsLogChangeRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogChangeRouter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsLogChangeRouter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsLogChangeRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogChangeRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogChangeRouter represents a LogChangeRouter event raised by the Contracts contract.
type ContractsLogChangeRouter struct {
	OldRouter common.Address
	NewRouter common.Address
	ChainID   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogChangeRouter is a free log retrieval operation binding the contract event 0x7eefe162042d50d604dca716bef4ff4c5e318a056f712c0195d016f78089955a.
//
// Solidity: event LogChangeRouter(address indexed oldRouter, address indexed newRouter, uint256 chainID)
func (_Contracts *ContractsFilterer) FilterLogChangeRouter(opts *bind.FilterOpts, oldRouter []common.Address, newRouter []common.Address) (*ContractsLogChangeRouterIterator, error) {

	var oldRouterRule []interface{}
	for _, oldRouterItem := range oldRouter {
		oldRouterRule = append(oldRouterRule, oldRouterItem)
	}
	var newRouterRule []interface{}
	for _, newRouterItem := range newRouter {
		newRouterRule = append(newRouterRule, newRouterItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogChangeRouter", oldRouterRule, newRouterRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogChangeRouterIterator{contract: _Contracts.contract, event: "LogChangeRouter", logs: logs, sub: sub}, nil
}

// WatchLogChangeRouter is a free log subscription operation binding the contract event 0x7eefe162042d50d604dca716bef4ff4c5e318a056f712c0195d016f78089955a.
//
// Solidity: event LogChangeRouter(address indexed oldRouter, address indexed newRouter, uint256 chainID)
func (_Contracts *ContractsFilterer) WatchLogChangeRouter(opts *bind.WatchOpts, sink chan<- *ContractsLogChangeRouter, oldRouter []common.Address, newRouter []common.Address) (event.Subscription, error) {

	var oldRouterRule []interface{}
	for _, oldRouterItem := range oldRouter {
		oldRouterRule = append(oldRouterRule, oldRouterItem)
	}
	var newRouterRule []interface{}
	for _, newRouterItem := range newRouter {
		newRouterRule = append(newRouterRule, newRouterItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogChangeRouter", oldRouterRule, newRouterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogChangeRouter)
				if err := _Contracts.contract.UnpackLog(event, "LogChangeRouter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogChangeRouter is a log parse operation binding the contract event 0x7eefe162042d50d604dca716bef4ff4c5e318a056f712c0195d016f78089955a.
//
// Solidity: event LogChangeRouter(address indexed oldRouter, address indexed newRouter, uint256 chainID)
func (_Contracts *ContractsFilterer) ParseLogChangeRouter(log types.Log) (*ContractsLogChangeRouter, error) {
	event := new(ContractsLogChangeRouter)
	if err := _Contracts.contract.UnpackLog(event, "LogChangeRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

func (_Contracts *ContractsTransactor) Swapout(opts *bind.TransactOpts, amount *big.Int, bindaddr common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "Swapout", amount, bindaddr)
}


func (_Contracts *ContractsTransactor) Transfer(opts *bind.TransactOpts, toAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer",  toAddress,amount)
}
