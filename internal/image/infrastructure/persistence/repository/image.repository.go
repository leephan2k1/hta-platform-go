package repository

import (
	"context"
	"fmt"
	"hta-platform/internal/image/domain/model/entity"
	"hta-platform/internal/image/domain/repository"
	"hta-platform/internal/image/infrastructure/streamer"
	"io"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type imageRepository struct {
	db        *gorm.DB
	streamers map[string]streamer.ImageStreamer
}

// StreamImage implements [repository.ImageRepository].
func (r *imageRepository) StreamImage(ctx context.Context, url string, source string) (io.ReadCloser, error) {
	s, ok := r.streamers[source]
	if !ok {
		return nil, fmt.Errorf("unsupported source: %s", source)
	}
	return s.Stream(ctx, url)
}

func (r *imageRepository) CreateImages(ctx context.Context, images []*entity.Image) ([]entity.Image, error) {
	err := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "url"}},
			DoUpdates: clause.AssignmentColumns([]string{"description", "resource_id", "source", "updated_at"}),
		}).Create(&images).Error

	if err != nil {
		return nil, err
	}

	res := make([]entity.Image, len(images))
	for i, img := range images {
		res[i] = *img
	}
	return res, nil
}

func (r *imageRepository) GetImagesByResourceId(ctx context.Context, resourceId string) ([]entity.Image, error) {
	var images []entity.Image
	err := r.db.WithContext(ctx).
		Where("resource_id = ?", resourceId).
		Find(&images).Error

	return images, err
}

func NewImageRepository(db *gorm.DB, streamers map[string]streamer.ImageStreamer) repository.ImageRepository {
	return &imageRepository{db: db, streamers: streamers}
}
