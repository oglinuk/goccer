package main

import (
	"log"

	"./utils"
)

func main() {
	cfg, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}

	log.Println(cfg)
}
