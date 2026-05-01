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

func (ah *UserHandler) BookmarkAuthor(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	var req dto.UserToResourceReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	if err := ah.service.BookmarkAuthor(ctx, userID, req.ResourceID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (ah *UserHandler) UnbookmarkAuthor(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	var req dto.UserToResourceReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	if err := ah.service.UnbookmarkAuthor(ctx, userID, req.ResourceID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (ah *UserHandler) GetBookmarkedAuthors(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	return ah.service.GetBookmarkedAuthors(ctx, userID)
}

func (ah *UserHandler) BookmarkMedia(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	var req dto.UserToResourceReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	if err := ah.service.BookmarkMedia(ctx, userID, req.ResourceID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (ah *UserHandler) UnbookmarkMedia(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	var req dto.UserToResourceReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	if err := ah.service.UnbookmarkMedia(ctx, userID, req.ResourceID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (ah *UserHandler) GetBookmarkedMedias(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	return ah.service.GetBookmarkedMedias(ctx, userID)
}

func (ah *UserHandler) IsBookmarkedAuthor(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	resourceID := ctx.Param("resourceId")
	exist, err := ah.service.IsBookmarkedAuthor(ctx, userID, resourceID)
	if err != nil {
		return nil, err
	}
	return gin.H{"exist": exist}, nil
}

func (ah *UserHandler) IsBookmarkedMedia(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	resourceID := ctx.Param("resourceId")
	exist, err := ah.service.IsBookmarkedMedia(ctx, userID, resourceID)
	if err != nil {
		return nil, err
	}
	return gin.H{"exist": exist}, nil
}

func (ah *UserHandler) UpsertReadingProgress(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	var req dto.UserReadingProgressReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	validation, exists := ctx.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Invalid request", "Validation not found in context")
	}

	apiErr := utils.ValidateStruct(&req, validation.(*validator.Validate))
	if apiErr != nil {
		return nil, apiErr
	}

	if err := ah.service.UpsertReadingProgress(ctx, userID, req); err != nil {
		return nil, err
	}
	return nil, nil
}

func (ah *UserHandler) GetReadingProgress(ctx *gin.Context) (res interface{}, err error) {
	userID := ctx.GetString("user_id")
	return ah.service.GetReadingProgress(ctx, userID)
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}
