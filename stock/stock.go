package stock

import (
	"fmt"
	"net/url"
)

// Stock 상장 주식의 기본값을 포함한다.
// 값의 생성은 DataProvider를 통해 가져온다.
type Stock struct {
	Code          int      // 종목코드
	URL           *url.URL // data source URL
	Name          string   // 종목명
	Market        string   // 코스피 / 코스닥
	Price         int      // 현재가 또는 종가
	YPrice        int      // 전일가
	Open          int      // 시가
	DayHigh       int      // 고가
	DayLow        int      // 저가
	DayUpperLimit int      // 상한가
	DayLowerLimit int      // 하한가
	Volume        int      // 거래량
}

// ToString 은 Stock 의 기본 정보를 string으로 전환한다.
func (e *Stock) ToString() string {
	return fmt.Sprintf("%s(%d, %s) : 현재가 %v원, 가격폭 %v~%v 상하한[%v,%v] 거래량 %v주\n", e.Name, e.Code, e.Market, e.Price, e.DayLow, e.DayHigh, e.DayLowerLimit, e.DayUpperLimit, e.Volume)
}
