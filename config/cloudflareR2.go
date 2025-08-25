package config

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/rs/zerolog/log"
)

func (cfg Config) LoadCloudflareR2Config() (aws.Config, error) {
	conf, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
		cfg.CloudFlareR2.APIKey, cfg.CloudFlareR2.APISecret, "", 
	)), awsConfig.WithRegion("auto"),)
	if err != nil {
		log.Fatal().Msgf("unable to load SDK config, " + err.Error())
	}

	log.Info().Msg("Cloudflare R2 configuration loaded")

	return conf, nil
	
}