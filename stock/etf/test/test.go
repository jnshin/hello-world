package main

import (
	"flag"
	"fmt"

	"github.com/jnshin/hello-world/stock/etf"
)

func main() {

	flag.Parse()

	/*************************************************************

	  일반 인터넷 환경에서는 아래 proxy 설정이 필요 없습니다.
		혹시 http proxy 설정이 필요한 경우만 설정합니다.

		*************************************************************/

	/*
		proxyUrl, err := url.Parse("http://www-proxy.jp.oracle.com:80")
		if err != nil {
			glog.Errorf("Error : proxyUrl 생성 실패. %v\n", err.Error())
			glog.Flush()
			os.Exit(1)
		}

		http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	*/

	/* HTTP proxy 설정 마지막 줄 */

	kodex기계조선, err := etf.NewEtf(102960)
	if err != nil {
		fmt.Println("" + err.Error())
	}

	kodex증권, _ := etf.NewEtf(102970)

	fmt.Printf("결과 : %s\n", kodex기계조선.ToString())
	fmt.Printf("결과 : %s\n", kodex증권.ToString())

}
