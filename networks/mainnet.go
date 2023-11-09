package mainnet

import (
	"context"
	"fmt"
	"log"
	"math/big"

	// "strings"

	"github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/accounts/abi"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/gofiber/fiber/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/useloopso/BMR/config"
)

var client *ethclient.Client
var cfg *config.Config

// @dev instantiate config for ethereum mainnet
func init() {
	var err error
	// Load the configuration
	cfg, err = config.LoadMainnetConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the Ethereum client
	client, err = ethclient.Dial(cfg.MAINNET_RPC_URL)
	if err != nil {
		log.Fatal(err)
	}
}

func FetchBalance(c *fiber.Ctx) {
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

	c.SendString(fmt.Sprintf("Latest block: %d\n", block.Number().Uint64()))
}

func ListenToEvents() {
	// @dev Replace with actual contract address
	contractAddress := common.HexToAddress("0x6B175474E89094C44Da98b954EedeAC495271d0F")

	// Receive events in a Go channel
	logs := make(chan types.Log)

	// Create a filter to capture the events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: big.NewInt(0),
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")}}, // topic hash
	}

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case event := <-logs:
			// transferID := event.Topics[1]
			// tokenAddress := common.BytesToAddress(event.Data[12:32]) // adjust byte offsets
			// amount := new(big.Int).SetBytes(event.Data[32:])

			fmt.Printf("Log Name: Approval\n")
			tokenAddress := event.Topics[1].Hex()
			amount := event.Topics[2].Big()

			// @todo Process the event data
			// @params transferID, tokenAddress, and amount

			// fmt.Printf("TransferID: %s\n", transferID)
			// fmt.Printf("Token Address: %s\n", tokenAddress.Hex())
			// fmt.Printf("Amount: %s\n", amount.String())

			fmt.Printf("Address: %s\n", tokenAddress)
			fmt.Printf("Amount: %s\n", amount)
		}
	}
}

func main() {
	// FetchBalance()
}
