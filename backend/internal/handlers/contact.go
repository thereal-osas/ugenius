package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ugenius/backend/internal/services"
	"github.com/ugenius/backend/pkg/response"
)

type ContactHandler struct {
	emailService *services.EmailService
}

func NewContactHandler(emailService *services.EmailService) *ContactHandler {
	return &ContactHandler{
		emailService: emailService,
	}
}

type ContactRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// SendContactEmail godoc
// @Summary Send contact form email
// @Tags contact
// @Accept json
// @Produce json
// @Param input body ContactRequest true "Contact form data"
// @Success 200 {object} response.Response
// @Router /contact [post]
func (h *ContactHandler) SendContactEmail(c *gin.Context) {
	var req ContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	// Send email
	if err := h.emailService.SendContactEmail(req.Name, req.Email, req.Subject, req.Message); err != nil {
		response.InternalError(c, "Failed to send email: "+err.Error())
		return
	}

	response.Success(c, "Email sent successfully", nil)
}
