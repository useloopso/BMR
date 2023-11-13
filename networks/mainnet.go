package mainnet

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	config "github.com/useloopso/BMR/config"
	loopso "github.com/useloopso/BMR/config/abi"
	model "github.com/useloopso/BMR/model"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var mumbaiClient *ethclient.Client
var luksoTestnetClient *ethclient.Client
var cfg *config.Config
var mumbaiBridgeInstance *loopso.Loopso
var luksoTestnetBridgeInstance *loopso.Loopso

// @dev instantiate config for ethereum mainnet
func init() {
	var err error
	// Load the configuration
	cfg, err = config.LoadMainnetConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the Mumbai client
	mumbaiClient, err = ethclient.Dial(cfg.MUMBAI_RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the Mumbai bridge contract
	mumbaiBridgeInstance, err = loopso.NewLoopso(cfg.MUMBAI_BRIDGE_ADDRESS, mumbaiClient)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the Lukso testnet client
	luksoTestnetClient, err = ethclient.Dial(cfg.LUKSO_TESTNET_RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the lukso testnet bridge contract
	luksoTestnetBridgeInstance, err = loopso.NewLoopso(cfg.LUKSO_TESTNET_BRIDGE_ADDRESS, luksoTestnetClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Contract is loaded")
}

func FetchBalance(c *fiber.Ctx) {
	// Get the balance of an account
	account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	balance, err := mumbaiClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	c.Status(200).SendString(fmt.Sprintf("Account balance: %d\n", balance)) // 25893180161173005034
}

func FetchBlock(c *fiber.Ctx) {
	// Get the latest known block
	block, err := mumbaiClient.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	main()

	c.SendString(fmt.Sprintf("Latest block: %d\n", block.Number().Uint64()))
}

func ListenToEvents(c *websocket.Conn) {
	// @dev Replace with actual contract address
	// contractAddress := common.HexToAddress("0xBb8d476fF7BEdCf2Eded931179196a17A3370A86")
	logs := make(chan types.Log)

	// Filter to capture the events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{cfg.MUMBAI_BRIDGE_ADDRESS},
		FromBlock: big.NewInt(0),
		ToBlock:   nil,
		Topics:    [][]common.Hash{{cfg.TOPIC_HASH}}, // topic hash
	}

	sub, err := mumbaiClient.SubscribeFilterLogs(context.Background(), query, logs)
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
				topicHash := event.Topics[1]

				// @todo Process the event data
				QueryBridge(&topicHash)
				if err = c.WriteMessage(mt, msg); err != nil {
					log.Printf("\nTransferID: %s\n", topicHash)
				}
			}
		}
	}()
}

func QueryBridge(topicHash *common.Hash) {
	var topicBytes [32]byte
	copy(topicBytes[:], topicHash[:])

	// Call the contract function
	transferId, err := mumbaiBridgeInstance.TokenTransfers(&bind.CallOpts{}, topicBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the keccak256 hash using Keccak256Hash
	// Expected: 0xa458b7ff0eb3c12e6e58c218f2a3111ab6cf26f757548bce2c887731f419675c
	// Current:  0xfc7fc10a4f0cda5df054a738f831d9278dd9d0226eb21e8960da9b7f9489d9d4
	hash := crypto.Keccak256Hash(transferId.TokenTransfer.TokenAddress.Bytes(), transferId.TokenTransfer.SrcChain.Bytes())

	attestationId, err := luksoTestnetBridgeInstance.AttestedTokens(&bind.CallOpts{}, hash)

	fmt.Println(hash, "heyyyyy", attestationId, "oooo2")

	// Parse the response into a struct
	tokenTransfer := model.TransferId{
		Timestamp:     transferId.TokenTransfer.Timestamp,
		SrcChain:      transferId.TokenTransfer.SrcChain,
		SrcAddress:    transferId.TokenTransfer.SrcAddress,
		DstChain:      transferId.TokenTransfer.DstChain,
		DstAddress:    transferId.TokenTransfer.DstAddress,
		TokenAddress:  transferId.TokenTransfer.TokenAddress,
		Amount:        transferId.Amount,
		AttestationId: hash,
	}

	WriteBridge(&tokenTransfer)
}

func WriteBridge(transferId *model.TransferId) {
	publicKey := cfg.PRIVATE_KEY.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("relayer address:", fromAddress)

	nonce, err := luksoTestnetClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := luksoTestnetClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("gas:", gasPrice, "nonce:", nonce)

	// a new keyed transactor
	// auth := bind.NewKeyedTransactor(cfg.PRIVATE_KEY)
	auth, err := bind.NewKeyedTransactorWithChainID(cfg.PRIVATE_KEY, big.NewInt(4201))

	// set transaction options attached to the keyed transactor
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// send tx & wait to get mined
	tx, err := luksoTestnetBridgeInstance.ReleaseWrappedTokens(auth, transferId.Amount, transferId.DstAddress, transferId.AttestationId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	// @todo verify it was mined

}

func main() {

}
