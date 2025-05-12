package domain

type User struct {
	ID           int
	Email        string
	PasswordHash string
	CreatedAt    string
	LastLoginAt  *string
}
