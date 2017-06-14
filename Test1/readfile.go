package Test1

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// 팁                    

// 파일 사용 후 filesystem에서 자동 삭제되도록 하려면, file 생성 후 handle을 남긴 상태에서 unlink 해 버리면 된다.
// memory 에만 상주시키려면, linux는 tmpfs를, windows는 FILE_ATTRIBUTE_TEMPORARY  를 고려해 보자.

func checkDirIsEmpty(dir string) (int, bool) {
	var fileCount int
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for index, file := range files {
		fmt.Printf("Index %d - name [%s]", index, file)
		fileCount = index
	}

	if fileCount == 0 {
		return 0, true
	} else {
		return fileCount, false
	}

}

// CreateTempFile : 임시 directory를 생성하고 그 안에 파일을 생성한다.
func CreateTempFile() {

	content := []byte("Temporary file's content")
	dir, err := ioutil.TempDir("", "golang_test")
	if err != nil {
		log.Fatal(err)
	}

	if fileCount, isEmpty := checkDirIsEmpty(dir); isEmpty {
		fmt.Println("New temporary directory " + dir + " is empty.")
	} else {
		fmt.Println("Directory " + dir + " contains some files. file count is " + strconv.Itoa(fileCount))
	}
	// defer os.RemoveAll(dir)

	// tmpfn := filepath.Join(dir, "tmpfile")

	tmpfile, err := ioutil.TempFile(dir, "tmpfile")
	if err != nil {
		log.Fatal(err)
	}

	// tmp file 생성 후 즉시 remove하고 write 해 보자.
	fmt.Println("만들어진 파일 지워보자. " + tmpfile.Name())

	// 그런데 생성 후 remove 시도하면 사용 중이라 지워지지 않는다.
	// 이런 경우 error handling은 어떻게 하지? error 내용이 뭔지 알아야하는데.
	if err := os.Remove(tmpfile.Name()); err != nil {

		if pError, ok := err.(*os.PathError); ok {
			fmt.Println("OK. patherror. " + pError.Op + ", " + pError.Err.Error())
			fmt.Println("Ignore remove error.")
		} else {
			log.Fatal(err)
		}

	}

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Write... [" + string(content) + "]")
	}

	// temp file의 처음으로 이동.
	if _, err := tmpfile.Seek(0, os.SEEK_SET); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Seek...")
	}

	// read buffer를 slice로 할당하는데 실제 영역을 할당해 주지 않으면
	// FILE.Read 는 len 0 의 slice readBuf를 argument로 받게되고,
	// len 0 이면 Read 는 에러 없이 읽지 않고 read bytes를 0로 반환한다.

	var readBuf []byte

	for i := 0; ; i++ {

		// slice를 초기화하는 것이 효과적일까? 아니면 새로 할당해 버릴까?
		// gc를 믿고 그냥 새로 할당해 봐...
		readBuf = make([]byte, 10)
		// Read 는 argument로 받는 slice의 크기 만큼만 읽는다.
		readBytes, err := tmpfile.Read(readBuf)
		if err == io.EOF {
			fmt.Println("No more data. break.")
			break // Breaking current loop
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Read result " + strconv.Itoa(i) + " : [" + string(readBuf) + "] read bytes : " + strconv.Itoa(readBytes))
	} // End of for loop

}
