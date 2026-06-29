package models

import (
	"errors"
	"strings"
	"time"
)

const (
	RoleAdministrador = "ADMINISTRADOR"
	RoleBibliotecario = "BIBLIOTECARIO"
	RoleLector        = "LECTOR"
)

// ValidRoles utiliza un map para validar los roles permitidos.
var ValidRoles = map[string]bool{
	RoleAdministrador: true,
	RoleBibliotecario: true,
	RoleLector:        true,
}

// User representa un usuario registrado en el sistema.
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null;default:LECTOR"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SetRole valida y asigna un rol al usuario.
func (u *User) SetRole(role string) error {
	role = strings.ToUpper(strings.TrimSpace(role))

	if !ValidRoles[role] {
		return errors.New("rol no válido")
	}

	u.Role = role
	return nil
}

// GetRole devuelve el rol actual del usuario.
func (u *User) GetRole() string {
	return u.Role
}

// HasRole verifica si el usuario posee un rol específico.
func (u *User) HasRole(role string) bool {
	return u.Role == strings.ToUpper(strings.TrimSpace(role))
}
