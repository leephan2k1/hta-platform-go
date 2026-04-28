package initialize

import (
	"hta-platform/internal/media/application/service"
	"hta-platform/internal/media/controller/http"
	"hta-platform/internal/media/infrastructure/persistence/repository"

	"gorm.io/gorm"
)

func InitMediaChapter(db *gorm.DB) *http.MediaChapterHandler {
	repo := repository.NewMediaChapterRepository(db)
	mediaRepo := repository.NewMediaRepository(db)
	svc := service.NewMediaChapterServiceImpl(repo, mediaRepo)
	handler := http.NewMediaChapterHandler(svc)
	return handler
}
