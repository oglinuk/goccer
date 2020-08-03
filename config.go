package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

const (
	configName string = "cfg.json"
)

// Default configurations
var (
	defaultHTTPDisk = &Config{
		MaxWorkers: runtime.GOMAXPROCS(0),
		Crawler:    "http",
		Writer:     "disk",
		Filters: []string{
			"facebook",
			"instagram",
			"google",
			"youtube",
			"amazon",
			"microsoft",
			"apple",
		},
		Paths: []string{
			"https://en.wikipedia.org/wiki/Chaos_Theory",
			"https://en.wikipedia.org/wiki/Machine_Learning",
		},
	}

	defaultFsDisk = &Config{
		MaxWorkers: runtime.GOMAXPROCS(0),
		Crawler:    "fs",
		Writer:     "disk",
		Filters: []string{
			".cache",
			".config",
			".Trash",
		},
		Paths: []string{
			"/home",
		},
	}
)

// Config file
type Config struct {
	MaxWorkers int
	Crawler    string
	Writer     string
	Filters    []string
	Paths      []string
}

// SaveConfig file
func SaveConfig(cf *Config) error {
	f, err := os.Create(configName)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t")
	encoder.Encode(cf)

	return nil
}

// LoadConfig file
func LoadConfig() (Config, error) {
	var cf Config
	f, err := os.Open(configName)
	if err != nil {
		var cfg *Config

		switch *config {
		case "httpdisk":
			cfg = defaultHTTPDisk
		case "fsdisk":
			cfg = defaultFsDisk
		default:
			cfg = defaultHTTPDisk
		}

		SaveConfig(cfg)
		return cf, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cf)

	return cf, err
}

// Aggregate uncrawled
func Aggregate() {
	log.Printf("Starting aggregation ...")

	var paths []string

	err := filepath.Walk("data/raw", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			f, err := os.Open(path)
			if err != nil {
				log.Printf("config.go::Aggregate::os.Open(%s)::ERROR: %s", path, err.Error())
			}
			defer f.Close()

			bs := bufio.NewScanner(f)
			for bs.Scan() {
				paths = append(paths, bs.Text())
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Failed to aggregate ...")
	}

	// TODO: Change below to read current cfg file, since *config wont be set
	// Im too lazy to do right now ...
	var cfg *Config

	switch *config {
	case "httpdisk":
		cfg = defaultHTTPDisk
	case "fsdisk":
		cfg = defaultFsDisk
	default:
		cfg = defaultFsDisk
	}

	cfg.Paths = paths

	SaveConfig(cfg)

	os.RemoveAll("data/raw")
}
