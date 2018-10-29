// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"net/rpc"
	"strings"

	"github.com/ethereum/go-ethereum"
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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OnepayABI is the input ABI used to generate the binding from.
const OnepayABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_rnd\",\"type\":\"uint256\"}],\"name\":\"getHistoryWinerInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_rnd\",\"type\":\"uint256\"},{\"name\":\"_plyer\",\"type\":\"address\"}],\"name\":\"getplyrTokenIdByAdde\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"buyCore\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"round_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_eth\",\"type\":\"uint256\"}],\"name\":\"alterMaxEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pot_\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentRoundInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MaxEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_rnd\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokenid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_winer\",\"type\":\"address\"}],\"name\":\"BuyCore\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_length\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_index\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_pot\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_rnd\",\"type\":\"uint256\"}],\"name\":\"EndRount\",\"type\":\"event\"}]"

// OnepayBin is the compiled bytecode used for deploying new contracts.
const OnepayBin = `60806040527308bc39742ba7398774768128eb285eb72bfd29e9600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160025567016345785d8a0000600455612710600b5534801561007c57600080fd5b5033600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610f8a806100cd6000396000f30060806040526004361061008e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630140d3f4146100a55780630ba5c57c14610127578063309e84fc146101c957806337c30a98146101d357806339b98764146101fe5780634b1cff0b1461022b578063747dff4214610256578063cd48578f1461028f575b34801561009a57600080fd5b506100a36102ba565b005b3480156100b157600080fd5b506100d0600480360381019080803590602001909291905050506106b9565b604051808581526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200194505050505060405180910390f35b34801561013357600080fd5b5061017260048036038101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061074c565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156101b557808201518184015260208101905061019a565b505050509050019250505060405180910390f35b6101d16102ba565b005b3480156101df57600080fd5b506101e86107f5565b6040518082815260200191505060405180910390f35b34801561020a57600080fd5b50610229600480360381019080803590602001909291905050506107fb565b005b34801561023757600080fd5b506102406108ca565b6040518082815260200191505060405180910390f35b34801561026257600080fd5b5061026b6108d0565b60405180848152602001838152602001828152602001935050505060405180910390f35b34801561029b57600080fd5b506102a4610908565b6040518082815260200191505060405180910390f35b60008060008060008060008034662386f26fc1000081101515156102dd57600080fd5b3398503497506102fd662386f26fc100008961090e90919063ffffffff16565b96506000662386f26fc100008981151561031357fe5b06111561039057610346610337662386f26fc100008961092990919063ffffffff16565b896109cd90919063ffffffff16565b95508873ffffffffffffffffffffffffffffffffffffffff166108fc879081150290604051600060405180830381858888f1935050505015801561038e573d6000803e3d6000fd5b505b6103be6103ad662386f26fc100008961092990919063ffffffff16565b600354610a5290919063ffffffff16565b600381905550600094505b868510156103e7576103da89610ada565b84806001019550506103c9565b6004546003541015156106ae576103fc610cdc565b9350600960006002548152602001908152602001600020600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692508273ffffffffffffffffffffffffffffffffffffffff166108fc61048a600a61047c600860035461092990919063ffffffff16565b61090e90919063ffffffff16565b9081150290604051600060405180830381858888f193505050501580156104b5573d6000803e3d6000fd5b506104f36104e2600a6104d4600860035461092990919063ffffffff16565b61090e90919063ffffffff16565b6003546109cd90919063ffffffff16565b9150600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f1935050505015801561055d573d6000803e3d6000fd5b50426000806002548152602001908152602001600020600001819055508360008060025481526020019081526020016000206001018190555082600080600254815260200190815260200160002060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610615600a610607600860035461092990919063ffffffff16565b61090e90919063ffffffff16565b6000806002548152602001908152602001600020600301819055507fe3d03f35aedf8f3d0610d5d596e5bb06dcc0223d0b6a22c64976987f91cd92d26002548585604051808481526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001935050505060405180910390a16106ad610eec565b5b505050505050505050565b6000806000806000859050600080828152602001908152602001600020600001546000808381526020019081526020016000206001015460008084815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600080858152602001908152602001600020600301549450945094509450509193509193565b60606006600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156107e857602002820191906000526020600020905b8154815260200190600101908083116107d4575b5050505050905092915050565b60025481565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f596f757220617265206e6f742061646d696e000000000000000000000000000081525060200191505060405180910390fd5b8060058190555050565b60035481565b600080600080600060025491506007600083815260200190815260200160002080549050905081816003549450945094505050909192565b60045481565b600080828481151561091c57fe5b0490508091505092915050565b60008083141561093c57600090506109c7565b818302905081838281151561094d57fe5b041415156109c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f536166654d617468206d756c206661696c65640000000000000000000000000081525060200191505060405180910390fd5b8090505b92915050565b6000828211151515610a47576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f536166654d61746820737562206661696c65640000000000000000000000000081525060200191505060405180910390fd5b818303905092915050565b60008183019050828110151515610ad1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260138152602001807f536166654d61746820616464206661696c65640000000000000000000000000081525060200191505060405180910390fd5b80905092915050565b600080610af36001600b54610a5290919063ffffffff16565b600b81905550600b54915060066000600254815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002082908060018154018082558091505090600182039060005260206000200160009091929091909150555060076000600254815260200190815260200160002080549050905060076000600254815260200190815260200160002082908060018154018082558091505090600182039060005260206000200160009091929091909150555082600960006002548152602001908152602001600020600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816008600060025481526020019081526020016000206000838152602001908152602001600020819055507fe671499be20312df72e417477296255457d7f4de16a88383df5446449000f004838383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a1505050565b60008060008060076000600254815260200190815260200160002080549050925082610de245604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083101515610d565780518252602082019150602081019050602083039250610d31565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060019004610dd444610dc63373ffffffffffffffffffffffffffffffffffffffff16610db84342610a5290919063ffffffff16565b610a5290919063ffffffff16565b610a5290919063ffffffff16565b610a5290919063ffffffff16565b604051602001808281526020019150506040516020818303038152906040526040518082805190602001908083835b602083101515610e365780518252602082019150602081019050602083039250610e11565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902060019004811515610e7057fe5b06915060086000600254815260200190815260200160002060008381526020019081526020016000205490507fa01a72713bf837059e3a668d28f0de277fb7f24f2a4e95bf926703c95b5f12b283838360405180848152602001838152602001828152602001935050505060405180910390a180935050505090565b7f017aec1d8a5efd0e01715e4b10aababb8795cd378b5744590eade4cf08310e58600354600254604051808381526020018281526020019250505060405180910390a1610f456001600254610a5290919063ffffffff16565b60028190555060006003819055506005546004819055505600a165627a7a7230582095e1d0f4053746398d61c9abac4d62354043a36beedefef6ebd8bc646392001a0029`

