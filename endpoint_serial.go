package gomavlib

import (
	"fmt"
	"github.com/tarm/serial"
	"io"
	"regexp"
	"strconv"
)

var reSerial = regexp.MustCompile("^(.+?):([0-9]+)$")

// EndpointSerial sets up a endpoint that works through a serial port.
type EndpointSerial struct {
	// the address of the serial port in format name:baudrate
	// example: /dev/ttyUSB0:57600
	Address string
}

type endpointSerial struct {
	io.ReadWriteCloser
}

func (conf EndpointSerial) init() (endpoint, error) {
	matches := reSerial.FindStringSubmatch(conf.Address)
	if matches == nil {
		return nil, fmt.Errorf("invalid address")
	}

	name := matches[1]
	baud, _ := strconv.Atoi(matches[2])

	port, err := serial.OpenPort(&serial.Config{
		Name: name,
		Baud: baud,
	})
	if err != nil {
		return nil, err
	}

	t := &endpointSerial{
		ReadWriteCloser: port,
	}
	return t, nil
}

func (t *endpointSerial) Desc() string {
	return fmt.Sprintf("serial")
}
