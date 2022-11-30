// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MultichainV7Router

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

// SwapInfo is an auto generated low-level Go binding around an user-defined struct.
type SwapInfo struct {
	SwapoutID   [32]byte
	Token       common.Address
	Receiver    common.Address
	Amount      *big.Int
	FromChainID *big.Int
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mpc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wNATIVE\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_anycallExecutor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_routerSecurity\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_old\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"ChangeAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"LogAnySwapInAndExec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"LogAnySwapOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"anycallProxy\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogAnySwapOutAndCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMPC\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMPC\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"applyTime\",\"type\":\"uint256\"}],\"name\":\"LogApplyMPC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMPC\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMPC\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"effectiveTime\",\"type\":\"uint256\"}],\"name\":\"LogChangeMPC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"anycallProxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogRetryExecRecord\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"dontExec\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"LogRetrySwapInAndExec\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Call_Paused_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Exec_Paused_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ALL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Retry_Paused_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Swapin_Paused_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Swapout_Paused_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"proxies\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"acceptAnyTokenFlags\",\"type\":\"bool[]\"}],\"name\":\"addAnycallProxies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"anySwapFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"internalType\":\"structSwapInfo\",\"name\":\"swapInfo\",\"type\":\"tuple\"}],\"name\":\"anySwapIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"internalType\":\"structSwapInfo\",\"name\":\"swapInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"anycallProxy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"anySwapInAndExec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"internalType\":\"structSwapInfo\",\"name\":\"swapInfo\",\"type\":\"tuple\"}],\"name\":\"anySwapInAuto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"internalType\":\"structSwapInfo\",\"name\":\"swapInfo\",\"type\":\"tuple\"}],\"name\":\"anySwapInNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"internalType\":\"structSwapInfo\",\"name\":\"swapInfo\",\"type\":\"tuple\"}],\"name\":\"anySwapInUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"internalType\":\"structSwapInfo\",\"name\":\"swapInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"anycallProxy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"anySwapInUnderlyingAndExec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"anycallProxy\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"anySwapOutAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"anycallProxy\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"anySwapOutNativeAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"}],\"name\":\"anySwapOutUnderlying\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"to\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"anycallProxy\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"anySwapOutUnderlyingAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"anycallExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"anycallProxyInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"acceptAnyToken\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"applyMPC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mpc\",\"type\":\"address\"}],\"name\":\"changeMPC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newVault\",\"type\":\"address\"}],\"name\":\"changeVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayMPC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mpc\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingMPC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"proxies\",\"type\":\"address[]\"}],\"name\":\"removeAnycallProxies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"retryRecords\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"swapID\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"swapoutID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainID\",\"type\":\"uint256\"}],\"internalType\":\"structSwapInfo\",\"name\":\"swapInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"anycallProxy\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"dontExec\",\"type\":\"bool\"}],\"name\":\"retrySwapinAndExec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"routerSecurity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_routerSecurity\",\"type\":\"address\"}],\"name\":\"setRouterSecurity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wNATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620058fe380380620058fe8339810160408190526200003491620001e5565b84848181806001600160a01b038116620000955760405162461bcd60e51b815260206004820152601c60248201527f4d50433a206d706320697320746865207a65726f20616464726573730000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b03831690811782556040514281529091907f581f388e3dd32e1bbf62a290f509c8245f9d0b71ef82614fb2b967ad0a10d5b99060200160405180910390a350600380546001600160a01b0319166001600160a01b0384169081179091556040516000907fcf9b665e0639e0b81a8db37b60ac7ddf45aeb1b484e11adeb7dff4bf4a3a6258908290a35050600160055550506001600160a01b038216620001935760405162461bcd60e51b815260206004820152601560248201527f7a65726f20616e7963616c6c206578656375746f72000000000000000000000060448201526064016200008c565b6001600160a01b0391821660a052918116608052600680546001600160a01b0319169290911691909117905550620002559050565b80516001600160a01b0381168114620001e057600080fd5b919050565b600080600080600060a08688031215620001fe57600080fd5b6200020986620001c8565b94506200021960208701620001c8565b93506200022960408701620001c8565b92506200023960608701620001c8565b91506200024960808701620001c8565b90509295509295909350565b60805160a05161561a620002e4600039600081816106b301528181611fc7015281816131880152613b6901526000818161024f015281816104f901528181610d0f01528181610d8d01528181610fd90152818161166f01528181611744015281816140260152818161409c0152818161416e015281816141df0152818161425d01526142aa015261561a6000f3fe60806040526004361061023f5760003560e01c80639ac25d081161012e578063e0e9048e116100ab578063f75c26641161006f578063f75c26641461077c578063f830e7b41461079c578063f851a440146107bc578063f91275b5146107dc578063f9ca3a5d146107fe57600080fd5b8063e0e9048e146106d5578063e2ea2ba9146106f5578063e94b714414610715578063ea0c968b14610749578063ed56531a1461075c57600080fd5b8063b63b38d0116100f2578063b63b38d01461062c578063c604b0b814610641578063cc95060a14610661578063d21c1cf514610681578063d2c7dfcc146106a157600080fd5b80639ac25d08146105875780639e9e46661461059c5780639ff1d3e8146105cc578063a413387a146105ec578063a66ec4431461060c57600080fd5b80636a42b8f8116101bc57806387cc6e2f1161018057806387cc6e2f146104a75780638f283970146104c75780638fd903f5146104e75780638fef848914610533578063912d857c1461055357600080fd5b80636a42b8f8146104035780636a6459d11461041a5780636b4b43761461044757806381aa7a8114610467578063872acd041461048757600080fd5b8063456862aa11610203578063456862aa1461035e578063540dd52c1461038e5780635598f119146103a15780635b7b018c146103c35780635de26385146103e357600080fd5b8063049b4e7e146102835780630c55b22e146102a3578063160f1053146102d85780631d5aa281146102ee5780632f4dae9f1461033e57600080fd5b3661027e57336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461027c5761027c6147e2565b005b600080fd5b34801561028f57600080fd5b5061027c61029e366004614856565b61081e565b3480156102af57600080fd5b506102c56000805160206155c583398151915281565b6040519081526020015b60405180910390f35b3480156102e457600080fd5b506102c560025481565b3480156102fa57600080fd5b506103276103093660046148bd565b60076020526000908152604090205460ff8082169161010090041682565b6040805192151583529015156020830152016102cf565b34801561034a57600080fd5b5061027c6103593660046148da565b61093f565b34801561036a57600080fd5b5061037e6103793660046148f3565b610975565b60405190151581526020016102cf565b61027c61039c36600461492c565b610a40565b3480156103ad57600080fd5b506102c56000805160206155a583398151915281565b3480156103cf57600080fd5b5061027c6103de3660046148bd565b610b56565b3480156103ef57600080fd5b5061027c6103fe3660046149a0565b610c4e565b34801561040f57600080fd5b506102c56202a30081565b34801561042657600080fd5b506102c56104353660046148da565b60086020526000908152604090205481565b34801561045357600080fd5b5061027c6104623660046149f5565b6110e3565b34801561047357600080fd5b5061027c6104823660046149a0565b61139a565b34801561049357600080fd5b5061027c6104a2366004614ac5565b61199a565b3480156104b357600080fd5b5061027c6104c2366004614b74565b6120f6565b3480156104d357600080fd5b5061027c6104e23660046148bd565b612232565b3480156104f357600080fd5b5061051b7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016102cf565b34801561053f57600080fd5b5061027c61054e3660046149a0565b6122b8565b34801561055f57600080fd5b506102c57f2db31f196e2df05c7a9363b1c9f780b80a1c446ac321107e34140944bfef822f81565b34801561059357600080fd5b506102c5600081565b3480156105a857600080fd5b5061037e6105b73660046148da565b60009081526004602052604090205460ff1690565b3480156105d857600080fd5b5061027c6105e73660046149a0565b61247d565b3480156105f857600080fd5b5060065461051b906001600160a01b031681565b34801561061857600080fd5b5061027c6106273660046148bd565b61277e565b34801561063857600080fd5b5061027c6127f6565b34801561064d57600080fd5b5061027c61065c366004614856565b612948565b34801561066d57600080fd5b5061027c61067c366004614ba0565b612b25565b34801561068d57600080fd5b5061027c61069c366004614c7f565b6132c4565b3480156106ad57600080fd5b5061051b7f000000000000000000000000000000000000000000000000000000000000000081565b3480156106e157600080fd5b5061027c6106f03660046149f5565b61337f565b34801561070157600080fd5b5061027c610710366004614cc1565b61353d565b34801561072157600080fd5b506102c57f42aeccb36e4cd8c38ec0b9ee052287345afef9d7d5211d495f4abc7e1950eb2681565b61027c610757366004614d2d565b6136a6565b34801561076857600080fd5b5061027c6107773660046148da565b61387d565b34801561078857600080fd5b5060005461051b906001600160a01b031681565b3480156107a857600080fd5b5060015461051b906001600160a01b031681565b3480156107c857600080fd5b5060035461051b906001600160a01b031681565b3480156107e857600080fd5b506102c560008051602061552583398151915281565b34801561080a57600080fd5b5061027c610819366004614ba0565b6138b0565b6002600554036108495760405162461bcd60e51b815260040161084090614de3565b60405180910390fd5b6002600555600061085a8684613cb3565b60065460405163078e2c7d60e51b81529192506000916001600160a01b039091169063f1c58fa09061089a908a9033908b908b9089908b90600401614e43565b6020604051808303816000875af11580156108b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108dd9190614eac565b9050336001600160a01b0316876001600160a01b0316827f0d969ae475ff6fcaf0dcfa760d4d8607244e8d95e9bf426f8d5d69f9a3e525af898987896040516109299493929190614ec5565b60405180910390a4505060016005555050505050565b6003546001600160a01b031633146109695760405162461bcd60e51b815260040161084090614eec565b61097281613edc565b50565b60006002600554036109995760405162461bcd60e51b815260040161084090614de3565b60026005556000546001600160a01b031633146109c85760405162461bcd60e51b815260040161084090614f23565b6040516360e232a960e01b81526001600160a01b0383811660048301528416906360e232a9906024016020604051808303816000875af1158015610a10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a349190614f4a565b60016005559392505050565b600260055403610a625760405162461bcd60e51b815260040161084090614de3565b60026005556000610a7285613fb3565b60065460405163078e2c7d60e51b81529192506000916001600160a01b039091169063f1c58fa090610ab290899033908a908a9089908b90600401614e43565b6020604051808303816000875af1158015610ad1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af59190614eac565b9050336001600160a01b0316866001600160a01b0316827f0d969ae475ff6fcaf0dcfa760d4d8607244e8d95e9bf426f8d5d69f9a3e525af88888789604051610b419493929190614ec5565b60405180910390a45050600160055550505050565b6000546001600160a01b03163314610b805760405162461bcd60e51b815260040161084090614f23565b6001600160a01b038116610bd65760405162461bcd60e51b815260206004820152601c60248201527f4d50433a206d706320697320746865207a65726f2061646472657373000000006044820152606401610840565b600180546001600160a01b0319166001600160a01b038316179055610bfe6202a30042614f7d565b60028190556001546000546040519283526001600160a01b03918216929116907f581f388e3dd32e1bbf62a290f509c8245f9d0b71ef82614fb2b967ad0a10d5b99060200160405180910390a350565b600080516020615525833981519152600081905260046020526000805160206155858339815191525460ff16158015610ca057506000805260046020526000805160206155458339815191525460ff16155b610cbc5760405162461bcd60e51b815260040161084090614f96565b600260055403610cde5760405162461bcd60e51b815260040161084090614de3565b60026005556000546001600160a01b03163314610d0d5760405162461bcd60e51b815260040161084090614f23565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610d835760405162461bcd60e51b815260206004820152601e60248201527f4d756c7469636861696e526f757465723a207a65726f20774e415449564500006044820152606401610840565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016610dbd60408401602085016148bd565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e1e9190614fcd565b6001600160a01b031614610e445760405162461bcd60e51b815260040161084090614fea565b60065460405163b1f05c5760e01b81526001600160a01b039091169063b1f05c5790610e7890879087908790600401615035565b600060405180830381600087803b158015610e9257600080fd5b505af1158015610ea6573d6000803e3d6000fd5b50610ebb9250505060408301602084016148bd565b6001600160a01b03166340c10f193084606001356040518363ffffffff1660e01b8152600401610eec9291906150a7565b6020604051808303816000875af1158015610f0b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f2f9190614f4a565b610f3b57610f3b6147e2565b610f4b60408301602084016148bd565b604051627b8a6760e11b8152606084013560048201523060248201526001600160a01b03919091169062f714ce906044016020604051808303816000875af1158015610f9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fbf9190614eac565b50604051632e1a7d4d60e01b8152606083013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561102557600080fd5b505af1158015611039573d6000803e3d6000fd5b5061105c925061105291505060608401604085016148bd565b836060013561434e565b61106c60608301604084016148bd565b6001600160a01b031661108560408401602085016148bd565b6001600160a01b031683600001357f164f647883b52834be7a5219336e455a23a358be27519d0442fc0ee5e1b1ce2e8787876060013588608001356040516110d09493929190614ec5565b60405180910390a4505060016005555050565b6000805160206155c5833981519152600081905260046020526000805160206155658339815191525460ff1615801561113557506000805260046020526000805160206155458339815191525460ff16155b6111515760405162461bcd60e51b815260040161084090614f96565b6000805160206155a5833981519152600081905260046020527f2bc77a4137c409d5e5d384844713fbba0af94ceb3bee1f816ae2b8a214afd84f5460ff161580156111b557506000805260046020526000805160206155458339815191525460ff16155b6111d15760405162461bcd60e51b815260040161084090614f96565b6002600554036111f35760405162461bcd60e51b815260040161084090614de3565b6002600555826112155760405162461bcd60e51b8152600401610840906150c0565b6000600660009054906101000a90046001600160a01b03166001600160a01b031663f1c58fa08d338e8e8e8e8e8e8e8e6040518b63ffffffff1660e01b815260040161126a9a999897969594939291906150e9565b6020604051808303816000875af1158015611289573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ad9190614eac565b604051632770a7eb60e21b81529091506001600160a01b038d1690639dc29fac906112de9033908d906004016150a7565b6020604051808303816000875af11580156112fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113219190614f4a565b61132d5761132d6147e2565b336001600160a01b03168c6001600160a01b0316827f968608314ec29f6fd1a9f6ef9e96247a4da1a683917569706e2d2b60ca7c0a6d8e8e8e8e8e8e8e8e60405161137f98979695949392919061515b565b60405180910390a45050600160055550505050505050505050565b600080516020615525833981519152600081905260046020526000805160206155858339815191525460ff161580156113ec57506000805260046020526000805160206155458339815191525460ff16155b6114085760405162461bcd60e51b815260040161084090614f96565b60026005540361142a5760405162461bcd60e51b815260040161084090614de3565b60026005556000546001600160a01b031633146114595760405162461bcd60e51b815260040161084090614f23565b60065460405163b1f05c5760e01b81526001600160a01b039091169063b1f05c579061148d90879087908790600401615035565b600060405180830381600087803b1580156114a757600080fd5b505af11580156114bb573d6000803e3d6000fd5b50600092506114d391505060408401602085016148bd565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b8152600401602060405180830381865afa158015611510573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115349190614fcd565b90506001600160a01b038116158015906115d8575060608301356001600160a01b0382166370a0823161156d60408701602088016148bd565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa1580156115b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d59190614eac565b10155b15611873576115ed60408401602085016148bd565b6001600160a01b03166340c10f193085606001356040518363ffffffff1660e01b815260040161161e9291906150a7565b6020604051808303816000875af115801561163d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116619190614f4a565b61166d5761166d6147e2565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316816001600160a01b0316036117cc576116b660408401602085016148bd565b604051627b8a6760e11b8152606085013560048201523060248201526001600160a01b03919091169062f714ce906044016020604051808303816000875af1158015611706573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061172a9190614eac565b50604051632e1a7d4d60e01b8152606084013560048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561179057600080fd5b505af11580156117a4573d6000803e3d6000fd5b506117c792506117bd91505060608501604086016148bd565b846060013561434e565b611912565b6117dc60408401602085016148bd565b6001600160a01b031662f714ce606085018035906117fd90604088016148bd565b6040516001600160e01b031960e085901b16815260048101929092526001600160a01b031660248201526044016020604051808303816000875af1158015611849573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061186d9190614eac565b50611912565b61188360408401602085016148bd565b6001600160a01b03166340c10f196118a160608601604087016148bd565b85606001356040518363ffffffff1660e01b81526004016118c39291906150a7565b6020604051808303816000875af11580156118e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119069190614f4a565b611912576119126147e2565b61192260608401604085016148bd565b6001600160a01b031661193b60408501602086016148bd565b6001600160a01b031684600001357f164f647883b52834be7a5219336e455a23a358be27519d0442fc0ee5e1b1ce2e8888886060013589608001356040516119869493929190614ec5565b60405180910390a450506001600555505050565b7f2db31f196e2df05c7a9363b1c9f780b80a1c446ac321107e34140944bfef822f600081905260046020527fe584ed1e59ddb832ebee6883112ccba942d9497c51eddef0bdc8765b36179e975460ff16158015611a1057506000805260046020526000805160206155458339815191525460ff16155b611a2c5760405162461bcd60e51b815260040161084090614f96565b600260055403611a4e5760405162461bcd60e51b815260040161084090614de3565b6002600555611a6360608701604088016148bd565b6001600160a01b0316336001600160a01b03161480611a8c57506003546001600160a01b031633145b611acc5760405162461bcd60e51b81526020600482015260116024820152700666f72626964207265747279207377617607c1b6044820152606401610840565b60065460405163047b1d6560e41b81526001600160a01b03909116906347b1d65090611b07908b908b908b359060808d013590600401614ec5565b602060405180830381865afa158015611b24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b489190614f4a565b611b895760405162461bcd60e51b81526020600482015260126024820152711cddd85c081b9bdd0818dbdb5c1b195d195960721b6044820152606401610840565b600088888835611b9f60408b0160208c016148bd565b611baf60608c0160408d016148bd565b8b606001358c608001358c8c8c604051602001611bd59a999897969594939291906151b2565b60405160208183030381529060405280519060200120905088888686604051602001611c049493929190615210565b60408051601f1981840301815291815281516020928301206000848152600890935291205414611c6f5760405162461bcd60e51b81526020600482015260166024820152751c995d1c9e481c9958dbdc99081b9bdd08195e1a5cdd60521b6044820152606401610840565b6000908152600860209081526040808320839055611c92919089019089016148bd565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b8152600401602060405180830381865afa158015611ccf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cf39190614fcd565b90506001600160a01b038116611d1b5760405162461bcd60e51b815260040161084090615237565b60608701356001600160a01b0382166370a08231611d3f60408b0160208c016148bd565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015611d83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611da79190614eac565b1015611df55760405162461bcd60e51b815260206004820152601e60248201527f4d756c7469636861696e526f757465723a207265747279206661696c656400006044820152606401610840565b611e0560408801602089016148bd565b6001600160a01b03166340c10f193089606001356040518363ffffffff1660e01b8152600401611e369291906150a7565b6020604051808303816000875af1158015611e55573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e799190614f4a565b611e8557611e856147e2565b600060608415611f3657611e9f60408a0160208b016148bd565b6001600160a01b031662f714ce60608b01803590611ec09060408e016148bd565b6040516001600160e01b031960e085901b16815260048101929092526001600160a01b031660248201526044016020604051808303816000875af1158015611f0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f309190614eac565b50612071565b611f4660408a0160208b016148bd565b604051627b8a6760e11b815260608b013560048201526001600160a01b038a81166024830152919091169062f714ce906044016020604051808303816000875af1158015611f98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fbc9190614eac565b506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016635b5120f78985611ffe60608e0160408f016148bd565b8d606001358c8c6040518763ffffffff1660e01b815260040161202696959493929190615278565b6000604051808303816000875af192505050801561206657506040513d6000823e601f3d908101601f1916820160405261206391908101906152fa565b60015b156120715790925090505b7f4024f72e00ae47f03ed1dd3ab595d04dabdc9d1f95f8c039bca61946d9da0eb38b8b8b356120a660408e0160208f016148bd565b8d60400160208101906120b991906148bd565b8e606001358f608001358c8a8a6040516120dc9a999897969594939291906153e9565b60405180910390a150506001600555505050505050505050565b6002600554036121185760405162461bcd60e51b815260040161084090614de3565b60026005556000546001600160a01b031633146121475760405162461bcd60e51b815260040161084090614f23565b6040516340c10f1960e01b81526001600160a01b038316906340c10f199061217590309085906004016150a7565b6020604051808303816000875af1158015612194573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121b89190614f4a565b50604051627b8a6760e11b8152600481018290523360248201526001600160a01b0383169062f714ce906044016020604051808303816000875af1158015612204573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122289190614eac565b5050600160055550565b6000546001600160a01b0316331461225c5760405162461bcd60e51b815260040161084090614f23565b6003546040516001600160a01b038084169216907fcf9b665e0639e0b81a8db37b60ac7ddf45aeb1b484e11adeb7dff4bf4a3a625890600090a3600380546001600160a01b0319166001600160a01b0392909216919091179055565b600080516020615525833981519152600081905260046020526000805160206155858339815191525460ff1615801561230a57506000805260046020526000805160206155458339815191525460ff16155b6123265760405162461bcd60e51b815260040161084090614f96565b6002600554036123485760405162461bcd60e51b815260040161084090614de3565b60026005556000546001600160a01b031633146123775760405162461bcd60e51b815260040161084090614f23565b60065460405163b1f05c5760e01b81526001600160a01b039091169063b1f05c57906123ab90879087908790600401615035565b600060405180830381600087803b1580156123c557600080fd5b505af11580156123d9573d6000803e3d6000fd5b506123ee9250505060408301602084016148bd565b6001600160a01b03166340c10f1961240c60608501604086016148bd565b84606001356040518363ffffffff1660e01b815260040161242e9291906150a7565b6020604051808303816000875af115801561244d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124719190614f4a565b61105c5761105c6147e2565b600080516020615525833981519152600081905260046020526000805160206155858339815191525460ff161580156124cf57506000805260046020526000805160206155458339815191525460ff16155b6124eb5760405162461bcd60e51b815260040161084090614f96565b60026005540361250d5760405162461bcd60e51b815260040161084090614de3565b60026005556000546001600160a01b0316331461253c5760405162461bcd60e51b815260040161084090614f23565b600061254e60408401602085016148bd565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b8152600401602060405180830381865afa15801561258b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125af9190614fcd565b6001600160a01b0316036125d55760405162461bcd60e51b815260040161084090615237565b60065460405163b1f05c5760e01b81526001600160a01b039091169063b1f05c579061260990879087908790600401615035565b600060405180830381600087803b15801561262357600080fd5b505af1158015612637573d6000803e3d6000fd5b5061264c9250505060408301602084016148bd565b6001600160a01b03166340c10f193084606001356040518363ffffffff1660e01b815260040161267d9291906150a7565b6020604051808303816000875af115801561269c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126c09190614f4a565b6126cc576126cc6147e2565b6126dc60408301602084016148bd565b6001600160a01b031662f714ce606084018035906126fd90604087016148bd565b6040516001600160e01b031960e085901b16815260048101929092526001600160a01b031660248201526044016020604051808303816000875af1158015612749573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061276d9190614eac565b5061106c60608301604084016148bd565b6002600554036127a05760405162461bcd60e51b815260040161084090614de3565b60026005556000546001600160a01b031633146127cf5760405162461bcd60e51b815260040161084090614f23565b600680546001600160a01b0319166001600160a01b03929092169190911790556001600555565b6001546001600160a01b031633148061282f57506000546001600160a01b03163314801561282f57506001546001600160a01b03163b15155b6128735760405162461bcd60e51b81526020600482015260156024820152744d50433a206f6e6c792070656e64696e67206d706360581b6044820152606401610840565b600060025411801561288757506002544210155b6128d35760405162461bcd60e51b815260206004820152601960248201527f4d50433a2074696d65206265666f72652064656c61794d5043000000000000006044820152606401610840565b6001546000546040514281526001600160a01b0392831692909116907f8d32c9dd498e08090b44a0f77fe9ec0278851f9dffc4b430428411243e7df0769060200160405180910390a360018054600080546001600160a01b03199081166001600160a01b038416178255909116909155600255565b6000805160206155c5833981519152600081905260046020526000805160206155658339815191525460ff1615801561299a57506000805260046020526000805160206155458339815191525460ff16155b6129b65760405162461bcd60e51b815260040161084090614f96565b6002600554036129d85760405162461bcd60e51b815260040161084090614de3565b600260055560065460405163078e2c7d60e51b81526000916001600160a01b03169063f1c58fa090612a18908a9033908b908b908b908b90600401614e43565b6020604051808303816000875af1158015612a37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a5b9190614eac565b604051632770a7eb60e21b81529091506001600160a01b03881690639dc29fac90612a8c90339088906004016150a7565b6020604051808303816000875af1158015612aab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612acf9190614f4a565b612adb57612adb6147e2565b336001600160a01b0316876001600160a01b0316827f0d969ae475ff6fcaf0dcfa760d4d8607244e8d95e9bf426f8d5d69f9a3e525af898989896040516109299493929190614ec5565b600080516020615525833981519152600081905260046020526000805160206155858339815191525460ff16158015612b7757506000805260046020526000805160206155458339815191525460ff16155b612b935760405162461bcd60e51b815260040161084090614f96565b7f42aeccb36e4cd8c38ec0b9ee052287345afef9d7d5211d495f4abc7e1950eb26600081905260046020527f87fe80ca05306c1534894bf0b3d49ed5cd5a64cfca7797f3e9903420e0675be65460ff16158015612c0957506000805260046020526000805160206155458339815191525460ff16155b612c255760405162461bcd60e51b815260040161084090614f96565b600260055403612c475760405162461bcd60e51b815260040161084090614de3565b60026005556000546001600160a01b03163314612c765760405162461bcd60e51b815260040161084090614f23565b6001600160a01b03851660009081526007602052604090205460ff16612cd95760405162461bcd60e51b8152602060048201526018602482015277756e737570706f7274656420616e63616c6c2070726f787960401b6044820152606401610840565b60065460405163b1f05c5760e01b81526001600160a01b039091169063b1f05c5790612d0d908b908b908b90600401615035565b600060405180830381600087803b158015612d2757600080fd5b505af1158015612d3b573d6000803e3d6000fd5b50505050600080876020016020810190612d5591906148bd565b6001600160a01b0316636f307dc36040518163ffffffff1660e01b8152600401602060405180830381865afa158015612d92573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612db69190614fcd565b90506001600160a01b038116612dde5760405162461bcd60e51b815260040161084090615237565b60608801356001600160a01b0382166370a08231612e0260408c0160208d016148bd565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602401602060405180830381865afa158015612e46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e6a9190614eac565b10612f8e57905080612e826040890160208a016148bd565b6001600160a01b03166340c10f19308a606001356040518363ffffffff1660e01b8152600401612eb39291906150a7565b6020604051808303816000875af1158015612ed2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ef69190614f4a565b612f0257612f026147e2565b612f126040890160208a016148bd565b604051627b8a6760e11b815260608a013560048201526001600160a01b038981166024830152919091169062f714ce906044016020604051808303816000875af1158015612f64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f889190614eac565b50613179565b6001600160a01b038716600090815260076020526040902054610100900460ff161561305b57612fc46040890160208a016148bd565b9150612fd66040890160208a016148bd565b6001600160a01b03166340c10f19888a606001356040518363ffffffff1660e01b81526004016130079291906150a7565b6020604051808303816000875af1158015613026573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061304a9190614f4a565b613056576130566147e2565b613179565b60008a8a8a3561307160408d0160208e016148bd565b61308160608e0160408f016148bd565b8d606001358e608001358e8e8e6040516020016130a79a999897969594939291906151b2565b6040516020818303038152906040528051906020012090508a8a88886040516020016130d69493929190615210565b60408051601f198184030181529181528151602092830120600084815260088452829020557f2d044017b61f24f5423ce5e0c62f9ead27cb38f1615069e703ba521d0b04696b918d918d918d359161313391908f01908f016148bd565b8d604001602081019061314691906148bd565b8e606001358f608001358f8f8f6040516131699a999897969594939291906151b2565b60405180910390a15050506132b5565b50600060606001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016635b5120f789856131be8d860160408f016148bd565b8d606001358c8c6040518763ffffffff1660e01b81526004016131e696959493929190615278565b6000604051808303816000875af192505050801561322657506040513d6000823e601f3d908101601f1916820160405261322391908101906152fa565b60015b156132315790925090505b61324160608a0160408b016148bd565b6001600160a01b031661325a60408b0160208c016148bd565b6001600160a01b03168a600001357f603ea9944a12c4ef108a97399c705891f182d169a361b6aa6455d14aa1cdd2588e8e8e606001358f6080013589896040516132a99695949392919061544f565b60405180910390a45050505b50506001600555505050505050565b6002600554036132e65760405162461bcd60e51b815260040161084090614de3565b60026005556003546001600160a01b031633146133155760405162461bcd60e51b815260040161084090614eec565b60005b81811015612228576007600084848481811061333657613336615496565b905060200201602081019061334b91906148bd565b6001600160a01b031681526020810191909152604001600020805461ffff1916905580613377816154ac565b915050613318565b6000805160206155a5833981519152600081905260046020527f2bc77a4137c409d5e5d384844713fbba0af94ceb3bee1f816ae2b8a214afd84f5460ff161580156133e357506000805260046020526000805160206155458339815191525460ff16155b6133ff5760405162461bcd60e51b815260040161084090614f96565b6002600554036134215760405162461bcd60e51b815260040161084090614de3565b6002600555816134435760405162461bcd60e51b8152600401610840906150c0565b600061344f8b89613cb3565b90506000600660009054906101000a90046001600160a01b03166001600160a01b031663f1c58fa08d338e8e878e8e8e8e8e6040518b63ffffffff1660e01b81526004016134a69a999897969594939291906150e9565b6020604051808303816000875af11580156134c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134e99190614eac565b9050336001600160a01b03168c6001600160a01b0316827f968608314ec29f6fd1a9f6ef9e96247a4da1a683917569706e2d2b60ca7c0a6d8e8e878e8e8e8e8e60405161137f98979695949392919061515b565b60026005540361355f5760405162461bcd60e51b815260040161084090614de3565b60026005556003546001600160a01b0316331461358e5760405162461bcd60e51b815260040161084090614eec565b828181146135d05760405162461bcd60e51b815260206004820152600f60248201526e0d8cadccee8d040dad2e6dac2e8c6d608b1b6044820152606401610840565b60005b8181101561369957604051806040016040528060011515815260200185858481811061360157613601615496565b905060200201602081019061361691906154c5565b151590526007600088888581811061363057613630615496565b905060200201602081019061364591906148bd565b6001600160a01b0316815260208082019290925260400160002082518154939092015115156101000261ff00199215159290921661ffff199093169290921717905580613691816154ac565b9150506135d3565b5050600160055550505050565b6000805160206155a5833981519152600081905260046020527f2bc77a4137c409d5e5d384844713fbba0af94ceb3bee1f816ae2b8a214afd84f5460ff1615801561370a57506000805260046020526000805160206155458339815191525460ff16155b6137265760405162461bcd60e51b815260040161084090614f96565b6002600554036137485760405162461bcd60e51b815260040161084090614de3565b60026005558161376a5760405162461bcd60e51b8152600401610840906150c0565b60006137758a613fb3565b90506000600660009054906101000a90046001600160a01b03166001600160a01b031663f1c58fa08c338d8d878e8e8e8e8e6040518b63ffffffff1660e01b81526004016137cc9a999897969594939291906150e9565b6020604051808303816000875af11580156137eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061380f9190614eac565b9050336001600160a01b03168b6001600160a01b0316827f968608314ec29f6fd1a9f6ef9e96247a4da1a683917569706e2d2b60ca7c0a6d8d8d878e8e8e8e8e60405161386398979695949392919061515b565b60405180910390a450506001600555505050505050505050565b6003546001600160a01b031633146138a75760405162461bcd60e51b815260040161084090614eec565b6109728161446c565b600080516020615525833981519152600081905260046020526000805160206155858339815191525460ff1615801561390257506000805260046020526000805160206155458339815191525460ff16155b61391e5760405162461bcd60e51b815260040161084090614f96565b7f42aeccb36e4cd8c38ec0b9ee052287345afef9d7d5211d495f4abc7e1950eb26600081905260046020527f87fe80ca05306c1534894bf0b3d49ed5cd5a64cfca7797f3e9903420e0675be65460ff1615801561399457506000805260046020526000805160206155458339815191525460ff16155b6139b05760405162461bcd60e51b815260040161084090614f96565b6002600554036139d25760405162461bcd60e51b815260040161084090614de3565b60026005556000546001600160a01b03163314613a015760405162461bcd60e51b815260040161084090614f23565b6001600160a01b03851660009081526007602052604090205460ff16613a645760405162461bcd60e51b8152602060048201526018602482015277756e737570706f7274656420616e63616c6c2070726f787960401b6044820152606401610840565b60065460405163b1f05c5760e01b81526001600160a01b039091169063b1f05c5790613a98908b908b908b90600401615035565b600060405180830381600087803b158015613ab257600080fd5b505af1158015613ac6573d6000803e3d6000fd5b50613adb9250505060408701602088016148bd565b6001600160a01b03166340c10f198688606001356040518363ffffffff1660e01b8152600401613b0c9291906150a7565b6020604051808303816000875af1158015613b2b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b4f9190614f4a565b613b5b57613b5b6147e2565b600060606001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016635b5120f788613b9f60408c0160208d016148bd565b613baf60608d0160408e016148bd565b8c606001358b8b6040518763ffffffff1660e01b8152600401613bd796959493929190615278565b6000604051808303816000875af1925050508015613c1757506040513d6000823e601f3d908101601f19168201604052613c1491908101906152fa565b60015b15613c225790925090505b613c326060890160408a016148bd565b6001600160a01b0316613c4b60408a0160208b016148bd565b6001600160a01b031689600001357f603ea9944a12c4ef108a97399c705891f182d169a361b6aa6455d14aa1cdd2588d8d8d606001358e608001358989604051613c9a9695949392919061544f565b60405180910390a4505060016005555050505050505050565b6000805160206155c5833981519152600081815260046020526000805160206155658339815191525490919060ff16158015613d0857506000805260046020526000805160206155458339815191525460ff16155b613d245760405162461bcd60e51b815260040161084090614f96565b6000846001600160a01b0316636f307dc36040518163ffffffff1660e01b8152600401602060405180830381865afa158015613d64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d889190614fcd565b90506001600160a01b038116613db05760405162461bcd60e51b815260040161084090615237565b6040516370a0823160e01b81526001600160a01b038681166004830152600091908316906370a0823190602401602060405180830381865afa158015613dfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e1e9190614eac565b9050613e356001600160a01b03831633888861450d565b6040516370a0823160e01b81526001600160a01b038781166004830152600091908416906370a0823190602401602060405180830381865afa158015613e7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ea39190614eac565b9050818110158015613ebe5750613eba8683614f7d565b8111155b613ec757600080fd5b613ed182826154e2565b979650505050505050565b600081815260046020526040902054819060ff1680613f1357506000805260046020526000805160206155458339815191525460ff165b613f5f5760405162461bcd60e51b815260206004820152601b60248201527f5061757361626c65436f6e74726f6c3a206e6f742070617573656400000000006044820152606401610840565b60008281526004602052604090819020805460ff19169055517fd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a90613fa79084815260200190565b60405180910390a15050565b6000805160206155c5833981519152600081815260046020526000805160206155658339815191525490919060ff1615801561400857506000805260046020526000805160206155458339815191525460ff16155b6140245760405162461bcd60e51b815260040161084090614f96565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661409a5760405162461bcd60e51b815260206004820152601e60248201527f4d756c7469636861696e526f757465723a207a65726f20774e415449564500006044820152606401610840565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316836001600160a01b0316636f307dc36040518163ffffffff1660e01b8152600401602060405180830381865afa158015614102573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141269190614fcd565b6001600160a01b03161461414c5760405162461bcd60e51b815260040161084090614fea565b6040516370a0823160e01b81526001600160a01b0384811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa1580156141b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141db9190614eac565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b15801561423857600080fd5b505af115801561424c573d6000803e3d6000fd5b506142889350506001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691508690503461457e565b6040516370a0823160e01b81526001600160a01b0385811660048301526000917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa1580156142f3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906143179190614eac565b9050818110158015614332575061432e3483614f7d565b8111155b61433b57600080fd5b61434582826154e2565b95945050505050565b8047101561439e5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610840565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146143eb576040519150601f19603f3d011682016040523d82523d6000602084013e6143f0565b606091505b50509050806144675760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610840565b505050565b600081815260046020526040902054819060ff161580156144a657506000805260046020526000805160206155458339815191525460ff16155b6144c25760405162461bcd60e51b815260040161084090614f96565b60008281526004602052604090819020805460ff19166001179055517f0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba3190613fa79084815260200190565b6040516001600160a01b03808516602483015283166044820152606481018290526145789085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261459d565b50505050565b6144678363a9059cbb60e01b84846040516024016145419291906150a7565b60006145f2826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661466f9092919063ffffffff16565b80519091501561446757808060200190518101906146109190614f4a565b6144675760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610840565b606061467e8484600085614688565b90505b9392505050565b6060824710156146e95760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610840565b6001600160a01b0385163b6147405760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610840565b600080866001600160a01b0316858760405161475c91906154f5565b60006040518083038185875af1925050503d8060008114614799576040519150601f19603f3d011682016040523d82523d6000602084013e61479e565b606091505b5091509150613ed1828286606083156147b8575081614681565b8251156147c85782518084602001fd5b8160405162461bcd60e51b81526004016108409190615511565b634e487b7160e01b600052600160045260246000fd5b6001600160a01b038116811461097257600080fd5b60008083601f84011261481f57600080fd5b50813567ffffffffffffffff81111561483757600080fd5b60208301915083602082850101111561484f57600080fd5b9250929050565b60008060008060006080868803121561486e57600080fd5b8535614879816147f8565b9450602086013567ffffffffffffffff81111561489557600080fd5b6148a18882890161480d565b9699909850959660408101359660609091013595509350505050565b6000602082840312156148cf57600080fd5b8135614681816147f8565b6000602082840312156148ec57600080fd5b5035919050565b6000806040838503121561490657600080fd5b8235614911816147f8565b91506020830135614921816147f8565b809150509250929050565b6000806000806060858703121561494257600080fd5b843561494d816147f8565b9350602085013567ffffffffffffffff81111561496957600080fd5b6149758782880161480d565b9598909750949560400135949350505050565b600060a0828403121561499a57600080fd5b50919050565b600080600060c084860312156149b557600080fd5b833567ffffffffffffffff8111156149cc57600080fd5b6149d88682870161480d565b90945092506149ec90508560208601614988565b90509250925092565b600080600080600080600080600060c08a8c031215614a1357600080fd5b8935614a1e816147f8565b985060208a013567ffffffffffffffff80821115614a3b57600080fd5b614a478d838e0161480d565b909a50985060408c0135975060608c0135965060808c0135915080821115614a6e57600080fd5b614a7a8d838e0161480d565b909650945060a08c0135915080821115614a9357600080fd5b50614aa08c828d0161480d565b915080935050809150509295985092959850929598565b801515811461097257600080fd5b6000806000806000806000610120888a031215614ae157600080fd5b873567ffffffffffffffff80821115614af957600080fd5b614b058b838c0161480d565b9099509750879150614b1a8b60208c01614988565b965060c08a01359150614b2c826147f8565b90945060e08901359080821115614b4257600080fd5b50614b4f8a828b0161480d565b909450925050610100880135614b6481614ab7565b8091505092959891949750929550565b60008060408385031215614b8757600080fd5b8235614b92816147f8565b946020939093013593505050565b6000806000806000806101008789031215614bba57600080fd5b863567ffffffffffffffff80821115614bd257600080fd5b614bde8a838b0161480d565b9098509650869150614bf38a60208b01614988565b955060c08901359150614c05826147f8565b90935060e08801359080821115614c1b57600080fd5b50614c2889828a0161480d565b979a9699509497509295939492505050565b60008083601f840112614c4c57600080fd5b50813567ffffffffffffffff811115614c6457600080fd5b6020830191508360208260051b850101111561484f57600080fd5b60008060208385031215614c9257600080fd5b823567ffffffffffffffff811115614ca957600080fd5b614cb585828601614c3a565b90969095509350505050565b60008060008060408587031215614cd757600080fd5b843567ffffffffffffffff80821115614cef57600080fd5b614cfb88838901614c3a565b90965094506020870135915080821115614d1457600080fd5b50614d2187828801614c3a565b95989497509550505050565b60008060008060008060008060a0898b031215614d4957600080fd5b8835614d54816147f8565b9750602089013567ffffffffffffffff80821115614d7157600080fd5b614d7d8c838d0161480d565b909950975060408b0135965060608b0135915080821115614d9d57600080fd5b614da98c838d0161480d565b909650945060808b0135915080821115614dc257600080fd5b50614dcf8b828c0161480d565b999c989b5096995094979396929594505050565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6001600160a01b0387811682528616602082015260e060408201819052600090614e709083018688614e1a565b8460608401528360808401528281038060a0850152600082526020810160c0850152506000602082015260408101915050979650505050505050565b600060208284031215614ebe57600080fd5b5051919050565b606081526000614ed9606083018688614e1a565b6020830194909452506040015292915050565b6020808252601a908201527f4d504341646d696e436f6e74726f6c3a206e6f742061646d696e000000000000604082015260600190565b6020808252600d908201526c4d50433a206f6e6c79206d706360981b604082015260600190565b600060208284031215614f5c57600080fd5b815161468181614ab7565b634e487b7160e01b600052601160045260246000fd5b80820180821115614f9057614f90614f67565b92915050565b60208082526017908201527f5061757361626c65436f6e74726f6c3a20706175736564000000000000000000604082015260600190565b600060208284031215614fdf57600080fd5b8151614681816147f8565b6020808252602b908201527f4d756c7469636861696e526f757465723a20756e6465726c79696e672069732060408201526a6e6f7420774e415449564560a81b606082015260800190565b60c08152600061504960c083018587614e1a565b9050823560208301526020830135615060816147f8565b6001600160a01b0390811660408481019190915284013590615081826147f8565b8082166060850152505060608301356080830152608083013560a0830152949350505050565b6001600160a01b03929092168252602082015260400190565b6020808252600f908201526e656d7074792063616c6c206461746160881b604082015260600190565b6001600160a01b038b811682528a16602082015260e0604082018190526000906151169083018a8c614e1a565b88606084015287608084015282810360a0840152615135818789614e1a565b905082810360c084015261514a818587614e1a565b9d9c50505050505050505050505050565b60a08152600061516f60a083018a8c614e1a565b886020840152876040840152828103606084015261518e818789614e1a565b905082810360808401526151a3818587614e1a565b9b9a5050505050505050505050565b60006101008083526151c78184018d8f614e1a565b602084018c90526001600160a01b038b811660408601528a81166060860152608085018a905260a08501899052871660c085015283810360e0850152905061514a818587614e1a565b604081526000615224604083018688614e1a565b8281036020840152613ed1818587614e1a565b60208082526021908201527f4d756c7469636861696e526f757465723a207a65726f20756e6465726c79696e6040820152606760f81b606082015260800190565b6001600160a01b0387811682528681166020830152851660408201526060810184905260a0608082018190526000906152b49083018486614e1a565b98975050505050505050565b634e487b7160e01b600052604160045260246000fd5b60005b838110156152f15781810151838201526020016152d9565b50506000910152565b6000806040838503121561530d57600080fd5b825161531881614ab7565b602084015190925067ffffffffffffffff8082111561533657600080fd5b818501915085601f83011261534a57600080fd5b81518181111561535c5761535c6152c0565b604051601f8201601f19908116603f01168101908382118183101715615384576153846152c0565b8160405282815288602084870101111561539d57600080fd5b6153ae8360208301602088016152d6565b80955050505050509250929050565b600081518084526153d58160208601602086016152d6565b601f01601f19169290920160200192915050565b60006101208083526153fe8184018d8f614e1a565b602084018c90526001600160a01b038b811660408601528a1660608501526080840189905260a0840188905286151560c085015285151560e0850152838103610100850152905061514a81856153bd565b60a08152600061546360a08301888a614e1a565b8660208401528560408401528415156060840152828103608084015261548981856153bd565b9998505050505050505050565b634e487b7160e01b600052603260045260246000fd5b6000600182016154be576154be614f67565b5060010190565b6000602082840312156154d757600080fd5b813561468181614ab7565b81810381811115614f9057614f90614f67565b600082516155078184602087016152d6565b9190910192915050565b60208152600061468160208301846153bd56fe9246b6a221bc8c80334eddce2febc232b5cab9fba8e96ff92acabffc5920ef3217ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec03e75eeefe1c4ab02b60c94ba8379778a686f441b318b906786c58dea7fab6500f29789c2adc3f170f466eb9b5f714ec7b82bbdeafb4c7e0b02581b188d6fa3ff1fb9f45325ee45a8708aacc208b1051219c6c4134780d77ec04c67aea218abf6353028a1e62f134c5974e9ee55a946683badabdaaaa0bfabd235a446273e047a2646970667358221220695d8c9f958d7f717b5f82be9e78938cf628c69c9d2db63ef2a9879cdbdc8b5964736f6c63430008110033000000000000000000000000fa9da51631268a30ec3ddd1ccbf46c65fad99251000000000000000000000000b3f91051cd787a0d14ef806fbb7a11c8479668690000000000000000000000000bd19f6447cd676255c7c7b00428462b3da67e3a0000000000000000000000006afcff9189e8ed3fcc1cffa184feb1276f6a82a5000000000000000000000000e56979f6ada241c1bed92e68535dcead9de2a5ef",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _admin common.Address, _mpc common.Address, _wNATIVE common.Address, _anycallExecutor common.Address, _routerSecurity common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _admin, _mpc, _wNATIVE, _anycallExecutor, _routerSecurity)
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

