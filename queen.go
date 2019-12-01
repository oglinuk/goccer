package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Queen is the crawling manager
type Queen struct {
	seed string
	pw   *URLWriter
	rw   *URLWriter
	ew   *URLWriter
	aw   *URLFile
}

// NewQueen constructor
func NewQueen(s string) *Queen {
	return &Queen{
		seed: s,
		pw:   NewURLWriter("data/crawled"),
		rw:   NewURLWriter("data/uncrawled"),
		ew:   NewURLWriter("data/errors"),
		aw:   NewURLFile(archiveName),
	}
}

// SpawnDrone crawls q.seed, then writes each uncrawled URL to the archive file
func (q *Queen) SpawnDrone() {
	if err := q.crawl(); err != nil {
		log.Printf("crawl err: %s", err)
		q.ew.write(q.seed)
	}

	for k := range q.rw.urlFiles {
		fileDir := filepath.Join(q.rw.path, k)

		fd, err := os.Open(fileDir)
		if err != nil {
			log.Printf("os.Open (Aggregate) err: %s", err.Error())
		}

		scanner := bufio.NewScanner(fd)
		for scanner.Scan() {
			scanned := scanner.Text()
			q.aw.file.WriteString(fmt.Sprintf("%s\n", scanned))
		}
		fd.Close()
	}
}

// crawl creates an http client, make a Get request, extracts hyperlinks,
// writes uncrawled URLs to q.rw and q.seed to q.pw
func (q *Queen) crawl() error {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 30,
	}

	resp, err := client.Get(q.seed)
	if err != nil {
		return err
	}

	if resp == nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		q.pw.write(q.seed)
		for _, URL := range q.extract(resp) {
			q.rw.write(URL)
		}
	} else {
		return err
	}

	return nil
}
