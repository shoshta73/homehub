package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/shoshta73/homehub/internal/auth"
	"github.com/shoshta73/homehub/internal/pastebin"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(corsOptions))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Mount("/auth", auth.Routes())
	r.Mount("/pastebin", pastebin.Routes())

	http.ListenAndServe(":3000", r)
	http.ListenAndServeTLS(":3001", "cert.pem", "key.pem", r)
}
