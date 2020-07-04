package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func httpGet(requestUrl string) (err error)  {
	client := &http.Client{}
	requestGet,_ := http.NewRequest("GET", requestUrl, nil)
	requestGet.Header.Add("cus-header","cus-value")
	requestGet.Header.Add("cus-header1","cus-value1")

	resp, err := client.Do(requestGet)

	if err != nil {
		fmt.Println("请求异常")
		return
	}

	defer resp.Body.Close()

	bodyContent,err := ioutil.ReadAll(resp.Body)
	fmt.Printf("status code is [%d]", resp.StatusCode)
	fmt.Printf("content info is [%s]", bodyContent)
	return

}

func main() {
	var url = "https://httpbin.org/get"
	httpGet(url)
}