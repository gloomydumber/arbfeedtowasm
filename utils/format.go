package utils

import (
	"regexp"
	"strings"
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