package main

import (
	"fmt"
	"log"
	"time"

	"github.com/OGLinuk/goccer/archive"
	"github.com/OGLinuk/goccer/utils"
)

func main() {
	utils.ParseFlags()

	cfg, err := utils.LoadConfig()
	if err != nil {
		log.Printf("LoadConfig err: %v", err)
	}

	if *utils.Store == "" {
		*utils.Store = fmt.Sprintf("%d", time.Now().Unix())
	}

	store := archive.NewLocalStore(*utils.Store)

	utils.InitProducer(cfg, store.Location)
	store.Archive()
}
