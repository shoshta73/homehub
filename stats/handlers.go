package stats

import (
	"github.com/labstack/echo/v4"
	"github.com/shoshta73/homehub/models/stats"
	"github.com/shoshta73/homehub/models/user"
)

func GetPastebinStats(c echo.Context) error {
	token, err := c.Cookie("token")
	if err != nil {
		return err
	}

	usr, err := user.GetUserByToken(token.Value)
	if err != nil {
		return err
	}

	stats, err := stats.GetPastebinStats(usr.Id)
	if err != nil {
		return err
	}

	return c.JSON(200, stats)
}
