package middleware

import (
	"blogServer/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

var bannedIPs []string
var seoPaths = []string{
	"/sitemap.xml",
	"/robots.txt",
	"/favicon.ico",
}

func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		path := c.Request.URL.Path

		if c.Request.Method == http.MethodGet {
			for _, seoPath := range seoPaths {
				if path == seoPath {
					c.Next()
					return
				}
			}

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
