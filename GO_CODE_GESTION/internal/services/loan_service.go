package services

import (
	"errors"

	"sistema-libros-electronicos/internal/models"
)

// CreateLoan valida los datos antes de registrar un préstamo.
func CreateLoan(loan models.Loan) error {

	// Validar usuario
	if loan.UserID == 0 {
		return errors.New("usuario inválido")
	}

	// Validar libro
	if loan.BookID == 0 {
		return errors.New("libro inválido")
	}

	// Validar estado
	if loan.Status == "" {
		return errors.New("el préstamo debe tener un estado")
	}

	return nil
}
