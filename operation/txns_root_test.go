package operation_test

import (
	"testing"

	"arbfeedtowasm/operation"
	"arbfeedtowasm/test"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TestCalculateTransactionsRoot verifies that the CalculateTransactionsRoot function correctly
// computes the Merkle root hash of a list of transactions.
//
// The test uses a sample set of parsed transactions from the Arbitrum sequencer feed, including
// the 'Start Block Transaction', to calculate the transactions root. It compares the computed
// transactions root hash with the expected hash (test.ExampleTransactionsRoot).
//
// If the calculated root does not match the expected root, the test fails and reports the discrepancy.
func TestCalculateTransactionsRoot(t *testing.T) {
	var txns types.Transactions = operation.ParseL2TransactionsWithStartTx(test.ExampleParsedMessage, test.ExampleLastTimestamp)
	var txnsRoot common.Hash = operation.CalculateTransactionsRoot(txns)

	if txnsRoot != test.ExampleTransactionsRoot {
		t.Errorf("Test failed for Calculate Transaction Root : expected %v, got %v", test.ExampleTransactionsRoot, txnsRoot)
	}
}
