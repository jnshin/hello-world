package fileio

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/golang/glog"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

/* FetchLines : 주어진 파일을 처음부터 읽어 한 줄씩 channel로 return해 준다. */
func FetchLines(filename string, code string) <-chan string {

	// file check 할 것

	var reader *bufio.Reader
	var file *os.File
	var err error

	if file, err = os.Open(filename); err != nil {
		glog.Fatal("Unable to open file : ", filename)
	}

	if len(code) > 0 {
		switch strings.ToUpper(code) {
		case "EUCKR":
			reader = bufio.NewReader(transform.NewReader(file, korean.EUCKR.NewDecoder()))
		default:
			glog.Fatal("Not Implemented code page.", code)
		}
	} else {
		reader = bufio.NewReader(file)
	}

	out := make(chan string, 10)
	go func() {
		var buf bytes.Buffer
		var isFirst, isPrefix bool = false, false
		var line []byte
		var err error
		for {
			if line, isPrefix, err = reader.ReadLine(); err != nil {
				if err == io.EOF {
					close(out)
					break
				}
				glog.Fatal("Reading error. ", err)
			}
			if isPrefix == true {
				buf.WriteString(string(line))
				isFirst = false
			} else if isFirst == true {
				out <- string(line)
			} else {
				buf.WriteString(string(line))
				out <- buf.String()
				buf.Reset()
			}
		}
	}()
	return out // end of
}
