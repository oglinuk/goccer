package main

import (
	"log"
	"os"

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

	if _, err := os.Stat("data"); err != nil {
		if err = os.MkdirAll("data", 0777); err != nil {
			log.Printf("MkdirAll err: %v", err)
		}
	}

	archiver := utils.NewArchiver()

	conducer.InitProducer(cfg.MaxWorkers, cfg.Seeds, archiver.ArchiveFile)

	archiver.Archive()
}
