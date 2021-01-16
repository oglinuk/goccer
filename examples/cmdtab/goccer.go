package main

import (
	"gitlab.com/rwxrob/cmdtab"
)

func init() {
	x := cmdtab.New("goccer", "crawl")
//	x.Default = ""
	x.Summary = `Crawl and collect links`
	x.Version = "1.0.0"
	x.Author = "oglinuk"
	x.Git = "gitlab.com/oglinuk/goccer"
//	x.Copyright = ""
	x.License = "Apache 2.0"
	x.Description = `TODO ...`
}
