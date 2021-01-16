package main

import (
	"fmt"

	"github.com/OGLinuk/goccer"
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("crawl")
	x.Usage = `<url>`
	x.Summary = `Crawl <url>`
	x.Description = `TODO ...`
	x.Method = func(args []string) error {
		if len(args) != 1 {
			return x.UsageError()
		}

		filters := map[string]struct{}{
			"google": {},
			"youtube": {},
			"facebook": {},
			"instagram": {},
			"amazon": {},
			"microsoft": {},
			"apple": {},
		}

		wp := goccer.NewWorkerPool(filters)
		wp.InitProducer()

		collected := wp.Queue(args)

		for _, link := range collected {
			fmt.Printf("%s\n", link)
		}
		return nil
	}
}
