package core

import (
	"fmt"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	TokensBridged                = crypto.Keccak256Hash([]byte("TokensBridged(bytes32,uint8)")).Hex()
	TokensBridgedBack            = crypto.Keccak256Hash([]byte("TokensBridgedBack(uint256,address,bytes32)")).Hex()
	NonFungibleTokensBridgedBack = crypto.Keccak256Hash([]byte("NonFungibleTokensBridgedBack(uint256,address,bytes32)")).Hex()
)

type TokenType uint8

const (
	Fungible TokenType = iota
	NonFungible
)

func HandleEvent(c *LoopsoClient, chain int, log ethtypes.Log) {
	eventSig := log.Topics[0].Hex()
	switch eventSig {
	case TokensBridged:
		fmt.Println("new tokens bridged event")
		handleTokensBridged(c, chain, log)
	case TokensBridgedBack:
		fmt.Println("new tokens bridged back event")
		handleTokensBridgedBack(c, chain, log)
	case NonFungibleTokensBridgedBack:
		fmt.Println("new non-fungible tokens bridged back event")
		handleNonFungibleTokensBridgedBack(c, chain, log)
	}
}

func handleTokensBridged(c *LoopsoClient, chain int, log ethtypes.Log) {
	e := praseTokensBridgedEvent(log)
	switch TokenType(e.TokenType) {
	case Fungible:
		bridgeFungibleTokens(c, chain, e.TransferID)
	case NonFungible:
		bridgeNonFungibleTokens(c, chain, e.TransferID)
	}
}

func handleTokensBridgedBack(c *LoopsoClient, chain int, log ethtypes.Log) {
	e := parseTokensBridgedBackEvent(log)
	bridgeFungibleTokensBack(c, chain, e.Amount, e.To, e.AttestationID)
}

func handleNonFungibleTokensBridgedBack(c *LoopsoClient, chain int, log ethtypes.Log) {
	e := parseNonFungibleTokensBridedBackEvent(log)
	bridgeNonFungibleTokensBack(c, chain, e.TokenId, e.To, e.AttestationID)
}
