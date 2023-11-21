package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/useloopso/BMR/config"
	"github.com/useloopso/BMR/types"
)

type LoopsoClient struct {
	conns      map[int]*ethclient.Client
	chainInfos map[int]types.ChainInfo
}

func NewClient(chains []types.ChainInfo) (*LoopsoClient, error) {
	var c LoopsoClient
	c.conns = make(map[int]*ethclient.Client)
	c.chainInfos = make(map[int]types.ChainInfo)

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
