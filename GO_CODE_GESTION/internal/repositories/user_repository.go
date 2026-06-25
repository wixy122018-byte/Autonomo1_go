package repositories

import (
	"errors"

	"sistema-libros-electronicos/internal/models"

	"gorm.io/gorm"
)

// Errores controlados del repositorio.
var (
	ErrUserNotFound      = errors.New("usuario no encontrado")
	ErrEmailAlreadyExist = errors.New("el correo ya está registrado")
)

// UserRepository define las operaciones que se pueden realizar con usuarios.
// Esta interfaz permite utilizar PostgreSQL u otra fuente de datos.
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(user *models.User) error
	DeleteByID(id uint) error
}

// GormUserRepository implementa UserRepository utilizando GORM y PostgreSQL.
type GormUserRepository struct {
	db *gorm.DB
}

// NewUserRepository crea una nueva instancia del repositorio.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{
		db: db,
	}
}

// Create registra un nuevo usuario.
func (r *GormUserRepository) Create(user *models.User) error {
	var count int64

	err := r.db.Model(&models.User{}).
		Where("LOWER(email) = LOWER(?)", user.Email).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return ErrEmailAlreadyExist
	}

	return r.db.Create(user).Error
}

// FindByEmail busca un usuario mediante su correo electrónico.
func (r *GormUserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.db.
		Where("LOWER(email) = LOWER(?)", email).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByID busca un usuario mediante su identificador.
func (r *GormUserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll devuelve todos los usuarios registrados.
func (r *GormUserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	err := r.db.Order("id ASC").Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

// Update actualiza la información de un usuario.
func (r *GormUserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// DeleteByID elimina un usuario mediante su identificador.
func (r *GormUserRepository) DeleteByID(id uint) error {
	result := r.db.Delete(&models.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}
