package main

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
)

// form形式的内容进行post提交

func httpPost(requestUrl string) (err error) {
	data := url.Values{}
	data.Add("username","yisan")
	data.Add("password", "123")

	resp, err := http.PostForm(requestUrl, data)

	if err != nil {
		fmt.Printf("请求异常,%s", err.Error())
		return
	}

	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("status code is [%d]\n", resp.StatusCode)
	fmt.Printf("info is [%s]", bodyContent)

	return
}

func main() {
	var url = "https://httpbin.org/post"
	httpPost(url)
}