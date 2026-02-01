package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationType string

const (
	NotificationSubmissionApproved  NotificationType = "submission_approved"
	NotificationSubmissionRejected  NotificationType = "submission_rejected"
	NotificationAchievementEarned   NotificationType = "achievement_earned"
	NotificationGoalCompleted       NotificationType = "goal_completed"
	NotificationGoalReminder        NotificationType = "goal_reminder"
	NotificationStudyGroupInvite    NotificationType = "study_group_invite"
	NotificationStudySessionReminder NotificationType = "study_session_reminder"
	NotificationWeeklyReport        NotificationType = "weekly_report"
	NotificationStreakWarning       NotificationType = "streak_warning"
	NotificationLeaderboardUpdate   NotificationType = "leaderboard_update"
	NotificationScholarshipUpdate   NotificationType = "scholarship_update"
	NotificationSystem              NotificationType = "system"
)

type Notification struct {
	ID         uuid.UUID        `gorm:"type:uuid;primary_key" json:"id"`
	UserID     uuid.UUID        `gorm:"type:uuid;not null;index" json:"user_id"`
	User       *User            `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Type       NotificationType `gorm:"type:varchar(50);not null" json:"type"`
	Title      string           `gorm:"not null" json:"title"`
	Message    string           `gorm:"not null" json:"message"`
	Data       string           `gorm:"type:jsonb" json:"data,omitempty"` // JSON metadata
	IsRead     bool             `gorm:"default:false" json:"is_read"`
	ReadAt     *time.Time       `json:"read_at,omitempty"`
	ActionURL  string           `json:"action_url,omitempty"`
	CreatedAt  time.Time        `json:"created_at"`
	DeletedAt  gorm.DeletedAt   `gorm:"index" json:"-"`
}

func (n *Notification) BeforeCreate(tx *gorm.DB) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return nil
}

// MarkAsRead marks the notification as read
func (n *Notification) MarkAsRead() {
	n.IsRead = true
	now := time.Now()
	n.ReadAt = &now
}

// NotificationPreferences stores user notification settings
type NotificationPreferences struct {
	ID                    uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID                uuid.UUID `gorm:"type:uuid;not null;uniqueIndex" json:"user_id"`
	EmailSubmissionStatus bool      `gorm:"default:true" json:"email_submission_status"`
	EmailAchievements     bool      `gorm:"default:true" json:"email_achievements"`
	EmailWeeklyReport     bool      `gorm:"default:true" json:"email_weekly_report"`
	EmailStudyReminders   bool      `gorm:"default:true" json:"email_study_reminders"`
	PushEnabled           bool      `gorm:"default:true" json:"push_enabled"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func (p *NotificationPreferences) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

