package middleware

import (
	"server/config"
	"server/core/response"

	"github.com/gin-gonic/gin"
)

// ApiKey验证中间件
func KeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ApiKey验证
		apiKey := c.Request.Header.Get("apiKey")
		if apiKey != config.Config.ApiKey {
			response.Fail(c, response.KeyInvalid)
			c.Abort()
			return
		}
	}
}
