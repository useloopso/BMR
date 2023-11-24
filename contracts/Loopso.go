// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ILoopsoTokenAttestation is an auto generated low-level Go binding around an user-defined struct.
type ILoopsoTokenAttestation struct {
	TokenAddress        common.Address
	TokenChain          *big.Int
	TokenType           uint8
	Decimals            uint8
	Symbol              string
	Name                string
	WrappedTokenAddress common.Address
}

// ILoopsoTokenTransferBase is an auto generated low-level Go binding around an user-defined struct.
type ILoopsoTokenTransferBase struct {
	Timestamp    *big.Int
	SrcChain     *big.Int
	SrcAddress   common.Address
	DstChain     *big.Int
	DstAddress   common.Address
	TokenAddress common.Address
}

// LoopsoABI is the input ABI used to generate the binding from.
const LoopsoABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"NativeTokensReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"}],\"name\":\"NonFungibleTokensBridgedBack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NonFungibleTokensReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"}],\"name\":\"TokenAttested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transferID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"enumILoopso.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"name\":\"TokensBridged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"}],\"name\":\"TokensBridgedBack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokensReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"}],\"name\":\"WrappedNonFungibleTokensReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationID\",\"type\":\"bytes32\"}],\"name\":\"WrappedTokensReleased\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_FUNGIBLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_NON_FUNGIBLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenChain\",\"type\":\"uint256\"},{\"internalType\":\"enumILoopso.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"}],\"internalType\":\"structILoopso.TokenAttestation\",\"name\":\"attestation\",\"type\":\"tuple\"}],\"name\":\"attestToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"attestationIds\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"attestedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenChain\",\"type\":\"uint256\"},{\"internalType\":\"enumILoopso.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"}],\"name\":\"bridgeNativeTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_dstChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"}],\"name\":\"bridgeNonFungibleTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_attestationID\",\"type\":\"bytes32\"}],\"name\":\"bridgeNonFungibleTokensBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"}],\"name\":\"bridgeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_attestationID\",\"type\":\"bytes32\"}],\"name\":\"bridgeTokensBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"discountNft\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSupportedTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenChain\",\"type\":\"uint256\"},{\"internalType\":\"enumILoopso.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"}],\"internalType\":\"structILoopso.TokenAttestation[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenChain\",\"type\":\"uint256\"}],\"name\":\"isTokenSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isWrappedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"releaseNativeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"releaseNonFungibleTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"releaseTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_attestationID\",\"type\":\"bytes32\"}],\"name\":\"releaseWrappedNonFungibleTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_attestationID\",\"type\":\"bytes32\"}],\"name\":\"releaseWrappedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_discountNft\",\"type\":\"address\"}],\"name\":\"setDiscountNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFungibleFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setNonFungibleFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITokenFactory\",\"name\":\"_tokenFactory\",\"type\":\"address\"}],\"name\":\"setTokenFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"contractITokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenTransfers\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dstAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structILoopso.TokenTransferBase\",\"name\":\"tokenTransfer\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenTransfersNonFungible\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"srcChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstChain\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dstAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structILoopso.TokenTransferBase\",\"name\":\"tokenTransfer\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wrappedToken\",\"type\":\"address\"}],\"name\":\"wrappedTokenInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenChain\",\"type\":\"uint256\"},{\"internalType\":\"enumILoopso.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"wrappedTokenAddress\",\"type\":\"address\"}],\"internalType\":\"structILoopso.TokenAttestation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Loopso is an auto generated Go binding around an Ethereum contract.
type Loopso struct {
	LoopsoCaller     // Read-only binding to the contract
	LoopsoTransactor // Write-only binding to the contract
	LoopsoFilterer   // Log filterer for contract events
}

// LoopsoCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoopsoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoopsoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoopsoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoopsoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoopsoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoopsoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoopsoSession struct {
	Contract     *Loopso           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoopsoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoopsoCallerSession struct {
	Contract *LoopsoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LoopsoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoopsoTransactorSession struct {
	Contract     *LoopsoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoopsoRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoopsoRaw struct {
	Contract *Loopso // Generic contract binding to access the raw methods on
}

// LoopsoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoopsoCallerRaw struct {
	Contract *LoopsoCaller // Generic read-only contract binding to access the raw methods on
}

// LoopsoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoopsoTransactorRaw struct {
	Contract *LoopsoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoopso creates a new instance of Loopso, bound to a specific deployed contract.
func NewLoopso(address common.Address, backend bind.ContractBackend) (*Loopso, error) {
	contract, err := bindLoopso(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Loopso{LoopsoCaller: LoopsoCaller{contract: contract}, LoopsoTransactor: LoopsoTransactor{contract: contract}, LoopsoFilterer: LoopsoFilterer{contract: contract}}, nil
}

// NewLoopsoCaller creates a new read-only instance of Loopso, bound to a specific deployed contract.
func NewLoopsoCaller(address common.Address, caller bind.ContractCaller) (*LoopsoCaller, error) {
	contract, err := bindLoopso(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoopsoCaller{contract: contract}, nil
}

// NewLoopsoTransactor creates a new write-only instance of Loopso, bound to a specific deployed contract.
func NewLoopsoTransactor(address common.Address, transactor bind.ContractTransactor) (*LoopsoTransactor, error) {
	contract, err := bindLoopso(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoopsoTransactor{contract: contract}, nil
}

// NewLoopsoFilterer creates a new log filterer instance of Loopso, bound to a specific deployed contract.
func NewLoopsoFilterer(address common.Address, filterer bind.ContractFilterer) (*LoopsoFilterer, error) {
	contract, err := bindLoopso(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoopsoFilterer{contract: contract}, nil
}

// bindLoopso binds a generic wrapper to an already deployed contract.
func bindLoopso(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoopsoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loopso *LoopsoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loopso.Contract.LoopsoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loopso *LoopsoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loopso.Contract.LoopsoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loopso *LoopsoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loopso.Contract.LoopsoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loopso *LoopsoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loopso.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loopso *LoopsoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loopso.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loopso *LoopsoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loopso.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Loopso *LoopsoCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Loopso *LoopsoSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Loopso.Contract.DEFAULTADMINROLE(&_Loopso.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Loopso *LoopsoCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Loopso.Contract.DEFAULTADMINROLE(&_Loopso.CallOpts)
}

// FEEFUNGIBLE is a free data retrieval call binding the contract method 0xe6be636f.
//
// Solidity: function FEE_FUNGIBLE() view returns(uint256)
func (_Loopso *LoopsoCaller) FEEFUNGIBLE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "FEE_FUNGIBLE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEFUNGIBLE is a free data retrieval call binding the contract method 0xe6be636f.
//
// Solidity: function FEE_FUNGIBLE() view returns(uint256)
func (_Loopso *LoopsoSession) FEEFUNGIBLE() (*big.Int, error) {
	return _Loopso.Contract.FEEFUNGIBLE(&_Loopso.CallOpts)
}

// FEEFUNGIBLE is a free data retrieval call binding the contract method 0xe6be636f.
//
// Solidity: function FEE_FUNGIBLE() view returns(uint256)
func (_Loopso *LoopsoCallerSession) FEEFUNGIBLE() (*big.Int, error) {
	return _Loopso.Contract.FEEFUNGIBLE(&_Loopso.CallOpts)
}

// FEENONFUNGIBLE is a free data retrieval call binding the contract method 0xd5abe9c3.
//
// Solidity: function FEE_NON_FUNGIBLE() view returns(uint256)
func (_Loopso *LoopsoCaller) FEENONFUNGIBLE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "FEE_NON_FUNGIBLE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEENONFUNGIBLE is a free data retrieval call binding the contract method 0xd5abe9c3.
//
// Solidity: function FEE_NON_FUNGIBLE() view returns(uint256)
func (_Loopso *LoopsoSession) FEENONFUNGIBLE() (*big.Int, error) {
	return _Loopso.Contract.FEENONFUNGIBLE(&_Loopso.CallOpts)
}

// FEENONFUNGIBLE is a free data retrieval call binding the contract method 0xd5abe9c3.
//
// Solidity: function FEE_NON_FUNGIBLE() view returns(uint256)
func (_Loopso *LoopsoCallerSession) FEENONFUNGIBLE() (*big.Int, error) {
	return _Loopso.Contract.FEENONFUNGIBLE(&_Loopso.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Loopso *LoopsoCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Loopso *LoopsoSession) RELAYERROLE() ([32]byte, error) {
	return _Loopso.Contract.RELAYERROLE(&_Loopso.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Loopso *LoopsoCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Loopso.Contract.RELAYERROLE(&_Loopso.CallOpts)
}

// AttestationIds is a free data retrieval call binding the contract method 0xb0aaf03f.
//
// Solidity: function attestationIds(uint256 ) view returns(bytes32)
func (_Loopso *LoopsoCaller) AttestationIds(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "attestationIds", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AttestationIds is a free data retrieval call binding the contract method 0xb0aaf03f.
//
// Solidity: function attestationIds(uint256 ) view returns(bytes32)
func (_Loopso *LoopsoSession) AttestationIds(arg0 *big.Int) ([32]byte, error) {
	return _Loopso.Contract.AttestationIds(&_Loopso.CallOpts, arg0)
}

// AttestationIds is a free data retrieval call binding the contract method 0xb0aaf03f.
//
// Solidity: function attestationIds(uint256 ) view returns(bytes32)
func (_Loopso *LoopsoCallerSession) AttestationIds(arg0 *big.Int) ([32]byte, error) {
	return _Loopso.Contract.AttestationIds(&_Loopso.CallOpts, arg0)
}

// AttestedTokens is a free data retrieval call binding the contract method 0x7f7b0109.
//
// Solidity: function attestedTokens(bytes32 ) view returns(address tokenAddress, uint256 tokenChain, uint8 tokenType, uint8 decimals, string symbol, string name, address wrappedTokenAddress)
func (_Loopso *LoopsoCaller) AttestedTokens(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TokenAddress        common.Address
	TokenChain          *big.Int
	TokenType           uint8
	Decimals            uint8
	Symbol              string
	Name                string
	WrappedTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "attestedTokens", arg0)

	outstruct := new(struct {
		TokenAddress        common.Address
		TokenChain          *big.Int
		TokenType           uint8
		Decimals            uint8
		Symbol              string
		Name                string
		WrappedTokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenChain = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TokenType = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Decimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Symbol = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.WrappedTokenAddress = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AttestedTokens is a free data retrieval call binding the contract method 0x7f7b0109.
//
// Solidity: function attestedTokens(bytes32 ) view returns(address tokenAddress, uint256 tokenChain, uint8 tokenType, uint8 decimals, string symbol, string name, address wrappedTokenAddress)
func (_Loopso *LoopsoSession) AttestedTokens(arg0 [32]byte) (struct {
	TokenAddress        common.Address
	TokenChain          *big.Int
	TokenType           uint8
	Decimals            uint8
	Symbol              string
	Name                string
	WrappedTokenAddress common.Address
}, error) {
	return _Loopso.Contract.AttestedTokens(&_Loopso.CallOpts, arg0)
}

// AttestedTokens is a free data retrieval call binding the contract method 0x7f7b0109.
//
// Solidity: function attestedTokens(bytes32 ) view returns(address tokenAddress, uint256 tokenChain, uint8 tokenType, uint8 decimals, string symbol, string name, address wrappedTokenAddress)
func (_Loopso *LoopsoCallerSession) AttestedTokens(arg0 [32]byte) (struct {
	TokenAddress        common.Address
	TokenChain          *big.Int
	TokenType           uint8
	Decimals            uint8
	Symbol              string
	Name                string
	WrappedTokenAddress common.Address
}, error) {
	return _Loopso.Contract.AttestedTokens(&_Loopso.CallOpts, arg0)
}

// DiscountNft is a free data retrieval call binding the contract method 0x647920d2.
//
// Solidity: function discountNft() view returns(address)
func (_Loopso *LoopsoCaller) DiscountNft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "discountNft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DiscountNft is a free data retrieval call binding the contract method 0x647920d2.
//
// Solidity: function discountNft() view returns(address)
func (_Loopso *LoopsoSession) DiscountNft() (common.Address, error) {
	return _Loopso.Contract.DiscountNft(&_Loopso.CallOpts)
}

// DiscountNft is a free data retrieval call binding the contract method 0x647920d2.
//
// Solidity: function discountNft() view returns(address)
func (_Loopso *LoopsoCallerSession) DiscountNft() (common.Address, error) {
	return _Loopso.Contract.DiscountNft(&_Loopso.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Loopso *LoopsoCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Loopso *LoopsoSession) FeeReceiver() (common.Address, error) {
	return _Loopso.Contract.FeeReceiver(&_Loopso.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_Loopso *LoopsoCallerSession) FeeReceiver() (common.Address, error) {
	return _Loopso.Contract.FeeReceiver(&_Loopso.CallOpts)
}

// GetAllSupportedTokens is a free data retrieval call binding the contract method 0x0107e472.
//
// Solidity: function getAllSupportedTokens() view returns((address,uint256,uint8,uint8,string,string,address)[])
func (_Loopso *LoopsoCaller) GetAllSupportedTokens(opts *bind.CallOpts) ([]ILoopsoTokenAttestation, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "getAllSupportedTokens")

	if err != nil {
		return *new([]ILoopsoTokenAttestation), err
	}

	out0 := *abi.ConvertType(out[0], new([]ILoopsoTokenAttestation)).(*[]ILoopsoTokenAttestation)

	return out0, err

}

// GetAllSupportedTokens is a free data retrieval call binding the contract method 0x0107e472.
//
// Solidity: function getAllSupportedTokens() view returns((address,uint256,uint8,uint8,string,string,address)[])
func (_Loopso *LoopsoSession) GetAllSupportedTokens() ([]ILoopsoTokenAttestation, error) {
	return _Loopso.Contract.GetAllSupportedTokens(&_Loopso.CallOpts)
}

// GetAllSupportedTokens is a free data retrieval call binding the contract method 0x0107e472.
//
// Solidity: function getAllSupportedTokens() view returns((address,uint256,uint8,uint8,string,string,address)[])
func (_Loopso *LoopsoCallerSession) GetAllSupportedTokens() ([]ILoopsoTokenAttestation, error) {
	return _Loopso.Contract.GetAllSupportedTokens(&_Loopso.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Loopso *LoopsoCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Loopso *LoopsoSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Loopso.Contract.GetRoleAdmin(&_Loopso.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Loopso *LoopsoCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Loopso.Contract.GetRoleAdmin(&_Loopso.CallOpts, role)
}

// GetSupportedTokensLength is a free data retrieval call binding the contract method 0xf960cd6b.
//
// Solidity: function getSupportedTokensLength() view returns(uint256)
func (_Loopso *LoopsoCaller) GetSupportedTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "getSupportedTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSupportedTokensLength is a free data retrieval call binding the contract method 0xf960cd6b.
//
// Solidity: function getSupportedTokensLength() view returns(uint256)
func (_Loopso *LoopsoSession) GetSupportedTokensLength() (*big.Int, error) {
	return _Loopso.Contract.GetSupportedTokensLength(&_Loopso.CallOpts)
}

// GetSupportedTokensLength is a free data retrieval call binding the contract method 0xf960cd6b.
//
// Solidity: function getSupportedTokensLength() view returns(uint256)
func (_Loopso *LoopsoCallerSession) GetSupportedTokensLength() (*big.Int, error) {
	return _Loopso.Contract.GetSupportedTokensLength(&_Loopso.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Loopso *LoopsoCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Loopso *LoopsoSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Loopso.Contract.HasRole(&_Loopso.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Loopso *LoopsoCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Loopso.Contract.HasRole(&_Loopso.CallOpts, role, account)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x832da333.
//
// Solidity: function isTokenSupported(address _tokenAddress, uint256 _tokenChain) view returns(bool)
func (_Loopso *LoopsoCaller) IsTokenSupported(opts *bind.CallOpts, _tokenAddress common.Address, _tokenChain *big.Int) (bool, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "isTokenSupported", _tokenAddress, _tokenChain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenSupported is a free data retrieval call binding the contract method 0x832da333.
//
// Solidity: function isTokenSupported(address _tokenAddress, uint256 _tokenChain) view returns(bool)
func (_Loopso *LoopsoSession) IsTokenSupported(_tokenAddress common.Address, _tokenChain *big.Int) (bool, error) {
	return _Loopso.Contract.IsTokenSupported(&_Loopso.CallOpts, _tokenAddress, _tokenChain)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x832da333.
//
// Solidity: function isTokenSupported(address _tokenAddress, uint256 _tokenChain) view returns(bool)
func (_Loopso *LoopsoCallerSession) IsTokenSupported(_tokenAddress common.Address, _tokenChain *big.Int) (bool, error) {
	return _Loopso.Contract.IsTokenSupported(&_Loopso.CallOpts, _tokenAddress, _tokenChain)
}

// IsWrappedToken is a free data retrieval call binding the contract method 0x64fb065b.
//
// Solidity: function isWrappedToken(address _token) view returns(bool)
func (_Loopso *LoopsoCaller) IsWrappedToken(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "isWrappedToken", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWrappedToken is a free data retrieval call binding the contract method 0x64fb065b.
//
// Solidity: function isWrappedToken(address _token) view returns(bool)
func (_Loopso *LoopsoSession) IsWrappedToken(_token common.Address) (bool, error) {
	return _Loopso.Contract.IsWrappedToken(&_Loopso.CallOpts, _token)
}

// IsWrappedToken is a free data retrieval call binding the contract method 0x64fb065b.
//
// Solidity: function isWrappedToken(address _token) view returns(bool)
func (_Loopso *LoopsoCallerSession) IsWrappedToken(_token common.Address) (bool, error) {
	return _Loopso.Contract.IsWrappedToken(&_Loopso.CallOpts, _token)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Loopso *LoopsoCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Loopso *LoopsoSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Loopso.Contract.OnERC721Received(&_Loopso.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Loopso *LoopsoCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Loopso.Contract.OnERC721Received(&_Loopso.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Loopso *LoopsoCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Loopso *LoopsoSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Loopso.Contract.SupportsInterface(&_Loopso.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Loopso *LoopsoCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Loopso.Contract.SupportsInterface(&_Loopso.CallOpts, interfaceId)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Loopso *LoopsoCaller) TokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "tokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Loopso *LoopsoSession) TokenFactory() (common.Address, error) {
	return _Loopso.Contract.TokenFactory(&_Loopso.CallOpts)
}

// TokenFactory is a free data retrieval call binding the contract method 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (_Loopso *LoopsoCallerSession) TokenFactory() (common.Address, error) {
	return _Loopso.Contract.TokenFactory(&_Loopso.CallOpts)
}

// TokenTransfers is a free data retrieval call binding the contract method 0x1385c670.
//
// Solidity: function tokenTransfers(bytes32 ) view returns((uint256,uint256,address,uint256,address,address) tokenTransfer, uint256 amount)
func (_Loopso *LoopsoCaller) TokenTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TokenTransfer ILoopsoTokenTransferBase
	Amount        *big.Int
}, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "tokenTransfers", arg0)

	outstruct := new(struct {
		TokenTransfer ILoopsoTokenTransferBase
		Amount        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenTransfer = *abi.ConvertType(out[0], new(ILoopsoTokenTransferBase)).(*ILoopsoTokenTransferBase)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenTransfers is a free data retrieval call binding the contract method 0x1385c670.
//
// Solidity: function tokenTransfers(bytes32 ) view returns((uint256,uint256,address,uint256,address,address) tokenTransfer, uint256 amount)
func (_Loopso *LoopsoSession) TokenTransfers(arg0 [32]byte) (struct {
	TokenTransfer ILoopsoTokenTransferBase
	Amount        *big.Int
}, error) {
	return _Loopso.Contract.TokenTransfers(&_Loopso.CallOpts, arg0)
}

// TokenTransfers is a free data retrieval call binding the contract method 0x1385c670.
//
// Solidity: function tokenTransfers(bytes32 ) view returns((uint256,uint256,address,uint256,address,address) tokenTransfer, uint256 amount)
func (_Loopso *LoopsoCallerSession) TokenTransfers(arg0 [32]byte) (struct {
	TokenTransfer ILoopsoTokenTransferBase
	Amount        *big.Int
}, error) {
	return _Loopso.Contract.TokenTransfers(&_Loopso.CallOpts, arg0)
}

// TokenTransfersNonFungible is a free data retrieval call binding the contract method 0xac3b78c8.
//
// Solidity: function tokenTransfersNonFungible(bytes32 ) view returns((uint256,uint256,address,uint256,address,address) tokenTransfer, uint256 tokenID, string tokenURI)
func (_Loopso *LoopsoCaller) TokenTransfersNonFungible(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TokenTransfer ILoopsoTokenTransferBase
	TokenID       *big.Int
	TokenURI      string
}, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "tokenTransfersNonFungible", arg0)

	outstruct := new(struct {
		TokenTransfer ILoopsoTokenTransferBase
		TokenID       *big.Int
		TokenURI      string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenTransfer = *abi.ConvertType(out[0], new(ILoopsoTokenTransferBase)).(*ILoopsoTokenTransferBase)
	outstruct.TokenID = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TokenURI = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// TokenTransfersNonFungible is a free data retrieval call binding the contract method 0xac3b78c8.
//
// Solidity: function tokenTransfersNonFungible(bytes32 ) view returns((uint256,uint256,address,uint256,address,address) tokenTransfer, uint256 tokenID, string tokenURI)
func (_Loopso *LoopsoSession) TokenTransfersNonFungible(arg0 [32]byte) (struct {
	TokenTransfer ILoopsoTokenTransferBase
	TokenID       *big.Int
	TokenURI      string
}, error) {
	return _Loopso.Contract.TokenTransfersNonFungible(&_Loopso.CallOpts, arg0)
}

// TokenTransfersNonFungible is a free data retrieval call binding the contract method 0xac3b78c8.
//
// Solidity: function tokenTransfersNonFungible(bytes32 ) view returns((uint256,uint256,address,uint256,address,address) tokenTransfer, uint256 tokenID, string tokenURI)
func (_Loopso *LoopsoCallerSession) TokenTransfersNonFungible(arg0 [32]byte) (struct {
	TokenTransfer ILoopsoTokenTransferBase
	TokenID       *big.Int
	TokenURI      string
}, error) {
	return _Loopso.Contract.TokenTransfersNonFungible(&_Loopso.CallOpts, arg0)
}

// WrappedTokenInfo is a free data retrieval call binding the contract method 0x3e38ac74.
//
// Solidity: function wrappedTokenInfo(address _wrappedToken) view returns((address,uint256,uint8,uint8,string,string,address))
func (_Loopso *LoopsoCaller) WrappedTokenInfo(opts *bind.CallOpts, _wrappedToken common.Address) (ILoopsoTokenAttestation, error) {
	var out []interface{}
	err := _Loopso.contract.Call(opts, &out, "wrappedTokenInfo", _wrappedToken)

	if err != nil {
		return *new(ILoopsoTokenAttestation), err
	}

	out0 := *abi.ConvertType(out[0], new(ILoopsoTokenAttestation)).(*ILoopsoTokenAttestation)

	return out0, err

}

// WrappedTokenInfo is a free data retrieval call binding the contract method 0x3e38ac74.
//
// Solidity: function wrappedTokenInfo(address _wrappedToken) view returns((address,uint256,uint8,uint8,string,string,address))
func (_Loopso *LoopsoSession) WrappedTokenInfo(_wrappedToken common.Address) (ILoopsoTokenAttestation, error) {
	return _Loopso.Contract.WrappedTokenInfo(&_Loopso.CallOpts, _wrappedToken)
}

// WrappedTokenInfo is a free data retrieval call binding the contract method 0x3e38ac74.
//
// Solidity: function wrappedTokenInfo(address _wrappedToken) view returns((address,uint256,uint8,uint8,string,string,address))
func (_Loopso *LoopsoCallerSession) WrappedTokenInfo(_wrappedToken common.Address) (ILoopsoTokenAttestation, error) {
	return _Loopso.Contract.WrappedTokenInfo(&_Loopso.CallOpts, _wrappedToken)
}

// AttestToken is a paid mutator transaction binding the contract method 0x31ba1648.
//
// Solidity: function attestToken((address,uint256,uint8,uint8,string,string,address) attestation) returns()
func (_Loopso *LoopsoTransactor) AttestToken(opts *bind.TransactOpts, attestation ILoopsoTokenAttestation) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "attestToken", attestation)
}

// AttestToken is a paid mutator transaction binding the contract method 0x31ba1648.
//
// Solidity: function attestToken((address,uint256,uint8,uint8,string,string,address) attestation) returns()
func (_Loopso *LoopsoSession) AttestToken(attestation ILoopsoTokenAttestation) (*types.Transaction, error) {
	return _Loopso.Contract.AttestToken(&_Loopso.TransactOpts, attestation)
}

// AttestToken is a paid mutator transaction binding the contract method 0x31ba1648.
//
// Solidity: function attestToken((address,uint256,uint8,uint8,string,string,address) attestation) returns()
func (_Loopso *LoopsoTransactorSession) AttestToken(attestation ILoopsoTokenAttestation) (*types.Transaction, error) {
	return _Loopso.Contract.AttestToken(&_Loopso.TransactOpts, attestation)
}

// BridgeNativeTokens is a paid mutator transaction binding the contract method 0x9bf22e81.
//
// Solidity: function bridgeNativeTokens(uint256 _dstChain, address _dstAddress) payable returns()
func (_Loopso *LoopsoTransactor) BridgeNativeTokens(opts *bind.TransactOpts, _dstChain *big.Int, _dstAddress common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "bridgeNativeTokens", _dstChain, _dstAddress)
}

// BridgeNativeTokens is a paid mutator transaction binding the contract method 0x9bf22e81.
//
// Solidity: function bridgeNativeTokens(uint256 _dstChain, address _dstAddress) payable returns()
func (_Loopso *LoopsoSession) BridgeNativeTokens(_dstChain *big.Int, _dstAddress common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeNativeTokens(&_Loopso.TransactOpts, _dstChain, _dstAddress)
}

// BridgeNativeTokens is a paid mutator transaction binding the contract method 0x9bf22e81.
//
// Solidity: function bridgeNativeTokens(uint256 _dstChain, address _dstAddress) payable returns()
func (_Loopso *LoopsoTransactorSession) BridgeNativeTokens(_dstChain *big.Int, _dstAddress common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeNativeTokens(&_Loopso.TransactOpts, _dstChain, _dstAddress)
}

// BridgeNonFungibleTokens is a paid mutator transaction binding the contract method 0x40c555dc.
//
// Solidity: function bridgeNonFungibleTokens(address _token, uint256 _tokenID, string tokenURI, uint256 _dstChain, address _dstAddress) payable returns()
func (_Loopso *LoopsoTransactor) BridgeNonFungibleTokens(opts *bind.TransactOpts, _token common.Address, _tokenID *big.Int, tokenURI string, _dstChain *big.Int, _dstAddress common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "bridgeNonFungibleTokens", _token, _tokenID, tokenURI, _dstChain, _dstAddress)
}

// BridgeNonFungibleTokens is a paid mutator transaction binding the contract method 0x40c555dc.
//
// Solidity: function bridgeNonFungibleTokens(address _token, uint256 _tokenID, string tokenURI, uint256 _dstChain, address _dstAddress) payable returns()
func (_Loopso *LoopsoSession) BridgeNonFungibleTokens(_token common.Address, _tokenID *big.Int, tokenURI string, _dstChain *big.Int, _dstAddress common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeNonFungibleTokens(&_Loopso.TransactOpts, _token, _tokenID, tokenURI, _dstChain, _dstAddress)
}

// BridgeNonFungibleTokens is a paid mutator transaction binding the contract method 0x40c555dc.
//
// Solidity: function bridgeNonFungibleTokens(address _token, uint256 _tokenID, string tokenURI, uint256 _dstChain, address _dstAddress) payable returns()
func (_Loopso *LoopsoTransactorSession) BridgeNonFungibleTokens(_token common.Address, _tokenID *big.Int, tokenURI string, _dstChain *big.Int, _dstAddress common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeNonFungibleTokens(&_Loopso.TransactOpts, _token, _tokenID, tokenURI, _dstChain, _dstAddress)
}

// BridgeNonFungibleTokensBack is a paid mutator transaction binding the contract method 0x41574977.
//
// Solidity: function bridgeNonFungibleTokensBack(uint256 _tokenId, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoTransactor) BridgeNonFungibleTokensBack(opts *bind.TransactOpts, _tokenId *big.Int, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "bridgeNonFungibleTokensBack", _tokenId, _to, _attestationID)
}

// BridgeNonFungibleTokensBack is a paid mutator transaction binding the contract method 0x41574977.
//
// Solidity: function bridgeNonFungibleTokensBack(uint256 _tokenId, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoSession) BridgeNonFungibleTokensBack(_tokenId *big.Int, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeNonFungibleTokensBack(&_Loopso.TransactOpts, _tokenId, _to, _attestationID)
}

// BridgeNonFungibleTokensBack is a paid mutator transaction binding the contract method 0x41574977.
//
// Solidity: function bridgeNonFungibleTokensBack(uint256 _tokenId, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoTransactorSession) BridgeNonFungibleTokensBack(_tokenId *big.Int, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeNonFungibleTokensBack(&_Loopso.TransactOpts, _tokenId, _to, _attestationID)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0x33b15aad.
//
// Solidity: function bridgeTokens(address _token, uint256 _amount, uint256 _dstChain, address _dstAddress) returns()
func (_Loopso *LoopsoTransactor) BridgeTokens(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _dstChain *big.Int, _dstAddress common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "bridgeTokens", _token, _amount, _dstChain, _dstAddress)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0x33b15aad.
//
// Solidity: function bridgeTokens(address _token, uint256 _amount, uint256 _dstChain, address _dstAddress) returns()
func (_Loopso *LoopsoSession) BridgeTokens(_token common.Address, _amount *big.Int, _dstChain *big.Int, _dstAddress common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeTokens(&_Loopso.TransactOpts, _token, _amount, _dstChain, _dstAddress)
}

// BridgeTokens is a paid mutator transaction binding the contract method 0x33b15aad.
//
// Solidity: function bridgeTokens(address _token, uint256 _amount, uint256 _dstChain, address _dstAddress) returns()
func (_Loopso *LoopsoTransactorSession) BridgeTokens(_token common.Address, _amount *big.Int, _dstChain *big.Int, _dstAddress common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeTokens(&_Loopso.TransactOpts, _token, _amount, _dstChain, _dstAddress)
}

// BridgeTokensBack is a paid mutator transaction binding the contract method 0x63a4e95f.
//
// Solidity: function bridgeTokensBack(uint256 _amount, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoTransactor) BridgeTokensBack(opts *bind.TransactOpts, _amount *big.Int, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "bridgeTokensBack", _amount, _to, _attestationID)
}

// BridgeTokensBack is a paid mutator transaction binding the contract method 0x63a4e95f.
//
// Solidity: function bridgeTokensBack(uint256 _amount, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoSession) BridgeTokensBack(_amount *big.Int, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeTokensBack(&_Loopso.TransactOpts, _amount, _to, _attestationID)
}

// BridgeTokensBack is a paid mutator transaction binding the contract method 0x63a4e95f.
//
// Solidity: function bridgeTokensBack(uint256 _amount, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoTransactorSession) BridgeTokensBack(_amount *big.Int, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.Contract.BridgeTokensBack(&_Loopso.TransactOpts, _amount, _to, _attestationID)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Loopso *LoopsoTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Loopso *LoopsoSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.GrantRole(&_Loopso.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Loopso *LoopsoTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.GrantRole(&_Loopso.TransactOpts, role, account)
}

// ReleaseNativeTokens is a paid mutator transaction binding the contract method 0x6cbd1c75.
//
// Solidity: function releaseNativeTokens(uint256 _amount, address _to) returns()
func (_Loopso *LoopsoTransactor) ReleaseNativeTokens(opts *bind.TransactOpts, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "releaseNativeTokens", _amount, _to)
}

// ReleaseNativeTokens is a paid mutator transaction binding the contract method 0x6cbd1c75.
//
// Solidity: function releaseNativeTokens(uint256 _amount, address _to) returns()
func (_Loopso *LoopsoSession) ReleaseNativeTokens(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseNativeTokens(&_Loopso.TransactOpts, _amount, _to)
}

// ReleaseNativeTokens is a paid mutator transaction binding the contract method 0x6cbd1c75.
//
// Solidity: function releaseNativeTokens(uint256 _amount, address _to) returns()
func (_Loopso *LoopsoTransactorSession) ReleaseNativeTokens(_amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseNativeTokens(&_Loopso.TransactOpts, _amount, _to)
}

// ReleaseNonFungibleTokens is a paid mutator transaction binding the contract method 0x15f5a8c5.
//
// Solidity: function releaseNonFungibleTokens(uint256 _tokenId, address _to, address _token) returns()
func (_Loopso *LoopsoTransactor) ReleaseNonFungibleTokens(opts *bind.TransactOpts, _tokenId *big.Int, _to common.Address, _token common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "releaseNonFungibleTokens", _tokenId, _to, _token)
}

// ReleaseNonFungibleTokens is a paid mutator transaction binding the contract method 0x15f5a8c5.
//
// Solidity: function releaseNonFungibleTokens(uint256 _tokenId, address _to, address _token) returns()
func (_Loopso *LoopsoSession) ReleaseNonFungibleTokens(_tokenId *big.Int, _to common.Address, _token common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseNonFungibleTokens(&_Loopso.TransactOpts, _tokenId, _to, _token)
}

// ReleaseNonFungibleTokens is a paid mutator transaction binding the contract method 0x15f5a8c5.
//
// Solidity: function releaseNonFungibleTokens(uint256 _tokenId, address _to, address _token) returns()
func (_Loopso *LoopsoTransactorSession) ReleaseNonFungibleTokens(_tokenId *big.Int, _to common.Address, _token common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseNonFungibleTokens(&_Loopso.TransactOpts, _tokenId, _to, _token)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0x62a00971.
//
// Solidity: function releaseTokens(uint256 _amount, address _to, address _token) returns()
func (_Loopso *LoopsoTransactor) ReleaseTokens(opts *bind.TransactOpts, _amount *big.Int, _to common.Address, _token common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "releaseTokens", _amount, _to, _token)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0x62a00971.
//
// Solidity: function releaseTokens(uint256 _amount, address _to, address _token) returns()
func (_Loopso *LoopsoSession) ReleaseTokens(_amount *big.Int, _to common.Address, _token common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseTokens(&_Loopso.TransactOpts, _amount, _to, _token)
}

// ReleaseTokens is a paid mutator transaction binding the contract method 0x62a00971.
//
// Solidity: function releaseTokens(uint256 _amount, address _to, address _token) returns()
func (_Loopso *LoopsoTransactorSession) ReleaseTokens(_amount *big.Int, _to common.Address, _token common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseTokens(&_Loopso.TransactOpts, _amount, _to, _token)
}

// ReleaseWrappedNonFungibleTokens is a paid mutator transaction binding the contract method 0xc1f865bb.
//
// Solidity: function releaseWrappedNonFungibleTokens(uint256 _tokenId, string _tokenURI, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoTransactor) ReleaseWrappedNonFungibleTokens(opts *bind.TransactOpts, _tokenId *big.Int, _tokenURI string, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "releaseWrappedNonFungibleTokens", _tokenId, _tokenURI, _to, _attestationID)
}

// ReleaseWrappedNonFungibleTokens is a paid mutator transaction binding the contract method 0xc1f865bb.
//
// Solidity: function releaseWrappedNonFungibleTokens(uint256 _tokenId, string _tokenURI, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoSession) ReleaseWrappedNonFungibleTokens(_tokenId *big.Int, _tokenURI string, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseWrappedNonFungibleTokens(&_Loopso.TransactOpts, _tokenId, _tokenURI, _to, _attestationID)
}

// ReleaseWrappedNonFungibleTokens is a paid mutator transaction binding the contract method 0xc1f865bb.
//
// Solidity: function releaseWrappedNonFungibleTokens(uint256 _tokenId, string _tokenURI, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoTransactorSession) ReleaseWrappedNonFungibleTokens(_tokenId *big.Int, _tokenURI string, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseWrappedNonFungibleTokens(&_Loopso.TransactOpts, _tokenId, _tokenURI, _to, _attestationID)
}

// ReleaseWrappedTokens is a paid mutator transaction binding the contract method 0x9a67e7f9.
//
// Solidity: function releaseWrappedTokens(uint256 _amount, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoTransactor) ReleaseWrappedTokens(opts *bind.TransactOpts, _amount *big.Int, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "releaseWrappedTokens", _amount, _to, _attestationID)
}

// ReleaseWrappedTokens is a paid mutator transaction binding the contract method 0x9a67e7f9.
//
// Solidity: function releaseWrappedTokens(uint256 _amount, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoSession) ReleaseWrappedTokens(_amount *big.Int, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseWrappedTokens(&_Loopso.TransactOpts, _amount, _to, _attestationID)
}

// ReleaseWrappedTokens is a paid mutator transaction binding the contract method 0x9a67e7f9.
//
// Solidity: function releaseWrappedTokens(uint256 _amount, address _to, bytes32 _attestationID) returns()
func (_Loopso *LoopsoTransactorSession) ReleaseWrappedTokens(_amount *big.Int, _to common.Address, _attestationID [32]byte) (*types.Transaction, error) {
	return _Loopso.Contract.ReleaseWrappedTokens(&_Loopso.TransactOpts, _amount, _to, _attestationID)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Loopso *LoopsoTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Loopso *LoopsoSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.RenounceRole(&_Loopso.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Loopso *LoopsoTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.RenounceRole(&_Loopso.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Loopso *LoopsoTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Loopso *LoopsoSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.RevokeRole(&_Loopso.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Loopso *LoopsoTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.RevokeRole(&_Loopso.TransactOpts, role, account)
}

// SetDiscountNft is a paid mutator transaction binding the contract method 0x12705222.
//
// Solidity: function setDiscountNft(address _discountNft) returns()
func (_Loopso *LoopsoTransactor) SetDiscountNft(opts *bind.TransactOpts, _discountNft common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "setDiscountNft", _discountNft)
}

// SetDiscountNft is a paid mutator transaction binding the contract method 0x12705222.
//
// Solidity: function setDiscountNft(address _discountNft) returns()
func (_Loopso *LoopsoSession) SetDiscountNft(_discountNft common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.SetDiscountNft(&_Loopso.TransactOpts, _discountNft)
}

// SetDiscountNft is a paid mutator transaction binding the contract method 0x12705222.
//
// Solidity: function setDiscountNft(address _discountNft) returns()
func (_Loopso *LoopsoTransactorSession) SetDiscountNft(_discountNft common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.SetDiscountNft(&_Loopso.TransactOpts, _discountNft)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Loopso *LoopsoTransactor) SetFeeReceiver(opts *bind.TransactOpts, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "setFeeReceiver", _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Loopso *LoopsoSession) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.SetFeeReceiver(&_Loopso.TransactOpts, _feeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address _feeReceiver) returns()
func (_Loopso *LoopsoTransactorSession) SetFeeReceiver(_feeReceiver common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.SetFeeReceiver(&_Loopso.TransactOpts, _feeReceiver)
}

// SetFungibleFee is a paid mutator transaction binding the contract method 0xd0e04bfa.
//
// Solidity: function setFungibleFee(uint256 _fee) returns()
func (_Loopso *LoopsoTransactor) SetFungibleFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "setFungibleFee", _fee)
}

// SetFungibleFee is a paid mutator transaction binding the contract method 0xd0e04bfa.
//
// Solidity: function setFungibleFee(uint256 _fee) returns()
func (_Loopso *LoopsoSession) SetFungibleFee(_fee *big.Int) (*types.Transaction, error) {
	return _Loopso.Contract.SetFungibleFee(&_Loopso.TransactOpts, _fee)
}

// SetFungibleFee is a paid mutator transaction binding the contract method 0xd0e04bfa.
//
// Solidity: function setFungibleFee(uint256 _fee) returns()
func (_Loopso *LoopsoTransactorSession) SetFungibleFee(_fee *big.Int) (*types.Transaction, error) {
	return _Loopso.Contract.SetFungibleFee(&_Loopso.TransactOpts, _fee)
}

// SetNonFungibleFee is a paid mutator transaction binding the contract method 0x4676f366.
//
// Solidity: function setNonFungibleFee(uint256 _fee) returns()
func (_Loopso *LoopsoTransactor) SetNonFungibleFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "setNonFungibleFee", _fee)
}

// SetNonFungibleFee is a paid mutator transaction binding the contract method 0x4676f366.
//
// Solidity: function setNonFungibleFee(uint256 _fee) returns()
func (_Loopso *LoopsoSession) SetNonFungibleFee(_fee *big.Int) (*types.Transaction, error) {
	return _Loopso.Contract.SetNonFungibleFee(&_Loopso.TransactOpts, _fee)
}

// SetNonFungibleFee is a paid mutator transaction binding the contract method 0x4676f366.
//
// Solidity: function setNonFungibleFee(uint256 _fee) returns()
func (_Loopso *LoopsoTransactorSession) SetNonFungibleFee(_fee *big.Int) (*types.Transaction, error) {
	return _Loopso.Contract.SetNonFungibleFee(&_Loopso.TransactOpts, _fee)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _tokenFactory) returns()
func (_Loopso *LoopsoTransactor) SetTokenFactory(opts *bind.TransactOpts, _tokenFactory common.Address) (*types.Transaction, error) {
	return _Loopso.contract.Transact(opts, "setTokenFactory", _tokenFactory)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _tokenFactory) returns()
func (_Loopso *LoopsoSession) SetTokenFactory(_tokenFactory common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.SetTokenFactory(&_Loopso.TransactOpts, _tokenFactory)
}

// SetTokenFactory is a paid mutator transaction binding the contract method 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _tokenFactory) returns()
func (_Loopso *LoopsoTransactorSession) SetTokenFactory(_tokenFactory common.Address) (*types.Transaction, error) {
	return _Loopso.Contract.SetTokenFactory(&_Loopso.TransactOpts, _tokenFactory)
}

// LoopsoNativeTokensReleasedIterator is returned from FilterNativeTokensReleased and is used to iterate over the raw logs and unpacked data for NativeTokensReleased events raised by the Loopso contract.
type LoopsoNativeTokensReleasedIterator struct {
	Event *LoopsoNativeTokensReleased // Event containing the contract specifics and raw log

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
func (it *LoopsoNativeTokensReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoNativeTokensReleased)
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
		it.Event = new(LoopsoNativeTokensReleased)
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
func (it *LoopsoNativeTokensReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoNativeTokensReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoNativeTokensReleased represents a NativeTokensReleased event raised by the Loopso contract.
type LoopsoNativeTokensReleased struct {
	Amount *big.Int
	To     common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeTokensReleased is a free log retrieval operation binding the contract event 0x59b630c025978c37eb64e54dcb9157eac7455f1cd69ce1989d7329b4ecbf0cbb.
//
// Solidity: event NativeTokensReleased(uint256 indexed amount, address indexed to)
func (_Loopso *LoopsoFilterer) FilterNativeTokensReleased(opts *bind.FilterOpts, amount []*big.Int, to []common.Address) (*LoopsoNativeTokensReleasedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "NativeTokensReleased", amountRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoNativeTokensReleasedIterator{contract: _Loopso.contract, event: "NativeTokensReleased", logs: logs, sub: sub}, nil
}

// WatchNativeTokensReleased is a free log subscription operation binding the contract event 0x59b630c025978c37eb64e54dcb9157eac7455f1cd69ce1989d7329b4ecbf0cbb.
//
// Solidity: event NativeTokensReleased(uint256 indexed amount, address indexed to)
func (_Loopso *LoopsoFilterer) WatchNativeTokensReleased(opts *bind.WatchOpts, sink chan<- *LoopsoNativeTokensReleased, amount []*big.Int, to []common.Address) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "NativeTokensReleased", amountRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoNativeTokensReleased)
				if err := _Loopso.contract.UnpackLog(event, "NativeTokensReleased", log); err != nil {
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

// ParseNativeTokensReleased is a log parse operation binding the contract event 0x59b630c025978c37eb64e54dcb9157eac7455f1cd69ce1989d7329b4ecbf0cbb.
//
// Solidity: event NativeTokensReleased(uint256 indexed amount, address indexed to)
func (_Loopso *LoopsoFilterer) ParseNativeTokensReleased(log types.Log) (*LoopsoNativeTokensReleased, error) {
	event := new(LoopsoNativeTokensReleased)
	if err := _Loopso.contract.UnpackLog(event, "NativeTokensReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoNonFungibleTokensBridgedBackIterator is returned from FilterNonFungibleTokensBridgedBack and is used to iterate over the raw logs and unpacked data for NonFungibleTokensBridgedBack events raised by the Loopso contract.
type LoopsoNonFungibleTokensBridgedBackIterator struct {
	Event *LoopsoNonFungibleTokensBridgedBack // Event containing the contract specifics and raw log

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
func (it *LoopsoNonFungibleTokensBridgedBackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoNonFungibleTokensBridgedBack)
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
		it.Event = new(LoopsoNonFungibleTokensBridgedBack)
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
func (it *LoopsoNonFungibleTokensBridgedBackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoNonFungibleTokensBridgedBackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoNonFungibleTokensBridgedBack represents a NonFungibleTokensBridgedBack event raised by the Loopso contract.
type LoopsoNonFungibleTokensBridgedBack struct {
	TokenId       *big.Int
	To            common.Address
	AttestationID [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNonFungibleTokensBridgedBack is a free log retrieval operation binding the contract event 0xfe42a38b0719dedad7ba864d935eab73d00c2b1709902dcfe410143190915e6d.
//
// Solidity: event NonFungibleTokensBridgedBack(uint256 indexed tokenId, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) FilterNonFungibleTokensBridgedBack(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address, attestationID [][32]byte) (*LoopsoNonFungibleTokensBridgedBackIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "NonFungibleTokensBridgedBack", tokenIdRule, toRule, attestationIDRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoNonFungibleTokensBridgedBackIterator{contract: _Loopso.contract, event: "NonFungibleTokensBridgedBack", logs: logs, sub: sub}, nil
}

// WatchNonFungibleTokensBridgedBack is a free log subscription operation binding the contract event 0xfe42a38b0719dedad7ba864d935eab73d00c2b1709902dcfe410143190915e6d.
//
// Solidity: event NonFungibleTokensBridgedBack(uint256 indexed tokenId, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) WatchNonFungibleTokensBridgedBack(opts *bind.WatchOpts, sink chan<- *LoopsoNonFungibleTokensBridgedBack, tokenId []*big.Int, to []common.Address, attestationID [][32]byte) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "NonFungibleTokensBridgedBack", tokenIdRule, toRule, attestationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoNonFungibleTokensBridgedBack)
				if err := _Loopso.contract.UnpackLog(event, "NonFungibleTokensBridgedBack", log); err != nil {
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

// ParseNonFungibleTokensBridgedBack is a log parse operation binding the contract event 0xfe42a38b0719dedad7ba864d935eab73d00c2b1709902dcfe410143190915e6d.
//
// Solidity: event NonFungibleTokensBridgedBack(uint256 indexed tokenId, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) ParseNonFungibleTokensBridgedBack(log types.Log) (*LoopsoNonFungibleTokensBridgedBack, error) {
	event := new(LoopsoNonFungibleTokensBridgedBack)
	if err := _Loopso.contract.UnpackLog(event, "NonFungibleTokensBridgedBack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoNonFungibleTokensReleasedIterator is returned from FilterNonFungibleTokensReleased and is used to iterate over the raw logs and unpacked data for NonFungibleTokensReleased events raised by the Loopso contract.
type LoopsoNonFungibleTokensReleasedIterator struct {
	Event *LoopsoNonFungibleTokensReleased // Event containing the contract specifics and raw log

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
func (it *LoopsoNonFungibleTokensReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoNonFungibleTokensReleased)
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
		it.Event = new(LoopsoNonFungibleTokensReleased)
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
func (it *LoopsoNonFungibleTokensReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoNonFungibleTokensReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoNonFungibleTokensReleased represents a NonFungibleTokensReleased event raised by the Loopso contract.
type LoopsoNonFungibleTokensReleased struct {
	TokenId *big.Int
	To      common.Address
	Token   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNonFungibleTokensReleased is a free log retrieval operation binding the contract event 0x8c398b928f8bac4ad1cf5448e6c4e602012df28bcc0e4bf2b9e715da4c3b653b.
//
// Solidity: event NonFungibleTokensReleased(uint256 indexed tokenId, address indexed to, address indexed token)
func (_Loopso *LoopsoFilterer) FilterNonFungibleTokensReleased(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address, token []common.Address) (*LoopsoNonFungibleTokensReleasedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "NonFungibleTokensReleased", tokenIdRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoNonFungibleTokensReleasedIterator{contract: _Loopso.contract, event: "NonFungibleTokensReleased", logs: logs, sub: sub}, nil
}

// WatchNonFungibleTokensReleased is a free log subscription operation binding the contract event 0x8c398b928f8bac4ad1cf5448e6c4e602012df28bcc0e4bf2b9e715da4c3b653b.
//
// Solidity: event NonFungibleTokensReleased(uint256 indexed tokenId, address indexed to, address indexed token)
func (_Loopso *LoopsoFilterer) WatchNonFungibleTokensReleased(opts *bind.WatchOpts, sink chan<- *LoopsoNonFungibleTokensReleased, tokenId []*big.Int, to []common.Address, token []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "NonFungibleTokensReleased", tokenIdRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoNonFungibleTokensReleased)
				if err := _Loopso.contract.UnpackLog(event, "NonFungibleTokensReleased", log); err != nil {
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

// ParseNonFungibleTokensReleased is a log parse operation binding the contract event 0x8c398b928f8bac4ad1cf5448e6c4e602012df28bcc0e4bf2b9e715da4c3b653b.
//
// Solidity: event NonFungibleTokensReleased(uint256 indexed tokenId, address indexed to, address indexed token)
func (_Loopso *LoopsoFilterer) ParseNonFungibleTokensReleased(log types.Log) (*LoopsoNonFungibleTokensReleased, error) {
	event := new(LoopsoNonFungibleTokensReleased)
	if err := _Loopso.contract.UnpackLog(event, "NonFungibleTokensReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Loopso contract.
type LoopsoRoleAdminChangedIterator struct {
	Event *LoopsoRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LoopsoRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoRoleAdminChanged)
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
		it.Event = new(LoopsoRoleAdminChanged)
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
func (it *LoopsoRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoRoleAdminChanged represents a RoleAdminChanged event raised by the Loopso contract.
type LoopsoRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Loopso *LoopsoFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LoopsoRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoRoleAdminChangedIterator{contract: _Loopso.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Loopso *LoopsoFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LoopsoRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoRoleAdminChanged)
				if err := _Loopso.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Loopso *LoopsoFilterer) ParseRoleAdminChanged(log types.Log) (*LoopsoRoleAdminChanged, error) {
	event := new(LoopsoRoleAdminChanged)
	if err := _Loopso.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Loopso contract.
type LoopsoRoleGrantedIterator struct {
	Event *LoopsoRoleGranted // Event containing the contract specifics and raw log

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
func (it *LoopsoRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoRoleGranted)
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
		it.Event = new(LoopsoRoleGranted)
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
func (it *LoopsoRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoRoleGranted represents a RoleGranted event raised by the Loopso contract.
type LoopsoRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Loopso *LoopsoFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LoopsoRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoRoleGrantedIterator{contract: _Loopso.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Loopso *LoopsoFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LoopsoRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoRoleGranted)
				if err := _Loopso.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Loopso *LoopsoFilterer) ParseRoleGranted(log types.Log) (*LoopsoRoleGranted, error) {
	event := new(LoopsoRoleGranted)
	if err := _Loopso.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Loopso contract.
type LoopsoRoleRevokedIterator struct {
	Event *LoopsoRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LoopsoRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoRoleRevoked)
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
		it.Event = new(LoopsoRoleRevoked)
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
func (it *LoopsoRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoRoleRevoked represents a RoleRevoked event raised by the Loopso contract.
type LoopsoRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Loopso *LoopsoFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LoopsoRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoRoleRevokedIterator{contract: _Loopso.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Loopso *LoopsoFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LoopsoRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoRoleRevoked)
				if err := _Loopso.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Loopso *LoopsoFilterer) ParseRoleRevoked(log types.Log) (*LoopsoRoleRevoked, error) {
	event := new(LoopsoRoleRevoked)
	if err := _Loopso.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoTokenAttestedIterator is returned from FilterTokenAttested and is used to iterate over the raw logs and unpacked data for TokenAttested events raised by the Loopso contract.
type LoopsoTokenAttestedIterator struct {
	Event *LoopsoTokenAttested // Event containing the contract specifics and raw log

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
func (it *LoopsoTokenAttestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoTokenAttested)
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
		it.Event = new(LoopsoTokenAttested)
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
func (it *LoopsoTokenAttestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoTokenAttestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoTokenAttested represents a TokenAttested event raised by the Loopso contract.
type LoopsoTokenAttested struct {
	AttestationID [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenAttested is a free log retrieval operation binding the contract event 0xa755bee123b19cfd95f30d3bed835ac8dab8a08f893508aa86140d1b66c77e13.
//
// Solidity: event TokenAttested(bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) FilterTokenAttested(opts *bind.FilterOpts, attestationID [][32]byte) (*LoopsoTokenAttestedIterator, error) {

	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "TokenAttested", attestationIDRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoTokenAttestedIterator{contract: _Loopso.contract, event: "TokenAttested", logs: logs, sub: sub}, nil
}

// WatchTokenAttested is a free log subscription operation binding the contract event 0xa755bee123b19cfd95f30d3bed835ac8dab8a08f893508aa86140d1b66c77e13.
//
// Solidity: event TokenAttested(bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) WatchTokenAttested(opts *bind.WatchOpts, sink chan<- *LoopsoTokenAttested, attestationID [][32]byte) (event.Subscription, error) {

	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "TokenAttested", attestationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoTokenAttested)
				if err := _Loopso.contract.UnpackLog(event, "TokenAttested", log); err != nil {
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

// ParseTokenAttested is a log parse operation binding the contract event 0xa755bee123b19cfd95f30d3bed835ac8dab8a08f893508aa86140d1b66c77e13.
//
// Solidity: event TokenAttested(bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) ParseTokenAttested(log types.Log) (*LoopsoTokenAttested, error) {
	event := new(LoopsoTokenAttested)
	if err := _Loopso.contract.UnpackLog(event, "TokenAttested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoTokensBridgedIterator is returned from FilterTokensBridged and is used to iterate over the raw logs and unpacked data for TokensBridged events raised by the Loopso contract.
type LoopsoTokensBridgedIterator struct {
	Event *LoopsoTokensBridged // Event containing the contract specifics and raw log

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
func (it *LoopsoTokensBridgedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoTokensBridged)
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
		it.Event = new(LoopsoTokensBridged)
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
func (it *LoopsoTokensBridgedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoTokensBridgedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoTokensBridged represents a TokensBridged event raised by the Loopso contract.
type LoopsoTokensBridged struct {
	TransferID [32]byte
	TokenType  uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokensBridged is a free log retrieval operation binding the contract event 0x11a762c44e982fe76f4f9b9c028673684f53601407c960c08a6f4472419373e6.
//
// Solidity: event TokensBridged(bytes32 indexed transferID, uint8 indexed tokenType)
func (_Loopso *LoopsoFilterer) FilterTokensBridged(opts *bind.FilterOpts, transferID [][32]byte, tokenType []uint8) (*LoopsoTokensBridgedIterator, error) {

	var transferIDRule []interface{}
	for _, transferIDItem := range transferID {
		transferIDRule = append(transferIDRule, transferIDItem)
	}
	var tokenTypeRule []interface{}
	for _, tokenTypeItem := range tokenType {
		tokenTypeRule = append(tokenTypeRule, tokenTypeItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "TokensBridged", transferIDRule, tokenTypeRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoTokensBridgedIterator{contract: _Loopso.contract, event: "TokensBridged", logs: logs, sub: sub}, nil
}

// WatchTokensBridged is a free log subscription operation binding the contract event 0x11a762c44e982fe76f4f9b9c028673684f53601407c960c08a6f4472419373e6.
//
// Solidity: event TokensBridged(bytes32 indexed transferID, uint8 indexed tokenType)
func (_Loopso *LoopsoFilterer) WatchTokensBridged(opts *bind.WatchOpts, sink chan<- *LoopsoTokensBridged, transferID [][32]byte, tokenType []uint8) (event.Subscription, error) {

	var transferIDRule []interface{}
	for _, transferIDItem := range transferID {
		transferIDRule = append(transferIDRule, transferIDItem)
	}
	var tokenTypeRule []interface{}
	for _, tokenTypeItem := range tokenType {
		tokenTypeRule = append(tokenTypeRule, tokenTypeItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "TokensBridged", transferIDRule, tokenTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoTokensBridged)
				if err := _Loopso.contract.UnpackLog(event, "TokensBridged", log); err != nil {
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

// ParseTokensBridged is a log parse operation binding the contract event 0x11a762c44e982fe76f4f9b9c028673684f53601407c960c08a6f4472419373e6.
//
// Solidity: event TokensBridged(bytes32 indexed transferID, uint8 indexed tokenType)
func (_Loopso *LoopsoFilterer) ParseTokensBridged(log types.Log) (*LoopsoTokensBridged, error) {
	event := new(LoopsoTokensBridged)
	if err := _Loopso.contract.UnpackLog(event, "TokensBridged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoTokensBridgedBackIterator is returned from FilterTokensBridgedBack and is used to iterate over the raw logs and unpacked data for TokensBridgedBack events raised by the Loopso contract.
type LoopsoTokensBridgedBackIterator struct {
	Event *LoopsoTokensBridgedBack // Event containing the contract specifics and raw log

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
func (it *LoopsoTokensBridgedBackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoTokensBridgedBack)
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
		it.Event = new(LoopsoTokensBridgedBack)
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
func (it *LoopsoTokensBridgedBackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoTokensBridgedBackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoTokensBridgedBack represents a TokensBridgedBack event raised by the Loopso contract.
type LoopsoTokensBridgedBack struct {
	Amount        *big.Int
	To            common.Address
	AttestationID [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensBridgedBack is a free log retrieval operation binding the contract event 0x80c5d77a092b389b97d2b4a5bc721315e540c267567756df827f2edab66facba.
//
// Solidity: event TokensBridgedBack(uint256 indexed amount, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) FilterTokensBridgedBack(opts *bind.FilterOpts, amount []*big.Int, to []common.Address, attestationID [][32]byte) (*LoopsoTokensBridgedBackIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "TokensBridgedBack", amountRule, toRule, attestationIDRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoTokensBridgedBackIterator{contract: _Loopso.contract, event: "TokensBridgedBack", logs: logs, sub: sub}, nil
}

// WatchTokensBridgedBack is a free log subscription operation binding the contract event 0x80c5d77a092b389b97d2b4a5bc721315e540c267567756df827f2edab66facba.
//
// Solidity: event TokensBridgedBack(uint256 indexed amount, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) WatchTokensBridgedBack(opts *bind.WatchOpts, sink chan<- *LoopsoTokensBridgedBack, amount []*big.Int, to []common.Address, attestationID [][32]byte) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "TokensBridgedBack", amountRule, toRule, attestationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoTokensBridgedBack)
				if err := _Loopso.contract.UnpackLog(event, "TokensBridgedBack", log); err != nil {
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

// ParseTokensBridgedBack is a log parse operation binding the contract event 0x80c5d77a092b389b97d2b4a5bc721315e540c267567756df827f2edab66facba.
//
// Solidity: event TokensBridgedBack(uint256 indexed amount, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) ParseTokensBridgedBack(log types.Log) (*LoopsoTokensBridgedBack, error) {
	event := new(LoopsoTokensBridgedBack)
	if err := _Loopso.contract.UnpackLog(event, "TokensBridgedBack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoTokensReleasedIterator is returned from FilterTokensReleased and is used to iterate over the raw logs and unpacked data for TokensReleased events raised by the Loopso contract.
type LoopsoTokensReleasedIterator struct {
	Event *LoopsoTokensReleased // Event containing the contract specifics and raw log

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
func (it *LoopsoTokensReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoTokensReleased)
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
		it.Event = new(LoopsoTokensReleased)
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
func (it *LoopsoTokensReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoTokensReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoTokensReleased represents a TokensReleased event raised by the Loopso contract.
type LoopsoTokensReleased struct {
	Amount *big.Int
	To     common.Address
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensReleased is a free log retrieval operation binding the contract event 0xac984453286939de6c31d3b8390199a972705a0d800108d513f2b825d23768fb.
//
// Solidity: event TokensReleased(uint256 indexed amount, address indexed to, address indexed token)
func (_Loopso *LoopsoFilterer) FilterTokensReleased(opts *bind.FilterOpts, amount []*big.Int, to []common.Address, token []common.Address) (*LoopsoTokensReleasedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "TokensReleased", amountRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoTokensReleasedIterator{contract: _Loopso.contract, event: "TokensReleased", logs: logs, sub: sub}, nil
}

// WatchTokensReleased is a free log subscription operation binding the contract event 0xac984453286939de6c31d3b8390199a972705a0d800108d513f2b825d23768fb.
//
// Solidity: event TokensReleased(uint256 indexed amount, address indexed to, address indexed token)
func (_Loopso *LoopsoFilterer) WatchTokensReleased(opts *bind.WatchOpts, sink chan<- *LoopsoTokensReleased, amount []*big.Int, to []common.Address, token []common.Address) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "TokensReleased", amountRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoTokensReleased)
				if err := _Loopso.contract.UnpackLog(event, "TokensReleased", log); err != nil {
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

// ParseTokensReleased is a log parse operation binding the contract event 0xac984453286939de6c31d3b8390199a972705a0d800108d513f2b825d23768fb.
//
// Solidity: event TokensReleased(uint256 indexed amount, address indexed to, address indexed token)
func (_Loopso *LoopsoFilterer) ParseTokensReleased(log types.Log) (*LoopsoTokensReleased, error) {
	event := new(LoopsoTokensReleased)
	if err := _Loopso.contract.UnpackLog(event, "TokensReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoWrappedNonFungibleTokensReleasedIterator is returned from FilterWrappedNonFungibleTokensReleased and is used to iterate over the raw logs and unpacked data for WrappedNonFungibleTokensReleased events raised by the Loopso contract.
type LoopsoWrappedNonFungibleTokensReleasedIterator struct {
	Event *LoopsoWrappedNonFungibleTokensReleased // Event containing the contract specifics and raw log

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
func (it *LoopsoWrappedNonFungibleTokensReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoWrappedNonFungibleTokensReleased)
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
		it.Event = new(LoopsoWrappedNonFungibleTokensReleased)
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
func (it *LoopsoWrappedNonFungibleTokensReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoWrappedNonFungibleTokensReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoWrappedNonFungibleTokensReleased represents a WrappedNonFungibleTokensReleased event raised by the Loopso contract.
type LoopsoWrappedNonFungibleTokensReleased struct {
	TokenId       *big.Int
	To            common.Address
	AttestationID [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWrappedNonFungibleTokensReleased is a free log retrieval operation binding the contract event 0xe436066797d3bcdab647d36495d87e5dba7a58a6c132890edcc038d783a596c1.
//
// Solidity: event WrappedNonFungibleTokensReleased(uint256 indexed tokenId, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) FilterWrappedNonFungibleTokensReleased(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address, attestationID [][32]byte) (*LoopsoWrappedNonFungibleTokensReleasedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "WrappedNonFungibleTokensReleased", tokenIdRule, toRule, attestationIDRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoWrappedNonFungibleTokensReleasedIterator{contract: _Loopso.contract, event: "WrappedNonFungibleTokensReleased", logs: logs, sub: sub}, nil
}

// WatchWrappedNonFungibleTokensReleased is a free log subscription operation binding the contract event 0xe436066797d3bcdab647d36495d87e5dba7a58a6c132890edcc038d783a596c1.
//
// Solidity: event WrappedNonFungibleTokensReleased(uint256 indexed tokenId, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) WatchWrappedNonFungibleTokensReleased(opts *bind.WatchOpts, sink chan<- *LoopsoWrappedNonFungibleTokensReleased, tokenId []*big.Int, to []common.Address, attestationID [][32]byte) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "WrappedNonFungibleTokensReleased", tokenIdRule, toRule, attestationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoWrappedNonFungibleTokensReleased)
				if err := _Loopso.contract.UnpackLog(event, "WrappedNonFungibleTokensReleased", log); err != nil {
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

// ParseWrappedNonFungibleTokensReleased is a log parse operation binding the contract event 0xe436066797d3bcdab647d36495d87e5dba7a58a6c132890edcc038d783a596c1.
//
// Solidity: event WrappedNonFungibleTokensReleased(uint256 indexed tokenId, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) ParseWrappedNonFungibleTokensReleased(log types.Log) (*LoopsoWrappedNonFungibleTokensReleased, error) {
	event := new(LoopsoWrappedNonFungibleTokensReleased)
	if err := _Loopso.contract.UnpackLog(event, "WrappedNonFungibleTokensReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LoopsoWrappedTokensReleasedIterator is returned from FilterWrappedTokensReleased and is used to iterate over the raw logs and unpacked data for WrappedTokensReleased events raised by the Loopso contract.
type LoopsoWrappedTokensReleasedIterator struct {
	Event *LoopsoWrappedTokensReleased // Event containing the contract specifics and raw log

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
func (it *LoopsoWrappedTokensReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoopsoWrappedTokensReleased)
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
		it.Event = new(LoopsoWrappedTokensReleased)
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
func (it *LoopsoWrappedTokensReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoopsoWrappedTokensReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoopsoWrappedTokensReleased represents a WrappedTokensReleased event raised by the Loopso contract.
type LoopsoWrappedTokensReleased struct {
	Amount        *big.Int
	To            common.Address
	AttestationID [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWrappedTokensReleased is a free log retrieval operation binding the contract event 0x7aebc2282a0784072e43f12d40afc658f82fda3d0e830865d50dd7b0196df7f1.
//
// Solidity: event WrappedTokensReleased(uint256 indexed amount, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) FilterWrappedTokensReleased(opts *bind.FilterOpts, amount []*big.Int, to []common.Address, attestationID [][32]byte) (*LoopsoWrappedTokensReleasedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.FilterLogs(opts, "WrappedTokensReleased", amountRule, toRule, attestationIDRule)
	if err != nil {
		return nil, err
	}
	return &LoopsoWrappedTokensReleasedIterator{contract: _Loopso.contract, event: "WrappedTokensReleased", logs: logs, sub: sub}, nil
}

// WatchWrappedTokensReleased is a free log subscription operation binding the contract event 0x7aebc2282a0784072e43f12d40afc658f82fda3d0e830865d50dd7b0196df7f1.
//
// Solidity: event WrappedTokensReleased(uint256 indexed amount, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) WatchWrappedTokensReleased(opts *bind.WatchOpts, sink chan<- *LoopsoWrappedTokensReleased, amount []*big.Int, to []common.Address, attestationID [][32]byte) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var attestationIDRule []interface{}
	for _, attestationIDItem := range attestationID {
		attestationIDRule = append(attestationIDRule, attestationIDItem)
	}

	logs, sub, err := _Loopso.contract.WatchLogs(opts, "WrappedTokensReleased", amountRule, toRule, attestationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoopsoWrappedTokensReleased)
				if err := _Loopso.contract.UnpackLog(event, "WrappedTokensReleased", log); err != nil {
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

// ParseWrappedTokensReleased is a log parse operation binding the contract event 0x7aebc2282a0784072e43f12d40afc658f82fda3d0e830865d50dd7b0196df7f1.
//
// Solidity: event WrappedTokensReleased(uint256 indexed amount, address indexed to, bytes32 indexed attestationID)
func (_Loopso *LoopsoFilterer) ParseWrappedTokensReleased(log types.Log) (*LoopsoWrappedTokensReleased, error) {
	event := new(LoopsoWrappedTokensReleased)
	if err := _Loopso.contract.UnpackLog(event, "WrappedTokensReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
