package services

import (
	"testing"

	"sistema-libros-electronicos/internal/models"
)

func TestFilterBooks(t *testing.T) {
	books := []models.Book{
		{ID: 1, Title: "Go Basico", Author: "Ana Perez", Category: "Tecnologia", Available: true, Active: true},
		{ID: 2, Title: "Historia Andina", Author: "Luis Mora", Category: "Historia", Available: false, Active: true},
		{ID: 3, Title: "Libro oculto", Author: "Ana Perez", Category: "Novela", Available: true, Active: false},
	}

	byTitle := FilterBooks(books, BookFilters{Title: "go"})
	if len(byTitle) != 1 || byTitle[0].ID != 1 {
		t.Fatalf("busqueda por titulo incorrecta: %#v", byTitle)
	}

	byAuthor := FilterBooks(books, BookFilters{Author: "ana"})
	if len(byAuthor) != 1 || byAuthor[0].ID != 1 {
		t.Fatalf("busqueda por autor incorrecta: %#v", byAuthor)
	}

	byCategory := FilterBooks(books, BookFilters{Category: "Historia"})
	if len(byCategory) != 1 || byCategory[0].ID != 2 {
		t.Fatalf("filtro por categoria incorrecto: %#v", byCategory)
	}

	available := FilterBooks(books, BookFilters{Availability: "true"})
	if len(available) != 1 || !available[0].IsAvailable() {
		t.Fatalf("filtro por disponibilidad incorrecto: %#v", available)
	}
}

func TestCountBooksByCategory(t *testing.T) {
	books := []models.Book{
		{Category: "Tecnologia", Active: true},
		{Category: "Tecnologia", Active: true},
		{Category: "Historia", Active: true},
		{Category: "Historia", Active: false},
	}

	counts := CountBooksByCategory(books)
	if counts["Tecnologia"] != 2 || counts["Historia"] != 1 {
		t.Fatalf("conteo por categoria incorrecto: %#v", counts)
	}
}
