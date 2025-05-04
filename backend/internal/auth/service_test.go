package auth

import (
	"context"
	"testing"

	"github.com/cpching/smart-recipe/backend/internal/model"
	"github.com/stretchr/testify/assert"
)

type fakeRepo struct{}

func (f *fakeRepo) CreateUser(ctx context.Context, u model.User) (model.User, error) {
	u.ID = 1
	u.CreatedAt = "2025-05-03T12:00:00"
	return u, nil
}

func (f *fakeRepo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	return model.User{}, nil
}

func TestRegister_InvalidEmail(t *testing.T) {
	svc := NewAuthService(&fakeRepo{})
	_, err := svc.Register(context.Background(), "bademail", "Password123")
	assert.ErrorIs(t, err, ErrInvalidEmail)
}

func TestRegister_WeakPassword(t *testing.T) {
	svc := NewAuthService(&fakeRepo{})
	_, err := svc.Register(context.Background(), "foo@example.com", "short")
	assert.ErrorIs(t, err, ErrWeakPassword)
	_, err = svc.Register(context.Background(), "foo@example.com", "weak0123")
	assert.ErrorIs(t, err, ErrWeakPassword)
}

func TestRegister_Success(t *testing.T) {
	svc := NewAuthService(&fakeRepo{})
	user, err := svc.Register(context.Background(), "foo@example.com", "Strong-Password1")
	assert.NoError(t, err)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "foo@example.com", user.Email)
	assert.NotEmpty(t, user.PasswordHash)
}
