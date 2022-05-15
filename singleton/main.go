package main

import "sync"

// 单例模式
type singleton struct{}

var (
	instance     *singleton
	instanceOnce sync.Once
)

func getInstance() *singleton {
	instanceOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}
