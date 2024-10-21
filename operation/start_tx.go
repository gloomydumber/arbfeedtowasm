// Package operation provides functions to parse and manipulate Arbitrum Layer 2 (L2) transactions,
// including calculating the Merkle root hash and handling 'Start Block Transactions' which are
// required for block initialization.
//
// This package is designed to work with the Arbitrum sequencer feed, decoding and processing
// transactions in a manner compatible with the Arbitrum Nitro framework.
package operation

import (
	"fmt"
	"math/big"

	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/utils"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/nitro/arbos/util"
)

// makeStartTransaction creates a 'Start Block Transaction' which is an ArbitrumInternalTx.
//
// This transaction is generated at the beginning of every block by the Arbitrum Sequencer
// and is essential for block initialization. The input parameters are as follows:
//
//   - l1BaseFee: The base fee for the L1 transaction. This should be provided as a *big.Int and
//     is also received in websocket message form.
//   - l1BlockNum: The L1 block number at which the transaction is created.
//   - l2BlockNum: The L2 block number corresponding to the transaction being processed.
//   - timePassed: The time difference between the current L2 block and the previous L2 block.
//
// This function packs the input parameters into the internal transaction data format and
// creates an ArbitrumInternalTx. If an error occurs during packing, the function will panic.
//
// For reference, see the Arbitrum Nitro internal transaction structure:
// https://github.com/OffchainLabs/nitro/blob/5f24df4740bdbf3f0dd07aa32ead06c2919087e4/arbos/internal_tx.go#L22
// https://github.com/OffchainLabs/nitro/blob/5f24df4740bdbf3f0dd07aa32ead06c2919087e4/arbos/block_processor.go#L200
func makeStartTransaction(
	l1BaseFee *big.Int,
	l1BlockNum uint64,
	l2BlockNum uint64,
	timePassed uint64,
) *types.ArbitrumInternalTx {
	data, err := util.PackInternalTxDataStartBlock(
		l1BaseFee,
		l1BlockNum,
		l2BlockNum,
		timePassed,
	)

	if err != nil {
		panic(fmt.Sprintf("Failed to pack internal tx %v", err))
	}

	startTx := &types.ArbitrumInternalTx{
		ChainId: big.NewInt(utils.ArbiturmChainId),
		Data:    data,
	}

	return startTx
}

// AppendStartTransaction appends the 'Start Block Transaction' to the list of L2 transactions.
//
// The 'Start Block Transaction' is created using information from the incoming Arbitrum sequencer
// feed message and is manually appended to the list of transactions since this transaction is not
// included in the sequencer feed.
//
// Parameters:
// - transactions: The current list of L2 transactions parsed from the message.
// - msg: The incoming message from the Arbitrum sequencer feed, containing the block's metadata.
// - lastTimestamp: The timestamp of the previous block, used to calculate the time passed.
//
// The function retrieves necessary information such as the L1 base fee, block numbers, and time passed
// from the message, then calls `makeStartTransaction` to create the 'Start Block Transaction'.
// This transaction is then appended to the very beginning of the transaction list and returned.
func AppendStartTransaction(
	transactions types.Transactions,
	msg feedtypes.IncomingMessage,
	lastTimestamp uint64,
) types.Transactions {
	l1BaseFee := msg.Message.Message.Header.L1BaseFee
	l1BlockNum := msg.Message.Message.Header.BlockNumber
	l2BlockNum := msg.SequenceNumber + utils.ArbiturmGenesisBlockNumber
	timePassed := msg.Message.Message.Header.Timestamp - lastTimestamp

	if l1BaseFee == nil {
		l1BaseFee = big.NewInt(0)
	}

	startTx := makeStartTransaction(l1BaseFee, l1BlockNum, l2BlockNum, timePassed)
	transactions = append(types.Transactions{types.NewTx(startTx)}, transactions...)

	return transactions
}
