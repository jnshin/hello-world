package windowsprinter

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

// Printer monitoring 설정 값.
type PrinterMonitoringConfig struct {
	// printer name. system에 조회되는 printer 이름
	PrinterName string
	// 모니터링 대상 여부
	EnableMonitoring bool `ini:"Monitoring"`
	// test page 출력 간격. 단위 hour
	PrintInterval int `ini:"IntervalHours"`
	// 출력 작업이 있다면 자동으로 timer를 PrintInterval로 reset함.
	// 이 기능은 color printer에서 흑백 출력만 반복하는 경우를 대비해 default disable 상태로 설정한다.
	EnableTimerReset bool
}

// Printer status
// 프린터 상태를 확인하려면 이 값을 참조.
type Printer struct {
	// Printer name
	Name       string
	Attributes int
	// Printer error information
	// 그런데 이 값은 좀 부정확해 보인다. 평상시 0.
	DetectedErrorState         int
	ExtendedPrinterStatus      int
	ExtendedDetectedErrorState int

	// virtual printer 들은 port 이름 끝에 ':'이 붙어 있다. 일반 printer는 USB 또는 network IP가 들어간다.
	// 'LPT1:', 'USB1:' 과 같은 형식이 여러개 나올 수 있다.
	// ':' 로 끝나지 않고 IP 형태로 나올 수 있다. ',' 구분자 사용.
	PortName []string
}

// description for printer.attributes
var printerAttrDesc = map[int]string{
	// bitmap of attributes for a windows-based printing device
	0x1:    "queued - Print jobs are buffered and queued",
	0x2:    "Direct - Document to be sent directly to the printer",
	0x4:    "Default - Default printer on a computer",
	0x8:    "Shared - Available as a shared network resource",
	0x10:   "Network - Attached to a network. If both local and Network bits are set, this indicates a network printer",
	0x20:   "Hidden - Hidden from some users on the network",
	0x40:   "Local - Directly connected to a computer",
	0x80:   "EnableDevQ - Enable the queue on the printer if Available",
	0x100:  "KeepPrintedJobs - Spooler should not delete documents after they are printed",
	0x200:  "DoCompleteFirst - Start jobs that are finished spooling first",
	0x400:  "WorkOffline - Queue print jobs when a printer is not available",
	0x800:  "EnableBIDI - Enable bidirectional printing",
	0x1000: "Allow only raw data type jobs to be SpoolEnabled",
	0x2000: "Published - Published in the network directory service"}

// description for printer.ExtendedDetectedErrorState
var printerExtErrStateDesc = map[int]string{
	// Report standard error information
	0:  "Unknown",
	1:  "Other",
	2:  "No Error",
	3:  "Low Paper",
	4:  "No Paper",
	5:  "Low Toner",
	6:  "No Toner",
	7:  "Door open",
	8:  "Jammed",
	9:  "Service Requested",
	10: "Output Bin Full",
	11: "Paper Problem",
	12: "Cannot Print page",
	13: "User intervention Required",
	14: "Out of memory",
	15: "Server Unknown"}

var extendedPrinterStatusDesc = map[int]string{
	// Status information for a printer that is different from information specified in the availability property
	1:  "Other",
	2:  "Unknown",
	3:  "Idle",
	4:  "Printing",
	5:  "Warming Up",
	6:  "Stopped Printing",
	7:  "Offline",
	8:  "Paused",
	9:  "Error",
	10: "Busy",
	11: "Not available",
	12: "Waiting",
	13: "Processing",
	14: "Initialization",
	15: "Power Save",
	16: "Pending Deletion",
	17: "I/O Active",
	18: "Manual Feed"}

// GetStatus : it returns new printer instance.
func (p *Printer) GetStatus() *Printer {

	// printer 껍데기 생성
	t := NewPrinter(p.Name)
	if t = getPrinterStatus(t); t == nil {
		glog.Error("Unable to get printer status : ", p.Name)
		return nil
	}
	glog.Flush()

	// ToDo
	// 받아온 status와 기존 P를 비교해 동작 또는 상태 변화가 있으면 조치해야...

	return t
}

// GetPrinterQueueStatus : it returns new printer spool statistics.
func (p *Printer) GetPrinterQueueStatus() *PrinterQueueStatus {

	return nil

}

// NewPrinter : 새로운 printer instance를 생성.
func NewPrinter(name string) *Printer {
	var t Printer
	t.Name = name
	return &t
}

func getWMICExitCode(err error) int {
	var errcode int
	errmsg := strings.Fields(err.Error())
	if errmsg[0] == "exit" && errmsg[1] == "status" {
		errcode, _ = strconv.Atoi(errmsg[2])
	} else {
		errcode = -1
	}

	glog.V(2).Info("WMIC error code : %v\n", errcode)
	return errcode
}

