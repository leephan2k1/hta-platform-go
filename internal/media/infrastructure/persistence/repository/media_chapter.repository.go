package repository

import (
	"context"
	"hta-platform/internal/media/domain/model/entity"
	"hta-platform/internal/media/domain/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type mediaChapterRepository struct {
	db *gorm.DB
}

// CreateMediaChapters implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) CreateMediaChapters(ctx context.Context, chapters []*entity.MediaChapter) ([]entity.MediaChapter, error) {
	err := m.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "url"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "order", "source", "language", "updated_at"}),
		}).Create(&chapters).Error

	if err != nil {
		return nil, err
	}

	res := make([]entity.MediaChapter, len(chapters))
	for i, ch := range chapters {
		res[i] = *ch
	}

	return res, nil
}

// GetChapterImagesByChapterUrl implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) GetChapterImagesByChapterUrl(ctx context.Context, url string) ([]entity.ChapterImage, error) {
	var images []entity.ChapterImage
	err := m.db.WithContext(ctx).
		Joins("JOIN hta.media_chapter mc ON mc.id = hta.chapter_image.chapter_id").
		Where("mc.url = ?", url).
		Order("hta.chapter_image.order ASC").
		Find(&images).Error
	return images, err
}

// GetMediaChaptersByMediaUrl implements [repository.MediaChapterRepository].
func (m *mediaChapterRepository) GetMediaChaptersByMediaUrl(ctx context.Context, url string) ([]entity.MediaChapter, error) {
	var chapters []entity.MediaChapter
	err := m.db.WithContext(ctx).
		Joins("JOIN hta.media m ON m.id = hta.media_chapter.media_id").
		Where("m.url = ?", url).
		Order("hta.media_chapter.order DESC").
		Find(&chapters).Error

	return chapters, err
}

func NewMediaChapterRepository(db *gorm.DB) repository.MediaChapterRepository {
	return &mediaChapterRepository{db: db}
}
