package image

import (
	"hta-platform/global"
	"hta-platform/internal/image/application/service"
	"hta-platform/internal/image/controller/http"
	persistence "hta-platform/internal/image/infrastructure/persistence/repository"
	"hta-platform/internal/image/infrastructure/streamer"

	"gorm.io/gorm"
)

func InitImage(db *gorm.DB) *http.ImageHandler {
	streamers := map[string]streamer.ImageStreamer{
		"MM":  streamer.NewMMStreamer(global.ConfigValue.MMReferer),
		"HTA": streamer.NewHTAStreamer(),
	}

	repo := persistence.NewImageRepository(db, streamers)
	svc := service.NewImageService(repo)
	return http.NewImageHandler(svc)
}
