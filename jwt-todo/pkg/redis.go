package pkg

import (
	"github.com/go-redis/redis/v7"
	"os"
)

var Client *redis.Client

func init() {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "18.140.56.139:56379"
	}
	Client = redis.NewClient(&redis.Options{
		Addr: dsn,
	})
	_, err := Client.Ping().Result() // ping 测试结果
	if err != nil {
		panic(err)
	}
}
