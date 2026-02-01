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

type ScholarshipsHandler struct {
	db *gorm.DB
}

func NewScholarshipsHandler(db *gorm.DB) *ScholarshipsHandler {
	return &ScholarshipsHandler{db: db}
}

// List godoc
// @Summary List available scholarships
// @Tags scholarships
// @Produce json
// @Success 200 {object} response.Response
// @Router /scholarships [get]
func (h *ScholarshipsHandler) List(c *gin.Context) {
	page := 1
	pageSize := 20
	if p, _ := strconv.Atoi(c.Query("page")); p > 0 {
		page = p
	}
	if ps, _ := strconv.Atoi(c.Query("page_size")); ps > 0 && ps <= 50 {
		pageSize = ps
	}

	query := h.db.Model(&models.Scholarship{}).Where("status = ? AND application_deadline > NOW()", models.ScholarshipStatusOpen)

	var total int64
	query.Count(&total)

	var scholarships []models.Scholarship
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("deadline ASC").Find(&scholarships).Error; err != nil {
		response.InternalError(c, "Failed to fetch scholarships")
		return
	}

	response.Paginated(c, scholarships, page, pageSize, total)
}

// GetByID godoc
// @Summary Get scholarship by ID
// @Tags scholarships
// @Param id path string true "Scholarship ID"
// @Success 200 {object} response.Response
// @Router /scholarships/{id} [get]
func (h *ScholarshipsHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var scholarship models.Scholarship
	if err := h.db.First(&scholarship, "id = ?", id).Error; err != nil {
		response.NotFound(c, "Scholarship not found")
		return
	}

	response.Success(c, "", scholarship)
}

// Apply godoc
// @Summary Apply for a scholarship
// @Tags scholarships
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Scholarship ID"
// @Param input body object{essay=string,supporting_docs=[]string} true "Application details"
// @Success 201 {object} response.Response
// @Router /scholarships/{id}/apply [post]
func (h *ScholarshipsHandler) Apply(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	scholarshipID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var scholarship models.Scholarship
	if err := h.db.First(&scholarship, "id = ?", scholarshipID).Error; err != nil {
		response.NotFound(c, "Scholarship not found")
		return
	}

	// Check if already applied
	var existingApp models.ScholarshipApplication
	if h.db.Where("user_id = ? AND scholarship_id = ?", userID, scholarshipID).First(&existingApp).Error == nil {
		response.BadRequest(c, "Already applied for this scholarship")
		return
	}

	var input struct {
		Statement string `json:"statement"`
		Documents string `json:"documents"` // JSON array of document URLs
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	application := &models.ScholarshipApplication{
		UserID:        userID,
		ScholarshipID: scholarshipID,
		Status:        models.ApplicationPending,
		Statement:     input.Statement,
		Documents:     input.Documents,
	}

	if err := h.db.Create(application).Error; err != nil {
		response.InternalError(c, "Failed to submit application")
		return
	}

	response.Created(c, "Application submitted successfully", application)
}

// MyApplications godoc
// @Summary Get user's scholarship applications
// @Tags scholarships
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response
// @Router /scholarships/my-applications [get]
func (h *ScholarshipsHandler) MyApplications(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	var applications []models.ScholarshipApplication
	if err := h.db.Preload("Scholarship").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&applications).Error; err != nil {
		response.InternalError(c, "Failed to fetch applications")
		return
	}

	response.Success(c, "", applications)
}

