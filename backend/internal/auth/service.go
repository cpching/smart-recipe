package auth

import (
	"context"

	"github.com/cpching/smart-recipe/backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, email, password string) (domain.User, error)
}
type authService struct {
	repo UserRepo
}

func NewAuthService(r UserRepo) AuthService {
	return &authService{repo: r}
}

func (s *authService) Register(ctx context.Context, email, password string) (domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{
		Email:        email,
		PasswordHash: string(hash),
	}
	return s.repo.CreateUser(ctx, user)
}
