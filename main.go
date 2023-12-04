package main

import (
	"blogServer/api"
	"blogServer/router"
	"flag"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	time.Local = time.FixedZone("CST", 8*3600)

	var protocol string
	var isReleaseMode bool
	flag.StringVar(&protocol, "protocol", "http", "目标协议,默认为http")
	flag.BoolVar(&isReleaseMode, "releaseMode", false, "release模式,默认为false")
	flag.Parse()
	slog.Info("log", "目标协议为:", protocol)
	slog.Info("log", "gin运行为release模式:", isReleaseMode)
	var runError error
	slog.Info("log", "blogServer is running on port 9000.")
	if isReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	ginServer := gin.Default()
	ginServer.Static("/uploads", "./uploads")
	ginServer.Use(api.Cors())
	ginServer.Use(api.TimeoutMiddleware(10000))
	ginServer.Use(api.InterceptorMiddleware())
	ginServer.Use(api.IPBanMiddleware())
	ginServer.Use(router.Default())

	if protocol == "https" {
		server := &http.Server{
			Addr:    ":9000",
			Handler: ginServer,
		}

		runError = server.ListenAndServeTLS("nano71.com_bundle.crt", "nano71.com.key")

	} else {
		//_ = ginServer.SetTrustedProxies([]string{"127.0.0.1"})
		runError = ginServer.Run(":9000")
	}
	if runError != nil {
		slog.Error("Failed to start HTTPS server", runError)
	}
}
