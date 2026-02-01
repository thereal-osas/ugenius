package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/internal/utils"
)

func (s *AuthService) VerifyEmail(token string) error {
	var user models.User
	if err := s.db.Where("verify_token = ?", token).First(&user).Error; err != nil {
		return ErrInvalidToken
	}

	user.EmailVerified = true
	user.VerifyToken = ""

	return s.db.Save(&user).Error
}

func (s *AuthService) ForgotPassword(email string) error {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		// Don't reveal if user exists
		return nil
	}

	resetToken := utils.GenerateResetToken()
	expiry := time.Now().Add(time.Hour) // 1 hour expiry

	user.ResetToken = resetToken
	user.ResetTokenExpiry = &expiry

	if err := s.db.Save(&user).Error; err != nil {
		return err
	}

	go s.emailService.SendPasswordResetEmail(user.Email, user.FirstName, resetToken)

	return nil
}

func (s *AuthService) ResetPassword(token, newPassword string) error {
	var user models.User
	if err := s.db.Where("reset_token = ? AND reset_token_expiry > ?", token, time.Now()).First(&user).Error; err != nil {
		return ErrInvalidToken
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.PasswordHash = hashedPassword
	user.ResetToken = ""
	user.ResetTokenExpiry = nil

	// Invalidate all refresh tokens for security
	s.db.Where("user_id = ?", user.ID).Delete(&models.RefreshToken{})

	return s.db.Save(&user).Error
}

func (s *AuthService) Logout(userID uuid.UUID, refreshToken string) error {
	return s.db.Where("user_id = ? AND token = ?", userID, refreshToken).Delete(&models.RefreshToken{}).Error
}

func (s *AuthService) LogoutAll(userID uuid.UUID) error {
	return s.db.Where("user_id = ?", userID).Delete(&models.RefreshToken{}).Error
}

func (s *AuthService) ResendVerification(email string) error {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return ErrUserNotFound
	}

	if user.EmailVerified {
		return nil // Already verified
	}

	// Generate new token
	user.VerifyToken = utils.GenerateVerificationToken()
	if err := s.db.Save(&user).Error; err != nil {
		return err
	}

	go s.emailService.SendVerificationEmail(user.Email, user.FirstName, user.VerifyToken)

	return nil
}

func (s *AuthService) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := s.db.Preload("Campus").First(&user, "id = ?", id).Error; err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

func (s *AuthService) UpdateProfile(userID uuid.UUID, firstName, lastName, phone, bio string) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, ErrUserNotFound
	}

	if firstName != "" {
		user.FirstName = firstName
	}
	if lastName != "" {
		user.LastName = lastName
	}
	if phone != "" {
		user.Phone = phone
	}
	if bio != "" {
		user.Bio = bio
	}

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) ChangePassword(userID uuid.UUID, currentPassword, newPassword string) error {
	var user models.User
	if err := s.db.First(&user, "id = ?", userID).Error; err != nil {
		return ErrUserNotFound
	}

	if !utils.CheckPassword(currentPassword, user.PasswordHash) {
		return ErrInvalidCredentials
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	user.PasswordHash = hashedPassword
	return s.db.Save(&user).Error
}