func convStrToInt(mp map[string]string, key string) int {

	rv, err := strconv.Atoi(mp[key])
	if err != nil {
		glog.Fatal("Atoi failed. ", key, "-", mp[key])
	}

	return rv
}

// getPrinterStatus : 실제 wmic 명령으로 windows 상태 값을 생성한다.
// caller는 argument p 에 name을 채워서 넘겨야한다.
func getPrinterStatus(p *Printer) *Printer {

	if len(p.Name) == 0 {
		glog.Error("Missing printer name")
		return nil
	}

	pm := queryWMI("printer", "where", "Name='"+p.Name+"'", "get", "/format:list")
	if pm == nil {
		return nil
	}

	p.Attributes = convStrToInt(pm, "Attributes")
	p.DetectedErrorState = convStrToInt(pm, "DetectedErrorState")
	p.ExtendedDetectedErrorState = convStrToInt(pm, "ExtendedDetectedErrorState")
	p.ExtendedPrinterStatus = convStrToInt(pm, "ExtendedPrinterStatus")
	// portname은 ',' 구분자로 다중 값이 올 수 있다. 보통은 1개 값만 존재.
	p.PortName = strings.Split(pm["PortName"], ",")

	return p
}

func queryWMI(args ...string) map[string]string {

	var errmsg string
	// wmic printer where "Name='36FMFD3-MP4054'" get /format:list
	// cmd := exec.Command("wmic", "printer", "where", "Name='"+p.Name+"'", "get", "/format:list")
	cmd := exec.Command("wmic", args...)
	// cmd := exec.Command("notepad")
	// fmt.Println("wmic", " ", "printer", " ", "where", " ", "Name='"+p.Name+"'", " ", "get", " ", "/format:list")

	var out bytes.Buffer
	fatalError := false
	cmd.Stdout = &out
	// 잘못된 대상에 대해 동작하는 경우 stderr에만 출력하고
	// stdout에는 빈줄만 들어가는 것으로 보여진다.
	// 제대로하려면 stderr로 살펴야하는데...
	// 여기서는 그냥 stdout이 빈줄이면 잘못된 것으로 간주하자.
	err := cmd.Run()
	if err != nil {
		errcode := getWMICExitCode(err)
		switch errcode {
		case 2147749911: // WBEM_E_INVALID_QUERY - Query was not syntactically valid
			fatalError = true
			errmsg = "Invalid WMI Query"
		default:
			fatalError = true
			errmsg = "Unexpected error"
		}

		if fatalError {
			glog.Fatalf("Error : unable to run wmic. [%s][%v]", err, errmsg)
		}
		return nil
	} // cmd.run() 에러 처리 끝.

	glog.V(3).Info("output [" + out.String() + "]")
	// fmt.Println("output [" + out.String() + "]")

	// wmic 명령의 결과가 공백인 경우.
	// 잘못된 대상에대해 실행하는 경우 등.
	if len(strings.TrimSpace(out.String())) == 0 {
		glog.V(2).Info("흠... 결과 값이 공백이야.")
		return nil
	} else {
		glog.V(2).Info("[%v]\n", strings.TrimSpace(out.String()))
	}
	r := bufio.NewReader(bytes.NewBuffer(out.Bytes()))
	// r := bufio.NewReader(out)

	rv := make(map[string]string)

	var (
		line []byte
	)
	for {
		// transform에서 생성한 reader는 buffer를 4k로 할당하기에
		// file open 시점에 할당한 buffer size가 무시된다.
		// test하려면 4k 넘게 할당해야한다.
		if line, _, err = r.ReadLine(); err != nil {
			if err == io.EOF {
				break
			}
			glog.Fatal("Reading error. ", err)
		}

		lineStr := strings.TrimSpace(string(line))
		if len(lineStr) == 0 {
			continue
		}

		tokens := strings.Split(lineStr, "=")

		glog.V(2).Info("%s - %s\n", tokens[0], tokens[1])
		if len(tokens[1]) > 0 {
			rv[tokens[0]] = tokens[1]
		}

	} // End of for loop.

	return rv

	// return nil
}

// PrinterQueueStatus : 프린터의 출력 건수 및 상태 변화를 보려면 이 값을 활용.
type PrinterQueueStatus struct {
	Name string
	// Current number of jobs in a print queue
	Jobs int
	// Current number of spooling jobs in a print queue  <-- 이건 필요 없을 듯...
	JobsSpooling      int
	TotalJobsPrinted  int
	TotalPagesPrinted int
	// Total number of out-of-papers erros in a print queue after the last restart.
	OutofPaperErrors int
	// Total number of printer-not-ready errors in a print queue after the last restart.
	NotReadyErrors int
}

type monitoringPrinters struct {
	printers []Printer
}
