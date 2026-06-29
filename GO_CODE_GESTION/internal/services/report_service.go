package services

import "sistema-libros-electronicos/internal/models"

// CountActiveLoans cuenta los préstamos
// que siguen activos.
func CountActiveLoans(loans []models.Loan) int {

	total := 0

	for _, loan := range loans {

		if loan.Status == "active" {
			total++
		}

	}

	return total
}
