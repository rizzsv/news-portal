package handler

import (
	"news-portal/internal/adapter/handler/request"
	"news-portal/internal/adapter/handler/response"
	"news-portal/internal/core/domain/entity"
	"news-portal/internal/core/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var err error
var code string
var errorResp response.ErrorResponseDefault
var validate = validator.New()

type AuthHandler interface {
	Login(c *fiber.Ctx) error
}

type authHandler struct {
	authService service.AuthService
}

// Login implements AuthHandler.
func (a *authHandler) Login(c *fiber.Ctx) error {
	req := request.LoginRequest{}
	resp := response.SuccessAuthResponse{}

	if err = c.BodyParser(&req); err != nil {
		code = "[HANDLER] Login - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	if err = validate.Struct(req); err != nil {
		code = "[HANDLER] Login - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		// Cek jika error validasi
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var msg string
			for _, fieldErr := range validationErrors {
				msg += fieldErr.Field() + ": " + fieldErr.Tag() + ", "
			}
			errorResp.Meta.Message = "Validation failed: " + msg
		} else {
			errorResp.Meta.Message = err.Error()
		}
		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	reqLogin := entity.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	result, err := a.authService.GetUserByEmail(c.Context(), reqLogin)
	if err != nil {
		code = "[HANDLER] Login - 3"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusBadRequest).JSON(errorResp)
	}

	resp.Meta.Status = true
	resp.Meta.Message = "Login Successful"
	resp.AccessToken = result.AccessToken
	resp.ExpiresAt = result.ExpiresAt

	return c.Status(fiber.StatusOK).JSON(resp)
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}
