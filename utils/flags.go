package utils

import (
	"flag"
	"log"
	"runtime"
)

var (
	SpecificSeed = flag.String("ss", "", "Specific seed to start from")
	Store        = flag.String("store", "", "Store to persist with")
)

func ParseFlags() {
	flag.Parse()

	if *SpecificSeed != "" {
		err := SaveConfig(&Config{
			MaxWorkers: runtime.GOMAXPROCS(0),
			Seeds:      []string{*SpecificSeed},
		})

		if err != nil {
			log.Printf("SaveConfig err: %v", err)
		}
	}
}
