package service

import (
	"context"
	"fmt"
	"hta-platform/internal/media/controller/dto"
	"hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/media/domain/repository"
	"strings"

	imageEntity "hta-platform/internal/image/domain/model/entity"

	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MediaChapterServiceImpl struct {
	repo      repository.MediaChapterRepository
	mediaRepo repository.MediaRepository
	db        *gorm.DB
}

// GetChapterImages implements [MediaChapterService].
func (m *MediaChapterServiceImpl) GetChapterImages(ctx context.Context, mediaUrl string, chapterUrl string) ([]dto.ChapterImageRes, error) {
	images, err := m.repo.GetChapterImages(ctx, mediaUrl, chapterUrl)
	if err != nil {
		return nil, err
	}

	res := make([]dto.ChapterImageRes, len(images))
	for i, img := range images {
		res[i].SetData(img)
	}

	return res, nil
}

// CreateChapterImages implements [MediaChapterService].
func (m *MediaChapterServiceImpl) CreateChapterImages(ctx context.Context, req *dto.CreateChapterImageReq) (*dto.ChapterImageRes, error) {
	// 1. Find the MediaChapter first by find MediaUrl -> ChapterOrder
	chapter, err := m.repo.FindChapterByMediaUrlAndOrder(ctx, req.MediaUrl, req.ChapterOrder)
	if err != nil {
		return nil, fmt.Errorf("media chapter not found: %w", err)
	}

	txErr := m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 2. Prepare ChapterImage entities
		chapterImages := make([]*entity.ChapterImage, len(req.ChapterImages))
		for i, ciReq := range req.ChapterImages {
			chapterImages[i] = &entity.ChapterImage{
				ChapterID: chapter.ID,
				Order:     ciReq.Order,
			}
		}

		// 3. Insert batch ChapterImage slice first (Upsert avoid duplicate)
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "order"}, {Name: "chapter_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
		}).Create(&chapterImages).Error; err != nil {
			return err
		}

		// 4. Prepare Image entities
		var allImages []*imageEntity.Image
		for i, ciReq := range req.ChapterImages {
			createdCI := chapterImages[i]
			for _, imgReq := range ciReq.Images {
				allImages = append(allImages, &imageEntity.Image{
					ResourceID:  createdCI.ID,
					URL:         imgReq.Url,
					Description: imgReq.Description,
					Source:      imgReq.Source,
				})
			}
		}

		// 5. Batch Insert Image records with OnConflict (idempotency on url)
		if len(allImages) > 0 {
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "url"}},
				DoUpdates: clause.AssignmentColumns([]string{"description", "resource_id", "source", "updated_at"}),
			}).Create(&allImages).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if txErr != nil {
		return nil, txErr
	}

	// Return response (mapping first created if any or just empty success)
	// Interface expects *dto.ChapterImageRes
	return nil, nil
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
			Source:  ch.Source,
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

func NewMediaChapterServiceImpl(repo repository.MediaChapterRepository, mediaRepo repository.MediaRepository, db *gorm.DB) MediaChapterService {
	return &MediaChapterServiceImpl{repo: repo, mediaRepo: mediaRepo, db: db}
}
