package main

import (
	"container/list"
	"context"
	"sync"
)

func main() {

}

type Service interface {
	// Name 服务名称
	Name()
	// Execute 服务运行函数
	Execute() error
	// Cancel 取消服务的运行
	Cancel()
}

// GoroutineRestartPolicy Goroutine 重启策略
type GoroutineRestartPolicy int

const (
	// GoroutineRestartPolicyNever 在任何情况下都不会重启
	GoroutineRestartPolicyNever GoroutineRestartPolicy = iota
	// GoroutineRestartPolicyAlways 总是尝试重启
	GoroutineRestartPolicyAlways
	// GoroutineRestartPolicyOnFailure 发生错误自动重启，但是在达到
	// 错误重启次数上限后，自动重启将会终止
	GoroutineRestartPolicyOnFailure
)

// JobRestartPolicy Job 的重启策略
type JobRestartPolicy struct {
	// Goroutine 重启策略
	GoroutineRestartPolicy GoroutineRestartPolicy
	// 最大重启次数
	MaxRetries int32
	// 每次重启延迟时间
	BackoffDelay int32
	// Goroutine 发生Panic的阀值，超过将不重启
	PanicThreshold int32
	// 最大执行时间
	Timeout int32
}

// JobPolicy 控制 Job 的最大并发数
type JobPolicy struct {
	// 可以并发运行的 Job 副本数量上限
	maximumConcurrency int32
	// 当前运行的 Job 副本数量
	currentConcurrency int32
}

type Job struct {
	service          Service
	JobRestartPolicy JobRestartPolicy
	JobPolicy        JobPolicy
}

// GoroutinePool 是一个Goroutine池，用于控制并发数量和任务执行。
type GoroutinePool struct {
	// 等待所有 Goroutine 终止运行
	wg sync.WaitGroup
	// Goroutine 关闭操作
	ctx context.Context
	// 对 Goroutine 进行取消、超时等操作
	cancel context.CancelFunc
	// 等待运行的 Jobs
	jobsWaitQueue list.List
	// 操作 jobsWaitQueue 时需要加锁
	jobMutex sync.Mutex
	// 允许运行的 Goroutine 数量上限
	maximumConcurrency int
	// 当前正在运行的 Goroutine 数量
	currentConcurrency int
}
