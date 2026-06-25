package routes

import (
	"net/http"

	"sistema-libros-electronicos/internal/handlers"

	"github.com/gin-gonic/gin"
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

	// Rutas de usuarios
	router.GET("/users", handlers.ListUsers)
	router.GET("/users/:id", handlers.GetUserByID)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	// Rutas de seguridad
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}
