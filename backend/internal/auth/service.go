package auth

import (
	"context"
	"errors"

	"github.com/cpching/smart-recipe/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (domain.User, error)
	// Login()
	// Logout()
	// Delete (terminate)
}
type authService struct {
	repo UserRepo
}

var (
	ErrEmailAlreadyExists = errors.New("EMAIL ALREADY EXISTS")
)

func NewAuthService(r UserRepo) AuthService {
	return &authService{repo: r}
}

func (s *authService) Register(ctx context.Context, email, password string) (domain.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)

	if err != nil {
		return domain.User{}, err
	}

	if user.Email != "" {
		return domain.User{}, ErrEmailAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return domain.User{}, err
	}

	user = domain.User{
		Email:        email,
		PasswordHash: string(hash),
	}

	return s.repo.CreateUser(ctx, user)
}
