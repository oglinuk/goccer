package hive

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
