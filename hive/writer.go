package hive

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

type URLFile struct {
	file *os.File
	urls map[string]struct{}
}

type URLWriter struct {
	path     string
	urlFiles map[string]*URLFile
}

func newURLFile(fPath string) *URLFile {
	file, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("os.OpenFile (newUrlFile) err: %s", err.Error())
	}
	return &URLFile{
		file: file,
		urls: make(map[string]struct{}),
	}
}

func newURLWriter(dPath string) *URLWriter {
	err := os.MkdirAll(dPath, 0777)
	if err != nil {
		log.Printf("os.MkdirAll (newUrlWriter) err: %s", err.Error())
	}
	return &URLWriter{
		path:     dPath,
		urlFiles: make(map[string]*URLFile),
	}
}

func (uw *URLWriter) write(URL string) error {
	u, err := url.Parse(URL)
	if err != nil {
		return err
	}

	base := u.Hostname()
	if base == "" || base == " " {
		base = "error"
	}
	fileDir := filepath.Join(uw.path, base)

	if _, ok := uw.urlFiles[base]; !ok {
		uw.urlFiles[base] = newURLFile(fileDir)
	}

	uf := uw.urlFiles[base]

	decoded, err := url.QueryUnescape(URL)
	if err != nil {
		return err
	}

	if _, exists := uf.urls[decoded]; !exists {
		uf.file.WriteString(fmt.Sprintf("%s\n", decoded))
		uf.urls[decoded] = struct{}{}
	}

	uf.file.Close()

	return nil
}
