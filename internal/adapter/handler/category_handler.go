package handler

import (
	"news-portal/internal/adapter/handler/response"
	"news-portal/internal/core/domain/entity"
	"news-portal/internal/core/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)	

var defaultSuccessResponse response.DefaultSuccessResponse

type CategoryHandler interface {
	GetCategory(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	UpdateCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

type categoryHandler struct {
	categoryService service.CategoryService
}

// CreateCategory implements CategoryHandler.
	func (ch *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteCategory implements CategoryHandler.
func (*categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetCategory implements CategoryHandler.
func (ch *categoryHandler) GetCategory(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	UserId := claims.UserId
	if UserId == 0 {
		code = "[HANDLER] GetCategory - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"
		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}
	result, err := ch.categoryService.GetCategory(c.Context())
	if err!= nil {
		code = "[HANDLER] GetCategory - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()
		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	categoryResponses := []response.SuccessCategoryResponse{}
	for _, result := range result {
		cr := response.SuccessCategoryResponse{
			ID: 		  result.ID,
			Title: result.Title,
			Slug: result.Slug,
			CreatedByName: result.User.Name,
		}
		categoryResponses = append(categoryResponses, cr)
	}
	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Categories fetched successfully"
	defaultSuccessResponse.Data = categoryResponses
	
	return c.JSON(defaultSuccessResponse)
}

// GetCategoryByID implements CategoryHandler.
func (ch *categoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateCategory implements CategoryHandler.
func (ch*categoryHandler) UpdateCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return &categoryHandler{categoryService: categoryService}
}
