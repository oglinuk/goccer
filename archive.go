package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

// Archive URLs into a compressed zip file
func Archive() (int, error) {
	unique := make(map[string]struct{})
	var uncrawled []string

	file, err := os.Open(archiveName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	af, err := createArchive()
	if err != nil {
		return 0, err
	}
	defer af.Close()

	aw, err := af.Create(fmt.Sprintf("%s.txt", archiveName))
	if err != nil {
		return 0, err
	}

	log.Println("Processing archival ...")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanned := scanner.Text()
		if _, exists := unique[scanned]; !exists {
			aw.Write([]byte(fmt.Sprintf("%s\n", scanned)))
			uncrawled = append(uncrawled, scanned)
			unique[scanned] = struct{}{}
		}
	}

	err = SaveConfig(&Config{
		MaxWorkers: runtime.GOMAXPROCS(0),
		Seeds:      uncrawled,
	})

	if err != nil {
		return 0, err
	}

	err = os.RemoveAll(archiveName)
	if err != nil {
		return 0, err
	}

	log.Println("Finished archival ...")

	return len(uncrawled), nil
}

func createArchive() (*zip.Writer, error) {
	af, err := os.Create(fmt.Sprintf("%s.zip", archiveName))
	if err != nil {
		return nil, err
	}

	return zip.NewWriter(af), nil
}
