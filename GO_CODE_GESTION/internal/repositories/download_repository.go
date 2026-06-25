package repositories

import (
	"sistema-libros-electronicos/internal/models"

	"gorm.io/gorm"
)

type DownloadRepository interface {
	Create(download models.Download) (models.Download, error)
	History() ([]models.Download, error)
}

type GormDownloadRepository struct {
	db *gorm.DB
}

func NewGormDownloadRepository(db *gorm.DB) *GormDownloadRepository {
	return &GormDownloadRepository{db: db}
}

func (r *GormDownloadRepository) Create(download models.Download) (models.Download, error) {
	if err := r.db.Create(&download).Error; err != nil {
		return models.Download{}, err
	}

	return download, nil
}

func (r *GormDownloadRepository) History() ([]models.Download, error) {
	var downloads []models.Download
	if err := r.db.Order("download_date DESC").Find(&downloads).Error; err != nil {
		return nil, err
	}

	return downloads, nil
}
