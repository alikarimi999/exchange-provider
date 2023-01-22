// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	"fmt"
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

// ExchangeAggregatorswapData is an auto generated low-level Go binding around an user-defined struct.
type ExchangeAggregatorswapData struct {
	Input       common.Address
	TotalAmount *big.Int
	FeeAmount   *big.Int
	Swapper     common.Address
	Data        []byte
	Sender      common.Address
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structExchangeAggregator.swapData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structExchangeAggregator.swapData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"input\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structExchangeAggregator.swapData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"swapNativeIn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620026ec380380620026ec8339818101604052810190620000379190620001d5565b620000576200004b6200009f60201b60201c565b620000a760201b60201c565b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000207565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200019d8262000170565b9050919050565b620001af8162000190565b8114620001bb57600080fd5b50565b600081519050620001cf81620001a4565b92915050565b600060208284031215620001ee57620001ed6200016b565b5b6000620001fe84828501620001be565b91505092915050565b6124d580620002176000396000f3fe60806040526004361061009c5760003560e01c8063715018a611610064578063715018a6146101755780638da5cb5b1461018c578063ad5c4648146101b7578063b0867883146101e2578063ecbdbb321461021f578063f2fde38b1461024a5761009c565b806301e33667146100a157806304599012146100ca5780632906799b146101075780634782f779146101305780635d86acf114610159575b600080fd5b3480156100ad57600080fd5b506100c860048036038101906100c391906113d1565b610273565b005b3480156100d657600080fd5b506100f160048036038101906100ec9190611424565b61028b565b6040516100fe9190611460565b60405180910390f35b34801561011357600080fd5b5061012e60048036038101906101299190611504565b61030e565b005b34801561013c57600080fd5b5061015760048036038101906101529190611580565b6104d4565b005b610173600480360381019061016e9190611504565b6104ea565b005b34801561018157600080fd5b5061018a6106a7565b005b34801561019857600080fd5b506101a16106bb565b6040516101ae91906115cf565b60405180910390f35b3480156101c357600080fd5b506101cc6106e4565b6040516101d991906115cf565b60405180910390f35b3480156101ee57600080fd5b506102096004803603810190610204919061172b565b61070a565b60405161021691906115cf565b60405180910390f35b34801561022b57600080fd5b5061023461074e565b6040516102419190611460565b60405180910390f35b34801561025657600080fd5b50610271600480360381019061026c9190611424565b610756565b005b61027b6107d9565b610286838383610857565b505050565b60008173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102c691906115cf565b602060405180830381865afa1580156102e3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030791906117b8565b9050919050565b3373ffffffffffffffffffffffffffffffffffffffff168360a00160208101906103389190611424565b73ffffffffffffffffffffffffffffffffffffffff161461038e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161038590611842565b60405180910390fd5b6103fb836040516020016103a29190611a19565b60405160208183030381529060405283838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061098d565b61041d8360000160208101906104119190611424565b33308660200135610b06565b61045f8360000160208101906104339190611424565b8460600160208101906104469190611424565b8560400135866020013561045a9190611a6a565b610c3f565b6104cf8360600160208101906104759190611424565b60008580608001906104879190611aad565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610d75565b505050565b6104dc6107d9565b6104e68282610e8c565b5050565b3373ffffffffffffffffffffffffffffffffffffffff168360a00160208101906105149190611424565b73ffffffffffffffffffffffffffffffffffffffff161461056a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056190611842565b60405180910390fd5b6105d78360405160200161057e9190611a19565b60405160208183030381529060405283838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505061098d565b826020013534101561061e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061590611b5c565b60405180910390fd5b60008360400135346106309190611a6a565b90506106a18460600160208101906106489190611424565b828680608001906106599190611aad565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050610d75565b50505050565b6106af6107d9565b6106b96000610f8c565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061074682610738856040516020016107249190611a19565b604051602081830303815290604052611050565b61106190919063ffffffff16565b905092915050565b600047905090565b61075e6107d9565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c490611bee565b60405180910390fd5b6107d681610f8c565b50565b6107e1611088565b73ffffffffffffffffffffffffffffffffffffffff166107ff6106bb565b73ffffffffffffffffffffffffffffffffffffffff1614610855576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084c90611c5a565b60405180910390fd5b565b6000808473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8585604051602401610889929190611c7a565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516108d79190611d14565b6000604051808303816000865af19150503d8060008114610914576040519150601f19603f3d011682016040523d82523d6000602084013e610919565b606091505b509150915081801561094757506000815114806109465750808060200190518101906109459190611d63565b5b5b610986576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097d90611e02565b60405180910390fd5b5050505050565b601b60f81b816040815181106109a6576109a5611e22565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506109dd6106bb565b73ffffffffffffffffffffffffffffffffffffffff16610a0e82610a0085611050565b61106190919063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff160315610b0257601c60f81b81604081518110610a4357610a42611e22565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610a7a6106bb565b73ffffffffffffffffffffffffffffffffffffffff16610aab82610a9d85611050565b61106190919063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1614610b01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af890611e9d565b60405180910390fd5b5b5050565b6000808573ffffffffffffffffffffffffffffffffffffffff166323b872dd868686604051602401610b3a93929190611ebd565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610b889190611d14565b6000604051808303816000865af19150503d8060008114610bc5576040519150601f19603f3d011682016040523d82523d6000602084013e610bca565b606091505b5091509150818015610bf85750600081511480610bf7575080806020019051810190610bf69190611d63565b5b5b610c37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2e90611f66565b60405180910390fd5b505050505050565b6000808473ffffffffffffffffffffffffffffffffffffffff1663095ea7b38585604051602401610c71929190611c7a565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610cbf9190611d14565b6000604051808303816000865af19150503d8060008114610cfc576040519150601f19603f3d011682016040523d82523d6000602084013e610d01565b606091505b5091509150818015610d2f5750600081511480610d2e575080806020019051810190610d2d9190611d63565b5b5b610d6e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6590611ff8565b60405180910390fd5b5050505050565b6000808473ffffffffffffffffffffffffffffffffffffffff168484604051610d9e9190611d14565b60006040518083038185875af1925050503d8060008114610ddb576040519150601f19603f3d011682016040523d82523d6000602084013e610de0565b606091505b509150915081610e8557604481511015610e2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e269061208a565b60405180910390fd5b60048101905080806020019051810190610e49919061214b565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7c91906121d8565b60405180910390fd5b5050505050565b60008273ffffffffffffffffffffffffffffffffffffffff1682600067ffffffffffffffff811115610ec157610ec0611600565b5b6040519080825280601f01601f191660200182016040528015610ef35781602001600182028036833780820191505090505b50604051610f019190611d14565b60006040518083038185875af1925050503d8060008114610f3e576040519150601f19603f3d011682016040523d82523d6000602084013e610f43565b606091505b5050905080610f87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f7e9061226c565b60405180910390fd5b505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081805190602001209050919050565b60008060006110708585611090565b9150915061107d816110e1565b819250505092915050565b600033905090565b60008060418351036110d15760008060006020860151925060408601519150606086015160001a90506110c587828585611247565b945094505050506110da565b60006002915091505b9250929050565b600060048111156110f5576110f461228c565b5b8160048111156111085761110761228c565b5b031561124457600160048111156111225761112161228c565b5b8160048111156111355761113461228c565b5b03611175576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161116c90612307565b60405180910390fd5b600260048111156111895761118861228c565b5b81600481111561119c5761119b61228c565b5b036111dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111d390612373565b60405180910390fd5b600360048111156111f0576111ef61228c565b5b8160048111156112035761120261228c565b5b03611243576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123a90612405565b60405180910390fd5b5b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c1115611282576000600391509150611320565b6000600187878787604051600081526020016040526040516112a7949392919061245a565b6020604051602081039080840390855afa1580156112c9573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361131757600060019250925050611320565b80600092509250505b94509492505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006113688261133d565b9050919050565b6113788161135d565b811461138357600080fd5b50565b6000813590506113958161136f565b92915050565b6000819050919050565b6113ae8161139b565b81146113b957600080fd5b50565b6000813590506113cb816113a5565b92915050565b6000806000606084860312156113ea576113e9611333565b5b60006113f886828701611386565b935050602061140986828701611386565b925050604061141a868287016113bc565b9150509250925092565b60006020828403121561143a57611439611333565b5b600061144884828501611386565b91505092915050565b61145a8161139b565b82525050565b60006020820190506114756000830184611451565b92915050565b600080fd5b600060c082840312156114965761149561147b565b5b81905092915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126114c4576114c361149f565b5b8235905067ffffffffffffffff8111156114e1576114e06114a4565b5b6020830191508360018202830111156114fd576114fc6114a9565b5b9250929050565b60008060006040848603121561151d5761151c611333565b5b600084013567ffffffffffffffff81111561153b5761153a611338565b5b61154786828701611480565b935050602084013567ffffffffffffffff81111561156857611567611338565b5b611574868287016114ae565b92509250509250925092565b6000806040838503121561159757611596611333565b5b60006115a585828601611386565b92505060206115b6858286016113bc565b9150509250929050565b6115c98161135d565b82525050565b60006020820190506115e460008301846115c0565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611638826115ef565b810181811067ffffffffffffffff8211171561165757611656611600565b5b80604052505050565b600061166a611329565b9050611676828261162f565b919050565b600067ffffffffffffffff82111561169657611695611600565b5b61169f826115ef565b9050602081019050919050565b82818337600083830152505050565b60006116ce6116c98461167b565b611660565b9050828152602081018484840111156116ea576116e96115ea565b5b6116f58482856116ac565b509392505050565b600082601f8301126117125761171161149f565b5b81356117228482602086016116bb565b91505092915050565b6000806040838503121561174257611741611333565b5b600083013567ffffffffffffffff8111156117605761175f611338565b5b61176c85828601611480565b925050602083013567ffffffffffffffff81111561178d5761178c611338565b5b611799858286016116fd565b9150509250929050565b6000815190506117b2816113a5565b92915050565b6000602082840312156117ce576117cd611333565b5b60006117dc848285016117a3565b91505092915050565b600082825260208201905092915050565b7f696e76616c65642073656e646572000000000000000000000000000000000000600082015250565b600061182c600e836117e5565b9150611837826117f6565b602082019050919050565b6000602082019050818103600083015261185b8161181f565b9050919050565b60006118716020840184611386565b905092915050565b6118828161135d565b82525050565b600061189760208401846113bc565b905092915050565b6118a88161139b565b82525050565b600080fd5b600080fd5b600080fd5b600080833560016020038436030381126118da576118d96118b8565b5b83810192508235915060208301925067ffffffffffffffff821115611902576119016118ae565b5b600182023603831315611918576119176118b3565b5b509250929050565b600082825260208201905092915050565b600061193d8385611920565b935061194a8385846116ac565b611953836115ef565b840190509392505050565b600060c083016119716000840184611862565b61197e6000860182611879565b5061198c6020840184611888565b611999602086018261189f565b506119a76040840184611888565b6119b4604086018261189f565b506119c26060840184611862565b6119cf6060860182611879565b506119dd60808401846118bd565b85830360808701526119f0838284611931565b92505050611a0160a0840184611862565b611a0e60a0860182611879565b508091505092915050565b60006020820190508181036000830152611a33818461195e565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611a758261139b565b9150611a808361139b565b9250828203905081811115611a9857611a97611a3b565b5b92915050565b600080fd5b600080fd5b600080fd5b60008083356001602003843603038112611aca57611ac9611a9e565b5b80840192508235915067ffffffffffffffff821115611aec57611aeb611aa3565b5b602083019250600182023603831315611b0857611b07611aa8565b5b509250929050565b7f696e73756666696369656e7420696e70757420616d6f756e7400000000000000600082015250565b6000611b466019836117e5565b9150611b5182611b10565b602082019050919050565b60006020820190508181036000830152611b7581611b39565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611bd86026836117e5565b9150611be382611b7c565b604082019050919050565b60006020820190508181036000830152611c0781611bcb565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611c446020836117e5565b9150611c4f82611c0e565b602082019050919050565b60006020820190508181036000830152611c7381611c37565b9050919050565b6000604082019050611c8f60008301856115c0565b611c9c6020830184611451565b9392505050565b600081519050919050565b600081905092915050565b60005b83811015611cd7578082015181840152602081019050611cbc565b60008484015250505050565b6000611cee82611ca3565b611cf88185611cae565b9350611d08818560208601611cb9565b80840191505092915050565b6000611d208284611ce3565b915081905092915050565b60008115159050919050565b611d4081611d2b565b8114611d4b57600080fd5b50565b600081519050611d5d81611d37565b92915050565b600060208284031215611d7957611d78611333565b5b6000611d8784828501611d4e565b91505092915050565b7f5472616e7366657248656c7065723a3a736166655472616e736665723a20747260008201527f616e73666572206661696c656400000000000000000000000000000000000000602082015250565b6000611dec602d836117e5565b9150611df782611d90565b604082019050919050565b60006020820190508181036000830152611e1b81611ddf565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f646174612074616d706572656400000000000000000000000000000000000000600082015250565b6000611e87600d836117e5565b9150611e9282611e51565b602082019050919050565b60006020820190508181036000830152611eb681611e7a565b9050919050565b6000606082019050611ed260008301866115c0565b611edf60208301856115c0565b611eec6040830184611451565b949350505050565b7f5472616e7366657248656c7065723a3a7472616e7366657246726f6d3a20747260008201527f616e7366657246726f6d206661696c6564000000000000000000000000000000602082015250565b6000611f506031836117e5565b9150611f5b82611ef4565b604082019050919050565b60006020820190508181036000830152611f7f81611f43565b9050919050565b7f5472616e7366657248656c7065723a3a73616665417070726f76653a2061707060008201527f726f7665206661696c6564000000000000000000000000000000000000000000602082015250565b6000611fe2602b836117e5565b9150611fed82611f86565b604082019050919050565b6000602082019050818103600083015261201181611fd5565b9050919050565b7f5361666543616c6c65723a3a7361666543616c6c3a2063616c6c206661696c6560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b60006120746021836117e5565b915061207f82612018565b604082019050919050565b600060208201905081810360008301526120a381612067565b9050919050565b600067ffffffffffffffff8211156120c5576120c4611600565b5b6120ce826115ef565b9050602081019050919050565b60006120ee6120e9846120aa565b611660565b90508281526020810184848401111561210a576121096115ea565b5b612115848285611cb9565b509392505050565b600082601f8301126121325761213161149f565b5b81516121428482602086016120db565b91505092915050565b60006020828403121561216157612160611333565b5b600082015167ffffffffffffffff81111561217f5761217e611338565b5b61218b8482850161211d565b91505092915050565b600081519050919050565b60006121aa82612194565b6121b481856117e5565b93506121c4818560208601611cb9565b6121cd816115ef565b840191505092915050565b600060208201905081810360008301526121f2818461219f565b905092915050565b7f5472616e7366657248656c7065723a3a736166655472616e736665724554483a60008201527f20455448207472616e73666572206661696c6564000000000000000000000000602082015250565b60006122566034836117e5565b9150612261826121fa565b604082019050919050565b6000602082019050818103600083015261228581612249565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b60006122f16018836117e5565b91506122fc826122bb565b602082019050919050565b60006020820190508181036000830152612320816122e4565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b600061235d601f836117e5565b915061236882612327565b602082019050919050565b6000602082019050818103600083015261238c81612350565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b60006123ef6022836117e5565b91506123fa82612393565b604082019050919050565b6000602082019050818103600083015261241e816123e2565b9050919050565b6000819050919050565b61243881612425565b82525050565b600060ff82169050919050565b6124548161243e565b82525050565b600060808201905061246f600083018761242f565b61247c602083018661244b565b612489604083018561242f565b612496606083018461242f565b9594505050505056fea2646970667358221220e7d92f7092140c18a55a91e1aafc47d621697eee9182d66633a26698f578ad1364736f6c63430008110033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _WETH common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _WETH)
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

