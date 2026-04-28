package repository

import (
	"context"
	"hta-platform/internal/image/domain/model/entity"
	"hta-platform/internal/image/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type imageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) repository.ImageRepository {
	return &imageRepository{db: db}
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
