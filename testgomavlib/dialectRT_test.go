package testgomavlib

import (
	"testing"
)

// DEFINE PUBLIC TYPES AND STRUCTURES.

// DEFINE PRIVATE TYPES AND STRUCTURES.
// DEFINE PUBLIC STATIC FUNCTIONS.
func TestCreateMessageById(t *testing.T) {
	CreateMessageByIdTest(t)
}

func TestCreateMessageByName(t *testing.T) {
	CreateMessageByNameTest(t)
}

func TestJSONMarshalAndUnmarshal(t *testing.T) {
	JSONMarshalAndUnmarshalTest(t)
}

func TestDialectRTCommonXML(t *testing.T) {
	DialectRTCommonXMLTest(t)
}

func TestDecodeAndEncodeRT(t *testing.T) {
	DecodeAndEncodeRTTest(t)
}

// DEFINE PUBLIC RECEIVER FUNCTIONS.
// DEFINE PRIVATE STATIC FUNCTIONS.
// DEFINE PRIVATE RECEIVER FUNCTIONS.
// ALL DONE.
