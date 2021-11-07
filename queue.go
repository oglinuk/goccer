package goccer

import (
	"log"
	"runtime"
	"sync"
)

type job struct {
	paths []string
}

type workerpool struct {
	jobs chan job
	wg *sync.WaitGroup
	w *memoryPool
	c *crawler
	mu *sync.Mutex
}

// NewWorkerpool constructor
func NewWorkerpool() *workerpool {
	wp := &workerpool{
		jobs: make(chan job, 100),
		wg: &sync.WaitGroup{},
		w: NewMemorypool(),
		c: NewCrawler(),
		mu: &sync.Mutex{},
	}

	for i := 0; i <= runtime.GOMAXPROCS(0); i++ {
		go wp.consume()
	}

	return wp
}

func (wp *workerpool) consume() {
	for {
		select {
		case job := <-wp.jobs:
			for index, path := range job.paths {
				go func(i int, p string) {
					log.Printf("[%d]Crawling %s ...\n", i, p)
					collected, err := wp.c.Crawl(p)
					// TODO: If an error occurs we should parse the error. if the
					// error is a malformed URL, remove it. If it was a timeout or
					// something just went wrong, retry the request X(2?) times.
					if err != nil {
						log.Printf("queue.go::consume::c.Crawl: %s", err.Error())
					}
					wp.mu.Lock()
					err = wp.w.write(collected)
					if err != nil {
						log.Printf("queue.go::consume::wp.w.write: %s", err.Error())
					}
					wp.mu.Unlock()
					wp.wg.Done()
				}(index, path)
			}
		}
	}
}

// Queue (p)ath(s) to be crawled
func (wp *workerpool) Queue(ps []string) []string {
	wp.wg.Add(len(ps))
	wp.jobs <- job{paths: ps}
	wp.wg.Wait()

	return wp.w.GetPaths()
}
