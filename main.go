package main

import (
	"fmt"
	"log"

	goccer "./include"
)

func main() {
	c := goccer.NewHTTPCrawler("https://en.wikipedia.org/wiki/Deep_Learning")

	collected, err := c.Crawl()
	if err != nil {
		log.Fatalf("Failed to c.Crawl: %s", err.Error())
	}

	for _, link := range collected {
		fmt.Printf("%s\n", link)
	}
	fmt.Printf("Collected %d links ...\n", len(collected))
}
