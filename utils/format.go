package utils

import (
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
