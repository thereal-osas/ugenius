package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BadgeType string

const (
	// Reading milestones
	BadgeFirstSubmission  BadgeType = "first_submission"
	Badge10Hours          BadgeType = "10_hours"
	Badge50Hours          BadgeType = "50_hours"
	Badge100Hours         BadgeType = "100_hours"
	Badge500Hours         BadgeType = "500_hours"
	Badge1000Hours        BadgeType = "1000_hours"

	// Streak badges
	Badge7DayStreak       BadgeType = "7_day_streak"
	Badge30DayStreak      BadgeType = "30_day_streak"
	Badge100DayStreak     BadgeType = "100_day_streak"

	// Weekly performance
	BadgeWeeklyTop3       BadgeType = "weekly_top_3"
	BadgeWeeklyChampion   BadgeType = "weekly_champion"

	// Goal achievements
	BadgeGoalCrusher      BadgeType = "goal_crusher"      // Complete 5 goals
	BadgeGoalMaster       BadgeType = "goal_master"       // Complete 20 goals

	// Community
	BadgeStudyBuddy       BadgeType = "study_buddy"       // Join first study group
	BadgeGroupLeader      BadgeType = "group_leader"      // Create a study group

	// Special
	BadgeEarlyBird        BadgeType = "early_bird"        // Log reading before 6 AM
	BadgeNightOwl         BadgeType = "night_owl"         // Log reading after 10 PM
	BadgeConsistent       BadgeType = "consistent"        // Submit every day for a week
)

// BadgeInfo contains metadata about each badge type
type BadgeInfo struct {
	Type        BadgeType `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Points      int       `json:"points"`
}

var BadgeCatalog = map[BadgeType]BadgeInfo{
	BadgeFirstSubmission: {Type: BadgeFirstSubmission, Title: "First Steps", Description: "Submit your first reading hour", Icon: "🎯", Points: 10},
	Badge10Hours:         {Type: Badge10Hours, Title: "Getting Started", Description: "Complete 10 hours of reading", Icon: "📖", Points: 25},
	Badge50Hours:         {Type: Badge50Hours, Title: "Dedicated Reader", Description: "Complete 50 hours of reading", Icon: "📚", Points: 50},
	Badge100Hours:        {Type: Badge100Hours, Title: "Century Club", Description: "Complete 100 hours of reading", Icon: "🏆", Points: 100},
	Badge500Hours:        {Type: Badge500Hours, Title: "Scholar", Description: "Complete 500 hours of reading", Icon: "🎓", Points: 250},
	Badge1000Hours:       {Type: Badge1000Hours, Title: "Academic Legend", Description: "Complete 1000 hours of reading", Icon: "👑", Points: 500},
	Badge7DayStreak:      {Type: Badge7DayStreak, Title: "Week Warrior", Description: "Maintain a 7-day reading streak", Icon: "🔥", Points: 30},
	Badge30DayStreak:     {Type: Badge30DayStreak, Title: "Month Master", Description: "Maintain a 30-day reading streak", Icon: "⚡", Points: 100},
	Badge100DayStreak:    {Type: Badge100DayStreak, Title: "Unstoppable", Description: "Maintain a 100-day reading streak", Icon: "💎", Points: 300},
	BadgeWeeklyTop3:      {Type: BadgeWeeklyTop3, Title: "Top Performer", Description: "Finish in top 3 on weekly leaderboard", Icon: "🥉", Points: 50},
	BadgeWeeklyChampion:  {Type: BadgeWeeklyChampion, Title: "Weekly Champion", Description: "Top the weekly leaderboard", Icon: "🥇", Points: 100},
	BadgeGoalCrusher:     {Type: BadgeGoalCrusher, Title: "Goal Crusher", Description: "Complete 5 study goals", Icon: "💪", Points: 75},
	BadgeGoalMaster:      {Type: BadgeGoalMaster, Title: "Goal Master", Description: "Complete 20 study goals", Icon: "🎖️", Points: 200},
	BadgeStudyBuddy:      {Type: BadgeStudyBuddy, Title: "Study Buddy", Description: "Join your first study group", Icon: "🤝", Points: 20},
	BadgeGroupLeader:     {Type: BadgeGroupLeader, Title: "Group Leader", Description: "Create a study group", Icon: "👨‍🏫", Points: 40},
	BadgeEarlyBird:       {Type: BadgeEarlyBird, Title: "Early Bird", Description: "Log reading before 6 AM", Icon: "🌅", Points: 15},
	BadgeNightOwl:        {Type: BadgeNightOwl, Title: "Night Owl", Description: "Log reading after 10 PM", Icon: "🦉", Points: 15},
	BadgeConsistent:      {Type: BadgeConsistent, Title: "Consistent", Description: "Submit reading hours every day for a week", Icon: "📅", Points: 35},
}

type Achievement struct {
	ID          uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	User        *User          `gorm:"foreignKey:UserID" json:"user,omitempty"`
	BadgeType   BadgeType      `gorm:"type:varchar(50);not null" json:"badge_type"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Icon        string         `json:"icon"`
	Points      int            `gorm:"default:0" json:"points"`
	EarnedAt    time.Time      `gorm:"not null" json:"earned_at"`
	CreatedAt   time.Time      `json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (a *Achievement) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	if a.EarnedAt.IsZero() {
		a.EarnedAt = time.Now()
	}
	return nil
}

// GetBadgeInfo returns the badge metadata for this achievement
func (a *Achievement) GetBadgeInfo() BadgeInfo {
	if info, ok := BadgeCatalog[a.BadgeType]; ok {
		return info
	}
	return BadgeInfo{Type: a.BadgeType, Title: string(a.BadgeType)}
}

