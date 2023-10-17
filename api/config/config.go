package config

import (
	"log"
	"os"
)

type Config struct {
  Port string
}

// Initialize configurations
func Initialize() Config {
  log.Println("Initialize configurations")

  var config Config

  config.Port = os.Getenv("PORT")

  log.Println("Read port from env: " + config.Port)

  if (config.Port == "") {
    config.Port = "8081"
  }

  return config
}
