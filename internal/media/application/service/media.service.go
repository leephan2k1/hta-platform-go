package service

import (
	"context"
	"hta-platform/internal/media/controller/dto"
	"hta-platform/internal/media/domain/model/entity"
)

type MediaService interface {
	GetMediaByUrl(ctx context.Context, url string) (dto.MediaResponse, error)

	GetMedias(ctx context.Context, req *dto.GetMediasReq) (dto.GetMediasRes, error)

	CreateMedia(ctx context.Context, req *dto.CreateMediaReq) (entity.Media, error)

	UpdateMedia(ctx context.Context, url string, req *dto.CreateMediaReq) (entity.Media, error)
}
