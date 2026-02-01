package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ugenius/backend/internal/middleware"
	"github.com/ugenius/backend/internal/services"
	"github.com/ugenius/backend/pkg/response"
)

// ResetPassword godoc
// @Summary Reset password with token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body object{token=string,password=string} true "Reset details"
// @Success 200 {object} response.Response
// @Router /auth/reset-password [post]
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var input struct {
		Token    string `json:"token" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.authService.ResetPassword(input.Token, input.Password); err != nil {
		response.BadRequest(c, "Invalid or expired reset token")
		return
	}

	response.Success(c, "Password reset successfully. You can now log in with your new password.", nil)
}

// Logout godoc
// @Summary Logout (invalidate refresh token)
// @Tags auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body object{refresh_token=string} true "Refresh token to invalidate"
// @Success 200 {object} response.Response
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_ = h.authService.Logout(userID, input.RefreshToken)

	response.Success(c, "Logged out successfully", nil)
}

// LogoutAll godoc
// @Summary Logout from all devices
// @Tags auth
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /auth/logout-all [post]
func (h *AuthHandler) LogoutAll(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	_ = h.authService.LogoutAll(userID)

	response.Success(c, "Logged out from all devices", nil)
}

// ResendVerification godoc
// @Summary Resend verification email
// @Tags auth
// @Accept json
// @Produce json
// @Param input body object{email=string} true "Email address"
// @Success 200 {object} response.Response
// @Router /auth/resend-verification [post]
func (h *AuthHandler) ResendVerification(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_ = h.authService.ResendVerification(input.Email)

	response.Success(c, "If an unverified account exists, a new verification email has been sent.", nil)
}

// GetMe godoc
// @Summary Get current user profile
// @Tags user
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /me [get]
func (h *AuthHandler) GetMe(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.Unauthorized(c, "")
		return
	}

	user, err := h.authService.GetUserByID(userID)
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}

	response.Success(c, "", user)
}

// UpdateProfile godoc
// @Summary Update user profile
// @Tags user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body object{first_name=string,last_name=string,phone=string,bio=string} true "Profile updates"
// @Success 200 {object} response.Response
// @Router /me [put]
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.Unauthorized(c, "")
		return
	}

	var input struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Phone     string `json:"phone"`
		Bio       string `json:"bio"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.authService.UpdateProfile(userID, input.FirstName, input.LastName, input.Phone, input.Bio)
	if err != nil {
		response.InternalError(c, "Failed to update profile")
		return
	}

	response.Success(c, "Profile updated successfully", user)
}

// ChangePassword godoc
// @Summary Change password
// @Tags user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body object{current_password=string,new_password=string} true "Password change"
// @Success 200 {object} response.Response
// @Router /me/password [put]
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		response.Unauthorized(c, "")
		return
	}

	var input struct {
		CurrentPassword string `json:"current_password" binding:"required"`
		NewPassword     string `json:"new_password" binding:"required,min=8"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.authService.ChangePassword(userID, input.CurrentPassword, input.NewPassword); err != nil {
		if err == services.ErrInvalidCredentials {
			response.BadRequest(c, "Current password is incorrect")
			return
		}
		response.InternalError(c, "Failed to change password")
		return
	}

	response.Success(c, "Password changed successfully", nil)
}

