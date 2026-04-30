package http

import (
	"hta-platform/internal/middleware"
	"hta-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterAuthorRoutes(rg *gin.RouterGroup, handler *AuthorHandler) {
	a := rg.Group("/authors")

	a.Use(middleware.Auth0Guard())
	a.Use(middleware.RolesGuard([]string{"MEMBER"}))

	a.POST("/", response.Wrap(handler.CreateAuthor))

	a.GET("/:url", response.Wrap(handler.GetAuthorByUrl))

	a.GET("", response.Wrap(handler.GetAuthors))
}
