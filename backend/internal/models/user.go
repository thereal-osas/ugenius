package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleStudent     Role = "student"
	RoleCampusAdmin Role = "campus_admin"
	RoleSuperAdmin  Role = "super_admin"
)

type User struct {
	ID              uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Email           string         `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash    string         `gorm:"not null" json:"-"`
	FirstName       string         `gorm:"not null" json:"first_name"`
	LastName        string         `gorm:"not null" json:"last_name"`
	Phone           string         `json:"phone"`
	Role            Role           `gorm:"type:varchar(20);not null;default:'student'" json:"role"`
	CampusID        *uuid.UUID     `gorm:"type:uuid" json:"campus_id"`
	Campus          *Campus        `gorm:"foreignKey:CampusID" json:"campus,omitempty"`
	Institution     string         `json:"institution"`
	Department      string         `json:"department"`
	Level           string         `json:"level"`
	EmailVerified   bool           `gorm:"default:false" json:"email_verified"`
	VerifyToken     string         `gorm:"index" json:"-"`
	ResetToken      string         `gorm:"index" json:"-"`
	ResetTokenExpiry *time.Time    `json:"-"`
	ProfilePicture  string         `json:"profile_picture"`
	Bio             string         `json:"bio"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	ReadingHours []ReadingHour  `gorm:"foreignKey:UserID" json:"reading_hours,omitempty"`
	StudyGoals   []StudyGoal    `gorm:"foreignKey:UserID" json:"study_goals,omitempty"`
	Achievements []Achievement  `gorm:"foreignKey:UserID" json:"achievements,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// UserStats represents aggregated statistics for a user
type UserStats struct {
	TotalReadingHours    float64 `json:"total_reading_hours"`
	WeeklyReadingHours   float64 `json:"weekly_reading_hours"`
	MonthlyReadingHours  float64 `json:"monthly_reading_hours"`
	TotalSubmissions     int64   `json:"total_submissions"`
	ApprovedSubmissions  int64   `json:"approved_submissions"`
	PendingSubmissions   int64   `json:"pending_submissions"`
	RejectedSubmissions  int64   `json:"rejected_submissions"`
	CurrentStreak        int     `json:"current_streak"`
	LongestStreak        int     `json:"longest_streak"`
	TotalAchievements    int64   `json:"total_achievements"`
	WeeklyRank           int     `json:"weekly_rank"`
	GoalsCompleted       int64   `json:"goals_completed"`
	GoalsInProgress      int64   `json:"goals_in_progress"`
}

// RefreshToken stores refresh tokens for JWT authentication
type RefreshToken struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Token     string         `gorm:"uniqueIndex;not null" json:"-"`
	ExpiresAt time.Time      `gorm:"not null" json:"expires_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (r *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

