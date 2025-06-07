// @title           Smart Recipe API
// @version         1.0
// @description     API for Smart Recipe backend
// @host            localhost:8080

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/cpching/smart-recipe/backend/docs"
	"github.com/cpching/smart-recipe/backend/internal/auth"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	v := auth.NewValidation()

	db, err := initDB()
	if err != nil {
		log.Fatalf("🚨 Cannot connect to DB: %v", err)
	}
	repo := auth.NewUserRepo(db)
	h := auth.NewHandler(auth.NewAuthService(repo), v)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.POST("/register", h.MiddlewareValidateUser, h.Register)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func initDB() (*sqlx.DB, error) {
	// 	dbConn := os.Getenv("DATABASE_URL")
	// 	return sqlx.Connect("postgres", dbConn)
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL not set")
	}

	db, err := waitForDB(dsn, 10)

	return db, err
}

func waitForDB(dsn string, s int) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error
	for i := 0; i < s; i++ {
		db, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			return db, nil
		}
		log.Println("Waiting for DB...", err)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to DB after retries: %w", err)
}
