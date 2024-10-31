//go:build !dev
// +build !dev

package auth

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/charmbracelet/log"
	"github.com/shoshta73/homehub/models/stats"
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

	usr, err := user.CreateUser(body.Username, body.Name, body.Email, body.Password)
	if err != nil {
		return err
	}

	go user.GenerateIdenticon(user.User{Username: body.Username, Avatar: usr.Avatar})

	log.Info("User registered")

	tkn, err := usr.GetClaims().GenerateToken()
	if err != nil {
		return err
	}

	go stats.InitUserStats(usr.Id)

	cookie := http.Cookie{
		Name:     "token",
		Value:    tkn,
		Expires:  time.Now().Add(time.Hour * 24 * 3),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	c.SetCookie(&cookie)

	return c.String(http.StatusOK, "OK")
}

func Login(c echo.Context) error {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	log.Info("Received login request")

	if !user.VerifyUser(body.Username, body.Email, body.Password) {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}

	usr, err := user.GetUserByEmail(body.Email)
	if err != nil {
		return err
	}

	tkn, err := usr.GetClaims().GenerateToken()
	if err != nil {
		return err
	}

	go stats.CheckUserStats(usr.Id)

	cookie := http.Cookie{
		Name:     "token",
		Value:    tkn,
		Expires:  time.Now().Add(time.Hour * 24 * 3),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	c.SetCookie(&cookie)

	return c.String(http.StatusOK, "OK")
}

func LoginWithEmail(c echo.Context) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	log.Info("Received login request")

	if body.Email == "" && body.Password == "" {
		return c.String(http.StatusBadRequest, "Email and password are required")
	}

	if body.Email == "" {
		return c.String(http.StatusBadRequest, "Email is required")
	}

	if body.Password == "" {
		return c.String(http.StatusBadRequest, "Password is required")
	}

	if !user.VerifyUserByEmail(body.Email, body.Password) {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}

	usr, err := user.GetUserByEmail(body.Email)
	if err != nil {
		return err
	}

	tkn, err := usr.GetClaims().GenerateToken()
	if err != nil {
		return err
	}

	go stats.CheckUserStats(usr.Id)

	cookie := http.Cookie{
		Name:     "token",
		Value:    tkn,
		Expires:  time.Now().Add(time.Hour * 24 * 3),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	c.SetCookie(&cookie)

	return c.String(http.StatusOK, "OK")
}

func LoginWithUsername(c echo.Context) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		return err
	}

	log.Info("Received login request")

	if body.Username == "" && body.Password == "" {
		return c.String(http.StatusBadRequest, "Username and password are required")
	}

	if body.Username == "" {
		return c.String(http.StatusBadRequest, "Username is required")
	}

	if body.Password == "" {
		return c.String(http.StatusBadRequest, "Password is required")
	}

	if !user.VerifyUserByUsername(body.Username, body.Password) {
		return c.String(http.StatusUnauthorized, "Invalid credentials")
	}

	usr, err := user.GetUserByUsername(body.Username)
	if err != nil {
		return err
	}

	tkn, err := usr.GetClaims().GenerateToken()
	if err != nil {
		return err
	}

	go stats.CheckUserStats(usr.Id)

	cookie := http.Cookie{
		Name:     "token",
		Value:    tkn,
		Expires:  time.Now().Add(time.Hour * 24 * 3),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	c.SetCookie(&cookie)

	return c.String(http.StatusOK, "OK")
}
