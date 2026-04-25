package repository

import (
	"context"
	"hta-platform/internal/category/domain/model/enity"
	"hta-platform/internal/category/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type categoryRepository struct {
	DB *gorm.DB
}

// CreateCategory implements [repository.CategoryRepository].
func (c *categoryRepository) CreateCategory(ctx context.Context, category *enity.Category) (enity.Category, error) {
	result := c.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "slug"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "slug", "updated_at"}),
	}).Create(category)

	if result.Error != nil {
		return enity.Category{}, result.Error
	}
	return *category, nil
}

// FindAllCategories implements [repository.CategoryRepository].
func (c *categoryRepository) FindAllCategories(ctx context.Context) ([]enity.Category, error) {
	var categories []enity.Category
	result := c.DB.Select("id", "name", "slug").Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepository{db}
}
