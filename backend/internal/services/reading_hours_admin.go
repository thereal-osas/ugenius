package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/models"
)

type ReviewInput struct {
	Status   models.SubmissionStatus `json:"status" binding:"required"`
	Feedback string                  `json:"feedback"`
}

func (s *ReadingHoursService) Review(adminID, readingHourID uuid.UUID, input *ReviewInput) (*models.ReadingHour, error) {
	readingHour, err := s.GetByID(readingHourID)
	if err != nil {
		return nil, err
	}

	// Update status
	readingHour.Status = input.Status
	if err := s.db.Save(readingHour).Error; err != nil {
		return nil, err
	}

	// Create review record
	now := time.Now()
	review := &models.AdminReview{
		ReadingHourID: readingHourID,
		AdminID:       adminID,
		Decision:      input.Status,
		Feedback:      input.Feedback,
		ReviewedAt:    now,
	}
	if err := s.db.Create(review).Error; err != nil {
		return nil, err
	}

	// Send notification email
	var user models.User
	if err := s.db.First(&user, "id = ?", readingHour.UserID).Error; err == nil {
		if input.Status == models.StatusApproved {
			go s.emailService.SendSubmissionApprovedEmail(user.Email, user.FirstName, readingHour.Subject, input.Feedback)
		} else if input.Status == models.StatusRejected {
			go s.emailService.SendSubmissionRejectedEmail(user.Email, user.FirstName, readingHour.Subject, input.Feedback)
		}
	}

	return readingHour, nil
}

func (s *ReadingHoursService) GetWeeklyLeaderboard(campusID *uuid.UUID, limit int) ([]models.LeaderboardEntry, error) {
	var entries []models.LeaderboardEntry

	weekAgo := time.Now().AddDate(0, 0, -7)

	query := s.db.Model(&models.ReadingHour{}).
		Select(`
			users.id as user_id,
			users.first_name,
			users.last_name,
			users.avatar_url,
			users.campus_id,
			campuses.name as campus_name,
			SUM(reading_hours.duration_minutes) / 60.0 as total_hours,
			COUNT(reading_hours.id) as submission_count
		`).
		Joins("JOIN users ON users.id = reading_hours.user_id").
		Joins("LEFT JOIN campuses ON campuses.id = users.campus_id").
		Where("reading_hours.status = ? AND reading_hours.reading_date >= ?", models.StatusApproved, weekAgo).
		Group("users.id, users.first_name, users.last_name, users.avatar_url, users.campus_id, campuses.name").
		Order("total_hours DESC")

	if campusID != nil {
		query = query.Where("users.campus_id = ?", *campusID)
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Scan(&entries).Error; err != nil {
		return nil, err
	}

	// Assign ranks
	for i := range entries {
		entries[i].Rank = i + 1
	}

	return entries, nil
}

func (s *ReadingHoursService) GetCampusStats(campusID uuid.UUID) (*models.CampusStats, error) {
	var stats models.CampusStats

	// Total students
	s.db.Model(&models.User{}).Where("campus_id = ? AND role = ?", campusID, models.RoleStudent).Count(&stats.TotalStudents)

	// Active students (submitted in last 7 days)
	weekAgo := time.Now().AddDate(0, 0, -7)
	s.db.Model(&models.ReadingHour{}).
		Joins("JOIN users ON users.id = reading_hours.user_id").
		Where("users.campus_id = ? AND reading_hours.created_at >= ?", campusID, weekAgo).
		Distinct("reading_hours.user_id").
		Count(&stats.ActiveStudents)

	// Total reading hours
	var totalMinutes int64
	s.db.Model(&models.ReadingHour{}).
		Joins("JOIN users ON users.id = reading_hours.user_id").
		Where("users.campus_id = ? AND reading_hours.status = ?", campusID, models.StatusApproved).
		Select("COALESCE(SUM(reading_hours.duration_minutes), 0)").
		Scan(&totalMinutes)
	stats.TotalReadingHours = float64(totalMinutes) / 60.0

	// Pending submissions
	s.db.Model(&models.ReadingHour{}).
		Joins("JOIN users ON users.id = reading_hours.user_id").
		Where("users.campus_id = ? AND reading_hours.status = ?", campusID, models.StatusPending).
		Count(&stats.PendingSubmissions)

	// Weekly reading hours
	var weeklyMinutes int64
	s.db.Model(&models.ReadingHour{}).
		Joins("JOIN users ON users.id = reading_hours.user_id").
		Where("users.campus_id = ? AND reading_hours.status = ? AND reading_hours.reading_date >= ?", campusID, models.StatusApproved, weekAgo).
		Select("COALESCE(SUM(reading_hours.duration_minutes), 0)").
		Scan(&weeklyMinutes)
	stats.WeeklyReadingHours = float64(weeklyMinutes) / 60.0

	return &stats, nil
}

