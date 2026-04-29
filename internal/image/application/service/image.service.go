package service

import (
	"hta-platform/internal/image/controller/dto"
	"io"
)

type ImageService interface {
	StreamImage(req *dto.StreamImageReq) (io.ReadCloser, error)
}
