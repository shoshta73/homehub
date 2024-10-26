package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/shoshta73/homehub/auth"
	"github.com/shoshta73/homehub/log"
)

func main() {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	e.POST("/auth/register", auth.Register)
	e.POST("/auth/login", auth.Login)

	log.Info("Starting server")
	err := e.Start(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
