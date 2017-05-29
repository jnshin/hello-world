package main

import (
	"flag"

	"golang.org/x/text/encoding/korean"

	"github.com/jnshin/hello-world/goquery"
)

func main() {

	flag.Parse()

	goquery.DumpUrl("http://finance.naver.com/item/main.nhn?code=102960", korean.EUCKR.NewDecoder(), "*")
}
