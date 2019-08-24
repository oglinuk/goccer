package hive

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

func (q *Queen) spawnDrone() {
	if err := q.crawl(); err != nil {
		log.Printf("crawl err: %s", err)
		q.ew.write(q.seed)
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
