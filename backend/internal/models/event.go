package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventType string

const (
	EventTypeWorkshop  EventType = "workshop"
	EventTypeSeminar   EventType = "seminar"
	EventTypeWebinar   EventType = "webinar"
	EventTypeMeetup    EventType = "meetup"
	EventTypeContest   EventType = "contest"
	EventTypeOther     EventType = "other"
)

type EventStatus string

const (
	EventStatusUpcoming  EventStatus = "upcoming"
	EventStatusOngoing   EventStatus = "ongoing"
	EventStatusCompleted EventStatus = "completed"
	EventStatusCancelled EventStatus = "cancelled"
)

type Event struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Type        EventType      `gorm:"type:varchar(50);not null;default:'other'" json:"type"`
	Status      EventStatus    `gorm:"type:varchar(50);not null;default:'upcoming'" json:"status"`
	StartTime   time.Time      `gorm:"not null" json:"start_time"`
	EndTime     time.Time      `gorm:"not null" json:"end_time"`
	Location    string         `json:"location"` // Physical location or "Online"
	VirtualLink string         `json:"virtual_link,omitempty"` // Zoom/Meet link for virtual events
	ImageURL    string         `json:"image_url,omitempty"`
	CampusID    *uuid.UUID     `gorm:"type:uuid" json:"campus_id,omitempty"` // nil = all campuses
	Campus      *Campus        `gorm:"foreignKey:CampusID" json:"campus,omitempty"`
	CreatedByID uuid.UUID      `gorm:"type:uuid;not null" json:"created_by_id"`
	CreatedBy   *User          `gorm:"foreignKey:CreatedByID" json:"created_by,omitempty"`
	MaxAttendees int           `json:"max_attendees,omitempty"` // 0 = unlimited
	IsFeatured  bool           `gorm:"default:false" json:"is_featured"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.New()
	}
	if e.Status == "" {
		e.Status = EventStatusUpcoming
	}
	if e.Type == "" {
		e.Type = EventTypeOther
	}
	return nil
}

// IsUpcoming returns true if the event hasn't started yet
func (e *Event) IsUpcoming() bool {
	return time.Now().Before(e.StartTime)
}

// IsOngoing returns true if the event is currently happening
func (e *Event) IsOngoing() bool {
	now := time.Now()
	return now.After(e.StartTime) && now.Before(e.EndTime)
}

// EventRegistration tracks user registrations for events
type EventRegistration struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	EventID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"event_id"`
	Event       *Event         `gorm:"foreignKey:EventID" json:"event,omitempty"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	User        *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	RegisteredAt time.Time     `json:"registered_at"`
	AttendedAt  *time.Time     `json:"attended_at,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (r *EventRegistration) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	if r.RegisteredAt.IsZero() {
		r.RegisteredAt = time.Now()
	}
	return nil
}

// EventFilter for querying events
type EventFilter struct {
	CampusID  *uuid.UUID
	Type      *EventType
	Status    *EventStatus
	StartDate *time.Time
	EndDate   *time.Time
	Featured  *bool
	Page      int
	PageSize  int
}

