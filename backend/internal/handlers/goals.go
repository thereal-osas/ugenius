package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/middleware"
	"github.com/ugenius/backend/internal/services"
	"github.com/ugenius/backend/pkg/response"
)

type GoalsHandler struct {
	service *services.GoalsService
}

func NewGoalsHandler(service *services.GoalsService) *GoalsHandler {
	return &GoalsHandler{service: service}
}

// Create godoc
// @Summary Create a new goal
// @Tags goals
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body services.CreateGoalInput true "Goal details"
// @Success 201 {object} response.Response
// @Router /goals [post]
func (h *GoalsHandler) Create(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	var input services.CreateGoalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	goal, err := h.service.Create(userID, &input)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Created(c, "Goal created successfully", goal)
}

// GetByID godoc
// @Summary Get goal by ID
// @Tags goals
// @Security BearerAuth
// @Produce json
// @Param id path string true "Goal ID"
// @Success 200 {object} response.Response
// @Router /goals/{id} [get]
func (h *GoalsHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	goal, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Goal not found")
		return
	}

	response.Success(c, "", goal)
}

// List godoc
// @Summary List user goals
// @Tags goals
// @Security BearerAuth
// @Produce json
// @Param active_only query bool false "Show only active goals"
// @Success 200 {object} response.Response
// @Router /goals [get]
func (h *GoalsHandler) List(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	activeOnly := c.Query("active_only") == "true"

	goals, err := h.service.List(userID, activeOnly)
	if err != nil {
		response.InternalError(c, "Failed to fetch goals")
		return
	}

	response.Success(c, "", goals)
}

// Update godoc
// @Summary Update a goal
// @Tags goals
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Goal ID"
// @Param input body services.CreateGoalInput true "Updated goal details"
// @Success 200 {object} response.Response
// @Router /goals/{id} [put]
func (h *GoalsHandler) Update(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var input services.CreateGoalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	goal, err := h.service.Update(userID, id, &input)
	if err != nil {
		switch err {
		case services.ErrGoalNotFound:
			response.NotFound(c, "Goal not found")
		case services.ErrGoalNotOwner:
			response.Forbidden(c, "You can only update your own goals")
		default:
			response.InternalError(c, "Failed to update goal")
		}
		return
	}

	response.Success(c, "Goal updated successfully", goal)
}

// Delete godoc
// @Summary Delete a goal
// @Tags goals
// @Security BearerAuth
// @Param id path string true "Goal ID"
// @Success 200 {object} response.Response
// @Router /goals/{id} [delete]
func (h *GoalsHandler) Delete(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	if err := h.service.Delete(userID, id); err != nil {
		switch err {
		case services.ErrGoalNotFound:
			response.NotFound(c, "Goal not found")
		case services.ErrGoalNotOwner:
			response.Forbidden(c, "You can only delete your own goals")
		default:
			response.InternalError(c, "Failed to delete goal")
		}
		return
	}

	response.Success(c, "Goal deleted successfully", nil)
}

