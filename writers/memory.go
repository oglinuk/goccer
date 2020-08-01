package writers

// MemoryStore is self explanitory
type MemoryStore struct {
	paths map[string]struct{}
}

// MemoryWriter is a base path and its subpaths
type MemoryWriter struct {
	base         string
	memoryStores map[string]*MemoryStore
}

// NewMemoryStore constructor
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		paths: make(map[string]struct{}),
	}
}

// NewMemoryWriter constructor
func NewMemoryWriter(basePath string) *MemoryWriter {
	return &MemoryWriter{
		base:         basePath,
		memoryStores: make(map[string]*MemoryStore),
	}
}

// Write paths to memory
func (mw *MemoryWriter) Write(paths []string) error {

	return nil
}
