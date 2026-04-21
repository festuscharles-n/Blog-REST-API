package models

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body" gorm:"not null"`
	Author    string    `json:"author" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
