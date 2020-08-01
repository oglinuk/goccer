package main

import (
	"flag"
	"log"
	"os"
)

var (
	config = flag.String("c", "", "Configuration: httpdisk, fsdisk")
	path   = flag.String("p", "", "Specific path to start from")
)

// ParseFlags if any
func ParseFlags() {
	flag.Parse()

	if *config != "" {
		var cfg *Config
		switch *config {
		case "httpdisk":
			cfg = defaultHTTPDisk
		case "fsdisk":
			cfg = defaultFsDisk
		}

		if *path != "" {
			cfg.Paths = []string{*path}
		}

		err := SaveConfig(cfg)

		if err != nil {
			log.Printf("flags.go::ParseFlags::SaveConfig::ERROR: %s", err.Error())
		}

		log.Printf("Successfully generated (%s) configuration file ... Please re-run goccer.", *config)
		os.Exit(0)
	}
}
