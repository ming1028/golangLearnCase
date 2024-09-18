package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Service struct {
	wg sync.WaitGroup
}

func (s *Service) FakeSendEmail() {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered panic: %v\n", err)
			}
		}()

		log.Println("Goroutine enter")
		time.Sleep(time.Second * 5)
		log.Println("Goroutine exit")
	}()
}

func (s *Service) GracefulStop(ctx context.Context) {
	log.Println("Waiting for service to finish")
	quit := make(chan struct{})
	go func() {
		s.wg.Wait() // 阻塞等待协程处理完
		close(quit) // 关闭
	}()
	select {
	case <-ctx.Done():
		log.Println("context was marked as done earlier, than user service has stopped")
	case <-quit:
		log.Println("Service finished")
	}
}

func (s *Service) Handler(w http.ResponseWriter, r *http.Request) {
	duration, err := time.ParseDuration(r.FormValue("duration"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	time.Sleep(duration)

	// 模拟需要异步执行的代码，比如注册接口异步发送邮件、发送 Kafka 消息等
	s.FakeSendEmail()

	_, _ = w.Write([]byte("Welcome HTTP Server"))
}

func main() {
	srv := &http.Server{
		Addr: ":8000",
	}

	svc := &Service{}
	http.HandleFunc("/sleep", svc.Handler)

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			// Error starting or closing listener:
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
		log.Println("Stopped serving new connections")
	}()

	// 错误写法
	// srv.RegisterOnShutdown(func() {
	//  svc.GracefulStop(ctx)
	// })

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// We received an SIGINT/SIGTERM/SIGQUIT signal, shut down.
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}

	// 优雅退出 service
	svc.GracefulStop(ctx)
	log.Println("HTTP server graceful shutdown completed")
}
