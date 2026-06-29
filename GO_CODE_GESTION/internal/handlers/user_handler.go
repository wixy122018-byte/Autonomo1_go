package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"sistema-libros-electronicos/internal/repositories"
	"sistema-libros-electronicos/internal/services"

	"github.com/gin-gonic/gin"
)

// UserHandler recibe las solicitudes HTTP relacionadas con usuarios.
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler crea una nueva instancia del handler de usuarios.
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsers devuelve todos los usuarios registrados.
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "no se pudieron obtener los usuarios",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// GetUserByID devuelve un usuario mediante su identificador.
func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "el identificador del usuario no es válido",
		})
		return
	}

	user, err := h.userService.GetUserByID(uint(id))

	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "no se pudo obtener el usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// Profile devuelve la información del usuario autenticado.
// El middleware JWT coloca el identificador del usuario en el contexto.
func (h *UserHandler) Profile(c *gin.Context) {
	userIDValue, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "usuario no autenticado",
		})
		return
	}

	userID, ok := userIDValue.(uint)

	if !ok || userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "identificador de usuario inválido",
		})
		return
	}

	user, err := h.userService.GetUserByID(userID)

	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "no se pudo obtener el perfil",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "perfil obtenido correctamente",
		"user":    user,
	})
}

// DeleteUser elimina un usuario mediante su identificador.
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "el identificador del usuario no es válido",
		})
		return
	}

	err = h.userService.DeleteUser(uint(id))

	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "no se pudo eliminar el usuario",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "usuario eliminado correctamente",
	})
}
