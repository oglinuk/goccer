package main

import (
	"log"
	"sync"

	"./crawlers"
	"./writers"
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

// consume all queued jobs in the WorkerPool
func consume(wp WorkerPool) {
	pw := writers.CreateWriter(wp.writer, "data/parsed", wp.filters)
	rw := writers.CreateWriter(wp.writer, "data/raw", wp.filters)

	for {
		select {
		case job, ok := <-wp.jobs:
			if !ok {
				return
			}

			if pw != nil {
				pw.Write([]string{job.path})
			}

			c := crawlers.CreateCrawler(wp.crawler, job.path)
			if c != nil && rw != nil {
				collection := c.Crawl()
				if collection != nil {
					rw.Write(collection)
					log.Printf("Collected %d paths from %s ...", len(collection), job.path)
				}
			}
			wp.wg.Done()
		}
	}
}
