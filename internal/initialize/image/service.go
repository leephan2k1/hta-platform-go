package image

import (
	"hta-platform/global"
	"hta-platform/internal/image/application/service"
	"hta-platform/internal/image/controller/http"
	persistence "hta-platform/internal/image/infrastructure/persistence/repository"
	"hta-platform/internal/image/infrastructure/streamer"
	mediaRepo "hta-platform/internal/media/infrastructure/persistence/repository"

	"gorm.io/gorm"
)

func InitImage(db *gorm.DB) *http.ImageHandler {
	streamers := map[string]streamer.ImageStreamer{
		"MM":  streamer.NewMMStreamer(global.ConfigValue.MMReferer),
		"HTA": streamer.NewHTAStreamer(),
	}

	repo := persistence.NewImageRepository(db, streamers)
	mediaRepository := mediaRepo.NewMediaRepository(db)
	svc := service.NewImageService(repo, mediaRepository)
	return http.NewImageHandler(svc)
}
