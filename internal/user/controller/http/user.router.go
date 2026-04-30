package http

import (
	"hta-platform/internal/middleware"
	"hta-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, handler *UserHandler) {
	auth := rg.Group("/auth")

	auth.POST("/register", response.Wrap(handler.RegisterUser))

	privateUser := rg.Group("/users")
	privateUser.Use(middleware.Auth0Guard())

	privateUser.GET("/authors", response.Wrap(handler.GetBookmarkedAuthors))
	privateUser.POST("/authors", response.Wrap(handler.BookmarkAuthor))
	privateUser.DELETE("/authors", response.Wrap(handler.UnbookmarkAuthor))

	privateUser.GET("/medias", response.Wrap(handler.GetBookmarkedMedias))
	privateUser.POST("/medias", response.Wrap(handler.BookmarkMedia))
	privateUser.DELETE("/medias", response.Wrap(handler.UnbookmarkMedia))
}
