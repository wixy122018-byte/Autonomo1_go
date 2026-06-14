package repositories

import (
	"errors"

	"sistema-libros-electronicos/internal/models"

	"gorm.io/gorm"
)

var ErrBookNotFound = errors.New("libro no encontrado")

type BookRepository interface {
	Create(book models.Book) (models.Book, error)
	FindAll() ([]models.Book, error)
	FindByID(id uint) (models.Book, error)
	Update(id uint, book models.Book) (models.Book, error)
	Deactivate(id uint) error
}

type GormBookRepository struct {
	db *gorm.DB
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{db: db}
}

func (r *GormBookRepository) Create(book models.Book) (models.Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r *GormBookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	if err := r.db.Where("active = ?", true).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *GormBookRepository) FindByID(id uint) (models.Book, error) {
	var book models.Book
	err := r.db.Where("id = ? AND active = ?", id, true).First(&book).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Book{}, ErrBookNotFound
	}

	return book, err
}

func (r *GormBookRepository) Update(id uint, book models.Book) (models.Book, error) {
	current, err := r.FindByID(id)
	if err != nil {
		return models.Book{}, err
	}

	book.ID = current.ID
	book.CreatedAt = current.CreatedAt
	book.Active = true

	if err := r.db.Save(&book).Error; err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r *GormBookRepository) Deactivate(id uint) error {
	result := r.db.Model(&models.Book{}).
		Where("id = ? AND active = ?", id, true).
		Updates(map[string]any{"active": false, "available": false})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrBookNotFound
	}

	return nil
}
