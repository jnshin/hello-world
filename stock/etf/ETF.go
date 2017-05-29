package etf

import (
  "github.com/jnshin/hello-world/stock"
  "net/url"
)

type Etf struct {
	Stock
	//
	// history_url *url.URL  // history는 다음에하자.
	펀드보수        float32
	nav         int // id on_board_last_nav
	수익률1개월      float32
	수익률3개월      float32
	수익률6개월      float32
	수익률12개월     float32
}

func NewStock(code int) (*Etf, error) {
  e := Etf{종목코드: code}
  err := FetchEtf(&e)
  return e, err
}

func FetchEtf(e *Etf) error {
  if e.종목코드 = 0 {
    return fmt.Errorf("FetchEtf : 잘못된 종목 코드. 0")
  }

  proxyUrl, err := url.Parse("http://www-proxy.jp.oracle.com:80")


  // http client는 사용 환경에 따라 설정이 다를 수 있으니, 이렇게 코드에서 일방적으로 설정하는 것은 옳지 않아 보임.
  //
  client := &http.Client{Timeout: 4 * time.Second, Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
  glog.V(2).Info("start dumpURL : ", target)
  resp, err := client.Get(target)

  if err != nil {
    glog.Error("Unable to fetch given url : ", target, " - Error : ", err.Error())
    return
  }

  defer resp.Body.Close()

  /* resp.Body는 reader type */

  var doc *goquery.Document

  if t != nil {
    doc, err = goquery.NewDocumentFromReader(transform.NewReader(resp.Body, t))
  } else {
    doc, err = goquery.NewDocumentFromReader(resp.Body)
  }

  if err != nil {
    glog.Error("Failed to open goquery doc. ", err.Error())
    return
  }

  doc.Find("dl").Find("dt").EachWithBreak(func(i int, s *goquery.Selection) {
    band, _ := s.Html()
    glog.Info("", i, " - ", band)
    fmt.Printf("%d - %s\n", i, band)
  })

  glog.Flush()


}
  종목코드는() int
	종목명는() string
	현재가는() int
	시가는() int
	고가는() int
	저가는() int
	거래량은() int
