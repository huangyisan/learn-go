package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func httpGet(url string) (err error){
	req, err := http.Get(url)
	fmt.Println(1)
	if err != nil {
		fmt.Println("get请求失败， %s", err.Error())
		return
	}
	defer req.Body.Close()

	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("读取content失败")
		return
	}
	fmt.Println(string(content))
	return

}

func main() {
	url := "https://httpbin.com/get"
	httpGet(url)
}