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

// IBridgeAggregatorbridgeInput is an auto generated low-level Go binding around an user-defined struct.
type IBridgeAggregatorbridgeInput struct {
	Bridge     common.Address
	TokenIn    common.Address
	Sender     common.Address
	BridgeFee  *big.Int
	AfterSwap  bool
	AmountIn   *big.Int
	FeeAmount  *big.Int
	BridgeData []byte
}

// IExchangeAggregatorswapInput is an auto generated low-level Go binding around an user-defined struct.
type IExchangeAggregatorswapInput struct {
	TokenIn      common.Address
	TokenOut     common.Address
	TotalAmount  *big.Int
	FeeAmount    *big.Int
	AmountIn     *big.Int
	FromContract bool
	Swapper      common.Address
	SwapperData  []byte
	Sender       common.Address
	Receiver     common.Address
	Native       bool
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BalanceETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"afterSwap\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"bridgeData\",\"type\":\"bytes\"}],\"internalType\":\"structIBridgeAggregator.bridgeInput\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"Bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeReciever\",\"type\":\"address\"}],\"name\":\"ChangeFeeReciever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"fromContract\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapperData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"native\",\"type\":\"bool\"}],\"internalType\":\"structIExchangeAggregator.swapInput\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"Swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"estimateAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReciever\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

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

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) view returns(uint256)
func (_Contracts *ContractsCaller) Balance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "Balance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) view returns(uint256)
func (_Contracts *ContractsSession) Balance(token common.Address) (*big.Int, error) {
	return _Contracts.Contract.Balance(&_Contracts.CallOpts, token)
}

// Balance is a free data retrieval call binding the contract method 0x239fcf0f.
//
// Solidity: function Balance(address token) view returns(uint256)
func (_Contracts *ContractsCallerSession) Balance(token common.Address) (*big.Int, error) {
	return _Contracts.Contract.Balance(&_Contracts.CallOpts, token)
}

// BalanceETH is a free data retrieval call binding the contract method 0x231da50d.
//
// Solidity: function BalanceETH() view returns(uint256)
func (_Contracts *ContractsCaller) BalanceETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "BalanceETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceETH is a free data retrieval call binding the contract method 0x231da50d.
//
// Solidity: function BalanceETH() view returns(uint256)
func (_Contracts *ContractsSession) BalanceETH() (*big.Int, error) {
	return _Contracts.Contract.BalanceETH(&_Contracts.CallOpts)
}

// BalanceETH is a free data retrieval call binding the contract method 0x231da50d.
//
// Solidity: function BalanceETH() view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceETH() (*big.Int, error) {
	return _Contracts.Contract.BalanceETH(&_Contracts.CallOpts)
}

