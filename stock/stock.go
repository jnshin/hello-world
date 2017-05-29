package stock

import "net/url"

type Stock struct {
	종목코드 int      // 종목코드
	URL  *url.URL // data source URL
	종목명  string   // 종목명
	현재가  int      // 현재가 또는 종가
	전일가  int      // 전일가
	시가   int      // 시가
	고가   int      // 고가
	저가   int      // 저가
	상한가  int      // 상한가
	하한가  int      // 하한가
	거래량  int      // 거래량
}

type StockInterface interface {
	종목코드는() int
	종목명는() string
	현재가는() int
	시가는() int
	고가는() int
	저가는() int
	거래량은() int
}

type 
