package service

import (
	"context"
	"fmt"
	"hta-platform/internal/media/controller/dto"
	"hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/media/domain/repository"
	"strings"

	"github.com/gosimple/slug"
)

type MediaChapterServiceImpl struct {
	repo      repository.MediaChapterRepository
	mediaRepo repository.MediaRepository
}

// CreateMediaChapters implements [MediaChapterService].
func (m *MediaChapterServiceImpl) CreateMediaChapters(ctx context.Context, req *dto.CreateMediaChapterReq) (*dto.MediaChapterRes, error) {
	// Find media by url
	media, err := m.mediaRepo.GetMediaByUrl(ctx, req.MediaUrl)
	if err != nil {
		return nil, fmt.Errorf("media not found: %w", err)
	}

	chapters := make([]*entity.MediaChapter, len(req.Chapters))
	for i, ch := range req.Chapters {
		chapterSlug := slug.Make(strings.ToLower(ch.Name))
		chapters[i] = &entity.MediaChapter{
			MediaID: media.ID,
			Name:    ch.Name,
			Order:   ch.Order,
			URL:     chapterSlug,
		}
	}

	createdChapters, err := m.repo.CreateMediaChapters(ctx, chapters)
	if err != nil {
		return nil, err
	}

	if len(createdChapters) > 0 {
		var res dto.MediaChapterRes
		res.SetData(createdChapters[0])
		return &res, nil
	}

	return nil, nil
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

func NewMediaChapterServiceImpl(repo repository.MediaChapterRepository, mediaRepo repository.MediaRepository) MediaChapterService {
	return &MediaChapterServiceImpl{repo: repo, mediaRepo: mediaRepo}
}
