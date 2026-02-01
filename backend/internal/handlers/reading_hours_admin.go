package handlers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/middleware"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/internal/services"
	"github.com/ugenius/backend/pkg/response"
)

// AdminList godoc
// @Summary List all reading hours (admin)
// @Tags admin
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param status query string false "Filter by status"
// @Param user_id query string false "Filter by user ID"
// @Success 200 {object} response.Response
// @Router /admin/reading-hours [get]
func (h *ReadingHoursHandler) AdminList(c *gin.Context) {
	campusID, hasCampus := middleware.GetCampusID(c)
	role, _ := middleware.GetUserRole(c)

	filter := &models.ReadingHourFilter{
		Page:     1,
		PageSize: 20,
	}

	// Campus admins can only see their campus
	if role == models.RoleCampusAdmin && hasCampus {
		filter.CampusID = &campusID
	}

	if page, _ := strconv.Atoi(c.Query("page")); page > 0 {
		filter.Page = page
	}
	if pageSize, _ := strconv.Atoi(c.Query("page_size")); pageSize > 0 && pageSize <= 100 {
		filter.PageSize = pageSize
	}
	if status := c.Query("status"); status != "" {
		s := models.SubmissionStatus(status)
		filter.Status = &s
	}
	if userIDStr := c.Query("user_id"); userIDStr != "" {
		if userID, err := uuid.Parse(userIDStr); err == nil {
			filter.UserID = &userID
		}
	}
	if subject := c.Query("subject"); subject != "" {
		filter.Subject = subject
	}
	if startDate := c.Query("start_date"); startDate != "" {
		if t, err := time.Parse("2006-01-02", startDate); err == nil {
			filter.StartDate = &t
		}
	}
	if endDate := c.Query("end_date"); endDate != "" {
		if t, err := time.Parse("2006-01-02", endDate); err == nil {
			filter.EndDate = &t
		}
	}

	readingHours, total, err := h.service.List(filter)
	if err != nil {
		response.InternalError(c, "Failed to fetch reading hours")
		return
	}

	response.Paginated(c, readingHours, filter.Page, filter.PageSize, total)
}

// Review godoc
// @Summary Review a reading hour submission
// @Tags admin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Reading hour ID"
// @Param input body services.ReviewInput true "Review details"
// @Success 200 {object} response.Response
// @Router /admin/reading-hours/{id}/review [post]
func (h *ReadingHoursHandler) Review(c *gin.Context) {
	adminID, _ := middleware.GetUserID(c)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var input services.ReviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if input.Status != models.StatusApproved && input.Status != models.StatusRejected {
		response.BadRequest(c, "Status must be 'approved' or 'rejected'")
		return
	}

	readingHour, err := h.service.Review(adminID, id, &input)
	if err != nil {
		response.InternalError(c, "Failed to review submission")
		return
	}

	response.Success(c, "Submission reviewed successfully", readingHour)
}

// GetLeaderboard godoc
// @Summary Get weekly leaderboard
// @Tags leaderboard
// @Produce json
// @Param campus_id query string false "Filter by campus ID"
// @Param limit query int false "Number of entries (default 10)"
// @Success 200 {object} response.Response
// @Router /leaderboard [get]
func (h *ReadingHoursHandler) GetLeaderboard(c *gin.Context) {
	var campusID *uuid.UUID
	if campusIDStr := c.Query("campus_id"); campusIDStr != "" {
		if id, err := uuid.Parse(campusIDStr); err == nil {
			campusID = &id
		}
	}

	limit := 10
	if l, _ := strconv.Atoi(c.Query("limit")); l > 0 && l <= 100 {
		limit = l
	}

	entries, err := h.service.GetWeeklyLeaderboard(campusID, limit)
	if err != nil {
		response.InternalError(c, "Failed to fetch leaderboard")
		return
	}

	response.Success(c, "", gin.H{
		"period":  "weekly",
		"entries": entries,
	})
}

// GetStats godoc
// @Summary Get user stats
// @Tags reading-hours
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response
// @Router /reading-hours/stats [get]
func (h *ReadingHoursHandler) GetStats(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	stats, err := h.service.GetUserStats(userID)
	if err != nil {
		response.InternalError(c, "Failed to fetch stats")
		return
	}

	response.Success(c, "", stats)
}

// GetCampusStats godoc
// @Summary Get campus stats (admin)
// @Tags admin
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response
// @Router /admin/stats [get]
func (h *ReadingHoursHandler) GetCampusStats(c *gin.Context) {
	campusID, ok := middleware.GetCampusID(c)
	if !ok {
		response.BadRequest(c, "Campus not found")
		return
	}

	stats, err := h.service.GetCampusStats(campusID)
	if err != nil {
		response.InternalError(c, "Failed to fetch campus stats")
		return
	}

	response.Success(c, "", stats)
}

