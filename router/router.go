package router

import (
	"github.com/gin-gonic/gin"
)

// Router 自定义路由器
type Router struct {
	// 路由表
	routes map[string]gin.HandlerFunc
}

// POST 添加路由
func (r *Router) POST(path string, handler gin.HandlerFunc) {
	r.add("POST", path, handler)
}

// 添加路由
func (r *Router) add(method string, path string, handler gin.HandlerFunc) {
	r.routes[method+" "+path] = handler
}

// GET 添加路由
func (r *Router) GET(path string, handler gin.HandlerFunc) {
	r.add("GET", path, handler)
}

// Preprocessor 处理请求
func (r *Router) Preprocessor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求方法
		method := c.Request.Method

		// 路由分发
		handler, ok := r.routes[method+" "+c.Request.URL.Path]

		if ok {
			handler(c)
		} else {
			// 不支持的请求方法
			c.AbortWithStatus(405)
		}
	}
}
