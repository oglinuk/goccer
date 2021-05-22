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
	wg   *sync.WaitGroup
	w    *memoryPool
}

// NewWorkerpool constructor
func NewWorkerpool() *workerpool {
	wp := &workerpool{
		jobs: make(chan job),
		wg:   &sync.WaitGroup{},
		w:    NewMemorypool(),
	}

	for i := 0; i <= runtime.GOMAXPROCS(0); i++ {
		go wp.consume()
	}

	return wp
}

func (wp *workerpool) consume() {
	for {
		select {
		case job, ok := <-wp.jobs:
			if !ok {
				return
			}

			for _, path := range job.paths {
				c := NewCrawler(path)
				collected := c.Crawl()
				if c.Err != nil {
					log.Printf("queue.go::consume::c.Crawl: %s", c.Err.Error())
				}
				wp.w.write(collected)
				wp.wg.Done()
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
