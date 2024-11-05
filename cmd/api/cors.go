//go:build !dev
// +build !dev

package main

import "github.com/go-chi/cors"

var corsOptions = cors.Options{
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"*"},
	ExposedHeaders:   []string{"*"},
	AllowCredentials: true,
	MaxAge:           300,
}
