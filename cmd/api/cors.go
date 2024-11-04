//go:build !dev
// +build !dev

package main

import "github.com/go-chi/cors"

var corsOptions = cors.Options{
	AllowCredentials: true,
	MaxAge:           300,
}
