package main

import (
	"context"
	_ "embed"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"sync"
)

//go:embed del_lock.lua
var luaScript string

func main() {
	rCmd := GetRedisClient()
	rCmd.Eval(context.Background(), luaScript, []string{"lock_key"}, 1)
}

var (
	redisClient     *redis.Client
	redisClientOnce sync.Once
)

func GetRedisClient() *redis.Client {
	redisClientOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: viper.GetString("redis.addr"),
		})
	})
	return redisClient
}
