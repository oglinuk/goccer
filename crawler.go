package main

import (
	"fmt"

	"github.com/oglinuk/goccer/crawlers"
)

// Crawler is the base for a thing that crawls the given path
type Crawler interface {
	Crawl()
}

// CreateCrawler of ctype with given path
func CreateCrawler(ctype, path string, w Writer) (Crawler, error) {
	switch ctype {
	case "http":
		return crawlers.NewHTTPCrawler(path, w), nil
	case "fs":
		return crawlers.NewFSCrawler(path), nil
	default:
		return nil, fmt.Errorf("crawler.go::CreateCrawler(%s, ...)::ERROR: Invalid crawler type", ctype)
	}
}
