package app

import (
	"news-portal/config"
	"news-portal/lib/auth"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

func RunServer() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("[RunServer-1] Failed to load config")
	}


	_, err = cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal().Err(err).Msg("[RunServer-2] Failed to connect to database")
	}

	log.Info().Msg("[RunServer-3] Server successfully started with database connection")

	cdfR2, err := cfg.LoadCloudflareR2Config()
	if err != nil {
		log.Fatal().Err(err).Msg("[RunServer-4] Failed to load Cloudflare R2 config")
	}
	_ = s3.NewFromConfig(cdfR2)

	_ = auth.NewJwt(cfg)
}
