package auth

type RegisterInput struct {
	Email    string `json:"email" example:"user@example.com" validate:"required,email"`
	Password string `json:"password" example:"Secret123!" validate:"required,password"`
}
