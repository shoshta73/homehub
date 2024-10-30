package main

import (
	"github.com/labstack/echo/v4"

	"github.com/shoshta73/homehub/auth"
	"github.com/shoshta73/homehub/models/user"
	"github.com/shoshta73/homehub/pastebin"
	"github.com/shoshta73/homehub/stats"
)

func routes(e *echo.Echo) {
	e.POST("/auth/register", auth.Register)
	e.POST("/auth/login", auth.Login)

	e.GET("/avatar", user.AvatarUrl)
	e.GET("/avatars/:username", user.Avatar)

	e.POST("/pastebin/create", pastebin.Create)
	e.GET("/pastebin/stats", stats.GetPastebinStats)
}
