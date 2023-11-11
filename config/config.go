package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SEPOLIA_RPC_URL string `json:"API_SEPOLIA"`
	MUMBAI_RPC_URL  string `json:"API_MUMBAI"`
	MAINNET_RPC_URL string `json:"API_MAINNET"`
	PORT            string `json:"PORT"`
}

func LoadMainnetConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	api_mainnet := os.Getenv("MAINNET_RPC_URL")
	if api_mainnet == "" {
		return nil, fmt.Errorf("MAINNET_RPC_URL is not set in the .env file")
	}

	api_sepolia := os.Getenv("SEPOLIA_RPC_URL")
	if api_sepolia == "" {
		return nil, fmt.Errorf("SEPOLIA_RPC_URL is not set in the .env file")
	}

	api_mumbai := os.Getenv("MUMBAI_RPC_URL")
	if api_sepolia == "" {
		return nil, fmt.Errorf("MUMBAI_RPC_URL is not set in the .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default value
	}

	return &Config{
		SEPOLIA_RPC_URL: api_sepolia,
		MUMBAI_RPC_URL:  api_mumbai,
		MAINNET_RPC_URL: api_mainnet,
		PORT:            port,
	}, nil
}
