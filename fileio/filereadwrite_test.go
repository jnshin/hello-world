package fileio

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var sampleFile = "sampledata.txt"

func TestFetchLines(t *testing.T) {
	var dataChan <-chan string
	var counter int
	var err error

	dataChan, err = FetchLines(sampleFile, "euckr")
	if err != nil {
		log.Fatalln("Error : " + err.Error())
	}

	counter = 0
	for a := range dataChan {
		counter++
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
