package utils

import (
	"sync"

	"github.com/OGLinuk/goccer/hive"
)

type Job struct {
	URL string
}

func Consumer(jobs <-chan Job, wg *sync.WaitGroup) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			q := hive.NewQueen(job.URL)
			q.SpawnDrone()
			q.Aggregate()
			wg.Done()
		}
	}
}
