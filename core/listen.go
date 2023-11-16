package core

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *LoopsoClient) ListenAll() {
	var wg sync.WaitGroup

	for _, chainInfo := range c.chainInfos {
		wg.Add(1)
		go c.Listen(chainInfo.ChainID, &wg)
	}

	wg.Wait()
}

func (c *LoopsoClient) Listen(chain int, wg *sync.WaitGroup) error {
	defer wg.Done()
	if !c.IsChainConnected(chain) {
		return fmt.Errorf("not connected to RPC on target chain with chain ID: %d", chain)
	}

	chainInfo := c.chainInfos[chain]
	address := common.HexToAddress(chainInfo.BridgeAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	ctx := context.Background()

	logs := make(chan ethtypes.Log)
	sub, err := c.conns[chain].SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return err
	}

	fmt.Println("listening on chain ID: ", chain)

	const maxReconnectAttempts = 3
	reconnectAttempts := 0

	for {
		select {
		case err := <-sub.Err():
			fmt.Println("Failed to get event log: ", err)

			reconnectAttempts++
			if reconnectAttempts > maxReconnectAttempts {
				return fmt.Errorf("maximum reconnection attempts reached")
			}

			time.Sleep(3 * time.Second)

			sub, err = c.conns[chain].SubscribeFilterLogs(ctx, query, logs)
			if err != nil {
				return err
			}

			fmt.Println("listening on chain ID: ", chain)
		case vLog := <-logs:
			go HandleEvent(c, chain, vLog)
		}
	}
}
