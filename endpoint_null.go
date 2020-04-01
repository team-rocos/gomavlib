package gomavlib

// NullEndpoint implements io.ReadWriteCloser.
// it does not read or write anything.
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
	return 0, nil
}

// Write function of NullEndpoint
func (c *NullEndpoint) Write(buf []byte) (int, error) {
	return 0, nil
}
