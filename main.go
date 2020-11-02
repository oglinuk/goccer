package main

import (
	"fmt"
	"log"
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
	c := goccer.NewHTTPCrawler([]string{
		"https://en.wikipedia.org/wiki/Deep_Learning",
		"https://en.wikipedia.org/wiki/Web_search_engine",
		"https://en.wikipedia.org/wiki/Chaos_Theory",
	})

	collected, err := c.Crawl()
	if err != nil {
		log.Fatalf("Failed to c.Crawl: %s", err.Error())
	}

	for _, link := range collected {
		fmt.Printf("%s\n", link)
	}
	fmt.Printf("Collected %d links in %s ...\n", len(collected), time.Since(timeComplexity))
}
