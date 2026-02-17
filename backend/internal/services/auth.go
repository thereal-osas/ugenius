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
	Email       string  `json:"email" binding:"required,email"`
	Password    string  `json:"password" binding:"required,min=8"`
	FirstName   string  `json:"first_name" binding:"required"`
	LastName    string  `json:"last_name" binding:"required"`
	Phone       string  `json:"phone"`
	CampusID    *string `json:"campus_id"`
	Institution string  `json:"institution"`
	Department  string  `json:"department"`
	Level       string  `json:"level"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateUserInput struct {
	Email       string
	Password    string
	FirstName   string
	LastName    string
	Phone       string
	CampusID    *uuid.UUID
	Institution string
	Department  string
	Level       string
	Role        models.Role
}

func (s *AuthService) CreateUser(input CreateUserInput) (*models.User, error) {
	// Validate campus if provided
	if input.CampusID != nil {
		var campus models.Campus
		if err := s.db.First(&campus, "id = ?", input.CampusID).Error; err != nil {
			return nil, ErrCampusNotFound
		}
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
		Role:         input.Role,
		CampusID:     input.CampusID,
		Institution:  input.Institution,
		Department:   input.Department,
		Level:        input.Level,
	}

	// Generate verification token for students
	if input.Role == models.RoleStudent {
		user.VerifyToken = utils.GenerateVerificationToken()
	} else {
		user.EmailVerified = true // Admins are verified by default
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Send verification email for students
	if input.Role == models.RoleStudent {
		go s.emailService.SendVerificationEmail(user.Email, user.FirstName, user.VerifyToken)
	}

	return user, nil
}

func (s *AuthService) Register(input *RegisterInput) (*models.User, error) {
	// Check if user exists
	var existingUser models.User
	if err := s.db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return nil, ErrUserExists
	}

	var campusUUID *uuid.UUID
	if input.CampusID != nil && *input.CampusID != "" {
		parsedUUID, err := uuid.Parse(*input.CampusID)
		if err != nil {
			return nil, ErrCampusNotFound
		}
		campusUUID = &parsedUUID
	}

	return s.CreateUser(CreateUserInput{
		Email:       input.Email,
		Password:    input.Password,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Phone:       input.Phone,
		CampusID:    campusUUID,
		Institution: input.Institution,
		Department:  input.Department,
		Level:       input.Level,
		Role:        models.RoleStudent,
	})
}

func (s *AuthService) RegisterAdmin(input *RegisterInput) (*models.User, error) {
	// Check if user exists
	var existingUser models.User
	if err := s.db.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return nil, ErrUserExists
	}

	var campusUUID *uuid.UUID
	if input.CampusID != nil && *input.CampusID != "" {
		parsedUUID, err := uuid.Parse(*input.CampusID)
		if err != nil {
			return nil, ErrCampusNotFound
		}
		campusUUID = &parsedUUID
	}

	return s.CreateUser(CreateUserInput{
		Email:       input.Email,
		Password:    input.Password,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Phone:       input.Phone,
		CampusID:    campusUUID,
		Institution: input.Institution,
		Department:  input.Department,
		Level:       input.Level,
		Role:        models.RoleCampusAdmin,
	})
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

func (s *AuthService) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s *AuthService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Preload("Campus").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *AuthService) DeleteUser(userID string) error {
	// Prevent deletion of super admin
	var user models.User
	if err := s.db.First(&user, "id = ?", userID).Error; err != nil {
		return err
	}

	if user.Role == models.RoleSuperAdmin {
		return errors.New("cannot delete super admin user")
	}

	// Soft delete the user
	return s.db.Delete(&models.User{}, "id = ?", userID).Error
}
