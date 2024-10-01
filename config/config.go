package config

import (
	"log"
)

type Config struct {
	Port string
}

var AppConfig Config

func LoadConfig() {
	AppConfig = Config{
		Port: ":8080",
	}
	log.Println("Configuration loaded")
}
