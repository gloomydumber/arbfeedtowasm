package feedtypes

import (
	"github.com/offchainlabs/nitro/arbos/arbostypes"
)

// broadcaster

// TODO: describe godocs here
type IncomingMessage struct {
	SequenceNumber      uint64  `json:"sequenceNumber"`
	Message             MessageWrapper `json:"message"`
	Signature           *string `json:"signature"`
}

// TODO: describe godocs here
type MessageWrapper struct {
	Message             Message `json:"message"`
	DelayedMessagesRead uint64  `json:"delayedMessagesRead"`
}

// TODO: describe godocs here
type Message struct {
	Header *arbostypes.L1IncomingMessageHeader `json:"header"`
	L2msg  string                              `json:"l2Msg"`
}