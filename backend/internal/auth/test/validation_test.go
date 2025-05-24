package test

import (
	"testing"

	"github.com/cpching/smart-recipe/backend/internal/auth"
	// validator "github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidateRegisterInput(t *testing.T) {
	v := auth.NewValidation()

	tests := []struct {
		name            string
		input           auth.RegisterInput
		wantErr         bool
		expectedMessage string
	}{
		{
			name: "valid input",
			input: auth.RegisterInput{
				Email:    "test@example.com",
				Password: "Strong-123",
			},
			wantErr: false,
		},
		{
			name: "missing email",
			input: auth.RegisterInput{
				Email:    "",
				Password: "Strong-123",
			},
			wantErr:         true,
			expectedMessage: "Email has an invalid format",
		},
		{
			name: "invalid email format",
			input: auth.RegisterInput{
				Email:    "invalid-email",
				Password: "Strong-123",
			},
			wantErr:         true,
			expectedMessage: "Email has an invalid format",
		},
		{
			name: "empty password",
			input: auth.RegisterInput{
				Email:    "test@example.com",
				Password: "",
			},
			wantErr:         true,
			expectedMessage: "Password is too weak",
		},
		{
			name: "weak password",
			input: auth.RegisterInput{
				Email:    "test@example.com",
				Password: "123",
			},
			wantErr:         true,
			expectedMessage: "Password is too weak",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Validate(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedMessage)

			} else {
				assert.NoError(t, err)
			}
		})
	}
}
