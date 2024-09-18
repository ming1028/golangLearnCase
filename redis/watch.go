package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const routineCount = 100

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			// 获取当前值或者零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			n++
			_, err = tx.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
				// pipe 处理错误情况
				pipeliner.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		for retries := routineCount; retries > 0; retries-- {
			err := redisClient.Watch(ctx, txf, key)
			fmt.Println(err)
			if err != redis.TxFailedErr {
				return err
			}
			return err
		}
		return errors.New("increment reached maximum number of retries")
	}

	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()

			if err := increment("counter3"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := redisClient.Get(ctx, "counter3").Int()
	fmt.Println("ended with", n, err)
}
