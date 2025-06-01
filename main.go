package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	dotEnv "github.com/joho/godotenv"
	adminService "github.com/shariarfaisal/order-ms/pkg/admin/service"
	brandService "github.com/shariarfaisal/order-ms/pkg/brand/service"
	hubService "github.com/shariarfaisal/order-ms/pkg/hub/service"
	marketService "github.com/shariarfaisal/order-ms/pkg/market/service"
	orderService "github.com/shariarfaisal/order-ms/pkg/order/service"
	riderService "github.com/shariarfaisal/order-ms/pkg/rider/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}

func initEnv() {
	err := dotEnv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initEnv()

	env := os.Getenv("ENV")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	sslMode := "require"
	if env == "production" {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
	host, user, password, dbName, port, sslMode)

	// Connect to the database
	dbRes, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	db = dbRes

	// Set up Gin
	r := gin.Default()

	// Trust localhost for development
	if err := r.SetTrustedProxies([]string{"localhost"}); err != nil {
		log.Printf("Warning: failed to set trusted proxies: %v", err)
	}

	r.Use(JSONMiddleware())

	// Ping route
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Initialize services
	adminService.Init(db, r)
	hubService.Init(db, r)
	orderService.Init(db, r)
	riderService.Init(db, r)
	brandService.Init(db, r)
	marketService.Init(db, r)

	// Start server
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
