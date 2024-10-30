package stats

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"

	pbs "github.com/shoshta73/homehub/models/stats"
	"github.com/shoshta73/homehub/models/user"
)

func GetPastebinStats(c echo.Context) error {
	log.Info("Getting pastebin stats")
	token, err := c.Cookie("token")
	if err != nil {
		return err
	}

	usr, err := user.GetUserByToken(token.Value)
	if err != nil {
		return err
	}

	log.Info("Getting pastebin stats", "id", usr.Id)
	stats, err := pbs.GetPastebinStats(usr.Id)
	if err != nil {
		return err
	}

	return c.JSON(200, stats)
}
