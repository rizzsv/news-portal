package repository

import (
	"context"
	"fmt"
	"news-portal/internal/core/domain/entity"
	"news-portal/internal/core/domain/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

var err error 
var code string

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, req entity.LoginRequest)(*entity.UserEntity, error)
}

type authRepository struct {
	db *gorm.DB
}

func (a *authRepository) GetUserByEmail(ctx context.Context, req entity.LoginRequest)(*entity.UserEntity, error) {
	var modelUser model.User

	err = a.db.Where("email = ?", req.Email).First(&modelUser).Error
	if err != nil {
		code = "[REPOSITORY] GetUserByEmail - 1"
		log.Errorw(code, err)
		return nil, err
	}

	resp := entity.UserEntity{
		ID:       fmt.Sprintf("%d", modelUser.ID),
		Email:    modelUser.Email,
		Name:     modelUser.Name,
		Password: modelUser.Password,
	} 
	return &resp, nil
}

func NewAuthReposiry(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}