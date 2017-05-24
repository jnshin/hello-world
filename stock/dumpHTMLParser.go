package main

/*
  공부 삼아 만들어 보고 든 생각인데, 이렇게 고생해 tokenizer 쓸 필요없이
	goquery를 사용하면 쉽게 처리된다.
*/

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

func dumpToken(url string) {

	client := &http.Client{Timeout: 4 * time.Second}
	link := url
	resp, err := client.Get(link)

	for err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc := html.NewTokenizer(resp.Body)

	found_price_summary := false

	for tokenType := doc.Next(); tokenType != html.ErrorToken; tokenType = doc.Next() {

		switch tokenType {
		case html.ErrorToken:
			doc.Err()
		case html.TextToken:
			fmt.Printf("%s[%s]\n", tokenType, doc.Token())
		case html.StartTagToken:
			tagName, _ := doc.TagName()
			for {
				tagAttr, tagAttrVal, moreAttr := doc.TagAttr()
				fmt.Printf("%s - tagName[%s], tagAttr[%s], tagAttrVal[%s] moreAttr[%v]\n", tokenType, tagName, tagAttr, tagAttrVal, moreAttr)
				if string(tagAttr) == "class" && string(tagAttrVal) == "type2 type_tax" {
					if found_price_summary != true {
						found_price_summary = true
					} else {
						log.Fatal(fmt.Errorf("Internal Error! : There are multiple 'class=type2 type_tax'\n"))
					}
					fmt.Println("Bingo!")

					var thVal, tdVal string
					thVal, tdVal = fetchTHTD(doc, "현재가")
					fmt.Printf("RV : [%v][%v]\n", thVal, tdVal)
					thVal, tdVal = fetchTHTD(doc, "매도호가")
					fmt.Printf("RV : [%v][%v]\n", thVal, tdVal)
					thVal, tdVal = fetchTHTD(doc, "매수호가")
					fmt.Printf("RV : [%v][%v]\n", thVal, tdVal)
				}
				if moreAttr != true {
					break
				}
			} // end of dump attribute
		case html.EndTagToken:
			fmt.Printf("%s[%s]\n", tokenType, doc.Token())
		} // end of switch

	} // end of for main loop
} // end of dumpToken

func fetchTHTD(doc *html.Tokenizer, thStr string) (thVal, tdVal string) {

	var foundTH bool = false
	var foundTD bool = false

	fmt.Printf("Enter fetchTHTD. target[%s] [%v][%v]\n", thStr, thVal, thVal)

	for {

		if len(thVal) > 0 && len(tdVal) > 0 {
			return
		}

		tokenType := doc.Next()
		switch tokenType {
		case html.ErrorToken:
			doc.Err()
		case html.StartTagToken:
			tagName, _ := doc.TagName()
			tagAttr, tagAttrVal, moreAttr := doc.TagAttr()
			fmt.Printf("%s - tagName[%s], tagAttr[%s], tagAttrVal[%s] moreAttr[%v]\n", tokenType, tagName, tagAttr, tagAttrVal, moreAttr)
			if foundTH == false && string(tagName) == "th" {
				foundTH = true
				fmt.Println("foundTH : true")
			} else if foundTH == false {
				continue
			}

			if foundTH == true && foundTD == false && string(tagName) == "td" {
				foundTD = true
				fmt.Println("foundTD : true")
			} else {
				continue
			}

		case html.TextToken:

			eucVal := []byte(doc.Token().Data)
			fmt.Printf("ok Text token : [%v] TH[%v-%v]TD[%v-%v]\n", string(eucVal), foundTH, thVal, foundTD, tdVal)
			reader := transform.NewReader(bytes.NewReader(eucVal), korean.EUCKR.NewDecoder())
			textVal, e := ioutil.ReadAll(reader)
			if e != nil {
				log.Fatalf("EUCKR 변환 실패.")
			}
			strVal := strings.TrimSpace(string(textVal))
			fmt.Printf("decode result [%v]\n", strVal)

			if foundTH == true && foundTD == false {

				if strVal == thStr && len(thVal) == 0 {
					// TH 값을 가져오지 않았으면 가져온다.
					fmt.Println("Found target th - " + thStr)
					thVal = strVal
				} else if len(strVal) == 0 {
					// TH를 가져오는 과정에 null text가 존재하는 경우 무시
					fmt.Println("Ignore null text during TH")
					continue
				} else if strVal != thStr {

					fmt.Printf("Not a target th. target TH[%v] curr[%v]\n", thStr, strVal)
					foundTH = false
					thVal = ""
				} else {
					fmt.Printf("Keep prev thval[%v] new val[%v]\n", thVal, strVal)
				}

			} else if foundTH == true && foundTD == true {
				if len(tdVal) == 0 {
					tdVal = strVal
				} else {
					fmt.Printf("Keep prev tdval[%v] new val[%v]\n", tdVal, strVal)
				}

			}
			// ??? ??????? ???????.
		} // end of switch
	} // end of for loop

}

func main() {
	dumpToken("http://finance.naver.com/item/sise.nhn?code=102960")
}
