package http

import (
	"hta-platform/internal/media/application/service"
	"hta-platform/internal/media/controller/dto"
	"hta-platform/pkg/response"
	"hta-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MediaChapterHandler struct {
	service service.MediaChapterService
}

func (h *MediaChapterHandler) GetMediaChaptersByMediaUrl(c *gin.Context) (interface{}, error) {
	url := c.Param("media-url")

	return h.service.GetMediaChaptersByMediaUrl(c, url)
}

func (h *MediaChapterHandler) GetChapterImagesByChapterUrl(c *gin.Context) (interface{}, error) {
	url := c.Param("chapter-url")

	return h.service.GetChapterImagesByChapterUrl(c, url)
}

func (h *MediaChapterHandler) CreateMediaChapters(c *gin.Context) (interface{}, error) {
	var req dto.CreateMediaChapterReq

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	validation, exists := c.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Invalid request", "Validation not found in context")
	}

	apiErr := utils.ValidateStruct(&req, validation.(*validator.Validate))
	if apiErr != nil {
		return nil, apiErr
	}

	return h.service.CreateMediaChapters(c, &req)
}

func NewMediaChapterHandler(service service.MediaChapterService) *MediaChapterHandler {
	return &MediaChapterHandler{service: service}
}
