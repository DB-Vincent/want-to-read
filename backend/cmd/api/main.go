package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/DB-Vincent/want-to-read/docs"
	"github.com/DB-Vincent/want-to-read/internal/database"
	"github.com/DB-Vincent/want-to-read/internal/handlers"
	"github.com/DB-Vincent/want-to-read/internal/models"
	"github.com/DB-Vincent/want-to-read/internal/seed"
	"github.com/DB-Vincent/want-to-read/internal/services"
)

// @title			Want to Read API
// @version		1.0
// @description	API for managing your reading list
// @host			localhost:8080
// @BasePath		/
func main() {
	r := gin.Default()

	// Enable CORS
	// Apply CORS middleware globally so it covers all routes, including /api/login
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	apiRoutes := r.Group("/api")

	// Initialize services
	healthService := services.NewHealthService()
	bookService := services.NewBookService()
	userService := services.NewUserService()

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler(healthService)
	bookHandler := handlers.NewBookHandler(bookService, userService)
	userHandler := handlers.NewUserHandler(userService)

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		healthHandler.Check(c.Writer, c.Request)
	})

	// User authentication endpoint
	apiRoutes.POST("/login", userHandler.Login)
	apiRoutes.POST("/register", userHandler.AuthMiddleware(), userHandler.SuperUserMiddleware(), userHandler.Register)

	// User management endpoints
	apiRoutes.GET("/users", userHandler.AuthMiddleware(), userHandler.SuperUserMiddleware(), userHandler.ListUsers)
	apiRoutes.PATCH("/user/:id", userHandler.AuthMiddleware(), userHandler.SuperUserMiddleware(), userHandler.EditUser)
	apiRoutes.POST("/change-password", userHandler.AuthMiddleware(), userHandler.ChangePassword)

	apiRoutes.Use(userHandler.AuthMiddleware())
	{
		apiRoutes.GET("/users/:user_id/books", bookHandler.ListBooks)         // List books for a specific user
		apiRoutes.POST("/users/:user_id/books", bookHandler.AddBook)          // Add book for a specific user
		apiRoutes.PATCH("/users/:user_id/books/:id", bookHandler.UpdateBook)  // Update a user's book
		apiRoutes.DELETE("/users/:user_id/books/:id", bookHandler.DeleteBook) // Delete a user's book
	}

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer database.CloseDB()

	if err := database.DB.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	for _, seed := range seed.All() {
		if err := seed.Run(database.DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
