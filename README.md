# (Go) (C)on(c)urrent Web Crawl(er)

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

```./main -s "https://en.wikipedia.org/wiki/Deep_Learning"```

## Todo
- Abstract archival to allow for different datastores
- Add ability to filter when crawling