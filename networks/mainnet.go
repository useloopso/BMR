package mainnet

import (
	"context"
	"fmt"
	"log"
	"math/big"

	config "github.com/useloopso/BMR/config"
	loopso "github.com/useloopso/BMR/config/abi"

	"github.com/ethereum/go-ethereum"

	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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
	client, err = ethclient.Dial(cfg.MUMBAI_RPC_URL)
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

	c.Status(200).SendString(fmt.Sprintf("Account balance: %d\n", balance)) // 25893180161173005034
}

func FetchBlock(c *fiber.Ctx) {
	// Get the latest known block
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	main()

	c.SendString(fmt.Sprintf("Latest block: %d\n", block.Number().Uint64()))
}

func ListenToEvents(c *websocket.Conn) {
	// @dev Replace with actual contract address
	contractAddress := common.HexToAddress("0xBb8d476fF7BEdCf2Eded931179196a17A3370A86")
	logs := make(chan types.Log)

	// Filter to capture the events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: big.NewInt(0),
		ToBlock:   nil,
		Topics:    [][]common.Hash{{common.HexToHash("0x11a762c44e982fe76f4f9b9c028673684f53601407c960c08a6f4472419373e6")}}, // topic hash
	}

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		var (
			mt  int
			msg []byte
		)
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case event := <-logs:
				transferID := event.Topics[1]

				// @todo Process the event data
				if err = c.WriteMessage(mt, msg); err != nil {
					log.Printf("TransferID: %s\n", transferID)
				}
			}
		}
	}()
}

func main() {
	mainAddress := common.HexToAddress("0xbb8d476ff7bedcf2eded931179196a17a3370a86")
	instance, err := loopso.NewLoopso(mainAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract is loaded")

	topic := "0xc353f0420caa96d0f2dbfd113ce48f76275dfa81ba13bdbd7bd8771afecb2318"

	targetAddressBytes := common.HexToHash(topic)

	fmt.Println(targetAddressBytes)

	// Call the contract function
	tuple, err := instance.TokenTransfers(&bind.CallOpts{}, targetAddressBytes)
	if err != nil {
		log.Fatal(err)
	}

	tokenTransfer := tuple.TokenTransfer
	amount := tuple.Amount

	fmt.Println("Token Transfer:", tokenTransfer)
	fmt.Println("Amount:", amount)
}

/*
function: tokenTransfer

types: tuple(block.timestamp: uint256, block.chainid: uint256, msg.sender: address,_dstChain: uint256,_dstAddress: address,_token,_amount: address)

output: (1699659701,80001,0x1c46D242755040a0032505fD33C6e8b83293a332,4201,0x1c46D242755040a0032505fD33C6e8b83293a332,0x8cBF42B6590614AbE7AB5ffc89aF153F5d620fC3)
*/