// NewContract creates a new instance of Contracts, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
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

// CallPausedROLE is a free data retrieval call binding the contract method 0x5598f119.
//
// Solidity: function Call_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) CallPausedROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "Call_Paused_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CallPausedROLE is a free data retrieval call binding the contract method 0x5598f119.
//
// Solidity: function Call_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) CallPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.CallPausedROLE(&_Contracts.CallOpts)
}

// CallPausedROLE is a free data retrieval call binding the contract method 0x5598f119.
//
// Solidity: function Call_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) CallPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.CallPausedROLE(&_Contracts.CallOpts)
}

// ExecPausedROLE is a free data retrieval call binding the contract method 0xe94b7144.
//
// Solidity: function Exec_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) ExecPausedROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "Exec_Paused_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ExecPausedROLE is a free data retrieval call binding the contract method 0xe94b7144.
//
// Solidity: function Exec_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) ExecPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.ExecPausedROLE(&_Contracts.CallOpts)
}

// ExecPausedROLE is a free data retrieval call binding the contract method 0xe94b7144.
//
// Solidity: function Exec_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) ExecPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.ExecPausedROLE(&_Contracts.CallOpts)
}

// PAUSEALLROLE is a free data retrieval call binding the contract method 0x9ac25d08.
//
// Solidity: function PAUSE_ALL_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) PAUSEALLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "PAUSE_ALL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEALLROLE is a free data retrieval call binding the contract method 0x9ac25d08.
//
// Solidity: function PAUSE_ALL_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) PAUSEALLROLE() ([32]byte, error) {
	return _Contracts.Contract.PAUSEALLROLE(&_Contracts.CallOpts)
}

