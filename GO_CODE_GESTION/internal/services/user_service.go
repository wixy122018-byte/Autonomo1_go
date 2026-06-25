package services

import (
	"errors"
	"net/mail"
	"strings"

	"sistema-libros-electronicos/internal/models"
	"sistema-libros-electronicos/internal/repositories"
)

// UserService contiene la lógica relacionada con la gestión de usuarios.
type UserService struct {
	repository repositories.UserRepository
}

// NewUserService crea una nueva instancia del servicio de usuarios.
func NewUserService(repository repositories.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

// ValidateEmail verifica que el correo no esté vacío y tenga un formato válido.
func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return errors.New("el correo electrónico es obligatorio")
	}

	parsedEmail, err := mail.ParseAddress(email)

	if err != nil || parsedEmail.Address != email {
		return errors.New("el formato del correo electrónico no es válido")
	}

	return nil
}

// ValidatePassword verifica que la contraseña cumpla la longitud mínima.
func ValidatePassword(password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("la contraseña es obligatoria")
	}

	if len(password) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}

	return nil
}

// GetUserByID obtiene un usuario mediante su identificador.
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("el identificador del usuario no es válido")
	}

	return s.repository.FindByID(id)
}

// GetAllUsers obtiene todos los usuarios registrados.
func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

// DeleteUser elimina un usuario mediante su identificador.
func (s *UserService) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("el identificador del usuario no es válido")
	}

	return s.repository.DeleteByID(id)
}
