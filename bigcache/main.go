package main

import (
	"context"
	"fmt"
	"github.com/allegro/bigcache/v3"
)

func main() {
	cacheConfig := bigcache.Config{
		Shards: 1024,
	}
	cache, err := bigcache.New(context.Background(), cacheConfig)
	if err != nil {
		panic(err)
	}
	err = cache.Set("key1", []byte("value1"))
	if err != nil {
		panic(err)
	}

	entry, err := cache.Get("key1")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(entry))
}
