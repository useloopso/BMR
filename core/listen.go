package core

import (
	"context"
	"fmt"
	"sync"

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

	logs := make(chan ethtypes.Log)

	c.conns[chain].SubscribeNewHead(context.Background(), make(chan *ethtypes.Header))

	sub, err := c.conns[chain].SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return err
	}

	fmt.Println("Listening on chain ID:", chain)

	for {
		select {
		case <-stopCh:
			return nil // Stop listening if signaled
		case err := <-sub.Err():
			fmt.Println("Failed to get event log: ", err)

			if err := c.tryReconnect(chain, 3, 1); err != nil {
				fmt.Println("failed to reconnect to RPC: ", err)
				if err := c.connectToFallback(chainInfo); err != nil {
					fmt.Println("all connections failed to chain ", chainInfo.ChainID)
					return err
				}
			}

			sub, err = c.conns[chain].SubscribeFilterLogs(context.Background(), query, logs)
			if err != nil {
				return err
			}

			fmt.Println("reconnected to RPC on chain ID: ", chain)

		case vLog := <-logs:
			go c.HandleEvent(chain, vLog)
		}
	}
}
