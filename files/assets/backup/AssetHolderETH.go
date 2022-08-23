// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package assets

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AssetHolderWithdrawalAuth is an auto generated low-level Go binding around an user-defined struct.
type AssetHolderWithdrawalAuth struct {
	ChannelID   [32]byte
	Participant common.Address
	Receiver    common.Address
	Amount      *big.Int
}

// AssetsABI is the input ABI used to generate the binding from.
const AssetsABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_adjudicator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"}],\"name\":\"OutcomeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"adjudicator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fundingID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"holdings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"parts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBals\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"subAllocs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"subBalances\",\"type\":\"uint256[]\"}],\"name\":\"setOutcome\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"settled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"channelID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAssetHolder.WithdrawalAuth\",\"name\":\"authorization\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AssetsBin is the compiled bytecode used for deploying new contracts.
var AssetsBin = "0x60806040523480156200001157600080fd5b50604051620019d7380380620019d7833981810160405262000037919081019062000098565b8080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000112565b6000815190506200009281620000f8565b92915050565b600060208284031215620000ab57600080fd5b6000620000bb8482850162000081565b91505092915050565b6000620000d182620000d8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200010381620000c4565b81146200010f57600080fd5b50565b6118b580620001226000396000f3fe6080604052600436106100555760003560e01c80631de26e161461005a5780634ed4283c1461007657806353c2ed8e1461009f57806379aad62e146100ca578063ae9ee18c146100f3578063d945af1d14610130575b600080fd5b610074600480360361006f9190810190610ea6565b61016d565b005b34801561008257600080fd5b5061009d60048036036100989190810190610ee2565b610226565b005b3480156100ab57600080fd5b506100b46104b8565b6040516100c19190611439565b60405180910390f35b3480156100d657600080fd5b506100f160048036036100ec9190810190610dbe565b6104de565b005b3480156100ff57600080fd5b5061011a60048036036101159190810190610d95565b610894565b604051610127919061165a565b60405180910390f35b34801561013c57600080fd5b5061015760048036036101529190810190610d95565b6108ac565b6040516101649190611454565b60405180910390f35b8034146101af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101a69061153f565b60405180910390fd5b6101d481600080858152602001908152602001600020546108cc90919063ffffffff16565b60008084815260200190815260200160002081905550817fcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a98260405161021a919061165a565b60405180910390a25050565b600160008460000135815260200190815260200160002060009054906101000a900460ff1661028a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610281906115ff565b60405180910390fd5b61030a8360405160200161029e919061163f565b60405160208183030381529060405283838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508560200160206103059190810190610d43565b610921565b610349576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610340906115bf565b60405180910390fd5b600061036b84600001358560200160206103669190810190610d43565b6109b8565b905083606001356000808381526020019081526020016000205410156103c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bd9061157f565b60405180910390fd5b6103ef8460600135600080848152602001908152602001600020546109eb90919063ffffffff16565b600080838152602001908152602001600020819055508360400160206104189190810190610d6c565b73ffffffffffffffffffffffffffffffffffffffff166108fc85606001359081150290604051600060405180830381858888f19350505050158015610461573d6000803e3d6000fd5b50807fd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81856060013586604001602061049c9190810190610d6c565b6040516104aa929190611675565b60405180910390a250505050565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461056e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105659061161f565b60405180910390fd5b8585905088889050146105b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ad9061155f565b60405180910390fd5b8181905084849050146105fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f59061151f565b60405180910390fd5b60008484905014610644576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063b906115df565b60405180910390fd5b60001515600160008b815260200190815260200160002060009054906101000a900460ff161515146106ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a29061159f565b60405180910390fd5b60008060008b815260200190815260200160002054905060008060008c815260200190815260200160002081905550600080905060608a8a90506040519080825280602002602001820160405280156107135781602001602082028038833980820191505090505b50905060008090505b8b8b90508110156107ca5760006107538e8e8e8581811061073957fe5b905060200201602061074e9190810190610d43565b6109b8565b90508083838151811061076257fe5b60200260200101818152505061079360008083815260200190815260200160002054866108cc90919063ffffffff16565b94506107ba8b8b848181106107a457fe5b90506020020135856108cc90919063ffffffff16565b935050808060010191505061071c565b5081831061082e5760008090505b8b8b905081101561082c578989828181106107ef57fe5b9050602002013560008084848151811061080557fe5b602002602001015181526020019081526020016000208190555080806001019150506107d8565b505b60018060008e815260200190815260200160002060006101000a81548160ff0219169083151502179055508b7fef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b860405160405180910390a2505050505050505050505050565b60006020528060005260406000206000915090505481565b60016020528060005260406000206000915054906101000a900460ff1681565b600080828401905083811015610917576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090e906114ff565b60405180910390fd5b8091505092915050565b6000806109348580519060200120610a35565b905060006109428286610a65565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561097e57600080fd5b8373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614925050509392505050565b600082826040516020016109cd92919061146f565b60405160208183030381529060405280519060200120905092915050565b6000610a2d83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610b51565b905092915050565b600081604051602001610a489190611413565b604051602081830303815290604052805190602001209050919050565b60006041825114610a795760009050610b4b565b60008060006020850151925060408501519150606085015160001a90507f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c1115610acd5760009350505050610b4b565b601b8160ff1614158015610ae55750601c8160ff1614155b15610af65760009350505050610b4b565b60018682858560405160008152602001604052604051610b199493929190611498565b6020604051602081039080840390855afa158015610b3b573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b6000838311158290610b99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9091906114dd565b60405180910390fd5b5060008385039050809150509392505050565b600081359050610bbb81611816565b92915050565b600081359050610bd08161182d565b92915050565b60008083601f840112610be857600080fd5b8235905067ffffffffffffffff811115610c0157600080fd5b602083019150836020820283011115610c1957600080fd5b9250929050565b60008083601f840112610c3257600080fd5b8235905067ffffffffffffffff811115610c4b57600080fd5b602083019150836020820283011115610c6357600080fd5b9250929050565b60008083601f840112610c7c57600080fd5b8235905067ffffffffffffffff811115610c9557600080fd5b602083019150836020820283011115610cad57600080fd5b9250929050565b600081359050610cc381611844565b92915050565b60008083601f840112610cdb57600080fd5b8235905067ffffffffffffffff811115610cf457600080fd5b602083019150836001820283011115610d0c57600080fd5b9250929050565b600060808284031215610d2557600080fd5b81905092915050565b600081359050610d3d8161185b565b92915050565b600060208284031215610d5557600080fd5b6000610d6384828501610bac565b91505092915050565b600060208284031215610d7e57600080fd5b6000610d8c84828501610bc1565b91505092915050565b600060208284031215610da757600080fd5b6000610db584828501610cb4565b91505092915050565b600080600080600080600080600060a08a8c031215610ddc57600080fd5b6000610dea8c828d01610cb4565b99505060208a013567ffffffffffffffff811115610e0757600080fd5b610e138c828d01610bd6565b985098505060408a013567ffffffffffffffff811115610e3257600080fd5b610e3e8c828d01610c6a565b965096505060608a013567ffffffffffffffff811115610e5d57600080fd5b610e698c828d01610c20565b945094505060808a013567ffffffffffffffff811115610e8857600080fd5b610e948c828d01610c6a565b92509250509295985092959850929598565b60008060408385031215610eb957600080fd5b6000610ec785828601610cb4565b9250506020610ed885828601610d2e565b9150509250929050565b600080600060a08486031215610ef757600080fd5b6000610f0586828701610d13565b935050608084013567ffffffffffffffff811115610f2257600080fd5b610f2e86828701610cc9565b92509250509250925092565b610f4381611792565b82525050565b610f5281611733565b82525050565b610f6181611721565b82525050565b610f7081611721565b82525050565b610f7f81611745565b82525050565b610f8e81611751565b82525050565b610f9d81611751565b82525050565b610fb4610faf82611751565b6117fb565b82525050565b6000610fc58261169e565b610fcf81856116a9565b9350610fdf8185602086016117c8565b610fe881611805565b840191505092915050565b6000611000601c836116ba565b91507f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000830152601c82019050919050565b6000611040601b836116a9565b91507f536166654d6174683a206164646974696f6e206f766572666c6f7700000000006000830152602082019050919050565b60006110806033836116a9565b91507f6c656e677468206f6620737562416c6c6f637320616e642073756242616c616e60008301527f6365732073686f756c6420626520657175616c000000000000000000000000006020830152604082019050919050565b60006110e6601f836116a9565b91507f77726f6e6720616d6f756e74206f662045544820666f72206465706f736974006000830152602082019050919050565b60006111266029836116a9565b91507f7061727469636970616e7473206c656e6774682073686f756c6420657175616c60008301527f2062616c616e63657300000000000000000000000000000000000000000000006020830152604082019050919050565b600061118c601f836116a9565b91507f696e73756666696369656e742045544820666f72207769746864726177616c006000830152602082019050919050565b60006111cc6025836116a9565b91507f747279696e6720746f2073657420616c726561647920736574746c656420636860008301527f616e6e656c0000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611232601d836116a9565b91507f7369676e617475726520766572696669636174696f6e206661696c65640000006000830152602082019050919050565b60006112726023836116a9565b91507f737562416c6c6f63732063757272656e746c79206e6f7420696d706c656d656e60008301527f74656400000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006112d86013836116a9565b91507f6368616e6e656c206e6f7420736574746c6564000000000000000000000000006000830152602082019050919050565b60006113186025836116a9565b91507f63616e206f6e6c792062652063616c6c6564206279207468652061646a75646960008301527f6361746f720000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6080820161138260008301836116f3565b61138f6000850182610f85565b5061139d60208301836116c5565b6113aa6020850182610f58565b506113b860408301836116dc565b6113c56040850182610f49565b506113d3606083018361170a565b6113e060608501826113e6565b50505050565b6113ef8161177b565b82525050565b6113fe8161177b565b82525050565b61140d81611785565b82525050565b600061141e82610ff3565b915061142a8284610fa3565b60208201915081905092915050565b600060208201905061144e6000830184610f67565b92915050565b60006020820190506114696000830184610f76565b92915050565b60006040820190506114846000830185610f94565b6114916020830184610f67565b9392505050565b60006080820190506114ad6000830187610f94565b6114ba6020830186611404565b6114c76040830185610f94565b6114d46060830184610f94565b95945050505050565b600060208201905081810360008301526114f78184610fba565b905092915050565b6000602082019050818103600083015261151881611033565b9050919050565b6000602082019050818103600083015261153881611073565b9050919050565b60006020820190508181036000830152611558816110d9565b9050919050565b6000602082019050818103600083015261157881611119565b9050919050565b600060208201905081810360008301526115988161117f565b9050919050565b600060208201905081810360008301526115b8816111bf565b9050919050565b600060208201905081810360008301526115d881611225565b9050919050565b600060208201905081810360008301526115f881611265565b9050919050565b60006020820190508181036000830152611618816112cb565b9050919050565b600060208201905081810360008301526116388161130b565b9050919050565b60006080820190506116546000830184611371565b92915050565b600060208201905061166f60008301846113f5565b92915050565b600060408201905061168a60008301856113f5565b6116976020830184610f3a565b9392505050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006116d46020840184610bac565b905092915050565b60006116eb6020840184610bc1565b905092915050565b60006117026020840184610cb4565b905092915050565b60006117196020840184610d2e565b905092915050565b600061172c8261175b565b9050919050565b600061173e8261175b565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b600061179d826117a4565b9050919050565b60006117af826117b6565b9050919050565b60006117c18261175b565b9050919050565b60005b838110156117e65780820151818401526020810190506117cb565b838111156117f5576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b61181f81611721565b811461182a57600080fd5b50565b61183681611733565b811461184157600080fd5b50565b61184d81611751565b811461185857600080fd5b50565b6118648161177b565b811461186f57600080fd5b5056fea365627a7a72315820a820c59f0888d9b3a03c793c8bf2a126c0e01a0af9e4c9a03d5cf9316564cd2c6c6578706572696d656e74616cf564736f6c63430005100040"

