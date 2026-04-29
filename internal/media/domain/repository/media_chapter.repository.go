package repository

import (
	"context"
	"hta-platform/internal/media/domain/model/entity"
)

type MediaChapterRepository interface {
	GetMediaChaptersByMediaUrl(ctx context.Context, url string) ([]entity.MediaChapter, error)

	GetChapterImages(ctx context.Context, mediaUrl string, chapterUrl string) ([]entity.ChapterImage, error)

	CreateMediaChapters(ctx context.Context, chapters []*entity.MediaChapter) ([]entity.MediaChapter, error)

	CreateChapterImages(ctx context.Context, images []*entity.ChapterImage) ([]entity.ChapterImage, error)

	FindChapterByMediaUrlAndOrder(ctx context.Context, mediaUrl string, order int64) (*entity.MediaChapter, error)
}
