package models

import "time"

type Loan struct {
	ID         uint       `json:"id" gorm:"primaryKey"`
	UserID     uint       `json:"user_id"`
	BookID     uint       `json:"book_id"`
	StartDate  time.Time  `json:"start_date"`
	DueDate    time.Time  `json:"due_date"`
	ReturnDate *time.Time `json:"return_date"`
	Status     string     `json:"status"`
}

// ReturnBook registra la devolución del libro
// y actualiza automáticamente su estado.
func (l *Loan) ReturnBook() {
	now := time.Now()
	l.ReturnDate = &now
	l.Status = "returned"
}
