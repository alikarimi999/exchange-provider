// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenBridge

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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"effectiveTime\",\"type\":\"uint256\"}],\"name\":\"LogChangeDCRMOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txhash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogSwapin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bindaddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogSwapout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"txhash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swapin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bindaddr\",\"type\":\"address\"}],\"name\":\"Swapout\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeDCRMOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"transferWithPermit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b5060405162001bcc38038062001bcc8339810160408190526200003491620002b1565b83516200004990600090602087019062000158565b5082516200005f90600190602086019062000158565b5060f882901b7fff0000000000000000000000000000000000000000000000000000000000000016608052600580546001600160a01b0319166001600160a01b0383161790554260065560405146907f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f90620000de9060009062000350565b60408051918290038220828201825260018352603160f81b60209384015290516200013193927fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6918691309101620003f5565b60408051601f19818403018152919052805160209091012060a05250620004809350505050565b82805462000166906200042d565b90600052602060002090601f0160209004810192826200018a5760008555620001d5565b82601f10620001a557805160ff1916838001178555620001d5565b82800160010185558215620001d5579182015b82811115620001d5578251825591602001919060010190620001b8565b50620001e3929150620001e7565b5090565b5b80821115620001e35760008155600101620001e8565b600082601f8301126200020f578081fd5b81516001600160401b03808211156200022c576200022c6200046a565b604051601f8301601f19908116603f011681019082821181831017156200025757620002576200046a565b8160405283815260209250868385880101111562000273578485fd5b8491505b8382101562000296578582018301518183018401529082019062000277565b83821115620002a757848385830101525b9695505050505050565b60008060008060808587031215620002c7578384fd5b84516001600160401b0380821115620002de578586fd5b620002ec88838901620001fe565b9550602087015191508082111562000302578485fd5b506200031187828801620001fe565b935050604085015160ff8116811462000328578283fd5b60608601519092506001600160a01b038116811462000345578182fd5b939692955090935050565b81546000908190600281046001808316806200036d57607f831692505b60208084108214156200038e57634e487b7160e01b87526022600452602487fd5b818015620003a55760018114620003b757620003e7565b60ff19861689528489019650620003e7565b620003c28a62000421565b885b86811015620003df5781548b820152908501908301620003c4565b505084890196505b509498975050505050505050565b9485526020850193909352604084019190915260608301526001600160a01b0316608082015260a00190565b60009081526020902090565b6002810460018216806200044257607f821691505b602082108114156200046457634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b60805160f81c60a05161171c620004b0600039600081816105cf0152610e46015260006105ab015261171c6000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c8063628d6cba116100b8578063a9059cbb1161007c578063a9059cbb1461024a578063b524f3a51461025d578063cae9ca5114610270578063d505accf14610283578063dd62ed3e14610298578063ec126c77146102ab57610136565b8063628d6cba146101f457806370a08231146102075780637ecebe001461021a5780638da5cb5b1461022d57806395d89b411461024257610136565b806330adf81f116100ff57806330adf81f146101a9578063313ce567146101b15780633644e515146101c65780634000aea0146101ce578063605629d6146101e157610136565b8062bf26f41461013b57806306fdde0314610159578063095ea7b31461016e57806318160ddd1461018e57806323b872dd14610196575b600080fd5b6101436102be565b604051610150919061138f565b60405180910390f35b6101616102e2565b60405161015091906113ea565b61018161017c3660046111cb565b610370565b6040516101509190611384565b6101436103c8565b6101816101a436600461111f565b6103cf565b610143610585565b6101b96105a9565b60405161015091906115fd565b6101436105cd565b6101816101dc3660046111f4565b6105f1565b6101816101ef36600461115a565b61074f565b6101816102023660046112ba565b61090b565b6101436102153660046110cc565b610980565b6101436102283660046110cc565b610992565b6102356109a4565b6040516101509190611328565b6101616109d1565b6101816102583660046111cb565b6109de565b61018161026b3660046110cc565b610ab8565b61018161027e3660046111f4565b610ba9565b61029661029136600461115a565b610c83565b005b6101436102a63660046110ed565b610da6565b6101816102b9366004611296565b610dc3565b7f42ce63790c28229c123925d83266e77c04d28784552ab68b350a9003226cbd5981565b600080546102ef9061163a565b80601f016020809104026020016040519081016040528092919081815260200182805461031b9061163a565b80156103685780601f1061033d57610100808354040283529160200191610368565b820191906000526020600020905b81548152906001019060200180831161034b57829003601f168201915b505050505081565b3360008181526008602090815260408083206001600160a01b038716808552925280832085905551919290916000805160206116c7833981519152906103b790869061138f565b60405180910390a350600192915050565b6003545b90565b60006001600160a01b0383161515806103f157506001600160a01b0383163014155b6103fa57600080fd5b6001600160a01b03841633146104c1576001600160a01b038416600090815260086020908152604080832033845290915290205460001981146104bf57828110156104605760405162461bcd60e51b81526004016104579061153e565b60405180910390fd5b600061046c8483611623565b6001600160a01b0387166000818152600860209081526040808320338085529252918290208490559051929350916000805160206116c7833981519152906104b590859061138f565b60405180910390a3505b505b6001600160a01b038416600090815260026020526040902054828110156104fa5760405162461bcd60e51b8152600401610457906115b6565b6105048382611623565b6001600160a01b03808716600090815260026020526040808220939093559086168152908120805485929061053a90849061160b565b92505081905550836001600160a01b0316856001600160a01b03166000805160206116a783398151915285604051610572919061138f565b60405180910390a3506001949350505050565b7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b60006001600160a01b03851615158061061357506001600160a01b0385163014155b61061c57600080fd5b336000908152600260205260409020548481101561064c5760405162461bcd60e51b8152600401610457906115b6565b6106568582611623565b33600090815260026020526040808220929092556001600160a01b0388168152908120805487929061068990849061160b565b90915550506040516001600160a01b0387169033906000805160206116a7833981519152906106b990899061138f565b60405180910390a3604051635260769b60e11b81526001600160a01b0387169063a4c0ed36906106f390339089908990899060040161133c565b602060405180830381600087803b15801561070d57600080fd5b505af1158015610721573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107459190611276565b9695505050505050565b6000844211156107715760405162461bcd60e51b815260040161045790611474565b6001600160a01b038816600090815260076020526040812080547f42ce63790c28229c123925d83266e77c04d28784552ab68b350a9003226cbd59918b918b918b9190866107be83611675565b919050558a6040516020016107d896959493929190611398565b6040516020818303038152906040528051906020012090506107fd8982878787610e41565b8061081057506108108982878787610f18565b61081957600080fd5b6001600160a01b03881615158061083957506001600160a01b0388163014155b61084257600080fd5b6001600160a01b0389166000908152600260205260409020548781101561087b5760405162461bcd60e51b8152600401610457906115b6565b6108858882611623565b6001600160a01b03808c1660009081526002602052604080822093909355908b16815290812080548a92906108bb90849061160b565b92505081905550886001600160a01b03168a6001600160a01b03166000805160206116a78339815191528a6040516108f3919061138f565b60405180910390a35060019998505050505050505050565b60006001600160a01b0382166109335760405162461bcd60e51b8152600401610457906114a4565b61093d3384610f4b565b816001600160a01b0316336001600160a01b03167f6b616089d04950dc06c45c6dd787d657980543f89651aec47924752c7d16c888856040516103b7919061138f565b60026020526000908152604090205481565b60076020526000908152604090205481565b600060065442106109c157506005546001600160a01b03166103cc565b506004546001600160a01b031690565b600180546102ef9061163a565b60006001600160a01b038316151580610a0057506001600160a01b0383163014155b610a0957600080fd5b3360009081526002602052604090205482811015610a395760405162461bcd60e51b8152600401610457906115b6565b610a438382611623565b33600090815260026020526040808220929092556001600160a01b03861681529081208054859290610a7690849061160b565b90915550506040516001600160a01b0385169033906000805160206116a783398151915290610aa690879061138f565b60405180910390a35060019392505050565b6000610ac26109a4565b6001600160a01b0316336001600160a01b031614610af25760405162461bcd60e51b8152600401610457906114d9565b6001600160a01b038216610b185760405162461bcd60e51b81526004016104579061143d565b610b206109a4565b600480546001600160a01b03199081166001600160a01b039384161790915560058054909116918416919091179055610b5c426202a30061160b565b60068190556005546004546040516001600160a01b0392831692909116907fe1968d4263a733e2597ef67ea6ad267343bba5f8bf0f99d85190e06b05d824d990600090a45060015b919050565b3360008181526008602090815260408083206001600160a01b038916808552925280832087905551919290916000805160206116c783398151915290610bf090889061138f565b60405180910390a360405162ba451f60e01b81526001600160a01b0386169062ba451f90610c2890339088908890889060040161133c565b602060405180830381600087803b158015610c4257600080fd5b505af1158015610c56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c7a9190611276565b95945050505050565b83421115610ca35760405162461bcd60e51b815260040161045790611474565b6001600160a01b038716600090815260076020526040812080547f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9918a918a918a919086610cf083611675565b9190505589604051602001610d0a96959493929190611398565b604051602081830303815290604052805190602001209050610d2f8882868686610e41565b80610d425750610d428882868686610f18565b610d4b57600080fd5b6001600160a01b038089166000818152600860209081526040808320948c168084529490915290819020899055516000805160206116c783398151915290610d94908a9061138f565b60405180910390a35050505050505050565b600860209081526000928352604080842090915290825290205481565b6000610dcd6109a4565b6001600160a01b0316336001600160a01b031614610dfd5760405162461bcd60e51b8152600401610457906114d9565b610e078383610fef565b826001600160a01b0316847f05d0634fe981be85c22e2942a880821b70095d84e152c3ea3c17a4e4250d9d6184604051610aa6919061138f565b6000807f000000000000000000000000000000000000000000000000000000000000000086604051602001610e7792919061130d565b604051602081830303815290604052805190602001209050600060018287878760405160008152602001604052604051610eb494939291906113cc565b6020604051602081039080840390855afa158015610ed6573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811615801590610f0c5750876001600160a01b0316816001600160a01b0316145b98975050505050505050565b600080610f2486611085565b9050600060018287878760405160008152602001604052604051610eb494939291906113cc565b6001600160a01b038216610f715760405162461bcd60e51b8152600401610457906114fd565b6001600160a01b03821660009081526002602052604081208054839290610f99908490611623565b925050819055508060036000828254610fb29190611623565b90915550506040516000906001600160a01b038416906000805160206116a783398151915290610fe390859061138f565b60405180910390a35050565b6001600160a01b0382166110155760405162461bcd60e51b81526004016104579061157f565b8060036000828254611027919061160b565b90915550506001600160a01b0382166000908152600260205260408120805483929061105490849061160b565b90915550506040516001600160a01b038316906000906000805160206116a783398151915290610fe390859061138f565b60008160405160200161109891906112dc565b604051602081830303815290604052805190602001209050919050565b80356001600160a01b0381168114610ba457600080fd5b6000602082840312156110dd578081fd5b6110e6826110b5565b9392505050565b600080604083850312156110ff578081fd5b611108836110b5565b9150611116602084016110b5565b90509250929050565b600080600060608486031215611133578081fd5b61113c846110b5565b925061114a602085016110b5565b9150604084013590509250925092565b600080600080600080600060e0888a031215611174578283fd5b61117d886110b5565b965061118b602089016110b5565b95506040880135945060608801359350608088013560ff811681146111ae578384fd5b9699959850939692959460a0840135945060c09093013592915050565b600080604083850312156111dd578182fd5b6111e6836110b5565b946020939093013593505050565b60008060008060608587031215611209578384fd5b611212856110b5565b935060208501359250604085013567ffffffffffffffff80821115611235578384fd5b818701915087601f830112611248578384fd5b813581811115611256578485fd5b886020828501011115611267578485fd5b95989497505060200194505050565b600060208284031215611287578081fd5b815180151581146110e6578182fd5b6000806000606084860312156112aa578283fd5b8335925061114a602085016110b5565b600080604083850312156112cc578182fd5b82359150611116602084016110b5565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c810191909152603c0190565b61190160f01b81526002810192909252602282015260420190565b6001600160a01b0391909116815260200190565b6001600160a01b0385168152602081018490526060604082018190528101829052600082846080840137818301608090810191909152601f909201601f191601019392505050565b901515815260200190565b90815260200190565b9586526001600160a01b0394851660208701529290931660408501526060840152608083019190915260a082015260c00190565b93845260ff9290921660208401526040830152606082015260800190565b6000602080835283518082850152825b81811015611416578581018301518582016040015282016113fa565b818111156114275783604083870101525b50601f01601f1916929092016040019392505050565b6020808252601d908201527f6e6577206f776e657220697320746865207a65726f2061646472657373000000604082015260600190565b60208082526016908201527515d15490cc4c0e88115e1c1a5c9959081c195c9b5a5d60521b604082015260600190565b6020808252818101527f62696e64206164647265737320697320746865207a65726f2061646472657373604082015260600190565b6020808252600a908201526937b7363c9037bbb732b960b11b604082015260600190565b60208082526021908201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736040820152607360f81b606082015260800190565b60208082526021908201527f5745524331303a2072657175657374206578636565647320616c6c6f77616e636040820152606560f81b606082015260800190565b6020808252601f908201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604082015260600190565b60208082526027908201527f5745524331303a207472616e7366657220616d6f756e7420657863656564732060408201526662616c616e636560c81b606082015260800190565b60ff91909116815260200190565b6000821982111561161e5761161e611690565b500190565b60008282101561163557611635611690565b500390565b60028104600182168061164e57607f821691505b6020821081141561166f57634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561168957611689611690565b5060010190565b634e487b7160e01b600052601160045260246000fdfeddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925a26469706673582212205052df72629d76d63845a29628826582047991598fc392546a7f8b34941e0a1664736f6c63430008010033000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000000006000000000000000000000000c564ee9f21ed8a2d8e7e76c085740d5e4c5fafbe000000000000000000000000000000000000000000000000000000000000000855534420436f696e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000045553444300000000000000000000000000000000000000000000000000000000",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _decimals uint8, _owner common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _name, _symbol, _decimals, _owner)
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

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contracts *ContractsCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contracts *ContractsSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contracts.Contract.DOMAINSEPARATOR(&_Contracts.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contracts *ContractsCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contracts.Contract.DOMAINSEPARATOR(&_Contracts.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Contracts.Contract.PERMITTYPEHASH(&_Contracts.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Contracts.Contract.PERMITTYPEHASH(&_Contracts.CallOpts)
}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsCaller) TRANSFERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "TRANSFER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsSession) TRANSFERTYPEHASH() ([32]byte, error) {
	return _Contracts.Contract.TRANSFERTYPEHASH(&_Contracts.CallOpts)
}

