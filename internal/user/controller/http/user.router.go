package http

import (
	"github.com/gin-gonic/gin"
	"github.com/leedev/go-rest-ddd/pkg/response"
)

func RegisterUserRoutes(rg *gin.RouterGroup, handler *UserHandler) {
	auth := rg.Group("/auth")
	auth.POST("/profile", response.Wrap(handler.GetUserProfile))
}
