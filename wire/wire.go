//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeEvent(ctx *gin.Context) Event {
	wire.Build(NewGreeter, NewEvent, NewMessage)
	return Event{}
}
