package conducer

import (
	"log"
	"sync"
)

func InitProducer(workers int, seeds []string, af string) {
	jobs := make(chan Job)

	wg := &sync.WaitGroup{}
	for i := 0; i <= workers; i++ {
		go consume(jobs, wg, af)
	}

	for i, seed := range seeds {
		wg.Add(1)
		go func(i int, seed string) {
			log.Printf("Fetching[%d]: %s", i, seed)
			jobs <- Job{URL: seed}
		}(i, seed)
	}
	wg.Wait()
	close(jobs)
}
