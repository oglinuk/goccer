# (Go) (C)on(c)urrent Crawl(er)

## How to use

### With config.json

```JSON
// Default config.json
{
	"MaxWorkers": 4,
	"Crawler": "http",
	"Filters": [
		"facebook",
		"instagram",
		"google",
		"youtube",
		"amazon",
		"microsoft",
		"apple",
	],
	"Paths": [
		"https://en.wikipedia.org/wiki/Chaos_Theory",
		"https://en.wikipedia.org/wiki/Machine_Learning"
	]
}
```

### With Seed Flag

```./main -ct http -p https://en.wikipedia.org/wiki/Deep_Learning```

## Todo
* [X] Abstract crawler to allow for different types of crawlers 
	* [X] Implement HTTP crawler
	* [ ] Implement Filesystem crawler
* [ ] Abstract writer to allow for different store options
	* [X] Implement writing to disk
	* [ ] Implement writing to database
* [ ] Implement compression?
* [X] Implement filters
* [ ] Store crawl errors
* [ ] Dockerize crawlers