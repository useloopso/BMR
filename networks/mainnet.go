package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/useloopso/BMR/config"
)

func main() {
    config, err := config.LoadMainnetConfig()
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

	// initialize client
	client, err := ethclient.Dial(config.ALCHEMY_API_URL)
	if err != nil {
		log.Fatal(err)
	}

	// Get the balance of an account
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Account balance: %d\n", balance) // 25893180161173005034

	// Get the latest known block
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Latest block: %d\n", block.Number().Uint64())
}
