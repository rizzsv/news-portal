package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	var errorMessages []string
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "email":
				errorMessages = append(errorMessages, "Invalid email format")
			case "required":
				errorMessages = append(errorMessages, "Field is required")
			case "min":
				if err.Field() == "Password" {
					errorMessages = append(errorMessages, "Password must be at least 8 characters long")
				}
			default:
				errorMessages = append(errorMessages, "Invalid field")
			}
		}
			return errors.New("Validation failed: " + joinMessage(errorMessages))
	}

	return nil
	
}

func joinMessage(messages []string) string {
	result := ""
	for i, message := range messages {
		if i > 0 {
			result += ", "
		}
		result += message
	}
	return result
}


