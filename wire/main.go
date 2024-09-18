package main

import "github.com/gin-gonic/gin"

func main() {
	ctx := &gin.Context{}
	e := InitializeEvent(ctx)
	e.Start()
}
