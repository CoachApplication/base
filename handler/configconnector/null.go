package configconnector

type NullConfigConnector struct {}

// List lists all configs
func (ncc *NullConfigConnector) List() []string {
	return []string{}
}

// ListScopes list all scopes for a config
func (ncc *NullConfigConnector) ListScopes(key string) []string {
	return []string{}
}

// Get retrieves ScopedConfig for the Config Get operation
func (ncc *NullConfigConnector) Get(key string, scope string) (io.ReaderCloser, error) {
	return NewBufferReaderCloser([]byte{}), nil
}

// Set Pushes provide ScopedConfig to a 
func (ncc *NullConfigConnector) Set(key string, scope string, io.ReaderCloser) error {
	return nil
}

/**
 * Buffer based ReaderCloser
 */

// BufferReaderCloser a bytes Buffer wrapper that adds a Close()
type BufferReaderCloser struct {
	bytes.Buffer
}

func NewBufferReaderCloser(b []byte) *BufferReaderCloser {
	return &BufferReaderCloser{
		Buffer: bytes.NewBuffer(b)
	}
}


// Close emulate a Close() to implement ReaderCloser
func (brc *BufferReaderCloser) Close() error {
	return nil
}
