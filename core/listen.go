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

func (c *LoopsoClient) ListenAll(stopCh <-chan struct{}) {
	var wg sync.WaitGroup

	for _, chainInfo := range c.chainInfos {
		wg.Add(1)
		go c.Listen(chainInfo.ChainID, &wg, stopCh)
	}

	wg.Wait()
}

func (c *LoopsoClient) Listen(chain int, wg *sync.WaitGroup, stopCh <-chan struct{}) error {
	defer wg.Done()
	if !c.IsChainConnected(chain) {
		fmt.Printf("Not connected to RPC on chain ID: %d\n", chain)
		return nil
	}

	chainInfo := c.chainInfos[chain]
	address := common.HexToAddress(chainInfo.BridgeAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logs := make(chan ethtypes.Log)
	sub, err := c.conns[chain].SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return err
	}

	fmt.Println("Listening on chain ID:", chain)

	const maxReconnectAttempts = 3
	reconnectAttempts := 0

	for {
		select {
		case <-stopCh:
			return nil // Stop listening if signaled
		case err := <-sub.Err():
			fmt.Println("Failed to get event log: ", err)

			reconnectAttempts++
			if reconnectAttempts > maxReconnectAttempts {
				return fmt.Errorf("maximum reconnection attempts reached")
			}

			backoff := time.Duration(reconnectAttempts) * time.Second
			fmt.Printf("Reconnecting in %v...\n", backoff)

			select {
			case <-time.After(backoff):
			case <-stopCh:
				return nil // Stop listening if signaled during backoff
			}

			sub, err = c.conns[chain].SubscribeFilterLogs(ctx, query, logs)
			if err != nil {
				return err
			}

			fmt.Println("listening on chain ID: ", chain)
		case vLog := <-logs:
			go c.HandleEvent(chain, vLog)
		}
	}
}
