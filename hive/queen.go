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
		pw:   newURLWriter("crawled"),
		rw:   newURLWriter("uncrawled"),
		ew:   newURLWriter("errors"),
		aw:   newURLFile(af),
	}
}
