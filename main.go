package main

import (
	"log"

	"./conducer"
	"./utils"
)

func main() {
	utils.ParseFlags()

	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Printf("LoadConfig err: %v", err)
		return
	}

	archiver := utils.NewArchiver()

	conducer.InitProducer(cfg.MaxWorkers, cfg.Seeds, archiver.ArchiveFile)

	archiver.Archive()
}
