package seeds

import (
	"news-portal/internal/core/domain/model"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := bcrypt.GenerateFromPassword([]byte("admin123"), 8)
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