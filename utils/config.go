package utils

import (
	"encoding/json"
	"os"
)

const (
	configName string = "config.json"
)

type Config struct {
	MaxWorkers int
	Seeds      []string
}

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

func LoadConfig() (Config, error) {
	var cf Config
	f, err := os.Open(configName)
	if err != nil {
		SaveConfig(&Config{
			MaxWorkers: 4,
			Seeds: []string{
				"https://en.wikipedia.org/wiki/Chaos_Theory",
			},
		})
		return cf, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cf)

	return cf, err
}
