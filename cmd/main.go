package main

import (
	"log"
	"wallet/config"
	"wallet/internal/app"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Error: %v", err)
	}

	app.Run(cfg)
}