package main

// 将get请求的参数作为变量,放到url中.

import (
	"net/http"
	"fmt"
	"net/url"
	"io/ioutil"
)

func httpGet(requestUrl string) (err error) {

	Url, err := url.Parse(requestUrl)
	if err != nil {
		fmt.Printf("url parse faild, err:[%s]", err.Error())
		return
	}

	params := url.Values{}

	// 将需要追加的请求参数使用set方式处理
	params.Set("query","googlesearch")
	params.Set("content","golang")
	Url.RawQuery = params.Encode()

	// 将encode的url转变为string方式,让http.Get()方法进行调用
	requestUrl = Url.String()

	resp ,err := http.Get(requestUrl)

	if err != nil {
		fmt.Printf("request faild [%s]", err.Error())
		return
	}

	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("res status code:[%d]\n", resp.StatusCode)
	fmt.Printf("res body content:[%s]\n", string(bodyContent))
	return

}

func main() {
	var url = "https://httpbin.org/get"
	httpGet(url)
}