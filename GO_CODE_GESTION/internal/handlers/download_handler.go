package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"sistema-libros-electronicos/internal/repositories"
	"sistema-libros-electronicos/internal/services"
)

type DownloadHandler struct {
	service *services.DownloadService
}

func NewDownloadHandler(service *services.DownloadService) *DownloadHandler {
	return &DownloadHandler{service: service}
}

func (h *DownloadHandler) Register(c *gin.Context) {
	var input services.DownloadInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json invalido"})
		return
	}

	download, err := h.service.Register(input)
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, repositories.ErrBookNotFound) {
			status = http.StatusNotFound
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, download)
}

func (h *DownloadHandler) History(c *gin.Context) {
	downloads, err := h.service.History()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, downloads)
}
