package models

import (
	"errors"
	"strings"
	"time"
)

var FixedCategories = [5]string{"Novela", "Tecnologia", "Ciencia", "Historia", "Educacion"}

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Author      string    `json:"author" gorm:"not null"`
	Category    string    `json:"category"`
	Editorial   string    `json:"editorial"`
	Year        int       `json:"year"`
	Description string    `json:"description"`
	Format      string    `json:"format"`
	ISBN        string    `json:"isbn" gorm:"unique"`
	Available   bool      `json:"available"`
	Active      bool      `json:"active" gorm:"default:true"`
	FilePath    string    `json:"file_path"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewBook(title, author, category, editorial, description, format, isbn, filePath string, year int, available bool) (Book, error) {
	book := Book{
		Title:       strings.TrimSpace(title),
		Author:      strings.TrimSpace(author),
		Category:    strings.TrimSpace(category),
		Editorial:   strings.TrimSpace(editorial),
		Year:        year,
		Description: strings.TrimSpace(description),
		Format:      strings.TrimSpace(format),
		ISBN:        strings.TrimSpace(isbn),
		Available:   available,
		Active:      true,
		FilePath:    strings.TrimSpace(filePath),
	}

	return book, book.Validate()
}

func (b Book) IsAvailable() bool {
	return b.Active && b.Available
}

func (b Book) Validate() error {
	if strings.TrimSpace(b.Title) == "" {
		return errors.New("el titulo es obligatorio")
	}

	if strings.TrimSpace(b.Author) == "" {
		return errors.New("el autor es obligatorio")
	}

	if strings.TrimSpace(b.Category) == "" {
		return errors.New("la categoria es obligatoria")
	}

	return nil
}
