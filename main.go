package main

import (
	"flag"
	"log"
	"runtime"
	"sync"

	"./hive"
	"./utils"
)

func main() {
	specifiedSeed := flag.String("s", "", "Specific seed to start from")
	flag.Parse()

	if *specifiedSeed != "" {
		err := utils.SaveConfig(&utils.Config{
			MaxWorkers: runtime.GOMAXPROCS(0),
			Seeds:      []string{*specifiedSeed},
		})

		if err != nil {
			panic(err)
		}
	}

	cfg, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}

	jobs := make(chan hive.Job)

	wg := &sync.WaitGroup{}
	for i := 0; i <= cfg.MaxWorkers; i++ {
		go hive.Worker(jobs, wg)
	}

	for i, seed := range cfg.Seeds {
		wg.Add(1)
		go func(i int, seed string) {
			log.Printf("Fetching[%d]: %s", i, seed)
			jobs <- hive.Job{URL: seed}
		}(i, seed)
	}

	wg.Wait()
	close(jobs)
	utils.AggregateConfig()
}
