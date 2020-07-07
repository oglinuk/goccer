package main

import (
	"flag"
	"log"
	"runtime"
)

var (
	crawlerType = flag.String("ct", "", "Crawler type")
	path        = flag.String("p", "", "Specific path to start from")
)

// ParseFlags if any
func ParseFlags() {
	flag.Parse()

	if *path != "" && *crawlerType != "" {
		err := SaveConfig(&Config{
			MaxWorkers: runtime.GOMAXPROCS(0),
			Crawler:    *crawlerType,
			Paths:      []string{*path},
		})

		if err != nil {
			log.Printf("SaveConfig err: %v", err)
		}
	}
}
