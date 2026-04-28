package repository

import (
	"context"
	"hta-platform/internal/media/domain/model/entity"
)

type MediaChapterRepository interface {
	GetMediaChaptersByMediaUrl(ctx context.Context, url string) ([]entity.MediaChapter, error)

	GetChapterImagesByChapterUrl(ctx context.Context, url string) ([]entity.ChapterImage, error)
}
