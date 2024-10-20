package operation_test

import (
	"testing"

	"arbfeedtowasm/operation"
	"arbfeedtowasm/test"

	"github.com/ethereum/go-ethereum/core/types"
)

func TestAppendStartTransaction(t *testing.T) {
	var txns types.Transactions = test.GetExampleSignedTxns()
	txns = operation.AppendStartTransaction(txns, test.ExampleParsedMessage, uint64(1728980884))
	var startTx = txns[0]

	if !test.CompareArbitrumInternalTx(startTx, test.ExampleStartTx) {
		t.Errorf("TestAppendStartTransaction (Using ArbitrumInternalTx comparison) failed. Expected transactions to be equal")
	}
}
