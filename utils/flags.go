package utils

import (
	"flag"
	"log"
	"runtime"
)

var (
	SpecifiedSeed = flag.String("seed", "", "Specific seed to start from")
)

func ParseFlags() {
	flag.Parse()

	if *SpecifiedSeed != "" {
		err := SaveConfig(&Config{
			MaxWorkers: runtime.GOMAXPROCS(0),
			Seeds:      []string{*SpecifiedSeed},
		})

		if err != nil {
			log.Printf("SaveConfig err: %v", err)
		}
	}
}
