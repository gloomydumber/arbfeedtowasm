package feedtypes

import (
	"github.com/offchainlabs/nitro/arbos/arbostypes"
)

// Verify if the struct type matches the actual code and documentation for reading sequencer feeds.
// Refer to the following links for more details:
// https://github.com/OffchainLabs/nitro/blob/fc2c2d806257b94b03bc57f5e4a3f7717c8c6011/broadcaster/message/message.go#L34
// https://docs.arbitrum.io/run-arbitrum-node/sequencer/read-sequencer-feed
type IncomingMessage struct {
	SequenceNumber uint64         `json:"sequenceNumber"`
	Message        MessageWrapper `json:"message"`
	Signature      *string        `json:"signature"`
}

type MessageWrapper struct {
	Message             Message `json:"message"`
	DelayedMessagesRead uint64  `json:"delayedMessagesRead"`
}

type Message struct {
	Header *arbostypes.L1IncomingMessageHeader `json:"header"`
	L2msg  string                              `json:"l2Msg"`
	// Only used for `L1MessageType_BatchPostingReport`
	BatchGasCost *uint64 `json:"batchGasCost,omitempty" rlp:"optional"`
}
