package crawlers

// FSCrawler for filesystems
type FSCrawler struct {
	path string
}

// NewFSCrawler constructor
func NewFSCrawler(p string) FSCrawler {
	return FSCrawler{
		path: p,
	}
}

// Crawl c.path
func (c FSCrawler) Crawl() []string {
	return nil
}
