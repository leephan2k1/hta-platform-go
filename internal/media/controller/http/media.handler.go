package http

import (
	"hta-platform/internal/media/application/service"
	"hta-platform/internal/media/controller/dto"
	"hta-platform/pkg/response"
	"hta-platform/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
)

type MediaHandler struct {
	mediaService service.MediaService
}

func NewMediaHandler(mediaService service.MediaService) *MediaHandler {
	return &MediaHandler{mediaService: mediaService}
}

func (h *MediaHandler) GetMediaByUrl(c *gin.Context) (interface{}, error) {
	url := c.Param("url")

	media, err := h.mediaService.GetMediaByUrl(c, url)
	if err != nil {
		return nil, err
	}

	return media, nil
}

func (h *MediaHandler) GetMedias(c *gin.Context) (interface{}, error) {
	var req dto.GetMediasReq

	if err := c.ShouldBindQuery(&req); err != nil {
		return nil, err
	}
	req.Normalize()

	validation, exists := c.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Invalid request", "Validation not found in context")
	}

	apiErr := utils.ValidateStruct(&req, validation.(*validator.Validate))
	if apiErr != nil {
		return nil, apiErr
	}

	medias, err := h.mediaService.GetMedias(c, &req)
	if err != nil {
		return nil, err
	}

	return medias, nil
}

func (h *MediaHandler) GenerateSlug(c *gin.Context) (interface{}, error) {
	var req dto.GenerateSlugReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		return nil, err
	}
	req.Normalize()

	validation, exists := c.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Invalid request", "Validation not found in context")
	}

	apiErr := utils.ValidateStruct(&req, validation.(*validator.Validate))
	if apiErr != nil {
		return nil, apiErr
	}

	slugVal := slug.Make(strings.ToLower(*req.Name))
	return slugVal, nil
}

func (h *MediaHandler) CreateMedia(c *gin.Context) (interface{}, error) {
	var req dto.CreateMediaReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		return nil, err
	}
	req.Normalize()

	validation, exists := c.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Invalid request", "Validation not found in context")
	}

	apiErr := utils.ValidateStruct(&req, validation.(*validator.Validate))
	if apiErr != nil {
		return nil, apiErr
	}

	newMedia, err := h.mediaService.CreateMedia(c, &req)
	if err != nil {
		return nil, err
	}

	return newMedia, nil
}

func (h *MediaHandler) UpdateMedia(c *gin.Context) (interface{}, error) {
	var req dto.CreateMediaReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		return nil, err
	}
	req.Normalize()

	validation, exists := c.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Invalid request", "Validation not found in context")
	}

	apiErr := utils.ValidateStruct(&req, validation.(*validator.Validate))
	if apiErr != nil {
		return nil, apiErr
	}

	url := slug.Make(strings.ToLower(*req.Name))
	updatedMedia, err := h.mediaService.UpdateMedia(c, url, &req)
	if err != nil {
		return nil, err
	}

	return updatedMedia, nil
}
