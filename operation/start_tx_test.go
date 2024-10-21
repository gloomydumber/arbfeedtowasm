package operation_test

import (
	"testing"

	"arbfeedtowasm/operation"
	"arbfeedtowasm/test"

	"github.com/ethereum/go-ethereum/core/types"
)

// TestAppendStartTransaction verifies that the AppendStartTransaction function correctly appends
// the 'Start Block Transaction' to the beginning of the L2 transactions list.
//
// The test uses example signed transactions and an example parsed message from the Arbitrum sequencer feed.
// It calls AppendStartTransaction to append the 'Start Block Transaction' and checks whether the
// first transaction in the resulting list matches the expected 'Start Block Transaction'.
//
// The comparison is made using a helper function, CompareArbitrumInternalTx, which ensures that all
// fields of the ArbitrumInternalTx are properly compared.
//
// If the comparison fails, the test outputs an error message.
func TestAppendStartTransaction(t *testing.T) {
	var txns types.Transactions = test.GetExampleSignedTxns()
	txns = operation.AppendStartTransaction(txns, test.ExampleParsedMessage, test.ExampleLastTimestamp)
	var startTx = txns[0]

	if !test.CompareArbitrumInternalTx(startTx, test.ExampleStartTx) {
		t.Errorf("TestAppendStartTransaction (Using ArbitrumInternalTx comparison) failed. Expected transactions to be equal")
	}
}
