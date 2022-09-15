package main

import (
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"statistics-service/config"
	"statistics-service/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.SetLogrus(cfg.Log.Level)

	// Run
	app.Run(cfg)
}
