package utils

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

// TODO: describe godocs here
func PrintTransaction(tx *types.Transaction) {
	jsonBytes, err := tx.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal transaction: %v", err)
	}

	fmt.Println(string(jsonBytes))
}

// TODO: describe godocs here
func PrintTransactionsRoot(txs types.Transactions) {
	root := types.DeriveSha(txs, trie.NewStackTrie(nil))

	fmt.Println("Transactions root:", root)
}

// TODO: describe godocs here
func PrintTransactionsLength(txs types.Transactions) {
	fmt.Println("Transactions length:", txs.Len())
}
