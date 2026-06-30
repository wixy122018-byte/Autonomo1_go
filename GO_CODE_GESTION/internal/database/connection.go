package database

import (
	"sistema-libros-electronicos/internal/config"
	"sistema-libros-electronicos/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := config.GetDatabaseDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	err = DB.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Loan{},
		&models.Reservation{},
		&models.Download{},
	)

	if err != nil {
		return err
	}

	return nil
}