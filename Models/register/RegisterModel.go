package register

type Register struct {
	Username string `validate:"required,email|numeric"`
	Password string `validate:"required,min=8"`
	ConfirmPassword string `validate:"required,eqfield=Password"`
}
