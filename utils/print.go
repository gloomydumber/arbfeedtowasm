// Package utils provides utility functions for handling Ethereum transactions,
// converting JavaScript-like object strings to valid JSON, and parsing incoming messages
// from the Arbitrum sequencer feed.
//
// This package is designed to support the processing and manipulation of data related
// to Ethereum transactions and feed messages, with a focus on the Arbitrum Layer 2 solution.
// It includes functions for converting input formats, calculating Merkle roots, and parsing
// JSON data into defined Go structures. Additionally, the package defines useful constants
// related to the Arbitrum network.
package utils

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

// PrintTransaction prints the JSON-encoded representation of an Ethereum transaction.
//
// This function takes a transaction (types.Transaction), marshals it into JSON format,
// and prints the resulting string. If the transaction cannot be marshaled due to an error,
// the function will log a fatal error and terminate the program.
//
// Parameters:
// - tx: The Ethereum transaction to be printed in JSON format.
func PrintTransaction(tx *types.Transaction) {
	jsonBytes, err := tx.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal transaction: %v", err)
	}

	fmt.Println(string(jsonBytes))
}

// PrintTransactionsRoot calculates and prints the Merkle root hash of a list of transactions.
//
// This function computes the Merkle trie root hash (the transaction root) of a set of transactions
// and prints it to the console. The root hash is a cryptographic summary of all transactions in the block.
//
// Parameters:
// - txs: The list of Ethereum transactions (types.Transactions) to calculate the root for.
func PrintTransactionsRoot(txs types.Transactions) {
	root := types.DeriveSha(txs, trie.NewStackTrie(nil))

	fmt.Println("Transactions root:", root)
}

// PrintTransactionsLength prints the number of transactions in a list of Ethereum transactions.
//
// This function prints the length (number of transactions) in the provided list of transactions.
//
// Parameters:
// - txs: The list of Ethereum transactions (types.Transactions) to count and print.
func PrintTransactionsLength(txs types.Transactions) {
	fmt.Println("Transactions length:", txs.Len())
}
