# (Go) (C)on(c)urrent Crawl(er)

## How to use

### With seeds.json

```JSON
// Default seeds.json
{
	"MaxWorkers": 4,
	"Filters": [
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
* [X] Filter option when crawling
* [ ] Change from net/http to [fasthttp](https://github.com/valyala/fasthttp)
* [ ] Change from JSON configuration to YAML
* [ ] Replace seeds in configuration file to a queue system
* [ ] Crawl other targets
	* [ ] Linux filesystems
	* [ ] Windows filesystems
* [ ] Abstract archival to allow for different stores
* [ ] Dockerize
