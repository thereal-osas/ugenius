package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ugenius/backend/internal/services"
	"github.com/ugenius/backend/pkg/response"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// @Summary Register a new student
// @Tags auth
// @Accept json
// @Produce json
// @Param input body services.RegisterInput true "Registration details"
// @Success 201 {object} response.Response
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var input services.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.authService.Register(&input)
	if err != nil {
		if err == services.ErrUserExists {
			response.Conflict(c, "User with this email already exists")
			return
		}
		response.InternalError(c, "Failed to register user: "+err.Error())
		return
	}

	response.Created(c, "User registered successfully", user)
}

// RegisterAdmin godoc
// @Summary Register a new admin
// @Tags auth
// @Accept json
// @Produce json
// @Param input body services.RegisterInput true "Admin registration details"
// @Success 201 {object} response.Response
// @Router /auth/register-admin [post]
func (h *AuthHandler) RegisterAdmin(c *gin.Context) {
	var input services.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.authService.RegisterAdmin(&input)
	if err != nil {
		if err == services.ErrUserExists {
			response.Conflict(c, "User with this email already exists")
			return
		}
		response.InternalError(c, "Failed to register admin: "+err.Error())
		return
	}

	response.Created(c, "Admin registered successfully", user)
}

// Login godoc
// @Summary Login to account
// @Tags auth
// @Accept json
// @Produce json
// @Param input body services.LoginInput true "Login credentials"
// @Success 200 {object} response.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var input services.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, tokens, err := h.authService.Login(&input)
	if err != nil {
		switch err {
		case services.ErrInvalidCredentials:
			response.Unauthorized(c, "Invalid email or password")
		default:
			response.InternalError(c, "Login failed")
		}
		return
	}

	response.Success(c, "Login successful", gin.H{
		"user":          user,
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"expires_at":    tokens.ExpiresAt,
	})
}

// RefreshToken godoc
// @Summary Refresh access token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body object{refresh_token=string} true "Refresh token"
// @Success 200 {object} response.Response
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, tokens, err := h.authService.RefreshToken(input.RefreshToken)
	if err != nil {
		response.Unauthorized(c, "Invalid or expired refresh token")
		return
	}

	response.Success(c, "Token refreshed", gin.H{
		"user":          user,
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"expires_at":    tokens.ExpiresAt,
	})
}

// VerifyEmail godoc
// @Summary Verify email address
// @Tags auth
// @Accept json
// @Produce json
// @Param input body object{token=string} true "Verification token"
// @Success 200 {object} response.Response
// @Router /auth/verify-email [post]
func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	var input struct {
		Token string `json:"token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.authService.VerifyEmail(input.Token); err != nil {
		response.BadRequest(c, "Invalid or expired verification token")
		return
	}

	response.Success(c, "Email verified successfully. You can now log in.", nil)
}

// ForgotPassword godoc
// @Summary Request password reset
// @Tags auth
// @Accept json
// @Produce json
// @Param input body object{email=string} true "Email address"
// @Success 200 {object} response.Response
// @Router /auth/forgot-password [post]
func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	_ = h.authService.ForgotPassword(input.Email)

	// Always return success to prevent email enumeration
	response.Success(c, "If an account exists with this email, you will receive a password reset link.", nil)
}
