package main

import (
	"fmt"
	"os"
)

func main() {
	osTmpDir := os.TempDir()
	fmt.Println("OS default temp is " + osTmpDir)
}
