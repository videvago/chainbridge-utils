// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package msg

import (
	"fmt"
	"math/big"
)

type ChainId uint8
type TransferType string
type Bytes32 [32]byte

func (r Bytes32) Hex() string {
	return fmt.Sprintf("%x", r)
}

type Nonce uint64

func (n Nonce) Big() *big.Int {
	return big.NewInt(int64(n))
}

var FungibleTransfer TransferType = "FungibleTransfer"
var NonFungibleTransfer TransferType = "NonFungibleTransfer"
var GenericTransfer TransferType = "GenericTransfer"

// Message is used as a generic format to communicate between chains
type Message struct {
	DepositKey   Bytes32      // The key which identifies the proposal
	Source       ChainId      // Source where message was initiated
	Destination  ChainId      // Destination chain of message
	Type         TransferType // type of bridge transfer
	DepositNonce Nonce        // Nonce for the deposit
	ResourceId   Bytes32
	Payload      []interface{} // data associated with event sequence
}

func NewGenericTransfer(source, dest ChainId, nonce Nonce, depositKey Bytes32, resourceId Bytes32, metadata []byte) Message {
	return Message{
		DepositKey:   depositKey,
		Source:       source,
		Destination:  dest,
		Type:         GenericTransfer,
		DepositNonce: nonce,
		ResourceId:   resourceId,
		Payload: []interface{}{
			metadata,
		},
	}
}

func Bytes32FromSlice(in []byte) Bytes32 {
	var res Bytes32
	copy(res[:], in)
	return res
}
