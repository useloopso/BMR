package core

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/useloopso/BMR/contracts"
)

func (c *LoopsoClient) bridgeFungibleTokens(chain int, transferID [32]byte) error {
	chainInfo := c.chainInfos[chain]
	client := c.conns[chain]

	srcBridge, err := contracts.NewLoopso(common.HexToAddress(chainInfo.BridgeAddress), client)
	if err != nil {
		fmt.Println("failed to init bridge contract: ", err)
		return err
	}

	tokenTransfer, err := srcBridge.TokenTransfers(nil, transferID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dstChainId := int(tokenTransfer.TokenTransfer.DstChain.Int64())
	dstBridgeAddress := c.chainInfos[dstChainId].BridgeAddress
	dstBridgeClient := c.conns[dstChainId]

	dstBridge, err := contracts.NewLoopso(common.HexToAddress(dstBridgeAddress), dstBridgeClient)
	if err != nil {
		fmt.Println("failed to init bridge: ", err)
		return err
	}

	// check whether token is supported already
	isTokenSupported, err := dstBridge.IsTokenSupported(nil, tokenTransfer.TokenTransfer.TokenAddress, big.NewInt(int64(chain)))
	if err != nil {
		fmt.Println("failed to check is token supported:", err)
		return err
	}

	auth, err := c.Auth(dstChainId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if !isTokenSupported {
		// if no --> all attestToken on destination chain
		att := attestationFromTokenTransfer(tokenTransfer, client)

		tx, err := dstBridge.AttestToken(auth, att)
		if err != nil {
			fmt.Println("failed to attest token: ", err)
			return err
		}
		fmt.Println("attest token tx hash: ", tx.Hash())

		_, err = bind.WaitMined(context.Background(), dstBridgeClient, tx)
		if err != nil {
			fmt.Println("failed to wait for attest token tx to be mined: ", err)
			return err
		}
	}

	attestationID := attestationID(tokenTransfer.TokenTransfer.TokenAddress, chain)

	tx, err := dstBridge.ReleaseWrappedTokens(auth, tokenTransfer.Amount, tokenTransfer.TokenTransfer.DstAddress, attestationID)
	if err != nil {
		fmt.Println("error releasing wrapped tokens: ", err)
		return err
	}
	fmt.Println("tx hash: ", tx.Hash())
	return nil
}

func (c *LoopsoClient) bridgeFungibleTokensBack(chain int, amount *big.Int, dstAddress common.Address, attestationID [32]byte) error {
	chainInfo := c.chainInfos[chain]
	client := c.conns[chain]

	srcBridge, err := contracts.NewLoopso(common.HexToAddress(chainInfo.BridgeAddress), client)
	if err != nil {
		fmt.Println("failed to init bridge contract: ", err)
		return err
	}

	attestedToken, err := srcBridge.AttestedTokens(nil, attestationID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dstChainId := int(attestedToken.TokenChain.Int64())
	dstBridgeAddress := c.chainInfos[dstChainId].BridgeAddress
	dstBridgeClient := c.conns[dstChainId]

	dstBridge, err := contracts.NewLoopso(common.HexToAddress(dstBridgeAddress), dstBridgeClient)
	if err != nil {
		fmt.Println("failed to init bridge: ", err)
		return err
	}

	auth, err := c.Auth(dstChainId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	tx, err := dstBridge.ReleaseTokens(auth, amount, dstAddress, attestedToken.TokenAddress)
	if err != nil {
		fmt.Println("error bridging tokens back: ", err)
		return err
	}
	fmt.Println("tx hash: ", tx.Hash())
	return nil
}
