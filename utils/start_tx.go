package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/nitro/arbos/util"
)

// TODO: define legit godoc here
// type of l1BaseFee is *big.Int, make sure this would be applied on also as received websocket message form
// https://github.com/OffchainLabs/nitro/blob/master/arbos/internal_tx.go#L22
// https://github.com/OffchainLabs/nitro/blob/master/arbos/block_processor.go#L200
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
		ChainId: big.NewInt(ArbiturmChainId),
		Data:    data,
	}

	return startTx
}

// TODO: define legit godoc here
func AppendStartTransaction(
	transactions types.Transactions,
	l1BaseFee *big.Int,
	l1BlockNum uint64,
	sequneceNumber uint64,
	currentTimestamp uint64,
	prevTimestamp uint64,
) types.Transactions {
	l2BlockNum := sequneceNumber + ArbiturmGenesisBlockNumber
	timePassed := currentTimestamp - prevTimestamp

	if l1BaseFee == nil {
		l1BaseFee = big.NewInt(0)
	}

	startTx := makeStartTransaction(l1BaseFee, l1BlockNum, l2BlockNum, timePassed)
	transactions = append(types.Transactions{types.NewTx(startTx)}, transactions...)

	return transactions
}
