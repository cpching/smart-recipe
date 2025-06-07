package test

import (
	"context"
	"testing"

	"github.com/cpching/smart-recipe/backend/internal/auth"
	"github.com/cpching/smart-recipe/backend/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepo struct {
	mock.Mock
}

func (m *mockRepo) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return domain.User{}, args.Error(1)
	}
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *mockRepo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return domain.User{}, args.Error(1)
	}
	return args.Get(0).(domain.User), args.Error(1)
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		setupRepo   func(*mockRepo)
		expectError error
	}{
		{
			name:     "successful register",
			email:    "test@example.com",
			password: "Abc-123-456",
			setupRepo: func(r *mockRepo) {
				r.On("GetByEmail", mock.Anything, "test@example.com").Return(nil, nil)
				r.On("CreateUser", mock.Anything, mock.Anything).Return(domain.User{Email: "test@example.com"}, nil)
			},
			expectError: nil,
		}, {
			name:     "duplicate email",
			email:    "taken@example.com",
			password: "Abc-123-456",
			setupRepo: func(r *mockRepo) {
				r.On("GetByEmail", mock.Anything, "taken@example.com").Return(domain.User{Email: "taken@example.com"}, nil)
			},
			expectError: auth.ErrEmailAlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(mockRepo)
			if tt.setupRepo != nil {
				tt.setupRepo(repo)
			}
			s := auth.NewAuthService(repo)
			_, err := s.Register(context.Background(), tt.email, tt.password)
			if tt.expectError != nil {
				assert.ErrorIs(t, err, tt.expectError)
			} else {
				assert.NoError(t, err)
			}

			repo.AssertExpectations(t)
		})
	}
}
