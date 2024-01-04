package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/qasperxyz/sdk-go/errors"
	"github.com/qasperxyz/sdk-go/types"
)

func Dispatch[T any](payload T, callback types.Handler[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&payload); err != nil {
			errors.ThrowBadRequest(c, err)
			return
		}
		callback(payload, c)
	}
}

func NoPayloadDispatch(callback types.NoPayloadHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		callback(c)
	}
}
