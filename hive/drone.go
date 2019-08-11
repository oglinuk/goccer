package hive

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func (q *Queen) SpawnDrone() {
	if err := q.crawl(); err != nil {
		log.Printf("Crawl err: %s", err)
		q.ew.Write(q.seed)
	}
}

func (q *Queen) crawl() error {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 5,
	}

	resp, err := client.Get(q.seed)
	if err != nil {
		return err
	}

	if resp == nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		q.pw.Write(q.seed)
		for _, URL := range q.Extract(resp) {
			q.rw.Write(URL)
		}
	} else {
		return err
	}

	return nil
}
