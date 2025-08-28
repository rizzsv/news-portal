package middleware

import (
	"news-portal/config"
	"news-portal/internal/adapter/handler/response"
	"news-portal/lib/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type middleware interface {
	CheckToken() fiber.Handler
}

type Options struct {
	authJwt auth.Jwt
}

// CheckToken implements middleware.
func (o *Options) CheckToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var errorResponse response.ErrorResponseDefault
		autHandler := c.Get("Authorization")
		if autHandler == "" {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Unauthorized"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)
		}

		tokenString := strings.Split(autHandler, "Bearer")[1]
		claims, err := o.authJwt.VerifyToken(tokenString)
		if err != nil {
			errorResponse.Meta.Status = false
			errorResponse.Meta.Message = "Unauthorized"
			return c.Status(fiber.StatusUnauthorized).JSON(errorResponse)

		}
		c.Locals("users", claims)

		return c.Next()
	}
}

func NewMiddleware(cfg *config.Config) middleware {
	opt := new(Options)
	opt.authJwt = auth.NewJwt(cfg)

	return opt
}
