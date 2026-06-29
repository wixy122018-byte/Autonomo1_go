package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type WebServiceInfo struct {
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	Description string   `json:"description"`
	RequestBody string   `json:"request_body,omitempty"`
	RequiresJWT bool     `json:"requires_jwt"`
	Roles       []string `json:"roles,omitempty"`
}

func WebServicesCatalog(c *gin.Context) {
	services := []WebServiceInfo{
		{Method: "POST", Path: "/register", Description: "Registra un nuevo usuario en la plataforma", RequestBody: "RegisterRequest", RequiresJWT: false},
		{Method: "POST", Path: "/login", Description: "Valida credenciales de usuario y devuelve el token JWT", RequestBody: "LoginRequest", RequiresJWT: false},
		{Method: "GET", Path: "/api/v1/services", Description: "Muestra este catálogo interactivo de servicios web", RequiresJWT: false},
		{Method: "GET", Path: "/profile", Description: "Obtiene la información de perfil del usuario autenticado", RequiresJWT: true},
		{Method: "GET", Path: "/books", Description: "Lista todos los libros activos, permitiendo filtros avanzados (título, autor, categoría, disponibilidad)", RequiresJWT: true},
		{Method: "GET", Path: "/books/search", Description: "Busca libros en tiempo real según un término general", RequiresJWT: true},
		{Method: "GET", Path: "/books/:id", Description: "Obtiene los detalles detallados de un libro por su ID", RequiresJWT: true},
		{Method: "POST", Path: "/books", Description: "Registra un nuevo libro físico o digital en el catálogo", RequestBody: "BookInput", RequiresJWT: true, Roles: []string{"ADMINISTRADOR", "BIBLIOTECARIO"}},
		{Method: "PUT", Path: "/books/:id", Description: "Actualiza los datos (título, autor, categoría, disponibilidad) de un libro existente", RequestBody: "BookInput", RequiresJWT: true, Roles: []string{"ADMINISTRADOR", "BIBLIOTECARIO"}},
		{Method: "DELETE", Path: "/books/:id", Description: "Desactiva un libro del catálogo (soft-delete para no romper registros)", RequiresJWT: true, Roles: []string{"ADMINISTRADOR", "BIBLIOTECARIO"}},
		{Method: "POST", Path: "/downloads", Description: "Registra la descarga de un libro por parte de un lector", RequestBody: "DownloadInput", RequiresJWT: true},
		{Method: "GET", Path: "/downloads/history", Description: "Consulta el historial de descargas realizadas por el lector autenticado", RequiresJWT: true},
		{Method: "GET", Path: "/users", Description: "Lista todos los usuarios registrados en el sistema", RequiresJWT: true, Roles: []string{"ADMINISTRADOR"}},
		{Method: "GET", Path: "/users/:id", Description: "Obtiene los detalles de un usuario específico por su ID", RequiresJWT: true, Roles: []string{"ADMINISTRADOR"}},
		{Method: "DELETE", Path: "/users/:id", Description: "Elimina permanentemente una cuenta de usuario del sistema", RequiresJWT: true, Roles: []string{"ADMINISTRADOR"}},
	}

	// Comprobar si el cliente prefiere HTML (ej. navegadores web)
	accept := c.GetHeader("Accept")
	format := c.Query("format")

	if strings.Contains(accept, "text/html") && format != "json" {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, renderHTML(services))
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"module":   "Sistema de Gestión de Libros Electrónicos",
		"type":     "REST JSON API",
		"services": services,
	})
}
