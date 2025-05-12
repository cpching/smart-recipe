package model

import "github.com/cpching/smart-recipe/backend/internal/domain"

type UserDB struct {
	ID           int     `db:"id"`
	Email        string  `db:"email"`
	PasswordHash string  `db:"password_hash"`
	CreatedAt    string  `db:"created_at"`
	LastLoginAt  *string `db:"last_login_at"`
}

// Mapper functions to convert between model and domain
func ToDomainUser(m UserDB) domain.User {
	return domain.User{
		ID:           m.ID,
		Email:        m.Email,
		PasswordHash: m.PasswordHash,
		CreatedAt:    m.CreatedAt,
		LastLoginAt:  m.LastLoginAt,
	}
}

func FromDomainUser(d domain.User) UserDB {
	return UserDB{
		ID:           d.ID,
		Email:        d.Email,
		PasswordHash: d.PasswordHash,
		CreatedAt:    d.CreatedAt,
		LastLoginAt:  d.LastLoginAt,
	}
}
