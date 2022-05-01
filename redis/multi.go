package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	// 事务
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pipe := redisClient.TxPipeline()
	incr := pipe.Incr(ctx, "pipe_counter")
	pipe.Expire(ctx, "pipe_counter", time.Hour)
	cmds, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Sprintf("TxPipeline err:%v\n", err)
		return
	}
	fmt.Println(cmds, incr.Val())
}
