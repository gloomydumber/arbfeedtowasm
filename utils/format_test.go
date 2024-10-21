package utils_test

import (
	"testing"

	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/test"
	"arbfeedtowasm/utils"

	"github.com/google/go-cmp/cmp"
)

// TestParseIncomingMessage tests the ParseIncomingMessage function with different input formats.
//
// This test checks that the ParseIncomingMessage function correctly parses different representations
// of incoming messages into the expected IncomingMessage struct. The test cases cover three formats:
// - A JavaScript-like object (without quotes around keys).
// - A valid JSON string.
// - A stringified object representation.
//
// For each test case, the input message is checked to determine if it is a JavaScript-like object
// using the IsJSObject function. If it is, the message is converted to valid JSON using ConvertToJSON.
// The converted or original message is then parsed using ParseIncomingMessage and compared against
// the expected result using cmp.Diff to identify any discrepancies.
func TestParseIncomingMessage(t *testing.T) {
	cases := []struct {
		name        string
		feedMessage string
		expected    feedtypes.IncomingMessage
	}{
		{
			name:        "JavaScript Object Case",
			feedMessage: test.ExampleFeedMessageJSObject,
			expected:    test.ExampleParsedMessage,
		},
		{
			name:        "JSON Case",
			feedMessage: test.ExampleFeedMessageJSON,
			expected:    test.ExampleParsedMessage,
		},
		{
			name:        "Stringified Case",
			feedMessage: test.ExampleFeedMessageString,
			expected:    test.ExampleParsedMessage,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			feedMessage := tc.feedMessage
			if utils.IsJSObject(feedMessage) {
				feedMessage = utils.ConvertToJSON(feedMessage)
			}
			parsedMessage := utils.ParseIncomingMessage(feedMessage)
			if diff := cmp.Diff(tc.expected, parsedMessage); diff != "" {
				t.Errorf("Test %s failed (-expected +got):\n%s", tc.name, diff)
			}
		})
	}
}
