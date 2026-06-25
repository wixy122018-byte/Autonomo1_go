package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"sistema-libros-electronicos/internal/database"
	"sistema-libros-electronicos/internal/handlers"
	authmiddleware "sistema-libros-electronicos/internal/middleware"
	"sistema-libros-electronicos/internal/models"
	"sistema-libros-electronicos/internal/repositories"
	"sistema-libros-electronicos/internal/services"
)

// RegisterRoutes registra todas las rutas disponibles en el sistema.
func RegisterRoutes(router *gin.Engine) {
	// Rutas iniciales creadas en la base del proyecto.
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

	// Inicialización de las capas del módulo de usuarios.
	userRepository := repositories.NewUserRepository(database.DB)

	userService := services.NewUserService(userRepository)

	jwtSecret := os.Getenv("JWT_SECRET")

	authService := services.NewAuthService(
		userRepository,
		jwtSecret,
	)

	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)

	authMiddleware := authmiddleware.NewAuthMiddleware(authService)

	// Rutas públicas: no necesitan token.
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	// Rutas protegidas: necesitan un token JWT válido.
	protected := router.Group("/")
	protected.Use(authMiddleware.RequireAuth())
	{
		protected.GET("/profile", userHandler.Profile)

		// Solo el administrador puede consultar todos los usuarios.
		protected.GET(
			"/users",
			authmiddleware.RequireRoles(models.RoleAdministrador),
			userHandler.GetUsers,
		)

		// Solo el administrador puede consultar usuarios por ID.
		protected.GET(
			"/users/:id",
			authmiddleware.RequireRoles(models.RoleAdministrador),
			userHandler.GetUserByID,
		)

		// Solo el administrador puede eliminar usuarios por ID.
		protected.DELETE(
			"/users/:id",
			authmiddleware.RequireRoles(models.RoleAdministrador),
			userHandler.DeleteUser,
		)
	}
}