// TRANSFERTYPEHASH is a free data retrieval call binding the contract method 0x00bf26f4.
//
// Solidity: function TRANSFER_TYPEHASH() view returns(bytes32)
func (_Contracts *ContractsCallerSession) TRANSFERTYPEHASH() ([32]byte, error) {
	return _Contracts.Contract.TRANSFERTYPEHASH(&_Contracts.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Contracts *ContractsCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Contracts *ContractsSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Contracts *ContractsCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Contracts *ContractsSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCallerSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Contracts *ContractsCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Contracts *ContractsSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Nonces(&_Contracts.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.Nonces(&_Contracts.CallOpts, arg0)
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

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// Swapin is a paid mutator transaction binding the contract method 0xec126c77.
//
// Solidity: function Swapin(bytes32 txhash, address account, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) Swapin(opts *bind.TransactOpts, txhash [32]byte, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "Swapin", txhash, account, amount)
}

// Swapin is a paid mutator transaction binding the contract method 0xec126c77.
//
// Solidity: function Swapin(bytes32 txhash, address account, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) Swapin(txhash [32]byte, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Swapin(&_Contracts.TransactOpts, txhash, account, amount)
}

// Swapin is a paid mutator transaction binding the contract method 0xec126c77.
//
// Solidity: function Swapin(bytes32 txhash, address account, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) Swapin(txhash [32]byte, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Swapin(&_Contracts.TransactOpts, txhash, account, amount)
}

