package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
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

func AggregateConfig() error {
	check := make(map[string]struct{})
	var uncrawled []string

	f, err := os.Open("to_crawl.txt")
	if err != nil {
		return err
	}

	fname := fmt.Sprintf("%d.txt", time.Now().Unix())
	ucf, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}

	log.Println("Processing aggregation ...")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		scanned := scanner.Text()

		ucf.WriteString(fmt.Sprintf("%s\n", scanned))
		if _, exists := check[scanned]; !exists {
			check[scanned] = struct{}{}
			uncrawled = append(uncrawled, scanned)
		}
	}

	err = SaveConfig(&Config{
		MaxWorkers: 4,
		Seeds:      uncrawled,
	})

	if err != nil {
		return err
	}

	err = os.RemoveAll("to_crawl.txt")
	if err != nil {
		return err
	}

	f.Close()
	ucf.Close()

	log.Println("Finished aggregation ...")

	return nil
}
