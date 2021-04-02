package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url = flag.String("url", "", "环境参数")

func main() {
	flag.Parse()
	if *url == "" {
		*url = "http://192.168.2.200"
	}
	count := 0

	for {
		count++
		response, err := http.Get(*url)
		defer response.Body.Close()
		if err != nil {
			// handle error
		}
		ioutil.ReadAll(response.Body)
		fmt.Println("请求地址：", *url, ";    请求次数：", count)
	}
}
