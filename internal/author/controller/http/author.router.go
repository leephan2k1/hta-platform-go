package http

import (
	"hta-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterAuthorRoutes(rg *gin.RouterGroup, handler *AuthorHandler) {
	a := rg.Group("/authors")

	a.POST("/", response.Wrap(handler.CreateAuthor))

	a.GET("/:url", response.Wrap(handler.GetAuthorByUrl))
}
