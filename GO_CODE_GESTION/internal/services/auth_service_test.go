package services

import (
	"errors"
	"strings"
	"testing"

	"sistema-libros-electronicos/internal/models"
	"sistema-libros-electronicos/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

// FakeUserRepository simula una base de datos para realizar pruebas
// sin conectarse directamente a PostgreSQL.
type FakeUserRepository struct {
	users  []models.User
	nextID uint
}

// NewFakeUserRepository crea un repositorio de prueba vacío.
func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{
		users:  []models.User{},
		nextID: 1,
	}
}

// Create guarda un usuario en memoria.
func (r *FakeUserRepository) Create(user *models.User) error {
	for _, existingUser := range r.users {
		if strings.EqualFold(existingUser.Email, user.Email) {
			return repositories.ErrEmailAlreadyExist
		}
	}

	user.ID = r.nextID
	r.nextID++

	r.users = append(r.users, *user)

	return nil
}

// FindByEmail busca un usuario por correo.
func (r *FakeUserRepository) FindByEmail(email string) (*models.User, error) {
	for i := range r.users {
		if strings.EqualFold(r.users[i].Email, email) {
			return &r.users[i], nil
		}
	}

	return nil, repositories.ErrUserNotFound
}

// FindByID busca un usuario por ID.
func (r *FakeUserRepository) FindByID(id uint) (*models.User, error) {
	for i := range r.users {
		if r.users[i].ID == id {
			return &r.users[i], nil
		}
	}

	return nil, repositories.ErrUserNotFound
}

// FindAll devuelve todos los usuarios simulados.
func (r *FakeUserRepository) FindAll() ([]models.User, error) {
	return r.users, nil
}

// Update actualiza un usuario simulado.
func (r *FakeUserRepository) Update(user *models.User) error {
	for i := range r.users {
		if r.users[i].ID == user.ID {
			r.users[i] = *user
			return nil
		}
	}

	return repositories.ErrUserNotFound
}

// TestRegisterUser verifica que el usuario se registre
// y que la contraseña quede cifrada con bcrypt.
func TestRegisterUser(t *testing.T) {
	repository := NewFakeUserRepository()

	authService := NewAuthService(
		repository,
		"clave-secreta-pruebas",
	)

	user := &models.User{
		Name:     "Michael",
		Email:    "michael@example.com",
		Password: "123456",
	}

	registeredUser, err := authService.Register(user)

	if err != nil {
		t.Fatalf("se esperaba un registro correcto, pero ocurrió: %v", err)
	}

	if registeredUser.ID == 0 {
		t.Error("el usuario debería tener un ID")
	}

	if registeredUser.Role != models.RoleLector {
		t.Errorf(
			"se esperaba el rol %s, pero se obtuvo %s",
			models.RoleLector,
			registeredUser.Role,
		)
	}

	if registeredUser.Password == "123456" {
		t.Error("la contraseña no debería guardarse en texto plano")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(registeredUser.Password),
		[]byte("123456"),
	)

	if err != nil {
		t.Error("el hash generado no corresponde a la contraseña")
	}
}

// TestLoginUser verifica el inicio de sesión y la generación del JWT.
func TestLoginUser(t *testing.T) {
	repository := NewFakeUserRepository()

	authService := NewAuthService(
		repository,
		"clave-secreta-pruebas",
	)

	user := &models.User{
		Name:     "Michael",
		Email:    "michael@example.com",
		Password: "123456",
	}

	_, err := authService.Register(user)

	if err != nil {
		t.Fatalf("no se pudo registrar el usuario: %v", err)
	}

	token, authenticatedUser, err := authService.Login(
		"michael@example.com",
		"123456",
	)

	if err != nil {
		t.Fatalf("no se pudo iniciar sesión: %v", err)
	}

	if token == "" {
		t.Error("se esperaba un token JWT")
	}

	if authenticatedUser.Email != "michael@example.com" {
		t.Error("el correo del usuario autenticado no coincide")
	}

	claims, err := authService.ValidateToken(token)

	if err != nil {
		t.Fatalf("el token generado debería ser válido: %v", err)
	}

	if claims.UserID != authenticatedUser.ID {
		t.Error("el ID almacenado en el token no coincide")
	}
}

// TestLoginIncorrectPassword verifica el manejo de una contraseña incorrecta.
func TestLoginIncorrectPassword(t *testing.T) {
	repository := NewFakeUserRepository()

	authService := NewAuthService(
		repository,
		"clave-secreta-pruebas",
	)

	user := &models.User{
		Name:     "Michael",
		Email:    "michael@example.com",
		Password: "123456",
	}

	_, err := authService.Register(user)

	if err != nil {
		t.Fatalf("no se pudo registrar el usuario: %v", err)
	}

	_, _, err = authService.Login(
		"michael@example.com",
		"contraseña-incorrecta",
	)

	if !errors.Is(err, ErrInvalidCredentials) {
		t.Errorf(
			"se esperaba ErrInvalidCredentials, pero se obtuvo: %v",
			err,
		)
	}
}

// TestInvalidEmail verifica la validación del correo.
func TestInvalidEmail(t *testing.T) {
	err := ValidateEmail("correo-invalido")

	if err == nil {
		t.Error("se esperaba un error para un correo inválido")
	}
}

// TestShortPassword verifica la longitud mínima de la contraseña.
func TestShortPassword(t *testing.T) {
	err := ValidatePassword("123")

	if err == nil {
		t.Error("se esperaba un error para una contraseña corta")
	}
}
