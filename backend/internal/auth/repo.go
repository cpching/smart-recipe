package auth

import (
	"context"
	"database/sql"

	"github.com/cpching/smart-recipe/backend/internal/domain"
	"github.com/cpching/smart-recipe/backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	dbUser := model.FromDomainUser(user)
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id, created_at`
	err := r.db.QueryRowxContext(ctx, query, dbUser.Email, dbUser.PasswordHash).Scan(&dbUser.ID, &dbUser.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}
	return model.ToDomainUser(dbUser), nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var dbUser model.UserDB
	err := r.db.GetContext(ctx, &dbUser, `SELECT * FROM users WHERE email = $1`, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, nil
		}
		return domain.User{}, err
	}
	return model.ToDomainUser(dbUser), nil
}
