package types

import "github.com/gin-gonic/gin"

type AbstractController interface {
	Route(r *gin.RouterGroup)
}

type Handler[T any] func(T, *gin.Context)

type NoPayloadHandler func(*gin.Context)
