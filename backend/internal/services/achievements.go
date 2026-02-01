package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/models"
	"gorm.io/gorm"
)

type AchievementsService struct {
	db           *gorm.DB
	emailService *EmailService
}

func NewAchievementsService(db *gorm.DB, emailService *EmailService) *AchievementsService {
	return &AchievementsService{
		db:           db,
		emailService: emailService,
	}
}

func (s *AchievementsService) GetUserAchievements(userID uuid.UUID) ([]models.Achievement, error) {
	var achievements []models.Achievement
	if err := s.db.Where("user_id = ?", userID).Order("earned_at DESC").Find(&achievements).Error; err != nil {
		return nil, err
	}
	return achievements, nil
}

func (s *AchievementsService) GetAllBadges(userID uuid.UUID) ([]map[string]interface{}, error) {
	// Get user's earned badges
	earnedBadges := make(map[models.BadgeType]bool)
	var achievements []models.Achievement
	s.db.Where("user_id = ?", userID).Find(&achievements)
	for _, a := range achievements {
		earnedBadges[a.BadgeType] = true
	}

	// Build response with all badges
	var badges []map[string]interface{}
	for badgeType, info := range models.BadgeCatalog {
		badges = append(badges, map[string]interface{}{
			"type":        badgeType,
			"title":       info.Title,
			"description": info.Description,
			"icon":        info.Icon,
			"points":      info.Points,
			"earned":      earnedBadges[badgeType],
		})
	}

	return badges, nil
}

func (s *AchievementsService) CheckAndAwardAchievements(userID uuid.UUID) ([]models.Achievement, error) {
	var newAchievements []models.Achievement

	// Get user stats
	var user models.User
	if err := s.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}

	// Get total approved reading hours
	var totalMinutes int64
	s.db.Model(&models.ReadingHour{}).
		Where("user_id = ? AND status = ?", userID, models.StatusApproved).
		Select("COALESCE(SUM(duration_minutes), 0)").
		Scan(&totalMinutes)
	totalHours := float64(totalMinutes) / 60.0

	// Get submission count
	var submissionCount int64
	s.db.Model(&models.ReadingHour{}).
		Where("user_id = ? AND status = ?", userID, models.StatusApproved).
		Count(&submissionCount)

	// Check each badge
	badgesToCheck := []struct {
		badgeType models.BadgeType
		condition bool
	}{
		{models.BadgeFirstSubmission, submissionCount >= 1},
		{models.BadgeConsistent, s.hasConsecutiveDays(userID, 7)},
		{models.Badge7DayStreak, s.hasConsecutiveDays(userID, 7)},
		{models.Badge30DayStreak, s.hasConsecutiveDays(userID, 30)},
		{models.Badge100DayStreak, s.hasConsecutiveDays(userID, 100)},
		{models.BadgeEarlyBird, s.hasEarlyMorningSubmissions(userID, 5)},
		{models.BadgeNightOwl, s.hasLateNightSubmissions(userID, 5)},
		{models.Badge10Hours, totalHours >= 10},
		{models.Badge50Hours, totalHours >= 50},
		{models.Badge100Hours, totalHours >= 100},
		{models.Badge500Hours, totalHours >= 500},
		{models.Badge1000Hours, totalHours >= 1000},
		{models.BadgeGoalCrusher, s.hasCompletedGoals(userID, 5)},
		{models.BadgeGoalMaster, s.hasCompletedGoals(userID, 20)},
		{models.BadgeStudyBuddy, s.isInStudyGroup(userID)},
		{models.BadgeGroupLeader, s.isGroupLeader(userID)},
		{models.BadgeWeeklyTop3, s.isInTopTen(userID)},
		{models.BadgeWeeklyChampion, s.isNumberOne(userID)},
	}

	for _, check := range badgesToCheck {
		if check.condition && !s.hasAchievement(userID, check.badgeType) {
			achievement := s.awardAchievement(userID, check.badgeType)
			if achievement != nil {
				newAchievements = append(newAchievements, *achievement)
			}
		}
	}

	return newAchievements, nil
}

func (s *AchievementsService) hasAchievement(userID uuid.UUID, badgeType models.BadgeType) bool {
	var count int64
	s.db.Model(&models.Achievement{}).Where("user_id = ? AND badge_type = ?", userID, badgeType).Count(&count)
	return count > 0
}

func (s *AchievementsService) awardAchievement(userID uuid.UUID, badgeType models.BadgeType) *models.Achievement {
	info, exists := models.BadgeCatalog[badgeType]
	if !exists {
		return nil
	}

	achievement := &models.Achievement{
		UserID:      userID,
		BadgeType:   badgeType,
		Title:       info.Title,
		Description: info.Description,
		Icon:        info.Icon,
		Points:      info.Points,
		EarnedAt:    time.Now(),
	}

	if err := s.db.Create(achievement).Error; err != nil {
		return nil
	}

	// Send email notification
	var user models.User
	if err := s.db.First(&user, "id = ?", userID).Error; err == nil {
		go s.emailService.SendAchievementEmail(user.Email, user.FirstName, info.Title, info.Description)
	}

	return achievement
}

func (s *AchievementsService) hasConsecutiveDays(userID uuid.UUID, days int) bool {
	// Check if user has submitted reading hours for consecutive days
	var dates []time.Time
	s.db.Model(&models.ReadingHour{}).
		Where("user_id = ? AND status = ?", userID, models.StatusApproved).
		Select("DISTINCT DATE(reading_date)").
		Order("reading_date DESC").
		Limit(days * 2).
		Pluck("reading_date", &dates)

	if len(dates) < days {
		return false
	}

	consecutive := 1
	for i := 1; i < len(dates); i++ {
		diff := dates[i-1].Sub(dates[i]).Hours() / 24
		if diff <= 1 {
			consecutive++
			if consecutive >= days {
				return true
			}
		} else {
			consecutive = 1
		}
	}
	return false
}

func (s *AchievementsService) hasEarlyMorningSubmissions(userID uuid.UUID, count int) bool {
	// Placeholder - would check for submissions before 6 AM
	return false
}

func (s *AchievementsService) hasLateNightSubmissions(userID uuid.UUID, count int) bool {
	// Placeholder - would check for submissions after 10 PM
	return false
}

func (s *AchievementsService) hasCompletedGoals(userID uuid.UUID, count int) bool {
	var completed int64
	s.db.Model(&models.StudyGoal{}).Where("user_id = ? AND is_completed = ?", userID, true).Count(&completed)
	return completed >= int64(count)
}

func (s *AchievementsService) isInStudyGroup(userID uuid.UUID) bool {
	var count int64
	s.db.Model(&models.StudyGroupMember{}).Where("user_id = ?", userID).Count(&count)
	return count > 0
}

func (s *AchievementsService) isGroupLeader(userID uuid.UUID) bool {
	var count int64
	s.db.Model(&models.StudyGroupMember{}).Where("user_id = ? AND role = ?", userID, models.MemberRoleLeader).Count(&count)
	return count > 0
}

func (s *AchievementsService) isInTopTen(userID uuid.UUID) bool {
	// Check weekly leaderboard
	return false // Placeholder
}

func (s *AchievementsService) isNumberOne(userID uuid.UUID) bool {
	// Check if user is #1 on weekly leaderboard
	return false // Placeholder
}

