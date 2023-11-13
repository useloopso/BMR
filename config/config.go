package config

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

type Config struct {
	MUMBAI_BRIDGE_ADDRESS        common.Address `json:"MUMBAI_BRIDGE_ADDRESS"`
	LUKSO_TESTNET_BRIDGE_ADDRESS common.Address `json:"LUKSO_TESTNET_BRIDGE_ADDRESS"`
	TOPIC_HASH                   common.Hash    `json:"TOPIC_HASH"`

	SEPOLIA_RPC_URL       string `json:"API_SEPOLIA"`
	MUMBAI_RPC_URL        string `json:"API_MUMBAI"`
	LUKSO_TESTNET_RPC_URL string `json:"API_LUKSO_TESTNET"`

	MAINNET_RPC_URL       string `json:"API_MAINNET"`
	LUKSO_MAINNET_RPC_URL string `json:"API_LUKSO_MAINNET"`

	PORT        string            `json:"PORT"`
	PRIVATE_KEY *ecdsa.PrivateKey `json:"PRIVATE_KEY"`
}

func LoadMainnetConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	/** VARIABLES **/
	var mumbai_address = common.HexToAddress("0xBb8d476fF7BEdCf2Eded931179196a17A3370A86")
	var lukso_testnet_address = common.HexToAddress("0x6C1aeA2C4933f040007a43Bc5683B0e068452c46")

	var topic_hash = common.HexToHash("0x11a762c44e982fe76f4f9b9c028673684f53601407c960c08a6f4472419373e6")

	/** .ENV **/
	api_mainnet := os.Getenv("MAINNET_RPC_URL")
	if api_mainnet == "" {
		return nil, fmt.Errorf("MAINNET_RPC_URL is not set in the .env file")
	}

	api_sepolia := os.Getenv("SEPOLIA_RPC_URL")
	if api_sepolia == "" {
		return nil, fmt.Errorf("SEPOLIA_RPC_URL is not set in the .env file")
	}

	api_mumbai := os.Getenv("MUMBAI_RPC_URL")
	if api_mumbai == "" {
		return nil, fmt.Errorf("MUMBAI_RPC_URL is not set in the .env file")
	}

	api_lukso_testnet := os.Getenv("LUKSO_TESTNET_RPC_URL")
	if api_lukso_testnet == "" {
		return nil, fmt.Errorf("LUKSO_TESTNET_RPC_URL is not set in the .env file")
	}

	api_lukso_mainnet := os.Getenv("LUKSO_MAINNET_RPC_URL")
	if api_lukso_mainnet == "" {
		return nil, fmt.Errorf("LUKSO_MAINNET_RPC_URL is not set in the .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default value
	}

	if os.Getenv("PRIVATE_KEY") == "" {
		return nil, fmt.Errorf("PRIVATE_KEY is not set in the .env file")
	}
	private_key, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		MUMBAI_BRIDGE_ADDRESS:        mumbai_address,
		LUKSO_TESTNET_BRIDGE_ADDRESS: lukso_testnet_address,
		TOPIC_HASH:                   topic_hash,

		SEPOLIA_RPC_URL:       api_sepolia,
		MUMBAI_RPC_URL:        api_mumbai,
		LUKSO_TESTNET_RPC_URL: api_lukso_testnet,

		MAINNET_RPC_URL:       api_mainnet,
		LUKSO_MAINNET_RPC_URL: api_lukso_mainnet,

		PORT:        port,
		PRIVATE_KEY: private_key,
	}, nil
}
