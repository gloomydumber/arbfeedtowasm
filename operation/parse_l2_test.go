package operation_test

import (
	"testing"

	"arbfeedtowasm/operation"
	"arbfeedtowasm/test"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/go-cmp/cmp"
)

func TestDecodeL2Message(t *testing.T) {
	var decodedL2Message []byte = operation.DecodeL2Message(test.ExampleParsedMessage.Message.Message.L2msg)

	if diff := cmp.Diff(test.ExampleDecodedL2Message, decodedL2Message); diff != "" {
		t.Errorf("Test %s failed (-expected +got):\n%s", "Decode L2 Message", diff)
	}
}

// TODO: Properly configure TestParseL2Transactions.
// The example transactions (test solution) are not signed yet, but the actual parsed messages are signed.
func TestParseL2Transactions(t *testing.T) {
	txns := operation.ParseL2Transactions(test.ExampleParsedMessage)

	// Compare fields for all transactions
	for i, expectedTx := range test.GetExampleSignedTxns() {
		gotTx := txns[i]
		test.PrintTransactionFields(gotTx, expectedTx)
		if !test.CompareTransactionFields(gotTx, expectedTx) {
			t.Errorf("Test failed for transaction %d: expected %v, got %v", i, expectedTx, gotTx)
		}
	}
}

func TestParseL2TransactionsWithStartTx(t *testing.T) {
	txns := operation.ParseL2TransactionsWithStartTx(test.ExampleParsedMessage, test.ExampleLastTimestamp)
	ExpectedTxns :=  append(types.Transactions{test.ExampleStartTx}, test.GetExampleSignedTxns()...)
	
	// Compare fields for all transactions
	for i, expectedTx := range ExpectedTxns {
		gotTx := txns[i]
		test.PrintTransactionFields(gotTx, expectedTx)
		if !test.CompareTransactionFields(gotTx, expectedTx) {
			t.Errorf("Test failed for transaction %d: expected %v, got %v", i, expectedTx, gotTx)
		}
	}
}
