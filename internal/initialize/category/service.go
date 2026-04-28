package initialize

import (
	"hta-platform/internal/category/application/service"
	"hta-platform/internal/category/controller/http"
	"hta-platform/internal/category/infrastructure/persistence/repository"

	"gorm.io/gorm"
)

func InitCategory(db *gorm.DB) *http.CategoryHandler {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	handler := http.NewCategoryHandler(categoryService)
	return handler
}
