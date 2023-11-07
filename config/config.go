package config

import (
    "os"
    "fmt"
    "github.com/joho/godotenv"
)

type Config struct {
    API_KEY string `json:"APIKey"`
    ALCHEMY_API_URL string `json:"APIURL"`
}

func LoadMainnetConfig() (*Config, error) {
    // load .env
    if err := godotenv.Load(); err != nil {
        return nil, err
    }

    // set API key
    apiKey := os.Getenv("ALCHEMY_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("ALCHEMY_API_KEY is not set in the .env file")
    }

    // set alchemy API URL
    apiUrl := fmt.Sprintf("https://eth-sepolia.g.alchemy.com/v2/%s", apiKey)

    return &Config{
        API_KEY: apiKey,
        ALCHEMY_API_URL: apiUrl,
    }, nil
}
