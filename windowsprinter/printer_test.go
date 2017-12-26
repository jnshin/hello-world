package windowsprinter

import (
	"fmt"
	"testing"

	"github.com/golang/glog"
)

func TestGetStatus(t *testing.T) {

	// var p = new(Printer{Name: "Test"})
	// 2017.09.04 36FMFD3-MP4054는 offline 상태.
	p := NewPrinter("36FMFD3-MP4054")
	p = p.GetStatus()
	// 정상적으로 실행될 것이라 기대
	// printer driver가 설치만 되어 있다면....  printer 상태와 무관함.
	if p == nil {
		t.Error("printer is nil. - 36FMFD3-MP4054")
	} else {
		fmt.Println("OK - 36FMFD3-MP4054")
		fmt.Printf("%v\n", p)
	}

	// var p = new(Printer{Name: "Test"})
	// 2017.09.04 36F2C-L8350 은 임시로 사용하는 color printer
	p = NewPrinter("36F2C-L8350")
	p = p.GetStatus()
	// 정상적으로 실행될 것이라 기대
	// printer driver가 설치만 되어 있다면....  printer 상태와 무관함.
	if p == nil {
		t.Error("printer is nil. - 36F2C-L8350")
	} else {
		fmt.Println("OK. - 36F2C-L8350")
		fmt.Printf("%v\n", p)
	}

	// 존재하지 않는 printer
	p = NewPrinter("NoSuchPrinter")
	p = p.GetStatus()
	// 없는 printer이니 실패하는 것이 당연. 성공하면 오히려 이상함.
	if p == nil {
		fmt.Println("OK - NoSuchPrinter")
	} else {
		t.Error("ERROR!")
	}

	glog.Flush()
}

func TestPrinterList(t *testing.T) {
	pl := GetPrinterList(true)
	if len(pl) == 0 {
		t.Error("Error! : GetPrinterList returns nil!")
	}

	for _, printer := range pl {

		if printer.IsDefault {
			fmt.Printf("Printer : %v is default printer.\n", printer.Name)
		} else {
			fmt.Printf("Printer : %v\n", printer.Name)
		}
	}

	fmt.Printf("Done - TestPrinterList\n")
}
