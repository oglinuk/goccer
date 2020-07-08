package crawlers

import "log"

// Crawler is the base for a thing that crawls the given path
type Crawler interface {
	Crawl() []string
}

// CreateCrawler of ctype with given path
func CreateCrawler(ctype, path string) Crawler {
	switch ctype {
	case "http":
		return NewHTTPCrawler(path)
	case "fs":
		return NewFSCrawler(path)
	default:
		log.Printf("crawler.go::CreateCrawler(%s, ...)::ERROR: Invalid crawler type", ctype)
		return nil
	}
}