// PAUSEALLROLE is a free data retrieval call binding the contract method 0x9ac25d08.
//
// Solidity: function PAUSE_ALL_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) PAUSEALLROLE() ([32]byte, error) {
	return _Contracts.Contract.PAUSEALLROLE(&_Contracts.CallOpts)
}

// RetryPausedROLE is a free data retrieval call binding the contract method 0x912d857c.
//
// Solidity: function Retry_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) RetryPausedROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "Retry_Paused_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RetryPausedROLE is a free data retrieval call binding the contract method 0x912d857c.
//
// Solidity: function Retry_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) RetryPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.RetryPausedROLE(&_Contracts.CallOpts)
}

// RetryPausedROLE is a free data retrieval call binding the contract method 0x912d857c.
//
// Solidity: function Retry_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) RetryPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.RetryPausedROLE(&_Contracts.CallOpts)
}

// SwapinPausedROLE is a free data retrieval call binding the contract method 0xf91275b5.
//
// Solidity: function Swapin_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) SwapinPausedROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "Swapin_Paused_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SwapinPausedROLE is a free data retrieval call binding the contract method 0xf91275b5.
//
// Solidity: function Swapin_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) SwapinPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.SwapinPausedROLE(&_Contracts.CallOpts)
}

// SwapinPausedROLE is a free data retrieval call binding the contract method 0xf91275b5.
//
// Solidity: function Swapin_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) SwapinPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.SwapinPausedROLE(&_Contracts.CallOpts)
}

