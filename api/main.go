package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

  r := chi.NewRouter()
  r.Use(middleware.Logger)

  r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("It's work"))
  })

  fmt.Println("Start server using 8080 port")
  http.ListenAndServe(":8080", r)
  
}
