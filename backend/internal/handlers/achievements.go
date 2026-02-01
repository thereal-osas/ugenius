package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ugenius/backend/internal/middleware"
	"github.com/ugenius/backend/internal/services"
	"github.com/ugenius/backend/pkg/response"
)

type AchievementsHandler struct {
	service *services.AchievementsService
}

func NewAchievementsHandler(service *services.AchievementsService) *AchievementsHandler {
	return &AchievementsHandler{service: service}
}

// List godoc
// @Summary List user achievements
// @Tags achievements
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response
// @Router /achievements [get]
func (h *AchievementsHandler) List(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	achievements, err := h.service.GetUserAchievements(userID)
	if err != nil {
		response.InternalError(c, "Failed to fetch achievements")
		return
	}

	response.Success(c, "", achievements)
}

// GetAllBadges godoc
// @Summary Get all available badges with earned status
// @Tags achievements
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response
// @Router /achievements/badges [get]
func (h *AchievementsHandler) GetAllBadges(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	badges, err := h.service.GetAllBadges(userID)
	if err != nil {
		response.InternalError(c, "Failed to fetch badges")
		return
	}

	response.Success(c, "", badges)
}

// CheckAchievements godoc
// @Summary Check and award new achievements
// @Tags achievements
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response
// @Router /achievements/check [post]
func (h *AchievementsHandler) CheckAchievements(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	newAchievements, err := h.service.CheckAndAwardAchievements(userID)
	if err != nil {
		response.InternalError(c, "Failed to check achievements")
		return
	}

	if len(newAchievements) > 0 {
		response.Success(c, "New achievements unlocked!", gin.H{
			"new_achievements": newAchievements,
		})
	} else {
		response.Success(c, "No new achievements", nil)
	}
}

