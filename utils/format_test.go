package utils_test

import (
	"reflect"
	"testing"

	"arbfeedtowasm/feedtypes"
	"arbfeedtowasm/test"
	"arbfeedtowasm/utils"
)

func TestParseIncomingMessage(t *testing.T) {
	var feedMessage string
	var parsedMessage feedtypes.IncomingMessage

	// Input format as JavaScript Object Case:
	feedMessage = test.ExampleFeedMessageJSObject
	if utils.IsJSObject(feedMessage) {
		feedMessage = utils.ConvertToJSON(feedMessage)
	}
	parsedMessage = utils.ParseIncomingMessage(feedMessage)
	if !reflect.DeepEqual(parsedMessage, test.ExampleParsedMessage) {
		t.Error("Wrong result: parsedMessage does not match ExampleParsedMessage")
	}

	// Input format as JSON parsed Case:
	feedMessage = test.ExampleFeedMessageJSON
	if utils.IsJSObject(feedMessage) {
		feedMessage = utils.ConvertToJSON(feedMessage)
	}
	parsedMessage = utils.ParseIncomingMessage(feedMessage)
	if !reflect.DeepEqual(parsedMessage, test.ExampleParsedMessage) {
		t.Error("Wrong result: parsedMessage does not match ExampleParsedMessage")
	}

	// Input format as stringified Case:
	feedMessage = test.ExampleFeedMessageString
	if utils.IsJSObject(feedMessage) {
		feedMessage = utils.ConvertToJSON(feedMessage)
	}
	parsedMessage = utils.ParseIncomingMessage(feedMessage)
	if !reflect.DeepEqual(parsedMessage, test.ExampleParsedMessage) {
		t.Error("Wrong result: parsedMessage does not match ExampleParsedMessage")
	}
}
