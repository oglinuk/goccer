package hive

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func (q *Queen) SpawnDrone() {
	q.crawl()
	q.rw.Aggregate()
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
		return nil
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		q.pw.Write(q.seed)
		for _, URL := range q.Extract(resp) {
			q.rw.Write(URL)
		}
	} else {
		log.Printf("Request failed: %s", q.seed)
	}

	return nil
}
