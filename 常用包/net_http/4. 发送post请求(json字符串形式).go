package main

import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
)

func httpPost(requestUrl string) (err error) {
	// 这种方式需要自己创建client,然后追加json内容去请求.
	client := &http.Client{}
	data := make(map[string]interface{})
	data["name"] = "yisan"
	data["pwd"] = "huang"
	jsonData, _ := json.Marshal(data)

	requestPost, err := http.NewRequest("POST", requestUrl, bytes.NewReader(jsonData))

	resp, err := client.Do(requestPost)

	if err != nil {
		fmt.Println("请求异常")
		return
	}

	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("status code is [%d]\n", resp.StatusCode)
	fmt.Printf("content info is [%s]\n", string(bodyContent))

	return
}


func main() {
	url := "https://httpbin.org/post"
	httpPost(url)
}
