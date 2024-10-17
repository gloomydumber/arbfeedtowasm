package operation

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

// TODO: describe godocs here
func CalculateTransactionsRoot(txs types.Transactions) common.Hash {
	root := types.DeriveSha(txs, trie.NewStackTrie(nil))

	return root
}