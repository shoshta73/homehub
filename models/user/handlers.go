package user

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/shoshta73/homehub/log"
)

func AvatarUrl(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return err
	}

	token := cookie.Value

	usr, err := getUserByToken(token)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, usr.GetAvatarURL())
}

func Avatar(c echo.Context) error {
	username := strings.TrimSuffix(c.Param("username"), ".png")
	log.Info("Getting avatar for", "username", username)

	usr, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	return c.File(usr.GetAvatar())
}
