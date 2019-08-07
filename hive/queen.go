package hive

type Queen struct {
	seed string
	pw   *URLWriter
	rw   *URLWriter
	aw   *URLFile
}

func NewQueen(URL string) *Queen {
	return &Queen{
		seed: URL,
		pw:   NewURLWriter("crawled"),
		rw:   NewURLWriter("uncrawled"),
		aw:   NewURLFile("to_crawl.txt"),
	}
}
