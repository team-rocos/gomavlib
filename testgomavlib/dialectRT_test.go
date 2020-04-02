package testgomavlib

import (
	"testing"
)

// DEFINE PUBLIC TYPES AND STRUCTURES.

// DEFINE PRIVATE TYPES AND STRUCTURES.
// DEFINE PUBLIC STATIC FUNCTIONS.
func TestCreateMessageById(t *testing.T) {
	CreateMessageByIdTest(t, "../mavlink-upstream/message_definitions/v1.0/common.xml", []string{"../mavlink-upstream/message_definitions/v1.0"})
}

func TestCreateMessageByName(t *testing.T) {
	CreateMessageByNameTest(t, "../mavlink-upstream/message_definitions/v1.0/common.xml", []string{"../mavlink-upstream/message_definitions/v1.0"})
}

func TestJSONMarshalAndUnmarshal(t *testing.T) {
	JSONMarshalAndUnmarshalTest(t, "../mavlink-upstream/message_definitions/v1.0/common.xml", []string{"../mavlink-upstream/message_definitions/v1.0"})
}

func TestDialectRTCommonXML(t *testing.T) {
	DialectRTCommonXMLTest(t, "../mavlink-upstream/message_definitions/v1.0/common.xml", []string{"../mavlink-upstream/message_definitions/v1.0"})
}

func TestDecodeAndEncodeRT(t *testing.T) {
	DecodeAndEncodeRTTest(t, "../mavlink-upstream/message_definitions/v1.0/common.xml", []string{"../mavlink-upstream/message_definitions/v1.0"})
}

// DEFINE PUBLIC RECEIVER FUNCTIONS.
// DEFINE PRIVATE STATIC FUNCTIONS.
// DEFINE PRIVATE RECEIVER FUNCTIONS.
// ALL DONE.
