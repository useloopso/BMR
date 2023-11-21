package core

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/useloopso/BMR/contracts"
)

func bridgeNonFungibleTokens(c *LoopsoClient, chain int, transferID [32]byte) {
	chainInfo := c.chainInfos[chain]
	client := c.conns[chain]

	srcBridge, err := contracts.NewLoopso(common.HexToAddress(chainInfo.BridgeAddress), client)
	if err != nil {
		fmt.Println("failed to init bridge contract: ", err)
		return
	}

	nonFungibleTokenTransfer, err := srcBridge.TokenTransfersNonFungible(nil, transferID)
	if err != nil {
		fmt.Println(err)
		return
	}

	dstChainId := int(nonFungibleTokenTransfer.TokenTransfer.DstChain.Int64())
	dstBridgeAddress := c.chainInfos[dstChainId].BridgeAddress
	dstBridgeClient := c.conns[dstChainId]

	dstBridge, err := contracts.NewLoopso(common.HexToAddress(dstBridgeAddress), dstBridgeClient)
	if err != nil {
		fmt.Println("failed to init bridge: ", err)
		return
	}

	isFungibleTokenSupported, err := dstBridge.IsTokenSupported(nil, nonFungibleTokenTransfer.TokenTransfer.TokenAddress, big.NewInt(int64(chain)))
	if err != nil {
		fmt.Println("failed to check is non-fungible token supported:", err)
		return
	}

	auth, err := c.Auth(dstChainId)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !isFungibleTokenSupported {
		// if no --> all attestToken on destination chain
		att := attestationFromNonFungibleTokenTransfer(nonFungibleTokenTransfer, client)
		tx, err := dstBridge.AttestToken(auth, att)
		if err != nil {
			fmt.Println("failed to attest token: ", err)
			return
		}
		fmt.Println("attest token tx hash: ", tx.Hash())

		_, err = bind.WaitMined(context.Background(), dstBridgeClient, tx)
		if err != nil {
			fmt.Println("failed to wait for attest token tx to be mined: ", err)
			return
		}
	}

	attestationID := attestationID(nonFungibleTokenTransfer.TokenTransfer.TokenAddress, chain)
	tx, err := dstBridge.ReleaseWrappedNonFungibleTokens(auth, nonFungibleTokenTransfer.TokenID, nonFungibleTokenTransfer.TokenURI, nonFungibleTokenTransfer.TokenTransfer.DstAddress, attestationID)
	if err != nil {
		fmt.Println("error releasing wrapped non-fungible tokens: ", err)
		return
	}
	fmt.Println("tx hash: ", tx.Hash())
}

func bridgeNonFungibleTokensBack(c *LoopsoClient, chain int, tokenId *big.Int, dstAddress common.Address, attestationID [32]byte) {
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

	auth, err := c.Auth(dstChainId)
	if err != nil {
		fmt.Println(err)
		return
	}

	tx, err := dstBridge.ReleaseNonFungibleTokens(auth, tokenId, dstAddress, attestedToken.TokenAddress)
	if err != nil {
		fmt.Println("error bridging non-fungible tokens back: ", err)
		return
	}
	fmt.Println("tx hash: ", tx.Hash())
}
