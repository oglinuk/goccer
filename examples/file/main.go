package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/oglinuk/goccer"
)

var (
	// Tested using a laptop with an i5-5200U CPU @ 2.20GHz
	// 100.txt // Collected 32610 links in 1.268676131s ...
	// 300.txt // Collected 79303 links in 1.771245201s ...
	// 500.txt // Collected 126395 links in 2.658882296s ...
	// 1000.txt // Collected 233647 links in 5.241596656s ...
	// 5000.txt // Collected 622264 links in 13.58724316s ...
	// 15000.txt // Collected 1746933 links in 32.152465588s ...

	sf = flag.String("sf", "1000.txt", "Seed file to use")
	timeComplexity time.Time
)

func init() {
	timeComplexity = time.Now()
}

func main() {
	flag.Parse()

	wp := goccer.NewWorkerpool()

	var seeds []string

	f, _ := os.Open(fmt.Sprintf("./seeds/%s", *sf))
	defer f.Close()

	bs := bufio.NewScanner(f)

	for bs.Scan() {
		seeds = append(seeds, bs.Text())
	}

	collected := wp.Queue(seeds)
	fmt.Printf("Crawled %d && collected %d links in %s ...\n", len(seeds), len(collected), time.Since(timeComplexity))

	outf, _ := os.Create("collected.txt")
	defer outf.Close()

	for _, c := range collected {
		outf.WriteString(fmt.Sprintf("%s\n", c))
	}
}
