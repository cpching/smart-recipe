package auth

import (
	"context"

	"github.com/cpching/smart-recipe/backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	CreateUser(ctx context.Context, u model.User) (model.User, error)
	GetByEmail(ctx context.Context, email string) (model.User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(ctx context.Context, u model.User) (model.User, error) {
	query := `INSERT INTO Users (email, password_hash) VALUES ($1, $2) RETURNING id, created_at`
	err := r.db.QueryRowContext(ctx, query, u.Email, u.PasswordHash).Scan(&u.ID, &u.CreatedAt)
	return u, err
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var u model.User
	err := r.db.GetContext(ctx, &u, `SELECT * FROM Users WHERE email = $1`, email)
	return u, err
}
