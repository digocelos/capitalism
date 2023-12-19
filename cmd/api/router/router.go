package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// Initialize routes
func Initialize() *chi.Mux {
  r := chi.NewRouter()
  r.Use(middleware.Logger)

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "It's work")
  })

  return r
}
