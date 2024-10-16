package utils

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"arbfeedtowasm/feedtypes"
)

// TODO: describe godocs here
// Function to convert a JavaScript-like object to valid JSON
func ConvertToJSON(jsLikeString string) string {
	// Regex to add quotes around keys
	re := regexp.MustCompile(`(\b\w+\b)\s*:`)
	result := re.ReplaceAllString(jsLikeString, `"$1":`)
	// Convert single quotes to double quotes, if any
	result = strings.ReplaceAll(result, `'`, `"`)
	return result
}

// TODO: describe godocs here
// Function to determine if input is JavaScript-like (no quotes around keys or single quotes)
func IsJSObject(input string) bool {
	// Look for keys without double quotes, followed by a colon (e.g., sequenceNumber: 123)
	re := regexp.MustCompile(`\b\w+\s*:`)
	return re.MatchString(input)
}

// TODO: describe godocs here
// ParseIncomingMessage parses the JSON data into an IncomingMessage struct
func ParseIncomingMessage(msg string) feedtypes.IncomingMessage {
	var parsedMsg feedtypes.IncomingMessage

	// Unmarshal JSON data into the IncomingMessage struct
	err := json.Unmarshal([]byte(msg), &parsedMsg)
	if err != nil {
		log.Fatalf("Failed to parse JSON input: %v", err)
	}

	// Return the parsed IncomingMessage struct
	return parsedMsg
}

// TODO: describe godocs here
// DecodeL2Message decodes a base64-encoded L2 message and handles errors internally.
func DecodeL2Message(l2MsgBase64 string) []byte {
	// Decode the base64-encoded string
	l2MsgData, err := base64.StdEncoding.DecodeString(l2MsgBase64)
	if err != nil {
		log.Fatalf("Error decoding l2Msg: %v", err)
	}

	// Return the decoded data
	return l2MsgData
}