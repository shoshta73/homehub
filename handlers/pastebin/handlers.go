package pastebin

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/charmbracelet/log"
	"github.com/shoshta73/homehub/models/paste"
	"github.com/shoshta73/homehub/models/user"
)

func Create(c echo.Context) error {
	log.Info("Received paste creation request")

	cookie, err := c.Cookie("token")
	if err != nil {
		return err
	}

	token := cookie.Value

	usr, err := user.GetUserByToken(token)
	if err != nil {
		return err
	}

	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	_, err = paste.CreatePaste(body.Title, body.Content, usr.Id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "OK")
}
