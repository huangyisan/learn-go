package main

import (
	"net"
	"fmt"
)

func main() {
	// 和服务端建联
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("建联失败")
	}

	defer conn.Close()

	//写数据
	//str := []byte{'h','e','l','l','o'}
	str := "hello world"

	// 将str转换为byte类型的切片
	conn.Write([]byte(str))

}
