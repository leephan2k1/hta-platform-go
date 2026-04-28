package http

import (
	"hta-platform/internal/media/application/service"

	"github.com/gin-gonic/gin"
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

func NewMediaChapterHandler(service service.MediaChapterService) *MediaChapterHandler {
	return &MediaChapterHandler{service: service}
}
