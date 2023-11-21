package utils

import "bytes"

func EncodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}
