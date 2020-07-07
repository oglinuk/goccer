package fs

// Crawler for filesystems
type Crawler struct {
	path string
}

// NewCrawler constructor
func NewCrawler(p string) Crawler {
	return Crawler{
		path: p,
	}
}

// Crawl c.path
func (c Crawler) Crawl() error {
	return nil
}
