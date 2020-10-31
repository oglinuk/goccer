package goccer

import (
	"log"
	"runtime"
	"strings"
	"sync"
)

// Job to be crawled
type Job struct {
	path string
}

// WorkerPool that consumes jobs
type WorkerPool struct {
	jobs    chan Job
	wg      *sync.WaitGroup
	filters []string
	w       *MemoryPool
}

// check if path contains any of the wp.filters
// TODO: Refactor below to do string.Contains concurrently
func (wp *WorkerPool) check(path string) bool {
	for _, filter := range wp.filters {
		if strings.Contains(path, filter) {
			return true
		}
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

			if wp.check(job.path) {
				log.Printf("Filtered: %s", job.path)
				wp.wg.Done()
				return
			}

			if _, exists := wp.w.roots[job.path]; !exists {
				wp.w.Write([]string{job.path})
			}

			c := NewHTTPCrawler(job.path)
			collection, err := c.Crawl()
			if err != nil {
				return // TODO: Need to do better ...
			}
			if collection != nil {
				wp.w.Write(collection)
			}

			wp.wg.Done()
		}
	}
}

// Queue paths
func (wp *WorkerPool) Queue(paths []string) []string {
	for i, path := range paths {
		wp.wg.Add(1)
		go func(i int, p string) {
			log.Printf("Crawling[%d]: %s", i, p)
			wp.jobs <- Job{path: p}
		}(i, path)
	}
	wp.wg.Wait()
	//close(wp.jobs)

	return wp.w.GetRoots()
}

// InitProducer starts GOMAXPROCS number of consumers
func (wp *WorkerPool) InitProducer() {
	for i := 0; i <= runtime.GOMAXPROCS(0); i++ {
		go wp.consume()
	}
}

// NewWorkerPool constructor
func NewWorkerPool(filters []string) *WorkerPool {
	return &WorkerPool{
		jobs:    make(chan Job),
		wg:      &sync.WaitGroup{},
		filters: filters,
		w:       NewMemoryPool(),
	}
}
