package service

import (
	"context"
	"hta-platform/internal/category/controller/dto"
	"hta-platform/internal/category/domain/model/entity"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, req *dto.CategoryReq) (entity.Category, error)

	GetAllCategories(ctx context.Context) ([]entity.Category, error)
}
