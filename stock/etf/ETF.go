package etf

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
	"github.com/jnshin/hello-world/jnshin"
	"github.com/jnshin/hello-world/stock"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

type Etf struct {
	stock.Stock
	//
	// history_url *url.URL  // history는 다음에하자.
	Brand        string // Brand. 운영사 코드
	ExpenseRatio float32
	Nav          int // id on_board_last_nav
	Perf1Mon     float32
	Perf3Mon     float32
	Perf6Mon     float32
	Perf12Mon    float32
}

func NewEtf(code int) (*Etf, error) {
	// e := &Etf{stock.Stock: stock.Stock{종목코드: code}}
	e := &Etf{Stock: stock.Stock{Code: code}}
	err := FetchEtf(e)
	return e, err
}

/* FetchEtf : naver stock에서 ETF의 시세 및 Nav를 가져온다.
   argument : e는 null이 아니어야하며, e.Code에 조회할 대상의 주식 코드를 미리 넣어야한다.
   보통의 경우 Etf.NewEtf() 를 통해 호출될 것을 기대한다.
   parse 과정에 문제가 있다면 glog v level 2 이상을 설정할 경우 info log에 parse 대상을 기록한다.

*/
func FetchEtf(e *Etf) error {

	var parseFailed bool = false

	if e.Code == 0 {
		return fmt.Errorf("FetchEtf : 잘못된 종목 코드. 0")
	} else {
		glog.V(2).Info("start FetchEtf : ", e.Code)
	}

	/*
		proxyUrl, err := url.Parse("http://www-proxy.jp.oracle.com:80")

	*/

	// client := &http.Client{Timeout: 4 * time.Second, Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	client := &http.Client{Timeout: 4 * time.Second}

	/* naver의 다른 페이지가 아닌 coinfo.nhn 을 사용한 것은 ETF history data를 가져올 목적으로 선택 했음.
	   history data를 parse하는 것은 나중에... */
	target := "http://finance.naver.com/item/coinfo.nhn?code=" + strconv.Itoa(e.Code)
	glog.V(2).Info("조회대상 URL : ", target)
	resp, err := client.Get(target)

	if err != nil {
		glog.Error("대상 URL 조회 실패 : ", target, " - Error : ", err.Error())
		parseFailed = true
		return err
	}

	defer resp.Body.Close()

	/* resp.Body는 reader type */

	var doc *goquery.Document

	// EUCKR 을 UTF8 로 변환 및 goquery doc open
	doc, err = goquery.NewDocumentFromReader(transform.NewReader(resp.Body, korean.EUCKR.NewDecoder()))

	if err != nil {
		glog.Error("Failed to open goquery doc. \n", err.Error())
		parseFailed = true
	}

	/* stock.Stock 의 기본 값 채우기 */
	doc.Find("dl").First().Children().Each(func(i int, s *goquery.Selection) {
		contents, err := s.Html()
		if err != nil {
			glog.Errorf("goquery failed at Html() : %v\n", err.Error())
			parseFailed = true
			return
		}

		/* 빈줄은 무시하자 */
		if len(contents) == 0 {
			return
		}

		words := strings.Split(contents, " ")

		// default는 일부러 두지않았음.
		switch words[0] {
		case "종목명":
			e.Name = strings.Join(words[1:], " ")
		case "종목코드":
			if tmpVal, err := strconv.Atoi(words[1]); tmpVal != e.Code || err != nil {
				if err != nil {
					glog.Errorf("종목코드 변환 실패. [%v]. %v\n", words[1], err.Error())
					parseFailed = true
				}
				glog.Errorf("가져온 page의 종목코드와 요청한 종목코드 불일치.\n")
				parseFailed = true
			} else {
				// 거래소 구분
				if words[2] == "코스피" {
					e.Market = "코스피"
				} else if words[2] == "코스닥" {
					e.Market = "코스닥"
				} else {
					glog.Errorf("알수 없는 거래소. %v\n", words[2])
					parseFailed = true
				}
			}
		case "현재가":
			if tmpVal, err := jnshin.Cntoi(words[1]); err != nil {
				glog.Errorf("현재가 변환 실패 [%v]. %v\n", words[1], err.Error())
			} else {
				e.Price = tmpVal
			}
		case "전일가":
			if tmpVal, err := jnshin.Cntoi(words[1]); err != nil {
				glog.Errorf("전일가 변환 실패 [%v]. %v\n", words[1], err.Error())
			} else {
				e.YPrice = tmpVal
			}
		case "시가":
			if tmpVal, err := jnshin.Cntoi(words[1]); err != nil {
				glog.Errorf("시가 변환 실패 [%v]. %v\n", words[1], err.Error())
			} else {
				e.Open = tmpVal
			}
		case "고가":
			if tmpVal, err := jnshin.Cntoi(words[1]); err != nil {
				glog.Errorf("고가 변환 실패 [%v]. %v\n", words[1], err.Error())
			} else {
				e.DayHigh = tmpVal
			}
		case "상한가":
			if tmpVal, err := jnshin.Cntoi(words[1]); err != nil {
				glog.Errorf("상한가 변환 실패 [%v]. %v\n", words[1], err.Error())
			} else {
				e.DayUpperLimit = tmpVal
			}
		case "저가":
			if tmpVal, err := jnshin.Cntoi(words[1]); err != nil {
				glog.Errorf("저가 변환 실패 [%v]. %v\n", words[1], err.Error())
			} else {
				e.DayLow = tmpVal
			}
		case "하한가":
			if tmpVal, err := jnshin.Cntoi(words[1]); err != nil {
				glog.Errorf("하한가 변환 실패 [%v]. %v\n", words[1], err.Error())
			} else {
				e.DayLowerLimit = tmpVal
			}
		case "거래량":
			if tmpVal, err := jnshin.Cntoi(words[1]); err != nil {
				glog.Errorf("거래량 변환 실패 [%v]. %v\n", words[1], err.Error())
			} else {
				e.Volume = tmpVal
			}
		case "거래대금":
			// 거래대금은 무시
			return

		}

		glog.V(2).Info("", i, " - ", contents)

	}) /* stock.Stock 기본값 채우기 끝 */

	/* Nav 가져오기 */
	doc.Find("#on_board_last_nav").Children().Children().Each(func(i int, s *goquery.Selection) {

		contents, err := s.Html()
		if err != nil {
			glog.Errorf("goquery failed at Html(), Nav 가져오기 : %v\n", err.Error())
			parseFailed = true
			return
		}

		// fmt.Printf("[%v]\n", contents)

		/* 빈줄은 무시하자 */
		if len(contents) == 0 {
			return
		}

		e.Nav, err = jnshin.Cntoi(contents)
		if err != nil {
			glog.Errorf("Parsing Nav 실패. [%s] : %v\n", contents, err.Error())
			parseFailed = true
		}

	}) /* End of 'Nav 가져오기' */

	/* 펀드보수 가져오기 */
	doc.Find("[summary='펀드보수 정보']").Find("td").Find("em").Each(func(i int, s *goquery.Selection) {

		contents, err := s.Html()
		if err != nil {
			glog.Errorf("goquery failed at Html(), 펀드보수 가져오기 : %v\n", err.Error())
			parseFailed = true
			return
		}

		// fmt.Printf("[%v]\n", contents)

		/* 빈줄은 무시하자 */
		if len(contents) == 0 {
			return
		}

		tmpER, err := jnshin.Cntof(strings.Replace(contents, "%", "", -1))
		if err != nil {
			glog.Errorf("Parsing 펀드보수 실패. [%s] : %v\n", contents, err.Error())
			parseFailed = true
		}
		e.ExpenseRatio = float32(tmpER)

	}) /* End of '펀드보수 가져오기' */

	glog.Flush()

	if parseFailed {
		return fmt.Errorf("주가정보 조회 실패. glog error 참조.\n")
	} else {
		return nil
	}
} /* End of FetchEtf */

func (e *Etf) ToString() string {
	return e.Stock.ToString() + fmt.Sprintf("Nav %d ExpenseRatio %4.2f%%", e.Nav, e.ExpenseRatio)
}
