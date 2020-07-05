package main

import (
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
)

func httpPost(url string) (err error) {
	data := make(map[string]interface{})
	data["name"] = "yisan"
	data["pwd"] = "huang"

	jsonData, _ := json.Marshal(data)


	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))

	if err != nil {
		fmt.Println("post请求异常")
		return
	}

	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("http code is %d\n", resp.StatusCode)
	fmt.Printf("http info is %s\n", string(bodyContent))
	return
}

func main() {
	url := "https://httpbin.org/post"
	httpPost(url)
}