package main

import (
	"log"
	"strings"
	"sync"

	"github.com/oglinuk/goccer/crawlers"
	"github.com/oglinuk/goccer/writers"
)

// Job to be crawled
type Job struct {
	path string
}

// WorkerPool that consumes jobs
type WorkerPool struct {
	jobs    chan Job
	wg      *sync.WaitGroup
	crawler string
	writer  string
	filters []string
}

// InitProducer queues all paths from config.json to the WorkerPool
func InitProducer(cfg Config) {
	wp := WorkerPool{
		jobs:    make(chan Job),
		wg:      &sync.WaitGroup{},
		crawler: cfg.Crawler,
		writer:  cfg.Writer,
		filters: cfg.Filters,
	}

	for i := 0; i <= cfg.MaxWorkers; i++ {
		go consume(wp)
	}

	for i, path := range cfg.Paths {
		wp.wg.Add(1)
		go func(i int, p string) {
			log.Printf("Crawling[%d]: %s", i, p)
			wp.jobs <- Job{path: p}
		}(i, path)
	}
	wp.wg.Wait()
	close(wp.jobs)
}

// check if path contains any of the wp.filters
// TODO: Refactor below to do string.Contains concurrently
func (wp WorkerPool) check(path string) bool {
	for _, filter := range wp.filters {
		if strings.Contains(path, filter) {
			return true
		}
	}

	return false
}

// consume all queued jobs in the WorkerPool
func consume(wp WorkerPool) {
	pw := writers.CreateWriter(wp.writer, "data/parsed")
	rw := writers.CreateWriter(wp.writer, "data/raw")

	for {
		select {
		case job, ok := <-wp.jobs:
			if !ok {
				return
			}

			for _, filter := range wp.filters {
				if strings.Contains(job.path, filter) {
					log.Printf("Filtered: %s", job.path)
					wp.wg.Done()
					return
				}
			}

			pw.Write([]string{job.path})

			c := crawlers.CreateCrawler(wp.crawler, job.path)
			collection := c.Crawl()
			if collection != nil {
				rw.Write(collection)
				log.Printf("Collected %d paths from %s ...", len(collection), job.path)
			}

			wp.wg.Done()
		}
	}
}
