package service

import (
	"context"
	"news-portal/internal/adapter/repository"
	"news-portal/internal/core/domain/entity"

	"github.com/gofiber/fiber/v2/log"
)

type CategoryService interface {
	GetCategory(ctx context.Context) ([]entity.CategoryEntity, error)
	GetCategoryByID(ctx context.Context, id int64) (*entity.CategoryEntity, error)
	CreateCategory(ctx context.Context, req entity.CategoryEntity) error
	UpdateCategory(ctx context.Context, req entity.CategoryEntity) error
	DeleteCategory(ctx context.Context, id int64) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

// CreateCategory implements CategoryService.
func (c *categoryService) CreateCategory(ctx context.Context, req entity.CategoryEntity) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryService.
func (c *categoryService) DeleteCategory(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetCategory implements CategoryService.
func (c *categoryService) GetCategory(ctx context.Context) ([]entity.CategoryEntity, error) {
	result, err := c.categoryRepository.GetCategory(ctx)
	if err != nil {
		code = "[SERVICE] GetCategory - 1"
		log.Errorw(code, err)
		return nil, err
	}
	return result, nil
}

// GetCategoryByID implements CategoryService.
func (c *categoryService) GetCategoryByID(ctx context.Context, id int64) (*entity.CategoryEntity, error) {
	panic("unimplemented")
}

// UpdateCategory implements CategoryService.
func (c *categoryService) UpdateCategory(ctx context.Context, req entity.CategoryEntity) error {
	panic("unimplemented")
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepo}
}
