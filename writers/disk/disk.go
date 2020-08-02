package disk

import (
	"log"
)

var (
	baseDiskDirName = "data"
)

// Disk writer interface
type Disk interface {
	Write([]string) error
}

// CreateDiskWriter of ctype
func CreateDiskWriter(ctype, path string) Disk {
	var w Disk

	switch ctype {
	case "http":
		w = NewHTTPDiskWriter(path)
	case "fs":
		w = NewFsDiskWriter(path)
	default:
		log.Fatalf("writers::disk::disk.go::CreateDiskWriter(%s, ...)::ERROR: Invalid crawler type", ctype)
	}

	return w
}
