package hive

type Queen struct {
	seed string
	pw   *URLWriter
	rw   *URLWriter
}

func NewQueen(URL string) *Queen {
	return &Queen{
		seed: URL,
		pw:   NewURLWriter("crawled"),
		rw:   NewURLWriter("uncrawled"),
	}
}
