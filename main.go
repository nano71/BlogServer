package main

import (
	"blogServer/api"
	"blogServer/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	ginServer := gin.Default()
	ginServer.Use(cors.Default())
	ginServer.Use(api.TimeoutMiddleware(10000))
	ginServer.Static("/uploads", "./uploads")
	ginServer.Use(router.Default())

	//_ = ginServer.Run(":9000")
	server := &http.Server{
		Addr:    ":9000",
		Handler: ginServer,
	}

	err := server.ListenAndServeTLS("nano71.com_bundle.crt", "nano71.com.key")
	if err != nil {
		slog.Error("Failed to start HTTPS server", err)
	}
	slog.Info("blogServer started.")
}
