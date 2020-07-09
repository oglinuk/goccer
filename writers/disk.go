package writers

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var (
	baseDiskDirName = "data"
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
// TODO: Figure out why there is an inconsistent number of files when crawling
// the two default config.json seeds
func NewDiskStore(fpath string) *DiskStore {
	file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("crawlers::disk.go::NewDiskWriter::os.OpenFile(%s)::ERROR: %s",
			fpath, err.Error())
	}
	return &DiskStore{
		file:  file,
		paths: make(map[string]struct{}),
	}
}

// NewDiskWriter constructor
func NewDiskWriter(dir string, filts []string) *DiskWriter {
	if _, err := os.Stat(baseDiskDirName); err != nil {
		if err = os.MkdirAll(baseDiskDirName, 0777); err != nil {
			log.Fatalf("crawlers::disk.go::NewDiskWriter::os.MkdirAll(%s)::ERROR: %s",
				baseDiskDirName, err.Error())
		}
	}

	err := os.MkdirAll(dir, 0777)
	if err != nil {
		log.Printf("crawlers::disk.go::NewDiskWriter::os.MkdirAll(%s)::ERROR: %s",
			dir, err.Error())
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
