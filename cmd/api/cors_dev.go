//go:build dev
// +build dev

package main

import "github.com/go-chi/cors"

var corsOptions = cors.Options{
	AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:4173"},
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"*"},
	ExposedHeaders:   []string{"*"},
	AllowCredentials: true,
	MaxAge:           300,
}
