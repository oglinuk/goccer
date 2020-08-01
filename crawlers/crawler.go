package crawlers

import "log"

// Crawler is the base for a thing that crawls the given path
type Crawler interface {
	Crawl() []string
}

// CreateCrawler of ctype with given path
func CreateCrawler(ctype, path string) Crawler {
	var c Crawler

	switch ctype {
	case "http":
		c = NewHTTPCrawler(path)
	case "fs":
		c = NewFSCrawler(path)
	default:
		log.Fatalf("crawler.go::CreateCrawler(%s, ...)::ERROR: Invalid crawler type", ctype)
	}

	return c
}
