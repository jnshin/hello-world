package stock

import (
	"fmt"
)


// MarketProvider 는 stock quote를 제공하는 provider의 기본 기능을 정의한다
// naver 또는 yahoo와 같이 data source에 따라 구현을 달리한다.
// ToDo : 정보제공 범위에 대한 config 정보를 표현할 방법 필요.
// 예를들어 어디는 수정주가를 제공하고 어디는 아닌 값을 제공한다거나... 
// 어디는 history 값을 가져올 방법이 없다거나.... 
// 장기적으로는 반복적인 GetQuote 를 너무 빈번하게 호출하지 못하도록 제한할 필요 있다. 
type MarketProvider interface {

	// GetQuote : 조회하고자하는 stock을 생성해 전달
	//   종목을 구분하기 위한 최소한의 정보가 포함되어야 함. 예) 종목코드
	GetQuote(i *Stock) (*Stock, error)

	// quoteList : 상장종목 리스트를 반환. 단 가격정보는 포함하지 않는다
	quoteList() (s []Stock)

	// quoteListWithPrice : quoteList 와 동일한 목록을 반환하지만 종목의 현재 가격을 포함한다
	quoteListWithPrice() (s []Stock)

	// isDelayedQuote : quote 조회에 delay가 존재하면 지연되는 초를 반환
	isDelayedQuote() (delay int)

	// providerName : data provider 이름을 반환한다.
	providerName() (name string)

}

// Naver Stock 

type NaverStock struct {}

func (p *NaverStock) GetQuote(s *Stock) (*Stock, error) {
	fmt.Printf("구현해라. NaverStock.GetQuote")
	return s, nil
}

// Yahoo Finance 

type YahooFinance struct {}

func (p *YahooFinance) GetQuote(s *Stock) (*Stock, error) {
	fmt.Printf("구현해라. YahooFinance.GetQuote")
	return s, nil
}

