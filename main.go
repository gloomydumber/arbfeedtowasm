package main

import (
	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/utils"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/nitro/arbos"
	"github.com/offchainlabs/nitro/arbos/arbostypes"
	"github.com/offchainlabs/nitro/arbos/util"
)

func main() {
	initialData := `{"sequenceNumber":241801449,"message":{"message":{"header":{"kind":3,"sender":"0xa4b000000000000000000073657175656e636572","blockNumber":20969771,"timestamp":1728980884,"requestId":null,"baseFeeL1":null},"l2Msg":"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"},"delayedMessagesRead":1719783}}`
	
	// Convert JavaScript-like object to JSON if necessary
	if utils.IsJSObject(initialData) {
		initialData = utils.ConvertToJSON(initialData)
	}

	var input feedtypes.IncomingMessage
	
	// Reading JSON from standard input
	err := json.Unmarshal([]byte(initialData), &input)
	if err != nil {
		log.Fatalf("Failed to parse JSON input: %v", err)
	}

	// Step 4: Decode L2msg from base64 string
	l2MsgData, err := base64.StdEncoding.DecodeString(input.Message.Message.L2msg) // Now correctly decoding the base64 string
	if err != nil {
		log.Fatalf("Error decoding l2Msg: %v", err)
	}

	// Step 5: Parse the transactions
	chainId := big.NewInt(42161)
	txs, err := arbos.ParseL2Transactions(&arbostypes.L1IncomingMessage{
		Header: input.Message.Message.Header,
		L2msg:  l2MsgData, // Now using the decoded L2msg data
	}, chainId)
	if err != nil {
		fmt.Println("Error parsing L2 transactions:", err)
	} else {
		fmt.Println("length of txns:", txs.Len())
		utils.PrintTransaction(txs[0])

		// calculation of tx root
		// startTx, err := util.PackInternalTxDataStartBlock(l1BaseFee, l1BlockNum, l2BlockNum, timePassed)
		// l1BaseFee : extractable from msg feed
		// l1BlockNum : extractable from msg feed
		// l2BlockNum : this can be done by adding the Arbitrum One genesis block number (22207817) to the sequence number of the feed message
		// timePassed : header.Time - lastHeader.Time
		// append(types.Transactions{types.NewTx(startTx)}, txes...)

		lastTimestamp := uint64(1728980884) // 1728980884
		data, err := util.PackInternalTxDataStartBlock(big.NewInt(0), input.Message.Message.Header.BlockNumber, input.SequenceNumber + uint64(22207817), input.Message.Message.Header.Timestamp - lastTimestamp)

		if err != nil {
			// fmt.Println("Error creating Start Transaction")
			panic(fmt.Sprintf("Failed to pack internal tx %v", err))
		}

		startTx := &types.ArbitrumInternalTx{
			ChainId: chainId,
			Data:    data,
		}

		txs = append(types.Transactions{types.NewTx(startTx)}, txs...)
		
		utils.PrintTransactionsRoot(txs)
	}
}