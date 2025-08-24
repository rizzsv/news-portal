package app

import (
	"news-portal/config"

	"github.com/rs/zerolog/log"
)

func RunServer() {
	// Load config dari .env (path = root project)
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("[RunServer-1] Failed to load config")
	}

	// Tes koneksi ke PostgreSQL
	_, err = cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal().Err(err).Msg("[RunServer-2] Failed to connect to database")
	}

	log.Info().Msg("[RunServer-3] Server successfully started with database connection")
}
