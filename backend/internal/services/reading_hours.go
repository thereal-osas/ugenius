package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/models"
	"gorm.io/gorm"
)

var (
	ErrReadingHourNotFound = errors.New("reading hour not found")
	ErrNotOwner            = errors.New("not the owner of this reading hour")
	ErrCannotModify        = errors.New("cannot modify approved or rejected submission")
)

type ReadingHoursService struct {
	db           *gorm.DB
	emailService *EmailService
}

func NewReadingHoursService(db *gorm.DB, emailService *EmailService) *ReadingHoursService {
	return &ReadingHoursService{
		db:           db,
		emailService: emailService,
	}
}

type CreateReadingHourInput struct {
	ReadingDate     string `json:"reading_date" binding:"required"`
	DurationMinutes int    `json:"duration_minutes" binding:"required,min=1"`
	Subject         string `json:"subject" binding:"required"`
	Topic           string `json:"topic"`
	Description     string `json:"description"`
}

func (s *ReadingHoursService) Create(userID uuid.UUID, input *CreateReadingHourInput) (*models.ReadingHour, error) {
	readingDate, err := time.Parse("2006-01-02", input.ReadingDate)
	if err != nil {
		return nil, errors.New("invalid date format, use YYYY-MM-DD")
	}

	readingHour := &models.ReadingHour{
		UserID:          userID,
		ReadingDate:     readingDate,
		DurationMinutes: input.DurationMinutes,
		Subject:         input.Subject,
		Topic:           input.Topic,
		Description:     input.Description,
		Status:          models.StatusPending,
	}

	if err := s.db.Create(readingHour).Error; err != nil {
		return nil, err
	}

	return readingHour, nil
}

func (s *ReadingHoursService) GetByID(id uuid.UUID) (*models.ReadingHour, error) {
	var readingHour models.ReadingHour
	if err := s.db.Preload("User").Preload("Review.Admin").First(&readingHour, "id = ?", id).Error; err != nil {
		return nil, ErrReadingHourNotFound
	}
	return &readingHour, nil
}

func (s *ReadingHoursService) List(filter *models.ReadingHourFilter) ([]models.ReadingHour, int64, error) {
	var readingHours []models.ReadingHour
	var total int64

	query := s.db.Model(&models.ReadingHour{}).Preload("Review")

	if filter.UserID != nil {
		query = query.Where("user_id = ?", *filter.UserID)
	}
	if filter.CampusID != nil {
		query = query.Joins("JOIN users ON users.id = reading_hours.user_id").
			Where("users.campus_id = ?", *filter.CampusID)
	}
	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}
	if filter.StartDate != nil {
		query = query.Where("reading_date >= ?", *filter.StartDate)
	}
	if filter.EndDate != nil {
		query = query.Where("reading_date <= ?", *filter.EndDate)
	}
	if filter.Subject != "" {
		query = query.Where("subject ILIKE ?", "%"+filter.Subject+"%")
	}

	query.Count(&total)

	if filter.Page > 0 && filter.PageSize > 0 {
		offset := (filter.Page - 1) * filter.PageSize
		query = query.Offset(offset).Limit(filter.PageSize)
	}

	if err := query.Order("created_at DESC").Find(&readingHours).Error; err != nil {
		return nil, 0, err
	}

	return readingHours, total, nil
}

func (s *ReadingHoursService) Update(userID, id uuid.UUID, input *CreateReadingHourInput) (*models.ReadingHour, error) {
	readingHour, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	if readingHour.UserID != userID {
		return nil, ErrNotOwner
	}

	if readingHour.Status != models.StatusPending {
		return nil, ErrCannotModify
	}

	readingDate, _ := time.Parse("2006-01-02", input.ReadingDate)

	readingHour.ReadingDate = readingDate
	readingHour.DurationMinutes = input.DurationMinutes
	readingHour.Subject = input.Subject
	readingHour.Topic = input.Topic
	readingHour.Description = input.Description

	if err := s.db.Save(readingHour).Error; err != nil {
		return nil, err
	}

	return readingHour, nil
}

func (s *ReadingHoursService) Delete(userID, id uuid.UUID) error {
	readingHour, err := s.GetByID(id)
	if err != nil {
		return err
	}

	if readingHour.UserID != userID {
		return ErrNotOwner
	}

	if readingHour.Status != models.StatusPending {
		return ErrCannotModify
	}

	return s.db.Delete(readingHour).Error
}

func (s *ReadingHoursService) GetUserStats(userID uuid.UUID) (*models.UserStats, error) {
	var stats models.UserStats

	// Total approved reading hours (in minutes)
	var totalMinutes int64
	s.db.Model(&models.ReadingHour{}).
		Where("user_id = ? AND status = ?", userID, models.StatusApproved).
		Select("COALESCE(SUM(duration_minutes), 0)").
		Scan(&totalMinutes)
	stats.TotalReadingHours = float64(totalMinutes) / 60.0

	// Weekly reading hours
	weekAgo := time.Now().AddDate(0, 0, -7)
	var weeklyMinutes int64
	s.db.Model(&models.ReadingHour{}).
		Where("user_id = ? AND status = ? AND reading_date >= ?", userID, models.StatusApproved, weekAgo).
		Select("COALESCE(SUM(duration_minutes), 0)").
		Scan(&weeklyMinutes)
	stats.WeeklyReadingHours = float64(weeklyMinutes) / 60.0

	// Monthly reading hours
	monthAgo := time.Now().AddDate(0, -1, 0)
	var monthlyMinutes int64
	s.db.Model(&models.ReadingHour{}).
		Where("user_id = ? AND status = ? AND reading_date >= ?", userID, models.StatusApproved, monthAgo).
		Select("COALESCE(SUM(duration_minutes), 0)").
		Scan(&monthlyMinutes)
	stats.MonthlyReadingHours = float64(monthlyMinutes) / 60.0

	// Submission counts
	s.db.Model(&models.ReadingHour{}).Where("user_id = ?", userID).Count(&stats.TotalSubmissions)
	s.db.Model(&models.ReadingHour{}).Where("user_id = ? AND status = ?", userID, models.StatusApproved).Count(&stats.ApprovedSubmissions)
	s.db.Model(&models.ReadingHour{}).Where("user_id = ? AND status = ?", userID, models.StatusPending).Count(&stats.PendingSubmissions)
	s.db.Model(&models.ReadingHour{}).Where("user_id = ? AND status = ?", userID, models.StatusRejected).Count(&stats.RejectedSubmissions)

	// Achievement count
	s.db.Model(&models.Achievement{}).Where("user_id = ?", userID).Count(&stats.TotalAchievements)

	// Goals
	s.db.Model(&models.StudyGoal{}).Where("user_id = ? AND is_completed = ?", userID, true).Count(&stats.GoalsCompleted)
	s.db.Model(&models.StudyGoal{}).Where("user_id = ? AND is_completed = ?", userID, false).Count(&stats.GoalsInProgress)

	return &stats, nil
}

