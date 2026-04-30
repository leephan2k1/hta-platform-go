package http

import (
	"hta-platform/internal/middleware"
	"hta-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, handler *CategoryHandler) {
	c := rg.Group("/categories")

	c.Use(middleware.Auth0Guard())
	c.Use(middleware.RolesGuard([]string{"MEMBER"}))

	c.POST("", response.Wrap(handler.CreateCategory))

	c.GET("", response.Wrap(handler.GetAllCategories))
}
