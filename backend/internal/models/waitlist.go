package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// WaitlistRegistration represents a user who has registered interest in U-Genius
type WaitlistRegistration struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	FullName  string         `gorm:"not null" json:"full_name"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Phone     string         `json:"phone,omitempty"`
	CampusID  *uuid.UUID     `gorm:"type:uuid" json:"campus_id,omitempty"`
	Level     string         `json:"level,omitempty"`
	Faculty   string         `json:"faculty,omitempty"`
	Department string        `json:"department,omitempty"`
	Address   string         `json:"address,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Campus *Campus `gorm:"foreignKey:CampusID" json:"campus,omitempty"`
}

func (w *WaitlistRegistration) BeforeCreate(tx *gorm.DB) error {
	if w.ID == uuid.Nil {
		w.ID = uuid.New()
	}
	return nil
}

// WaitlistRegistrationRequest represents the API request to register for the waitlist
type WaitlistRegistrationRequest struct {
	FullName   string `json:"full_name" binding:"required,min=2,max=100"`
	Email      string `json:"email" binding:"required,email"`
	Phone      string `json:"phone,omitempty"`
	CampusID   string `json:"campus_id,omitempty"`
	Level      string `json:"level,omitempty"`
	Faculty    string `json:"faculty,omitempty"`
	Department string `json:"department,omitempty"`
	Address    string `json:"address,omitempty"`
}

