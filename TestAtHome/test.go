package main

import (
	"fmt"
	"log"

	"github.com/doneland/yquotes"
)

// test


/*
  yahoo page 일부가 동작이 불안정해 보인다.
	또 한국 장 마감 후 자정 경에 모든 값을 0로 클리어하기 떄문에 종가 등 정보에 오류도 불러올 수 있다.
	차라리 goquery를 이용해 naver page를 parse하는게 더 좋을 것 같다.
*/
func main() {
	fmt.Println("Test@home")

	stock, err := yquotes.NewStock("102960.KS", false)
	if err != nil {
		log.Fatalf("Error on NewStock : %s", err)
	}

	fmt.Printf("Symbol : %s, name : %s, price %f",
		stock.Symbol, stock.Name, stock.History
		stock.Price)
}
