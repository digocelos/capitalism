package main

import (
	"log"
	"net/http"

	"github.com/digocelos/capitalism/config"
	"github.com/digocelos/capitalism/router"
)

func main() {

  config := config.Initialize()

  r := router.Initialize()
  log.Println("Listening..." + config.Port)
  log.Fatal(http.ListenAndServe(":" + config.Port, r))
  
}
