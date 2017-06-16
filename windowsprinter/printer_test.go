package windowsprinter

import (
	"fmt"
	"testing"
)

func TestGetStatus(t *testing.T) {

	// var p = new(Printer{Name: "Test"})
	p := NewPrinter("36FMFD3-MP4054")
	p = p.GetStatus()
	if p == nil {
		fmt.Println("printer is nil.")
	} else {
		fmt.Println("OK.")

		fmt.Printf("%v\n", p)
	}

	// 존재하지 않는 printer
	p = NewPrinter("NoSuchPrinter")
	p = p.GetStatus()
	if p == nil {
		fmt.Println("OK")
	} else {
		fmt.Println("ERROR!")
	}
}
