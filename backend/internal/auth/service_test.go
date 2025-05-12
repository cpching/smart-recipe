package auth

import (
	"context"
	"testing"

	"github.com/cpching/smart-recipe/backend/internal/domain"
	// "github.com/stretchr/testify/assert"
)

type fakeRepo struct{}

func (f *fakeRepo) CreateUser(ctx context.Context, u domain.User) (domain.User, error) {
	u.ID = 1
	u.CreatedAt = "2025-05-03T12:00:00"
	return u, nil
}

func (f *fakeRepo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	return domain.User{}, nil
}

func TestRegister_InvalidEmail(t *testing.T) {
	// svc := NewAuthService(&fakeRepo{})
}

func TestRegister_WeakPassword(t *testing.T) {
	// svc := NewAuthService(&fakeRepo{})
	// _, err := svc.Register(context.Background(), "foo@example.com", "short")
	// assert.ErrorIs(t, err, ErrWeakPassword)
	// _, err = svc.Register(context.Background(), "foo@example.com", "weak0123")
	// assert.ErrorIs(t, err, ErrWeakPassword)
	// _, err = svc.Register(context.Background(), "foo@example.com", "suk-suk-suk")
	// assert.NoError(t, err)
}

func TestRegister_Success(t *testing.T) {
	// svc := NewAuthService(&fakeRepo{})
	// user, err := svc.Register(context.Background(), "foo@example.com", "Strong-Password1")
	// assert.ErrorIs(t, err, ErrWeakPassword)
	// assert.NoError(t, err)
	// assert.Equal(t, 1, user.ID)
	// assert.Equal(t, "foo@example.com", user.Email)
	// assert.NotEmpty(t, user.PasswordHash)
}
