package utils

import (
	"log"
	"sync"
)

func InitProducer(cfg Config, store string) {
	jobs := make(chan Job)

	wg := &sync.WaitGroup{}
	for i := 0; i <= cfg.MaxWorkers; i++ {
		go consume(jobs, wg, store)
	}

	for i, seed := range cfg.Seeds {
		wg.Add(1)
		go func(i int, seed string) {
			log.Printf("Fetching[%d]: %s", i, seed)
			jobs <- Job{URL: seed}
		}(i, seed)
	}
	wg.Wait()
	close(jobs)
}
