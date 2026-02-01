package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/config"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/internal/utils"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user with this email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrEmailNotVerified   = errors.New("email not verified")
	ErrInvalidToken       = errors.New("invalid or expired token")
	ErrCampusNotFound     = errors.New("campus not found")
)

type AuthService struct {
	db           *gorm.DB
	cfg          *config.Config
	emailService *EmailService
}

func NewAuthService(db *gorm.DB, cfg *config.Config, emailService *EmailService) *AuthService {
	return &AuthService{
		db:           db,
		cfg:          cfg,
		emailService: emailService,
	}
}

type RegisterInput struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=8"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Phone       string `json:"phone"`
	CampusID    string `json:"campus_id" binding:"required"`
	Institution string `json:"institution"`
	Department  string `json:"department"`
	Level       string `json:"level"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (s *AuthService) Register(input *RegisterInput) (*models.User, error) {
	// Check if user exists
	var existingUser models.User
	if err := s.db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return nil, ErrUserExists
	}

	// Validate campus
	campusUUID, err := uuid.Parse(input.CampusID)
	if err != nil {
		return nil, ErrCampusNotFound
	}

	var campus models.Campus
	if err := s.db.First(&campus, "id = ?", campusUUID).Error; err != nil {
		return nil, ErrCampusNotFound
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &models.User{
		Email:        input.Email,
		PasswordHash: hashedPassword,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Phone:        input.Phone,
		Role:         models.RoleStudent,
		CampusID:     &campusUUID,
		Institution:  input.Institution,
		Department:   input.Department,
		Level:        input.Level,
		VerifyToken:  utils.GenerateVerificationToken(),
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	// Send verification email
	go s.emailService.SendVerificationEmail(user.Email, user.FirstName, user.VerifyToken)

	return user, nil
}

func (s *AuthService) Login(input *LoginInput) (*models.User, *utils.TokenPair, error) {
	var user models.User
	if err := s.db.Preload("Campus").Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, nil, ErrInvalidCredentials
	}

	if !utils.CheckPassword(input.Password, user.PasswordHash) {
		return nil, nil, ErrInvalidCredentials
	}

	// Generate tokens
	tokenPair, err := s.generateTokenPair(&user)
	if err != nil {
		return nil, nil, err
	}

	return &user, tokenPair, nil
}

func (s *AuthService) generateTokenPair(user *models.User) (*utils.TokenPair, error) {
	accessToken, expiresAt, err := utils.GenerateAccessToken(user, s.cfg.JWT.Secret, s.cfg.JWT.ExpiryHours)
	if err != nil {
		return nil, err
	}

	refreshTokenStr, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Store refresh token
	refreshToken := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshTokenStr,
		ExpiresAt: time.Now().AddDate(0, 0, s.cfg.JWT.RefreshExpiryDays),
	}
	if err := s.db.Create(refreshToken).Error; err != nil {
		return nil, err
	}

	return &utils.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		ExpiresAt:    expiresAt,
	}, nil
}

func (s *AuthService) RefreshToken(refreshTokenStr string) (*models.User, *utils.TokenPair, error) {
	var refreshToken models.RefreshToken
	if err := s.db.Where("token = ? AND expires_at > ?", refreshTokenStr, time.Now()).First(&refreshToken).Error; err != nil {
		return nil, nil, ErrInvalidToken
	}

	// Delete old refresh token
	s.db.Delete(&refreshToken)

	var user models.User
	if err := s.db.Preload("Campus").First(&user, "id = ?", refreshToken.UserID).Error; err != nil {
		return nil, nil, ErrUserNotFound
	}

	// Generate new token pair
	tokenPair, err := s.generateTokenPair(&user)
	if err != nil {
		return nil, nil, err
	}

	return &user, tokenPair, nil
}
