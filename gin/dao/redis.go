package dao

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var redisClient = &redis.Client{}

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               "10.0.1.71:6379", // Redis 服务器地址和端口号
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           "", // Redis 密码，如果没有设置则为空
		DB:                 6,  // Redis 数据库编号，默认为 0
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolFIFO:           false,
		PoolSize:           10, // Redis Pool Size
		MinIdleConns:       2,  // 最小空闲连接数，默认为 0
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})
}

func WriteToRedis(ctx *gin.Context) {
	cmd := redisClient.Set(context.TODO(), "name", "jerry", -1)

	if cmd.Err() != nil {
		panic(cmd.Err())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Redis 写入成功",
	})
}

func GetFromRedis(ctx *gin.Context) {
	panic("not implemented")
	cmd := redisClient.Get(context.TODO(), "name")

	if cmd.Err() != nil {
		panic(cmd.Err())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Redis 查询成功，数据:%s", cmd.String()),
	})
}
