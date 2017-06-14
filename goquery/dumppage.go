package goquery

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/text/transform"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/glog"
)

/* DumpUrl : 주어진 URL 내용을 tag 와 index를 달아 dump */
func DumpUrl(target string, t transform.Transformer, queryString string) {

	/* Set Proxy */

	proxyUrl, err := url.Parse("http://www-proxy.jp.oracle.com:80")

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

	doc.Find("dl").Find("dt").Each(func(i int, s *goquery.Selection) {
		band, _ := s.Html()
		glog.Info("", i, " - ", band)
		fmt.Printf("%d - %s\n", i, band)
	})   

	glog.Flush()
	return

}
