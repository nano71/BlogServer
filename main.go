package main

import (
	"blogServer/middleware"
	"blogServer/router"
	"crypto/tls"
	"flag"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

type CertReloader struct 
type CertReloader struct {
	mu       sync.RWMutex
	cert     *tls.Certificate
	certFile string
	keyFile  string
}

func NewCertReloader(certFile, keyFile string) (*CertReloader, error) {
	cert, err := loadCertificate(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	cr := &CertReloader{
		cert:     cert,
		certFile: certFile,
		keyFile:  keyFile,
	}
	go cr.watch()
	return cr, nil
}

func (cr *CertReloader) watch() {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	for range ticker.C {
		cert, err := loadCertificate(cr.certFile, cr.keyFile)
		if err != nil {
			slog.Error("SSL reload failed", "error", err)
			continue
		}
		cr.mu.Lock()
		cr.cert = cert
		cr.mu.Unlock()
		slog.Info("SSL certificate reloaded successfully.")
	}
}

func (cr *CertReloader) GetCertificate(clientHello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	return cr.cert, nil
}

func loadCertificate(certFile, keyFile string) (*tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func main() {
	time.Local = time.FixedZone("CST", 8*3600)

	var protocol string
	var isReleaseMode bool
	flag.StringVar(&protocol, "protocol", "http", "目标协议,默认为http")
	flag.BoolVar(&isReleaseMode, "releaseMode", false, "release模式,默认为false")
	flag.Parse()

	slog.Info("log", "目标协议为:", protocol)
	slog.Info("log", "gin运行为release模式:", isReleaseMode)

	if isReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	ginServer := gin.Default()
	ginServer.StaticFile("/robots.txt", "./robots.txt")
	ginServer.Static("/uploads", "./uploads")
	ginServer.Use(middleware.Cors())
	ginServer.Use(middleware.AccessLog())
	ginServer.Use(middleware.Timeout(10000))
	ginServer.Use(middleware.Interceptor())
	ginServer.Use(middleware.IPBan())
	ginServer.Use(router.Default())

	slog.Info("log", "blogServer is running on port 9000.")

	var runError error
	if protocol == "https" {
		certFile := "/etc/nginx/ssl/nano71.com.crt"
		keyFile := "/etc/nginx/ssl/nano71.com.key"

		certReloader, err := NewCertReloader(certFile, keyFile)
		if err != nil {
			slog.Error("Failed to load initial certificate", err)
			return
		}

		tlsConfig := &tls.Config{
			GetCertificate: certReloader.GetCertificate,
			MinVersion:     tls.VersionTLS12,
		}

		server := &http.Server{
			Addr:      ":9000",
			Handler:   ginServer,
			TLSConfig: tlsConfig,
		}

		runError = server.ListenAndServeTLS("", "")
	} else {
		runError = ginServer.Run(":9000")
	}

	if runError != nil {
		slog.Error("Failed to start server", runError)
	}
}
