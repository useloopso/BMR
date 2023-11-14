package core

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/useloopso/BMR/contracts"
)

func bridgeFungibleTokens(c *LoopsoClient, chain int, transferID [32]byte) {
	chainInfo := c.chainInfos[chain]
	client := c.conns[chain]

	srcBridge, err := contracts.NewLoopso(common.HexToAddress(chainInfo.BridgeAddress), client)
	if err != nil {
		fmt.Println("failed to init bridge contract: ", err)
		return
	}

	tokenTransfer, err := srcBridge.TokenTransfers(nil, transferID)
	if err != nil {
		fmt.Println(err)
		return
	}

	dstChainId := int(tokenTransfer.TokenTransfer.DstChain.Int64())
	dstBridgeAddress := c.chainInfos[dstChainId].BridgeAddress
	dstBridgeClient := c.conns[dstChainId]

	dstBridge, err := contracts.NewLoopso(common.HexToAddress(dstBridgeAddress), dstBridgeClient)
	if err != nil {
		fmt.Println("failed to init bridge: ", err)
		return
	}

	// check whether token is supported already
	isTokenSupported, err := dstBridge.IsTokenSupported(nil, tokenTransfer.TokenTransfer.TokenAddress, big.NewInt(int64(chain)))
	if err != nil {
		fmt.Println("failed to check is token supported:", err)
		return
	}

	auth, err := c.Auth(dstChainId)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !isTokenSupported {
		// if no --> all attestToken on destination chain
		att := attestationFromTokenTransfer(tokenTransfer, client)
		tx, err := dstBridge.AttestToken(auth, att)
		if err != nil {
			fmt.Println("failed to attest token: ", err)
			return
		}
		fmt.Println("attest token tx hash: ", tx.Hash())
	}

	attestationID := attestationID(tokenTransfer.TokenTransfer.TokenAddress, chain)
	tx, err := dstBridge.ReleaseWrappedTokens(auth, tokenTransfer.Amount, tokenTransfer.TokenTransfer.DstAddress, attestationID)
	if err != nil {
		fmt.Println("error releasing wrapped tokens: ", err)
		return
	}
	fmt.Println("tx hash: ", tx.Hash())
}

func bridgeFungibleTokensBack(c *LoopsoClient, chain int, amount *big.Int, dstAddress common.Address, attestationID [32]byte) {
	chainInfo := c.chainInfos[chain]
	client := c.conns[chain]

	srcBridge, err := contracts.NewLoopso(common.HexToAddress(chainInfo.BridgeAddress), client)
	if err != nil {
		fmt.Println("failed to init bridge contract: ", err)
		return
	}

	attestedToken, err := srcBridge.AttestedTokens(nil, attestationID)
	if err != nil {
		fmt.Println(err)
		return
	}

	dstChainId := int(attestedToken.TokenChain.Int64())
	dstBridgeAddress := c.chainInfos[dstChainId].BridgeAddress
	dstBridgeClient := c.conns[dstChainId]

	dstBridge, err := contracts.NewLoopso(common.HexToAddress(dstBridgeAddress), dstBridgeClient)
	if err != nil {
		fmt.Println("failed to init bridge: ", err)
		return
	}

	auth, err := c.Auth(chain)
	if err != nil {
		fmt.Println(err)
		return
	}

	tx, err := dstBridge.ReleaseTokens(auth, amount, dstAddress, attestedToken.TokenAddress)
	if err != nil {
		fmt.Println("error bridging tokens back: ", err)
		return
	}
	fmt.Println("tx hash: ", tx.Hash())
}
