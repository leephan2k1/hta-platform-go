package initialize

import (
	"hta-platform/internal/user/application/service"
	"hta-platform/internal/user/controller/http"

	userRepo "hta-platform/internal/user/infrastructure/persistence/repository"

	"gorm.io/gorm"
)

// initializes service, repository, and handler for user
func InitUser(db *gorm.DB) *http.UserHandler {
	repository := userRepo.NewUserRepository(db)
	service := service.NewUserService(repository)
	handler := http.NewUserHandler(service)
	return handler
}
