package crawlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
	var paths []string

	if stat, err := os.Stat(c.path); err == nil && stat.IsDir() {
		infos, err := ioutil.ReadDir(c.path)
		if err != nil {
			log.Printf("crawlers::fs.go::Crawl::ioutil.ReadDir(%s)::ERROR: %s", c.path, err.Error())
			return nil
		}

		for _, info := range infos {
			path := fmt.Sprintf("%s/%s", c.path, info.Name())
			paths = append(paths, path)
		}
	} else {
		// Do something with the file?
	}

	return paths
}
