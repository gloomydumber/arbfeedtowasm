package operation_test

import (
	"testing"

	"arbfeedtowasm/operation"
	"arbfeedtowasm/test"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestCalculateTransactionsRoot(t *testing.T) {
	var txns types.Transactions = operation.ParseL2TransactionsWithStartTx(test.ExampleParsedMessage, test.ExampleLastTimestamp)
	var txnsRoot common.Hash = operation.CalculateTransactionsRoot(txns)

	if txnsRoot != test.ExampleTransactionsRoot {
		t.Errorf("Test failed for Calculate Transaction Root : expected %v, got %v", test.ExampleTransactionsRoot, txnsRoot)
	}
}
