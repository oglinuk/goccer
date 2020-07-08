# (Go) (C)on(c)urrent Crawl(er)

## How to use

### With config.json

```JSON
// Default config.json
{
	"MaxWorkers": 4,
	"Crawler": "http",
	"Writer": "disk",
	"Filters": [
		"facebook",
		"instagram",
		"google",
		"youtube",
		"amazon",
		"microsoft",
		"apple"
	],
	"Paths": [
		"https://en.wikipedia.org/wiki/Chaos_Theory",
		"https://en.wikipedia.org/wiki/Machine_Learning"
	]
}
```

### With Flags

```./goccer -ct http -wt disk -p https://en.wikipedia.org/wiki/Deep_Learning```

## Todo
* [X] Abstract crawler to allow for different types of crawlers 
	* [X] Implement HTTP crawler
	* [ ] Implement Filesystem crawler
* [X] Abstract writer to allow for different types of writers
	* [X] Implement disk writer
	* [ ] Implement memory writer
	* [ ] Implement database writer
* [ ] Implement compression?
* [X] Implement filters
* [ ] Store crawl errors
* [ ] Dockerize
* [ ] Create examples?