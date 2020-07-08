package writers

import (
	"log"
)

// Writer is a thing that writes the given path to somewhere
type Writer interface {
	Write(path []string) error
}

// CreateWriter of wtype
func CreateWriter(wtype, path string, filters []string) Writer {
	switch wtype {
	case "disk":
		return NewDiskWriter(path, filters)
	case "memory":
		return NewMemoryWriter(path, filters)
	default:
		log.Printf("writer.go::CreateWriter(%s, ...)::ERROR: Invalid crawler type", wtype)
		return nil
	}
}
