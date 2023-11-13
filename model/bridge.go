package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

/**
 * @dev function: tokenTransfer
 * @dev types: tuple(block.timestamp: uint256, block.chainid: uint256, msg.sender: address, _dstChain: uint256,_dstAddress: address, _token: address), _amount: uint256
 * @dev output: (1699659701,80001,0x1c46D242755040a0032505fD33C6e8b83293a332,4201,0x1c46D242755040a0032505fD33C6e8b83293a332,0x8cBF42B6590614AbE7AB5ffc89aF153F5d620fC3)
 */
type TransferId struct {
	Timestamp     *big.Int
	SrcChain      *big.Int
	SrcAddress    common.Address
	DstChain      *big.Int
	DstAddress    common.Address
	TokenAddress  common.Address
	Amount        *big.Int
	AttestationId common.Hash
}
