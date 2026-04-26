package service

import (
	"context"
	"hta-platform/internal/media/controller/dto"
	"hta-platform/internal/media/domain/model/entity"
)

type MediaService interface {
	CreateMedia(ctx context.Context, req *dto.MediaReq) (entity.Media, error)

	UpdateMedia(ctx context.Context, url string, req *dto.MediaReq) (entity.Media, error)
}
