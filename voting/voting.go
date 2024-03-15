// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package voting

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VotingMetaData contains all meta data concerning the Voting contract.
var VotingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"color\",\"type\":\"string\"}],\"name\":\"getVoteCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"color\",\"type\":\"string\"}],\"name\":\"voteForColor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506104108061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c8063069953a71461004357806327fa0b7614610073578063b99ef1fa1461008f575b5f80fd5b61005d60048036038101906100589190610282565b6100bf565b60405161006a91906102e1565b60405180910390f35b61008d60048036038101906100889190610282565b6100e5565b005b6100a960048036038101906100a49190610282565b610109565b6040516100b691906102e1565b60405180910390f35b5f80826040516100cf9190610366565b9081526020016040518091039020549050919050565b60045f6040516100f4906103c6565b90815260200160405180910390208190555050565b5f818051602081018201805184825260208301602085012081835280955050505050505f915090505481565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101948261014e565b810181811067ffffffffffffffff821117156101b3576101b261015e565b5b80604052505050565b5f6101c5610135565b90506101d1828261018b565b919050565b5f67ffffffffffffffff8211156101f0576101ef61015e565b5b6101f98261014e565b9050602081019050919050565b828183375f83830152505050565b5f610226610221846101d6565b6101bc565b9050828152602081018484840111156102425761024161014a565b5b61024d848285610206565b509392505050565b5f82601f83011261026957610268610146565b5b8135610279848260208601610214565b91505092915050565b5f602082840312156102975761029661013e565b5b5f82013567ffffffffffffffff8111156102b4576102b3610142565b5b6102c084828501610255565b91505092915050565b5f819050919050565b6102db816102c9565b82525050565b5f6020820190506102f45f8301846102d2565b92915050565b5f81519050919050565b5f81905092915050565b5f5b8381101561032b578082015181840152602081019050610310565b5f8484015250505050565b5f610340826102fa565b61034a8185610304565b935061035a81856020860161030e565b80840191505092915050565b5f6103718284610336565b915081905092915050565b7f72656400000000000000000000000000000000000000000000000000000000005f82015250565b5f6103b0600383610304565b91506103bb8261037c565b600382019050919050565b5f6103d0826103a4565b915081905091905056fea2646970667358221220694c6180e1a5a983509b2e69059d625cc742238a83296ad5c7a4691957c702ee64736f6c63430008180033",
}

// VotingABI is the input ABI used to generate the binding from.
// Deprecated: Use VotingMetaData.ABI instead.
var VotingABI = VotingMetaData.ABI

// VotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VotingMetaData.Bin instead.
var VotingBin = VotingMetaData.Bin

// DeployVoting deploys a new Ethereum contract, binding an instance of Voting to it.
func DeployVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Voting, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// Voting is an auto generated Go binding around an Ethereum contract.
type Voting struct {
	VotingCaller     // Read-only binding to the contract
	VotingTransactor // Write-only binding to the contract
	VotingFilterer   // Log filterer for contract events
}

// VotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type VotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VotingSession struct {
	Contract     *Voting           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VotingCallerSession struct {
	Contract *VotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VotingTransactorSession struct {
	Contract     *VotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type VotingRaw struct {
	Contract *Voting // Generic contract binding to access the raw methods on
}

// VotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VotingCallerRaw struct {
	Contract *VotingCaller // Generic read-only contract binding to access the raw methods on
}

// VotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VotingTransactorRaw struct {
	Contract *VotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoting creates a new instance of Voting, bound to a specific deployed contract.
func NewVoting(address common.Address, backend bind.ContractBackend) (*Voting, error) {
	contract, err := bindVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Voting{VotingCaller: VotingCaller{contract: contract}, VotingTransactor: VotingTransactor{contract: contract}, VotingFilterer: VotingFilterer{contract: contract}}, nil
}

// NewVotingCaller creates a new read-only instance of Voting, bound to a specific deployed contract.
func NewVotingCaller(address common.Address, caller bind.ContractCaller) (*VotingCaller, error) {
	contract, err := bindVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VotingCaller{contract: contract}, nil
}

// NewVotingTransactor creates a new write-only instance of Voting, bound to a specific deployed contract.
func NewVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*VotingTransactor, error) {
	contract, err := bindVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VotingTransactor{contract: contract}, nil
}

// NewVotingFilterer creates a new log filterer instance of Voting, bound to a specific deployed contract.
func NewVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*VotingFilterer, error) {
	contract, err := bindVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VotingFilterer{contract: contract}, nil
}

// bindVoting binds a generic wrapper to an already deployed contract.
func bindVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.VotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.VotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Voting *VotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Voting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Voting *VotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Voting *VotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Voting.Contract.contract.Transact(opts, method, params...)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(string color) view returns(uint256)
func (_Voting *VotingCaller) GetVoteCount(opts *bind.CallOpts, color string) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "getVoteCount", color)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(string color) view returns(uint256)
func (_Voting *VotingSession) GetVoteCount(color string) (*big.Int, error) {
	return _Voting.Contract.GetVoteCount(&_Voting.CallOpts, color)
}

// GetVoteCount is a free data retrieval call binding the contract method 0x069953a7.
//
// Solidity: function getVoteCount(string color) view returns(uint256)
func (_Voting *VotingCallerSession) GetVoteCount(color string) (*big.Int, error) {
	return _Voting.Contract.GetVoteCount(&_Voting.CallOpts, color)
}

// Votes is a free data retrieval call binding the contract method 0xb99ef1fa.
//
// Solidity: function votes(string ) view returns(uint256)
func (_Voting *VotingCaller) Votes(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Voting.contract.Call(opts, &out, "votes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xb99ef1fa.
//
// Solidity: function votes(string ) view returns(uint256)
func (_Voting *VotingSession) Votes(arg0 string) (*big.Int, error) {
	return _Voting.Contract.Votes(&_Voting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xb99ef1fa.
//
// Solidity: function votes(string ) view returns(uint256)
func (_Voting *VotingCallerSession) Votes(arg0 string) (*big.Int, error) {
	return _Voting.Contract.Votes(&_Voting.CallOpts, arg0)
}

// VoteForColor is a paid mutator transaction binding the contract method 0x27fa0b76.
//
// Solidity: function voteForColor(string color) returns()
func (_Voting *VotingTransactor) VoteForColor(opts *bind.TransactOpts, color string) (*types.Transaction, error) {
	return _Voting.contract.Transact(opts, "voteForColor", color)
}

// VoteForColor is a paid mutator transaction binding the contract method 0x27fa0b76.
//
// Solidity: function voteForColor(string color) returns()
func (_Voting *VotingSession) VoteForColor(color string) (*types.Transaction, error) {
	return _Voting.Contract.VoteForColor(&_Voting.TransactOpts, color)
}

// VoteForColor is a paid mutator transaction binding the contract method 0x27fa0b76.
//
// Solidity: function voteForColor(string color) returns()
func (_Voting *VotingTransactorSession) VoteForColor(color string) (*types.Transaction, error) {
	return _Voting.Contract.VoteForColor(&_Voting.TransactOpts, color)
}
