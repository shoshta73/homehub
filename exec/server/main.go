//go:build !dev
// +build !dev

package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/charmbracelet/log"
)

var liveCertDir string

func init() {
	liveCertDir = os.Getenv("LIVE_CERT")
}

func main() {
	e := echo.New()

	e.Pre(middleware.HTTPSRedirect())
	e.Pre(middleware.RemoveTrailingSlash())

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper:    nil,
		Root:       "dist",
		Index:      "index.html",
		HTML5:      true,
		Browse:     false,
		IgnoreBase: false,
		Filesystem: nil,
	}))

	certFile := filepath.Join(liveCertDir, "fullchain.pem")
	keyFile := filepath.Join(liveCertDir, "privkey.pem")

	routes(e)

	log.Info("Starting server")
	go func() {
		if err := e.Start(":80"); err != nil && err != http.ErrServerClosed {
			log.Fatal("HTTP server failed:", err)
		}
	}()

	err := e.StartTLS(":443", certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
}