// Swapout is a paid mutator transaction binding the contract method 0x628d6cba.
//
// Solidity: function Swapout(uint256 amount, address bindaddr) returns(bool)
func (_Contracts *ContractsTransactor) Swapout(opts *bind.TransactOpts, amount *big.Int, bindaddr common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "Swapout", amount, bindaddr)
}

// Swapout is a paid mutator transaction binding the contract method 0x628d6cba.
//
// Solidity: function Swapout(uint256 amount, address bindaddr) returns(bool)
func (_Contracts *ContractsSession) Swapout(amount *big.Int, bindaddr common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Swapout(&_Contracts.TransactOpts, amount, bindaddr)
}

// Swapout is a paid mutator transaction binding the contract method 0x628d6cba.
//
// Solidity: function Swapout(uint256 amount, address bindaddr) returns(bool)
func (_Contracts *ContractsTransactorSession) Swapout(amount *big.Int, bindaddr common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Swapout(&_Contracts.TransactOpts, amount, bindaddr)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Contracts *ContractsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Contracts *ContractsSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Contracts *ContractsTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_Contracts *ContractsTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveAndCall", spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_Contracts *ContractsSession) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveAndCall(&_Contracts.TransactOpts, spender, value, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes data) returns(bool)
func (_Contracts *ContractsTransactorSession) ApproveAndCall(spender common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveAndCall(&_Contracts.TransactOpts, spender, value, data)
}

// ChangeDCRMOwner is a paid mutator transaction binding the contract method 0xb524f3a5.
//
// Solidity: function changeDCRMOwner(address newOwner) returns(bool)
func (_Contracts *ContractsTransactor) ChangeDCRMOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "changeDCRMOwner", newOwner)
}

