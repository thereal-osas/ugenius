package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/models"
	"gorm.io/gorm"
)

type WaitlistHandler struct {
	db *gorm.DB
}

func NewWaitlistHandler(db *gorm.DB) *WaitlistHandler {
	return &WaitlistHandler{db: db}
}

// Register handles waitlist registration
func (h *WaitlistHandler) Register(c *gin.Context) {
	var req models.WaitlistRegistrationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request",
			"details": err.Error(),
		})
		return
	}

	// Normalize email
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))

	// Check if email already registered
	var existingCount int64
	h.db.Model(&models.WaitlistRegistration{}).Where("email = ?", req.Email).Count(&existingCount)
	if existingCount > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This email is already registered on our waitlist",
		})
		return
	}

	// Create waitlist registration
	registration := models.WaitlistRegistration{
		FullName:   strings.TrimSpace(req.FullName),
		Email:      req.Email,
		Phone:      strings.TrimSpace(req.Phone),
		Level:      req.Level,
		Faculty:    strings.TrimSpace(req.Faculty),
		Department: strings.TrimSpace(req.Department),
		Address:    strings.TrimSpace(req.Address),
	}

	// Parse and set campus ID if provided
	if req.CampusID != "" {
		campusID, err := uuid.Parse(req.CampusID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid campus ID",
			})
			return
		}

		// Verify campus exists
		var campus models.Campus
		if err := h.db.First(&campus, "id = ?", campusID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Campus not found",
			})
			return
		}
		registration.CampusID = &campusID
	}

	// Save to database
	if err := h.db.Create(&registration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register. Please try again.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful! We'll notify you when U-Genius launches at your campus.",
		"data": gin.H{
			"id":        registration.ID,
			"full_name": registration.FullName,
			"email":     registration.Email,
		},
	})
}

