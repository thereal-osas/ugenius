package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/middleware"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/pkg/response"
	"gorm.io/gorm"
)

type NotificationsHandler struct {
	db *gorm.DB
}

func NewNotificationsHandler(db *gorm.DB) *NotificationsHandler {
	return &NotificationsHandler{db: db}
}

// List godoc
// @Summary List user notifications
// @Tags notifications
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Param unread_only query bool false "Show only unread"
// @Success 200 {object} response.Response
// @Router /notifications [get]
func (h *NotificationsHandler) List(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	page := 1
	pageSize := 20
	if p, _ := strconv.Atoi(c.Query("page")); p > 0 {
		page = p
	}
	if ps, _ := strconv.Atoi(c.Query("page_size")); ps > 0 && ps <= 100 {
		pageSize = ps
	}

	query := h.db.Where("user_id = ?", userID)
	if c.Query("unread_only") == "true" {
		query = query.Where("is_read = ?", false)
	}

	var total int64
	query.Model(&models.Notification{}).Count(&total)

	var notifications []models.Notification
	offset := (page - 1) * pageSize
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&notifications).Error; err != nil {
		response.InternalError(c, "Failed to fetch notifications")
		return
	}

	response.Paginated(c, notifications, page, pageSize, total)
}

// MarkAsRead godoc
// @Summary Mark notification as read
// @Tags notifications
// @Security BearerAuth
// @Param id path string true "Notification ID"
// @Success 200 {object} response.Response
// @Router /notifications/{id}/read [post]
func (h *NotificationsHandler) MarkAsRead(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var notification models.Notification
	if err := h.db.Where("id = ? AND user_id = ?", id, userID).First(&notification).Error; err != nil {
		response.NotFound(c, "Notification not found")
		return
	}

	notification.MarkAsRead()
	h.db.Save(&notification)

	response.Success(c, "Notification marked as read", nil)
}

// MarkAllAsRead godoc
// @Summary Mark all notifications as read
// @Tags notifications
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Router /notifications/read-all [post]
func (h *NotificationsHandler) MarkAllAsRead(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	h.db.Model(&models.Notification{}).
		Where("user_id = ? AND is_read = ?", userID, false).
		Update("is_read", true)

	response.Success(c, "All notifications marked as read", nil)
}

// GetUnreadCount godoc
// @Summary Get unread notification count
// @Tags notifications
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response
// @Router /notifications/unread-count [get]
func (h *NotificationsHandler) GetUnreadCount(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	var count int64
	h.db.Model(&models.Notification{}).Where("user_id = ? AND is_read = ?", userID, false).Count(&count)

	response.Success(c, "", gin.H{"unread_count": count})
}

// Delete godoc
// @Summary Delete a notification
// @Tags notifications
// @Security BearerAuth
// @Param id path string true "Notification ID"
// @Success 200 {object} response.Response
// @Router /notifications/{id} [delete]
func (h *NotificationsHandler) Delete(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	result := h.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Notification{})
	if result.RowsAffected == 0 {
		response.NotFound(c, "Notification not found")
		return
	}

	response.Success(c, "Notification deleted", nil)
}

