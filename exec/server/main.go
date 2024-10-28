package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/shoshta73/homehub/auth"
	"github.com/shoshta73/homehub/log"
	"github.com/shoshta73/homehub/models/user"
	"github.com/shoshta73/homehub/pastebin"
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

	e.POST("/auth/register", auth.Register)
	e.POST("/auth/login", auth.Login)

	e.GET("/avatar", user.AvatarUrl)
	e.GET("/avatars/:username", user.Avatar)

	e.POST("/pastebin/create", pastebin.Create)

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
