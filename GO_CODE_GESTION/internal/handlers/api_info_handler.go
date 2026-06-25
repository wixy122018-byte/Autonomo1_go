package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebServiceInfo struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Description string `json:"description"`
	RequestBody string `json:"request_body,omitempty"`
}

func WebServicesCatalog(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"module": "Sistema de Gestión de Libros Electrónicos",
		"type":   "REST JSON API",
		"services": []WebServiceInfo{
			{Method: "POST", Path: "/register", Description: "Registra un nuevo usuario", RequestBody: "RegisterRequest"},
			{Method: "POST", Path: "/login", Description: "Valida credenciales y devuelve JWT", RequestBody: "LoginRequest"},
			{Method: "GET", Path: "/profile", Description: "Perfil del usuario autenticado (requiere JWT)"},
			{Method: "GET", Path: "/users", Description: "Lista todos los usuarios (solo admin, requiere JWT)"},
			{Method: "GET", Path: "/users/:id", Description: "Consulta un usuario por ID (solo admin, requiere JWT)"},
			{Method: "POST", Path: "/books", Description: "Registra un libro nuevo (admin/bibliotecario)", RequestBody: "BookInput"},
			{Method: "GET", Path: "/books", Description: "Lista libros activos con filtros por query params (requiere JWT)"},
			{Method: "GET", Path: "/books/search", Description: "Busca libros por título, autor, categoría o disponibilidad"},
			{Method: "GET", Path: "/books/:id", Description: "Consulta un libro por ID (requiere JWT)"},
			{Method: "PUT", Path: "/books/:id", Description: "Actualiza los datos de un libro (admin/bibliotecario)", RequestBody: "BookInput"},
			{Method: "DELETE", Path: "/books/:id", Description: "Desactiva un libro sin eliminarlo (admin/bibliotecario)"},
			{Method: "POST", Path: "/downloads", Description: "Registra una descarga si el libro está disponible (requiere JWT)", RequestBody: "DownloadInput"},
			{Method: "GET", Path: "/downloads/history", Description: "Consulta el historial de descargas (requiere JWT)"},
		},
	})
}
