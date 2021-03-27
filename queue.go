package goccer

import (
	"log"
	"runtime"
	"sync"
)

// Job to be crawled
type Job struct {
	paths []string
}

// WorkerPool that consumes jobs
type WorkerPool struct {
	jobs    chan Job
	wg      *sync.WaitGroup
	filters map[string]struct{}
	w       *memoryPool
}

// check if path contains any of the wp.filters
func (wp *WorkerPool) check(path string) bool {
	if _, exists := wp.filters[path]; exists {
		return true
	}
	return false
}

// consume all queued jobs in the WorkerPool
func (wp *WorkerPool) consume() {
	for {
		select {
		case job, ok := <-wp.jobs:
			if !ok {
				return
			}

			for _, path := range job.paths {
				if wp.check(path) {
					log.Printf("Filtered: %s", path)
					wp.wg.Done()
					return
				}

				if _, exists := wp.w.roots[path]; !exists {
					wp.w.write([]string{path})
				}

				c := newHTTPCrawler(path)
				collection, err := c.crawl()
				if err != nil {
					log.Printf("consume::c.crawl: %s", err.Error())
					return // TODO: Need to do better ...
				}

				if collection != nil {
					wp.w.write(collection)
				}

				wp.wg.Done()
			}
		}
	}
}

// Queue paths
func (wp *WorkerPool) Queue(paths []string) []string {
	wp.wg.Add(len(paths))
	wp.jobs <- Job{paths: paths}
	wp.wg.Wait()

	roots := wp.w.getRoots()

	return roots
}

// InitProducer starts GOMAXPROCS number of consumers
func (wp *WorkerPool) InitProducer() {
	for i := 0; i <= runtime.GOMAXPROCS(0); i++ {
		go wp.consume()
	}
}

// NewWorkerPool constructor
func NewWorkerPool(filters map[string]struct{}) *WorkerPool {
	return &WorkerPool{
		jobs:    make(chan Job),
		wg:      &sync.WaitGroup{},
		filters: filters,
		w:       newMemoryPool(),
	}
}
