package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/shoshta73/homehub/log"
	"github.com/shoshta73/homehub/models/user"
)

func Register(c echo.Context) error {
	var body struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	log.Info("Received registration request")

	ue, err := user.IsExistingByEmail(body.Email)
	if err != nil {
		return err
	}

	if ue {
		return c.String(http.StatusConflict, "User with this email already exists")
	}

	ue, err = user.IsExistingByUsername(body.Username)
	if err != nil {
		return err
	}

	if ue {
		return c.String(http.StatusConflict, "User with this username already exists")
	}

	_, err = user.CreateUser(body.Username, body.Name, body.Email, body.Password)
	if err != nil {
		return err
	}

	log.Info("User registered")

	return c.String(http.StatusOK, "Hello, world!")
}
