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
		pw:   newURLWriter("data/crawled"),
		rw:   newURLWriter("data/uncrawled"),
		ew:   newURLWriter("data/errors"),
		aw:   newURLFile(af),
	}
}
