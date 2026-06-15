package services

import (
	"errors"
	"time"

	"sistema-libros-electronicos/interno/modelos"
	"sistema-libros-electronicos/interno/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("biblioteca_virtual_secret")

type AuthService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(user *modelos.User) error {

	if user.Email == "" {
		return errors.New("correo requerido")
	}

	if len(user.Password) < 6 {
		return errors.New("la contraseña debe tener mínimo 6 caracteres")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.repo.Create(user)
}

func (s *AuthService) Login(email string, password string) (string, error) {

	user, err := s.repo.GetByEmail(email)

	if err != nil {
		return "", errors.New("usuario no encontrado")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("contraseña incorrecta")
	}

	token, err := GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

func GenerateToken(user *modelos.User) (string, error) {

	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(jwtKey)
}
