package initialize

import (
	"hta-platform/internal/author/application/service"
	"hta-platform/internal/author/controller/http"

	authRepo "hta-platform/internal/author/infrastructure/persistence/repository"

	"gorm.io/gorm"
)

// initializes service, repository, and handler for auth
func InitAuthor(db *gorm.DB) *http.AuthorHandler {
	authorRepo := authRepo.NewAuthorRepository(db)
	service := service.NewAuthorService(authorRepo)
	handler := http.NewAuthorHandler(service)
	return handler
}