// SwapoutPausedROLE is a free data retrieval call binding the contract method 0x0c55b22e.
//
// Solidity: function Swapout_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCaller) SwapoutPausedROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "Swapout_Paused_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SwapoutPausedROLE is a free data retrieval call binding the contract method 0x0c55b22e.
//
// Solidity: function Swapout_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsSession) SwapoutPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.SwapoutPausedROLE(&_Contracts.CallOpts)
}

// SwapoutPausedROLE is a free data retrieval call binding the contract method 0x0c55b22e.
//
// Solidity: function Swapout_Paused_ROLE() view returns(bytes32)
func (_Contracts *ContractsCallerSession) SwapoutPausedROLE() ([32]byte, error) {
	return _Contracts.Contract.SwapoutPausedROLE(&_Contracts.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contracts *ContractsCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contracts *ContractsSession) Admin() (common.Address, error) {
	return _Contracts.Contract.Admin(&_Contracts.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Contracts *ContractsCallerSession) Admin() (common.Address, error) {
	return _Contracts.Contract.Admin(&_Contracts.CallOpts)
}

// AnycallExecutor is a free data retrieval call binding the contract method 0xd2c7dfcc.
//
// Solidity: function anycallExecutor() view returns(address)
func (_Contracts *ContractsCaller) AnycallExecutor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "anycallExecutor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AnycallExecutor is a free data retrieval call binding the contract method 0xd2c7dfcc.
//
// Solidity: function anycallExecutor() view returns(address)
func (_Contracts *ContractsSession) AnycallExecutor() (common.Address, error) {
	return _Contracts.Contract.AnycallExecutor(&_Contracts.CallOpts)
}

// AnycallExecutor is a free data retrieval call binding the contract method 0xd2c7dfcc.
//
// Solidity: function anycallExecutor() view returns(address)
func (_Contracts *ContractsCallerSession) AnycallExecutor() (common.Address, error) {
	return _Contracts.Contract.AnycallExecutor(&_Contracts.CallOpts)
}

// AnycallProxyInfo is a free data retrieval call binding the contract method 0x1d5aa281.
//
// Solidity: function anycallProxyInfo(address ) view returns(bool supported, bool acceptAnyToken)
func (_Contracts *ContractsCaller) AnycallProxyInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Supported      bool
	AcceptAnyToken bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "anycallProxyInfo", arg0)

	outstruct := new(struct {
		Supported      bool
		AcceptAnyToken bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Supported = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.AcceptAnyToken = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// AnycallProxyInfo is a free data retrieval call binding the contract method 0x1d5aa281.
//
// Solidity: function anycallProxyInfo(address ) view returns(bool supported, bool acceptAnyToken)
func (_Contracts *ContractsSession) AnycallProxyInfo(arg0 common.Address) (struct {
	Supported      bool
	AcceptAnyToken bool
}, error) {
	return _Contracts.Contract.AnycallProxyInfo(&_Contracts.CallOpts, arg0)
}

// AnycallProxyInfo is a free data retrieval call binding the contract method 0x1d5aa281.
//
// Solidity: function anycallProxyInfo(address ) view returns(bool supported, bool acceptAnyToken)
func (_Contracts *ContractsCallerSession) AnycallProxyInfo(arg0 common.Address) (struct {
	Supported      bool
	AcceptAnyToken bool
}, error) {
	return _Contracts.Contract.AnycallProxyInfo(&_Contracts.CallOpts, arg0)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Contracts *ContractsCaller) Delay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "delay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Contracts *ContractsSession) Delay() (*big.Int, error) {
	return _Contracts.Contract.Delay(&_Contracts.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() view returns(uint256)
func (_Contracts *ContractsCallerSession) Delay() (*big.Int, error) {
	return _Contracts.Contract.Delay(&_Contracts.CallOpts)
}

// DelayMPC is a free data retrieval call binding the contract method 0x160f1053.
//
// Solidity: function delayMPC() view returns(uint256)
func (_Contracts *ContractsCaller) DelayMPC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "delayMPC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayMPC is a free data retrieval call binding the contract method 0x160f1053.
//
// Solidity: function delayMPC() view returns(uint256)
func (_Contracts *ContractsSession) DelayMPC() (*big.Int, error) {
	return _Contracts.Contract.DelayMPC(&_Contracts.CallOpts)
}

// DelayMPC is a free data retrieval call binding the contract method 0x160f1053.
//
// Solidity: function delayMPC() view returns(uint256)
func (_Contracts *ContractsCallerSession) DelayMPC() (*big.Int, error) {
	return _Contracts.Contract.DelayMPC(&_Contracts.CallOpts)
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

// Paused is a free data retrieval call binding the contract method 0x9e9e4666.
//
// Solidity: function paused(bytes32 role) view returns(bool)
func (_Contracts *ContractsCaller) Paused(opts *bind.CallOpts, role [32]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "paused", role)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x9e9e4666.
//
// Solidity: function paused(bytes32 role) view returns(bool)
func (_Contracts *ContractsSession) Paused(role [32]byte) (bool, error) {
	return _Contracts.Contract.Paused(&_Contracts.CallOpts, role)
}

// Paused is a free data retrieval call binding the contract method 0x9e9e4666.
//
// Solidity: function paused(bytes32 role) view returns(bool)
func (_Contracts *ContractsCallerSession) Paused(role [32]byte) (bool, error) {
	return _Contracts.Contract.Paused(&_Contracts.CallOpts, role)
}

// PendingMPC is a free data retrieval call binding the contract method 0xf830e7b4.
//
// Solidity: function pendingMPC() view returns(address)
func (_Contracts *ContractsCaller) PendingMPC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "pendingMPC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingMPC is a free data retrieval call binding the contract method 0xf830e7b4.
//
// Solidity: function pendingMPC() view returns(address)
func (_Contracts *ContractsSession) PendingMPC() (common.Address, error) {
	return _Contracts.Contract.PendingMPC(&_Contracts.CallOpts)
}

// PendingMPC is a free data retrieval call binding the contract method 0xf830e7b4.
//
// Solidity: function pendingMPC() view returns(address)
func (_Contracts *ContractsCallerSession) PendingMPC() (common.Address, error) {
	return _Contracts.Contract.PendingMPC(&_Contracts.CallOpts)
}

// RetryRecords is a free data retrieval call binding the contract method 0x6a6459d1.
//
// Solidity: function retryRecords(bytes32 ) view returns(bytes32)
func (_Contracts *ContractsCaller) RetryRecords(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "retryRecords", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RetryRecords is a free data retrieval call binding the contract method 0x6a6459d1.
//
// Solidity: function retryRecords(bytes32 ) view returns(bytes32)
func (_Contracts *ContractsSession) RetryRecords(arg0 [32]byte) ([32]byte, error) {
	return _Contracts.Contract.RetryRecords(&_Contracts.CallOpts, arg0)
}

// RetryRecords is a free data retrieval call binding the contract method 0x6a6459d1.
//
// Solidity: function retryRecords(bytes32 ) view returns(bytes32)
func (_Contracts *ContractsCallerSession) RetryRecords(arg0 [32]byte) ([32]byte, error) {
	return _Contracts.Contract.RetryRecords(&_Contracts.CallOpts, arg0)
}

// RouterSecurity is a free data retrieval call binding the contract method 0xa413387a.
//
// Solidity: function routerSecurity() view returns(address)
func (_Contracts *ContractsCaller) RouterSecurity(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "routerSecurity")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RouterSecurity is a free data retrieval call binding the contract method 0xa413387a.
//
// Solidity: function routerSecurity() view returns(address)
func (_Contracts *ContractsSession) RouterSecurity() (common.Address, error) {
	return _Contracts.Contract.RouterSecurity(&_Contracts.CallOpts)
}

// RouterSecurity is a free data retrieval call binding the contract method 0xa413387a.
//
// Solidity: function routerSecurity() view returns(address)
func (_Contracts *ContractsCallerSession) RouterSecurity() (common.Address, error) {
	return _Contracts.Contract.RouterSecurity(&_Contracts.CallOpts)
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

// AddAnycallProxies is a paid mutator transaction binding the contract method 0xe2ea2ba9.
//
// Solidity: function addAnycallProxies(address[] proxies, bool[] acceptAnyTokenFlags) returns()
func (_Contracts *ContractsTransactor) AddAnycallProxies(opts *bind.TransactOpts, proxies []common.Address, acceptAnyTokenFlags []bool) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addAnycallProxies", proxies, acceptAnyTokenFlags)
}

// AddAnycallProxies is a paid mutator transaction binding the contract method 0xe2ea2ba9.
//
// Solidity: function addAnycallProxies(address[] proxies, bool[] acceptAnyTokenFlags) returns()
func (_Contracts *ContractsSession) AddAnycallProxies(proxies []common.Address, acceptAnyTokenFlags []bool) (*types.Transaction, error) {
	return _Contracts.Contract.AddAnycallProxies(&_Contracts.TransactOpts, proxies, acceptAnyTokenFlags)
}

// AddAnycallProxies is a paid mutator transaction binding the contract method 0xe2ea2ba9.
//
// Solidity: function addAnycallProxies(address[] proxies, bool[] acceptAnyTokenFlags) returns()
func (_Contracts *ContractsTransactorSession) AddAnycallProxies(proxies []common.Address, acceptAnyTokenFlags []bool) (*types.Transaction, error) {
	return _Contracts.Contract.AddAnycallProxies(&_Contracts.TransactOpts, proxies, acceptAnyTokenFlags)
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

// AnySwapIn is a paid mutator transaction binding the contract method 0x8fef8489.
//
// Solidity: function anySwapIn(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsTransactor) AnySwapIn(opts *bind.TransactOpts, swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapIn", swapID, swapInfo)
}

// AnySwapIn is a paid mutator transaction binding the contract method 0x8fef8489.
//
// Solidity: function anySwapIn(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsSession) AnySwapIn(swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapIn(&_Contracts.TransactOpts, swapID, swapInfo)
}

// AnySwapIn is a paid mutator transaction binding the contract method 0x8fef8489.
//
// Solidity: function anySwapIn(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsTransactorSession) AnySwapIn(swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapIn(&_Contracts.TransactOpts, swapID, swapInfo)
}

// AnySwapInAndExec is a paid mutator transaction binding the contract method 0xf9ca3a5d.
//
// Solidity: function anySwapInAndExec(string swapID, (bytes32,address,address,uint256,uint256) swapInfo, address anycallProxy, bytes data) returns()
func (_Contracts *ContractsTransactor) AnySwapInAndExec(opts *bind.TransactOpts, swapID string, swapInfo SwapInfo, anycallProxy common.Address, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapInAndExec", swapID, swapInfo, anycallProxy, data)
}

// AnySwapInAndExec is a paid mutator transaction binding the contract method 0xf9ca3a5d.
//
// Solidity: function anySwapInAndExec(string swapID, (bytes32,address,address,uint256,uint256) swapInfo, address anycallProxy, bytes data) returns()
func (_Contracts *ContractsSession) AnySwapInAndExec(swapID string, swapInfo SwapInfo, anycallProxy common.Address, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInAndExec(&_Contracts.TransactOpts, swapID, swapInfo, anycallProxy, data)
}

// AnySwapInAndExec is a paid mutator transaction binding the contract method 0xf9ca3a5d.
//
// Solidity: function anySwapInAndExec(string swapID, (bytes32,address,address,uint256,uint256) swapInfo, address anycallProxy, bytes data) returns()
func (_Contracts *ContractsTransactorSession) AnySwapInAndExec(swapID string, swapInfo SwapInfo, anycallProxy common.Address, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInAndExec(&_Contracts.TransactOpts, swapID, swapInfo, anycallProxy, data)
}

// AnySwapInAuto is a paid mutator transaction binding the contract method 0x81aa7a81.
//
// Solidity: function anySwapInAuto(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsTransactor) AnySwapInAuto(opts *bind.TransactOpts, swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapInAuto", swapID, swapInfo)
}

// AnySwapInAuto is a paid mutator transaction binding the contract method 0x81aa7a81.
//
// Solidity: function anySwapInAuto(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsSession) AnySwapInAuto(swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInAuto(&_Contracts.TransactOpts, swapID, swapInfo)
}

// AnySwapInAuto is a paid mutator transaction binding the contract method 0x81aa7a81.
//
// Solidity: function anySwapInAuto(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsTransactorSession) AnySwapInAuto(swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInAuto(&_Contracts.TransactOpts, swapID, swapInfo)
}

// AnySwapInNative is a paid mutator transaction binding the contract method 0x5de26385.
//
// Solidity: function anySwapInNative(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsTransactor) AnySwapInNative(opts *bind.TransactOpts, swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapInNative", swapID, swapInfo)
}

