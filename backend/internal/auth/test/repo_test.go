package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/cpching/smart-recipe/backend/internal/auth"
)

func setupTestPostgres(t *testing.T) (*sqlx.DB, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{ContainerRequest: req, Started: true})
	require.NoError(t, err)

	host, err := postgresC.Host(ctx)
	require.NoError(t, err)

	port, err := postgresC.MappedPort(ctx, "5432")
	require.NoError(t, err)

	dsn := fmt.Sprintf("postgres://test:test@%s:%s/testdb?sslmode=disable", host, port.Port())
	db, err := sqlx.Connect("postgres", dsn)
	require.NoError(t, err)

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	require.NoError(t, err)

	m, err := migrate.NewWithDatabaseInstance("file://../../../migrations", "postgres", driver)
	require.NoError(t, err)

	err = m.Up()
	require.NoError(t, err)

	cleanup := func() {
		_ = db.Close()
		_ = postgresC.Terminate(ctx)
	}

	return db, cleanup
}

func TestUserRepo_GetByEmail(t *testing.T) {
	db, teardown := setupTestPostgres(t)
	defer teardown()

	repo := auth.NewUserRepo(db)

	_, err := db.Exec(`INSERT INTO users (email, password_hash) VALUES ($1, $2)`, "test@example.com", "hashed-pass")
	require.NoError(t, err)

	user, err := repo.GetByEmail(context.Background(), "test@example.com")
	require.NoError(t, err)
	require.Equal(t, "test@example.com", user.Email)
}
