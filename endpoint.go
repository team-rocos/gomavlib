package gomavlib

import (
	"io"
)

// EndpointChannel is a channel provided by a endpoint.
type EndpointChannel struct {
	desc      string
	rwc       io.ReadWriteCloser
	writeChan chan interface{}
}

// String implements fmt.Stringer and returns infos about channel.
func (e *EndpointChannel) String() string {
	return e.desc
}

// EndpointConf is the interface implemented by all endpoints.
type EndpointConf interface {
	init() (endpoint, error)
}

// a endpoint must also implement one of the following:
// - endpointChannelSingle
// - endpointChannelAccepter
type endpoint interface{}

// endpoint that provides a single channel.
// Read() must not return any error unless Close() is called.
type endpointChannelSingle interface {
	endpoint
	Desc() string
	io.ReadWriteCloser
}

// endpoint that provides multiple channels.
type endpointChannelAccepter interface {
	endpoint
	Close() error
	Accept() (endpointChannelSingle, error)
}
