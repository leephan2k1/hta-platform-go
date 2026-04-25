package service

import (
	"context"
	"hta-platform/internal/category/controller/dto"
	"hta-platform/internal/category/domain/model/enity"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, req *dto.CategoryReq) (enity.Category, error)

	GetAllCategories(ctx context.Context) ([]enity.Category, error)
}
