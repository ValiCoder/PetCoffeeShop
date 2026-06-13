package main

import (
	"log"

	_ "coffee-project/docs"
	"coffee-project/internal/database"
	"coffee-project/internal/handler"
	"coffee-project/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Coffee Shop API
// @version         1.0
// @description     Backend API for managing coffee shop items, users, and roles.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, processing with system env")
	}

	db := database.ConnectDB()
	defer db.Close()

	r := gin.Default()

	r.Use(middleware.Logger())

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Инициализируем хендлер
	userHandler := handler.NewUserHandler()

	// API route group
	v1 := r.Group("/api/v1")
	{
		// Передаем метод нашей структуры
		v1.GET("/ping", userHandler.PingHandler)
	}

	r.Run(":8080")
}
