package middleware

import (
	"blogServer/database"
	"github.com/gin-gonic/gin"
	"time"
)

type Log struct {
	Id         int
	Ip         string
	CreateTime time.Time
	Url        string
	Ua         string
	Latency    string
	Status     int
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start).String()
		go func() {
			log := &Log{
				Ip:         c.ClientIP(),
				CreateTime: start,
				Url:        c.Request.RequestURI,
				Ua:         c.Request.UserAgent(),
				Latency:    latency,
				Status:     c.Writer.Status(),
			}
			database.GetDB().Create(log)
		}()
	}
}
