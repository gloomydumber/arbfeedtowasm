package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/offchainlabs/nitro/arbos"
	"github.com/offchainlabs/nitro/arbos/arbostypes"
)

func printTransaction(tx *types.Transaction) {
	jsonBytes, err := tx.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal transaction: %v", err)
	}

	fmt.Println(string(jsonBytes))
}

func calculateTransactionsRoot(txs types.Transactions) {
	root := types.DeriveSha(txs, trie.NewStackTrie(nil))
	fmt.Println(root)
}

func main() {
		l2MsgBase64 := "AwAAAAAAAABuBPhrgaSDmJaAgwP3C5Rjh4QWTIIu+dbgs9wi8v8FXgxsHoCEZt8BWoMBSYWgGyIWIElNYM/NFiKm6WTRlIZdJvok0OFLwoDREanI5s6gF7c+cc/rzUV/7jAzSW35ZPLDUXSKg1M7jUfqObNcU5YAAAAAAAAAkQQC+I2CpLGBuYCDzf5ggwKHj5R4ByiJ7k1/4aEAwlKWqrvqMukr6oCkREOInAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABHu7wICgvOUmFkLv7EUi1IU/OtcQJDyV8cHJfntYstww1eaapwygc1Dwp1lHN3BbiXTgk4Z/CuekGVVjwsQPXTMK3g7h3eYAAAAAAAAAdAQC+HCCpLGCAtmAg83+YIMBqmmUXoCahaoYKpkh7dEKQWN0W7PjYoSHBdyqj+EgDYDAAaCjLxkNBk0z5Lahw6OJq2Zt90e4O33ELOlqoHTZnsFlU59iP3WHj6RvumpXeOZyc/P1YkPlItISsA8h8qU2YjCkAAAAAAAAAhUEAvkCEIKksYIV44CDmJaAgwqZ8JTDUscRlDlB1gjQVF7y5v5X2ecBVYC5AaRX7P0oAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEBAWcI020AAAAAAAKGmgAAAPAAAQAAAAAAAAAAAAAAAB8Oo7Y/P8oFcZ5U50ae+Jd1TvZmAAAApKAOqkgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAgAAAAAAAAAAAAAAAIkGCzHbIcbLTpRurLKO/v8IXCdaKEfn8oI6UEj0rizYCKXpeKps5B/Ltufnu7sbZERrBjkAAAAAAAAAAAAAAADTTFgNPRH67mOqAzUgWUjC6bzkXwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAKaIkGvYsAAAAAKF6AACAAAAAAAAAAAAAAAAiQYLMdshxstOlG6sso7+/whcJ1oAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADAgKBzjcghg/XkxCdiM9xYsJypgKgcaqmJIB2GUv0Y6crhEKB9jsdSC+WQPFuoeXrerqLhpTIhpERE4AOQIJdyD/9PbwAAAAAAAAC4BAL4tIKksYILY4R3NZQAhHfOKoCDAqXdlMR5ZvYwoaISccRLMV4LsPXCmsH7gLhEj/6dRgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABFAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAHAgKByVyN6njcUHXBqVg9i5AowcO6bdib9VSD6Ctx0VSRQSqBuYOM8Ne8wOcvv6E718WLAmEmP1g/l6uXRDAi1EP/JIwAAAAAAAAB6BAL4doKksYMiDEmAg+ThwIMPQkCUzZNLgY3dfaSoLNPnPAFFaQJMJviAiwYABAD/BQM3o4HMwICgJ640l/QRP8kUZYZFMKOeRAiqBEZQSndYxI5Rdi/FBBagDDAaHoZGlj9PEH8hvSuBCTf8Gwe5gBgM0m9Y+4+makcAAAAAAAAApAT4oYMQmAWGBDCLYG7TgwprW5TzPxi7wiS082yhXqdId6CXT1wJ44C1AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAuDtGT03B3wAAAAAAAAAAAAAzeje8PIK6Dk1wCDAUmGoB3bWFP+tj1px/08PcudegzGIvLwQDZHWL+fdonr7rKroCb8F57RpQ5CGrxeSY8RYcjta3cuYYVOknsI+LRleWVoAAAAAAAAALIEAviugqSxgICEATEtAIMCxNmUr4jQZed8jMIjkyfF7bOkMiaOWDGAuEQJXqezAAAAAAAAAAAAAAAAKKesZpmPTENJy1rFMCSF+uMytVgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAvrwgMCAoPOBkMPvdxDPPpWE1vygynqlQYbtxN8osq+NTjckhD2xoCanRxYFdu/dZAvF2Epp49ePo6TCR6kle53HnNdVXJEe"

		l2MsgData, err := base64.StdEncoding.DecodeString(l2MsgBase64)
		if err != nil {
			fmt.Println("Error decoding l2Msg:", err)
			return
		}

	msg := &arbostypes.L1IncomingMessage{
		Header: &arbostypes.L1IncomingMessageHeader{
			Kind:        3,
			Poster:      common.HexToAddress("0xa4b000000000000000000073657175656e636572"),
			BlockNumber: 20940885,
			Timestamp:   1728631672,
			RequestId:   nil,
			L1BaseFee:   nil,
		},
		L2msg:        l2MsgData,
		BatchGasCost: nil, // Optional
	}
	
	chainId := big.NewInt(42161)
    
	txs, err := arbos.ParseL2Transactions(msg, chainId)
    if err != nil {
        fmt.Println("Error parsing L2 transactions:", err)
    } else {
		printTransaction((txs[0]))
		calculateTransactionsRoot(txs)
    }
}