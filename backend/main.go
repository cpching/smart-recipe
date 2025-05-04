package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// TODO: Add user registration, login, and recipe endpoints

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func initDB() (*sqlx.DB, error) {
	dbConn := os.Getenv("DATABASE_URL")
	return sqlx.Connect("postgres", dbConn)
}
