package utils

import (
	"sync"

	"../hive"
)

type Job struct {
	URL string
}

func consume(jobs <-chan Job, wg *sync.WaitGroup) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			q := hive.NewQueen(job.URL, ArchiveFile)
			q.SpawnDrone()
			wg.Done()
		}
	}
}
