package middleware

import (
	"blogServer/response"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"time"
)

func Timeout(limit time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(limit*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(response.Timeout),
	)
}
