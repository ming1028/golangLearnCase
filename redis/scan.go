package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})

	// 获得所有的key
	keys, err := redisClient.Keys(ctx, "*").Result()
	if err != nil {
		fmt.Sprintf("keys err:%v\n", err)
		return
	}
	for _, key := range keys {
		fmt.Println(key)
	}

	// scan
	var cursor uint64
	for {
		var (
			keys []string
			err  error
		)
		keys, cursor, err = redisClient.Scan(ctx, cursor, "*", 0).Result()
		if err != nil {
			fmt.Sprintf("scan err:%v\n", err)
			return
		}
		for _, key := range keys {
			fmt.Println(key)
		}
		fmt.Println("cursor", cursor)
		if cursor == 0 {
			break
		}
	}

	// 简化模式
	iter := redisClient.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys: ", iter.Val())
	}
	if err := iter.Err(); err != nil {
		fmt.Sprintf("iterator next err:%v\n", err)
		return
	}
}
