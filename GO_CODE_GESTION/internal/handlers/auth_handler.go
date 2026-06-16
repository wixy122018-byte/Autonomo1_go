package handlers

import (
	"errors"
	"net/http"

	"sistema-libros-electronicos/internal/models"
	"sistema-libros-electronicos/internal/repositories"
	"sistema-libros-electronicos/internal/services"

	"github.com/gin-gonic/gin"
)

// AuthHandler recibe las solicitudes HTTP de registro y login.
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler crea una nueva instancia del handler de autenticación.
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// RegisterRequest representa los datos requeridos para registrar un usuario.
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

// LoginRequest representa los datos necesarios para iniciar sesión.
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register registra un nuevo usuario en el sistema.
func (h *AuthHandler) Register(c *gin.Context) {
	var request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "los datos enviados no son válidos",
		})
		return
	}

	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
	}

	registeredUser, err := h.authService.Register(user)

	if err != nil {
		if errors.Is(err, repositories.ErrEmailAlreadyExist) {
			c.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "usuario registrado correctamente",
		"user":    registeredUser,
	})
}

// Login autentica al usuario y devuelve un token JWT.
func (h *AuthHandler) Login(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "correo y contraseña son obligatorios",
		})
		return
	}

	token, user, err := h.authService.Login(
		request.Email,
		request.Password,
	)

	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if errors.Is(err, services.ErrJWTSecretRequired) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "inicio de sesión correcto",
		"token":   token,
		"user":    user,
	})
}
