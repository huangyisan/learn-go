package main

import (
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("Get error")
		return
	}

	// 最后关闭body
	defer resp.Body.Close()

	// 创建4k缓存
	buf := make([]byte, 4*1024)
	var tmp string

	// 不停读取，直到返回为0，表示没有数据了
	for  {
		// 往buff读
		n, err := resp.Body.Read(buf)
		// 此时读取完毕, 则break
		if n == 0 {
			break
		}


		tmp += string(buf[:n])
	}

	fmt.Printf("content is %s \n", tmp)


}