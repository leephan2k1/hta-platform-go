package service

import (
	"context"
	"hta-platform/internal/category/controller/dto"
	"hta-platform/internal/category/domain/model/enity"
	"hta-platform/internal/category/domain/repository"

	"github.com/gosimple/slug"
)

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

// CreateCategory implements [CategoryService].
func (s *categoryService) CreateCategory(ctx context.Context, req *dto.CategoryReq) (enity.Category, error) {
	category := enity.Category{
		Name: req.Name,
		Slug: slug.Make(req.Name),
	}
	return s.categoryRepo.CreateCategory(ctx, &category)
}

// GetAllCategories implements [CategoryService].
func (s *categoryService) GetAllCategories(ctx context.Context) ([]enity.Category, error) {
	return s.categoryRepo.FindAllCategories(ctx)
}
