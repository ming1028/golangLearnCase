package main

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// 将多个命令放入 pipeline 中，然后使用1次读写操作像执行单个命令一样执行它们，节省了执行命令的网络往返时间
	pipe := redisClient.Pipeline()
	incr := pipe.Incr(ctx, "pipe_counter")
	pipe.Expire(ctx, "pipe_counter", time.Hour)

	cmds, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Sprintf("pipe exec err:%v\n", err)
		return
	}
	fmt.Println(cmds, incr.Val())

	// 闭包函数模式
	var incrs *redis.IntCmd
	cmdss, err := redisClient.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		incrs = pipeliner.Incr(ctx, "pipe_counter")
		pipeliner.Expire(ctx, "pipe_counter", time.Hour)
		return nil
	})
	if err != nil {
		fmt.Sprintf("pipelined err:%v\n", err)
	}
	for _, cmd := range cmdss {
		fmt.Println(cmd.String())
	}
	fmt.Println(incrs.Val())
}
