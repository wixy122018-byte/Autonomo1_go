package services

import (
	"errors"

	"sistema-libros-electronicos/internal/models"
	"sistema-libros-electronicos/internal/repositories"
)

var ErrBookUnavailable = errors.New("el libro no esta disponible para descarga")

type DownloadService struct {
	downloadRepo repositories.DownloadRepository
	bookRepo     repositories.BookRepository
}

func NewDownloadService(downloadRepo repositories.DownloadRepository, bookRepo repositories.BookRepository) *DownloadService {
	return &DownloadService{downloadRepo: downloadRepo, bookRepo: bookRepo}
}

type DownloadInput struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

func (s *DownloadService) Register(input DownloadInput) (models.Download, error) {
	book, err := s.bookRepo.FindByID(input.BookID)
	if err != nil {
		return models.Download{}, err
	}

	if !book.IsAvailable() {
		return models.Download{}, ErrBookUnavailable
	}

	return s.downloadRepo.Create(models.NewDownload(input.UserID, book))
}

func (s *DownloadService) History() ([]models.Download, error) {
	return s.downloadRepo.History()
}
