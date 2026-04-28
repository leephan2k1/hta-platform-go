package http

import (
	"hta-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterMediaChapterRoutes(rg *gin.RouterGroup, handler *MediaChapterHandler) {
	mc := rg.Group("/chapters")

	mc.GET("/by-media/:media-url", response.Wrap(handler.GetMediaChaptersByMediaUrl))

	mc.GET("/:chapter-url", response.Wrap(handler.GetChapterImagesByChapterUrl))
}
