package main

import (
	"./hive"
	"./utils"
)

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}

	for _, seed := range cfg.Seeds {
		q := hive.NewQueen(seed)
		q.SpawnDrone()
	}
}
