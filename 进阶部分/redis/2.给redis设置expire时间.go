package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func main() {
	// go 发起和redis连接
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	// 关闭redis连接
	defer c.Close()

	// 给name设置有效时间为10s, 只能对已经存在的key进行设置.
	_, err = c.Do("expire", "name", 10)
	if err != nil {
		fmt.Println("设置expire失败")
		return
	}
}