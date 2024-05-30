package api

import (
	"ryosantouchh/gym-management-backend/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type ApiContext struct {
	// embedded field for gin context
	*gin.Context
}

func NewApiContext(ctx *gin.Context) *ApiContext {
	return &ApiContext{Context: ctx}
}

func (c *ApiContext) JSON(statusCode int, response interface{}) {
	c.Context.JSON(statusCode, response)
}

func GinHandler(handler func(ports.HTTPContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler(NewApiContext(ctx))
	}
}
