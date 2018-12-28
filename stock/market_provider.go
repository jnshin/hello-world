package stock

/*
import (
	"fmt"
)
*/

// MarketProvider 는 stock quote를 제공하는 provider의 기본 기능을 정의한다
// naver 또는 yahoo와 같이 data source에 따라 구현을 달리한다.
type MarketProvider interface {

	// getQuote : 조회하고자하는 stock을 생성해 전달
	//   종목을 구분하기 위한 최소한의 정보가 포함되어야 함. 예) 종목코드
	quote(i Stock) (Stock, error)

	// quoteList : 상장종목 리스트를 반환. 단 가격정보는 포함하지 않는다
	quoteList() (s []Stock)

	// quoteListWithPrice : quoteList 와 동일한 목록을 반환하지만 종목의 현재 가격을 포함한다
	quoteListWithPrice() (s []Stock)

	// isDelayedQuote : quote 조회에 delay가 존재하면 지연되는 초를 반환
	isDelayedQuote() (delay int)

	// providerName : data provider 이름을 반환한다.
	providerName() (name string)
}
