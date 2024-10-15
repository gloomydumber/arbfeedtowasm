package utils

import (
	"regexp"
	"strings"
)

// TODO: describe godocs here
// Function to convert a JavaScript-like object to valid JSON
func ConvertToJSON(jsLikeString string) string {
	// Regex to add quotes around keys
	re := regexp.MustCompile(`(\w+):`)
	result := re.ReplaceAllString(jsLikeString, `"$1":`)
	result = strings.ReplaceAll(result, `'`, `"`)
	return result
}

// TODO: describe godocs here
// Function to determine if input is JavaScript-like (no quotes around keys or single quotes)
func IsJSLike(input string) bool {
	// If there are no double quotes around any key or string, we assume it's JS-like
	return !regexp.MustCompile(`"\w+"`).MatchString(input) || strings.Contains(input, `'`)
}