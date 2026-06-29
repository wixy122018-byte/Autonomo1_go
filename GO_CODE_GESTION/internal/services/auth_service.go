package services

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"sistema-libros-electronicos/internal/models"
	"sistema-libros-electronicos/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("correo o contraseña incorrectos")
	ErrJWTSecretRequired  = errors.New("la clave secreta JWT no está configurada")
	ErrInvalidToken       = errors.New("token inválido o expirado")
)

// UserClaims representa la información almacenada dentro del token JWT.
type UserClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// AuthService contiene la lógica de registro, login y autenticación.
type AuthService struct {
	repository    repositories.UserRepository
	jwtSecret     []byte
	tokenDuration time.Duration
}

// NewAuthService crea una nueva instancia del servicio de autenticación.
func NewAuthService(
	repository repositories.UserRepository,
	jwtSecret string,
) *AuthService {
	return &AuthService{
		repository:    repository,
		jwtSecret:     []byte(strings.TrimSpace(jwtSecret)),
		tokenDuration: 24 * time.Hour,
	}
}

// Register valida los datos, cifra la contraseña y registra al usuario.
func (s *AuthService) Register(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("los datos del usuario son obligatorios")
	}

	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

	if user.Name == "" {
		return nil, errors.New("el nombre es obligatorio")
	}

	if err := ValidateEmail(user.Email); err != nil {
		return nil, err
	}

	if err := ValidatePassword(user.Password); err != nil {
		return nil, err
	}

	// Si el usuario no envía un rol, se asigna LECTOR por defecto.
	if strings.TrimSpace(user.Role) == "" {
		if err := user.SetRole(models.RoleLector); err != nil {
			return nil, err
		}
	} else {
		if err := user.SetRole(user.Role); err != nil {
			return nil, err
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, errors.New("no se pudo cifrar la contraseña")
	}

	// La contraseña en texto plano se reemplaza por su hash.
	user.Password = string(hashedPassword)

	if err := s.repository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login valida las credenciales y genera un token JWT.
func (s *AuthService) Login(
	email string,
	password string,
) (string, *models.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	if err := ValidateEmail(email); err != nil {
		return "", nil, err
	}

	if strings.TrimSpace(password) == "" {
		return "", nil, errors.New("la contraseña es obligatoria")
	}

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return "", nil, ErrInvalidCredentials
		}

		return "", nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err != nil {
		return "", nil, ErrInvalidCredentials
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

// GenerateToken genera un JWT con los datos principales del usuario.
func (s *AuthService) GenerateToken(user *models.User) (string, error) {
	if len(s.jwtSecret) == 0 {
		return "", ErrJWTSecretRequired
	}

	now := time.Now()

	claims := UserClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "sistema-libros-electronicos",
			Subject:   strconv.FormatUint(uint64(user.ID), 10),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.tokenDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", errors.New("no se pudo generar el token")
	}

	return signedToken, nil
}

// ValidateToken verifica la firma, el método y la vigencia del token JWT.
func (s *AuthService) ValidateToken(tokenString string) (*UserClaims, error) {
	if len(s.jwtSecret) == 0 {
		return nil, ErrJWTSecretRequired
	}

	if strings.TrimSpace(tokenString) == "" {
		return nil, ErrInvalidToken
	}

	claims := &UserClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, ErrInvalidToken
			}

			return s.jwtSecret, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuer("sistema-libros-electronicos"),
	)

	if err != nil || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
