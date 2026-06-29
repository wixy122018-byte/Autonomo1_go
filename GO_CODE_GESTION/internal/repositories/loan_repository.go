package repositories

import "sistema-libros-electronicos/internal/models"

// LoanRepository define las operaciones básicas
// para gestionar préstamos.
type LoanRepository interface {
	Create(*models.Loan) error
	GetAll() ([]models.Loan, error)
	GetByID(uint) (*models.Loan, error)
	Update(*models.Loan) error
}
