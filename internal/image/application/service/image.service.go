package service

import (
	"context"
	"hta-platform/internal/image/controller/dto"
	"io"
)

type ImageService interface {
	StreamImage(req *dto.StreamImageReq) (io.ReadCloser, error)

	MigrateThumbnail(ctx context.Context, req *dto.MigrateThumbnailReq) error
}
