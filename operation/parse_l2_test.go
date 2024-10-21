package operation_test

import (
	"testing"

	"arbfeedtowasm/operation"
	"arbfeedtowasm/test"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/go-cmp/cmp"
)

// TestDecodeL2Message verifies that the DecodeL2Message function correctly decodes
// a base64-encoded L2 message into its byte representation.
//
// The test uses example data from the test package, which provides a sample parsed message
// and its corresponding decoded output. The test checks whether the decoded L2 message matches
// the expected byte array.
//
// If there is a difference between the expected and actual results, the test fails and
// outputs a diff showing where the mismatch occurred.
func TestDecodeL2Message(t *testing.T) {
	var decodedL2Message []byte = operation.DecodeL2Message(test.ExampleParsedMessage.Message.Message.L2msg)

	if diff := cmp.Diff(test.ExampleDecodedL2Message, decodedL2Message); diff != "" {
		t.Errorf("Test %s failed (-expected +got):\n%s", "Decode L2 Message", diff)
	}
}

// TestParseL2Transactions ensures that the ParseL2Transactions function correctly decodes
// and parses an L2 message into a list of transactions.
//
// The test uses example parsed messages and compares the parsed transactions against
// a set of expected transactions.
//
// The test iterates over each transaction, printing and comparing each field in the transaction.
// If a mismatch is found, the test fails and provides detailed information about the difference.
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

// TestParseL2TransactionsWithStartTx verifies that the ParseL2TransactionsWithStartTx function
// correctly appends the 'Start Block Transaction' to the list of parsed transactions.
//
// The test constructs the expected list of transactions, where the 'Start Block Transaction'
// is added to the beginning, followed by the parsed transactions.
//
// The test compares each transaction in the result with the expected transactions, printing
// and comparing individual fields. If a mismatch is found, the test fails and outputs the
// relevant details.
func TestParseL2TransactionsWithStartTx(t *testing.T) {
	txns := operation.ParseL2TransactionsWithStartTx(test.ExampleParsedMessage, test.ExampleLastTimestamp)
	ExpectedTxns := append(types.Transactions{test.ExampleStartTx}, test.GetExampleSignedTxns()...)

	// Compare fields for all transactions
	for i, expectedTx := range ExpectedTxns {
		gotTx := txns[i]
		test.PrintTransactionFields(gotTx, expectedTx)
		if !test.CompareTransactionFields(gotTx, expectedTx) {
			t.Errorf("Test failed for transaction %d: expected %v, got %v", i, expectedTx, gotTx)
		}
	}
}
