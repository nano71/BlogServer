package router

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
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
func ParseJWT(tokenString string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
func randomString(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomString := base64.StdEncoding.EncodeToString(randomBytes)
	return randomString[:length], nil
}
func (r *Router) Preprocessor() gin.HandlerFunc {
	return func(c *gin.Context) {
		//time.Sleep(2000 * time.Millisecond)
		method := c.Request.Method

		handler, ok := r.routes[method+" "+c.Request.URL.Path]

		if ok {
			handler(c)
		} else {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
		}
	}
}
