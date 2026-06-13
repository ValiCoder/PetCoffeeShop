package main

import (
	"log"

	_ "coffee-project/docs"
	"coffee-project/internal/database"
	"coffee-project/internal/handler"
	"coffee-project/internal/middleware"

	"coffee-project/internal/repository"
	"coffee-project/internal/usecase"

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

	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Инициализируем слои
	menuRepo := repository.NewPostgresMenuRepository(db.DB)
	menuUsecase := usecase.NewMenuUsecase(menuRepo)

	menuHandler := handler.NewMenuHandler(menuUsecase)
	userHandler := handler.NewUserHandler()
	// authHandler := handler.NewAuthHandler() // Закомментировано, пока не создашь auth.go

	// API route group
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", userHandler.PingHandler)

		v1.GET("/menu", menuHandler.GetMenu)
		v1.POST("/menu", menuHandler.CreateItem)
	}

	r.Run(":8080")
}
