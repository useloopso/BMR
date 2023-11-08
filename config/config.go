package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SEPOLIA_RPC_URL string `json:"API_TESTNET"`
	MAINNET_RPC_URL string `json:"API_MAINNET"`
}

func LoadMainnetConfig() (*Config, error) {
	// load .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// set mainnet
	api_mainnet := os.Getenv("MAINNET_RPC_URL")
	if api_mainnet == "" {
		return nil, fmt.Errorf("MAINNET_RPC_URL is not set in the .env file")
	}

	// set sepolia
	api_sepolia := os.Getenv("SEPOLIA_RPC_URL")
	if api_sepolia == "" {
		return nil, fmt.Errorf("SEPOLIA_RPC_URL is not set in the .env file")
	}

	return &Config{
		SEPOLIA_RPC_URL: api_sepolia,
		MAINNET_RPC_URL: api_mainnet,
	}, nil
}
