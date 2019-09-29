# (Go) (C)on(c)urrent Crawl(er)

## How to use

### With Config.json

```JSON
// Default config.json
{
    "MaxWorkers": 4,
	"Seeds": [
		"https://en.wikipedia.org/wiki/Chaos_Theory",
		"https://en.wikipedia.org/wiki/Machine_Learning"
	]
}
```

### With Specific Seed Flag

```./main -s https://en.wikipedia.org/wiki/Deep_Learning```

## Todo
* [ ] Crawl a device directory
* [ ] Abstract archival to allow for different datastores
* [ ] xstore backend for archival
* [ ] Configuration for the compression of the archive
* [ ] Filter/Blacklist option when crawling
* [ ] Archive crawl errors
* [ ] Replace ```config.json``` with a queue system
* [ ] Dockerize

