package writers

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// DiskStore is self explanitory
type DiskStore struct {
	file  *os.File
	paths map[string]struct{}
}

// DiskWriter is a directory containing DiskStores
type DiskWriter struct {
	path       string
	filters    []string
	diskStores map[string]*DiskStore
}

// NewDiskStore constructor
func NewDiskStore(fPath string) *DiskStore {
	file, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("os.OpenFile (NewDiskStore) err: %s", err.Error())
	}
	return &DiskStore{
		file:  file,
		paths: make(map[string]struct{}),
	}
}

// NewDiskWriter constructor
func NewDiskWriter(dir string, filts []string) *DiskWriter {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Printf("os.MkdirAll (NewDiskWriter) err: %s", err.Error())
	}

	return &DiskWriter{
		path:       dir,
		filters:    filts,
		diskStores: make(map[string]*DiskStore),
	}
}

// Write paths to disk
func (dw *DiskWriter) Write(paths []string) error {
	for _, p := range paths {
		for _, f := range dw.filters {
			if strings.Contains(strings.ToLower(p), f) {
				return nil
			}
		}

		u, err := url.Parse(p)
		if err != nil {
			return err
		}

		base := u.Hostname()
		if base == "" || base == " " {
			base = "error"
		}
		fpath := filepath.Join(dw.path, base)

		if _, ok := dw.diskStores[base]; !ok {
			dw.diskStores[base] = NewDiskStore(fpath)
		}

		ds := dw.diskStores[base]

		decoded, err := url.QueryUnescape(p)
		if err != nil {
			return err
		}

		if _, exists := ds.paths[decoded]; !exists {
			ds.file.WriteString(fmt.Sprintf("%s\n", decoded))
			ds.paths[decoded] = struct{}{}
		}

		ds.file.Close()
	}

	return nil
}
