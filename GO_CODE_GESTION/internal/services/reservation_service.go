package services

import (
	"errors"

	"sistema-libros-electronicos/internal/models"
)

// CreateReservation valida la información
// antes de registrar una reserva.
func CreateReservation(reservation models.Reservation) error {

	if reservation.UserID == 0 {
		return errors.New("usuario inválido")
	}

	if reservation.BookID == 0 {
		return errors.New("libro inválido")
	}

	if reservation.Status == "" {
		return errors.New("la reserva debe tener un estado")
	}

	return nil
}
