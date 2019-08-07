package hive

import (
	"sync"
)

type Job struct {
	URL string
}

func Worker(jobs <-chan Job, wg *sync.WaitGroup) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			q := NewQueen(job.URL)
			q.SpawnDrone()
			q.Aggregate()
			wg.Done()
		}
	}
}
