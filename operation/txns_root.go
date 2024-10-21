// Package operation provides functions to parse and manipulate Arbitrum Layer 2 (L2) transactions,
// including calculating the Merkle root hash and handling 'Start Block Transactions' which are
// required for block initialization.
//
// This package is designed to work with the Arbitrum sequencer feed, decoding and processing
// transactions in a manner compatible with the Arbitrum Nitro framework.
package operation

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

// CalculateTransactionsRoot calculates the Merkle root hash of a list of transactions.
//
// This function computes the root of the Merkle trie (also known as the transaction root)
// for a given set of Ethereum transactions using the DeriveSha function from the go-ethereum library.
//
// The transactions root is an essential part of the block header, providing a cryptographic
// summary of all transactions in the block.
//
// Parameters:
// - txs: A list of Ethereum transactions (types.Transactions) to compute the root for.
//
// Returns:
// - common.Hash: The resulting Merkle root hash of the transactions.
func CalculateTransactionsRoot(txs types.Transactions) common.Hash {
	root := types.DeriveSha(txs, trie.NewStackTrie(nil))

	return root
}
