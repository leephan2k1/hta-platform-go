package http

import (
	"hta-platform/internal/image/application/service"
	"hta-platform/internal/image/controller/dto"
	"hta-platform/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	imgService service.ImageService
}

func NewImageHandler(imgService service.ImageService) *ImageHandler {
	return &ImageHandler{imgService: imgService}
}

func (h *ImageHandler) MigrateThumbnail(c *gin.Context) {
	var req dto.MigrateThumbnailReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.imgService.MigrateThumbnail(c.Request.Context(), &req); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Thumbnail migration completed successfully"})
}

func (h *ImageHandler) StreamImage(c *gin.Context) {
	var req dto.StreamImageReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	reader, err := h.imgService.StreamImage(&req)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, "Failed to stream image", err.Error())
		return
	}
	defer reader.Close()

	c.Header("Content-Type", "image/jpeg") // Default, could be dynamic
	c.DataFromReader(http.StatusOK, -1, "image/jpeg", reader, nil)
}
