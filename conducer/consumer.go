package conducer

import (
	"sync"

	"../hive"
)

type Job struct {
	URL string
}

func consume(jobs <-chan Job, wg *sync.WaitGroup, af string) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			q := hive.NewQueen(job.URL, af)
			q.SpawnDrone()
			wg.Done()
		}
	}
}
