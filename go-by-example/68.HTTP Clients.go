package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    resp, err := http.Get("https://www.baidu.com")
    if err != nil {
        panic(err)
    }
    // defer最后关闭resp.Body内容
    defer resp.Body.Close()

    // 获取status
    fmt.Println(resp.StatusCode)

    // 读取body全部内容
    //copyBody := resp.Body
    content, _ := ioutil.ReadAll(resp.Body)
    // 读取到的为二进制，通过string()转成字符串
    fmt.Println(string(content))

    //// 前五行内容
    //fmt.Printf("read first 5 line\n")
    //scanner := bufio.NewScanner(resp.Body)
    //for i := 0; scanner.Scan() && i < 5 ; i++ {
    //    fmt.Println(scanner.Text())
    //}
}