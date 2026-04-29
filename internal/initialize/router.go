package initialize

import (
	authorHttp "hta-platform/internal/author/controller/http"
	categoryHttp "hta-platform/internal/category/controller/http"
	initializeAuthor "hta-platform/internal/initialize/author"
	initializeCategory "hta-platform/internal/initialize/category"
	initializeMedia "hta-platform/internal/initialize/media"
	initializeMediaChapter "hta-platform/internal/initialize/media_chapter"
	mediaChapterHttp "hta-platform/internal/media/controller/http"
	mediaHttp "hta-platform/internal/media/controller/http"
	"hta-platform/internal/middleware"
	"hta-platform/pkg/response"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, isLogger string) *gin.Engine {
	// Initialize the router
	// This function will set up the routes and middleware for the application
	// It will return a gin.Engine instance that can be used to run the server

	var r *gin.Engine
	// Set the mode based on the environment
	if isLogger == "debug" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	r.Use(middleware.CORS) // cross
	r.Use(middleware.ValidatorMiddleware())
	// r.Use() // logging

	// r.Use() // limiter global
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	// r.Use(middleware.Validator())      // middleware

	// r.Use(middlewares.NewRateLimiter().GlobalRateLimiter()) // 100 req/s
	r.GET("/ping/100", func(ctx *gin.Context) {
		response.SuccessResponse(ctx, "pong")
	})

	r.GET("/ping/200", response.Wrap(func(ctx *gin.Context) (res interface{}, err error) {
		return "pong", nil
	}))

	// === register routes theo module
	v1 := r.Group("/v1")

	// Register the auth routes
	// === DI các handler
	authorHandler := initializeAuthor.InitAuthor(db)
	authorHttp.RegisterAuthorRoutes(v1, authorHandler)

	categoryHandler := initializeCategory.InitCategory(db)
	categoryHttp.RegisterCategoryRoutes(v1, categoryHandler)

	mediaHandler := initializeMedia.InitMedia(db)
	mediaHttp.RegisterMediaRoutes(v1, mediaHandler)

	mediaChapterHandler := initializeMediaChapter.InitMediaChapter(db)
	mediaChapterHttp.RegisterMediaChapterRoutes(v1, mediaChapterHandler)

	return r
}
