# Goccer

Go concurrent crawler(s) library

## Usage

```Go
package main

import (
	"log"

	"github.com/oglinuk/goccer"
)

func main() {
	wp := goccer.NewWorkerpool()

	collected := wp.Queue([]string{"https://fourohfournotfound.com"})

	for _, c := range collected {
		log.Println(c)
	}
}
```

## Examples

See [examples](examples) directory.
