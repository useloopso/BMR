package core

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/useloopso/BMR/config"
	"github.com/useloopso/BMR/types"
)

type LoopsoClient struct {
	conns      map[int]*ethclient.Client
	chainInfos map[int]types.ChainInfo
	retryCount int
}

func NewClient(chains []types.ChainInfo) (*LoopsoClient, error) {
	var c LoopsoClient
	c.conns = make(map[int]*ethclient.Client)
	c.chainInfos = make(map[int]types.ChainInfo)
	c.retryCount = 3

	for _, cInfo := range chains {
		client, err := ethclient.Dial(cInfo.RpcURL)
		if err != nil {
			return nil, err
		}
		c.conns[cInfo.ChainID] = client
		c.chainInfos[cInfo.ChainID] = cInfo
	}

	return &c, nil
}

type loopsoEventHandler func(c *LoopsoClient, chain int, log ethtypes.Log) error

/*
retry on error will be called to handle errors on any of the event handlers
- go through the fallback rpcs
- replace current RPC with fallback RPC
- call failed func again
- if error == nil, return nil
- else try again with next fallback rpc
*/
func (c *LoopsoClient) RetryOnError(handler loopsoEventHandler, chain int, log ethtypes.Log) error {
	if len(c.chainInfos[chain].FallbackRpcs) == 0 {
		return errors.New("no fallback RPCs in chain config")
	}

	for i := 0; i < c.retryCount; i++ {
		fmt.Println("retrying...")
		for j, fallbackRpc := range c.chainInfos[chain].FallbackRpcs {
			client, err := ethclient.Dial(fallbackRpc)
			if err != nil {
				fmt.Println("Fallback rpc connect error: ", err)
				continue
			}

			oldClient := c.conns[chain]
			c.conns[chain] = client

			err = handler(c, chain, log)
			if err == nil {
				oldClient.Close()
				fmt.Println("call successful with fallback RPC")
				return nil
			}

			fmt.Println("call failed with fallback RPC no. ", j, " error: ", err)
			c.conns[chain].Close()
			c.conns[chain] = oldClient
		}
	}

	return errors.New("call failed with all fallback RPCs")
}

func (c *LoopsoClient) Auth(chainId int) (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(config.Config.PrivKey)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(int64(chainId)))
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (c *LoopsoClient) IsChainConnected(chain int) bool {
	return c.conns[chain] != nil
}
