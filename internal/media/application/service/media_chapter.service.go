package service

import (
	"context"
	"hta-platform/internal/media/controller/dto"
)

type MediaChapterService interface {
	GetMediaChaptersByMediaUrl(ctx context.Context, url string) ([]dto.MediaChapterRes, error)

	GetChapterImages(ctx context.Context, mediaUrl string, chapterUrl string) ([]dto.ChapterImageRes, error)

	CreateMediaChapters(ctx context.Context, req *dto.CreateMediaChapterReq) (*dto.MediaChapterRes, error)

	CreateChapterImages(ctx context.Context, req *dto.CreateChapterImageReq) (*dto.ChapterImageRes, error)
}
