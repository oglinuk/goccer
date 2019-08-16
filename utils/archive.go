package utils

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func Archive() error {
	var uncrawled []string

	file, err := os.Open("to_crawl.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	fname := fmt.Sprintf("%d", time.Now().Unix())
	af, err := createArchive(fname)
	if err != nil {
		return err
	}
	defer af.Close()

	aw, err := af.Create(fmt.Sprintf("%s.txt", fname))
	if err != nil {
		return err
	}

	log.Println("Processing archival ...")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanned := scanner.Text()

		aw.Write([]byte(fmt.Sprintf("%s\n", scanned)))
		uncrawled = append(uncrawled, scanned)
	}

	err = SaveConfig(&Config{
		MaxWorkers: runtime.GOMAXPROCS(0),
		Seeds:      uncrawled,
	})

	if err != nil {
		return err
	}

	err = os.RemoveAll("to_crawl.txt")
	if err != nil {
		return err
	}

	log.Println("Finished archival ...")

	return nil
}

func createArchive(archiveName string) (*zip.Writer, error) {
	af, err := os.Create(fmt.Sprintf("%s.zip", archiveName))
	if err != nil {
		return nil, err
	}

	return zip.NewWriter(af), nil
}
