package main

import (
	"encoding/json"
	"os"
	"runtime"
)

const (
	configName string = "seeds.json"
)

// Config file
type Config struct {
	MaxWorkers int
	Filters    []string
	Seeds      []string
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
		SaveConfig(&Config{
			MaxWorkers: runtime.GOMAXPROCS(0),
			Filters: []string{
				"facebook",
				"instagram",
				"youtube",
				"google",
				"amazon",
				"microsoft",
				"azure",
			},
			Seeds: []string{
				"https://en.wikipedia.org/wiki/Chaos_Theory",
				"https://en.wikipedia.org/wiki/Machine_Learning",
			},
		})
		return cf, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cf)

	return cf, err
}
