package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/models"
	"gorm.io/gorm"
)

var (
	ErrGoalNotFound = errors.New("goal not found")
	ErrGoalNotOwner = errors.New("not the owner of this goal")
)

type GoalsService struct {
	db *gorm.DB
}

func NewGoalsService(db *gorm.DB) *GoalsService {
	return &GoalsService{db: db}
}

type CreateGoalInput struct {
	Title       string            `json:"title" binding:"required"`
	Description string            `json:"description"`
	GoalType    models.GoalType   `json:"goal_type" binding:"required"`
	Period      models.GoalPeriod `json:"period" binding:"required"`
	TargetValue int               `json:"target_value" binding:"required,min=1"`
	StartDate   string            `json:"start_date" binding:"required"`
	EndDate     string            `json:"end_date" binding:"required"`
}

func (s *GoalsService) Create(userID uuid.UUID, input *CreateGoalInput) (*models.StudyGoal, error) {
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return nil, errors.New("invalid start date format, use YYYY-MM-DD")
	}

	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return nil, errors.New("invalid end date format, use YYYY-MM-DD")
	}

	if endDate.Before(startDate) {
		return nil, errors.New("end date must be after start date")
	}

	goal := &models.StudyGoal{
		UserID:       userID,
		Title:        input.Title,
		Description:  input.Description,
		GoalType:     input.GoalType,
		Period:       input.Period,
		TargetValue:  input.TargetValue,
		CurrentValue: 0,
		StartDate:    startDate,
		EndDate:      endDate,
		IsCompleted:  false,
	}

	if err := s.db.Create(goal).Error; err != nil {
		return nil, err
	}

	return goal, nil
}

func (s *GoalsService) GetByID(id uuid.UUID) (*models.StudyGoal, error) {
	var goal models.StudyGoal
	if err := s.db.First(&goal, "id = ?", id).Error; err != nil {
		return nil, ErrGoalNotFound
	}
	return &goal, nil
}

func (s *GoalsService) List(userID uuid.UUID, activeOnly bool) ([]models.StudyGoal, error) {
	var goals []models.StudyGoal
	query := s.db.Where("user_id = ?", userID)

	if activeOnly {
		query = query.Where("is_completed = ? AND end_date >= ?", false, time.Now())
	}

	if err := query.Order("created_at DESC").Find(&goals).Error; err != nil {
		return nil, err
	}

	return goals, nil
}

func (s *GoalsService) Update(userID, id uuid.UUID, input *CreateGoalInput) (*models.StudyGoal, error) {
	goal, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	if goal.UserID != userID {
		return nil, ErrGoalNotOwner
	}

	startDate, _ := time.Parse("2006-01-02", input.StartDate)
	endDate, _ := time.Parse("2006-01-02", input.EndDate)

	goal.Title = input.Title
	goal.Description = input.Description
	goal.GoalType = input.GoalType
	goal.Period = input.Period
	goal.TargetValue = input.TargetValue
	goal.StartDate = startDate
	goal.EndDate = endDate

	if err := s.db.Save(goal).Error; err != nil {
		return nil, err
	}

	return goal, nil
}

func (s *GoalsService) Delete(userID, id uuid.UUID) error {
	goal, err := s.GetByID(id)
	if err != nil {
		return err
	}

	if goal.UserID != userID {
		return ErrGoalNotOwner
	}

	return s.db.Delete(goal).Error
}

func (s *GoalsService) UpdateProgress(userID uuid.UUID) error {
	var goals []models.StudyGoal
	s.db.Where("user_id = ? AND is_completed = ?", userID, false).Find(&goals)

	for _, goal := range goals {
		var currentValue int

		switch goal.GoalType {
		case models.GoalTypeReadingHours:
			var totalMinutes int64
			s.db.Model(&models.ReadingHour{}).
				Where("user_id = ? AND status = ? AND reading_date BETWEEN ? AND ?",
					userID, models.StatusApproved, goal.StartDate, goal.EndDate).
				Select("COALESCE(SUM(duration_minutes), 0)").
				Scan(&totalMinutes)
			currentValue = int(totalMinutes / 60)

		case models.GoalTypeSubmissions:
			var count int64
			s.db.Model(&models.ReadingHour{}).
				Where("user_id = ? AND status = ? AND reading_date BETWEEN ? AND ?",
					userID, models.StatusApproved, goal.StartDate, goal.EndDate).
				Count(&count)
			currentValue = int(count)

		case models.GoalTypeStreak:
			// Count consecutive days
			currentValue = s.calculateStreak(userID, goal.StartDate, goal.EndDate)
		}

		goal.CurrentValue = currentValue
		if currentValue >= goal.TargetValue {
			goal.IsCompleted = true
			now := time.Now()
			goal.CompletedAt = &now
		}

		s.db.Save(&goal)
	}

	return nil
}

func (s *GoalsService) calculateStreak(userID uuid.UUID, startDate, endDate time.Time) int {
	var dates []time.Time
	s.db.Model(&models.ReadingHour{}).
		Where("user_id = ? AND status = ? AND reading_date BETWEEN ? AND ?",
			userID, models.StatusApproved, startDate, endDate).
		Select("DISTINCT DATE(reading_date)").
		Order("reading_date DESC").
		Pluck("reading_date", &dates)

	if len(dates) == 0 {
		return 0
	}

	maxStreak := 1
	currentStreak := 1
	for i := 1; i < len(dates); i++ {
		diff := dates[i-1].Sub(dates[i]).Hours() / 24
		if diff <= 1 {
			currentStreak++
			if currentStreak > maxStreak {
				maxStreak = currentStreak
			}
		} else {
			currentStreak = 1
		}
	}
	return maxStreak
}

