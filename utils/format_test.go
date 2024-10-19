package utils_test

import (
	"testing"

	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/test"
	"arbfeedtowasm/utils"

	"github.com/google/go-cmp/cmp"
)

func TestParseIncomingMessage(t *testing.T) {
    cases := []struct {
        name          string
        feedMessage   string
        expected      feedtypes.IncomingMessage
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