// AnySwapInNative is a paid mutator transaction binding the contract method 0x5de26385.
//
// Solidity: function anySwapInNative(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsSession) AnySwapInNative(swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInNative(&_Contracts.TransactOpts, swapID, swapInfo)
}

// AnySwapInNative is a paid mutator transaction binding the contract method 0x5de26385.
//
// Solidity: function anySwapInNative(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsTransactorSession) AnySwapInNative(swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInNative(&_Contracts.TransactOpts, swapID, swapInfo)
}

// AnySwapInUnderlying is a paid mutator transaction binding the contract method 0x9ff1d3e8.
//
// Solidity: function anySwapInUnderlying(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsTransactor) AnySwapInUnderlying(opts *bind.TransactOpts, swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapInUnderlying", swapID, swapInfo)
}

// AnySwapInUnderlying is a paid mutator transaction binding the contract method 0x9ff1d3e8.
//
// Solidity: function anySwapInUnderlying(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsSession) AnySwapInUnderlying(swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInUnderlying(&_Contracts.TransactOpts, swapID, swapInfo)
}

// AnySwapInUnderlying is a paid mutator transaction binding the contract method 0x9ff1d3e8.
//
// Solidity: function anySwapInUnderlying(string swapID, (bytes32,address,address,uint256,uint256) swapInfo) returns()
func (_Contracts *ContractsTransactorSession) AnySwapInUnderlying(swapID string, swapInfo SwapInfo) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInUnderlying(&_Contracts.TransactOpts, swapID, swapInfo)
}

// AnySwapInUnderlyingAndExec is a paid mutator transaction binding the contract method 0xcc95060a.
//
// Solidity: function anySwapInUnderlyingAndExec(string swapID, (bytes32,address,address,uint256,uint256) swapInfo, address anycallProxy, bytes data) returns()
func (_Contracts *ContractsTransactor) AnySwapInUnderlyingAndExec(opts *bind.TransactOpts, swapID string, swapInfo SwapInfo, anycallProxy common.Address, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapInUnderlyingAndExec", swapID, swapInfo, anycallProxy, data)
}

// AnySwapInUnderlyingAndExec is a paid mutator transaction binding the contract method 0xcc95060a.
//
// Solidity: function anySwapInUnderlyingAndExec(string swapID, (bytes32,address,address,uint256,uint256) swapInfo, address anycallProxy, bytes data) returns()
func (_Contracts *ContractsSession) AnySwapInUnderlyingAndExec(swapID string, swapInfo SwapInfo, anycallProxy common.Address, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInUnderlyingAndExec(&_Contracts.TransactOpts, swapID, swapInfo, anycallProxy, data)
}

