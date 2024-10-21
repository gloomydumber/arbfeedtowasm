// Package utils provides utility functions for handling Ethereum transactions,
// converting JavaScript-like object strings to valid JSON, and parsing incoming messages
// from the Arbitrum sequencer feed.
//
// This package is designed to support the processing and manipulation of data related
// to Ethereum transactions and feed messages, with a focus on the Arbitrum Layer 2 solution.
// It includes functions for converting input formats, calculating Merkle roots, and parsing
// JSON data into defined Go structures. Additionally, the package defines useful constants
// related to the Arbitrum network.
package utils

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"arbfeedtowasm/feedtypes"
)

// ConvertToJSON converts a JavaScript-like object string to valid JSON format.
//
// This function adds double quotes around keys in a string that is formatted like a JavaScript object
// (without quotes around the keys) and replaces single quotes with double quotes. It uses regular
// expressions to detect keys and applies the appropriate transformation.
//
// Example Input: {sequenceNumber: 123, 'key': 'value'}
// Example Output: {"sequenceNumber": 123, "key": "value"}
//
// Parameters:
// - jsLikeString: A string representing a JavaScript-like object.
//
// Returns:
// - A string with valid JSON format.
func ConvertToJSON(jsLikeString string) string {
	// Regex to add quotes around keys
	re := regexp.MustCompile(`(\b\w+\b)\s*:`)
	result := re.ReplaceAllString(jsLikeString, `"$1":`)
	// Convert single quotes to double quotes, if any
	result = strings.ReplaceAll(result, `'`, `"`)
	return result
}

// IsJSObject determines if a string is formatted like a JavaScript object.
//
// This function checks whether the input string contains keys without double quotes followed by colons,
// which is common in JavaScript object literals (e.g., sequenceNumber: 123). It uses regular expressions
// to match such patterns.
//
// Parameters:
// - input: The string to be checked.
//
// Returns:
// - A boolean value indicating whether the input is a JavaScript-like object (true) or not (false).
func IsJSObject(input string) bool {
	// Look for keys without double quotes, followed by a colon (e.g., sequenceNumber: 123)
	re := regexp.MustCompile(`\b\w+\s*:`)
	return re.MatchString(input)
}

// ParseIncomingMessage parses a JSON string into an IncomingMessage struct.
//
// This function takes a JSON string as input and unmarshals it into an IncomingMessage struct
// defined in the feedtypes package. If the input is invalid JSON, the function logs a fatal error.
//
// Parameters:
// - msg: A JSON string representing an IncomingMessage.
//
// Returns:
// - An IncomingMessage struct populated with the parsed data.
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
