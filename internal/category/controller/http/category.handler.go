package http

import (
	"hta-platform/internal/category/application/service"
	"hta-platform/internal/category/controller/dto"
	"hta-platform/pkg/response"
	"hta-platform/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) (res interface{}, err error) {
	var req dto.CategoryReq

	err = c.ShouldBindJSON(&req)
	if err != nil {
		return nil, err
	}
	req.Normalize()

	validation, exists := c.Get("validation")
	if !exists {
		return nil, response.NewAPIError(http.StatusInternalServerError, "Invalid request", "Validation not found in context")
	}

	apiErr := utils.ValidateStruct(&req, validation.(*validator.Validate))
	if apiErr != nil {
		return nil, apiErr
	}

	newCategory, err := h.categoryService.CreateCategory(c, &req)
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) (interface{}, error) {
	categories, err := h.categoryService.GetAllCategories(c)
	if err != nil {
		return nil, err
	}

	res := make([]dto.CategoryRes, len(categories))
	for i, cat := range categories {
		res[i] = dto.CategoryRes{
			ID:   cat.ID.String(),
			Name: cat.Name,
			Slug: cat.Slug,
		}
	}

	return res, nil
}
