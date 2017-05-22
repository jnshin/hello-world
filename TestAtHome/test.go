package main

import (
	"fmt"
	"log"

	"github.com/doneland/yquotes"
)

// test

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
