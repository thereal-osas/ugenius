package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GoalType string
type GoalPeriod string

const (
	GoalTypeReadingHours GoalType = "reading_hours"
	GoalTypeSubmissions  GoalType = "submissions"
	GoalTypeStreak       GoalType = "streak"

	GoalPeriodDaily   GoalPeriod = "daily"
	GoalPeriodWeekly  GoalPeriod = "weekly"
	GoalPeriodMonthly GoalPeriod = "monthly"
	GoalPeriodCustom  GoalPeriod = "custom"
)

type StudyGoal struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID        uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	User          *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Title         string         `gorm:"not null" json:"title"`
	Description   string         `json:"description"`
	GoalType      GoalType       `gorm:"type:varchar(30);not null" json:"goal_type"`
	TargetValue   int            `gorm:"not null" json:"target_value"` // e.g., target hours or submissions
	CurrentValue  int            `gorm:"default:0" json:"current_value"`
	Period        GoalPeriod     `gorm:"type:varchar(20);not null" json:"period"`
	StartDate     time.Time      `gorm:"type:date;not null" json:"start_date"`
	EndDate       time.Time      `gorm:"type:date" json:"end_date"`
	IsCompleted   bool           `gorm:"default:false" json:"is_completed"`
	CompletedAt   *time.Time     `json:"completed_at,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *StudyGoal) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

// Progress returns the percentage progress towards the goal
func (s *StudyGoal) Progress() float64 {
	if s.TargetValue == 0 {
		return 0
	}
	progress := float64(s.CurrentValue) / float64(s.TargetValue) * 100
	if progress > 100 {
		return 100
	}
	return progress
}

// IsActive returns true if the goal is within its date range and not completed
func (s *StudyGoal) IsActive() bool {
	now := time.Now()
	return !s.IsCompleted && now.After(s.StartDate) && (s.EndDate.IsZero() || now.Before(s.EndDate))
}