// DeployAssets deploys a new Ethereum contract, binding an instance of Assets to it.
func DeployAssets(auth *bind.TransactOpts, backend bind.ContractBackend, _adjudicator common.Address) (common.Address, *types.Transaction, *Assets, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssetsBin), backend, _adjudicator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Assets{AssetsCaller: AssetsCaller{contract: contract}, AssetsTransactor: AssetsTransactor{contract: contract}, AssetsFilterer: AssetsFilterer{contract: contract}}, nil
}

// Assets is an auto generated Go binding around an Ethereum contract.
type Assets struct {
	AssetsCaller     // Read-only binding to the contract
	AssetsTransactor // Write-only binding to the contract
	AssetsFilterer   // Log filterer for contract events
}

// AssetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetsSession struct {
	Contract     *Assets           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetsCallerSession struct {
	Contract *AssetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AssetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetsTransactorSession struct {
	Contract     *AssetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetsRaw struct {
	Contract *Assets // Generic contract binding to access the raw methods on
}

// AssetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetsCallerRaw struct {
	Contract *AssetsCaller // Generic read-only contract binding to access the raw methods on
}

// AssetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetsTransactorRaw struct {
	Contract *AssetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssets creates a new instance of Assets, bound to a specific deployed contract.
func NewAssets(address common.Address, backend bind.ContractBackend) (*Assets, error) {
	contract, err := bindAssets(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Assets{AssetsCaller: AssetsCaller{contract: contract}, AssetsTransactor: AssetsTransactor{contract: contract}, AssetsFilterer: AssetsFilterer{contract: contract}}, nil
}

// NewAssetsCaller creates a new read-only instance of Assets, bound to a specific deployed contract.
func NewAssetsCaller(address common.Address, caller bind.ContractCaller) (*AssetsCaller, error) {
	contract, err := bindAssets(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetsCaller{contract: contract}, nil
}

// NewAssetsTransactor creates a new write-only instance of Assets, bound to a specific deployed contract.
func NewAssetsTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetsTransactor, error) {
	contract, err := bindAssets(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetsTransactor{contract: contract}, nil
}

// NewAssetsFilterer creates a new log filterer instance of Assets, bound to a specific deployed contract.
func NewAssetsFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetsFilterer, error) {
	contract, err := bindAssets(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetsFilterer{contract: contract}, nil
}

// bindAssets binds a generic wrapper to an already deployed contract.
func bindAssets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assets *AssetsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assets.Contract.AssetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assets *AssetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assets.Contract.AssetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assets *AssetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assets.Contract.AssetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assets *AssetsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Assets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assets *AssetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assets *AssetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assets.Contract.contract.Transact(opts, method, params...)
}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assets *AssetsCaller) Adjudicator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Assets.contract.Call(opts, &out, "adjudicator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assets *AssetsSession) Adjudicator() (common.Address, error) {
	return _Assets.Contract.Adjudicator(&_Assets.CallOpts)
}

// Adjudicator is a free data retrieval call binding the contract method 0x53c2ed8e.
//
// Solidity: function adjudicator() view returns(address)
func (_Assets *AssetsCallerSession) Adjudicator() (common.Address, error) {
	return _Assets.Contract.Adjudicator(&_Assets.CallOpts)
}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assets *AssetsCaller) Holdings(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Assets.contract.Call(opts, &out, "holdings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assets *AssetsSession) Holdings(arg0 [32]byte) (*big.Int, error) {
	return _Assets.Contract.Holdings(&_Assets.CallOpts, arg0)
}

// Holdings is a free data retrieval call binding the contract method 0xae9ee18c.
//
// Solidity: function holdings(bytes32 ) view returns(uint256)
func (_Assets *AssetsCallerSession) Holdings(arg0 [32]byte) (*big.Int, error) {
	return _Assets.Contract.Holdings(&_Assets.CallOpts, arg0)
}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assets *AssetsCaller) Settled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Assets.contract.Call(opts, &out, "settled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assets *AssetsSession) Settled(arg0 [32]byte) (bool, error) {
	return _Assets.Contract.Settled(&_Assets.CallOpts, arg0)
}

