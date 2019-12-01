package main

import (
	"log"
	"sync"
)

// Job to be crawled
type Job struct {
	URL string
}

// InitProducer queues all of the seeds to the worker pool
func InitProducer(workers int, seeds, filters []string) {
	jobs := make(chan Job)

	wg := &sync.WaitGroup{}
	for i := 0; i <= workers; i++ {
		go consume(jobs, wg, filters)
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

// consume all the queued jobs
func consume(jobs <-chan Job, wg *sync.WaitGroup, filters []string) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			q := NewQueen(job.URL, filters)
			q.SpawnDrone()
			wg.Done()
		}
	}
}