// EstimateAmountOut is a free data retrieval call binding the contract method 0xa299ed7d.
//
// Solidity: function estimateAmountOut(address priceProvider, address provider, address tA, address tB, uint256 amountIn, uint8 version) view returns(uint256 amountOut, uint24 fee)
func (_Contracts *ContractsCaller) EstimateAmountOut(opts *bind.CallOpts, priceProvider common.Address, provider common.Address, tA common.Address, tB common.Address, amountIn *big.Int, version uint8) (struct {
	AmountOut *big.Int
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "estimateAmountOut", priceProvider, provider, tA, tB, amountIn, version)

	outstruct := new(struct {
		AmountOut *big.Int
		Fee       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EstimateAmountOut is a free data retrieval call binding the contract method 0xa299ed7d.
//
// Solidity: function estimateAmountOut(address priceProvider, address provider, address tA, address tB, uint256 amountIn, uint8 version) view returns(uint256 amountOut, uint24 fee)
func (_Contracts *ContractsSession) EstimateAmountOut(priceProvider common.Address, provider common.Address, tA common.Address, tB common.Address, amountIn *big.Int, version uint8) (struct {
	AmountOut *big.Int
	Fee       *big.Int
}, error) {
	return _Contracts.Contract.EstimateAmountOut(&_Contracts.CallOpts, priceProvider, provider, tA, tB, amountIn, version)
}

// EstimateAmountOut is a free data retrieval call binding the contract method 0xa299ed7d.
//
// Solidity: function estimateAmountOut(address priceProvider, address provider, address tA, address tB, uint256 amountIn, uint8 version) view returns(uint256 amountOut, uint24 fee)
func (_Contracts *ContractsCallerSession) EstimateAmountOut(priceProvider common.Address, provider common.Address, tA common.Address, tB common.Address, amountIn *big.Int, version uint8) (struct {
	AmountOut *big.Int
	Fee       *big.Int
}, error) {
	return _Contracts.Contract.EstimateAmountOut(&_Contracts.CallOpts, priceProvider, provider, tA, tB, amountIn, version)
}

// FeeReciever is a free data retrieval call binding the contract method 0xf61db740.
//
// Solidity: function feeReciever() view returns(address)
func (_Contracts *ContractsCaller) FeeReciever(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "feeReciever")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReciever is a free data retrieval call binding the contract method 0xf61db740.
//
// Solidity: function feeReciever() view returns(address)
func (_Contracts *ContractsSession) FeeReciever() (common.Address, error) {
	return _Contracts.Contract.FeeReciever(&_Contracts.CallOpts)
}

// FeeReciever is a free data retrieval call binding the contract method 0xf61db740.
//
// Solidity: function feeReciever() view returns(address)
func (_Contracts *ContractsCallerSession) FeeReciever() (common.Address, error) {
	return _Contracts.Contract.FeeReciever(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// SwapAmountOut is a free data retrieval call binding the contract method 0x22cedf7d.
//
// Solidity: function swapAmountOut() view returns(uint256)
func (_Contracts *ContractsCaller) SwapAmountOut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "swapAmountOut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapAmountOut is a free data retrieval call binding the contract method 0x22cedf7d.
//
// Solidity: function swapAmountOut() view returns(uint256)
func (_Contracts *ContractsSession) SwapAmountOut() (*big.Int, error) {
	return _Contracts.Contract.SwapAmountOut(&_Contracts.CallOpts)
}

// SwapAmountOut is a free data retrieval call binding the contract method 0x22cedf7d.
//
// Solidity: function swapAmountOut() view returns(uint256)
func (_Contracts *ContractsCallerSession) SwapAmountOut() (*big.Int, error) {
	return _Contracts.Contract.SwapAmountOut(&_Contracts.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0x5cedfda0.
//
// Solidity: function Bridge((address,address,address,uint256,bool,uint256,uint256,bytes) data, bytes sig) payable returns()
func (_Contracts *ContractsTransactor) Bridge(opts *bind.TransactOpts, data IBridgeAggregatorbridgeInput, sig []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "Bridge", data, sig)
}

// Bridge is a paid mutator transaction binding the contract method 0x5cedfda0.
//
// Solidity: function Bridge((address,address,address,uint256,bool,uint256,uint256,bytes) data, bytes sig) payable returns()
func (_Contracts *ContractsSession) Bridge(data IBridgeAggregatorbridgeInput, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Bridge(&_Contracts.TransactOpts, data, sig)
}

// Bridge is a paid mutator transaction binding the contract method 0x5cedfda0.
//
// Solidity: function Bridge((address,address,address,uint256,bool,uint256,uint256,bytes) data, bytes sig) payable returns()
func (_Contracts *ContractsTransactorSession) Bridge(data IBridgeAggregatorbridgeInput, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Bridge(&_Contracts.TransactOpts, data, sig)
}

// ChangeFeeReciever is a paid mutator transaction binding the contract method 0xddd0fcd5.
//
// Solidity: function ChangeFeeReciever(address _feeReciever) returns()
func (_Contracts *ContractsTransactor) ChangeFeeReciever(opts *bind.TransactOpts, _feeReciever common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "ChangeFeeReciever", _feeReciever)
}

// ChangeFeeReciever is a paid mutator transaction binding the contract method 0xddd0fcd5.
//
// Solidity: function ChangeFeeReciever(address _feeReciever) returns()
func (_Contracts *ContractsSession) ChangeFeeReciever(_feeReciever common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeFeeReciever(&_Contracts.TransactOpts, _feeReciever)
}

// ChangeFeeReciever is a paid mutator transaction binding the contract method 0xddd0fcd5.
//
// Solidity: function ChangeFeeReciever(address _feeReciever) returns()
func (_Contracts *ContractsTransactorSession) ChangeFeeReciever(_feeReciever common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeFeeReciever(&_Contracts.TransactOpts, _feeReciever)
}

// Swap is a paid mutator transaction binding the contract method 0x3f464ff5.
//
// Solidity: function Swap((address,address,uint256,uint256,uint256,bool,address,bytes,address,address,bool) data, bytes sig) payable returns()
func (_Contracts *ContractsTransactor) Swap(opts *bind.TransactOpts, data IExchangeAggregatorswapInput, sig []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "Swap", data, sig)
}

// Swap is a paid mutator transaction binding the contract method 0x3f464ff5.
//
// Solidity: function Swap((address,address,uint256,uint256,uint256,bool,address,bytes,address,address,bool) data, bytes sig) payable returns()
func (_Contracts *ContractsSession) Swap(data IExchangeAggregatorswapInput, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Swap(&_Contracts.TransactOpts, data, sig)
}

// Swap is a paid mutator transaction binding the contract method 0x3f464ff5.
//
// Solidity: function Swap((address,address,uint256,uint256,uint256,bool,address,bytes,address,address,bool) data, bytes sig) payable returns()
func (_Contracts *ContractsTransactorSession) Swap(data IExchangeAggregatorswapInput, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Swap(&_Contracts.TransactOpts, data, sig)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b1bfa7f.
//
// Solidity: function Withdraw(address token, address to, uint256 amount) returns()
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "Withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b1bfa7f.
//
// Solidity: function Withdraw(address token, address to, uint256 amount) returns()
func (_Contracts *ContractsSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9b1bfa7f.
//
// Solidity: function Withdraw(address token, address to, uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, token, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x566e45b1.
//
// Solidity: function WithdrawETH(address to, uint256 amount) payable returns()
func (_Contracts *ContractsTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "WithdrawETH", to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x566e45b1.
//
// Solidity: function WithdrawETH(address to, uint256 amount) payable returns()
func (_Contracts *ContractsSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawETH(&_Contracts.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x566e45b1.
//
// Solidity: function WithdrawETH(address to, uint256 amount) payable returns()
func (_Contracts *ContractsTransactorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawETH(&_Contracts.TransactOpts, to, amount)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Contracts *ContractsTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Contracts *ContractsSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Multicall(&_Contracts.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Contracts *ContractsTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Multicall(&_Contracts.TransactOpts, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contracts *ContractsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.RenounceOwnership(&_Contracts.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contracts *ContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.TransferOwnership(&_Contracts.TransactOpts, newOwner)
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

// ContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contracts contract.
type ContractsOwnershipTransferredIterator struct {
	Event *ContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnershipTransferred)
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
		it.Event = new(ContractsOwnershipTransferred)
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
func (it *ContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnershipTransferred represents a OwnershipTransferred event raised by the Contracts contract.
type ContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractsOwnershipTransferredIterator{contract: _Contracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnershipTransferred)
				if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contracts *ContractsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractsOwnershipTransferred, error) {
	event := new(ContractsOwnershipTransferred)
	if err := _Contracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
