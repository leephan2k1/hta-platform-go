package repository

import (
	"context"
	"hta-platform/internal/category/domain/model/enity"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *enity.Category) (enity.Category, error)

	FindAllCategories(ctx context.Context) ([]enity.Category, error)
}
