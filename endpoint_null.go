package gomavlib

import "fmt"

// this is an example struct that implements io.ReadWriteCloser.
// it does not read anything and prints what it receives.
// the only requirement is that Close() must release Read().
type NullEndpoint struct {
	readChan chan []byte
}

func NewNullEndpoint() *NullEndpoint {
	return &NullEndpoint{
		readChan: make(chan []byte),
	}
}

func (c *NullEndpoint) Close() error {
	close(c.readChan)
	return nil
}

func (c *NullEndpoint) Read(buf []byte) (int, error) {
	read, ok := <-c.readChan
	if ok == false {
		return 0, fmt.Errorf("all right")
	}
	n := copy(buf, read)
	return n, nil
}

func (c *NullEndpoint) Write(buf []byte) (int, error) {
	return len(buf), nil
}
