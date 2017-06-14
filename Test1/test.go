package Test1

import (
	"fmt"
	"os"
)

func main() {
	osTmpDir := os.TempDir()
	fmt.Println("OS default temp is " + osTmpDir)
	fmt.Println("test")
}
