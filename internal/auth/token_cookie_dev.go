//go:build dev
// +build dev

package auth

import (
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 3),
	})
}
