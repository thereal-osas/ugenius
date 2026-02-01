package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/pkg/response"
)

type EventsHandler struct {
	db *gorm.DB
}

func NewEventsHandler(db *gorm.DB) *EventsHandler {
	return &EventsHandler{db: db}
}

// List godoc
// @Summary List upcoming events
// @Tags events
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param type query string false "Filter by event type"
// @Param featured query bool false "Filter by featured status"
// @Success 200 {object} response.Response
// @Router /events [get]
func (h *EventsHandler) List(c *gin.Context) {
	page := 1
	pageSize := 10

	if p := c.Query("page"); p != "" {
		if parsed, err := parseInt(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	if ps := c.Query("page_size"); ps != "" {
		if parsed, err := parseInt(ps); err == nil && parsed > 0 && parsed <= 50 {
			pageSize = parsed
		}
	}

	query := h.db.Model(&models.Event{}).
		Where("status IN ?", []models.EventStatus{models.EventStatusUpcoming, models.EventStatusOngoing}).
		Where("start_time >= ?", time.Now().Add(-24*time.Hour)). // Include events from last 24 hours
		Order("start_time ASC")

	// Filter by type
	if eventType := c.Query("type"); eventType != "" {
		query = query.Where("type = ?", eventType)
	}

	// Filter by featured
	if featured := c.Query("featured"); featured == "true" {
		query = query.Where("is_featured = ?", true)
	}

	var total int64
	query.Count(&total)

	var events []models.Event
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&events).Error; err != nil {
		response.InternalError(c, "Failed to fetch events")
		return
	}

	response.Paginated(c, events, page, pageSize, total)
}

// GetByID godoc
// @Summary Get event by ID
// @Tags events
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} response.Response
// @Router /events/{id} [get]
func (h *EventsHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid event ID")
		return
	}

	var event models.Event
	if err := h.db.Preload("Campus").First(&event, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response.NotFound(c, "Event not found")
			return
		}
		response.InternalError(c, "Failed to fetch event")
		return
	}

	response.Success(c, "Event retrieved successfully", event)
}

// GetFeatured godoc
// @Summary Get featured upcoming events
// @Tags events
// @Produce json
// @Param limit query int false "Number of events to return (default 3)"
// @Success 200 {object} response.Response
// @Router /events/featured [get]
func (h *EventsHandler) GetFeatured(c *gin.Context) {
	limit := 3
	if l := c.Query("limit"); l != "" {
		if parsed, err := parseInt(l); err == nil && parsed > 0 && parsed <= 10 {
			limit = parsed
		}
	}

	var events []models.Event
	if err := h.db.
		Where("status = ?", models.EventStatusUpcoming).
		Where("start_time >= ?", time.Now()).
		Where("is_featured = ?", true).
		Order("start_time ASC").
		Limit(limit).
		Find(&events).Error; err != nil {
		response.InternalError(c, "Failed to fetch featured events")
		return
	}

	response.Success(c, "Featured events retrieved successfully", events)
}

// parseInt helper function
func parseInt(s string) (int, error) {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	return i, err
}
