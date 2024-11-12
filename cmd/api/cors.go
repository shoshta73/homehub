//go:build !dev
// +build !dev

package main

import (
	"github.com/go-chi/cors"
	"net/http"
	"os"
	"strings"
)

var corsOptions = cors.Options{
	AllowOriginFunc: func(r *http.Request, origin string) bool {
		domain := os.Getenv("DOMAIN")
		allowedDomains := os.Getenv("ALLOWED_DOMAINS")

		if domain == "" {
			logger.Fatal("DOMAIN not set")
		}

		if allowedDomains == "" {
			logger.Fatal("ALLOWED_DOMAINS not set")
		}

		if domain == origin {
			logger.Fatal("Server currently does not support serving frontend")
		}

		for _, d := range strings.Split(allowedDomains, ",") {
			d = strings.TrimSpace(d)

			if d == domain {
				logger.Warn("")
				logger.Fatal("Please do not set you api domain to the allowed domain.\nServer currently does not support serving frontend")
			}

			if origin == d {
				return true
			}
		}
		return false
	},
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"*"},
	ExposedHeaders:   []string{"*"},
	AllowCredentials: true,
	MaxAge:           300,
}
