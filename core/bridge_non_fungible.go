package core

import (
	"fmt"

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

	tokenTransferNFT, err := srcBridge.TokenTransfersNonFungible(nil, transferID)
	if err != nil {
		fmt.Println(err)
		return
	}

	dstChainId := int(tokenTransferNFT.TokenTransfer.DstChain.Int64())
	dstBridgeAddress := c.chainInfos[dstChainId].BridgeAddress
	dstBridgeClient := c.conns[dstChainId]

	dstBridge, err := contracts.NewLoopso(common.HexToAddress(dstBridgeAddress), dstBridgeClient)
	if err != nil {
		fmt.Println("failed to init bridge: ", err)
		return
	}

	// @todo: check whether NFT is supported

	auth, err := c.Auth(dstChainId)
	if err != nil {
		fmt.Println(err)
		return
	}

	attestationID := attestationID(tokenTransferNFT.TokenTransfer.TokenAddress, chain)
	tx, err := dstBridge.ReleaseWrappedNonFungibleTokens(auth, tokenTransferNFT.TokenID, tokenTransferNFT.TokenURI, tokenTransferNFT.TokenTransfer.DstAddress, attestationID)
	if err != nil {
		fmt.Println("error releasing wrapped non-fungible tokens: ", err)
		return
	}
	fmt.Println("tx hash: ", tx.Hash())
}
