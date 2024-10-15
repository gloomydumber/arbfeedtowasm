package main

import (
	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/offchainlabs/nitro/arbos"
	"github.com/offchainlabs/nitro/arbos/arbostypes"
)

func main() {
	initialData := `{
   "message":{
      "header":{
         "kind":3,
         "sender":"0xa4b000000000000000000073657175656e636572",
         "blockNumber":20964528,
         "timestamp":1728917679,
         "requestId":null,
         "baseFeeL1":null
      },
      "l2Msg":"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"
   },
   "delayedMessagesRead":1718868
}`

	var jsonInput string
	if utils.IsJSLike(initialData) {
		jsonInput = utils.ConvertToJSON(initialData)
		fmt.Println("Converted JS-like input to valid JSON")
	} else {
		jsonInput = initialData
		fmt.Println("Input is already valid JSON")
	}

	var input feedtypes.IncomingMessage

	// Reading JSON from standard input
	err := json.Unmarshal([]byte(jsonInput), &input)
	if err != nil {
		log.Fatalf("Failed to parse JSON input: %v", err)
	}

	// Step 4: Decode L2msg from base64 string
	l2MsgData, err := base64.StdEncoding.DecodeString(input.Message.L2msg) // Now correctly decoding the base64 string
	if err != nil {
		log.Fatalf("Error decoding l2Msg: %v", err)
	}

	// Step 5: Parse the transactions
	chainId := big.NewInt(42161)
	txs, err := arbos.ParseL2Transactions(&arbostypes.L1IncomingMessage{
		Header: input.Message.Header,
		L2msg:  l2MsgData, // Now using the decoded L2msg data
	}, chainId)
	if err != nil {
		fmt.Println("Error parsing L2 transactions:", err)
	} else {
		fmt.Println("length of txns:", txs.Len())
		utils.PrintTransaction(txs[0])
		utils.PrintTransactionsRoot(txs)
	}
}