package feedtypes

import (
	"github.com/ethereum/go-ethereum/core/types"
)

// Define a struct for the response, including both success (result) and error cases
type Response struct {
	Result string         `json:"result"`          // Either "success" or "error"
	Data   *ResponseData  `json:"data,omitempty"`  // Only present if result is "success"
	Error  *ErrorResponse `json:"error,omitempty"` // Only present if result is "error"
}

type ResponseData struct {
	Transactions     types.Transactions `json:"transactions"`
	TransactionsRoot string             `json:"transactionsRoot"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
