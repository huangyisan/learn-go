package main

import (
	"fmt"
	"net"
	_ "io"
)

func process(conn net.Conn) {
	// 这里循环的接受客户端发送的数据


	// 处理完就关闭连接
	defer conn.Close()

	// 读取传过来的信息
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		fmt.Println("开始读取..")
		//n表示多少个字节, 这边Read函数会把conn的内容写入到buf中
		n, err := conn.Read(buf) // 等待客户端通过conn 发送信息,如果客户端没有发送(Write),则该协程一直阻塞在这里
		if err != nil {
			fmt.Println("读内容出异常", err)
			// 如果异常就退出
			return
		}
		// 将客户端发送的内容显示在服务器终端
		fmt.Print(string(buf[:n]))
	}

}

func main() {
	// 服务器监听
	fmt.Println("服务器开始监听...")
	// tcp表示网络协议使用tcp
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer listen.Close() //延时关闭连接

	// 循环等待客户端来连接
	for {
		// 等待客户端连接
		fmt.Println("等待客户端连接...")
		// Accept在没接受客户端的时候会阻塞
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept 出错", err)
			// 这边不需要return. 连接错了 就等待下一次.
		} else {
			fmt.Println("连接成功")
			fmt.Println("客户端ip为", conn.RemoteAddr().String())
		}
		// 创建一个协程, 将conn递交出去, 为客户端服务
		go process(conn)
	}

	fmt.Println("")
}
