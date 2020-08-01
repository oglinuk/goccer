package writers

import (
	"log"
)

// Writer is a thing that writes the given path to somewhere
type Writer interface {
	Write(path []string) error
}

// CreateWriter of wtype
func CreateWriter(wtype, path string) Writer {
	var w Writer

	switch wtype {
	case "disk":
		w = NewDiskWriter(path)
	case "memory":
		w = NewMemoryWriter(path)
	default:
		log.Fatalf("writer.go::CreateWriter(%s, ...)::ERROR: Invalid crawler type", wtype)
	}

	return w
}
