package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// URLFile of a domain subdomains as its content
type URLFile struct {
	file *os.File
	urls map[string]struct{}
}

// URLWriter is a directory containing URLFiles
type URLWriter struct {
	path     string
	filters  []string
	urlFiles map[string]*URLFile
}

// NewURLFile constructor
func NewURLFile(fPath string) *URLFile {
	file, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("os.OpenFile (newUrlFile) err: %s", err.Error())
	}
	return &URLFile{
		file: file,
		urls: make(map[string]struct{}),
	}
}

//NewURLWriter constructor
func NewURLWriter(dPath string, filts []string) *URLWriter {
	err := os.MkdirAll(dPath, 0777)
	if err != nil {
		log.Printf("os.MkdirAll (newUrlWriter) err: %s", err.Error())
	}
	return &URLWriter{
		path:     dPath,
		filters:  filts,
		urlFiles: make(map[string]*URLFile),
	}
}

// write checks if the URL base exists as a URLFile; if not create
// then checks if the URL exists in the URLFile; if not write it
func (uw *URLWriter) write(URL string) error {
	if uw.filter(URL) {
		return nil
	}

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
		uw.urlFiles[base] = NewURLFile(fileDir)
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

// filter check if contains one of the filters
func (uw *URLWriter) filter(check string) bool {
	for _, f := range uw.filters {
		if strings.Contains(strings.ToLower(check), f) {
			return true
		}
	}
	return false
}
