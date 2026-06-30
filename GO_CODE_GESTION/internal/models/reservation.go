package models

import "time"

type Reservation struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id"`
	BookID          uint      `json:"book_id"`
	ReservationDate time.Time `json:"reservation_date"`
	Status          string    `json:"status"`
}