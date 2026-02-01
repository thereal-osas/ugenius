package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudyGroup struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	CampusID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"campus_id"`
	Campus      *Campus        `gorm:"foreignKey:CampusID" json:"campus,omitempty"`
	CreatedByID uuid.UUID      `gorm:"type:uuid;not null" json:"created_by_id"`
	CreatedBy   *User          `gorm:"foreignKey:CreatedByID" json:"created_by,omitempty"`
	MaxMembers  int            `gorm:"default:10" json:"max_members"`
	Subject     string         `json:"subject"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Members  []StudyGroupMember `gorm:"foreignKey:GroupID" json:"members,omitempty"`
	Sessions []StudySession     `gorm:"foreignKey:GroupID" json:"sessions,omitempty"`
}

func (s *StudyGroup) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

// MemberCount returns the current number of members
func (s *StudyGroup) MemberCount() int {
	return len(s.Members)
}

// IsFull returns true if the group has reached max capacity
func (s *StudyGroup) IsFull() bool {
	return len(s.Members) >= s.MaxMembers
}

type MemberRole string

const (
	MemberRoleLeader MemberRole = "leader"
	MemberRoleMember MemberRole = "member"
)

type StudyGroupMember struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	GroupID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"group_id"`
	Group     *StudyGroup    `gorm:"foreignKey:GroupID" json:"group,omitempty"`
	UserID    uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	User      *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Role      MemberRole     `gorm:"type:varchar(20);default:'member'" json:"role"`
	JoinedAt  time.Time      `gorm:"not null" json:"joined_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (m *StudyGroupMember) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	if m.JoinedAt.IsZero() {
		m.JoinedAt = time.Now()
	}
	return nil
}

type StudySession struct {
	ID              uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	GroupID         uuid.UUID      `gorm:"type:uuid;not null;index" json:"group_id"`
	Group           *StudyGroup    `gorm:"foreignKey:GroupID" json:"group,omitempty"`
	Title           string         `gorm:"not null" json:"title"`
	Description     string         `json:"description"`
	ScheduledAt     time.Time      `gorm:"not null" json:"scheduled_at"`
	DurationMinutes int            `gorm:"not null" json:"duration_minutes"`
	Location        string         `json:"location"` // Physical or virtual meeting link
	CreatedByID     uuid.UUID      `gorm:"type:uuid;not null" json:"created_by_id"`
	CreatedBy       *User          `gorm:"foreignKey:CreatedByID" json:"created_by,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (s *StudySession) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

// IsUpcoming returns true if the session hasn't happened yet
func (s *StudySession) IsUpcoming() bool {
	return time.Now().Before(s.ScheduledAt)
}

