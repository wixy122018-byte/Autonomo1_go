package models

import "time"

type Download struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id"`
	BookID       uint      `json:"book_id"`
	DownloadDate time.Time `json:"download_date"`
}