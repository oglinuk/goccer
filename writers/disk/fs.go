package disk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// FsStore is a file and a map containing paths in the file
type FsStore struct {
	file  *os.File
	paths map[string]struct{}
}

// FsWriter is a directory and a map containing Stores
type FsWriter struct {
	path       string
	diskStores map[string]*FsStore
}

// NewFsDiskStore constructor
func NewFsDiskStore(fpath string) *FsStore {
	file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("writers::disk.go::NewDiskWriter::os.OpenFile(%s)::ERROR: %s", fpath, err.Error())
	}
	return &FsStore{
		file:  file,
		paths: make(map[string]struct{}),
	}
}

// NewFsDiskWriter constructor
func NewFsDiskWriter(dir string) *FsWriter {
	if _, err := os.Stat(baseDiskDirName); err != nil {
		if err = os.MkdirAll(baseDiskDirName, 0777); err != nil {
			log.Fatalf("writers::disk.go::NewDiskWriter::os.MkdirAll(%s)::ERROR: %s", baseDiskDirName, err.Error())
		}
	}

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Printf("writers::disk.go::NewDiskWriter::os.MkdirAll(%s)::ERROR: %s", dir, err.Error())
	}

	return &FsWriter{
		path:       dir,
		diskStores: make(map[string]*FsStore),
	}
}

// Write fs paths to disk
func (dw *FsWriter) Write(paths []string) error {
	for _, p := range paths {
		cleaned := filepath.Clean(p)

		base := filepath.Base(cleaned)
		if base == "" || base == " " {
			base = "error"
		}
		fpath := filepath.Join(dw.path, base)

		if _, ok := dw.diskStores[base]; !ok {
			dw.diskStores[base] = NewFsDiskStore(fpath)
		}

		ds := dw.diskStores[base]

		if _, exists := ds.paths[cleaned]; !exists {
			ds.file.WriteString(fmt.Sprintf("%s\n", cleaned))
			ds.paths[cleaned] = struct{}{}
		}

		ds.file.Close()
	}

	return nil
}
