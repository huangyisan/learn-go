package main

import (
	"net"
	"fmt"
)

func main() {
	// 建立tcp监听
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("监听异常")
		return
	}

	defer listener.Close()

	// 建立连接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("建联异常")
		return
	}

	defer conn.Close()

	// 准备从conn中读取数据
	buf := make([]byte, 1024)
	// 返回两个参数，第一个为接收到的字符串长度，第二个为错误信息
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取数据异常")
		return
	}

	fmt.Printf("读取的内容为: %s", string(buf[:n]))

}