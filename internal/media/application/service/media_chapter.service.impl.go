package service

import (
	"context"
	"hta-platform/internal/media/controller/dto"
	"hta-platform/internal/media/domain/repository"
)

type MediaChapterServiceImpl struct {
	repo repository.MediaChapterRepository
}

// GetChapterImagesByChapterUrl implements [MediaChapterService].
func (m *MediaChapterServiceImpl) GetChapterImagesByChapterUrl(ctx context.Context, url string) ([]dto.ChapterImageRes, error) {
	images, err := m.repo.GetChapterImagesByChapterUrl(ctx, url)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ChapterImageRes, len(images))
	for i, img := range images {
		res[i].SetData(img)
	}

	return res, nil
}

// GetMediaChaptersByMediaUrl implements [MediaChapterService].
func (m *MediaChapterServiceImpl) GetMediaChaptersByMediaUrl(ctx context.Context, url string) ([]dto.MediaChapterRes, error) {
	chapters, err := m.repo.GetMediaChaptersByMediaUrl(ctx, url)
	if err != nil {
		return nil, err
	}

	res := make([]dto.MediaChapterRes, len(chapters))
	for i, ch := range chapters {
		res[i].SetData(ch)
	}

	return res, nil
}

func NewMediaChapterServiceImpl(repo repository.MediaChapterRepository) MediaChapterService {
	return &MediaChapterServiceImpl{repo: repo}
}
