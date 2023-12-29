package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qasperxyz/sdk-go/types"
)

// print to http response internal server error
func ThrowInternalServerError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, types.NewResponseBuilder().AddMessage(types.NewMessage(types.Fatal, err.Error())).Build())
}

// print to http response bad request error
func ThrowBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, types.NewResponseBuilder().AddMessage(types.NewMessage(types.Fatal, err.Error())).Build())
}

// print to http response unauthorized error
func ThrowUnauthorize(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, types.NewResponseBuilder().AddMessage(types.NewMessage(types.Fatal, err.Error())).Build())
}
