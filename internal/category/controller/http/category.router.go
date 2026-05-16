package http

import (
	"hta-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, handler *CategoryHandler) {
	c := rg.Group("/categories")

	c.POST("", response.Wrap(handler.CreateCategory))

	c.GET("", response.Wrap(handler.GetAllCategories))
}
