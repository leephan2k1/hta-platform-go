package http

import (
	"hta-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterMediaChapterRoutes(rg *gin.RouterGroup, handler *MediaChapterHandler) {
	mc := rg.Group("/chapters")

	mc.GET("/by-media/:media-url", response.Wrap(handler.GetMediaChaptersByMediaUrl))

	mc.GET("/:chapter-url/images", response.Wrap(handler.GetChapterImages))

	mc.POST("", response.Wrap(handler.CreateMediaChapters))

	mc.POST("/images", response.Wrap(handler.CreateChapterImages))
}
