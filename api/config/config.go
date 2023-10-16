package config

import "os"

type Config struct {
  Port string
}

// Initialize configurations
func Initialize() Config {
  var config Config

  config.Port = os.Getenv("PORT")

  if (config.Port == "") {
    config.Port = "8080"
  }

  return config
}
