package middleware

import (
	"blogServer/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

var bannedIPs []string

func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if c.Request.Method == http.MethodGet {
			bannedIPs = append(bannedIPs, clientIP)
			response.Forbidden(c)
			c.Abort()
			return
		}

		c.Next()
	}
}

func IPBan() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		for _, bannedIP := range bannedIPs {
			if clientIP == bannedIP {
				response.Forbidden(c)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
