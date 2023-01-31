package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go handle(ctx, 3*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("main:", ctx.Err())
	}
}

func handle(
	ctx context.Context,
	duration time.Duration,
) {
	select {
	case <-ctx.Done():
		fmt.Println("handle:", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
