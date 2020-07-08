package main

import (
	"fmt"

	"github.com/oglinuk/goccer/writers"
)

// Writer is a thing that writes the given path to somewhere
type Writer interface {
	write(path string) error
}

// CreateWriter of wtype
func CreateWriter(wtype string, filters []string) (Writer, error) {
	switch wtype {
	case "disk":
		return writers.NewDiskWriter(path, filters), nil
	case "memory":
		return writers.NewMemoryWriter(path, filters), nil
	default:
		return nil, fmt.Errorf("writer.go::CreateWriter(%s, ...)::ERROR: Invalid crawler type", ctype)
	}
}
