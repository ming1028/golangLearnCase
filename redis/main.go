package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	fmt.Println(redisClient)

	// 字符串解析
	opt, err := redis.ParseURL("redis://'':''@localhost:6379")
	if err != nil {
		fmt.Sprintf("redis parse error:%v", err)
	}
	rdb := redis.NewClient(opt)
	fmt.Println(rdb)
}
