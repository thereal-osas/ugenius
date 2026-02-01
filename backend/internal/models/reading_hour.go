package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubmissionStatus string

const (
	StatusPending  SubmissionStatus = "pending"
	StatusApproved SubmissionStatus = "approved"
	StatusRejected SubmissionStatus = "rejected"
)

type ReadingHour struct {
	ID              uuid.UUID        `gorm:"type:uuid;primary_key" json:"id"`
	UserID          uuid.UUID        `gorm:"type:uuid;not null;index" json:"user_id"`
	User            *User            `gorm:"foreignKey:UserID" json:"user,omitempty"`
	ReadingDate     time.Time        `gorm:"type:date;not null" json:"reading_date"`
	DurationMinutes int              `gorm:"not null" json:"duration_minutes"`
	Subject         string           `gorm:"not null" json:"subject"`
	Topic           string           `json:"topic"`
	Description     string           `json:"description"`
	Status          SubmissionStatus `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DeletedAt       gorm.DeletedAt   `gorm:"index" json:"-"`

	// Review relationship
	Review *AdminReview `gorm:"foreignKey:ReadingHourID" json:"review,omitempty"`
}

func (r *ReadingHour) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	if r.Status == "" {
		r.Status = StatusPending
	}
	return nil
}

// DurationHours returns the duration in hours (float)
func (r *ReadingHour) DurationHours() float64 {
	return float64(r.DurationMinutes) / 60.0
}

type AdminReview struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	ReadingHourID uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex" json:"reading_hour_id"`
	AdminID       uuid.UUID      `gorm:"type:uuid;not null" json:"admin_id"`
	Admin         *User          `gorm:"foreignKey:AdminID" json:"admin,omitempty"`
	Decision      SubmissionStatus `gorm:"type:varchar(20);not null" json:"decision"`
	Feedback      string         `json:"feedback"`
	ReviewedAt    time.Time      `gorm:"not null" json:"reviewed_at"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (a *AdminReview) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	if a.ReviewedAt.IsZero() {
		a.ReviewedAt = time.Now()
	}
	return nil
}

// ReadingHourFilter for querying reading hours
type ReadingHourFilter struct {
	UserID    *uuid.UUID
	CampusID  *uuid.UUID
	Status    *SubmissionStatus
	StartDate *time.Time
	EndDate   *time.Time
	Subject   string
	Page      int
	PageSize  int
}

// LeaderboardEntry represents a single entry in the leaderboard
type LeaderboardEntry struct {
	Rank             int       `json:"rank"`
	UserID           uuid.UUID `json:"user_id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	ProfilePicture   string    `json:"profile_picture"`
	TotalMinutes     int       `json:"total_minutes"`
	TotalHours       float64   `json:"total_hours"`
	SubmissionCount  int       `json:"submission_count"`
}

