package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScholarshipStatus string

const (
	ScholarshipStatusOpen      ScholarshipStatus = "open"
	ScholarshipStatusClosed    ScholarshipStatus = "closed"
	ScholarshipStatusReviewing ScholarshipStatus = "reviewing"
)

type ApplicationStatus string

const (
	ApplicationPending   ApplicationStatus = "pending"
	ApplicationReviewing ApplicationStatus = "reviewing"
	ApplicationApproved  ApplicationStatus = "approved"
	ApplicationRejected  ApplicationStatus = "rejected"
)

type Scholarship struct {
	ID                  uuid.UUID         `gorm:"type:uuid;primary_key" json:"id"`
	Title               string            `gorm:"not null" json:"title"`
	Description         string            `gorm:"not null" json:"description"`
	Provider            string            `json:"provider"` // Organization providing the scholarship
	Amount              float64           `json:"amount"`
	Currency            string            `gorm:"default:'NGN'" json:"currency"`
	Status              ScholarshipStatus `gorm:"type:varchar(20);default:'open'" json:"status"`
	
	// Eligibility Requirements
	MinReadingHours     int     `json:"min_reading_hours"`     // Minimum approved reading hours
	MinWeeklyHours      float64 `json:"min_weekly_hours"`      // Minimum weekly average
	MinAchievements     int     `json:"min_achievements"`      // Minimum badges earned
	RequiredBadges      string  `gorm:"type:jsonb" json:"required_badges"` // JSON array of required badge types
	
	ApplicationDeadline time.Time `gorm:"not null" json:"application_deadline"`
	StartDate           time.Time `json:"start_date"`
	MaxRecipients       int       `gorm:"default:1" json:"max_recipients"`
	
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Applications []ScholarshipApplication `gorm:"foreignKey:ScholarshipID" json:"applications,omitempty"`
}

func (s *Scholarship) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

// IsOpen returns true if applications are still being accepted
func (s *Scholarship) IsOpen() bool {
	return s.Status == ScholarshipStatusOpen && time.Now().Before(s.ApplicationDeadline)
}

type ScholarshipApplication struct {
	ID            uuid.UUID         `gorm:"type:uuid;primary_key" json:"id"`
	ScholarshipID uuid.UUID         `gorm:"type:uuid;not null;index" json:"scholarship_id"`
	Scholarship   *Scholarship      `gorm:"foreignKey:ScholarshipID" json:"scholarship,omitempty"`
	UserID        uuid.UUID         `gorm:"type:uuid;not null;index" json:"user_id"`
	User          *User             `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Status        ApplicationStatus `gorm:"type:varchar(20);default:'pending'" json:"status"`
	
	// Application data
	Statement     string `json:"statement"`          // Personal statement
	Documents     string `gorm:"type:jsonb" json:"documents"` // JSON array of document URLs
	
	// Stats at time of application (snapshot)
	TotalReadingHours float64 `json:"total_reading_hours"`
	WeeklyAverage     float64 `json:"weekly_average"`
	AchievementCount  int     `json:"achievement_count"`
	
	// Review
	ReviewedByID  *uuid.UUID `gorm:"type:uuid" json:"reviewed_by_id"`
	ReviewedBy    *User      `gorm:"foreignKey:ReviewedByID" json:"reviewed_by,omitempty"`
	ReviewNotes   string     `json:"review_notes"`
	ReviewedAt    *time.Time `json:"reviewed_at"`
	
	AppliedAt time.Time      `gorm:"not null" json:"applied_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (a *ScholarshipApplication) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	if a.AppliedAt.IsZero() {
		a.AppliedAt = time.Now()
	}
	return nil
}

