package main

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	timeComplexity time.Time
)

func init() {
	timeComplexity = time.Now()

	logFile, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("main.go::init::os.OpenFile::ERROR: %s", err.Error())
	}

	mw := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(mw)

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
