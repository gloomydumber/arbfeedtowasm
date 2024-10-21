// Package operation provides functions to parse and manipulate Arbitrum Layer 2 (L2) transactions,
// including calculating the Merkle root hash and handling 'Start Block Transactions' which are
// required for block initialization.
//
// This package is designed to work with the Arbitrum sequencer feed, decoding and processing
// transactions in a manner compatible with the Arbitrum Nitro framework.
package operation

import (
	"encoding/base64"
	"log"
	"math/big"

	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/utils"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/nitro/arbos"
	"github.com/offchainlabs/nitro/arbos/arbostypes"
)

// DecodeL2Message decodes a base64-encoded L2 message and handles errors internally.
func DecodeL2Message(l2MsgBase64 string) []byte {
	l2MsgData, err := base64.StdEncoding.DecodeString(l2MsgBase64)
	if err != nil {
		log.Fatalf("Error decoding l2Msg: %v", err)
	}

	return l2MsgData
}

// ParseL2Transactions decodes a base64-encoded L2 message and parses it into transactions.
// It uses the arbos package from Nitro to handle the decoding and parsing logic.
//
// The input is an IncomingMessage from the Arbitrum sequencer feed, and the function
// returns the parsed transactions.
//
// See also: https://github.com/OffchainLabs/nitro/blob/5f24df4740bdbf3f0dd07aa32ead06c2919087e4/arbos/parse_l2.go#L20
func ParseL2Transactions(msg feedtypes.IncomingMessage) types.Transactions {
	l2MsgData := DecodeL2Message(msg.Message.Message.L2msg)

	// Initialize the L1IncomingMessage
	l1IncomingMsg := &arbostypes.L1IncomingMessage{
		Header: msg.Message.Message.Header,
		L2msg:  l2MsgData,
	}

	// Check if BatchGasCost exists, and if so, assign it
	if msg.Message.Message.BatchGasCost != nil {
		l1IncomingMsg.BatchGasCost = msg.Message.Message.BatchGasCost
	}

	// Call ParseL2Transactions with the conditional l1IncomingMsg
	txns, err := arbos.ParseL2Transactions(l1IncomingMsg, big.NewInt(utils.ArbiturmChainId))

	if err != nil {
		log.Fatalf("Error parsing txns: %v", err)
	}

	return txns
}

// ParseL2TransactionsWithStartTx decodes a base64-encoded L2 message and parses it into transactions,
// including the 'Start Block Transaction'.
//
// The 'Start Block Transaction' is an ArbitrumInternalTx added by the Sequencer at the beginning of every block.
// This transaction can also be viewed on Arbiscan for every block.
//
// Since the Arbitrum Sequencer feed does not include information about the 'Start Block Transaction',
// it must be manually appended to the very first position of the parsed transactions.
//
// The function takes an IncomingMessage from the Arbitrum sequencer feed and the lastTimestamp, which
// is the timestamp from the previous sequencer feed message, used to construct the 'Start Block Transaction'.
//
// Returns the transactions with the manually appended 'Start Block Transaction'.
func ParseL2TransactionsWithStartTx(msg feedtypes.IncomingMessage, lastTimestamp uint64) types.Transactions {
	txns := ParseL2Transactions(msg)
	txns = AppendStartTransaction(txns, msg, lastTimestamp)

	return txns
}
