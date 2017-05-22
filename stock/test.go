package main

import (
	"fmt"
)

func main() {
	func1("test")
}

func func1(str string) {
	fmt.Printf("한글 : [%v]\n", str)
}
