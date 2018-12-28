package stock

/*
import (
	"fmt"
)
*/

type MarketProvider interface {

	// getQuote : 조회하고자하는 stock을 생성해 전달 
	//   종목을 구분하기 위한 최소한의 정보가 포함되어야 함. 예) 종목코드
	quote(s *Stock) (s *Stock, err error)

	// quoteList : 상장종목 리스트를 반환. 단 가격정보는 포함하지 않는다
	quoteList() (s []Stock)

	// quoteListWithPrice : quoteList 와 동일한 목록을 반환하지만 종목의 현재 가격을 포함한다
	quoteListWithPrice() (s []Stock) 

	// isDelayedQuote : quote 조회에 delay가 존재하면 지연되는 초를 반환
	isDelayedQuote() (delay int)

	// providerName : data provider 이름을 반환
	providerName() (name string)

	newfunc()

}