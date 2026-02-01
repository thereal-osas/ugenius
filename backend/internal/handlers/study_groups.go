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

type StudyGroupsHandler struct {
	db *gorm.DB
}

func NewStudyGroupsHandler(db *gorm.DB) *StudyGroupsHandler {
	return &StudyGroupsHandler{db: db}
}

// Create godoc
// @Summary Create a new study group
// @Tags study-groups
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body object{name=string,description=string,subject=string,max_members=int} true "Group details"
// @Success 201 {object} response.Response
// @Router /study-groups [post]
func (h *StudyGroupsHandler) Create(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	campusID, _ := middleware.GetCampusID(c)

	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Subject     string `json:"subject"`
		MaxMembers  int    `json:"max_members"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if input.MaxMembers == 0 {
		input.MaxMembers = 10
	}

	group := &models.StudyGroup{
		Name:        input.Name,
		Description: input.Description,
		Subject:     input.Subject,
		CampusID:    campusID,
		CreatedByID: userID,
		MaxMembers:  input.MaxMembers,
		IsActive:    true,
	}

	if err := h.db.Create(group).Error; err != nil {
		response.InternalError(c, "Failed to create study group")
		return
	}

	// Add creator as leader
	member := &models.StudyGroupMember{
		GroupID: group.ID,
		UserID:  userID,
		Role:    models.MemberRoleLeader,
	}
	h.db.Create(member)

	response.Created(c, "Study group created successfully", group)
}

// List godoc
// @Summary List study groups
// @Tags study-groups
// @Security BearerAuth
// @Produce json
// @Param campus_id query string false "Filter by campus"
// @Param subject query string false "Filter by subject"
// @Success 200 {object} response.Response
// @Router /study-groups [get]
func (h *StudyGroupsHandler) List(c *gin.Context) {
	page := 1
	pageSize := 20
	if p, _ := strconv.Atoi(c.Query("page")); p > 0 {
		page = p
	}
	if ps, _ := strconv.Atoi(c.Query("page_size")); ps > 0 && ps <= 50 {
		pageSize = ps
	}

	query := h.db.Model(&models.StudyGroup{}).Where("is_active = ?", true)

	if campusID := c.Query("campus_id"); campusID != "" {
		if id, err := uuid.Parse(campusID); err == nil {
			query = query.Where("campus_id = ?", id)
		}
	}
	if subject := c.Query("subject"); subject != "" {
		query = query.Where("subject ILIKE ?", "%"+subject+"%")
	}

	var total int64
	query.Count(&total)

	var groups []models.StudyGroup
	offset := (page - 1) * pageSize
	if err := query.Preload("Members").Preload("Members.User").
		Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&groups).Error; err != nil {
		response.InternalError(c, "Failed to fetch study groups")
		return
	}

	response.Paginated(c, groups, page, pageSize, total)
}

// GetByID godoc
// @Summary Get study group by ID
// @Tags study-groups
// @Security BearerAuth
// @Param id path string true "Group ID"
// @Success 200 {object} response.Response
// @Router /study-groups/{id} [get]
func (h *StudyGroupsHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var group models.StudyGroup
	if err := h.db.Preload("Members").Preload("Members.User").First(&group, "id = ?", id).Error; err != nil {
		response.NotFound(c, "Study group not found")
		return
	}

	response.Success(c, "", group)
}

// Join godoc
// @Summary Join a study group
// @Tags study-groups
// @Security BearerAuth
// @Param id path string true "Group ID"
// @Success 200 {object} response.Response
// @Router /study-groups/{id}/join [post]
func (h *StudyGroupsHandler) Join(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	groupID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	var group models.StudyGroup
	if err := h.db.Preload("Members").First(&group, "id = ?", groupID).Error; err != nil {
		response.NotFound(c, "Study group not found")
		return
	}

	// Check if already a member
	for _, m := range group.Members {
		if m.UserID == userID {
			response.BadRequest(c, "Already a member of this group")
			return
		}
	}

	// Check max members
	if len(group.Members) >= group.MaxMembers {
		response.BadRequest(c, "Group is full")
		return
	}

	member := &models.StudyGroupMember{
		GroupID: groupID,
		UserID:  userID,
		Role:    models.MemberRoleMember,
	}
	if err := h.db.Create(member).Error; err != nil {
		response.InternalError(c, "Failed to join group")
		return
	}

	response.Success(c, "Joined group successfully", nil)
}

// Leave godoc
// @Summary Leave a study group
// @Tags study-groups
// @Security BearerAuth
// @Param id path string true "Group ID"
// @Success 200 {object} response.Response
// @Router /study-groups/{id}/leave [post]
func (h *StudyGroupsHandler) Leave(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	groupID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID")
		return
	}

	result := h.db.Where("group_id = ? AND user_id = ?", groupID, userID).Delete(&models.StudyGroupMember{})
	if result.RowsAffected == 0 {
		response.NotFound(c, "Not a member of this group")
		return
	}

	response.Success(c, "Left group successfully", nil)
}

