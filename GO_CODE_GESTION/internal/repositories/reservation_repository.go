package repositories

import "sistema-libros-electronicos/internal/models"

// ReservationRepository define las operaciones
// para gestionar reservas.
type ReservationRepository interface {
	Create(*models.Reservation) error
	GetAll() ([]models.Reservation, error)
	GetByID(uint) (*models.Reservation, error)
	Update(*models.Reservation) error
}
