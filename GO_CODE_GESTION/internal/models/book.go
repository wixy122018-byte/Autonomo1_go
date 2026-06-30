package models

import "time"

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Category    string    `json:"category"`
	Editorial   string    `json:"editorial"`
	Year        int       `json:"year"`
	Description string    `json:"description"`
	Format      string    `json:"format"`
	ISBN        string    `json:"isbn" gorm:"unique"`
	Available   bool      `json:"available"`
	FilePath    string    `json:"file_path"`
	CreatedAt   time.Time `json:"created_at"`
}