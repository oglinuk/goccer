package writers

import (
	"log"

	"github.com/oglinuk/goccer/writers/disk"
)

// Writer is a thing that writes the given path to somewhere
type Writer interface {
	Write(path []string) error
}

// CreateWriter of wtype
func CreateWriter(ctype, wtype, path string) Writer {
	var w Writer

	switch wtype {
	case "disk":
		w = disk.CreateDiskWriter(ctype, path)
	case "memory":
		w = NewMemoryWriter(path)
	default:
		log.Fatalf("writer.go::CreateWriter(%s, ...)::ERROR: Invalid crawler type", wtype)
	}

	return w
}
