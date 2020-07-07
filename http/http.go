package http

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"
)

// Crawler for HTTP URLs
type Crawler struct {
	seed string
}

// NewCrawler constructor
func NewCrawler(s string) Crawler {
	return Crawler{
		seed: s,
	}
}

// Crawl c.seed and extracts all URLs from the response
func (c Crawler) Crawl() error {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * 7,
	}

	resp, err := client.Get(c.seed)
	if err != nil {
		return err
	}

	if resp == nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		for _, URL := range c.extract(resp) {
			// TODO: Replace below with refactored http writer implementation
			log.Println(URL)
		}
	} else {
		return err
	}

	return nil
}
