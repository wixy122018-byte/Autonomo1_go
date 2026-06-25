package services

import (
	"errors"
	"strings"
	"time"

	"sistema-libros-electronicos/internal/models"
	"sistema-libros-electronicos/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

var userRepository = repositories.NewUserRepository()

// validateUserFields verifica que los campos obligatorios
// del usuario no estén vacíos.
func validateUserFields(name string, email string, password string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("el nombre es obligatorio")
	}

	if strings.TrimSpace(email) == "" {
		return errors.New("el correo es obligatorio")
	}

	if strings.TrimSpace(password) == "" {
		return errors.New("la contraseña es obligatoria")
	}

	return nil
}

// hashPassword cifra la contraseña antes de guardarla.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CreateUser registra un nuevo usuario aplicando validaciones,
// evitando correos duplicados y cifrando la contraseña.
func CreateUser(name string, email string, password string, role string) (*models.User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)
	role = strings.TrimSpace(role)

	if err := validateUserFields(name, email, password); err != nil {
		return nil, err
	}

	exists, err := userRepository.EmailExists(email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("el correo ya está registrado")
	}

	if role == "" {
		role = "lector"
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		Role:      role,
		CreatedAt: time.Now(),
	}

	err = userRepository.Create(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsers() ([]models.User, error) {
	return userRepository.GetAll()
}

func GetUserByID(id uint) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("id de usuario inválido")
	}

	return userRepository.GetByID(id)
}

// UpdateUser actualiza los datos de un usuario.
// Si se envía una nueva contraseña, se cifra antes de guardarla.
func UpdateUser(id uint, name string, email string, password string, role string) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("id de usuario inválido")
	}

	currentUser, err := userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	email = strings.TrimSpace(email)

	if email != "" && email != currentUser.Email {
		exists, err := userRepository.EmailExists(email)
		if err != nil {
			return nil, err
		}

		if exists {
			return nil, errors.New("el correo ya está registrado por otro usuario")
		}
	}

	user := models.User{
		Name:  strings.TrimSpace(name),
		Email: email,
		Role:  strings.TrimSpace(role),
	}

	if strings.TrimSpace(password) != "" {
		hashedPassword, err := hashPassword(password)
		if err != nil {
			return nil, err
		}

		user.Password = hashedPassword
	}

	return userRepository.Update(id, &user)
}

func DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("id de usuario inválido")
	}

	return userRepository.Delete(id)
}

// Login valida las credenciales del usuario.
func Login(email string, password string) (*models.User, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return nil, errors.New("el correo es obligatorio")
	}

	if strings.TrimSpace(password) == "" {
		return nil, errors.New("la contraseña es obligatoria")
	}

	user, err := userRepository.GetByEmail(email)
	if err != nil {
		return nil, errors.New("credenciales incorrectas")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("credenciales incorrectas")
	}

	return user, nil
}
