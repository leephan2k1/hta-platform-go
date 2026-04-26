package initialize

import (
	"hta-platform/internal/media/application/service"
	"hta-platform/internal/media/controller/http"
	mediaRepo "hta-platform/internal/media/infrastructure/persistence/repository"

	"gorm.io/gorm"
)

func InitMedia(db *gorm.DB) *http.MediaHandler {
	repo := mediaRepo.NewMediaRepository(db)
	mediaService := service.NewMediaService(repo, db)
	handler := http.NewMediaHandler(mediaService)
	return handler
}
