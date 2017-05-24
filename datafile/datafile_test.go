package datafile

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestOpenDatafile(t *testing.T) {
	var testFile *os.File
	var err error

	// OpenDatafile 은 temporary directory에 파일을 생성 또는 append mode로 open한다.
	// 생성과 동작을 시험하기 위해 동일 argument로 2번 호출해야한다.
	if testFile, err = OpenDatafile("test_golang"); testFile == nil || err != nil {
		t.Errorf("첫번째 OpenDatafile(test_golang) 호출 실패. testFile[%#v], err[%s]", testFile, err)
	} else {
		t.Logf("첫번째 OpenDatafile(test_golang) 호출 성공. testFile[%#v], err[%s]", testFile, err)
		t.Logf("file path - " + testFile.Name())
	}

	testFile.Close()

	if testFile, err = OpenDatafile("test_golang"); testFile == nil || err != nil {
		t.Errorf("두번째 OpenDatafile(test_golang) 호출 실패. testFile[%#v], err[%s]", testFile, err)
	} else {
		t.Logf("두번째 OpenDatafile(test_golang) 호출 성공. testFile[%#v], err[%s]", testFile, err)
		t.Logf("file path - " + testFile.Name())
	}

	// Cleanup
	testFileName := testFile.Name()
	if err = testFile.Close(); err != nil {
		t.Logf("Cleanup failed. close failed. %s", err)
	}
	if err = os.Remove(testFileName); err != nil {
		t.Logf("Cleanup failed. rm failed. %s", err)
	}

}

func TestSetBaseDir(t *testing.T) {
	var tmpDir string
	// var testFile *os.File
	var err error

	tmpDir, err = ioutil.TempDir("", "golang_test")
	if err != nil {
		t.Errorf("Unable to Create temporary director. err[%s]", err)
	}

	SetBaseDir(tmpDir)
	TestOpenDatafile(t)

	// cleanup
	os.Remove(tmpDir)

}
