package fileio

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang/glog"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

/* FetchLines : 지정한 파일에서 data를 읽어 line 단위로 channel에 넣어 반환한다.
   만약 code를 제공하면 file data를 해당 code로 간주하고 UTF8로 변환해 반환한다.
	 현재는 'EUCKR'만 구현했다.
*/
func FetchLines(filename string, code string) (<-chan string, error) {

	// file check 할 것

	var file *os.File
	var err error

	const bufferSize = 8192

	if file, err = os.Open(filename); err != nil {
		glog.Fatal("Unable to open file : ", filename)
	}

	// NewReaderSize의 최소 buffer size는 16 byte 이다. 이하 크기로 설정시 자동으로 16.
	reader := bufio.NewReaderSize(file, bufferSize)

	if len(code) > 0 {
		switch strings.ToUpper(code) {
		case "EUCKR":
			// transform.NewReader 는 내부에서 buffer size를 4K로 조정한다.
			reader = bufio.NewReader(transform.NewReader(reader, korean.EUCKR.NewDecoder()))
		default:
			glog.Error("Not Implemented code page.", code)
			return nil, fmt.Errorf("Given code[%v] is not implemented yet.\n", code)

		}
	} else {
		reader = bufio.NewReader(reader)
	}

	out := make(chan string)
	go func() {
		var buf bytes.Buffer
		var isFirst, isPrefix bool = true, false
		var line []byte
		var err error
		for {
			// transform에서 생성한 reader는 buffer를 4k로 할당하기에
			// file open 시점에 할당한 buffer size가 무시된다.
			// test하려면 4k 넘게 할당해야한다.
			if line, isPrefix, err = reader.ReadLine(); err != nil {
				if err == io.EOF {
					close(out)
					break
				}
				glog.Fatal("Reading error. ", err)
			}

			if isPrefix == true {
				// buf.WriteString(string(line))   // string으로 변환하지 말고 그대로 []bytes를 넣어도 될 듯.
				buf.Write(line)
				isFirst = false
			} else if isFirst == true {
				out <- string(line)
			} else {
				buf.Write(line)
				out <- buf.String()
				buf.Reset()
				isFirst = true
			}
		} // End of for loop
	}()
	return out, nil // end of
}
