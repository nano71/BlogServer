package middleware

import (
	"blogServer/response"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func Timeout(limit time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(limit*time.Millisecond),
		timeout.WithResponse(response.Timeout),
	)
}
