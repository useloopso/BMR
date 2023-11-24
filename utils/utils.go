package utils

import (
	"bytes"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func EncodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func IsClientConnected(client *ethclient.Client) bool {
	_, err := client.BlockNumber(context.Background())
	return err == nil
}

func IsZeroAddress(address common.Address) bool {
	return address == common.HexToAddress("0x0000000000000000000000000000000000000000")
}
