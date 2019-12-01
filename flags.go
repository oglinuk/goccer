package main

import (
	"flag"
	"log"
	"runtime"
)

var (
	specificSeed = flag.String("s", "", "Specific seed to start from")
)

// ParseFlags if any
func ParseFlags() {
	flag.Parse()

	if *specificSeed != "" {
		err := SaveConfig(&Config{
			MaxWorkers: runtime.GOMAXPROCS(0),
			Seeds:      []string{*specificSeed},
		})

		if err != nil {
			log.Printf("SaveConfig err: %v", err)
		}
	}
}
