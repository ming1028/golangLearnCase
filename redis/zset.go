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
	fmt.Println(redisClient)

	zsetKey := "zsetRank"
	rank := []*redis.Z{
		{
			Score:  80,
			Member: "chinese",
		},
		{
			Score:  90,
			Member: "math",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// zAdd
	err := redisClient.ZAdd(ctx, zsetKey, rank...).Err()
	if err != nil {
		fmt.Sprintf("zadd err:%v\n", err)
		return
	}

	// zIncrBy
	newCore, err := redisClient.ZIncrBy(ctx, zsetKey, 10, "math").Result()
	if err != nil {
		fmt.Sprintf("zincrBy err:%v\n", err)
		return
	}
	fmt.Println("newCore is:", newCore)

	// 取分数最高的
	fmt.Println("倒序：")
	ret := redisClient.ZRevRangeWithScores(ctx, zsetKey, 0, 0).Val()
	for _, val := range ret {
		fmt.Println(val.Score, val.Member)
	}

	fmt.Println("正序排序")
	ret, err = redisClient.ZRangeWithScores(ctx, zsetKey, 0, 1).Result()
	if err != nil {
		fmt.Sprintf("ZRangeWithScores err:%v\n", err)
		return
	}
	for _, val := range ret {
		fmt.Println(val.Score, val.Member)
	}

	// 取分数 85-95
	op := &redis.ZRangeBy{
		Min: "85",
		Max: "95",
	}
	ret, err = redisClient.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
	if err != nil {
		fmt.Sprintf("ZRangeByScoreWithScores err:%v\n", err)
		return
	}
	fmt.Println("范围：")
	for _, val := range ret {
		fmt.Println(val.Score, val.Member)
	}
}
