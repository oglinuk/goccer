package main

import (
	"github.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("goccer", "crawl")
	x.Summary = `Crawl and collect links`
	x.Version = "1.0.0"
	x.Author = "oglinuk"
	x.License = "Apache 2.0"
}
