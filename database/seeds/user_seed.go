package seeds

import (
	"news-portal/internal/core/domain/model"
	"news-portal/lib/conv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to hash password")
	}

	admin := model.User{
		Name:    "rizzz",
		Email:   "rizq.syafriano@gmail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "rizq.syafriano@gmail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("failed to seed user")
	} else {
		log.Info().Msg("user seeded")
	}
}