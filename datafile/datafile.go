package datafile

import (
	"fmt"
	"os"
	"path/filepath"
)

var baseDir string

func init() {
	baseDir = os.TempDir()
}

// SetBaseDir : 기본 작업 위치를 지정
// 미설정시 os의 기본 temporary directory.
func SetBaseDir(dir string) {
	baseDir = dir
}

// OpenDatafile : baseDir에 file을 생성 또는 append mode로 open 함.
func OpenDatafile(filename string) (f *os.File, err error) {
	if len(filename) == 0 {
		return nil, fmt.Errorf("OpenDatafile : filename is null")
	}

	absFilename := filepath.Join(baseDir, filename)

	f, err = os.OpenFile(absFilename, os.O_RDWR+os.O_APPEND+os.O_CREATE, 0666)

	return
}
