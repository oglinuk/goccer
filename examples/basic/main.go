package main

import (
	"fmt"
	"time"

	"github.com/oglinuk/goccer"
)

var (
	timeComplexity time.Time
)

func init() {
	timeComplexity = time.Now()
}

func main() {
	wp := goccer.NewWorkerpool()

	seeds := []string{
		"https://en.wikipedia.org/wiki/Deep_Learning",
		"https://en.wikipedia.org/wiki/Web_search_engine",
		"https://en.wikipedia.org/wiki/Chaos_Theory",
	}

	collected := wp.Queue(seeds)
	t := time.Since(timeComplexity)

	for _, link := range collected {
		fmt.Printf("%s\n", link)
	}
	fmt.Printf("Crawled %d && collected %d links in %s ...\n", len(seeds), len(collected), t)
}
