package model

type User struct {
	ID           int     `db:"id" json:"id"`
	Email        string  `db:"email" json:"email"`
	PasswordHash string  `db:"password_hash" json:"-"`
	CreatedAt    string  `db:"create_at" json:"created_at"`
	LastLoginAt  *string `db:"last_login_at" json:"last_login_at,omitempty"`
}
