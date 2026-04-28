package http

import (
	"hta-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterMediaRoutes(rg *gin.RouterGroup, handler *MediaHandler) {
	m := rg.Group("/medias")

	m.GET("/:url", response.Wrap(handler.GetMediaByUrl))

	m.GET("", response.Wrap(handler.GetMedias))

	m.POST("", response.Wrap(handler.CreateMedia))

	m.PATCH("/:url", response.Wrap(handler.UpdateMedia))
}
