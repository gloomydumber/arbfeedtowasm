package jsutil

import (
	"encoding/json"
	"syscall/js"

	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/operation"
	"arbfeedtowasm/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func CalculateTransactionsRootWithStartTx(this js.Value, p []js.Value) interface{} {
	var sequencerMessage = p[0].String()
	var lastTimestamp = uint64(p[1].Int())

	var msg feedtypes.IncomingMessage = utils.ParseIncomingMessage(sequencerMessage)
	var txns types.Transactions = operation.ParseL2TransactionsWithStartTx(msg, lastTimestamp)
	var txnsRoot common.Hash = operation.CalculateTransactionsRoot(txns)

	return js.ValueOf(txnsRoot.String())
}

func GetTransactionDetailsWithRoot(this js.Value, p []js.Value) interface{} {
	var sequencerMessage = p[0].String()
	var lastTimestamp = uint64(p[1].Int())

	var msg feedtypes.IncomingMessage = utils.ParseIncomingMessage(sequencerMessage)
	var txns types.Transactions = operation.ParseL2TransactionsWithStartTx(msg, lastTimestamp)
	var txnsRoot common.Hash = operation.CalculateTransactionsRoot(txns)

	response := feedtypes.Response{
		Result: "success",
		Data: &feedtypes.ResponseData{
			Transactions:     txns,
			TransactionsRoot: txnsRoot.String(),
		},
	}

	jsonResult, err := json.Marshal(response)
	if err != nil {
		response = feedtypes.Response{
			Result: "error",
			Error: &feedtypes.ErrorResponse{
				Message: "Failed to marshal response to JSON",
			},
		}
		jsonResult, _ = json.Marshal(response)
		return js.ValueOf(string(jsonResult))
	}

	return js.ValueOf(string(jsonResult))
}