// Signer is a free data retrieval call binding the contract method 0xb0867883.
//
// Solidity: function signer((address,uint256,uint256,address,bytes,address) data, bytes sig) pure returns(address)
func (_Contracts *ContractsCaller) Signer(opts *bind.CallOpts, data ExchangeAggregatorswapData, sig []byte) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "signer", data, sig)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0xb0867883.
//
// Solidity: function signer((address,uint256,uint256,address,bytes,address) data, bytes sig) pure returns(address)
func (_Contracts *ContractsSession) Signer(data ExchangeAggregatorswapData, sig []byte) (common.Address, error) {
	return _Contracts.Contract.Signer(&_Contracts.CallOpts, data, sig)
}

// Signer is a free data retrieval call binding the contract method 0xb0867883.
//
// Solidity: function signer((address,uint256,uint256,address,bytes,address) data, bytes sig) pure returns(address)
func (_Contracts *ContractsCallerSession) Signer(data ExchangeAggregatorswapData, sig []byte) (common.Address, error) {
	return _Contracts.Contract.Signer(&_Contracts.CallOpts, data, sig)
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
func (_Contracts *ContractsTransactor) Swap(opts *bind.TransactOpts, data ExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "swap", data, sig)
}

// Swap is a paid mutator transaction binding the contract method 0x2906799b.
//
// Solidity: function swap((address,uint256,uint256,address,bytes,address) data, bytes sig) returns()
func (_Contracts *ContractsSession) Swap(data ExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Swap(&_Contracts.TransactOpts, data, sig)
}

