package core

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/useloopso/BMR/contracts"
)

func praseTokensBridgedEvent(l ethtypes.Log) contracts.LoopsoTokensBridged {
	var transferID [32]byte
	copy(transferID[:], l.Topics[1].Bytes()[:32])
	tokenType := uint8(binary.LittleEndian.Uint32(l.Topics[2].Bytes()))
	return contracts.LoopsoTokensBridged{
		TransferID: transferID,
		TokenType:  tokenType,
		Raw:        l,
	}
}

func parseTokensBridgedBackEvent(l ethtypes.Log) contracts.LoopsoTokensBridgedBack {
	var attestationID [32]byte
	amount := new(big.Int).SetBytes(l.Topics[1].Bytes())
	to := common.BytesToAddress(l.Topics[2].Bytes())
	copy(attestationID[:], l.Topics[3].Bytes()[:32])
	return contracts.LoopsoTokensBridgedBack{
		Amount:        amount,
		To:            to,
		AttestationID: attestationID,
		Raw:           l,
	}
}

func parseNonFungibleTokensBridedBackEvent(l ethtypes.Log) contracts.LoopsoNonFungibleTokensBridgedBack {
	// TODO
	return contracts.LoopsoNonFungibleTokensBridgedBack{}
}

func attestationFromTokenTransfer(tokenTransfer struct {
	TokenTransfer contracts.ILoopsoTokenTransferBase
	Amount        *big.Int
}, client *ethclient.Client) contracts.ILoopsoTokenAttestation {
	token, err := contracts.NewERC20(tokenTransfer.TokenTransfer.TokenAddress, client)
	if err != nil {
		fmt.Println("failed to init erc20 token: ", err)
		return contracts.ILoopsoTokenAttestation{}
	}

	decimals, err := token.Decimals(nil)
	if err != nil {
		fmt.Println("failed to get token decimals: ", err)
		return contracts.ILoopsoTokenAttestation{}
	}

	name, err := token.Name(nil)
	if err != nil {
		fmt.Println("failed to get token name: ", err)
		return contracts.ILoopsoTokenAttestation{}
	}

	symbol, err := token.Symbol(nil)
	if err != nil {
		fmt.Println("failed to get token symbol: ", err)
		return contracts.ILoopsoTokenAttestation{}
	}

	return contracts.ILoopsoTokenAttestation{
		TokenAddress:        tokenTransfer.TokenTransfer.TokenAddress,
		TokenChain:          tokenTransfer.TokenTransfer.SrcChain,
		TokenType:           0,
		Decimals:            decimals,
		Symbol:              symbol,
		Name:                name,
		WrappedTokenAddress: common.Address{},
	}
}

func attestationFromNonFungibleTokenTransfer(nonFungibleTokenTransfer struct {
	TokenTransfer contracts.ILoopsoTokenTransferBase
	TokenID       *big.Int
	TokenURI      string
}, client *ethclient.Client) contracts.ILoopsoTokenAttestation {
	nonFungibleToken, err := contracts.NewERC721(nonFungibleTokenTransfer.TokenTransfer.TokenAddress, client)
	if err != nil {
		fmt.Println("failed to init erc721 token: ", err)
		return contracts.ILoopsoTokenAttestation{}
	}

	name, err := nonFungibleToken.Name(nil)
	if err != nil {
		fmt.Println("failed to get token name: ", err)
		return contracts.ILoopsoTokenAttestation{}
	}

	symbol, err := nonFungibleToken.Symbol(nil)
	if err != nil {
		fmt.Println("failed to get token symbol: ", err)
		return contracts.ILoopsoTokenAttestation{}
	}

	return contracts.ILoopsoTokenAttestation{
		TokenAddress:        nonFungibleTokenTransfer.TokenTransfer.TokenAddress,
		TokenChain:          nonFungibleTokenTransfer.TokenTransfer.SrcChain,
		TokenType:           1,
		Decimals:            uint8(0),
		Symbol:              symbol,
		Name:                name,
		WrappedTokenAddress: common.Address{},
	}
}

func attestationID(tokenAddress common.Address, tokenChain int) common.Hash {
	chainBytes := common.LeftPadBytes(big.NewInt(int64(tokenChain)).Bytes(), 32)
	return crypto.Keccak256Hash(tokenAddress[:], chainBytes)
}
