package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sistema-libros-electronicos/internal/repositories"
	"sistema-libros-electronicos/internal/services"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) Create(c *gin.Context) {
	var input services.BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json invalido"})
		return
	}

	book, err := h.service.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) List(c *gin.Context) {
	books, err := h.service.List(filtersFromQuery(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) FindByID(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	book, err := h.service.FindByID(id)
	if err != nil {
		writeBookError(c, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Update(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	var input services.BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "json invalido"})
		return
	}

	book, err := h.service.Update(id, input)
	if err != nil {
		writeBookError(c, err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) Deactivate(c *gin.Context) {
	id, ok := parseID(c)
	if !ok {
		return
	}

	if err := h.service.Deactivate(id); err != nil {
		writeBookError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *BookHandler) Search(c *gin.Context) {
	h.List(c)
}

func filtersFromQuery(c *gin.Context) services.BookFilters {
	return services.BookFilters{
		Title:        c.Query("title"),
		Author:       c.Query("author"),
		Category:     c.Query("category"),
		Availability: c.Query("available"),
	}
}

func parseID(c *gin.Context) (uint, bool) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return 0, false
	}

	return uint(id), true
}

func writeBookError(c *gin.Context, err error) {
	if errors.Is(err, repositories.ErrBookNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
