package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	archiveName    = fmt.Sprintf("data/%d", time.Now().Unix())
	timeComplexity time.Time
)

func init() {
	timeComplexity = time.Now()
}

func main() {
	ParseFlags()

	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig err: %v", err)
	}

	if _, err := os.Stat("data"); err != nil {
		if err = os.MkdirAll("data", 0777); err != nil {
			log.Fatalf("MkdirAll err: %v", err)
		}
	}

	InitProducer(cfg.MaxWorkers, cfg.Seeds)

	collected, err := Archive()
	if err != nil {
		log.Fatalf("Failed to Archive: %s", err.Error())
	}

	log.Printf("Crawled [%d] and collected [%d] in %s ...",
		len(cfg.Seeds), collected, time.Since(timeComplexity))
}
