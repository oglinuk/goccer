package goccer

import (
	"log"
	"runtime"
	"sync"
	"syscall"
)

// job is an array of strings to crawl
type job struct {
	paths []string
}

// Workerpool is a wrapper for the jobs chan to send seeds, the wg
// sync.WaitGroup to wait for goroutines, the w(riter) which is an
// in-memory implementation (only at the moment, see TODO above crawler),
// the (c)rawler, and a mu(tex) to prevent concurrent map writes
type Workerpool struct {
	jobs chan job
	wg *sync.WaitGroup
	w *memoryPool
	c *crawler
	mu *sync.Mutex
}

// init to set the ulimit to max
// TODO: Find out if this is the correct place to do it
func init() {
	var rLimit syscall.Rlimit
	syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	rLimit.Cur = rLimit.Max
	syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}

// NewWorkerpool constructor
func NewWorkerpool() *Workerpool {
	wp := &Workerpool{
		jobs: make(chan job, 100),
		wg: &sync.WaitGroup{},
		w: newMemorypool(),
		c: newCrawler(),
		mu: &sync.Mutex{},
	}

	for i := 0; i <= runtime.GOMAXPROCS(0); i++ {
		go wp.consume()
	}

	return wp
}

// consume crawls any job.paths sent over the jobs chan, then writes the
// collected to whatever wp.w is (in-memory in this case)
func (wp *Workerpool) consume() {
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
func (wp *Workerpool) Queue(ps []string) []string {
	wp.wg.Add(len(ps))
	wp.jobs <- job{paths: ps}
	wp.wg.Wait()

	return wp.w.GetPaths()
}
