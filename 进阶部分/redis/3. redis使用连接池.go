package main

import (
	"github.com/gomodule/redigo/redis"
	"fmt"
)


// 定义一个全局的pool

var pool *redis.Pool

// 在init方法里面初始化连接池

func init() {
	pool = &redis.Pool{
		MaxIdle: 10, // 最大空闲连接数
		MaxActive: 0, // 最大连接数, 如果为0, 则不设限制
		IdleTimeout: 100, // 最大空闲时间, 单位s, 必须设置比server端来的短
		Dial: func() (redis.Conn, error) { // 初始化连接代码, 连接哪个redis
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	// 最后关闭连接池
	defer pool.Close()

	// 先从pool中取出连接
	conn := pool.Get()
	// 关闭取出的连接
	defer conn.Close()

	_, err := conn.Do("set", "name", "yisan")
	if err != nil {
		fmt.Println("set 失败",err)
		return
	}

	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get 失败", err)
	}
	fmt.Println(r)


}