// DeployOnepay deploys a new Ethereum contract, binding an instance of Onepay to it.
func DeployOnepay(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Onepay, error) {
	parsed, err := abi.JSON(strings.NewReader(OnepayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OnepayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Onepay{OnepayCaller: OnepayCaller{contract: contract}, OnepayTransactor: OnepayTransactor{contract: contract}, OnepayFilterer: OnepayFilterer{contract: contract}}, nil
}

// Onepay is an auto generated Go binding around an Ethereum contract.
type Onepay struct {
	OnepayCaller     // Read-only binding to the contract
	OnepayTransactor // Write-only binding to the contract
	OnepayFilterer   // Log filterer for contract events
}

// OnepayCaller is an auto generated read-only Go binding around an Ethereum contract.
type OnepayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnepayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OnepayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnepayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OnepayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OnepaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OnepaySession struct {
	Contract     *Onepay           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnepayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OnepayCallerSession struct {
	Contract *OnepayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OnepayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OnepayTransactorSession struct {
	Contract     *OnepayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OnepayRaw is an auto generated low-level Go binding around an Ethereum contract.
type OnepayRaw struct {
	Contract *Onepay // Generic contract binding to access the raw methods on
}

// OnepayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OnepayCallerRaw struct {
	Contract *OnepayCaller // Generic read-only contract binding to access the raw methods on
}

// OnepayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OnepayTransactorRaw struct {
	Contract *OnepayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOnepay creates a new instance of Onepay, bound to a specific deployed contract.
func NewOnepay(address common.Address, backend bind.ContractBackend) (*Onepay, error) {
	contract, err := bindOnepay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Onepay{OnepayCaller: OnepayCaller{contract: contract}, OnepayTransactor: OnepayTransactor{contract: contract}, OnepayFilterer: OnepayFilterer{contract: contract}}, nil
}

// NewOnepayCaller creates a new read-only instance of Onepay, bound to a specific deployed contract.
func NewOnepayCaller(address common.Address, caller bind.ContractCaller) (*OnepayCaller, error) {
	contract, err := bindOnepay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OnepayCaller{contract: contract}, nil
}

// NewOnepayTransactor creates a new write-only instance of Onepay, bound to a specific deployed contract.
func NewOnepayTransactor(address common.Address, transactor bind.ContractTransactor) (*OnepayTransactor, error) {
	contract, err := bindOnepay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OnepayTransactor{contract: contract}, nil
}

// NewOnepayFilterer creates a new log filterer instance of Onepay, bound to a specific deployed contract.
func NewOnepayFilterer(address common.Address, filterer bind.ContractFilterer) (*OnepayFilterer, error) {
	contract, err := bindOnepay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OnepayFilterer{contract: contract}, nil
}

// bindOnepay binds a generic wrapper to an already deployed contract.
func bindOnepay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OnepayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Onepay *OnepayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Onepay.Contract.OnepayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Onepay *OnepayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Onepay.Contract.OnepayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Onepay *OnepayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Onepay.Contract.OnepayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Onepay *OnepayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Onepay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Onepay *OnepayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Onepay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Onepay *OnepayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Onepay.Contract.contract.Transact(opts, method, params...)
}

// MaxEth is a free data retrieval call binding the contract method 0xcd48578f.
//
// Solidity: function MaxEth() constant returns(uint256)
func (_Onepay *OnepayCaller) MaxEth(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Onepay.contract.Call(opts, out, "MaxEth")
	return *ret0, err
}

// MaxEth is a free data retrieval call binding the contract method 0xcd48578f.
//
// Solidity: function MaxEth() constant returns(uint256)
func (_Onepay *OnepaySession) MaxEth() (*big.Int, error) {
	return _Onepay.Contract.MaxEth(&_Onepay.CallOpts)
}

// MaxEth is a free data retrieval call binding the contract method 0xcd48578f.
//
// Solidity: function MaxEth() constant returns(uint256)
func (_Onepay *OnepayCallerSession) MaxEth() (*big.Int, error) {
	return _Onepay.Contract.MaxEth(&_Onepay.CallOpts)
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() constant returns(uint256, uint256, uint256)
func (_Onepay *OnepayCaller) GetCurrentRoundInfo(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Onepay.contract.Call(opts, out, "getCurrentRoundInfo")
	return *ret0, *ret1, *ret2, err
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() constant returns(uint256, uint256, uint256)
func (_Onepay *OnepaySession) GetCurrentRoundInfo() (*big.Int, *big.Int, *big.Int, error) {
	return _Onepay.Contract.GetCurrentRoundInfo(&_Onepay.CallOpts)
}

// GetCurrentRoundInfo is a free data retrieval call binding the contract method 0x747dff42.
//
// Solidity: function getCurrentRoundInfo() constant returns(uint256, uint256, uint256)
func (_Onepay *OnepayCallerSession) GetCurrentRoundInfo() (*big.Int, *big.Int, *big.Int, error) {
	return _Onepay.Contract.GetCurrentRoundInfo(&_Onepay.CallOpts)
}

// GetHistoryWinerInfo is a free data retrieval call binding the contract method 0x0140d3f4.
//
// Solidity: function getHistoryWinerInfo(_rnd uint256) constant returns(uint256, uint256, address, uint256)
func (_Onepay *OnepayCaller) GetHistoryWinerInfo(opts *bind.CallOpts, _rnd *big.Int) (*big.Int, *big.Int, common.Address, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Onepay.contract.Call(opts, out, "getHistoryWinerInfo", _rnd)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetHistoryWinerInfo is a free data retrieval call binding the contract method 0x0140d3f4.
//
// Solidity: function getHistoryWinerInfo(_rnd uint256) constant returns(uint256, uint256, address, uint256)
func (_Onepay *OnepaySession) GetHistoryWinerInfo(_rnd *big.Int) (*big.Int, *big.Int, common.Address, *big.Int, error) {
	return _Onepay.Contract.GetHistoryWinerInfo(&_Onepay.CallOpts, _rnd)
}

// GetHistoryWinerInfo is a free data retrieval call binding the contract method 0x0140d3f4.
//
// Solidity: function getHistoryWinerInfo(_rnd uint256) constant returns(uint256, uint256, address, uint256)
func (_Onepay *OnepayCallerSession) GetHistoryWinerInfo(_rnd *big.Int) (*big.Int, *big.Int, common.Address, *big.Int, error) {
	return _Onepay.Contract.GetHistoryWinerInfo(&_Onepay.CallOpts, _rnd)
}

// GetplyrTokenIdByAdde is a free data retrieval call binding the contract method 0x0ba5c57c.
//
// Solidity: function getplyrTokenIdByAdde(_rnd uint256, _plyer address) constant returns(uint256[])
func (_Onepay *OnepayCaller) GetplyrTokenIdByAdde(opts *bind.CallOpts, _rnd *big.Int, _plyer common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Onepay.contract.Call(opts, out, "getplyrTokenIdByAdde", _rnd, _plyer)
	return *ret0, err
}

// GetplyrTokenIdByAdde is a free data retrieval call binding the contract method 0x0ba5c57c.
//
// Solidity: function getplyrTokenIdByAdde(_rnd uint256, _plyer address) constant returns(uint256[])
func (_Onepay *OnepaySession) GetplyrTokenIdByAdde(_rnd *big.Int, _plyer common.Address) ([]*big.Int, error) {
	return _Onepay.Contract.GetplyrTokenIdByAdde(&_Onepay.CallOpts, _rnd, _plyer)
}

// GetplyrTokenIdByAdde is a free data retrieval call binding the contract method 0x0ba5c57c.
//
// Solidity: function getplyrTokenIdByAdde(_rnd uint256, _plyer address) constant returns(uint256[])
func (_Onepay *OnepayCallerSession) GetplyrTokenIdByAdde(_rnd *big.Int, _plyer common.Address) ([]*big.Int, error) {
	return _Onepay.Contract.GetplyrTokenIdByAdde(&_Onepay.CallOpts, _rnd, _plyer)
}

// Pot is a free data retrieval call binding the contract method 0x4b1cff0b.
//
// Solidity: function pot_() constant returns(uint256)
func (_Onepay *OnepayCaller) Pot(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Onepay.contract.Call(opts, out, "pot_")
	return *ret0, err
}

// Pot is a free data retrieval call binding the contract method 0x4b1cff0b.
//
// Solidity: function pot_() constant returns(uint256)
func (_Onepay *OnepaySession) Pot() (*big.Int, error) {
	return _Onepay.Contract.Pot(&_Onepay.CallOpts)
}

// Pot is a free data retrieval call binding the contract method 0x4b1cff0b.
//
// Solidity: function pot_() constant returns(uint256)
func (_Onepay *OnepayCallerSession) Pot() (*big.Int, error) {
	return _Onepay.Contract.Pot(&_Onepay.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x37c30a98.
//
// Solidity: function round_() constant returns(uint256)
func (_Onepay *OnepayCaller) Round(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Onepay.contract.Call(opts, out, "round_")
	return *ret0, err
}

// Round is a free data retrieval call binding the contract method 0x37c30a98.
//
// Solidity: function round_() constant returns(uint256)
func (_Onepay *OnepaySession) Round() (*big.Int, error) {
	return _Onepay.Contract.Round(&_Onepay.CallOpts)
}

// Round is a free data retrieval call binding the contract method 0x37c30a98.
//
// Solidity: function round_() constant returns(uint256)
func (_Onepay *OnepayCallerSession) Round() (*big.Int, error) {
	return _Onepay.Contract.Round(&_Onepay.CallOpts)
}

// AlterMaxEth is a paid mutator transaction binding the contract method 0x39b98764.
//
// Solidity: function alterMaxEth(_eth uint256) returns()
func (_Onepay *OnepayTransactor) AlterMaxEth(opts *bind.TransactOpts, _eth *big.Int) (*types.Transaction, error) {
	return _Onepay.contract.Transact(opts, "alterMaxEth", _eth)
}

// AlterMaxEth is a paid mutator transaction binding the contract method 0x39b98764.
//
// Solidity: function alterMaxEth(_eth uint256) returns()
func (_Onepay *OnepaySession) AlterMaxEth(_eth *big.Int) (*types.Transaction, error) {
	return _Onepay.Contract.AlterMaxEth(&_Onepay.TransactOpts, _eth)
}

// AlterMaxEth is a paid mutator transaction binding the contract method 0x39b98764.
//
// Solidity: function alterMaxEth(_eth uint256) returns()
func (_Onepay *OnepayTransactorSession) AlterMaxEth(_eth *big.Int) (*types.Transaction, error) {
	return _Onepay.Contract.AlterMaxEth(&_Onepay.TransactOpts, _eth)
}

// BuyCore is a paid mutator transaction binding the contract method 0x309e84fc.
//
// Solidity: function buyCore() returns()
func (_Onepay *OnepayTransactor) BuyCore(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Onepay.contract.Transact(opts, "buyCore")
}

// BuyCore is a paid mutator transaction binding the contract method 0x309e84fc.
//
// Solidity: function buyCore() returns()
func (_Onepay *OnepaySession) BuyCore() (*types.Transaction, error) {
	return _Onepay.Contract.BuyCore(&_Onepay.TransactOpts)
}

// BuyCore is a paid mutator transaction binding the contract method 0x309e84fc.
//
// Solidity: function buyCore() returns()
func (_Onepay *OnepayTransactorSession) BuyCore() (*types.Transaction, error) {
	return _Onepay.Contract.BuyCore(&_Onepay.TransactOpts)
}

// OnepayBuyCoreIterator is returned from FilterBuyCore and is used to iterate over the raw logs and unpacked data for BuyCore events raised by the Onepay contract.
type OnepayBuyCoreIterator struct {
	Event *OnepayBuyCore // Event containing the contract specifics and raw log

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
func (it *OnepayBuyCoreIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnepayBuyCore)
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
		it.Event = new(OnepayBuyCore)
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
func (it *OnepayBuyCoreIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnepayBuyCoreIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnepayBuyCore represents a BuyCore event raised by the Onepay contract.
type OnepayBuyCore struct {
	Rnd     *big.Int
	Tokenid *big.Int
	Winer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyCore is a free log retrieval operation binding the contract event 0xe3d03f35aedf8f3d0610d5d596e5bb06dcc0223d0b6a22c64976987f91cd92d2.
//
// Solidity: e BuyCore(_rnd uint256, _tokenid uint256, _winer address)
func (_Onepay *OnepayFilterer) FilterBuyCore(opts *bind.FilterOpts) (*OnepayBuyCoreIterator, error) {

	logs, sub, err := _Onepay.contract.FilterLogs(opts, "BuyCore")
	if err != nil {
		return nil, err
	}
	return &OnepayBuyCoreIterator{contract: _Onepay.contract, event: "BuyCore", logs: logs, sub: sub}, nil
}

// WatchBuyCore is a free log subscription operation binding the contract event 0xe3d03f35aedf8f3d0610d5d596e5bb06dcc0223d0b6a22c64976987f91cd92d2.
//
// Solidity: e BuyCore(_rnd uint256, _tokenid uint256, _winer address)
func (_Onepay *OnepayFilterer) WatchBuyCore(opts *bind.WatchOpts, sink chan<- *OnepayBuyCore) (event.Subscription, error) {

	logs, sub, err := _Onepay.contract.WatchLogs(opts, "BuyCore")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnepayBuyCore)
				if err := _Onepay.contract.UnpackLog(event, "BuyCore", log); err != nil {
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

// OnepayEndRountIterator is returned from FilterEndRount and is used to iterate over the raw logs and unpacked data for EndRount events raised by the Onepay contract.
type OnepayEndRountIterator struct {
	Event *OnepayEndRount // Event containing the contract specifics and raw log

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
func (it *OnepayEndRountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnepayEndRount)
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
		it.Event = new(OnepayEndRount)
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
func (it *OnepayEndRountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnepayEndRountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnepayEndRount represents a EndRount event raised by the Onepay contract.
type OnepayEndRount struct {
	Pot *big.Int
	Rnd *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEndRount is a free log retrieval operation binding the contract event 0x017aec1d8a5efd0e01715e4b10aababb8795cd378b5744590eade4cf08310e58.
//
// Solidity: e EndRount(_pot uint256, _rnd uint256)
func (_Onepay *OnepayFilterer) FilterEndRount(opts *bind.FilterOpts) (*OnepayEndRountIterator, error) {

	logs, sub, err := _Onepay.contract.FilterLogs(opts, "EndRount")
	if err != nil {
		return nil, err
	}
	return &OnepayEndRountIterator{contract: _Onepay.contract, event: "EndRount", logs: logs, sub: sub}, nil
}

// WatchEndRount is a free log subscription operation binding the contract event 0x017aec1d8a5efd0e01715e4b10aababb8795cd378b5744590eade4cf08310e58.
//
// Solidity: e EndRount(_pot uint256, _rnd uint256)
func (_Onepay *OnepayFilterer) WatchEndRount(opts *bind.WatchOpts, sink chan<- *OnepayEndRount) (event.Subscription, error) {

	logs, sub, err := _Onepay.contract.WatchLogs(opts, "EndRount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnepayEndRount)
				if err := _Onepay.contract.UnpackLog(event, "EndRount", log); err != nil {
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

// OnepayWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Onepay contract.
type OnepayWithdrawIterator struct {
	Event *OnepayWithdraw // Event containing the contract specifics and raw log

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
func (it *OnepayWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnepayWithdraw)
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
		it.Event = new(OnepayWithdraw)
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
func (it *OnepayWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnepayWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnepayWithdraw represents a Withdraw event raised by the Onepay contract.
type OnepayWithdraw struct {
	Length  *big.Int
	Index   *big.Int
	Tokenid *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xa01a72713bf837059e3a668d28f0de277fb7f24f2a4e95bf926703c95b5f12b2.
//
// Solidity: e Withdraw(_length uint256, _index uint256, _tokenid uint256)
func (_Onepay *OnepayFilterer) FilterWithdraw(opts *bind.FilterOpts) (*OnepayWithdrawIterator, error) {

	logs, sub, err := _Onepay.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &OnepayWithdrawIterator{contract: _Onepay.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xa01a72713bf837059e3a668d28f0de277fb7f24f2a4e95bf926703c95b5f12b2.
//
// Solidity: e Withdraw(_length uint256, _index uint256, _tokenid uint256)
func (_Onepay *OnepayFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *OnepayWithdraw) (event.Subscription, error) {

	logs, sub, err := _Onepay.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnepayWithdraw)
				if err := _Onepay.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// OnepayBuyTokenIterator is returned from FilterBuyToken and is used to iterate over the raw logs and unpacked data for BuyToken events raised by the Onepay contract.
type OnepayBuyTokenIterator struct {
	Event *OnepayBuyToken // Event containing the contract specifics and raw log

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
func (it *OnepayBuyTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OnepayBuyToken)
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
		it.Event = new(OnepayBuyToken)
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
func (it *OnepayBuyTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OnepayBuyTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OnepayBuyToken represents a BuyToken event raised by the Onepay contract.
type OnepayBuyToken struct {
	From    common.Address
	Tokenid *big.Int
	Index   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuyToken is a free log retrieval operation binding the contract event 0xe671499be20312df72e417477296255457d7f4de16a88383df5446449000f004.
//
// Solidity: e buyToken(_from address, _tokenid uint256, _index uint256)
func (_Onepay *OnepayFilterer) FilterBuyToken(opts *bind.FilterOpts) (*OnepayBuyTokenIterator, error) {

	logs, sub, err := _Onepay.contract.FilterLogs(opts, "buyToken")
	if err != nil {
		return nil, err
	}
	return &OnepayBuyTokenIterator{contract: _Onepay.contract, event: "buyToken", logs: logs, sub: sub}, nil
}

// WatchBuyToken is a free log subscription operation binding the contract event 0xe671499be20312df72e417477296255457d7f4de16a88383df5446449000f004.
//
// Solidity: e buyToken(_from address, _tokenid uint256, _index uint256)
func (_Onepay *OnepayFilterer) WatchBuyToken(opts *bind.WatchOpts, sink chan<- *OnepayBuyToken) (event.Subscription, error) {

	logs, sub, err := _Onepay.contract.WatchLogs(opts, "buyToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OnepayBuyToken)
				if err := _Onepay.contract.UnpackLog(event, "buyToken", log); err != nil {
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
const key = `{"address":"6b9161be3d3f711e495bcedbf937512897f3ac17","crypto":{"cipher":"aes-128-ctr","cipherparams":{"iv":"e85b7091599ca649d28e8d01e83e669f"},"ciphertext":"d3e7db0d73c1278688eff7051dd4c356021756fcd235ff3d69d1779193b0db53","kdf":"scrypt","kdfparams":{"dklen":32,"n":65536,"p":1,"r":8,"salt":"fbc6ac9ba88c208a1c3b462cf3a2b1750360a2bdc4516d614f5d031ae61f71c2"},"mac":"b4bf3228f5b8143eb9289e8af103d9ca84cfd0b9a9e45389bc5b89206b48719e"},"id":"bbe3fb84-fbc2-47fb-9f25-c2790a93214f","version":3}`

func main() {
	//创建基于IPC的RPC连接到远程节点和授权的交易者
	conn, err := rpc.Server{}
}