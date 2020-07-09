package main

import (
	"log"
	"time"
)

var (
	timeComplexity time.Time
)

func init() {
	timeComplexity = time.Now()

	ParseFlags()
}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("main.go::main::LoadConfig::ERROR: %s", err.Error())
	}

	InitProducer(cfg)

	log.Printf("Crawled [%d] in %s ...", len(cfg.Paths), time.Since(timeComplexity))
}
