package main

import (
	"log"

	"./hive"
	"./utils"
)

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}

	for i, seed := range cfg.Seeds {
		log.Printf("[%d]Fetching: %s", i, seed)
		q := hive.NewQueen(seed)
		q.SpawnDrone()
	}

	utils.AggregateConfig()
}
