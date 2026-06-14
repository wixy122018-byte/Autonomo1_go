package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"sistema-libros-electronicos/internal/config"
	"sistema-libros-electronicos/internal/models"
	"sistema-libros-electronicos/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("credenciales invalidas")
	ErrEmailAlreadyExists = errors.New("el correo ya esta registrado")
	ErrInvalidToken       = errors.New("token invalido")
)

type AuthService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

type TokenClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Exp    int64  `json:"exp"`
}

func (s *AuthService) Register(input RegisterInput) (AuthResponse, error) {
	name := strings.TrimSpace(input.Name)
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)
	role := strings.TrimSpace(input.Role)

	if name == "" || email == "" || password == "" {
		return AuthResponse{}, errors.New("nombre, correo y contraseña son obligatorios")
	}
	if len(password) < 6 {
		return AuthResponse{}, errors.New("la contraseña debe tener al menos 6 caracteres")
	}
	if role == "" {
		role = "user"
	}

	if _, err := s.userRepo.FindByEmail(email); err == nil {
		return AuthResponse{}, ErrEmailAlreadyExists
	} else if !errors.Is(err, repositories.ErrUserNotFound) {
		return AuthResponse{}, err
	}

	hash, err := HashPassword(password)
	if err != nil {
		return AuthResponse{}, err
	}

	user, err := s.userRepo.Create(models.User{
		Name:     name,
		Email:    email,
		Password: hash,
		Role:     role,
	})
	if err != nil {
		return AuthResponse{}, err
	}

	token, err := GenerateJWT(user)
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{Token: token, User: user}, nil
}

func (s *AuthService) Login(input LoginInput) (AuthResponse, error) {
	email := strings.ToLower(strings.TrimSpace(input.Email))
	password := strings.TrimSpace(input.Password)

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return AuthResponse{}, ErrInvalidCredentials
	}

	if !CheckPassword(password, user.Password) {
		return AuthResponse{}, ErrInvalidCredentials
	}

	token, err := GenerateJWT(user)
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{Token: token, User: user}, nil
}

func (s *AuthService) FindUserFromToken(token string) (models.User, error) {
	claims, err := VerifyJWT(token)
	if err != nil {
		return models.User{}, err
	}

	return s.userRepo.FindByID(claims.UserID)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateJWT(user models.User) (string, error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	claims := TokenClaims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Exp:    time.Now().Add(24 * time.Hour).Unix(),
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	unsigned := encodeSegment(headerJSON) + "." + encodeSegment(claimsJSON)
	return unsigned + "." + signJWT(unsigned), nil
}

func VerifyJWT(token string) (TokenClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return TokenClaims{}, ErrInvalidToken
	}

	unsigned := parts[0] + "." + parts[1]
	expectedSignature := signJWT(unsigned)
	if !hmac.Equal([]byte(expectedSignature), []byte(parts[2])) {
		return TokenClaims{}, ErrInvalidToken
	}

	claimsJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return TokenClaims{}, ErrInvalidToken
	}

	var claims TokenClaims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return TokenClaims{}, ErrInvalidToken
	}
	if claims.Exp < time.Now().Unix() {
		return TokenClaims{}, errors.New("token expirado")
	}

	return claims, nil
}

func BearerToken(authHeader string) (string, error) {
	parts := strings.Fields(authHeader)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", ErrInvalidToken
	}

	return parts[1], nil
}

func encodeSegment(value []byte) string {
	return base64.RawURLEncoding.EncodeToString(value)
}

func signJWT(unsigned string) string {
	mac := hmac.New(sha256.New, []byte(config.GetJWTSecret()))
	mac.Write([]byte(unsigned))
	return encodeSegment(mac.Sum(nil))
}
