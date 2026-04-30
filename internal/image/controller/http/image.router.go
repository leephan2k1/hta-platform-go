package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterImageRoutes(rg *gin.RouterGroup, handler *ImageHandler) {
	i := rg.Group("/images")

	i.GET("/stream", handler.StreamImage)

	i.GET("/migrate-thumbnail", handler.MigrateThumbnail)
}
