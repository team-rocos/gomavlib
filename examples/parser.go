// +build ignore

package main

import (
	"bytes"
	"fmt"

	"github.com/team-rocos/gomavlib"
	"github.com/team-rocos/gomavlib/dialects/ardupilotmega"
)

func main() {
	inBuf := bytes.NewBuffer(
		[]byte("\xfd\t\x01\x00\x00\x00\x00\x00\x00\x00\x04\x00\x00\x00\x01\x02\x03\x05\x03\xd9\xd1\x01\x02\x00\x00\x00\x00\x00\x0eG\x04\x0c\xef\x9b"))
	outBuf := bytes.NewBuffer(nil)

	// if NewNode() is not flexible enough, the library provides a low-level Mavlink
	// frame parser, that can be allocated with NewParser().
	parser, err := gomavlib.NewParser(gomavlib.ParserConf{
		Reader:      inBuf,
		Writer:      outBuf,
		Dialect:     ardupilotmega.Dialect,
		OutVersion:  gomavlib.V2, // change to V1 if you're unable to write to the target
		OutSystemId: 10,
	})
	if err != nil {
		panic(err)
	}

	// read a message, encapsulated in a frame
	frame, err := parser.Read()
	if err != nil {
		panic(err)
	}

	fmt.Printf("decoded: %+v\n", frame)

	// write a message
	err = parser.WriteMessage(&ardupilotmega.MessageParamValue{
		ParamId:    "test_parameter",
		ParamValue: 123456,
		ParamType:  ardupilotmega.MAV_PARAM_TYPE_UINT32,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("encoded: %v\n", outBuf.Bytes())
}
