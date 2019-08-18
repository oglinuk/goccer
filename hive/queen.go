package hive

type Queen struct {
	seed string
	pw   *URLWriter
	rw   *URLWriter
	ew   *URLWriter
	aw   *URLFile
}

func NewQueen(URL, af string) *Queen {
	return &Queen{
		seed: URL,
		pw:   NewURLWriter("crawled"),
		rw:   NewURLWriter("uncrawled"),
		ew:   NewURLWriter("errors"),
		aw:   NewURLFile(af),
	}
}
