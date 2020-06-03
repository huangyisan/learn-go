package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
)

func main() {
	// 主动连接使用Dial方法
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err !=nil {
		fmt.Println("无法连接...")
		return
	}

	fmt.Println("连接成功....",conn)
	// 客户端发送单行数据,然后退出
	reader := bufio.NewReader(os.Stdin) // os.Stdin标准输入[终端]
	// 从终端读取一行用户输入, 并发送给服务器. 读到'\n'后退出
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败,",err)

	}
	// 在将line发送给服务器
	_ ,err = conn.Write([]byte(line))
	if err != nil {
		fmt.Println("写服务端失败", err)
	}


}