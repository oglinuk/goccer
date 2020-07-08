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

// InitProducer queues all paths to the worker pool
func InitProducer(workers int, c, w string, paths, filters []string) {
	jobs := make(chan Job)

	wg := &sync.WaitGroup{}
	for i := 0; i <= workers; i++ {
		go consume(jobs, wg, filters)
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
func consume(jobs <-chan Job, wg *sync.WaitGroup, filters []string) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}

			pw := writers.CreateWriter(job.writer, "data/parsed", filters)
			if pw != nil {
				pw.Write([]string{job.path})
			}

			c := crawlers.CreateCrawler(job.crawler, job.path)
			rw := writers.CreateWriter(job.writer, "data/raw", filters)
			if c != nil && rw != nil {
				collection := c.Crawl()
				rw.Write(collection)
				log.Printf("Collected %d paths from %s ...", len(collection), job.path)
			}
			wg.Done()
		}
	}
}
