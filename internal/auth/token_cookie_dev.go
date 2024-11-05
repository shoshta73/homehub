//go:build dev
// +build dev

package auth

import (
	"net/http"
	"time"
)

func getCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 3),
		SameSite: http.SameSiteNoneMode,
	}
}
