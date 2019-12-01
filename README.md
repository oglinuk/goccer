# (Go) (C)on(c)urrent Crawl(er)

## How to use

### With seeds.json

```JSON
// Default seeds.json
{
	"MaxWorkers": 4,
	"Filter": [
		"facebook",
		"instagram",
		"youtube",
		"google",
		"amazon",
		"microsoft",
		"azure"
	],
	"Seeds": [
		"https://en.wikipedia.org/wiki/Chaos_Theory",
		"https://en.wikipedia.org/wiki/Machine_Learning"
	]
}
```

### With Specific Seed Flag

```./goccer -s https://en.wikipedia.org/wiki/Deep_Learning```

## Todo
* [ ] Crawl other targets
	* [ ] Device filesystems
* [ ] Abstract archival to allow for different datastores
* [ ] gokv backend for archival
* [ ] Configuration for the compression of the archive
* [*] Filter/Blacklist option when crawling
* [ ] Archive crawl errors
* [ ] Replace ```seeds.json``` with a queue system
* [ ] Dockerize
