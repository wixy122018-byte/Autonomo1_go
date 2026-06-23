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
		"module": "Integrante 3 - libros, busqueda, descargas y autenticacion",
		"type":   "REST JSON API",
		"services": []WebServiceInfo{
			{Method: "POST", Path: "/api/v1/books", Description: "Registra un libro nuevo", RequestBody: "BookInput"},
			{Method: "GET", Path: "/api/v1/books", Description: "Lista libros activos y permite filtros por query params"},
			{Method: "GET", Path: "/api/v1/books/:id", Description: "Consulta un libro por ID"},
			{Method: "PUT", Path: "/api/v1/books/:id", Description: "Actualiza los datos de un libro", RequestBody: "BookInput"},
			{Method: "DELETE", Path: "/api/v1/books/:id", Description: "Desactiva un libro sin eliminarlo fisicamente"},
			{Method: "GET", Path: "/api/v1/books/search", Description: "Busca libros por titulo, autor, categoria o disponibilidad"},
			{Method: "POST", Path: "/api/v1/downloads", Description: "Registra una descarga si el libro esta disponible", RequestBody: "DownloadInput"},
			{Method: "GET", Path: "/api/v1/downloads/history", Description: "Consulta el historial de descargas"},
			{Method: "POST", Path: "/api/v1/users/register", Description: "Registra usuario con password encriptado", RequestBody: "RegisterInput"},
			{Method: "POST", Path: "/api/v1/login", Description: "Valida credenciales y devuelve JWT", RequestBody: "LoginInput"},
			{Method: "GET", Path: "/api/v1/users/me", Description: "Devuelve el usuario autenticado usando Authorization Bearer"},
		},
	})
}
