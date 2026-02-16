package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/middleware"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/internal/services"
	"github.com/ugenius/backend/pkg/response"
	"gorm.io/gorm"
)

type AdminHandler struct {
	db          *gorm.DB
	authService *services.AuthService
}

func NewAdminHandler(db *gorm.DB, authService *services.AuthService) *AdminHandler {
	return &AdminHandler{
		db:          db,
		authService: authService,
	}
}

type CreateAdminInput struct {
	Email     string    `json:"email" binding:"required,email"`
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Password  string    `json:"password" binding:"required,min=8"`
	CampusID  uuid.UUID `json:"campus_id" binding:"required"`
}

// CreateAdmin godoc
// @Summary      Create a new campus admin
// @Description  Create a new campus admin account. This is only accessible to super admins.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        input body CreateAdminInput true "Admin creation data"
// @Success      201 {object} response.SuccessResponse{data=models.User}
// @Failure      400 {object} response.ErrorResponse
// @Failure      401 {object} response.ErrorResponse
// @Failure      403 {object} response.ErrorResponse
// @Failure      409 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /admin/users [post]
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var input CreateAdminInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	// Check if user with this email already exists
	_, err := h.authService.FindUserByEmail(input.Email)
	if err == nil {
		response.Conflict(c, "User with this email already exists")
		return
	}

	// Create user
	user, err := h.authService.CreateUser(services.CreateUserInput{
		Email:     input.Email,
		Password:  input.Password,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		CampusID:  &input.CampusID,
		Role:      models.RoleCampusAdmin,
	})

	if err != nil {
		response.InternalError(c, "Failed to create admin user: "+err.Error())
		return
	}

	response.Created(c, "Admin created successfully", user)
}

// GetCampusUsers godoc
// @Summary      Get all users for the admin's campus
// @Description  Get a list of all users belonging to the campus of the authenticated admin. Super admins can see all users.
// @Tags         Admin
// @Produce      json
// @Success      200 {object} response.SuccessResponse{data=[]models.User}
// @Failure      401 {object} response.ErrorResponse
// @Failure      403 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /admin/users [get]
func (h *AdminHandler) GetCampusUsers(c *gin.Context) {
	userRole, exists := middleware.GetUserRole(c)
	if !exists {
		response.Forbidden(c, "User role not found.")
		return
	}

	var users []models.User
	var err error

	if userRole == models.RoleSuperAdmin {
		// Super admins can see all users
		err = h.db.Preload("Campus").Find(&users).Error
	} else {
		// Campus admins can only see users from their campus
		campusID, exists := middleware.GetCampusID(c)
		if !exists {
			response.Forbidden(c, "You are not associated with a campus.")
			return
		}
		err = h.db.Preload("Campus").Where("campus_id = ?", campusID).Find(&users).Error
	}

	if err != nil {
		response.InternalError(c, "Failed to fetch users: "+err.Error())
		return
	}

	response.Success(c, "Users retrieved successfully", users)
}

type PromoteUserInput struct {
	UserID   uuid.UUID `json:"user_id" binding:"required"`
	CampusID uuid.UUID `json:"campus_id" binding:"required"`
}

// PromoteUser godoc
// @Summary      Promote a student to campus admin
// @Description  Promote a student user to campus admin role. Only accessible to super admins.
// @Tags         Admin
// @Accept       json
// @Produce      json
// @Param        input body PromoteUserInput true "User promotion data"
// @Success      200 {object} response.SuccessResponse{data=models.User}
// @Failure      400 {object} response.ErrorResponse
// @Failure      401 {object} response.ErrorResponse
// @Failure      403 {object} response.ErrorResponse
// @Failure      404 {object} response.ErrorResponse
// @Failure      500 {object} response.ErrorResponse
// @Router       /admin/users/promote [post]
func (h *AdminHandler) PromoteUser(c *gin.Context) {
	var input PromoteUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	// Find the user
	var user models.User
	if err := h.db.First(&user, "id = ?", input.UserID).Error; err != nil {
		response.NotFound(c, "User not found")
		return
	}

	// Check if user is already an admin
	if user.Role == models.RoleCampusAdmin || user.Role == models.RoleSuperAdmin {
		response.BadRequest(c, "User is already an administrator")
		return
	}

	// Update user role and campus assignment
	user.Role = models.RoleCampusAdmin
	user.CampusID = &input.CampusID

	if err := h.db.Save(&user).Error; err != nil {
		response.InternalError(c, "Failed to promote user: "+err.Error())
		return
	}

	// Load campus information
	h.db.Preload("Campus").First(&user, user.ID)

	response.Success(c, "User promoted to campus admin successfully", user)
}
