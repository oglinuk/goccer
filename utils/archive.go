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

var (
	ArchiveFile = fmt.Sprintf("%d", time.Now().Unix())
)

func Archive() error {
	var uncrawled []string

	file, err := os.Open(ArchiveFile)
	if err != nil {
		return err
	}
	defer file.Close()

	af, err := createArchive(ArchiveFile)
	if err != nil {
		return err
	}
	defer af.Close()

	aw, err := af.Create(fmt.Sprintf("%s.txt", ArchiveFile))
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

	err = os.RemoveAll(ArchiveFile)
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