// Settled is a free data retrieval call binding the contract method 0xd945af1d.
//
// Solidity: function settled(bytes32 ) view returns(bool)
func (_Assets *AssetsCallerSession) Settled(arg0 [32]byte) (bool, error) {
	return _Assets.Contract.Settled(&_Assets.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assets *AssetsTransactor) Deposit(opts *bind.TransactOpts, fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assets.contract.Transact(opts, "deposit", fundingID, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assets *AssetsSession) Deposit(fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assets.Contract.Deposit(&_Assets.TransactOpts, fundingID, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x1de26e16.
//
// Solidity: function deposit(bytes32 fundingID, uint256 amount) payable returns()
func (_Assets *AssetsTransactorSession) Deposit(fundingID [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Assets.Contract.Deposit(&_Assets.TransactOpts, fundingID, amount)
}

// SetOutcome is a paid mutator transaction binding the contract method 0x79aad62e.
//
// Solidity: function setOutcome(bytes32 channelID, address[] parts, uint256[] newBals, bytes32[] subAllocs, uint256[] subBalances) returns()
func (_Assets *AssetsTransactor) SetOutcome(opts *bind.TransactOpts, channelID [32]byte, parts []common.Address, newBals []*big.Int, subAllocs [][32]byte, subBalances []*big.Int) (*types.Transaction, error) {
	return _Assets.contract.Transact(opts, "setOutcome", channelID, parts, newBals, subAllocs, subBalances)
}

// SetOutcome is a paid mutator transaction binding the contract method 0x79aad62e.
//
// Solidity: function setOutcome(bytes32 channelID, address[] parts, uint256[] newBals, bytes32[] subAllocs, uint256[] subBalances) returns()
func (_Assets *AssetsSession) SetOutcome(channelID [32]byte, parts []common.Address, newBals []*big.Int, subAllocs [][32]byte, subBalances []*big.Int) (*types.Transaction, error) {
	return _Assets.Contract.SetOutcome(&_Assets.TransactOpts, channelID, parts, newBals, subAllocs, subBalances)
}

// SetOutcome is a paid mutator transaction binding the contract method 0x79aad62e.
//
// Solidity: function setOutcome(bytes32 channelID, address[] parts, uint256[] newBals, bytes32[] subAllocs, uint256[] subBalances) returns()
func (_Assets *AssetsTransactorSession) SetOutcome(channelID [32]byte, parts []common.Address, newBals []*big.Int, subAllocs [][32]byte, subBalances []*big.Int) (*types.Transaction, error) {
	return _Assets.Contract.SetOutcome(&_Assets.TransactOpts, channelID, parts, newBals, subAllocs, subBalances)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4ed4283c.
//
// Solidity: function withdraw((bytes32,address,address,uint256) authorization, bytes signature) returns()
func (_Assets *AssetsTransactor) Withdraw(opts *bind.TransactOpts, authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assets.contract.Transact(opts, "withdraw", authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4ed4283c.
//
// Solidity: function withdraw((bytes32,address,address,uint256) authorization, bytes signature) returns()
func (_Assets *AssetsSession) Withdraw(authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assets.Contract.Withdraw(&_Assets.TransactOpts, authorization, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4ed4283c.
//
// Solidity: function withdraw((bytes32,address,address,uint256) authorization, bytes signature) returns()
func (_Assets *AssetsTransactorSession) Withdraw(authorization AssetHolderWithdrawalAuth, signature []byte) (*types.Transaction, error) {
	return _Assets.Contract.Withdraw(&_Assets.TransactOpts, authorization, signature)
}

// AssetsDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Assets contract.
type AssetsDepositedIterator struct {
	Event *AssetsDeposited // Event containing the contract specifics and raw log

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
func (it *AssetsDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetsDeposited)
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
		it.Event = new(AssetsDeposited)
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
func (it *AssetsDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetsDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetsDeposited represents a Deposited event raised by the Assets contract.
type AssetsDeposited struct {
	FundingID [32]byte
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assets *AssetsFilterer) FilterDeposited(opts *bind.FilterOpts, fundingID [][32]byte) (*AssetsDepositedIterator, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assets.contract.FilterLogs(opts, "Deposited", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return &AssetsDepositedIterator{contract: _Assets.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assets *AssetsFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *AssetsDeposited, fundingID [][32]byte) (event.Subscription, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assets.contract.WatchLogs(opts, "Deposited", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetsDeposited)
				if err := _Assets.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xcd2fe07293de5928c5df9505b65a8d6506f8668dfe81af09090920687edc48a9.
//
// Solidity: event Deposited(bytes32 indexed fundingID, uint256 amount)
func (_Assets *AssetsFilterer) ParseDeposited(log types.Log) (*AssetsDeposited, error) {
	event := new(AssetsDeposited)
	if err := _Assets.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetsOutcomeSetIterator is returned from FilterOutcomeSet and is used to iterate over the raw logs and unpacked data for OutcomeSet events raised by the Assets contract.
type AssetsOutcomeSetIterator struct {
	Event *AssetsOutcomeSet // Event containing the contract specifics and raw log

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
func (it *AssetsOutcomeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetsOutcomeSet)
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
		it.Event = new(AssetsOutcomeSet)
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
func (it *AssetsOutcomeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetsOutcomeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetsOutcomeSet represents a OutcomeSet event raised by the Assets contract.
type AssetsOutcomeSet struct {
	ChannelID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutcomeSet is a free log retrieval operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assets *AssetsFilterer) FilterOutcomeSet(opts *bind.FilterOpts, channelID [][32]byte) (*AssetsOutcomeSetIterator, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Assets.contract.FilterLogs(opts, "OutcomeSet", channelIDRule)
	if err != nil {
		return nil, err
	}
	return &AssetsOutcomeSetIterator{contract: _Assets.contract, event: "OutcomeSet", logs: logs, sub: sub}, nil
}

// WatchOutcomeSet is a free log subscription operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assets *AssetsFilterer) WatchOutcomeSet(opts *bind.WatchOpts, sink chan<- *AssetsOutcomeSet, channelID [][32]byte) (event.Subscription, error) {

	var channelIDRule []interface{}
	for _, channelIDItem := range channelID {
		channelIDRule = append(channelIDRule, channelIDItem)
	}

	logs, sub, err := _Assets.contract.WatchLogs(opts, "OutcomeSet", channelIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetsOutcomeSet)
				if err := _Assets.contract.UnpackLog(event, "OutcomeSet", log); err != nil {
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

// ParseOutcomeSet is a log parse operation binding the contract event 0xef898d6cd3395b6dfe67a3c1923e5c726c1b154e979fb0a25a9c41d0093168b8.
//
// Solidity: event OutcomeSet(bytes32 indexed channelID)
func (_Assets *AssetsFilterer) ParseOutcomeSet(log types.Log) (*AssetsOutcomeSet, error) {
	event := new(AssetsOutcomeSet)
	if err := _Assets.contract.UnpackLog(event, "OutcomeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AssetsWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Assets contract.
type AssetsWithdrawnIterator struct {
	Event *AssetsWithdrawn // Event containing the contract specifics and raw log

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
func (it *AssetsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetsWithdrawn)
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
		it.Event = new(AssetsWithdrawn)
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
func (it *AssetsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetsWithdrawn represents a Withdrawn event raised by the Assets contract.
type AssetsWithdrawn struct {
	FundingID [32]byte
	Amount    *big.Int
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assets *AssetsFilterer) FilterWithdrawn(opts *bind.FilterOpts, fundingID [][32]byte) (*AssetsWithdrawnIterator, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assets.contract.FilterLogs(opts, "Withdrawn", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return &AssetsWithdrawnIterator{contract: _Assets.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assets *AssetsFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *AssetsWithdrawn, fundingID [][32]byte) (event.Subscription, error) {

	var fundingIDRule []interface{}
	for _, fundingIDItem := range fundingID {
		fundingIDRule = append(fundingIDRule, fundingIDItem)
	}

	logs, sub, err := _Assets.contract.WatchLogs(opts, "Withdrawn", fundingIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetsWithdrawn)
				if err := _Assets.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xd0b6e7d0170f56c62f87de6a8a47a0ccf41c86ffb5084d399d8eb62e823f2a81.
//
// Solidity: event Withdrawn(bytes32 indexed fundingID, uint256 amount, address receiver)
func (_Assets *AssetsFilterer) ParseWithdrawn(log types.Log) (*AssetsWithdrawn, error) {
	event := new(AssetsWithdrawn)
	if err := _Assets.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
