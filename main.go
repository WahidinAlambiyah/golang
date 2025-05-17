package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"user-management/handler"
)

func main() {
	r := chi.NewRouter()

	// Middleware bawaan
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Route dasar
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Post("/register", handler.Register)
	r.Post("/login", handler.Login)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
