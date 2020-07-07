package main

import (
	"fmt"

	"./fs"
	"./http"
)

// Crawler is the base for a thing that crawls the given path
type Crawler interface {
	Crawl() error
}

// CreateCrawler of ctype with given path
func CreateCrawler(ctype, path string) (Crawler, error) {
	switch ctype {
	case "http":
		return http.NewCrawler(path), nil
	case "fs":
		return fs.NewCrawler(path), nil
	default:
		return nil, fmt.Errorf("crawler.go::CreateCrawler(%s, ...)::ERROR: Invalid crawler type", ctype)
	}
}
