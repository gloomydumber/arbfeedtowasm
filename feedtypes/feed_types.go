package feedtypes

import (
	"github.com/offchainlabs/nitro/arbos/arbostypes"
)

// TODO: describe godocs here
type IncomingMessage struct {
	Message             Message `json:"message"`
	DelayedMessagesRead uint64  `json:"delayedMessagesRead"`
}

// TODO: describe godocs here
type Message struct {
	Header *arbostypes.L1IncomingMessageHeader `json:"header"`
	L2msg  string                              `json:"l2Msg"` // Base64-encoded string
}