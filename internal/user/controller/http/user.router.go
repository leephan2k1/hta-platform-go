package http

import (
	"hta-platform/pkg/response"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, handler *UserHandler) {
	auth := rg.Group("/auth")

	auth.POST("/register", response.Wrap(handler.RegisterUser))
}
