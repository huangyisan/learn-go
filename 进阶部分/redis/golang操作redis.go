package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
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
	fmt.Println(c)

	// 通过go往redis写入数据, 写入string [k-v]
	_, err = c.Do("set","name","yisan")
	if err != nil {
		fmt.Println("set出错", err)
	} else {
		fmt.Println("set成功")
	}

	// 取出value
	// 因为DO函数返回接口类型和err,r是接口类型, 需要断言.但这里断言得用到redis.类型方法包裹Do函数.
	r, err := redis.String(c.Do("get", "name"))
	if err != nil {
		fmt.Println("获取失败", err)
		return
	}
	fmt.Println("值为",r)



}

