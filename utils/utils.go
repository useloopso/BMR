package utils

import (
	"bytes"
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

func EncodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func IsClientConnected(client *ethclient.Client) bool {
	_, err := client.BlockNumber(context.Background())
	return err == nil
}
