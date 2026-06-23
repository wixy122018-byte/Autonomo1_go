package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sistema-libros-electronicos/internal/database"
	"sistema-libros-electronicos/internal/handlers"
	"sistema-libros-electronicos/internal/repositories"
	"sistema-libros-electronicos/internal/services"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Sistema de Gestión de Libros Electrónicos",
			"status":  "Servidor en línea",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	bookRepository := repositories.NewGormBookRepository(database.DB)
	downloadRepository := repositories.NewGormDownloadRepository(database.DB)
	userRepository := repositories.NewGormUserRepository(database.DB)

	bookService := services.NewBookService(bookRepository)
	downloadService := services.NewDownloadService(downloadRepository, bookRepository)
	authService := services.NewAuthService(userRepository)

	bookHandler := handlers.NewBookHandler(bookService)
	downloadHandler := handlers.NewDownloadHandler(downloadService)
	authHandler := handlers.NewAuthHandler(authService)

	router.GET("/api/v1/services", handlers.WebServicesCatalog)

	router.POST("/users/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.GET("/users/me", authHandler.Me)

	router.POST("/books", bookHandler.Create)
	router.GET("/books", bookHandler.List)
	router.GET("/books/search", bookHandler.Search)
	router.GET("/books/:id", bookHandler.FindByID)
	router.PUT("/books/:id", bookHandler.Update)
	router.DELETE("/books/:id", bookHandler.Deactivate)

	router.POST("/downloads", downloadHandler.Register)
	router.GET("/downloads/history", downloadHandler.History)

	api := router.Group("/api/v1")
	api.POST("/users/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
	api.GET("/users/me", authHandler.Me)
	api.POST("/books", bookHandler.Create)
	api.GET("/books", bookHandler.List)
	api.GET("/books/search", bookHandler.Search)
	api.GET("/books/:id", bookHandler.FindByID)
	api.PUT("/books/:id", bookHandler.Update)
	api.DELETE("/books/:id", bookHandler.Deactivate)
	api.POST("/downloads", downloadHandler.Register)
	api.GET("/downloads/history", downloadHandler.History)
}
