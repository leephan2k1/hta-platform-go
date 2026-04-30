package service

import (
	"context"
	"hta-platform/global"
	"hta-platform/internal/image/controller/dto"
	imageEntity "hta-platform/internal/image/domain/model/entity"
	"hta-platform/internal/image/domain/repository"
	mediaDto "hta-platform/internal/media/controller/dto"
	mediaRepo "hta-platform/internal/media/domain/repository"
	"io"
	"sync"
	"go.uber.org/zap"
)

type imageService struct {
	imgRepo   repository.ImageRepository
	mediaRepo mediaRepo.MediaRepository
}

// StreamImage implements [ImageService].
func (i *imageService) StreamImage(req *dto.StreamImageReq) (io.ReadCloser, error) {
	return i.imgRepo.StreamImage(context.Background(), req.URL, req.Source)
}

// MigrateThumbnail implements [ImageService].
func (i *imageService) MigrateThumbnail(ctx context.Context, req *dto.MigrateThumbnailReq) error {
	const limit = 5000
	global.Logger.Info("Starting thumbnail migration", zap.String("source", req.Source), zap.Int("limit", limit))

	// We need to handle both NSFW and non-NSFW media
	for _, nsfw := range []bool{false, true} {
		global.Logger.Info("Processing media batch", zap.Bool("is_nsfw", nsfw))

		// 1. Get first batch to determine total for this NSFW status
		firstReq := &mediaDto.GetMediasReq{
			Limit:  limit,
			Page:   1,
			IsNSFW: nsfw,
		}
		_, total, err := i.mediaRepo.GetMedias(ctx, firstReq)
		if err != nil {
			global.Logger.Error("Failed to get media total", zap.Error(err), zap.Bool("is_nsfw", nsfw))
			return err
		}

		if total == 0 {
			global.Logger.Info("No media found for status", zap.Bool("is_nsfw", nsfw))
			continue
		}

		totalPages := int((total + int64(limit) - 1) / int64(limit))
		global.Logger.Info("Migration details", zap.Int64("total_records", total), zap.Int("total_pages", totalPages), zap.Bool("is_nsfw", nsfw))

		var wg sync.WaitGroup
		errChan := make(chan error, totalPages)

		for page := 1; page <= totalPages; page++ {
			wg.Add(1)
			go func(p int, isNSFW bool) {
				defer wg.Done()

				global.Logger.Debug("Processing batch", zap.Int("page", p), zap.Bool("is_nsfw", isNSFW))

				// Fetch batch
				batchReq := &mediaDto.GetMediasReq{
					Limit:  limit,
					Page:   p,
					IsNSFW: isNSFW,
				}
				medias, _, err := i.mediaRepo.GetMedias(ctx, batchReq)
				if err != nil {
					global.Logger.Error("Failed to fetch batch", zap.Int("page", p), zap.Error(err))
					errChan <- err
					return
				}

				// Prepare entities
				var images []*imageEntity.Image
				for _, m := range medias {
					if m.Thumbnail != "" {
						images = append(images, &imageEntity.Image{
							URL:         m.Thumbnail,
							Description: req.Description,
							ResourceID:  m.ID,
							Source:      req.Source,
						})
					}
				}

				// Bulk insert
				if len(images) > 0 {
					global.Logger.Debug("Bulk inserting images", zap.Int("count", len(images)), zap.Int("page", p))
					_, err = i.imgRepo.CreateImages(ctx, images)
					if err != nil {
						global.Logger.Error("Failed to bulk insert images", zap.Int("page", p), zap.Error(err))
						errChan <- err
						return
					}
				}
			}(page, nsfw)
		}

		wg.Wait()
		close(errChan)

		// Check for errors in this pass
		for e := range errChan {
			if e != nil {
				return e
			}
		}
	}

	global.Logger.Info("Thumbnail migration completed successfully")
	return nil
}

func NewImageService(imgRepo repository.ImageRepository, mediaRepo mediaRepo.MediaRepository) ImageService {
	return &imageService{imgRepo: imgRepo, mediaRepo: mediaRepo}
}
