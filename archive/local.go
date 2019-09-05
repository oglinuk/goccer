package archive

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/OGLinuk/goccer/utils"
)

var (
	_ ArchiveStore = (*LocalStore)(nil)
)

type LocalStore struct {
	Location string
}

func NewLocalStore(l string) *LocalStore {
	return &LocalStore{
		Location: l,
	}
}

func (ls *LocalStore) Archive() error {
	if err := localArchive(ls.location); err != nil {
		return err
	}

	return nil
}

func localArchive(archiveFile string) error {
	var uncrawled []string

	file, err := os.Open(archiveFile)
	if err != nil {
		return err
	}
	defer file.Close()

	af, err := createArchive(archiveFile)
	if err != nil {
		return err
	}
	defer af.Close()

	aw, err := af.Create(fmt.Sprintf("%s.txt", archiveFile))
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

	err = utils.SaveConfig(&utils.Config{
		MaxWorkers: runtime.GOMAXPROCS(0),
		Seeds:      uncrawled,
	})

	if err != nil {
		return err
	}

	err = os.RemoveAll(archiveFile)
	if err != nil {
		return err
	}

	log.Println("Finished local archival ...")
	return nil
}

func createArchive(archiveName string) (*zip.Writer, error) {
	af, err := os.Create(fmt.Sprintf("%s.zip", archiveName))
	if err != nil {
		return nil, err
	}

	return zip.NewWriter(af), nil
}
