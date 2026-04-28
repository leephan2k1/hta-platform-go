package repository

import (
	"context"

	"hta-platform/internal/category/domain/model/entity"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *entity.Category) (entity.Category, error)

	FindAllCategories(ctx context.Context) ([]entity.Category, error)
}
