package utils_test

import (
	"reflect"
	"testing"

	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/test"
	"arbfeedtowasm/utils"
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
            if !reflect.DeepEqual(parsedMessage, tc.expected) {
                t.Errorf("Wrong result for %s: got %+v, expected %+v", tc.name, parsedMessage, tc.expected)
            }
        })
    }
}
