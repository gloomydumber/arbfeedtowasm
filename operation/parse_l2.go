package operation

import (
	"encoding/base64"
	"log"
	"math/big"

	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/utils"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/nitro/arbos"
	"github.com/offchainlabs/nitro/arbos/arbostypes"
)

// TODO: describe godocs here
// DecodeL2Message decodes a base64-encoded L2 message and handles errors internally.
func decodeL2Message(l2MsgBase64 string) []byte {
	// Decode the base64-encoded string
	l2MsgData, err := base64.StdEncoding.DecodeString(l2MsgBase64)
	if err != nil {
		log.Fatalf("Error decoding l2Msg: %v", err)
	}

	// Return the decoded data
	return l2MsgData
}

// TODO: describe godocs here
func ParseL2Transactions(msg feedtypes.IncomingMessage) types.Transactions {
	l2MsgData := decodeL2Message(msg.Message.Message.L2msg)

	txns, err := arbos.ParseL2Transactions(&arbostypes.L1IncomingMessage{
		Header: msg.Message.Message.Header,
		L2msg:  l2MsgData,
	}, big.NewInt(utils.ArbiturmChainId))

	if err != nil {
		log.Fatalf("Error parsing txns: %v", err)
	}

	return txns
}
