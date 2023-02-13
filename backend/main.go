package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	// r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/health", health)
		r.Post("/login", login)
		r.Get("/register", register)
	})

	// Private routes
	r.Group(func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/{userId}", getUserById)
		})
		r.Route("/project", func(r chi.Router) {
			r.Get("/{projectId}", getProjectById)
			r.Post("/", postProject)
			r.Get("/", getProjects)
		})
	})

	log.Fatal(http.ListenAndServe(":4044", r))
}
