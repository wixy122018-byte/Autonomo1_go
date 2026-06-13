package services

import (
	"testing"

	"sistema-libros-electronicos/internal/models"
)

func TestHashAndCheckPassword(t *testing.T) {
	hash, err := HashPassword("secreto123")
	if err != nil {
		t.Fatal(err)
	}

	if hash == "secreto123" {
		t.Fatal("la contraseña no debe guardarse en texto plano")
	}
	if !CheckPassword("secreto123", hash) {
		t.Fatal("la contraseña correcta no fue validada")
	}
	if CheckPassword("incorrecta", hash) {
		t.Fatal("una contraseña incorrecta no debe ser aceptada")
	}
}

func TestGenerateAndVerifyJWT(t *testing.T) {
	user := models.User{
		ID:    7,
		Name:  "Estudiante",
		Email: "estudiante@example.com",
		Role:  "user",
	}

	token, err := GenerateJWT(user)
	if err != nil {
		t.Fatal(err)
	}

	claims, err := VerifyJWT(token)
	if err != nil {
		t.Fatal(err)
	}

	if claims.UserID != user.ID || claims.Email != user.Email || claims.Role != user.Role {
		t.Fatalf("claims incorrectos: %#v", claims)
	}
}
