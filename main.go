package main

import (
	"log"

	"github.com/OGLinuk/goccer/utils"
)

func main() {
	utils.ParseFlags()

	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Printf("LoadConfig err: %v", err)
	}

	utils.InitProducer(cfg)
	utils.Archive()
}
