package writers

// MemoryStore is self explanitory
type MemoryStore struct {
	paths map[string]struct{}
}

// MemoryWriter is a base path and its subpaths
type MemoryWriter struct {
	base         string
	filts        []string
	memoryStores map[string]*MemoryStore
}

// NewMemoryStore constructor
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		paths: make(map[string]struct{}),
	}
}

// NewMemoryWriter constructor
func NewMemoryWriter(basePath string, filters []string) *MemoryWriter {
	return &MemoryWriter{
		base:         basePath,
		filts:        filters,
		memoryStores: make(map[string]*MemoryStore),
	}
}

func (mw *MemoryWriter) Write(path []string) error {

	return nil
}