// Swap is a paid mutator transaction binding the contract method 0x2906799b.
//
// Solidity: function swap((address,uint256,uint256,address,bytes,address) data, bytes sig) returns()
func (_Contracts *ContractsTransactorSession) Swap(data ExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.Swap(&_Contracts.TransactOpts, data, sig)
}

// SwapNativeIn is a paid mutator transaction binding the contract method 0x5d86acf1.
//
// Solidity: function swapNativeIn((address,uint256,uint256,address,bytes,address) data, bytes sig) payable returns()
func (_Contracts *ContractsTransactor) SwapNativeIn(opts *bind.TransactOpts, data ExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "swapNativeIn", data, sig)
}

// SwapNativeIn is a paid mutator transaction binding the contract method 0x5d86acf1.
//
// Solidity: function swapNativeIn((address,uint256,uint256,address,bytes,address) data, bytes sig) payable returns()
func (_Contracts *ContractsSession) SwapNativeIn(data ExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SwapNativeIn(&_Contracts.TransactOpts, data, sig)
}

// SwapNativeIn is a paid mutator transaction binding the contract method 0x5d86acf1.
//
// Solidity: function swapNativeIn((address,uint256,uint256,address,bytes,address) data, bytes sig) payable returns()
func (_Contracts *ContractsTransactorSession) SwapNativeIn(data ExchangeAggregatorswapData, sig []byte) (*types.Transaction, error) {
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


func (c *Contracts) ErrSTF() error {
	return fmt.Errorf("execution reverted: ExchangeAggregator::TransferHelper:safeTransferFrom")
}