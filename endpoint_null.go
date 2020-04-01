package gomavlib

import "fmt"

// NullEndpoint implements io.ReadWriteCloser.
// it does not read anything and prints what it receives.
// the only requirement is that Close() must release Read().
type NullEndpoint struct {
	readChan chan []byte
}

// NewNullEndpoint returns a pointer to a new NullEndpoint
func NewNullEndpoint() *NullEndpoint {
	return &NullEndpoint{
		readChan: make(chan []byte),
	}
}

// Close closes the channel c.readChan
func (c *NullEndpoint) Close() error {
	close(c.readChan)
	return nil
}

// Read function of NullEndpoint
func (c *NullEndpoint) Read(buf []byte) (int, error) {
	read, ok := <-c.readChan
	if ok == false {
		return 0, fmt.Errorf("failed to read from c.readChan")
	}
	n := copy(buf, read)
	return n, nil
}

// Write function of NullEndpoint
func (c *NullEndpoint) Write(buf []byte) (int, error) {
	return len(buf), nil
}
