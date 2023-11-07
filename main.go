package main

import (
	"blogServer/api"
	"blogServer/router"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	time.Local = time.FixedZone("CST", 8*3600)
	//gin.SetMode(gin.ReleaseMode)
	ginServer := gin.Default()
	ginServer.Static("/uploads", "./uploads")
	ginServer.Use(api.Cors())
	ginServer.Use(api.TimeoutMiddleware(10000))
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
	slog.Info("blogServer is running on port 9000.")
}
