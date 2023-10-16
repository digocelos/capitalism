package main

import (
	"log"
	"net/http"

	"github.com/digocelos/capitalism/router"
)

func main() {

  r := router.Initialize()
  log.Println("Listening...")
  log.Fatal(http.ListenAndServe(":8080", r))
  
}
