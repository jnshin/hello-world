package fileio

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var sampleFile string = "sampledata.txt"

func TestFetchLines(t *testing.T) {
	var dataChan <-chan string
	var counter int

	dataChan = FetchLines(sampleFile, "euckr")

	counter = 0
	for a := range dataChan {
		counter += 1
		fmt.Printf("Data - [%v]\n", a)
	}
}

func TestMain(m *testing.M) {

	fileInfo, err := os.Stat(sampleFile)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("Sample file does not exists.")
		}
	}
	_ = fileInfo
	code := m.Run()
	os.Exit(code)
}
