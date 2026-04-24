package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/leedev/go-rest-ddd/internal/user/application/service"
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
