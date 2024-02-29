package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

// RequestLogger 中间件记录请求时间
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 执行下一个处理程序
		c.Next()

		// 请求结束后计算时间
		latency := time.Since(start)
		path := c.Request.URL.Path
		method := c.Request.Method

		// 记录请求时间
		c.Writer.Header().Set("X-Request-Time", latency.String())
		c.Writer.Header().Set("X-Request-Path", path)
		c.Writer.Header().Set("X-Request-Method", method)
	}
}
