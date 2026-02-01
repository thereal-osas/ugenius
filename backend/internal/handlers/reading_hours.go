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

type ReadingHoursHandler struct {
	service *services.ReadingHoursService
}

func NewReadingHoursHandler(service *services.ReadingHoursService) *ReadingHoursHandler {
	return &ReadingHoursHandler{service: service}
}

// Create godoc
// @Summary Submit reading hours
// @Tags reading-hours
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body services.CreateReadingHourInput true "Reading hour details"
// @Success 201 {object} response.Response
// @Router /reading-hours [post]
func (h *ReadingHoursHandler) Create(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	var input services.CreateReadingHourInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	readingHour, err := h.service.Create(userID, &input)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Created(c, "Reading hours submitted successfully", readingHour)
}

// GetByID godoc
// @Summary Get reading hour by ID
// @Tags reading-hours
// @Security BearerAuth
// @Produce json
// @Param id path string true "Reading hour ID"
// @Success 200 {object} response.Response
// @Router /reading-hours/{id} [get]
func (h *ReadingHoursHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	readingHour, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Reading hour not found")
		return
	}

	response.Success(c, "", readingHour)
}

// List godoc
// @Summary List reading hours
// @Tags reading-hours
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param status query string false "Filter by status"
// @Param subject query string false "Filter by subject"
// @Success 200 {object} response.Response
// @Router /reading-hours [get]
func (h *ReadingHoursHandler) List(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	filter := &models.ReadingHourFilter{
		UserID:   &userID,
		Page:     1,
		PageSize: 20,
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

// Update godoc
// @Summary Update reading hour
// @Tags reading-hours
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Reading hour ID"
// @Param input body services.CreateReadingHourInput true "Updated details"
// @Success 200 {object} response.Response
// @Router /reading-hours/{id} [put]
func (h *ReadingHoursHandler) Update(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var input services.CreateReadingHourInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	readingHour, err := h.service.Update(userID, id, &input)
	if err != nil {
		switch err {
		case services.ErrReadingHourNotFound:
			response.NotFound(c, "Reading hour not found")
		case services.ErrNotOwner:
			response.Forbidden(c, "You can only update your own submissions")
		case services.ErrCannotModify:
			response.BadRequest(c, "Cannot modify approved or rejected submissions")
		default:
			response.InternalError(c, "Failed to update reading hour")
		}
		return
	}

	response.Success(c, "Reading hour updated successfully", readingHour)
}

// Delete godoc
// @Summary Delete reading hour
// @Tags reading-hours
// @Security BearerAuth
// @Param id path string true "Reading hour ID"
// @Success 200 {object} response.Response
// @Router /reading-hours/{id} [delete]
func (h *ReadingHoursHandler) Delete(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.service.Delete(userID, id); err != nil {
		switch err {
		case services.ErrReadingHourNotFound:
			response.NotFound(c, "Reading hour not found")
		case services.ErrNotOwner:
			response.Forbidden(c, "You can only delete your own submissions")
		case services.ErrCannotModify:
			response.BadRequest(c, "Cannot delete approved or rejected submissions")
		default:
			response.InternalError(c, "Failed to delete reading hour")
		}
		return
	}

	response.Success(c, "Reading hour deleted successfully", nil)
}

