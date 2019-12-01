package main

import (
	"archive/zip"
	"bufio"
	"fmt"
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

	za, err := zipArchive()
	if err != nil {
		return 0, err
	}
	defer za.Close()

	aw, err := za.Create(fmt.Sprintf("%s.txt", archiveName))
	if err != nil {
		return 0, err
	}

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
		Filters: []string{
			"facebook",
			"instagram",
			"youtube",
			"google",
			"amazon",
			"microsoft",
			"azure",
		},
		Seeds: uncrawled,
	})

	if err != nil {
		return 0, err
	}

	err = os.RemoveAll(archiveName)
	if err != nil {
		return 0, err
	}

	return len(uncrawled), nil
}

// zipArchive creates/returns a zip writer
func zipArchive() (*zip.Writer, error) {
	af, err := os.Create(fmt.Sprintf("%s.zip", archiveName))
	if err != nil {
		return nil, err
	}

	return zip.NewWriter(af), nil
}