// AnySwapInUnderlyingAndExec is a paid mutator transaction binding the contract method 0xcc95060a.
//
// Solidity: function anySwapInUnderlyingAndExec(string swapID, (bytes32,address,address,uint256,uint256) swapInfo, address anycallProxy, bytes data) returns()
func (_Contracts *ContractsTransactorSession) AnySwapInUnderlyingAndExec(swapID string, swapInfo SwapInfo, anycallProxy common.Address, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapInUnderlyingAndExec(&_Contracts.TransactOpts, swapID, swapInfo, anycallProxy, data)
}

// AnySwapOut is a paid mutator transaction binding the contract method 0xc604b0b8.
//
// Solidity: function anySwapOut(address token, string to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOut(opts *bind.TransactOpts, token common.Address, to string, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOut", token, to, amount, toChainID)
}

// AnySwapOut is a paid mutator transaction binding the contract method 0xc604b0b8.
//
// Solidity: function anySwapOut(address token, string to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOut(token common.Address, to string, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOut(&_Contracts.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOut is a paid mutator transaction binding the contract method 0xc604b0b8.
//
// Solidity: function anySwapOut(address token, string to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOut(token common.Address, to string, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOut(&_Contracts.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOutAndCall is a paid mutator transaction binding the contract method 0x6b4b4376.
//
// Solidity: function anySwapOutAndCall(address token, string to, uint256 amount, uint256 toChainID, string anycallProxy, bytes data) returns()
func (_Contracts *ContractsTransactor) AnySwapOutAndCall(opts *bind.TransactOpts, token common.Address, to string, amount *big.Int, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutAndCall", token, to, amount, toChainID, anycallProxy, data)
}

// AnySwapOutAndCall is a paid mutator transaction binding the contract method 0x6b4b4376.
//
// Solidity: function anySwapOutAndCall(address token, string to, uint256 amount, uint256 toChainID, string anycallProxy, bytes data) returns()
func (_Contracts *ContractsSession) AnySwapOutAndCall(token common.Address, to string, amount *big.Int, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutAndCall(&_Contracts.TransactOpts, token, to, amount, toChainID, anycallProxy, data)
}

// AnySwapOutAndCall is a paid mutator transaction binding the contract method 0x6b4b4376.
//
// Solidity: function anySwapOutAndCall(address token, string to, uint256 amount, uint256 toChainID, string anycallProxy, bytes data) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutAndCall(token common.Address, to string, amount *big.Int, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutAndCall(&_Contracts.TransactOpts, token, to, amount, toChainID, anycallProxy, data)
}

// AnySwapOutNative is a paid mutator transaction binding the contract method 0x540dd52c.
//
// Solidity: function anySwapOutNative(address token, string to, uint256 toChainID) payable returns()
func (_Contracts *ContractsTransactor) AnySwapOutNative(opts *bind.TransactOpts, token common.Address, to string, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutNative", token, to, toChainID)
}

// AnySwapOutNative is a paid mutator transaction binding the contract method 0x540dd52c.
//
// Solidity: function anySwapOutNative(address token, string to, uint256 toChainID) payable returns()
func (_Contracts *ContractsSession) AnySwapOutNative(token common.Address, to string, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutNative(&_Contracts.TransactOpts, token, to, toChainID)
}

// AnySwapOutNative is a paid mutator transaction binding the contract method 0x540dd52c.
//
// Solidity: function anySwapOutNative(address token, string to, uint256 toChainID) payable returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutNative(token common.Address, to string, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutNative(&_Contracts.TransactOpts, token, to, toChainID)
}

// AnySwapOutNativeAndCall is a paid mutator transaction binding the contract method 0xea0c968b.
//
// Solidity: function anySwapOutNativeAndCall(address token, string to, uint256 toChainID, string anycallProxy, bytes data) payable returns()
func (_Contracts *ContractsTransactor) AnySwapOutNativeAndCall(opts *bind.TransactOpts, token common.Address, to string, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutNativeAndCall", token, to, toChainID, anycallProxy, data)
}

// AnySwapOutNativeAndCall is a paid mutator transaction binding the contract method 0xea0c968b.
//
// Solidity: function anySwapOutNativeAndCall(address token, string to, uint256 toChainID, string anycallProxy, bytes data) payable returns()
func (_Contracts *ContractsSession) AnySwapOutNativeAndCall(token common.Address, to string, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutNativeAndCall(&_Contracts.TransactOpts, token, to, toChainID, anycallProxy, data)
}

// AnySwapOutNativeAndCall is a paid mutator transaction binding the contract method 0xea0c968b.
//
// Solidity: function anySwapOutNativeAndCall(address token, string to, uint256 toChainID, string anycallProxy, bytes data) payable returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutNativeAndCall(token common.Address, to string, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutNativeAndCall(&_Contracts.TransactOpts, token, to, toChainID, anycallProxy, data)
}

// AnySwapOutUnderlying is a paid mutator transaction binding the contract method 0x049b4e7e.
//
// Solidity: function anySwapOutUnderlying(address token, string to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsTransactor) AnySwapOutUnderlying(opts *bind.TransactOpts, token common.Address, to string, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutUnderlying", token, to, amount, toChainID)
}

// AnySwapOutUnderlying is a paid mutator transaction binding the contract method 0x049b4e7e.
//
// Solidity: function anySwapOutUnderlying(address token, string to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsSession) AnySwapOutUnderlying(token common.Address, to string, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlying(&_Contracts.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOutUnderlying is a paid mutator transaction binding the contract method 0x049b4e7e.
//
// Solidity: function anySwapOutUnderlying(address token, string to, uint256 amount, uint256 toChainID) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutUnderlying(token common.Address, to string, amount *big.Int, toChainID *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlying(&_Contracts.TransactOpts, token, to, amount, toChainID)
}

// AnySwapOutUnderlyingAndCall is a paid mutator transaction binding the contract method 0xe0e9048e.
//
// Solidity: function anySwapOutUnderlyingAndCall(address token, string to, uint256 amount, uint256 toChainID, string anycallProxy, bytes data) returns()
func (_Contracts *ContractsTransactor) AnySwapOutUnderlyingAndCall(opts *bind.TransactOpts, token common.Address, to string, amount *big.Int, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "anySwapOutUnderlyingAndCall", token, to, amount, toChainID, anycallProxy, data)
}

// AnySwapOutUnderlyingAndCall is a paid mutator transaction binding the contract method 0xe0e9048e.
//
// Solidity: function anySwapOutUnderlyingAndCall(address token, string to, uint256 amount, uint256 toChainID, string anycallProxy, bytes data) returns()
func (_Contracts *ContractsSession) AnySwapOutUnderlyingAndCall(token common.Address, to string, amount *big.Int, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlyingAndCall(&_Contracts.TransactOpts, token, to, amount, toChainID, anycallProxy, data)
}

// AnySwapOutUnderlyingAndCall is a paid mutator transaction binding the contract method 0xe0e9048e.
//
// Solidity: function anySwapOutUnderlyingAndCall(address token, string to, uint256 amount, uint256 toChainID, string anycallProxy, bytes data) returns()
func (_Contracts *ContractsTransactorSession) AnySwapOutUnderlyingAndCall(token common.Address, to string, amount *big.Int, toChainID *big.Int, anycallProxy string, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.AnySwapOutUnderlyingAndCall(&_Contracts.TransactOpts, token, to, amount, toChainID, anycallProxy, data)
}

// ApplyMPC is a paid mutator transaction binding the contract method 0xb63b38d0.
//
// Solidity: function applyMPC() returns()
func (_Contracts *ContractsTransactor) ApplyMPC(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "applyMPC")
}

// ApplyMPC is a paid mutator transaction binding the contract method 0xb63b38d0.
//
// Solidity: function applyMPC() returns()
func (_Contracts *ContractsSession) ApplyMPC() (*types.Transaction, error) {
	return _Contracts.Contract.ApplyMPC(&_Contracts.TransactOpts)
}

