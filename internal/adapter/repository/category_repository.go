package repository

import (
	"context"
	"errors"
	"fmt"
	"news-portal/internal/core/domain/entity"
	"news-portal/internal/core/domain/model"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategory(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryByID(ctx context.Context, id int64) (*entity.CategoryEntity, error)
	CreateCategory(ctx context.Context, req entity.CategoryEntity) error
	UpdateCategory(ctx context.Context, req entity.CategoryEntity) error
	DeleteCategory(ctx context.Context, id int64) error
}

type categoryRepository struct {
	db *gorm.DB
}

// CreateCategory implements CategoryRepository.
func (c *categoryRepository) CreateCategory(ctx context.Context, req entity.CategoryEntity) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryRepository.
func (c *categoryRepository) DeleteCategory(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetCategory implements CategoryRepository.
func (c *categoryRepository) GetCategory(ctx context.Context) ([]entity.CategoryEntity, error) {
	var modelCategories []model.Category
	var err error
	var code string

	err = c.db.Order("created_at DESC").Preload("User").Find(&modelCategories).Error
	if err != nil {
		code = "[REPOSITORY] GetCategory - 1"
		log.Errorw(code, err)
		return nil, err
	}
	if len(modelCategories) == 0 {
		code = "[REPOSITORY] GetCategory - 2"
		err = errors.New("categories not found")
		log.Errorw(code, err)
		return nil, err
	}
	var resps []entity.CategoryEntity
	for _, val := range modelCategories {
		resps = append(resps, entity.CategoryEntity{
			ID:   fmt.Sprintf("%d", val.ID),
			Title: val.Title,
			Slug:  val.Slug,
			User: entity.UserEntity{
				ID:   fmt.Sprintf("%d", val.User.ID),
				Name: val.User.Name,
				Email: val.User.Email,
				Password: val.User.Password,
			},
		})
	}
	return resps, nil
}

// GetCategoryByID implements CategoryRepository.
func (c *categoryRepository) GetCategoryByID(ctx context.Context, id int64) (*entity.CategoryEntity, error) {
	panic("unimplemented")
}

// UpdateCategory implements CategoryRepository.
func (c *categoryRepository) UpdateCategory(ctx context.Context, req entity.CategoryEntity) error {
	panic("unimplemented")
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}
