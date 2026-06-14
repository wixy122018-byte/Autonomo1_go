package services

import (
	"strconv"
	"strings"

	"sistema-libros-electronicos/internal/models"
	"sistema-libros-electronicos/internal/repositories"
)

type BookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) *BookService {
	return &BookService{repo: repo}
}

type BookInput struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Editorial   string `json:"editorial"`
	Year        int    `json:"year"`
	Description string `json:"description"`
	Format      string `json:"format"`
	ISBN        string `json:"isbn"`
	Available   bool   `json:"available"`
	FilePath    string `json:"file_path"`
}

type BookFilters struct {
	Title        string
	Author       string
	Category     string
	Availability string
}

func (s *BookService) Create(input BookInput) (models.Book, error) {
	book, err := models.NewBook(
		input.Title,
		input.Author,
		input.Category,
		input.Editorial,
		input.Description,
		input.Format,
		input.ISBN,
		input.FilePath,
		input.Year,
		input.Available,
	)
	if err != nil {
		return models.Book{}, err
	}

	return s.repo.Create(book)
}

func (s *BookService) List(filters BookFilters) ([]models.Book, error) {
	books, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return FilterBooks(books, filters), nil
}

func (s *BookService) FindByID(id uint) (models.Book, error) {
	return s.repo.FindByID(id)
}

func (s *BookService) Update(id uint, input BookInput) (models.Book, error) {
	book, err := models.NewBook(
		input.Title,
		input.Author,
		input.Category,
		input.Editorial,
		input.Description,
		input.Format,
		input.ISBN,
		input.FilePath,
		input.Year,
		input.Available,
	)
	if err != nil {
		return models.Book{}, err
	}

	return s.repo.Update(id, book)
}

func (s *BookService) Deactivate(id uint) error {
	return s.repo.Deactivate(id)
}

func FilterBooks(books []models.Book, filters BookFilters) []models.Book {
	filterMap := map[string]string{
		"title":        strings.TrimSpace(filters.Title),
		"author":       strings.TrimSpace(filters.Author),
		"category":     strings.TrimSpace(filters.Category),
		"availability": strings.TrimSpace(filters.Availability),
	}

	filtered := []models.Book{}
	for _, book := range books {
		if !book.Active {
			continue
		}
		if filterMap["title"] != "" && !containsText(book.Title, filterMap["title"]) {
			continue
		}
		if filterMap["author"] != "" && !containsText(book.Author, filterMap["author"]) {
			continue
		}
		if filterMap["category"] != "" && !strings.EqualFold(book.Category, filterMap["category"]) {
			continue
		}
		if filterMap["availability"] != "" {
			available, err := strconv.ParseBool(filterMap["availability"])
			if err == nil && book.IsAvailable() != available {
				continue
			}
		}

		filtered = append(filtered, book)
	}

	return filtered
}

func CountBooksByCategory(books []models.Book) map[string]int {
	counts := map[string]int{}
	for _, book := range books {
		if book.Active {
			counts[book.Category]++
		}
	}

	return counts
}

func containsText(value, query string) bool {
	return strings.Contains(strings.ToLower(value), strings.ToLower(query))
}
