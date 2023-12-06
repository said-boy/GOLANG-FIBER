package auth

type Auth struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
} 