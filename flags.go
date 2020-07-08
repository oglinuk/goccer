package main

import (
	"flag"
	"log"
	"runtime"
)

var (
	crawlerType = flag.String("ct", "", "Crawler type")
	writerType  = flag.String("wt", "", "Writer type")
	path        = flag.String("p", "", "Specific path to start from")
)

// ParseFlags if any
func ParseFlags() {
	flag.Parse()

	if *path != "" && *crawlerType != "" && *writerType != "" {
		err := SaveConfig(&Config{
			MaxWorkers: runtime.GOMAXPROCS(0),
			Crawler:    *crawlerType,
			Writer:     *writerType,
			// TODO: Im lazy, but need to add flag for filters eventually ...
			Filters: []string{
				"facebook",
				"instagram",
				"google",
				"youtube",
				"amazon",
				"microsoft",
				"apple",
			},
			Paths: []string{*path},
		})

		if err != nil {
			log.Printf("SaveConfig err: %v", err)
		}
	}
}
