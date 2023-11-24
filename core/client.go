package core

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/useloopso/BMR/config"
	"github.com/useloopso/BMR/types"
	"github.com/useloopso/BMR/utils"
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

func (c *LoopsoClient) RetryOnError(handler loopsoEventHandler, chain int, log ethtypes.Log) error {
	for _, chainInfo := range c.chainInfos {
		if err := c.connectToFallback(chainInfo); err != nil {
			fmt.Println(err)
		}
	}

	for i := 0; i < c.retryCount; i++ {
		fmt.Println("retrying...")
		err := handler(c, chain, log)
		if err == nil {
			fmt.Println("call successful with fallback RPC")
			return nil
		}
	}

	return errors.New("call failed with fallback RPCs")
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

func (c *LoopsoClient) connectToFallback(chainInfo types.ChainInfo) error {
	if !utils.IsClientConnected(c.conns[chainInfo.ChainID]) {
		c.conns[chainInfo.ChainID].Close()
		fmt.Println("RPC error on chain: ", chainInfo.ChainID)

		for _, fallbackRpc := range chainInfo.FallbackRpcs {
			client, err := ethclient.Dial(fallbackRpc)
			if err != nil {
				fmt.Println("Fallback rpc connect error: ", err)
				continue
			}

			c.conns[chainInfo.ChainID] = client
			fmt.Println("reconnected to RPC on chain: ", chainInfo.ChainID)
			return nil
		}
		return fmt.Errorf("failed to connect to any fallback RPC on chain %d", chainInfo.ChainID)
	}
	return nil
}

func (c *LoopsoClient) tryReconnect(chain int, attempts int, reconnectDelay time.Duration) error {
	c.conns[chain].Close()
	chainInfo := c.chainInfos[chain]

	var errOut error

	for i := 0; i < attempts; i++ {
		client, err := ethclient.Dial(chainInfo.RpcURL)
		if err != nil {
			errOut = err
			time.Sleep(time.Second * reconnectDelay)
			continue
		}
		c.conns[chain] = client
		return nil
	}

	return errOut
}
