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

// IExchangeAggregatorswapData is an auto generated low-level Go binding around an user-defined struct.
type IExchangeAggregatorswapData struct {
	Input       common.Address
	TotalAmount *big.Int
	FeeAmount   *big.Int
	Swapper     common.Address
	Data        []byte
	Sender      common.Address
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_PriceProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PriceProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_PriceProvider\",\"type\":\"address\"}],\"name\":\"changePriceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"estimateAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structIExchangeAggregator.swapData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structIExchangeAggregator.swapData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"swapNativeIn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200293738038062002937833981810160405281019062000037919062000217565b620000576200004b620000e160201b60201c565b620000e960201b60201c565b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050506200025e565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001df82620001b2565b9050919050565b620001f181620001d2565b8114620001fd57600080fd5b50565b6000815190506200021181620001e6565b92915050565b60008060408385031215620002315762000230620001ad565b5b6000620002418582860162000200565b9250506020620002548582860162000200565b9150509250929050565b6126c9806200026e6000396000f3fe6080604052600436106100c25760003560e01c80636110358d1161007f578063ad5c464811610059578063ad5c464814610246578063ae4f5be814610271578063ecbdbb321461029a578063f2fde38b146102c5576100c2565b80636110358d146101c6578063715018a6146102045780638da5cb5b1461021b576100c2565b806301e33667146100c757806304599012146100f05780632906799b1461012d5780634782f779146101565780634e1a67f51461017f5780635d86acf1146101aa575b600080fd5b3480156100d357600080fd5b506100ee60048036038101906100e99190611531565b6102ee565b005b3480156100fc57600080fd5b5061011760048036038101906101129190611584565b610306565b60405161012491906115c0565b60405180910390f35b34801561013957600080fd5b50610154600480360381019061014f9190611664565b610389565b005b34801561016257600080fd5b5061017d600480360381019061017891906116e0565b610557565b005b34801561018b57600080fd5b5061019461056d565b6040516101a1919061172f565b60405180910390f35b6101c460048036038101906101bf9190611664565b610593565b005b3480156101d257600080fd5b506101ed60048036038101906101e89190611783565b610758565b6040516101fb92919061181c565b60405180910390f35b34801561021057600080fd5b5061021961080c565b005b34801561022757600080fd5b50610230610820565b60405161023d919061172f565b60405180910390f35b34801561025257600080fd5b5061025b610849565b604051610268919061172f565b60405180910390f35b34801561027d57600080fd5b5061029860048036038101906102939190611584565b61086f565b005b3480156102a657600080fd5b506102af6108bb565b6040516102bc91906115c0565b60405180910390f35b3480156102d157600080fd5b506102ec60048036038101906102e79190611584565b6108c3565b005b6102f6610946565b6103018383836109c4565b505050565b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610341919061172f565b602060405180830381865afa15801561035e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610382919061185a565b9050919050565b3373ffffffffffffffffffffffffffffffffffffffff168360a00160208101906103b39190611584565b73ffffffffffffffffffffffffffffffffffffffff1614610409576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610400906118e4565b60405180910390fd5b61047e610414610820565b846040516020016104259190611adb565b60405160208183030381529060405284848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610afa565b6104a08360000160208101906104949190611584565b33308660200135610c66565b6104e28360000160208101906104b69190611584565b8460600160208101906104c99190611584565b856040013586602001356104dd9190611b2c565b610d9f565b6105528360600160208101906104f89190611584565b600085806080019061050a9190611b6f565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610ed5565b505050565b61055f610946565b6105698282610fec565b5050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b3373ffffffffffffffffffffffffffffffffffffffff168360a00160208101906105bd9190611584565b73ffffffffffffffffffffffffffffffffffffffff1614610613576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060a906118e4565b60405180910390fd5b61068861061e610820565b8460405160200161062f9190611adb565b60405160208183030381529060405284848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610afa565b82602001353410156106cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c690611c1e565b60405180910390fd5b60008360400135346106e19190611b2c565b90506107528460600160208101906106f99190611584565b8286806080019061070a9190611b6f565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610ed5565b50505050565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636110358d88888888886040518663ffffffff1660e01b81526004016107be959493929190611c4d565b6040805180830381865afa1580156107da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fe9190611ccc565b915091509550959350505050565b610814610946565b61081e60006110ec565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610877610946565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600047905090565b6108cb610946565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361093a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093190611d7e565b60405180910390fd5b610943816110ec565b50565b61094e6111b0565b73ffffffffffffffffffffffffffffffffffffffff1661096c610820565b73ffffffffffffffffffffffffffffffffffffffff16146109c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b990611dea565b60405180910390fd5b565b6000808473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040516024016109f6929190611e0a565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610a449190611ea4565b6000604051808303816000865af19150503d8060008114610a81576040519150601f19603f3d011682016040523d82523d6000602084013e610a86565b606091505b5091509150818015610ab45750600081511480610ab3575080806020019051810190610ab29190611ef3565b5b5b610af3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aea90611f92565b60405180910390fd5b5050505050565b601b60f81b81604081518110610b1357610b12611fb2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508273ffffffffffffffffffffffffffffffffffffffff16610b7482610b66856111b8565b6111c990919063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff160315610c6157601c60f81b81604081518110610ba957610ba8611fb2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508273ffffffffffffffffffffffffffffffffffffffff16610c0a82610bfc856111b8565b6111c990919063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1614610c60576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c579061202d565b60405180910390fd5b5b505050565b6000808573ffffffffffffffffffffffffffffffffffffffff166323b872dd868686604051602401610c9a9392919061204d565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610ce89190611ea4565b6000604051808303816000865af19150503d8060008114610d25576040519150601f19603f3d011682016040523d82523d6000602084013e610d2a565b606091505b5091509150818015610d585750600081511480610d57575080806020019051810190610d569190611ef3565b5b5b610d97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8e906120f6565b60405180910390fd5b505050505050565b6000808473ffffffffffffffffffffffffffffffffffffffff1663095ea7b38585604051602401610dd1929190611e0a565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610e1f9190611ea4565b6000604051808303816000865af19150503d8060008114610e5c576040519150601f19603f3d011682016040523d82523d6000602084013e610e61565b606091505b5091509150818015610e8f5750600081511480610e8e575080806020019051810190610e8d9190611ef3565b5b5b610ece576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ec590612188565b60405180910390fd5b5050505050565b6000808473ffffffffffffffffffffffffffffffffffffffff168484604051610efe9190611ea4565b60006040518083038185875af1925050503d8060008114610f3b576040519150601f19603f3d011682016040523d82523d6000602084013e610f40565b606091505b509150915081610fe557604481511015610f8f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f869061221a565b60405180910390fd5b60048101905080806020019051810190610fa9919061235b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fdc91906123e8565b60405180910390fd5b5050505050565b60008273ffffffffffffffffffffffffffffffffffffffff1682600067ffffffffffffffff8111156110215761102061223f565b5b6040519080825280601f01601f1916602001820160405280156110535781602001600182028036833780820191505090505b506040516110619190611ea4565b60006040518083038185875af1925050503d806000811461109e576040519150601f19603f3d011682016040523d82523d6000602084013e6110a3565b606091505b50509050806110e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110de9061247c565b60405180910390fd5b505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600081805190602001209050919050565b60008060006111d885856111f0565b915091506111e581611241565b819250505092915050565b60008060418351036112315760008060006020860151925060408601519150606086015160001a9050611225878285856113a7565b9450945050505061123a565b60006002915091505b9250929050565b600060048111156112555761125461249c565b5b8160048111156112685761126761249c565b5b03156113a457600160048111156112825761128161249c565b5b8160048111156112955761129461249c565b5b036112d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112cc90612517565b60405180910390fd5b600260048111156112e9576112e861249c565b5b8160048111156112fc576112fb61249c565b5b0361133c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133390612583565b60405180910390fd5b600360048111156113505761134f61249c565b5b8160048111156113635761136261249c565b5b036113a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161139a90612615565b60405180910390fd5b5b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c11156113e2576000600391509150611480565b600060018787878760405160008152602001604052604051611407949392919061264e565b6020604051602081039080840390855afa158015611429573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361147757600060019250925050611480565b80600092509250505b94509492505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006114c88261149d565b9050919050565b6114d8816114bd565b81146114e357600080fd5b50565b6000813590506114f5816114cf565b92915050565b6000819050919050565b61150e816114fb565b811461151957600080fd5b50565b60008135905061152b81611505565b92915050565b60008060006060848603121561154a57611549611493565b5b6000611558868287016114e6565b9350506020611569868287016114e6565b925050604061157a8682870161151c565b9150509250925092565b60006020828403121561159a57611599611493565b5b60006115a8848285016114e6565b91505092915050565b6115ba816114fb565b82525050565b60006020820190506115d560008301846115b1565b92915050565b600080fd5b600060c082840312156115f6576115f56115db565b5b81905092915050565b600080fd5b600080fd5b600080fd5b60008083601f840112611624576116236115ff565b5b8235905067ffffffffffffffff81111561164157611640611604565b5b60208301915083600182028301111561165d5761165c611609565b5b9250929050565b60008060006040848603121561167d5761167c611493565b5b600084013567ffffffffffffffff81111561169b5761169a611498565b5b6116a7868287016115e0565b935050602084013567ffffffffffffffff8111156116c8576116c7611498565b5b6116d48682870161160e565b92509250509250925092565b600080604083850312156116f7576116f6611493565b5b6000611705858286016114e6565b92505060206117168582860161151c565b9150509250929050565b611729816114bd565b82525050565b60006020820190506117446000830184611720565b92915050565b600060ff82169050919050565b6117608161174a565b811461176b57600080fd5b50565b60008135905061177d81611757565b92915050565b600080600080600060a0868803121561179f5761179e611493565b5b60006117ad888289016114e6565b95505060206117be888289016114e6565b94505060406117cf888289016114e6565b93505060606117e08882890161151c565b92505060806117f18882890161176e565b9150509295509295909350565b600062ffffff82169050919050565b611816816117fe565b82525050565b600060408201905061183160008301856115b1565b61183e602083018461180d565b9392505050565b60008151905061185481611505565b92915050565b6000602082840312156118705761186f611493565b5b600061187e84828501611845565b91505092915050565b600082825260208201905092915050565b7f696e76616c65642073656e646572000000000000000000000000000000000000600082015250565b60006118ce600e83611887565b91506118d982611898565b602082019050919050565b600060208201905081810360008301526118fd816118c1565b9050919050565b600061191360208401846114e6565b905092915050565b611924816114bd565b82525050565b6000611939602084018461151c565b905092915050565b61194a816114fb565b82525050565b600080fd5b600080fd5b600080fd5b6000808335600160200384360303811261197c5761197b61195a565b5b83810192508235915060208301925067ffffffffffffffff8211156119a4576119a3611950565b5b6001820236038313156119ba576119b9611955565b5b509250929050565b600082825260208201905092915050565b82818337600083830152505050565b6000601f19601f8301169050919050565b60006119ff83856119c2565b9350611a0c8385846119d3565b611a15836119e2565b840190509392505050565b600060c08301611a336000840184611904565b611a40600086018261191b565b50611a4e602084018461192a565b611a5b6020860182611941565b50611a69604084018461192a565b611a766040860182611941565b50611a846060840184611904565b611a91606086018261191b565b50611a9f608084018461195f565b8583036080870152611ab28382846119f3565b92505050611ac360a0840184611904565b611ad060a086018261191b565b508091505092915050565b60006020820190508181036000830152611af58184611a20565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b37826114fb565b9150611b42836114fb565b9250828203905081811115611b5a57611b59611afd565b5b92915050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112611b8c57611b8b611b60565b5b80840192508235915067ffffffffffffffff821115611bae57611bad611b65565b5b602083019250600182023603831315611bca57611bc9611b6a565b5b509250929050565b7f696e73756666696369656e7420696e70757420616d6f756e7400000000000000600082015250565b6000611c08601983611887565b9150611c1382611bd2565b602082019050919050565b60006020820190508181036000830152611c3781611bfb565b9050919050565b611c478161174a565b82525050565b600060a082019050611c626000830188611720565b611c6f6020830187611720565b611c7c6040830186611720565b611c8960608301856115b1565b611c966080830184611c3e565b9695505050505050565b611ca9816117fe565b8114611cb457600080fd5b50565b600081519050611cc681611ca0565b92915050565b60008060408385031215611ce357611ce2611493565b5b6000611cf185828601611845565b9250506020611d0285828601611cb7565b9150509250929050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611d68602683611887565b9150611d7382611d0c565b604082019050919050565b60006020820190508181036000830152611d9781611d5b565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611dd4602083611887565b9150611ddf82611d9e565b602082019050919050565b60006020820190508181036000830152611e0381611dc7565b9050919050565b6000604082019050611e1f6000830185611720565b611e2c60208301846115b1565b9392505050565b600081519050919050565b600081905092915050565b60005b83811015611e67578082015181840152602081019050611e4c565b60008484015250505050565b6000611e7e82611e33565b611e888185611e3e565b9350611e98818560208601611e49565b80840191505092915050565b6000611eb08284611e73565b915081905092915050565b60008115159050919050565b611ed081611ebb565b8114611edb57600080fd5b50565b600081519050611eed81611ec7565b92915050565b600060208284031215611f0957611f08611493565b5b6000611f1784828501611ede565b91505092915050565b7f45786368616e676541676772656761746f723a3a5472616e7366657248656c7060008201527f65723a736166655472616e736665720000000000000000000000000000000000602082015250565b6000611f7c602f83611887565b9150611f8782611f20565b604082019050919050565b60006020820190508181036000830152611fab81611f6f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f646174612074616d706572656400000000000000000000000000000000000000600082015250565b6000612017600d83611887565b915061202282611fe1565b602082019050919050565b600060208201905081810360008301526120468161200a565b9050919050565b60006060820190506120626000830186611720565b61206f6020830185611720565b61207c60408301846115b1565b949350505050565b7f45786368616e676541676772656761746f723a3a5472616e7366657248656c7060008201527f65723a736166655472616e7366657246726f6d00000000000000000000000000602082015250565b60006120e0603383611887565b91506120eb82612084565b604082019050919050565b6000602082019050818103600083015261210f816120d3565b9050919050565b7f45786368616e676541676772656761746f723a3a5472616e7366657248656c7060008201527f65723a73616665417070726f7665000000000000000000000000000000000000602082015250565b6000612172602e83611887565b915061217d82612116565b604082019050919050565b600060208201905081810360008301526121a181612165565b9050919050565b7f45786368616e676541676772656761746f723a3a5361666543616c6c65723a7360008201527f61666543616c6c00000000000000000000000000000000000000000000000000602082015250565b6000612204602783611887565b915061220f826121a8565b604082019050919050565b60006020820190508181036000830152612233816121f7565b9050919050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b612277826119e2565b810181811067ffffffffffffffff821117156122965761229561223f565b5b80604052505050565b60006122a9611489565b90506122b5828261226e565b919050565b600067ffffffffffffffff8211156122d5576122d461223f565b5b6122de826119e2565b9050602081019050919050565b60006122fe6122f9846122ba565b61229f565b90508281526020810184848401111561231a5761231961223a565b5b612325848285611e49565b509392505050565b600082601f830112612342576123416115ff565b5b81516123528482602086016122eb565b91505092915050565b60006020828403121561237157612370611493565b5b600082015167ffffffffffffffff81111561238f5761238e611498565b5b61239b8482850161232d565b91505092915050565b600081519050919050565b60006123ba826123a4565b6123c48185611887565b93506123d4818560208601611e49565b6123dd816119e2565b840191505092915050565b6000602082019050818103600083015261240281846123af565b905092915050565b7f45786368616e676541676772656761746f723a3a5472616e7366657248656c7060008201527f65723a736166655472616e736665724554480000000000000000000000000000602082015250565b6000612466603283611887565b91506124718261240a565b604082019050919050565b6000602082019050818103600083015261249581612459565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b6000612501601883611887565b915061250c826124cb565b602082019050919050565b60006020820190508181036000830152612530816124f4565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b600061256d601f83611887565b915061257882612537565b602082019050919050565b6000602082019050818103600083015261259c81612560565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b60006125ff602283611887565b915061260a826125a3565b604082019050919050565b6000602082019050818103600083015261262e816125f2565b9050919050565b6000819050919050565b61264881612635565b82525050565b6000608082019050612663600083018761263f565b6126706020830186611c3e565b61267d604083018561263f565b61268a606083018461263f565b9594505050505056fea26469706673582212201873c143207726faa6ac6fa84ce9dafa3dfac4729656e9f45ed158e04e86e42264736f6c63430008110033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _WETH common.Address, _PriceProvider common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _WETH, _PriceProvider)
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

// PriceProvider is a free data retrieval call binding the contract method 0x4e1a67f5.
//
// Solidity: function PriceProvider() view returns(address)
func (_Contracts *ContractsCaller) PriceProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "PriceProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceProvider is a free data retrieval call binding the contract method 0x4e1a67f5.
//
// Solidity: function PriceProvider() view returns(address)
func (_Contracts *ContractsSession) PriceProvider() (common.Address, error) {
	return _Contracts.Contract.PriceProvider(&_Contracts.CallOpts)
}

// PriceProvider is a free data retrieval call binding the contract method 0x4e1a67f5.
//
// Solidity: function PriceProvider() view returns(address)
func (_Contracts *ContractsCallerSession) PriceProvider() (common.Address, error) {
	return _Contracts.Contract.PriceProvider(&_Contracts.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contracts *ContractsCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contracts *ContractsSession) WETH() (common.Address, error) {
	return _Contracts.Contract.WETH(&_Contracts.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Contracts *ContractsCallerSession) WETH() (common.Address, error) {
	return _Contracts.Contract.WETH(&_Contracts.CallOpts)
}

// BalanceETH is a free data retrieval call binding the contract method 0xecbdbb32.
//
// Solidity: function balanceETH() view returns(uint256)
func (_Contracts *ContractsCaller) BalanceETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceETH is a free data retrieval call binding the contract method 0xecbdbb32.
//
// Solidity: function balanceETH() view returns(uint256)
func (_Contracts *ContractsSession) BalanceETH() (*big.Int, error) {
	return _Contracts.Contract.BalanceETH(&_Contracts.CallOpts)
}

// BalanceETH is a free data retrieval call binding the contract method 0xecbdbb32.
//
// Solidity: function balanceETH() view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceETH() (*big.Int, error) {
	return _Contracts.Contract.BalanceETH(&_Contracts.CallOpts)
}

// BalanceToken is a free data retrieval call binding the contract method 0x04599012.
//
// Solidity: function balanceToken(address token) view returns(uint256)
func (_Contracts *ContractsCaller) BalanceToken(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceToken", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceToken is a free data retrieval call binding the contract method 0x04599012.
//
// Solidity: function balanceToken(address token) view returns(uint256)
func (_Contracts *ContractsSession) BalanceToken(token common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceToken(&_Contracts.CallOpts, token)
}

// BalanceToken is a free data retrieval call binding the contract method 0x04599012.
//
// Solidity: function balanceToken(address token) view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceToken(token common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceToken(&_Contracts.CallOpts, token)
}

// EstimateAmountOut is a free data retrieval call binding the contract method 0x6110358d.
//
// Solidity: function estimateAmountOut(address provider, address tA, address tB, uint256 amountIn, uint8 version) view returns(uint256 amountOut, uint24 fee)
func (_Contracts *ContractsCaller) EstimateAmountOut(opts *bind.CallOpts, provider common.Address, tA common.Address, tB common.Address, amountIn *big.Int, version uint8) (struct {
	AmountOut *big.Int
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "estimateAmountOut", provider, tA, tB, amountIn, version)

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

// EstimateAmountOut is a free data retrieval call binding the contract method 0x6110358d.
//
// Solidity: function estimateAmountOut(address provider, address tA, address tB, uint256 amountIn, uint8 version) view returns(uint256 amountOut, uint24 fee)
func (_Contracts *ContractsSession) EstimateAmountOut(provider common.Address, tA common.Address, tB common.Address, amountIn *big.Int, version uint8) (struct {
	AmountOut *big.Int
	Fee       *big.Int
}, error) {
	return _Contracts.Contract.EstimateAmountOut(&_Contracts.CallOpts, provider, tA, tB, amountIn, version)
}

// EstimateAmountOut is a free data retrieval call binding the contract method 0x6110358d.
//
// Solidity: function estimateAmountOut(address provider, address tA, address tB, uint256 amountIn, uint8 version) view returns(uint256 amountOut, uint24 fee)
func (_Contracts *ContractsCallerSession) EstimateAmountOut(provider common.Address, tA common.Address, tB common.Address, amountIn *big.Int, version uint8) (struct {
	AmountOut *big.Int
	Fee       *big.Int
}, error) {
	return _Contracts.Contract.EstimateAmountOut(&_Contracts.CallOpts, provider, tA, tB, amountIn, version)
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

// ChangePriceProvider is a paid mutator transaction binding the contract method 0xae4f5be8.
//
// Solidity: function changePriceProvider(address _PriceProvider) returns()
func (_Contracts *ContractsTransactor) ChangePriceProvider(opts *bind.TransactOpts, _PriceProvider common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changePriceProvider", _PriceProvider)
}

// ChangePriceProvider is a paid mutator transaction binding the contract method 0xae4f5be8.
//
// Solidity: function changePriceProvider(address _PriceProvider) returns()
func (_Contracts *ContractsSession) ChangePriceProvider(_PriceProvider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangePriceProvider(&_Contracts.TransactOpts, _PriceProvider)
}

// ChangePriceProvider is a paid mutator transaction binding the contract method 0xae4f5be8.
//
// Solidity: function changePriceProvider(address _PriceProvider) returns()
func (_Contracts *ContractsTransactorSession) ChangePriceProvider(_PriceProvider common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangePriceProvider(&_Contracts.TransactOpts, _PriceProvider)
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

// Swap is a paid mutator transaction binding the contract method 0x2906799b.
//
// Solidity: function swap((address,uint256,uint256,address,bytes,address) data, bytes sig) returns()
func (_Contracts *ContractsTransactor) Swap(opts *bind.TransactOpts, data IExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "swap", data, sig)
}

// Swap is a paid mutator transaction binding the contract method 0x2906799b.
//
// Solidity: function swap((address,uint256,uint256,address,bytes,address) data, bytes sig) returns()
func (_Contracts *ContractsSession) Swap(data IExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Swap(&_Contracts.TransactOpts, data, sig)
}

// Swap is a paid mutator transaction binding the contract method 0x2906799b.
//
// Solidity: function swap((address,uint256,uint256,address,bytes,address) data, bytes sig) returns()
func (_Contracts *ContractsTransactorSession) Swap(data IExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Swap(&_Contracts.TransactOpts, data, sig)
}

// SwapNativeIn is a paid mutator transaction binding the contract method 0x5d86acf1.
//
// Solidity: function swapNativeIn((address,uint256,uint256,address,bytes,address) data, bytes sig) payable returns()
func (_Contracts *ContractsTransactor) SwapNativeIn(opts *bind.TransactOpts, data IExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "swapNativeIn", data, sig)
}

// SwapNativeIn is a paid mutator transaction binding the contract method 0x5d86acf1.
//
// Solidity: function swapNativeIn((address,uint256,uint256,address,bytes,address) data, bytes sig) payable returns()
func (_Contracts *ContractsSession) SwapNativeIn(data IExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SwapNativeIn(&_Contracts.TransactOpts, data, sig)
}

// SwapNativeIn is a paid mutator transaction binding the contract method 0x5d86acf1.
//
// Solidity: function swapNativeIn((address,uint256,uint256,address,bytes,address) data, bytes sig) payable returns()
func (_Contracts *ContractsTransactorSession) SwapNativeIn(data IExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SwapNativeIn(&_Contracts.TransactOpts, data, sig)
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

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Contracts *ContractsTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawETH", to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Contracts *ContractsSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawETH(&_Contracts.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawETH(&_Contracts.TransactOpts, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_Contracts *ContractsTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdrawToken", token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_Contracts *ContractsSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawToken(&_Contracts.TransactOpts, token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.WithdrawToken(&_Contracts.TransactOpts, token, to, amount)
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
