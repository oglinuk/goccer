package hive

type Queen struct {
	seed string
}

func NewQueen(URL string) *Queen {
	return &Queen{
		seed: URL,
	}
}
