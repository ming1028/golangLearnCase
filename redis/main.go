package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
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
		fmt.Sprintf("redis parse error:%v\n", err)
		return
	}
	rdb := redis.NewClient(opt)
	fmt.Println(rdb)

	// go-redis使用
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	// 执行命令获得结果
	val := redisClient.Get(ctx, "name")
	if err != nil {
		fmt.Sprintf("redis get err key:name err:%v\n", val)
		return
	}
	fmt.Println(val)

	// 获得命令对象
	cmder := redisClient.Get(ctx, "name")
	fmt.Println(cmder.Val(), cmder.Err())

	// 直接执行命令获取错误
	err = redisClient.Set(ctx, "name", "张三", time.Minute*10).Err()
	if err != nil {
		fmt.Sprintf("redis set err:%v\n", err)
		return
	}
	val = redisClient.Get(ctx, "name")
	fmt.Println("redis get val:", val.Val(), val.Err())

	// 执行自定义命令
	vals, err := redisClient.Do(ctx, "set", "name", "王五", "EX", 3600).Result()
	if err != nil {
		fmt.Sprintf("redis do set err:%v\n", err)
		return
	}
	fmt.Println(vals.(string))

	vals, err = redisClient.Do(ctx, "get", "name").Result()
	if err != nil {
		fmt.Sprintf("redis do get err:%v\n", err)
		return
	}
	fmt.Println(vals)

	nameVal, err := redisClient.Get(ctx, "names").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println("defaultVal", err)
			return
		}
		fmt.Sprintf("redis get err:%v\n", err)
		return
	}
	fmt.Println(nameVal)
}
