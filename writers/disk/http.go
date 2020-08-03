package disk

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

// HTTPStore is a file and a map containing paths in the file
type HTTPStore struct {
	file  *os.File
	paths map[string]struct{}
}

// HTTPWriter is a directory and a map containing Stores
type HTTPWriter struct {
	path       string
	diskStores map[string]*HTTPStore
}

// NewHTTPDiskStore constructor
func NewHTTPDiskStore(fpath string) *HTTPStore {
	file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("writers::disk::fs.go::NewHTTPDiskStore::os.OpenFile(%s)::ERROR: %s", fpath, err.Error())
	}
	return &HTTPStore{
		file:  file,
		paths: make(map[string]struct{}),
	}
}

// NewHTTPDiskWriter constructor
func NewHTTPDiskWriter(dir string) *HTTPWriter {
	if _, err := os.Stat(baseDiskDirName); err != nil {
		if err = os.MkdirAll(baseDiskDirName, 0777); err != nil {
			log.Fatalf("writers::disk::fs.go::NewHTTPDiskWriter::os.MkdirAll(%s)::ERROR: %s", baseDiskDirName, err.Error())
		}
	}

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Printf("writers::disk::fs.go::NewHTTPDiskWriter::os.MkdirAll(%s)::ERROR: %s", dir, err.Error())
	}

	return &HTTPWriter{
		path:       dir,
		diskStores: make(map[string]*HTTPStore),
	}
}

// Write http paths to disk
func (dw *HTTPWriter) Write(paths []string) error {
	for _, p := range paths {
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
			dw.diskStores[base] = NewHTTPDiskStore(fpath)
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
