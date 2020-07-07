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

	ParseFlags()

	if _, err := os.Stat("data"); err != nil {
		if err = os.MkdirAll("data", 0777); err != nil {
			log.Fatalf("MkdirAll err: %v", err)
		}
	}
}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("LoadConfig err: %v", err)
	}

	InitProducer(cfg.MaxWorkers, cfg.Crawler, cfg.Paths, cfg.Filters)

	log.Printf("Crawled [%d] in %s ...",
		len(cfg.Paths), time.Since(timeComplexity))
}