// ChangeDCRMOwner is a paid mutator transaction binding the contract method 0xb524f3a5.
//
// Solidity: function changeDCRMOwner(address newOwner) returns(bool)
func (_Contracts *ContractsSession) ChangeDCRMOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeDCRMOwner(&_Contracts.TransactOpts, newOwner)
}

// ChangeDCRMOwner is a paid mutator transaction binding the contract method 0xb524f3a5.
//
// Solidity: function changeDCRMOwner(address newOwner) returns(bool)
func (_Contracts *ContractsTransactorSession) ChangeDCRMOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.ChangeDCRMOwner(&_Contracts.TransactOpts, newOwner)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address target, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contracts *ContractsTransactor) Permit(opts *bind.TransactOpts, target common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "permit", target, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address target, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contracts *ContractsSession) Permit(target common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Permit(&_Contracts.TransactOpts, target, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address target, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Contracts *ContractsTransactorSession) Permit(target common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Permit(&_Contracts.TransactOpts, target, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Contracts *ContractsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Contracts *ContractsSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Contracts *ContractsTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, to, value)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_Contracts *ContractsTransactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferAndCall", to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_Contracts *ContractsSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.TransferAndCall(&_Contracts.TransactOpts, to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool)
func (_Contracts *ContractsTransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.TransferAndCall(&_Contracts.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Contracts *ContractsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Contracts *ContractsSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Contracts *ContractsTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, value)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0x605629d6.
//
// Solidity: function transferWithPermit(address target, address to, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Contracts *ContractsTransactor) TransferWithPermit(opts *bind.TransactOpts, target common.Address, to common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferWithPermit", target, to, value, deadline, v, r, s)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0x605629d6.
//
// Solidity: function transferWithPermit(address target, address to, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Contracts *ContractsSession) TransferWithPermit(target common.Address, to common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.TransferWithPermit(&_Contracts.TransactOpts, target, to, value, deadline, v, r, s)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0x605629d6.
//
// Solidity: function transferWithPermit(address target, address to, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_Contracts *ContractsTransactorSession) TransferWithPermit(target common.Address, to common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.TransferWithPermit(&_Contracts.TransactOpts, target, to, value, deadline, v, r, s)
}

// ContractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contracts contract.
type ContractsApprovalIterator struct {
	Event *ContractsApproval // Event containing the contract specifics and raw log

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
func (it *ContractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApproval)
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
		it.Event = new(ContractsApproval)
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
func (it *ContractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApproval represents a Approval event raised by the Contracts contract.
type ContractsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalIterator{contract: _Contracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApproval)
				if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) ParseApproval(log types.Log) (*ContractsApproval, error) {
	event := new(ContractsApproval)
	if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogChangeDCRMOwnerIterator is returned from FilterLogChangeDCRMOwner and is used to iterate over the raw logs and unpacked data for LogChangeDCRMOwner events raised by the Contracts contract.
type ContractsLogChangeDCRMOwnerIterator struct {
	Event *ContractsLogChangeDCRMOwner // Event containing the contract specifics and raw log

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
func (it *ContractsLogChangeDCRMOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogChangeDCRMOwner)
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
		it.Event = new(ContractsLogChangeDCRMOwner)
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
func (it *ContractsLogChangeDCRMOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogChangeDCRMOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogChangeDCRMOwner represents a LogChangeDCRMOwner event raised by the Contracts contract.
type ContractsLogChangeDCRMOwner struct {
	OldOwner      common.Address
	NewOwner      common.Address
	EffectiveTime *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogChangeDCRMOwner is a free log retrieval operation binding the contract event 0xe1968d4263a733e2597ef67ea6ad267343bba5f8bf0f99d85190e06b05d824d9.
//
// Solidity: event LogChangeDCRMOwner(address indexed oldOwner, address indexed newOwner, uint256 indexed effectiveTime)
func (_Contracts *ContractsFilterer) FilterLogChangeDCRMOwner(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address, effectiveTime []*big.Int) (*ContractsLogChangeDCRMOwnerIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogChangeDCRMOwner", oldOwnerRule, newOwnerRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogChangeDCRMOwnerIterator{contract: _Contracts.contract, event: "LogChangeDCRMOwner", logs: logs, sub: sub}, nil
}

// WatchLogChangeDCRMOwner is a free log subscription operation binding the contract event 0xe1968d4263a733e2597ef67ea6ad267343bba5f8bf0f99d85190e06b05d824d9.
//
// Solidity: event LogChangeDCRMOwner(address indexed oldOwner, address indexed newOwner, uint256 indexed effectiveTime)
func (_Contracts *ContractsFilterer) WatchLogChangeDCRMOwner(opts *bind.WatchOpts, sink chan<- *ContractsLogChangeDCRMOwner, oldOwner []common.Address, newOwner []common.Address, effectiveTime []*big.Int) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}
	var effectiveTimeRule []interface{}
	for _, effectiveTimeItem := range effectiveTime {
		effectiveTimeRule = append(effectiveTimeRule, effectiveTimeItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogChangeDCRMOwner", oldOwnerRule, newOwnerRule, effectiveTimeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogChangeDCRMOwner)
				if err := _Contracts.contract.UnpackLog(event, "LogChangeDCRMOwner", log); err != nil {
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

// ParseLogChangeDCRMOwner is a log parse operation binding the contract event 0xe1968d4263a733e2597ef67ea6ad267343bba5f8bf0f99d85190e06b05d824d9.
//
// Solidity: event LogChangeDCRMOwner(address indexed oldOwner, address indexed newOwner, uint256 indexed effectiveTime)
func (_Contracts *ContractsFilterer) ParseLogChangeDCRMOwner(log types.Log) (*ContractsLogChangeDCRMOwner, error) {
	event := new(ContractsLogChangeDCRMOwner)
	if err := _Contracts.contract.UnpackLog(event, "LogChangeDCRMOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogSwapinIterator is returned from FilterLogSwapin and is used to iterate over the raw logs and unpacked data for LogSwapin events raised by the Contracts contract.
type ContractsLogSwapinIterator struct {
	Event *ContractsLogSwapin // Event containing the contract specifics and raw log

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
func (it *ContractsLogSwapinIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogSwapin)
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
		it.Event = new(ContractsLogSwapin)
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
func (it *ContractsLogSwapinIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogSwapinIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogSwapin represents a LogSwapin event raised by the Contracts contract.
type ContractsLogSwapin struct {
	Txhash  [32]byte
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogSwapin is a free log retrieval operation binding the contract event 0x05d0634fe981be85c22e2942a880821b70095d84e152c3ea3c17a4e4250d9d61.
//
// Solidity: event LogSwapin(bytes32 indexed txhash, address indexed account, uint256 amount)
func (_Contracts *ContractsFilterer) FilterLogSwapin(opts *bind.FilterOpts, txhash [][32]byte, account []common.Address) (*ContractsLogSwapinIterator, error) {

	var txhashRule []interface{}
	for _, txhashItem := range txhash {
		txhashRule = append(txhashRule, txhashItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogSwapin", txhashRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogSwapinIterator{contract: _Contracts.contract, event: "LogSwapin", logs: logs, sub: sub}, nil
}

// WatchLogSwapin is a free log subscription operation binding the contract event 0x05d0634fe981be85c22e2942a880821b70095d84e152c3ea3c17a4e4250d9d61.
//
// Solidity: event LogSwapin(bytes32 indexed txhash, address indexed account, uint256 amount)
func (_Contracts *ContractsFilterer) WatchLogSwapin(opts *bind.WatchOpts, sink chan<- *ContractsLogSwapin, txhash [][32]byte, account []common.Address) (event.Subscription, error) {

	var txhashRule []interface{}
	for _, txhashItem := range txhash {
		txhashRule = append(txhashRule, txhashItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogSwapin", txhashRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogSwapin)
				if err := _Contracts.contract.UnpackLog(event, "LogSwapin", log); err != nil {
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

// ParseLogSwapin is a log parse operation binding the contract event 0x05d0634fe981be85c22e2942a880821b70095d84e152c3ea3c17a4e4250d9d61.
//
// Solidity: event LogSwapin(bytes32 indexed txhash, address indexed account, uint256 amount)
func (_Contracts *ContractsFilterer) ParseLogSwapin(log types.Log) (*ContractsLogSwapin, error) {
	event := new(ContractsLogSwapin)
	if err := _Contracts.contract.UnpackLog(event, "LogSwapin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsLogSwapoutIterator is returned from FilterLogSwapout and is used to iterate over the raw logs and unpacked data for LogSwapout events raised by the Contracts contract.
type ContractsLogSwapoutIterator struct {
	Event *ContractsLogSwapout // Event containing the contract specifics and raw log

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
func (it *ContractsLogSwapoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsLogSwapout)
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
		it.Event = new(ContractsLogSwapout)
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
func (it *ContractsLogSwapoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsLogSwapoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsLogSwapout represents a LogSwapout event raised by the Contracts contract.
type ContractsLogSwapout struct {
	Account  common.Address
	Bindaddr common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSwapout is a free log retrieval operation binding the contract event 0x6b616089d04950dc06c45c6dd787d657980543f89651aec47924752c7d16c888.
//
// Solidity: event LogSwapout(address indexed account, address indexed bindaddr, uint256 amount)
func (_Contracts *ContractsFilterer) FilterLogSwapout(opts *bind.FilterOpts, account []common.Address, bindaddr []common.Address) (*ContractsLogSwapoutIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var bindaddrRule []interface{}
	for _, bindaddrItem := range bindaddr {
		bindaddrRule = append(bindaddrRule, bindaddrItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "LogSwapout", accountRule, bindaddrRule)
	if err != nil {
		return nil, err
	}
	return &ContractsLogSwapoutIterator{contract: _Contracts.contract, event: "LogSwapout", logs: logs, sub: sub}, nil
}

// WatchLogSwapout is a free log subscription operation binding the contract event 0x6b616089d04950dc06c45c6dd787d657980543f89651aec47924752c7d16c888.
//
// Solidity: event LogSwapout(address indexed account, address indexed bindaddr, uint256 amount)
func (_Contracts *ContractsFilterer) WatchLogSwapout(opts *bind.WatchOpts, sink chan<- *ContractsLogSwapout, account []common.Address, bindaddr []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var bindaddrRule []interface{}
	for _, bindaddrItem := range bindaddr {
		bindaddrRule = append(bindaddrRule, bindaddrItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "LogSwapout", accountRule, bindaddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsLogSwapout)
				if err := _Contracts.contract.UnpackLog(event, "LogSwapout", log); err != nil {
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

// ParseLogSwapout is a log parse operation binding the contract event 0x6b616089d04950dc06c45c6dd787d657980543f89651aec47924752c7d16c888.
//
// Solidity: event LogSwapout(address indexed account, address indexed bindaddr, uint256 amount)
func (_Contracts *ContractsFilterer) ParseLogSwapout(log types.Log) (*ContractsLogSwapout, error) {
	event := new(ContractsLogSwapout)
	if err := _Contracts.contract.UnpackLog(event, "LogSwapout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contracts contract.
type ContractsTransferIterator struct {
	Event *ContractsTransfer // Event containing the contract specifics and raw log

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
func (it *ContractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer)
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
		it.Event = new(ContractsTransfer)
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
func (it *ContractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer represents a Transfer event raised by the Contracts contract.
type ContractsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransferIterator{contract: _Contracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer)
				if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) ParseTransfer(log types.Log) (*ContractsTransfer, error) {
	event := new(ContractsTransfer)
	if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
