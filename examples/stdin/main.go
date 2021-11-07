package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/oglinuk/goccer"
)

func main() {
	wp := goccer.NewWorkerpool()

	bs := bufio.NewScanner(os.Stdin)

	var seeds []string

	for bs.Scan() {
		if bs.Text() != "start" {
			seeds = append(seeds, bs.Text())
		}
	}

	collected := wp.Queue(seeds)

	for _, c := range collected {
		fmt.Println(c)
	}
}
