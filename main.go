package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/OGLinuk/goccer/utils"
)

func main() {
	specifiedSeed := flag.String("s", "", "Specific seed to start from")
	flag.Parse()

	if *specifiedSeed != "" {
		err := utils.SaveConfig(&utils.Config{
			MaxWorkers: runtime.GOMAXPROCS(0),
			Seeds:      []string{*specifiedSeed},
		})

		if err != nil {
			log.Printf("SaveConfig err: %v", err)
		}
	}

	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Printf("LoadConfig err: %v", err)
	}

	utils.InitProducer(cfg)
	utils.Archive()
}
