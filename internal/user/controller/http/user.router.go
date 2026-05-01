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
	privateUser.GET("/authors/:resourceId", response.Wrap(handler.IsBookmarkedAuthor))
	privateUser.POST("/authors", response.Wrap(handler.BookmarkAuthor))
	privateUser.DELETE("/authors", response.Wrap(handler.UnbookmarkAuthor))

	privateUser.GET("/medias", response.Wrap(handler.GetBookmarkedMedias))
	privateUser.GET("/medias/:resourceId", response.Wrap(handler.IsBookmarkedMedia))
	privateUser.POST("/medias", response.Wrap(handler.BookmarkMedia))
	privateUser.DELETE("/medias", response.Wrap(handler.UnbookmarkMedia))

	privateUser.GET("/reading/progress", response.Wrap(handler.GetReadingProgress))
	privateUser.POST("/reading/progress", response.Wrap(handler.UpsertReadingProgress))

	privateUser.POST("/reading/session/start", response.Wrap(handler.StartReadingSession))
	privateUser.POST("/reading/session/end", response.Wrap(handler.EndReadingSession))
	privateUser.GET("/reading/session", response.Wrap(handler.GetReadingSessions))
}
