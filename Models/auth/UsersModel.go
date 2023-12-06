package auth

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

type User struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
} 