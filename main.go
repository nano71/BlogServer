package main

import (
	"blogServer/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	// 创建 Gin 路由
	ginServer := gin.Default()
	ginServer.Use(router.Middleware())

	//db := database.GetDB()

	// 启动服务器
	_ = ginServer.Run()
	//defer db.Close()
}
