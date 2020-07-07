package main

import (
	"log"
	"sync"
)

// Job to be crawled
type Job struct {
	Crawler string
	Path    string
}

// InitProducer queues all of the paths to the worker pool
func InitProducer(workers int, crawler string, paths, filters []string) {
	jobs := make(chan Job)

	wg := &sync.WaitGroup{}
	for i := 0; i <= workers; i++ {
		go consume(jobs, wg)
	}

	for i, path := range paths {
		wg.Add(1)
		go func(i int, path string) {
			log.Printf("Crawling[%d]: %s", i, path)
			jobs <- Job{Crawler: crawler, Path: path}
		}(i, path)
	}
	wg.Wait()
	close(jobs)
}

// consume all the queued jobs
func consume(jobs <-chan Job, wg *sync.WaitGroup) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			log.Println(job)
			c, err := CreateCrawler(job.Crawler, job.Path)
			if err != nil {
				log.Printf("queue.go::consume::CreateCrawler::ERROR: %s", err.Error())
			}
			c.Crawl()
			wg.Done()
		}
	}
}
