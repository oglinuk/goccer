package main

import (
	"log"
	"sync"

	"./crawlers"
	"./writers"
)

// Job to be crawled
type Job struct {
	crawler string
	writer  string
	path    string
}

// InitProducer queues all of the paths to the worker pool
func InitProducer(workers int, c, w, d string, paths, filters []string) {
	jobs := make(chan Job)

	wg := &sync.WaitGroup{}
	for i := 0; i <= workers; i++ {
		go consume(jobs, wg, filters, d)
	}

	for i, path := range paths {
		wg.Add(1)
		go func(i int, c, w, path string) {
			log.Printf("Crawling[%d]: %s", i, path)
			jobs <- Job{crawler: c, writer: w, path: path}
		}(i, c, w, path)
	}
	wg.Wait()
	close(jobs)
}

// consume all the queued jobs
func consume(jobs <-chan Job, wg *sync.WaitGroup, filters []string, dir string) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			c := crawlers.CreateCrawler(job.crawler, job.path)
			w := writers.CreateWriter(job.writer, dir, filters)
			if c != nil && w != nil {
				w.Write(c.Crawl())
			}
			wg.Done()
		}
	}
}
