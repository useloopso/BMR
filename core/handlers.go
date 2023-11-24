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

func (c *LoopsoClient) HandleEvent(chain int, log ethtypes.Log) {
	eventSig := log.Topics[0].Hex()
	switch eventSig {
	case TokensBridged:
		fmt.Println("new tokens bridged event")
		if err := handleTokensBridged(c, chain, log); err != nil {
			fmt.Println(err)
			if err := c.RetryOnError(handleTokensBridged, chain, log); err != nil {
				fmt.Println(err)
			}
		}
	case TokensBridgedBack:
		fmt.Println("new tokens bridged back event")
		if err := handleTokensBridgedBack(c, chain, log); err != nil {
			fmt.Println(err)
			if err := c.RetryOnError(handleTokensBridgedBack, chain, log); err != nil {
				fmt.Println(err)
			}
		}
	case NonFungibleTokensBridgedBack:
		fmt.Println("new non-fungible tokens bridged back event")
		if err := handleNonFungibleTokensBridgedBack(c, chain, log); err != nil {
			fmt.Println(err)
			if err := c.RetryOnError(handleNonFungibleTokensBridgedBack, chain, log); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func handleTokensBridged(c *LoopsoClient, chain int, log ethtypes.Log) error {
	e := praseTokensBridgedEvent(log)
	switch TokenType(e.TokenType) {
	case Fungible:
		return c.bridgeFungibleTokens(chain, e.TransferID)
	case NonFungible:
		return c.bridgeNonFungibleTokens(chain, e.TransferID)
	}
	return nil
}

func handleTokensBridgedBack(c *LoopsoClient, chain int, log ethtypes.Log) error {
	e := parseTokensBridgedBackEvent(log)
	return c.bridgeFungibleTokensBack(chain, e.Amount, e.To, e.AttestationID)
}

func handleNonFungibleTokensBridgedBack(c *LoopsoClient, chain int, log ethtypes.Log) error {
	e := parseNonFungibleTokensBridedBackEvent(log)
	return c.bridgeNonFungibleTokensBack(chain, e.TokenId, e.To, e.AttestationID)
}
