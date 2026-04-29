package service

import (
	"context"
	"hta-platform/internal/image/controller/dto"
	"hta-platform/internal/image/domain/repository"
	"io"
)

type imageService struct {
	imgRepo repository.ImageRepository
}

// StreamImage implements [ImageService].
func (i *imageService) StreamImage(req *dto.StreamImageReq) (io.ReadCloser, error) {
	return i.imgRepo.StreamImage(context.Background(), req.URL, req.Source)
}

func NewImageService(imgRepo repository.ImageRepository) ImageService {
	return &imageService{imgRepo: imgRepo}
}
