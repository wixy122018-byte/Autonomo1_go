package middleware

import (
	"net/http"
	"strings"

	"sistema-libros-electronicos/internal/services"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware valida la autenticación mediante JWT.
type AuthMiddleware struct {
	authService *services.AuthService
}

// NewAuthMiddleware crea una instancia del middleware.
func NewAuthMiddleware(authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// RequireAuth protege las rutas que necesitan un usuario autenticado.
func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token de autenticación requerido",
			})
			c.Abort()
			return
		}

		// El encabezado debe tener el formato: Bearer TOKEN
		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") ||
			strings.TrimSpace(parts[1]) == "" {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "formato de token inválido",
			})
			c.Abort()
			return
		}

		claims, err := m.authService.ValidateToken(
			strings.TrimSpace(parts[1]),
		)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// Se guardan los datos del usuario para usarlos en los handlers.
		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// RequireRoles permite el acceso únicamente a los roles autorizados.
func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleValue, exists := c.Get("role")

		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "rol del usuario no disponible",
			})
			c.Abort()
			return
		}

		userRole, ok := roleValue.(string)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "rol del usuario inválido",
			})
			c.Abort()
			return
		}

		for _, allowedRole := range allowedRoles {
			if strings.EqualFold(userRole, allowedRole) {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": "no tiene permisos para acceder a esta ruta",
		})
		c.Abort()
	}
}
