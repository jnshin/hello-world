package main

import (
	"flag"
	"fmt"

	"github.com/jnshin/hello-world/stock/etf"
)

func main() {

	flag.Parse()

	kodex기계조선, err := etf.NewEtf(102960)
	if err != nil {
		fmt.Println("" + err.Error())
	}

	kodex증권, _ := etf.NewEtf(102970)

	fmt.Printf("결과 : %s\n", kodex기계조선.ToString())
	fmt.Printf("결과 : %s\n", kodex증권.ToString())

}