// ApplyMPC is a paid mutator transaction binding the contract method 0xb63b38d0.
//
// Solidity: function applyMPC() returns()
func (_Contracts *ContractsTransactorSession) ApplyMPC() (*types.Transaction, error) {
	return _Contracts.Contract.ApplyMPC(&_Contracts.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Contracts *ContractsTransactor) ChangeAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changeAdmin", _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Contracts *ContractsSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeAdmin(&_Contracts.TransactOpts, _admin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address _admin) returns()
func (_Contracts *ContractsTransactorSession) ChangeAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeAdmin(&_Contracts.TransactOpts, _admin)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address _mpc) returns()
func (_Contracts *ContractsTransactor) ChangeMPC(opts *bind.TransactOpts, _mpc common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changeMPC", _mpc)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address _mpc) returns()
func (_Contracts *ContractsSession) ChangeMPC(_mpc common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeMPC(&_Contracts.TransactOpts, _mpc)
}

// ChangeMPC is a paid mutator transaction binding the contract method 0x5b7b018c.
//
// Solidity: function changeMPC(address _mpc) returns()
func (_Contracts *ContractsTransactorSession) ChangeMPC(_mpc common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeMPC(&_Contracts.TransactOpts, _mpc)
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

// Pause is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 role) returns()
func (_Contracts *ContractsTransactor) Pause(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "pause", role)
}

// Pause is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 role) returns()
func (_Contracts *ContractsSession) Pause(role [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Pause(&_Contracts.TransactOpts, role)
}

// Pause is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 role) returns()
func (_Contracts *ContractsTransactorSession) Pause(role [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Pause(&_Contracts.TransactOpts, role)
}

// RemoveAnycallProxies is a paid mutator transaction binding the contract method 0xd21c1cf5.
//
// Solidity: function removeAnycallProxies(address[] proxies) returns()
func (_Contracts *ContractsTransactor) RemoveAnycallProxies(opts *bind.TransactOpts, proxies []common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeAnycallProxies", proxies)
}

// RemoveAnycallProxies is a paid mutator transaction binding the contract method 0xd21c1cf5.
//
// Solidity: function removeAnycallProxies(address[] proxies) returns()
func (_Contracts *ContractsSession) RemoveAnycallProxies(proxies []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveAnycallProxies(&_Contracts.TransactOpts, proxies)
}

// RemoveAnycallProxies is a paid mutator transaction binding the contract method 0xd21c1cf5.
//
// Solidity: function removeAnycallProxies(address[] proxies) returns()
func (_Contracts *ContractsTransactorSession) RemoveAnycallProxies(proxies []common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveAnycallProxies(&_Contracts.TransactOpts, proxies)
}

// RetrySwapinAndExec is a paid mutator transaction binding the contract method 0x872acd04.
//
// Solidity: function retrySwapinAndExec(string swapID, (bytes32,address,address,uint256,uint256) swapInfo, address anycallProxy, bytes data, bool dontExec) returns()
func (_Contracts *ContractsTransactor) RetrySwapinAndExec(opts *bind.TransactOpts, swapID string, swapInfo SwapInfo, anycallProxy common.Address, data []byte, dontExec bool) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "retrySwapinAndExec", swapID, swapInfo, anycallProxy, data, dontExec)
}

// RetrySwapinAndExec is a paid mutator transaction binding the contract method 0x872acd04.
//
// Solidity: function retrySwapinAndExec(string swapID, (bytes32,address,address,uint256,uint256) swapInfo, address anycallProxy, bytes data, bool dontExec) returns()
func (_Contracts *ContractsSession) RetrySwapinAndExec(swapID string, swapInfo SwapInfo, anycallProxy common.Address, data []byte, dontExec bool) (*types.Transaction, error) {
	return _Contracts.Contract.RetrySwapinAndExec(&_Contracts.TransactOpts, swapID, swapInfo, anycallProxy, data, dontExec)
}

// RetrySwapinAndExec is a paid mutator transaction binding the contract method 0x872acd04.
//
// Solidity: function retrySwapinAndExec(string swapID, (bytes32,address,address,uint256,uint256) swapInfo, address anycallProxy, bytes data, bool dontExec) returns()
func (_Contracts *ContractsTransactorSession) RetrySwapinAndExec(swapID string, swapInfo SwapInfo, anycallProxy common.Address, data []byte, dontExec bool) (*types.Transaction, error) {
	return _Contracts.Contract.RetrySwapinAndExec(&_Contracts.TransactOpts, swapID, swapInfo, anycallProxy, data, dontExec)
}

// SetRouterSecurity is a paid mutator transaction binding the contract method 0xa66ec443.
//
// Solidity: function setRouterSecurity(address _routerSecurity) returns()
func (_Contracts *ContractsTransactor) SetRouterSecurity(opts *bind.TransactOpts, _routerSecurity common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setRouterSecurity", _routerSecurity)
}

// SetRouterSecurity is a paid mutator transaction binding the contract method 0xa66ec443.
//
// Solidity: function setRouterSecurity(address _routerSecurity) returns()
func (_Contracts *ContractsSession) SetRouterSecurity(_routerSecurity common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetRouterSecurity(&_Contracts.TransactOpts, _routerSecurity)
}

// SetRouterSecurity is a paid mutator transaction binding the contract method 0xa66ec443.
//
// Solidity: function setRouterSecurity(address _routerSecurity) returns()
func (_Contracts *ContractsTransactorSession) SetRouterSecurity(_routerSecurity common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetRouterSecurity(&_Contracts.TransactOpts, _routerSecurity)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 role) returns()
func (_Contracts *ContractsTransactor) Unpause(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unpause", role)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 role) returns()
func (_Contracts *ContractsSession) Unpause(role [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Unpause(&_Contracts.TransactOpts, role)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 role) returns()
func (_Contracts *ContractsTransactorSession) Unpause(role [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Unpause(&_Contracts.TransactOpts, role)
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

// ContractsChangeAdminIterator is returned from FilterChangeAdmin and is used to iterate over the raw logs and unpacked data for ChangeAdmin events raised by the Contracts contract.
type ContractsChangeAdminIterator struct {
	Event *ContractsChangeAdmin // Event containing the contract specifics and raw log

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
func (it *ContractsChangeAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsChangeAdmin)
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
		it.Event = new(ContractsChangeAdmin)
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
func (it *ContractsChangeAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsChangeAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsChangeAdmin represents a ChangeAdmin event raised by the Contracts contract.
type ContractsChangeAdmin struct {
	Old common.Address
	New common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChangeAdmin is a free log retrieval operation binding the contract event 0xcf9b665e0639e0b81a8db37b60ac7ddf45aeb1b484e11adeb7dff4bf4a3a6258.
//
// Solidity: event ChangeAdmin(address indexed _old, address indexed _new)
func (_Contracts *ContractsFilterer) FilterChangeAdmin(opts *bind.FilterOpts, _old []common.Address, _new []common.Address) (*ContractsChangeAdminIterator, error) {

	var _oldRule []interface{}
	for _, _oldItem := range _old {
		_oldRule = append(_oldRule, _oldItem)
	}
	var _newRule []interface{}
	for _, _newItem := range _new {
		_newRule = append(_newRule, _newItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ChangeAdmin", _oldRule, _newRule)
	if err != nil {
		return nil, err
	}
	return &ContractsChangeAdminIterator{contract: _Contracts.contract, event: "ChangeAdmin", logs: logs, sub: sub}, nil
}

// WatchChangeAdmin is a free log subscription operation binding the contract event 0xcf9b665e0639e0b81a8db37b60ac7ddf45aeb1b484e11adeb7dff4bf4a3a6258.
//
// Solidity: event ChangeAdmin(address indexed _old, address indexed _new)
func (_Contracts *ContractsFilterer) WatchChangeAdmin(opts *bind.WatchOpts, sink chan<- *ContractsChangeAdmin, _old []common.Address, _new []common.Address) (event.Subscription, error) {

	var _oldRule []interface{}
	for _, _oldItem := range _old {
		_oldRule = append(_oldRule, _oldItem)
	}
	var _newRule []interface{}
	for _, _newItem := range _new {
		_newRule = append(_newRule, _newItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ChangeAdmin", _oldRule, _newRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsChangeAdmin)
				if err := _Contracts.contract.UnpackLog(event, "ChangeAdmin", log); err != nil {
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

// ParseChangeAdmin is a log parse operation binding the contract event 0xcf9b665e0639e0b81a8db37b60ac7ddf45aeb1b484e11adeb7dff4bf4a3a6258.
//
// Solidity: event ChangeAdmin(address indexed _old, address indexed _new)
func (_Contracts *ContractsFilterer) ParseChangeAdmin(log types.Log) (*ContractsChangeAdmin, error) {
	event := new(ContractsChangeAdmin)
	if err := _Contracts.contract.UnpackLog(event, "ChangeAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	SwapID      string
	SwapoutID   [32]byte
	Token       common.Address
	Receiver    common.Address
	Amount      *big.Int
	FromChainID *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapIn is a free log retrieval operation binding the contract event 0x164f647883b52834be7a5219336e455a23a358be27519d0442fc0ee5e1b1ce2e.
//
// Solidity: event LogAnySwapIn(string swapID, bytes32 indexed swapoutID, address indexed token, address indexed receiver, uint256 amount, uint256 fromChainID)
func (_Contracts *ContractsFilterer) FilterLogAnySwapIn(opts *bind.FilterOpts, swapoutID [][32]byte, token []common.Address, receiver []common.Address) (*ContractsLogAnySwapInIterator, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogAnySwapIn", swapoutIDRule, tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogAnySwapInIterator{contract: _Contracts.contract, event: "LogAnySwapIn", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapIn is a free log subscription operation binding the contract event 0x164f647883b52834be7a5219336e455a23a358be27519d0442fc0ee5e1b1ce2e.
//
// Solidity: event LogAnySwapIn(string swapID, bytes32 indexed swapoutID, address indexed token, address indexed receiver, uint256 amount, uint256 fromChainID)
func (_Contracts *ContractsFilterer) WatchLogAnySwapIn(opts *bind.WatchOpts, sink chan<- *ContractsLogAnySwapIn, swapoutID [][32]byte, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogAnySwapIn", swapoutIDRule, tokenRule, receiverRule)
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

// ParseLogAnySwapIn is a log parse operation binding the contract event 0x164f647883b52834be7a5219336e455a23a358be27519d0442fc0ee5e1b1ce2e.
//
// Solidity: event LogAnySwapIn(string swapID, bytes32 indexed swapoutID, address indexed token, address indexed receiver, uint256 amount, uint256 fromChainID)
func (_Contracts *ContractsFilterer) ParseLogAnySwapIn(log types.Log) (*ContractsLogAnySwapIn, error) {
	event := new(ContractsLogAnySwapIn)
	if err := _Contracts.contract.UnpackLog(event, "LogAnySwapIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogAnySwapInAndExecIterator is returned from FilterLogAnySwapInAndExec and is used to iterate over the raw logs and unpacked data for LogAnySwapInAndExec events raised by the Contracts contract.
type ContractsLogAnySwapInAndExecIterator struct {
	Event *ContractsLogAnySwapInAndExec // Event containing the contract specifics and raw log

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
func (it *ContractsLogAnySwapInAndExecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogAnySwapInAndExec)
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
		it.Event = new(ContractsLogAnySwapInAndExec)
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
func (it *ContractsLogAnySwapInAndExecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogAnySwapInAndExecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogAnySwapInAndExec represents a LogAnySwapInAndExec event raised by the Contracts contract.
type ContractsLogAnySwapInAndExec struct {
	SwapID      string
	SwapoutID   [32]byte
	Token       common.Address
	Receiver    common.Address
	Amount      *big.Int
	FromChainID *big.Int
	Success     bool
	Result      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapInAndExec is a free log retrieval operation binding the contract event 0x603ea9944a12c4ef108a97399c705891f182d169a361b6aa6455d14aa1cdd258.
//
// Solidity: event LogAnySwapInAndExec(string swapID, bytes32 indexed swapoutID, address indexed token, address indexed receiver, uint256 amount, uint256 fromChainID, bool success, bytes result)
func (_Contracts *ContractsFilterer) FilterLogAnySwapInAndExec(opts *bind.FilterOpts, swapoutID [][32]byte, token []common.Address, receiver []common.Address) (*ContractsLogAnySwapInAndExecIterator, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogAnySwapInAndExec", swapoutIDRule, tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogAnySwapInAndExecIterator{contract: _Contracts.contract, event: "LogAnySwapInAndExec", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapInAndExec is a free log subscription operation binding the contract event 0x603ea9944a12c4ef108a97399c705891f182d169a361b6aa6455d14aa1cdd258.
//
// Solidity: event LogAnySwapInAndExec(string swapID, bytes32 indexed swapoutID, address indexed token, address indexed receiver, uint256 amount, uint256 fromChainID, bool success, bytes result)
func (_Contracts *ContractsFilterer) WatchLogAnySwapInAndExec(opts *bind.WatchOpts, sink chan<- *ContractsLogAnySwapInAndExec, swapoutID [][32]byte, token []common.Address, receiver []common.Address) (event.Subscription, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogAnySwapInAndExec", swapoutIDRule, tokenRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogAnySwapInAndExec)
				if err := _Contracts.contract.UnpackLog(event, "LogAnySwapInAndExec", log); err != nil {
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

// ParseLogAnySwapInAndExec is a log parse operation binding the contract event 0x603ea9944a12c4ef108a97399c705891f182d169a361b6aa6455d14aa1cdd258.
//
// Solidity: event LogAnySwapInAndExec(string swapID, bytes32 indexed swapoutID, address indexed token, address indexed receiver, uint256 amount, uint256 fromChainID, bool success, bytes result)
func (_Contracts *ContractsFilterer) ParseLogAnySwapInAndExec(log types.Log) (*ContractsLogAnySwapInAndExec, error) {
	event := new(ContractsLogAnySwapInAndExec)
	if err := _Contracts.contract.UnpackLog(event, "LogAnySwapInAndExec", log); err != nil {
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
	SwapoutID [32]byte
	Token     common.Address
	From      common.Address
	Receiver  string
	Amount    *big.Int
	ToChainID *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapOut is a free log retrieval operation binding the contract event 0x0d969ae475ff6fcaf0dcfa760d4d8607244e8d95e9bf426f8d5d69f9a3e525af.
//
// Solidity: event LogAnySwapOut(bytes32 indexed swapoutID, address indexed token, address indexed from, string receiver, uint256 amount, uint256 toChainID)
func (_Contracts *ContractsFilterer) FilterLogAnySwapOut(opts *bind.FilterOpts, swapoutID [][32]byte, token []common.Address, from []common.Address) (*ContractsLogAnySwapOutIterator, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogAnySwapOut", swapoutIDRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogAnySwapOutIterator{contract: _Contracts.contract, event: "LogAnySwapOut", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapOut is a free log subscription operation binding the contract event 0x0d969ae475ff6fcaf0dcfa760d4d8607244e8d95e9bf426f8d5d69f9a3e525af.
//
// Solidity: event LogAnySwapOut(bytes32 indexed swapoutID, address indexed token, address indexed from, string receiver, uint256 amount, uint256 toChainID)
func (_Contracts *ContractsFilterer) WatchLogAnySwapOut(opts *bind.WatchOpts, sink chan<- *ContractsLogAnySwapOut, swapoutID [][32]byte, token []common.Address, from []common.Address) (event.Subscription, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogAnySwapOut", swapoutIDRule, tokenRule, fromRule)
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

// ParseLogAnySwapOut is a log parse operation binding the contract event 0x0d969ae475ff6fcaf0dcfa760d4d8607244e8d95e9bf426f8d5d69f9a3e525af.
//
// Solidity: event LogAnySwapOut(bytes32 indexed swapoutID, address indexed token, address indexed from, string receiver, uint256 amount, uint256 toChainID)
func (_Contracts *ContractsFilterer) ParseLogAnySwapOut(log types.Log) (*ContractsLogAnySwapOut, error) {
	event := new(ContractsLogAnySwapOut)
	if err := _Contracts.contract.UnpackLog(event, "LogAnySwapOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogAnySwapOutAndCallIterator is returned from FilterLogAnySwapOutAndCall and is used to iterate over the raw logs and unpacked data for LogAnySwapOutAndCall events raised by the Contracts contract.
type ContractsLogAnySwapOutAndCallIterator struct {
	Event *ContractsLogAnySwapOutAndCall // Event containing the contract specifics and raw log

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
func (it *ContractsLogAnySwapOutAndCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogAnySwapOutAndCall)
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
		it.Event = new(ContractsLogAnySwapOutAndCall)
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
func (it *ContractsLogAnySwapOutAndCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogAnySwapOutAndCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogAnySwapOutAndCall represents a LogAnySwapOutAndCall event raised by the Contracts contract.
type ContractsLogAnySwapOutAndCall struct {
	SwapoutID    [32]byte
	Token        common.Address
	From         common.Address
	Receiver     string
	Amount       *big.Int
	ToChainID    *big.Int
	AnycallProxy string
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogAnySwapOutAndCall is a free log retrieval operation binding the contract event 0x968608314ec29f6fd1a9f6ef9e96247a4da1a683917569706e2d2b60ca7c0a6d.
//
// Solidity: event LogAnySwapOutAndCall(bytes32 indexed swapoutID, address indexed token, address indexed from, string receiver, uint256 amount, uint256 toChainID, string anycallProxy, bytes data)
func (_Contracts *ContractsFilterer) FilterLogAnySwapOutAndCall(opts *bind.FilterOpts, swapoutID [][32]byte, token []common.Address, from []common.Address) (*ContractsLogAnySwapOutAndCallIterator, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogAnySwapOutAndCall", swapoutIDRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogAnySwapOutAndCallIterator{contract: _Contracts.contract, event: "LogAnySwapOutAndCall", logs: logs, sub: sub}, nil
}

// WatchLogAnySwapOutAndCall is a free log subscription operation binding the contract event 0x968608314ec29f6fd1a9f6ef9e96247a4da1a683917569706e2d2b60ca7c0a6d.
//
// Solidity: event LogAnySwapOutAndCall(bytes32 indexed swapoutID, address indexed token, address indexed from, string receiver, uint256 amount, uint256 toChainID, string anycallProxy, bytes data)
func (_Contracts *ContractsFilterer) WatchLogAnySwapOutAndCall(opts *bind.WatchOpts, sink chan<- *ContractsLogAnySwapOutAndCall, swapoutID [][32]byte, token []common.Address, from []common.Address) (event.Subscription, error) {

	var swapoutIDRule []interface{}
	for _, swapoutIDItem := range swapoutID {
		swapoutIDRule = append(swapoutIDRule, swapoutIDItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogAnySwapOutAndCall", swapoutIDRule, tokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogAnySwapOutAndCall)
				if err := _Contracts.contract.UnpackLog(event, "LogAnySwapOutAndCall", log); err != nil {
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

// ParseLogAnySwapOutAndCall is a log parse operation binding the contract event 0x968608314ec29f6fd1a9f6ef9e96247a4da1a683917569706e2d2b60ca7c0a6d.
//
// Solidity: event LogAnySwapOutAndCall(bytes32 indexed swapoutID, address indexed token, address indexed from, string receiver, uint256 amount, uint256 toChainID, string anycallProxy, bytes data)
func (_Contracts *ContractsFilterer) ParseLogAnySwapOutAndCall(log types.Log) (*ContractsLogAnySwapOutAndCall, error) {
	event := new(ContractsLogAnySwapOutAndCall)
	if err := _Contracts.contract.UnpackLog(event, "LogAnySwapOutAndCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogApplyMPCIterator is returned from FilterLogApplyMPC and is used to iterate over the raw logs and unpacked data for LogApplyMPC events raised by the Contracts contract.
type ContractsLogApplyMPCIterator struct {
	Event *ContractsLogApplyMPC // Event containing the contract specifics and raw log

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
func (it *ContractsLogApplyMPCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogApplyMPC)
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
		it.Event = new(ContractsLogApplyMPC)
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
func (it *ContractsLogApplyMPCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogApplyMPCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogApplyMPC represents a LogApplyMPC event raised by the Contracts contract.
type ContractsLogApplyMPC struct {
	OldMPC    common.Address
	NewMPC    common.Address
	ApplyTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogApplyMPC is a free log retrieval operation binding the contract event 0x8d32c9dd498e08090b44a0f77fe9ec0278851f9dffc4b430428411243e7df076.
//
// Solidity: event LogApplyMPC(address indexed oldMPC, address indexed newMPC, uint256 applyTime)
func (_Contracts *ContractsFilterer) FilterLogApplyMPC(opts *bind.FilterOpts, oldMPC []common.Address, newMPC []common.Address) (*ContractsLogApplyMPCIterator, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogApplyMPC", oldMPCRule, newMPCRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogApplyMPCIterator{contract: _Contracts.contract, event: "LogApplyMPC", logs: logs, sub: sub}, nil
}

// WatchLogApplyMPC is a free log subscription operation binding the contract event 0x8d32c9dd498e08090b44a0f77fe9ec0278851f9dffc4b430428411243e7df076.
//
// Solidity: event LogApplyMPC(address indexed oldMPC, address indexed newMPC, uint256 applyTime)
func (_Contracts *ContractsFilterer) WatchLogApplyMPC(opts *bind.WatchOpts, sink chan<- *ContractsLogApplyMPC, oldMPC []common.Address, newMPC []common.Address) (event.Subscription, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogApplyMPC", oldMPCRule, newMPCRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogApplyMPC)
				if err := _Contracts.contract.UnpackLog(event, "LogApplyMPC", log); err != nil {
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

// ParseLogApplyMPC is a log parse operation binding the contract event 0x8d32c9dd498e08090b44a0f77fe9ec0278851f9dffc4b430428411243e7df076.
//
// Solidity: event LogApplyMPC(address indexed oldMPC, address indexed newMPC, uint256 applyTime)
func (_Contracts *ContractsFilterer) ParseLogApplyMPC(log types.Log) (*ContractsLogApplyMPC, error) {
	event := new(ContractsLogApplyMPC)
	if err := _Contracts.contract.UnpackLog(event, "LogApplyMPC", log); err != nil {
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
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogChangeMPC is a free log retrieval operation binding the contract event 0x581f388e3dd32e1bbf62a290f509c8245f9d0b71ef82614fb2b967ad0a10d5b9.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 effectiveTime)
func (_Contracts *ContractsFilterer) FilterLogChangeMPC(opts *bind.FilterOpts, oldMPC []common.Address, newMPC []common.Address) (*ContractsLogChangeMPCIterator, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogChangeMPC", oldMPCRule, newMPCRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogChangeMPCIterator{contract: _Contracts.contract, event: "LogChangeMPC", logs: logs, sub: sub}, nil
}

// WatchLogChangeMPC is a free log subscription operation binding the contract event 0x581f388e3dd32e1bbf62a290f509c8245f9d0b71ef82614fb2b967ad0a10d5b9.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 effectiveTime)
func (_Contracts *ContractsFilterer) WatchLogChangeMPC(opts *bind.WatchOpts, sink chan<- *ContractsLogChangeMPC, oldMPC []common.Address, newMPC []common.Address) (event.Subscription, error) {

	var oldMPCRule []interface{}
	for _, oldMPCItem := range oldMPC {
		oldMPCRule = append(oldMPCRule, oldMPCItem)
	}
	var newMPCRule []interface{}
	for _, newMPCItem := range newMPC {
		newMPCRule = append(newMPCRule, newMPCItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogChangeMPC", oldMPCRule, newMPCRule)
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

// ParseLogChangeMPC is a log parse operation binding the contract event 0x581f388e3dd32e1bbf62a290f509c8245f9d0b71ef82614fb2b967ad0a10d5b9.
//
// Solidity: event LogChangeMPC(address indexed oldMPC, address indexed newMPC, uint256 effectiveTime)
func (_Contracts *ContractsFilterer) ParseLogChangeMPC(log types.Log) (*ContractsLogChangeMPC, error) {
	event := new(ContractsLogChangeMPC)
	if err := _Contracts.contract.UnpackLog(event, "LogChangeMPC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogRetryExecRecordIterator is returned from FilterLogRetryExecRecord and is used to iterate over the raw logs and unpacked data for LogRetryExecRecord events raised by the Contracts contract.
type ContractsLogRetryExecRecordIterator struct {
	Event *ContractsLogRetryExecRecord // Event containing the contract specifics and raw log

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
func (it *ContractsLogRetryExecRecordIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogRetryExecRecord)
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
		it.Event = new(ContractsLogRetryExecRecord)
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
func (it *ContractsLogRetryExecRecordIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogRetryExecRecordIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogRetryExecRecord represents a LogRetryExecRecord event raised by the Contracts contract.
type ContractsLogRetryExecRecord struct {
	SwapID       string
	SwapoutID    [32]byte
	Token        common.Address
	Receiver     common.Address
	Amount       *big.Int
	FromChainID  *big.Int
	AnycallProxy common.Address
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogRetryExecRecord is a free log retrieval operation binding the contract event 0x2d044017b61f24f5423ce5e0c62f9ead27cb38f1615069e703ba521d0b04696b.
//
// Solidity: event LogRetryExecRecord(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, address anycallProxy, bytes data)
func (_Contracts *ContractsFilterer) FilterLogRetryExecRecord(opts *bind.FilterOpts) (*ContractsLogRetryExecRecordIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogRetryExecRecord")
	if err != nil {
		return nil, err
	}
	return &ContractsLogRetryExecRecordIterator{contract: _Contracts.contract, event: "LogRetryExecRecord", logs: logs, sub: sub}, nil
}

// WatchLogRetryExecRecord is a free log subscription operation binding the contract event 0x2d044017b61f24f5423ce5e0c62f9ead27cb38f1615069e703ba521d0b04696b.
//
// Solidity: event LogRetryExecRecord(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, address anycallProxy, bytes data)
func (_Contracts *ContractsFilterer) WatchLogRetryExecRecord(opts *bind.WatchOpts, sink chan<- *ContractsLogRetryExecRecord) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogRetryExecRecord")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogRetryExecRecord)
				if err := _Contracts.contract.UnpackLog(event, "LogRetryExecRecord", log); err != nil {
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

// ParseLogRetryExecRecord is a log parse operation binding the contract event 0x2d044017b61f24f5423ce5e0c62f9ead27cb38f1615069e703ba521d0b04696b.
//
// Solidity: event LogRetryExecRecord(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, address anycallProxy, bytes data)
func (_Contracts *ContractsFilterer) ParseLogRetryExecRecord(log types.Log) (*ContractsLogRetryExecRecord, error) {
	event := new(ContractsLogRetryExecRecord)
	if err := _Contracts.contract.UnpackLog(event, "LogRetryExecRecord", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogRetrySwapInAndExecIterator is returned from FilterLogRetrySwapInAndExec and is used to iterate over the raw logs and unpacked data for LogRetrySwapInAndExec events raised by the Contracts contract.
type ContractsLogRetrySwapInAndExecIterator struct {
	Event *ContractsLogRetrySwapInAndExec // Event containing the contract specifics and raw log

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
func (it *ContractsLogRetrySwapInAndExecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogRetrySwapInAndExec)
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
		it.Event = new(ContractsLogRetrySwapInAndExec)
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
func (it *ContractsLogRetrySwapInAndExecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogRetrySwapInAndExecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogRetrySwapInAndExec represents a LogRetrySwapInAndExec event raised by the Contracts contract.
type ContractsLogRetrySwapInAndExec struct {
	SwapID      string
	SwapoutID   [32]byte
	Token       common.Address
	Receiver    common.Address
	Amount      *big.Int
	FromChainID *big.Int
	DontExec    bool
	Success     bool
	Result      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogRetrySwapInAndExec is a free log retrieval operation binding the contract event 0x4024f72e00ae47f03ed1dd3ab595d04dabdc9d1f95f8c039bca61946d9da0eb3.
//
// Solidity: event LogRetrySwapInAndExec(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, bool dontExec, bool success, bytes result)
func (_Contracts *ContractsFilterer) FilterLogRetrySwapInAndExec(opts *bind.FilterOpts) (*ContractsLogRetrySwapInAndExecIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogRetrySwapInAndExec")
	if err != nil {
		return nil, err
	}
	return &ContractsLogRetrySwapInAndExecIterator{contract: _Contracts.contract, event: "LogRetrySwapInAndExec", logs: logs, sub: sub}, nil
}

// WatchLogRetrySwapInAndExec is a free log subscription operation binding the contract event 0x4024f72e00ae47f03ed1dd3ab595d04dabdc9d1f95f8c039bca61946d9da0eb3.
//
// Solidity: event LogRetrySwapInAndExec(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, bool dontExec, bool success, bytes result)
func (_Contracts *ContractsFilterer) WatchLogRetrySwapInAndExec(opts *bind.WatchOpts, sink chan<- *ContractsLogRetrySwapInAndExec) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogRetrySwapInAndExec")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogRetrySwapInAndExec)
				if err := _Contracts.contract.UnpackLog(event, "LogRetrySwapInAndExec", log); err != nil {
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

// ParseLogRetrySwapInAndExec is a log parse operation binding the contract event 0x4024f72e00ae47f03ed1dd3ab595d04dabdc9d1f95f8c039bca61946d9da0eb3.
//
// Solidity: event LogRetrySwapInAndExec(string swapID, bytes32 swapoutID, address token, address receiver, uint256 amount, uint256 fromChainID, bool dontExec, bool success, bytes result)
func (_Contracts *ContractsFilterer) ParseLogRetrySwapInAndExec(log types.Log) (*ContractsLogRetrySwapInAndExec, error) {
	event := new(ContractsLogRetrySwapInAndExec)
	if err := _Contracts.contract.UnpackLog(event, "LogRetrySwapInAndExec", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Contracts contract.
type ContractsPausedIterator struct {
	Event *ContractsPaused // Event containing the contract specifics and raw log

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
func (it *ContractsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPaused)
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
		it.Event = new(ContractsPaused)
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
func (it *ContractsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPaused represents a Paused event raised by the Contracts contract.
type ContractsPaused struct {
	Role [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 role)
func (_Contracts *ContractsFilterer) FilterPaused(opts *bind.FilterOpts) (*ContractsPausedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ContractsPausedIterator{contract: _Contracts.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 role)
func (_Contracts *ContractsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractsPaused) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPaused)
				if err := _Contracts.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 role)
func (_Contracts *ContractsFilterer) ParsePaused(log types.Log) (*ContractsPaused, error) {
	event := new(ContractsPaused)
	if err := _Contracts.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Contracts contract.
type ContractsUnpausedIterator struct {
	Event *ContractsUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsUnpaused)
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
		it.Event = new(ContractsUnpaused)
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
func (it *ContractsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsUnpaused represents a Unpaused event raised by the Contracts contract.
type ContractsUnpaused struct {
	Role [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 role)
func (_Contracts *ContractsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ContractsUnpausedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ContractsUnpausedIterator{contract: _Contracts.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 role)
func (_Contracts *ContractsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractsUnpaused) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsUnpaused)
				if err := _Contracts.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 role)
func (_Contracts *ContractsFilterer) ParseUnpaused(log types.Log) (*ContractsUnpaused, error) {
	event := new(ContractsUnpaused)
	if err := _Contracts.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
