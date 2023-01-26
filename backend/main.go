package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Backend is healthy."))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/health", health)

	log.Fatal(http.ListenAndServe(":4044", r))
}
