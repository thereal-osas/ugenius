package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campus struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Code      string         `gorm:"uniqueIndex;not null" json:"code"`
	Location  string         `json:"location"`
	State     string         `json:"state"`
	Country   string         `gorm:"default:'Nigeria'" json:"country"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Users       []User       `gorm:"foreignKey:CampusID" json:"users,omitempty"`
	StudyGroups []StudyGroup `gorm:"foreignKey:CampusID" json:"study_groups,omitempty"`
}

func (c *Campus) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

// CampusStats represents aggregated statistics for a campus
type CampusStats struct {
	TotalStudents       int64   `json:"total_students"`
	ActiveStudents      int64   `json:"active_students"`
	TotalReadingHours   float64 `json:"total_reading_hours"`
	WeeklyReadingHours  float64 `json:"weekly_reading_hours"`
	TotalSubmissions    int64   `json:"total_submissions"`
	PendingSubmissions  int64   `json:"pending_submissions"`
	ApprovedSubmissions int64   `json:"approved_submissions"`
	ActiveStudyGroups   int64   `json:"active_study_groups"`
	TopPerformers       []User  `json:"top_performers,omitempty"`
}

