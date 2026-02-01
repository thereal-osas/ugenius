package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/pkg/response"
	"gorm.io/gorm"
)

type CampusHandler struct {
	db *gorm.DB
}

func NewCampusHandler(db *gorm.DB) *CampusHandler {
	return &CampusHandler{db: db}
}

// List godoc
// @Summary List all campuses
// @Tags campuses
// @Produce json
// @Success 200 {object} response.Response
// @Router /campuses [get]
func (h *CampusHandler) List(c *gin.Context) {
	var campuses []models.Campus
	if err := h.db.Where("is_active = ?", true).Order("name ASC").Find(&campuses).Error; err != nil {
		response.InternalError(c, "Failed to fetch campuses")
		return
	}

	response.Success(c, "", campuses)
}

// GetByID godoc
// @Summary Get campus by ID
// @Tags campuses
// @Produce json
// @Param id path string true "Campus ID"
// @Success 200 {object} response.Response
// @Router /campuses/{id} [get]
func (h *CampusHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var campus models.Campus
	if err := h.db.First(&campus, "id = ?", id).Error; err != nil {
		response.NotFound(c, "Campus not found")
		return
	}

	response.Success(c, "", campus)
}

// Create godoc
// @Summary Create a new campus (super admin only)
// @Tags campuses
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body object{name=string,code=string,location=string,state=string,country=string} true "Campus details"
// @Success 201 {object} response.Response
// @Router /admin/campuses [post]
func (h *CampusHandler) Create(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Code     string `json:"code" binding:"required"`
		Location string `json:"location"`
		State    string `json:"state"`
		Country  string `json:"country"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	campus := &models.Campus{
		Name:     input.Name,
		Code:     input.Code,
		Location: input.Location,
		State:    input.State,
		Country:  input.Country,
		IsActive: true,
	}

	if err := h.db.Create(campus).Error; err != nil {
		response.InternalError(c, "Failed to create campus")
		return
	}

	response.Created(c, "Campus created successfully", campus)
}

// Update godoc
// @Summary Update a campus (super admin only)
// @Tags campuses
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Campus ID"
// @Param input body object{name=string,code=string,location=string,state=string,country=string,is_active=bool} true "Campus details"
// @Success 200 {object} response.Response
// @Router /admin/campuses/{id} [put]
func (h *CampusHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var campus models.Campus
	if err := h.db.First(&campus, "id = ?", id).Error; err != nil {
		response.NotFound(c, "Campus not found")
		return
	}

	var input struct {
		Name     string `json:"name"`
		Code     string `json:"code"`
		Location string `json:"location"`
		State    string `json:"state"`
		Country  string `json:"country"`
		IsActive *bool  `json:"is_active"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if input.Name != "" {
		campus.Name = input.Name
	}
	if input.Code != "" {
		campus.Code = input.Code
	}
	if input.Location != "" {
		campus.Location = input.Location
	}
	if input.State != "" {
		campus.State = input.State
	}
	if input.Country != "" {
		campus.Country = input.Country
	}
	if input.IsActive != nil {
		campus.IsActive = *input.IsActive
	}

	if err := h.db.Save(&campus).Error; err != nil {
		response.InternalError(c, "Failed to update campus")
		return
	}

	response.Success(c, "Campus updated successfully", campus)
}

