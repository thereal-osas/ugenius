package handlers

import (
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/pkg/response"
	"gorm.io/gorm"
)

type GalleryHandler struct {
	db *gorm.DB
}

func NewGalleryHandler(db *gorm.DB) *GalleryHandler {
	return &GalleryHandler{db: db}
}

type CreateGalleryInput struct {
	Title    string `json:"title" binding:"required"`
	ImageURL string `json:"image_url" binding:"required"`
	Caption  string `json:"caption"`
}

// CreateGallery godoc
// @Summary Create new gallery image (Super Admin only)
// @Tags gallery
// @Accept json
// @Produce json
// @Param input body CreateGalleryInput true "Gallery data"
// @Success 200 {object} response.Response
// @Router /admin/gallery [post]
func (h *GalleryHandler) CreateGallery(c *gin.Context) {
	var input CreateGalleryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid input: "+err.Error())
		return
	}

	gallery := models.Gallery{
		Title:     input.Title,
		ImageURL:  input.ImageURL,
		Caption:   input.Caption,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.db.Create(&gallery).Error; err != nil {
		response.InternalError(c, "Failed to create gallery item: "+err.Error())
		return
	}

	response.Success(c, "Gallery item created successfully", gallery)
}

// GetGallery godoc
// @Summary Get all gallery images
// @Tags gallery
// @Produce json
// @Success 200 {object} response.Response
// @Router /gallery [get]
func (h *GalleryHandler) GetGallery(c *gin.Context) {
	var galleries []models.Gallery

	if err := h.db.Order("created_at DESC").Find(&galleries).Error; err != nil {
		response.InternalError(c, "Failed to fetch gallery: "+err.Error())
		return
	}

	response.Success(c, "Gallery fetched successfully", galleries)
}

// DeleteGallery godoc
// @Summary Delete gallery image (Super Admin only)
// @Tags gallery
// @Produce json
// @Param id path int true "Gallery ID"
// @Success 200 {object} response.Response
// @Router /admin/gallery/:id [delete]
func (h *GalleryHandler) DeleteGallery(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid gallery ID")
		return
	}

	if err := h.db.Delete(&models.Gallery{}, id).Error; err != nil {
		response.InternalError(c, "Failed to delete gallery item: "+err.Error())
		return
	}

	response.Success(c, "Gallery item deleted successfully", nil)
}

// UploadGalleryImage godoc
// @Summary Upload gallery image (Super Admin only)
// @Tags gallery
// @Accept multipart/form-data
// @Produce json
// @Param image formData file true "Image file"
// @Success 200 {object} response.Response
// @Router /admin/gallery/upload [post]
func (h *GalleryHandler) UploadGalleryImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		response.BadRequest(c, "No file uploaded")
		return
	}

	// Generate unique filename
	filename := filepath.Base(file.Filename)
	uniqueFilename := strconv.FormatInt(time.Now().Unix(), 10) + "_" + filename

	// Save file to uploads directory
	uploadPath := "uploads/gallery/"
	if err := c.SaveUploadedFile(file, uploadPath+uniqueFilename); err != nil {
		response.InternalError(c, "Failed to save file: "+err.Error())
		return
	}

	imageURL := "/uploads/gallery/" + uniqueFilename
	response.Success(c, "Image uploaded successfully", map[string]string{
		"image_url": imageURL,
	})
}
