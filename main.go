package main

import (
	"fmt"
	"time"

	goccer "./include"
)

var (
	timeComplexity time.Time
)

func init() {
	timeComplexity = time.Now()
}

func main() {
	filters := map[string]struct{}{
		"facebook":  {},
		"instagram": {},
		"google":    {},
		"youtube":   {},
		"amazon":    {},
		"microsoft": {},
		"apple":     {},
	}

	wp := goccer.NewWorkerPool(filters)
	wp.InitProducer()

	seeds := []string{
		"https://en.wikipedia.org/wiki/Deep_Learning",
		"https://en.wikipedia.org/wiki/Web_search_engine",
		"https://en.wikipedia.org/wiki/Chaos_Theory",
	}

	collected := wp.Queue(seeds)

	for _, link := range collected {
		fmt.Printf("%s\n", link)
	}
	fmt.Printf("Collected %d links in %s ...\n", len(collected), time.Since(timeComplexity))
}
