package http

import (
	"fmt"

	"hta-platform/internal/user/application/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (ah *UserHandler) GetUserProfile(ctx *gin.Context) (res interface{}, err error) {
	fmt.Println("---> GetUserProfile")

	return nil, nil
}
