package main

import (
	"github.com/labstack/echo/v4"

	"github.com/shoshta73/homehub/handlers/auth"
	"github.com/shoshta73/homehub/handlers/pastebin"
	"github.com/shoshta73/homehub/handlers/stats"

	"github.com/shoshta73/homehub/models/user"
)

func routes(e *echo.Echo) {
	e.POST("/auth/register", auth.Register)
	e.POST("/auth/login/username", auth.LoginWithUsername)
	e.POST("/auth/login/email", auth.LoginWithEmail)
	e.POST("/auth/validate", auth.Validate)

	e.GET("/avatar", user.AvatarUrl)
	e.GET("/avatars/:username", user.Avatar)

	e.POST("/pastebin/create", pastebin.Create)
	e.GET("/pastebin/stats", stats.GetPastebinStats)
}
