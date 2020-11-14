package main

import "net/http"

// w用于应答客户端
// req用于接收客户端请求
func HandConn(w http.ResponseWriter, req *http.Request) {
	str := "hello world"
w.Write([]byte(str))

}


func main() {
	// 注册处理函数
	// 第一个参数为路径
	// 第二个参数为func(ResponseWriter, *Request)， 这里符合HandConn，直接传入
	http.HandleFunc("/", HandConn)

	// 监听绑定
	http.ListenAndServe("127.0.0.1:7777", nil)
}