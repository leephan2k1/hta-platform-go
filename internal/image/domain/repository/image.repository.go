package repository

import (
	"context"
	"hta-platform/internal/image/domain/model/entity"
)

type ImageRepository interface {
	CreateImages(ctx context.Context, images []*entity.Image) ([]entity.Image, error)

	GetImagesByResourceId(ctx context.Context, resourceId string) ([]entity.Image, error)
}
