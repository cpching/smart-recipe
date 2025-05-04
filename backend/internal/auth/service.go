package auth

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"

	"github.com/cpching/smart-recipe/backend/internal/model"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrWeakPassword = errors.New("password too weak")
)

var emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

type AuthService interface {
	Register(ctx context.Context, email, password string) (model.User, error)
}

type authService struct {
	repo UserRepo
}

func NewAuthService(r UserRepo) AuthService {
	return &authService{repo: r}
}

func (s *authService) Register(ctx context.Context, email, password string) (model.User, error) {
	if !emailRegex.MatchString(email) {
		return model.User{}, ErrInvalidEmail
	}

	if len(password) < 8 {
		return model.User{}, ErrWeakPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	return s.repo.CreateUser(ctx, model.User{
		Email:        email,
		PasswordHash: string(hash),
	})
}
