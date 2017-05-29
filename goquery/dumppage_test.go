package goquery

import (
	"testing"

	"golang.org/x/text/encoding/korean"
)

func TestFetchLines(t *testing.T) {
	DumpUrl("http://finance.naver.com/item/main.nhn?code=102960", korean.EUCKR.NewDecoder(), "*")
}
