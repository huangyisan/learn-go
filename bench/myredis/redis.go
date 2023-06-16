package myredis

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx context.Context

func init() {
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
func GetKey(c *gin.Context) {
	key := c.DefaultQuery("key", "defaultValue")
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		c.JSON(200, gin.H{
			"message": val,
		})
	} else if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"message": val,
	})
}

func SetKey(c *gin.Context) {
	key := c.DefaultQuery("key", "defaultValue")
	value := c.DefaultQuery("value", "defaultValue")
	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"message": "success",
	})
}
