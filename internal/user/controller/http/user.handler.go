package http

import (
	"hta-platform/internal/user/application/service"
	"hta-platform/internal/user/controller/dto"
	"hta-platform/pkg/response"
	"hta-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service service.UserService
}

func (ah *UserHandler) RegisterUser(ctx *gin.Context) (res interface{}, err error) {
	var req dto.RegisterUserReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	req.Normalize()

	validation, exists := ctx.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Invalid request", "Validation not found in context")
	}

	apiErr := utils.ValidateStruct(&req, validation.(*validator.Validate))
	if apiErr != nil {
		return nil, apiErr
	}

	if err := ah.service.RegisterUser(ctx, req); err != nil {
		return nil, err
	}
	return nil, nil
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}
