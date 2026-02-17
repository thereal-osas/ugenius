package models

import (
	"time"
)

type Gallery struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	ImageURL  string    `json:"image_url" gorm:"not null"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Gallery) TableName() string {
	return "galleries"
}
