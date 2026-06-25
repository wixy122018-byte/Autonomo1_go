package repositories

import (
	"errors"

	"sistema-libros-electronicos/internal/database"
	"sistema-libros-electronicos/internal/models"

	"gorm.io/gorm"
)

// UserRepository define las operaciones que se pueden realizar
// sobre los usuarios dentro del sistema.
type UserRepository interface {
	Create(*models.User) error
	GetAll() ([]models.User, error)
	GetByID(uint) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	Update(uint, *models.User) (*models.User, error)
	Delete(uint) error
	EmailExists(string) (bool, error)
}

type GormUserRepository struct{}

// NewUserRepository crea una instancia del repositorio de usuarios.
func NewUserRepository() UserRepository {
	return &GormUserRepository{}
}

func (r *GormUserRepository) Create(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *GormUserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func (r *GormUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GormUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *GormUserRepository) Update(id uint, user *models.User) (*models.User, error) {
	existingUser, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user.Name != "" {
		existingUser.Name = user.Name
	}

	if user.Email != "" {
		existingUser.Email = user.Email
	}

	if user.Password != "" {
		existingUser.Password = user.Password
	}

	if user.Role != "" {
		existingUser.Role = user.Role
	}

	err = database.DB.Save(existingUser).Error
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (r *GormUserRepository) Delete(id uint) error {
	result := database.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *GormUserRepository) EmailExists(email string) (bool, error) {
	var user models.User

	err := database.DB.Where("email = ?", email).First(&user).Error

	if err == nil {
		return true, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return false, err
}
