package hive

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
)

type URLFile struct {
	file       *os.File
	hasWritten bool
}

type URLWriter struct {
	path     string
	urlFiles map[string]*URLFile
}

func NewURLFile(fPath string) *URLFile {
	file, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("os.OpenFile NewUrlFile err: %s", err.Error())
	}
	return &URLFile{
		file:       file,
		hasWritten: false,
	}
}

func NewURLWriter(dPath string) *URLWriter {
	err := os.MkdirAll(dPath, 0777)
	if err != nil {
		log.Printf("os.MkdirAll NewUrlWriter err: %s", err.Error())
	}
	return &URLWriter{
		path:     dPath,
		urlFiles: make(map[string]*URLFile),
	}
}

func (uw *URLWriter) Write(URL string) error {
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

	uf.file.WriteString(fmt.Sprintf("%s\n", decoded))
	uf.hasWritten = true

	return nil
}

func (uw *URLWriter) Aggregate() {
	check := make(map[string]struct{})
	//var uncrawled []string

	f, err := os.OpenFile("to_crawl.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		log.Printf("os.OpenFile NewUrlFile err: %s", err.Error())
	}

	for k := range uw.urlFiles {
		fileDir := filepath.Join(uw.path, k)

		fd, err := os.Open(fileDir)
		if err != nil {
			log.Printf("os.Open Aggregate err: %s", err.Error())
		}
		defer fd.Close()

		scanner := bufio.NewScanner(fd)
		for scanner.Scan() {
			scanned := scanner.Text()
			if _, exists := check[scanned]; !exists {
				check[scanned] = struct{}{}
				f.WriteString(fmt.Sprintf("%s\n", scanned))
				//uncrawled = append(uncrawled, scanner.Text())
			}
		}
	}
	/*
		utils.SaveConfig(&utils.Config{
			MaxWorkers: 4,
			Seeds:      uncrawled,
		})
	*/
}
